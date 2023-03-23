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
	"sync"
	"time"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
	"github.com/google/uuid"
	"github.com/spf13/pflag"

	"github.com/spf13/cobra"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/logging"
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
	BackupCmd.Flags().StringVarP(&BackupMode, "dn-backup-mode", "b", "", "openGauss data backup mode (FULL|PTRACK)")
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
	proxy, err := pkg.NewShardingSphereProxy(Username, Password, pkg.DefaultDbName, Host, Port)
	if err != nil {
		return xerr.NewCliErr("create ss-proxy connect failed")
	}

	ls, err := pkg.NewLocalStorage(pkg.DefaultRootDir())
	if err != nil {
		return xerr.NewCliErr("create local storage failed")
	}

	// Step1. lock cluster
	if err := proxy.LockForBackup(); err != nil {
		return xerr.NewCliErr("lock for backup failed")
	}

	// Step2. Get cluster info and save local backup info
	lsBackup, err := exportData(proxy, ls)
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("export backup data failed, err:%s", err.Error()))
	}

	logging.Info(fmt.Sprintf("export backup data success, backup filename: %s", filename))

	// Step3. send backup command to agent-server.
	if err := execBackup(lsBackup); err != nil {
		// if backup failed, still need to unlock cluster.
		if err := proxy.Unlock(); err != nil {
			logging.Error(fmt.Sprintf("coz exec backup failed, try to unlock cluster, but still failed, err:%s", err.Error()))
		}
		return xerr.NewCliErr(fmt.Sprintf("exec backup failed, err:%s", err.Error()))
	}

	// Step4. unlock cluster
	if err := proxy.Unlock(); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("unlock cluster failed, err:%s", err.Error()))
	}

	// Step5. update backup file
	if err := ls.WriteByJSON(filename, lsBackup); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("update backup file failed, err:%s", err.Error()))
	}

	// Step6. check agent server backup status
	status := checkBackupStatus(lsBackup)
	logging.Info(fmt.Sprintf("backup result:%s", status))

	// Step7. finished backup and update backup file
	if err := ls.WriteByJSON(filename, lsBackup); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("update backup file failed, err: %s", err.Error()))
	}

	logging.Info("backup finished")
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
			ID:        uuid.New().String(), // generate uuid for this backup
			CSN:       csn,
			StartTime: time.Now().Unix(),
			EndTime:   0,
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
	var (
		wg       sync.WaitGroup
		sNodes   = lsBackup.SsBackup.StorageNodes
		dnCh     = make(chan *model.DataNode, len(sNodes))
		failSnCh = make(chan *model.StorageNode, len(sNodes))
		success  = true
	)
	logging.Info("Starting send backup command to agent server...")

	for _, node := range sNodes {
		wg.Add(1)
		go func(wg *sync.WaitGroup, node *model.StorageNode) {
			defer wg.Done()
			agentHost := node.IP
			if agentHost == "127.0.0.1" {
				agentHost = Host
			}
			as := pkg.NewAgentServer(fmt.Sprintf("%s:%d", agentHost, AgentPort))
			_execBackup(as, node, dnCh, failSnCh)
		}(&wg, node)
	}

	wg.Wait()
	close(dnCh)
	close(failSnCh)

	// TODO format print data like a table
	for errN := range failSnCh {
		success = false
		logging.Error(fmt.Sprintf("failed node detail: [IP:%s, PORT:%d]", errN.IP, errN.Port))
	}

	if !success {
		lsBackup.SsBackup.Status = model.SsBackupStatusFailed
		return xerr.NewCliErr("backup failed")
	}

	// save data node list to lsBackup
	for dn := range dnCh {
		lsBackup.DnList = append(lsBackup.DnList, dn)
	}

	lsBackup.SsBackup.Status = model.SsBackupStatusRunning
	return nil
}

func _execBackup(as pkg.IAgentServer, node *model.StorageNode, dnCh chan *model.DataNode, failSnCh chan *model.StorageNode) {
	in := &model.BackupIn{
		DbPort:       node.Port,
		DbName:       node.Database,
		Username:     node.Username,
		Password:     node.Password,
		DnBackupPath: BackupPath,
		DnThreadsNum: ThreadsNum,
		DnBackupMode: model.BDBackModeFull,
		Instance:     defaultInstance,
	}
	backupID, err := as.Backup(in)
	if err != nil {
		logging.Error(fmt.Sprintf("backup failed, %s", err.Error()))
		failSnCh <- node
		return
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
}

func checkBackupStatus(lsBackup *model.LsBackup) model.BackupStatus {
	var (
		wg                sync.WaitGroup
		dataNodeMap       = make(map[string]*model.DataNode)
		backupFinalStatus = model.SsBackupStatusCompleted
		statusCh          = make(chan *model.DataNode, len(lsBackup.DnList))
	)

	// DataNode.IP -> DataNode
	for _, dn := range lsBackup.DnList {
		dataNodeMap[dn.IP] = dn
	}

	for _, sn := range lsBackup.SsBackup.StorageNodes {
		wg.Add(1)
		go func(wg *sync.WaitGroup, sn *model.StorageNode) {
			defer wg.Done()
			agentHost := sn.IP
			if agentHost == "127.0.0.1" {
				agentHost = Host
			}
			as := pkg.NewAgentServer(fmt.Sprintf("%s:%d", agentHost, AgentPort))
			dn := dataNodeMap[sn.IP]

			// check backup status
			status := checkStatus(as, sn, dn.BackupID, model.BackupStatus(""), defaultShowDetailRetryTimes)

			// update DataNode status
			dn.Status = status
			dn.EndTime = time.Now().Unix()
			statusCh <- dn
		}(&wg, sn)
	}

	wg.Wait()
	close(statusCh)

	for dn := range statusCh {
		logging.Info(fmt.Sprintf("data node backup final status: [IP:%s, backupID:%s] ==> %s", dn.IP, dn.BackupID, dn.Status))
	}

	for _, dn := range lsBackup.DnList {
		if dn.Status != model.SsBackupStatusCompleted {
			backupFinalStatus = model.SsBackupStatusFailed
		}
	}

	lsBackup.SsBackup.Status = backupFinalStatus
	return backupFinalStatus
}

func checkStatus(as pkg.IAgentServer, sn *model.StorageNode, backupId string, status model.BackupStatus, retryTimes uint8) model.BackupStatus {
	if retryTimes+1 == 0 {
		return status
	}
	if status == model.SsBackupStatusCompleted || status == model.SsBackupStatusFailed {
		return status
	}

	// todo: how often to check backup status
	time.Sleep(time.Second * 2)

	in := &model.ShowDetailIn{
		DbPort:       sn.Port,
		DbName:       sn.Database,
		Username:     sn.Username,
		Password:     sn.Password,
		DnBackupId:   backupId,
		DnBackupPath: BackupPath,
		Instance:     defaultInstance,
	}
	backupInfo, err := as.ShowDetail(in)
	if err != nil {
		logging.Error(fmt.Sprintf("get storage node [IP:%s] backup detail from agent server failed, will retry %d times.\n%s", sn.IP, retryTimes, err.Error()))
		return checkStatus(as, sn, backupId, model.SsBackupStatusCheckError, retryTimes-1)
	}
	return checkStatus(as, sn, backupId, backupInfo.Status, retryTimes)
}
