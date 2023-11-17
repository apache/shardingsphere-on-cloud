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
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

//nolint:dupl
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a backup record",
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

		if err := deleteRecord(); err != nil {
			logging.Error(err.Error())
		}
	},
}

//nolint:dupl
func init() {
	RootCmd.AddCommand(DeleteCmd)

	DeleteCmd.Flags().StringVarP(&Host, "host", "H", "", "ss-proxy hostname or ip")
	_ = DeleteCmd.MarkFlagRequired("host")
	DeleteCmd.Flags().Uint16VarP(&Port, "port", "P", 0, "ss-proxy port")
	_ = DeleteCmd.MarkFlagRequired("port")
	DeleteCmd.Flags().StringVarP(&Username, "username", "u", "", "ss-proxy username")
	_ = DeleteCmd.MarkFlagRequired("username")
	DeleteCmd.Flags().StringVarP(&Password, "password", "p", "", "ss-proxy password")
	_ = DeleteCmd.MarkFlagRequired("password")
	DeleteCmd.Flags().StringVarP(&BackupPath, "dn-backup-path", "B", "", "openGauss data backup path")
	_ = DeleteCmd.MarkFlagRequired("dn-backup-path")
	DeleteCmd.Flags().Uint16VarP(&AgentPort, "agent-port", "a", 443, "agent server port")
	_ = DeleteCmd.MarkFlagRequired("agent-port")

	DeleteCmd.Flags().StringVarP(&CSN, "csn", "", "", "commit sequence number")
	DeleteCmd.Flags().StringVarP(&RecordID, "id", "", "", "backup record id")
}

//nolint:dupl
func deleteRecord() error {
	// init local storage
	ls, err := pkg.NewLocalStorage(pkg.DefaultRootDir())
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("new local storage failed. err: %s", err.Error()))
	}

	// get backup record
	var baks []*model.LsBackup
	baks, err = validate(ls, CSN, RecordID)
	if err != nil {
		return err
	}

	bak := baks[0]
	// check agent server status
	logging.Info("Checking agent server status...")
	if available := checkAgentServerStatus(bak); !available {
		return xerr.NewCliErr("one or more agent server are not available.")
	}

	prompt := fmt.Sprintf(
		"The backup record(ID: %s, CSN: %s) will be deleted forever.\n"+
			"Are you sure to continue? (Y/N)", bak.Info.ID, bak.Info.CSN)
	err = getUserApproveInTerminal(prompt)
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("%s", err))
	}

	// mark the target backup record to be deleted
	// meanwhile this record cannot be restored
	if err := ls.HideByName(bak.Info.FileName); err != nil {
		return xerr.NewCliErr("cannot mark backup record.")
	}

	// exec delete
	logging.Info("Start delete backup data to openGauss...")
	if err := _execDelete(bak); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("exec delete failed. err: %s", err))
	}
	logging.Info("Delete backup data success!")

	// delete the backup record
	if err := ls.DeleteByHidedName(bak.Info.FileName); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("exec delete backup record failed. err: %s", err))
	}

	logging.Info("Delete success!")
	return nil
}

func _execDelete(lsBackup *model.LsBackup) error {
	var (
		dataNodeMap       = make(map[string]*model.DataNode)
		totalNum          = len(lsBackup.SsBackup.StorageNodes)
		resultCh          = make(chan *model.DeleteBackupResult, totalNum)
		dnResult          = make([]*model.DeleteBackupResult, 0)
		deleteFinalStatus = "Completed"
	)
	for _, dn := range lsBackup.DnList {
		dataNodeMap[dn.IP] = dn
	}

	if totalNum == 0 {
		logging.Info("No data node need to delete backup files")
		return nil
	}

	pw := prettyoutput.NewPW(totalNum)
	go pw.Render()

	for _, storagenode := range lsBackup.SsBackup.StorageNodes {
		sn := storagenode
		if dn, ok := dataNodeMap[sn.IP]; !ok {
			logging.Warn(fmt.Sprintf("SKIPPED! data node %s:%d not found in backup info.", sn.IP, sn.Port))
			continue
		} else {
			as := pkg.NewAgentServer(fmt.Sprintf("%s:%d", convertLocalhost(sn.IP), AgentPort))
			go doDelete(as, sn, dn, resultCh, pw)
		}
	}

	time.Sleep(time.Millisecond * 100)

	for pw.IsRenderInProgress() {
		time.Sleep(time.Millisecond * 100)
	}

	close(resultCh)

	for result := range resultCh {
		dnResult = append(dnResult, result)
		if result.Status != "Completed" {
			deleteFinalStatus = "Failed"
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Delete Backup Files Result")
	t.AppendHeader(table.Row{"#", "Node IP", "Node Port", "Result", "Message"})
	t.SetColumnConfigs([]table.ColumnConfig{{Number: 5, WidthMax: 50}})

	for i, result := range dnResult {
		t.AppendRow([]interface{}{i + 1, result.IP, result.Port, result.Status, result.Msg})
		t.AppendSeparator()
	}

	t.Render()

	if deleteFinalStatus == "Failed" {
		return xerr.NewCliErr("delete failed, please check the log for more details.")
	}

	return nil
}
