/*
* Licensed to the Apache Software Foundation (ASF) under one or more
* contributor license agreements.  See the NOTICE file distributed with
* this work for additional information regarding copyright ownership.
* The ASF licenses this file to You under the Apache License, Version 2.0
* (the "License"); you may not use this file except in compliance with
* the License.  You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/logging"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/prettyoutput"
	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/sync/errgroup"
)

const (
	// defaultInstance is used to set backup instance name in openGauss, we can modify it in the future.
	defaultInstance = "ins-default-ss"
	// defaultShowDetailRetryTimes retry times of check backup detail from agent server
	defaultShowDetailRetryTimes = 3
)

var filename string

var BackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup a database cluster",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			fmt.Printf("Flag: %s Value: %s\n", flag.Name, flag.Value)
		})

		// convert BackupModeStr to BackupMode
		switch BackupModeStr {
		case "FULL", "full":
			BackupMode = model.BDBackModeFull
		case "PTRACK", "ptrack":
			BackupMode = model.DBBackModePTrack
		}
		if BackupMode == model.DBBackModePTrack {
			logging.Warn("Please make sure all openGauss nodes have been set correct configuration about ptrack. You can refer to https://support.huaweicloud.com/intl/zh-cn/devg-opengauss/opengauss_devg_1362.html for more details.")
		}

		logging.Info(fmt.Sprintf("Default backup path: %s", pkg.DefaultRootDir()))

		// Start backup
		if err := backup(); err != nil {
			logging.Error(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(BackupCmd)

	BackupCmd.Flags().StringVarP(&Host, "host", "H", "", "ss-proxy hostname or ip")
	_ = BackupCmd.MarkFlagRequired("host")
	BackupCmd.Flags().Uint16VarP(&Port, "port", "P", 0, "ss-proxy port")
	_ = BackupCmd.MarkFlagRequired("port")
	BackupCmd.Flags().StringVarP(&Username, "username", "u", "", "ss-proxy username")
	_ = BackupCmd.MarkFlagRequired("username")
	BackupCmd.Flags().StringVarP(&Password, "password", "p", "", "ss-proxy password")
	_ = BackupCmd.MarkFlagRequired("password")
	BackupCmd.Flags().StringVarP(&BackupPath, "dn-backup-path", "B", "", "openGauss data backup path")
	_ = BackupCmd.MarkFlagRequired("dn-backup-path")
	BackupCmd.Flags().StringVarP(&BackupModeStr, "dn-backup-mode", "b", "", "openGauss data backup mode (FULL|PTRACK)")
	_ = BackupCmd.MarkFlagRequired("dn-backup-mode")
	BackupCmd.Flags().Uint8VarP(&ThreadsNum, "dn-threads-num", "j", 1, "openGauss data backup threads nums")
	BackupCmd.Flags().Uint16VarP(&AgentPort, "agent-port", "a", 443, "agent server port")
	_ = BackupCmd.MarkFlagRequired("agent-port")

}

// Steps of backup:
// 1. lock cluster
// 2. Get cluster info and save local backup info
// 3. Operate backup by agent-server
// 4. unlock cluster
// 5. Waiting for backups finished
// 6. Update local backup info
// 7. Double check backups all finished
func backup() error {
	var err error
	proxy, err := pkg.NewShardingSphereProxy(Username, Password, pkg.DefaultDBName, Host, Port)
	if err != nil {
		return xerr.NewCliErr("create ss-proxy connect failed")
	}

	ls, err := pkg.NewLocalStorage(pkg.DefaultRootDir())
	if err != nil {
		return xerr.NewCliErr("create local storage failed")
	}

	defer func() {
		if err != nil {
			logging.Info("try to unlock cluster ...")
			if err := proxy.Unlock(); err != nil {
				logging.Error(fmt.Sprintf("coz backup failed, try to unlock cluster, but still failed, err:%s", err.Error()))
			}
		}
	}()

	// Step1. lock cluster
	logging.Info("Starting lock cluster ...")
	err = proxy.LockForBackup()
	if err != nil {
		return xerr.NewCliErr("lock for backup failed")
	}

	// Step2. Get cluster info and save local backup info
	logging.Info("Starting export metadata ...")
	lsBackup, err := exportData(proxy, ls)
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("export backup data failed, err:%s", err.Error()))
	}
	logging.Info(fmt.Sprintf("Export backup data success, backup filename: %s", filename))

	// Step3. Check agent server status
	logging.Info("Checking agent server status...")
	if available := checkAgentServerStatus(lsBackup); !available {
		err = xerr.NewCliErr("one or more agent server are not available.")
		return err
	}

	// Step4. Show disk space
	logging.Info("Checking disk space...")
	err = checkDiskSpace(lsBackup)
	if err != nil {
		return xerr.NewCliErr(err.Error())
	}

	// Step5. send backup command to agent-server.
	logging.Info("Starting backup ...")
	err = execBackup(lsBackup)
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("exec backup failed, err:%s", err.Error()))
	}

	// Step6. unlock cluster
	logging.Info("Starting unlock cluster ...")
	err = proxy.Unlock()
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("unlock cluster failed, err:%s", err.Error()))
	}

	// Step7. update backup file
	logging.Info("Starting update backup file ...")
	err = ls.WriteByJSON(filename, lsBackup)
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("update backup file failed, err:%s", err.Error()))
	}

	// Step8. check agent server backup
	logging.Info("Starting check backup status ...")
	status := checkBackupStatus(lsBackup)
	logging.Info(fmt.Sprintf("Backup result: %s", status))

	// Step9. finished backup and update backup file
	logging.Info("Starting update backup file ...")
	err = ls.WriteByJSON(filename, lsBackup)
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("update backup file failed, err: %s", err.Error()))
	}

	logging.Info("Backup finished!")
	return nil
}

func exportData(proxy pkg.IShardingSphereProxy, ls pkg.ILocalStorage) (lsBackup *model.LsBackup, err error) {
	// Step1. export cluster metadata from ss-proxy
	cluster, err := proxy.ExportMetaData()
	if err != nil {
		return nil, xerr.NewCliErr("export meta data failed")
	}

	// Step2. export storage nodes from ss-proxy
	nodes, err := proxy.ExportStorageNodes()
	if err != nil {
		return nil, xerr.NewCliErr("export storage nodes failed")
	}

	// Step3. combine the backup contents
	filename = ls.GenFilename(pkg.ExtnJSON)
	csn := ""
	if cluster.SnapshotInfo != nil {
		csn = cluster.SnapshotInfo.Csn
	}

	contents := &model.LsBackup{
		Info: &model.BackupMetaInfo{
			ID:         uuid.New().String(), // generate uuid for this backup
			CSN:        csn,
			StartTime:  time.Now().Unix(),
			EndTime:    0,
			BackupMode: BackupMode,
		},
		SsBackup: &model.SsBackup{
			Status:       model.SsBackupStatusWaiting, // default status of backup is model.SsBackupStatusWaiting
			ClusterInfo:  cluster,
			StorageNodes: nodes,
		},
	}

	// Step4. finally, save data with json to local
	if err := ls.WriteByJSON(filename, contents); err != nil {
		return nil, xerr.NewCliErr("write backup info by json failed")
	}

	return contents, nil
}

func execBackup(lsBackup *model.LsBackup) error {
	sNodes := lsBackup.SsBackup.StorageNodes
	dnCh := make(chan *model.DataNode, len(sNodes))
	g := new(errgroup.Group)

	logging.Info("Starting send backup command to agent server...")

	for _, node := range sNodes {
		sn := node
		as := pkg.NewAgentServer(fmt.Sprintf("%s:%d", convertLocalhost(sn.IP), AgentPort))
		g.Go(func() error {
			return _execBackup(as, sn, dnCh)
		})
	}

	err := g.Wait()
	close(dnCh)

	// if backup failed, return error
	if err != nil {
		lsBackup.SsBackup.Status = model.SsBackupStatusFailed
		return xerr.NewCliErr(fmt.Sprintf("node backup failed, err:%s", err.Error()))
	}

	// save data node list to lsBackup
	for dn := range dnCh {
		lsBackup.DnList = append(lsBackup.DnList, dn)
	}

	lsBackup.SsBackup.Status = model.SsBackupStatusRunning
	return nil
}

func _execBackup(as pkg.IAgentServer, node *model.StorageNode, dnCh chan *model.DataNode) error {
	in := &model.BackupIn{
		DBPort:       node.Port,
		DBName:       node.Database,
		Username:     node.Username,
		Password:     node.Password,
		DnBackupPath: BackupPath,
		DnThreadsNum: ThreadsNum,
		DnBackupMode: BackupMode,
		Instance:     defaultInstance,
	}
	backupID, err := as.Backup(in)
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("backup failed, err:%s", err.Error()))
	}

	// update DnList of lsBackup
	dn := &model.DataNode{
		IP:        node.IP,
		Port:      node.Port,
		Status:    model.SsBackupStatusRunning,
		BackupID:  backupID,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	dnCh <- dn
	return nil
}

func checkBackupStatus(lsBackup *model.LsBackup) model.BackupStatus {
	var (
		dataNodeMap       = make(map[string]*model.DataNode)
		dnCh              = make(chan *model.DataNode, len(lsBackup.DnList))
		backupFinalStatus = model.SsBackupStatusCompleted
		totalNum          = len(lsBackup.SsBackup.StorageNodes)
		dnResult          = make([]*model.DataNode, 0)
	)

	for _, dn := range lsBackup.DnList {
		dataNodeMap[dn.IP] = dn
	}

	pw := prettyoutput.NewPW(totalNum)
	go pw.Render()
	for idx := 0; idx < totalNum; idx++ {
		sn := lsBackup.SsBackup.StorageNodes[idx]
		as := pkg.NewAgentServer(fmt.Sprintf("%s:%d", convertLocalhost(sn.IP), AgentPort))
		dn := dataNodeMap[sn.IP]
		go checkStatus(as, sn, dn, dnCh, pw)
	}

	// wait for all data node backup finished
	time.Sleep(time.Millisecond * 100)
	for pw.IsRenderInProgress() {
		if pw.LengthActive() == 0 {
			pw.Stop()
		}
		time.Sleep(time.Millisecond * 100)
	}

	close(dnCh)

	for dn := range dnCh {
		dnResult = append(dnResult, dn)
		if dn.Status != model.SsBackupStatusCompleted {
			backupFinalStatus = model.SsBackupStatusFailed
		}
	}

	// print backup result formatted
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Backup Task Result: %s", backupFinalStatus)
	t.AppendHeader(table.Row{"#", "Data Node IP", "Data Node Port", "Result"})

	for i, dn := range dnResult {
		t.AppendRow([]interface{}{i + 1, dn.IP, dn.Port, dn.Status})
		t.AppendSeparator()
	}

	t.Render()

	lsBackup.DnList = dnResult
	lsBackup.SsBackup.Status = backupFinalStatus
	lsBackup.Info.EndTime = time.Now().Unix()
	return backupFinalStatus
}

func checkStatus(as pkg.IAgentServer, sn *model.StorageNode, dn *model.DataNode, dnCh chan *model.DataNode, pw progress.Writer) {
	var (
		// mark check status is done, time ticker should break.
		done = make(chan bool)
		// time ticker, try to doCheck request every 2 seconds.
		ticker = time.Tick(time.Second * 2)
		// progress bar.
		tracker = progress.Tracker{Message: fmt.Sprintf("Checking backup status  # %s:%d", sn.IP, AgentPort), Total: 0, Units: progress.UnitsDefault}
	)

	pw.AppendTracker(&tracker)

	for !tracker.IsDone() {
		select {
		case <-done:
			return
		case <-ticker:
			status, err := doCheck(as, sn, dn.BackupID, defaultShowDetailRetryTimes)
			if err != nil {
				tracker.MarkAsErrored()
				dn.Status = status
				dn.EndTime = time.Now().Unix()
				dnCh <- dn
				done <- true
			}
			if status == model.SsBackupStatusCompleted || status == model.SsBackupStatusFailed {
				tracker.MarkAsDone()
				dn.Status = status
				dn.EndTime = time.Now().Unix()
				dnCh <- dn
				done <- true
			}
		}
	}
}

func doCheck(as pkg.IAgentServer, sn *model.StorageNode, backupID string, retries int) (status model.BackupStatus, err error) {
	in := &model.ShowDetailIn{
		DBPort:       sn.Port,
		DBName:       sn.Database,
		Username:     sn.Username,
		Password:     sn.Password,
		DnBackupID:   backupID,
		DnBackupPath: BackupPath,
		Instance:     defaultInstance,
	}
	backupInfo, err := as.ShowDetail(in)
	if err != nil {
		if retries == 0 {
			return model.SsBackupStatusCheckError, err
		}
		time.Sleep(time.Second * 1)
		return doCheck(as, sn, backupID, retries-1)
	}

	return backupInfo.Status, nil
}
