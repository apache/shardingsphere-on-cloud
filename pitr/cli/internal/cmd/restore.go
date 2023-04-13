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
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
	"github.com/spf13/pflag"

	"github.com/spf13/cobra"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/logging"
)

var (
	// database names exist in ss proxy and backup, will to be dropped
	databaseNamesExist []string
)

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
		return xerr.NewCliErr(fmt.Sprintf("new local storage failed, err:%s", err.Error()))
	}
	proxy, err := pkg.NewShardingSphereProxy(Username, Password, pkg.DefaultDBName, Host, Port)
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("new ss-proxy failed, err:%s", err.Error()))
	}

	// get backup record
	var bak *model.LsBackup
	if CSN != "" {
		bak, err = ls.ReadByCSN(CSN)
		if err != nil {
			return xerr.NewCliErr("read backup record by csn failed")
		}
	}

	if RecordID != "" {
		bak, err = ls.ReadByID(RecordID)
		if err != nil {
			return xerr.NewCliErr("read backup record by id failed")
		}
	}
	if bak == nil {
		return xerr.NewCliErr("backup record not found")
	}

	// check if the backup logic database exits,
	// if exits, we need to warning user that we will drop the database.
	if err := checkDatabaseExist(proxy, bak); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("check database exist failed:%s", err.Error()))
	}

	// exec restore
	logging.Info("Start restore backup data to openGauss...")
	if err := execRestore(bak); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("exec restore failed:%s", err.Error()))
	}

	logging.Info("Restore backup data to openGauss success!")
	// restore metadata to ss-proxy
	if err := restoreDataToSSProxy(proxy, bak); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("restore metadata to ss-proxy failed:%s", err.Error()))
	}
	logging.Info("Restore success!")
	return nil
}

func checkDatabaseExist(proxy pkg.IShardingSphereProxy, bak *model.LsBackup) error {
	clusterNow, err := proxy.ExportMetaData()
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("get cluster metadata failed:%s", err.Error()))
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
	return getUserApproveInTerminal()
}

func getUserApproveInTerminal() error {
	fmt.Printf("Detected that the database [%s] already exists in shardingsphere-proxy metadata.\nThe logic database will be DROPPED and then insert backup's metadata into shardingsphere-proxy after restoring the backup data.\nPLEASE MAKE SURE OF THIS ACTION, CONTINUE? (Y|N)\n", strings.Join(databaseNamesExist, ","))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("read user input failed:%s", err.Error()))
	}
	if scanner.Text() != "Y" && scanner.Text() != "y" && scanner.Text() != "yes" && scanner.Text() != "YES" && scanner.Text() != "Yes" {
		return xerr.NewCliErr("User abort")
	}
	return nil
}

func execRestore(lsBackup *model.LsBackup) error {
	var (
		wg           sync.WaitGroup
		storageNodes = lsBackup.SsBackup.StorageNodes
		dataNodeMap  = make(map[string]*model.DataNode)
		failedCh     = make(chan error, len(storageNodes))
	)

	for _, dataNode := range lsBackup.DnList {
		dataNodeMap[dataNode.IP] = dataNode
	}

	for _, storageNode := range storageNodes {
		wg.Add(1)
		storageNode := storageNode
		agentHost := storageNode.IP
		if agentHost == "127.0.0.1" {
			agentHost = Host
		}
		as := pkg.NewAgentServer(fmt.Sprintf("%s:%d", agentHost, AgentPort))
		dataNode, ok := dataNodeMap[storageNode.IP]
		if !ok {
			return xerr.NewCliErr(fmt.Sprintf("data node not found:%s", storageNode.IP))
		}
		go func() {
			defer wg.Done()
			_execRestore(as, storageNode, dataNode.BackupID, failedCh)
		}()
	}
	wg.Wait()
	close(failedCh)
	if len(failedCh) > 0 {
		var errMsg string
		for err := range failedCh {
			errMsg += err.Error() + "\n"
		}
		return xerr.NewCliErr(errMsg)
	}
	return nil
}

func _execRestore(as pkg.IAgentServer, node *model.StorageNode, backupID string, failedCh chan error) {
	in := &model.RestoreIn{
		DBPort:       node.Port,
		DBName:       node.Database,
		Username:     node.Username,
		Password:     node.Password,
		Instance:     defaultInstance,
		DnBackupPath: BackupPath,
		DnBackupID:   backupID,
	}

	if err := as.Restore(in); err != nil {
		failedCh <- xerr.NewCliErr(fmt.Sprintf("restore node:[IP:%s] failed:%s", node.IP, err.Error()))
	}
}

func restoreDataToSSProxy(proxy pkg.IShardingSphereProxy, lsBackup *model.LsBackup) error {
	// drop database if exists
	for _, shardingDBName := range databaseNamesExist {
		logging.Info(fmt.Sprintf("Dropping database: [%s] ...", shardingDBName))
		if err := proxy.DropDatabase(shardingDBName); err != nil {
			return xerr.NewCliErr(fmt.Sprintf("drop database failed:%s", err.Error()))
		}
	}

	// import metadata
	if err := proxy.ImportMetaData(lsBackup.SsBackup.ClusterInfo); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("Import metadata to ss-proxy failed:%s", err.Error()))
	}

	return nil
}
