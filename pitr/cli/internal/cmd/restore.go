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
	"strings"
	"time"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/logging"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/prettyoutput"

	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	// database names exist in ss proxy and backup, will to be dropped
	databaseNamesExist []string
)

//nolint:dupl
var RestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a database cluster ",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			fmt.Printf("Flag: %s Value: %s\n", flag.Name, flag.Value)
		})

		if CSN == "" && RecordID == "" {
			logging.Error("Please specify csn or record id")
			return
		}

		if CSN != "" && RecordID != "" {
			logging.Error("Please specify only one of csn and record id")
			return
		}

		if err := restore(); err != nil {
			logging.Error(err.Error())
		}
	},
}

//nolint:dupl
func init() {
	RootCmd.AddCommand(RestoreCmd)

	RestoreCmd.Flags().StringVarP(&Host, "host", "H", "", "ss-proxy hostname or ip")
	_ = RestoreCmd.MarkFlagRequired("host")
	RestoreCmd.Flags().Uint16VarP(&Port, "port", "P", 0, "ss-proxy port")
	_ = RestoreCmd.MarkFlagRequired("port")
	RestoreCmd.Flags().StringVarP(&Username, "username", "u", "", "ss-proxy username")
	_ = RestoreCmd.MarkFlagRequired("username")
	RestoreCmd.Flags().StringVarP(&Password, "password", "p", "", "ss-proxy password")
	_ = RestoreCmd.MarkFlagRequired("password")
	RestoreCmd.Flags().StringVarP(&BackupPath, "dn-backup-path", "B", "", "openGauss data backup path")
	_ = RestoreCmd.MarkFlagRequired("dn-backup-path")
	RestoreCmd.Flags().Uint16VarP(&AgentPort, "agent-port", "a", 443, "agent server port")
	_ = RestoreCmd.MarkFlagRequired("agent-port")

	RestoreCmd.Flags().StringVarP(&CSN, "csn", "", "", "commit sequence number")
	RestoreCmd.Flags().StringVarP(&RecordID, "id", "", "", "backup record id")
}

func restore() error {
	// init local storage
	ls, err := pkg.NewLocalStorage(pkg.DefaultRootDir())
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("new local storage failed. err: %s", err.Error()))
	}
	proxy, err := pkg.NewShardingSphereProxy(Username, Password, pkg.DefaultDBName, Host, Port)
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("new ss-proxy failed. err: [%s", err.Error()))
	}

	// get backup record
	var bak *model.LsBackup
	if CSN != "" {
		bak, err = ls.ReadByCSN(CSN)
		if err != nil {
			return xerr.NewCliErr(fmt.Sprintf("read backup record by csn failed. err: %s", err))
		}
	}

	if RecordID != "" {
		bak, err = ls.ReadByID(RecordID)
		if err != nil {
			return xerr.NewCliErr(fmt.Sprintf("read backup record by id failed. err: %s", err))
		}
	}
	if bak == nil {
		return xerr.NewCliErr(fmt.Sprintf("backup record not found. err: %s", err))
	}

	// check if the backup logic database exits,
	// if exits, we need to warning user that we will drop the database.
	if err := checkDatabaseExist(proxy, bak); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("check database exist failed. err: %s", err))
	}

	// check agent server status
	logging.Info("Checking agent server status...")
	if available := checkAgentServerStatus(bak); !available {
		return xerr.NewCliErr("one or more agent server are not available.")
	}

	// exec restore
	logging.Info("Start restore backup data to openGauss...")
	if err := execRestore(bak); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("exec restore failed. err: %s", err))
	}

	logging.Info("Restore backup data to openGauss success!")
	// restore metadata to ss-proxy
	if err := restoreDataToSSProxy(proxy, bak); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("restore metadata to shardingsphere proxy failed. err: %s", err))
	}
	logging.Info("Restore success!")
	return nil
}

func checkDatabaseExist(proxy pkg.IShardingSphereProxy, bak *model.LsBackup) error {
	clusterNow, err := proxy.ExportMetaData()
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("get cluster metadata failed. err: %s", err))
	}

	for k := range bak.SsBackup.ClusterInfo.MetaData.Databases {
		if _, ok := clusterNow.MetaData.Databases[k]; ok {
			databaseNamesExist = append(databaseNamesExist, k)
		}
	}

	if len(databaseNamesExist) == 0 {
		return nil
	}

	// get user input to confirm
	prompt := fmt.Sprintf(
		"Detected That The Database [%s] Already Exists In ShardingSphere-Proxy Metadata.\n"+
			"The Logic Database Will Be DROPPED And Then Insert Backup's Metadata Into ShardingSphere-Proxy After Restoring The Backup Data.\n"+
			"Are you sure to continue? (Y|N)", strings.Join(databaseNamesExist, ","))
	return getUserApproveInTerminal(prompt)
}

func restoreDataToSSProxy(proxy pkg.IShardingSphereProxy, lsBackup *model.LsBackup) error {
	// drop database if exists
	for _, shardingDBName := range databaseNamesExist {
		logging.Info(fmt.Sprintf("Dropping database: [%s] ...", shardingDBName))
		if err := proxy.DropDatabase(shardingDBName); err != nil {
			return xerr.NewCliErr(fmt.Sprintf("drop database failed. err: %s", err))
		}
	}

	// import metadata
	if err := proxy.ImportMetaData(lsBackup.SsBackup.ClusterInfo); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("Import metadata to ss-proxy failed. err: %s", err))
	}

	return nil
}

func execRestore(lsBackup *model.LsBackup) error {
	var (
		totalNum           = len(lsBackup.SsBackup.StorageNodes)
		dataNodeMap        = make(map[string]*model.DataNode)
		resultCh           = make(chan *model.RestoreResult, totalNum)
		dnResult           = make([]*model.RestoreResult, 0)
		restoreFinalStatus = "Completed"
	)

	for _, dataNode := range lsBackup.DnList {
		dataNodeMap[dataNode.IP] = dataNode
	}

	if totalNum == 0 {
		return xerr.NewCliErr(fmt.Sprintf("no storage node found, please check backup record [%s].", lsBackup.Info.ID))
	}

	pw := prettyoutput.NewPW(totalNum)
	go pw.Render()
	for i := 0; i < totalNum; i++ {
		sn := lsBackup.SsBackup.StorageNodes[i]
		dn := dataNodeMap[sn.IP]
		as := pkg.NewAgentServer(fmt.Sprintf("%s:%d", convertLocalhost(sn.IP), AgentPort))
		go doRestore(as, sn, dn.BackupID, resultCh, pw)
	}

	time.Sleep(time.Millisecond * 100)
	for pw.IsRenderInProgress() {
		time.Sleep(time.Millisecond * 100)
	}

	close(resultCh)

	for result := range resultCh {
		dnResult = append(dnResult, result)
		if result.Status != "Completed" {
			restoreFinalStatus = "Failed"
		}
	}

	// print result formatted
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Restore Task Result: %s", restoreFinalStatus)
	t.AppendHeader(table.Row{"#", "Data Node IP", "Data Node Port", "Result"})

	for i, dn := range dnResult {
		t.AppendRow([]interface{}{i + 1, dn.IP, dn.Port, dn.Status})
		t.AppendSeparator()
	}

	t.Render()

	if restoreFinalStatus == "Failed" {
		return xerr.NewCliErr("restore failed, please check the log for more details.")
	}

	return nil
}

func doRestore(as pkg.IAgentServer, sn *model.StorageNode, backupID string, resultCh chan *model.RestoreResult, pw progress.Writer) {
	tracker := &progress.Tracker{Message: fmt.Sprintf("Restore data to openGauss: %s", sn.IP)}
	result := ""

	in := &model.RestoreIn{
		DBPort:       sn.Port,
		DBName:       sn.Database,
		Username:     sn.Username,
		Password:     sn.Password,
		Instance:     defaultInstance,
		DnBackupPath: BackupPath,
		DnBackupID:   backupID,
	}

	pw.AppendTracker(tracker)

	if err := as.Restore(in); err != nil {
		tracker.MarkAsErrored()
		result = "Failed"
	} else {
		tracker.MarkAsDone()
		result = "Completed"
	}

	resultCh <- &model.RestoreResult{
		IP:     sn.IP,
		Port:   sn.Port,
		Status: result,
	}
}
