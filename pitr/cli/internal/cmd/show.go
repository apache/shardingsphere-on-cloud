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
	"encoding/json"
	"fmt"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/logging"
	"github.com/spf13/cobra"
	"os"
)

var (
	CSN      string
	RecordID string
)

var Show = &cobra.Command{
	Use:   "show",
	Short: "Show backup history",
	Run: func(cmd *cobra.Command, args []string) {
		host, err := cmd.Flags().GetString(host)
		if err != nil {
			logging.Error(err.Error())
		}
		logging.Info(fmt.Sprintf("flags:host:%s", host))

		port, err := cmd.Flags().GetUint16(port)
		if err != nil {
			logging.Error(err.Error())
		}
		logging.Info(fmt.Sprintf("flags:port:%d", port))

		un, err := cmd.Flags().GetString(username)
		if err != nil {
			logging.Error(err.Error())
		}
		logging.Info(fmt.Sprintf("flags:username:%s", un))

		pw, err := cmd.Flags().GetString(password)
		if err != nil {
			logging.Error(err.Error())
		}
		logging.Info(fmt.Sprintf("flags:password:%s", pw))

		agentPort, err := cmd.Flags().GetUint16(agentPort)
		if err != nil {
			logging.Error(err.Error())
		}
		logging.Info(fmt.Sprintf("flags:agentPort:%d", agentPort))

		CSN, err = cmd.Flags().GetString(csn)
		if err != nil {
			logging.Error(err.Error())
		}
		logging.Info(fmt.Sprintf("flags:csn:%s", CSN))

		RecordID, err = cmd.Flags().GetString(backupRecordID)
		if err != nil {
			logging.Error(err.Error())
		}
		logging.Info(fmt.Sprintf("flags:id:%s", RecordID))

		logging.Info("Show backup history ...")

		if err := show(); err != nil {
			logging.Error(err.Error())
		}
	},
}

func show() error {
	root := fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".gs_pitr")
	logging.Info(fmt.Sprintf("Default backup path: %s\n", root))

	ls, err := pkg.NewLocalStorage(root)
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("connect to local storage failed, err:%s", err.Error()))
	}
	// show backup record by csn
	if CSN != "" {
		bak, err := ls.ReadByCSN(csn)
		if bak == nil {
			fmt.Println("Didn't find backup record by csn: ", CSN)
			return nil
		}
		if err != nil {
			return xerr.NewCliErr(fmt.Sprintf("read backup record failed, err:%s", err.Error()))
		}
		if err := formatRecord([]model.LsBackup{*bak}); err != nil {
			return err
		}
		return nil
	}
	// show backup record by id
	if RecordID != "" {
		bak, err := ls.ReadByID(RecordID)
		if bak == nil {
			fmt.Println("Didn't find backup record by record id: ", RecordID)
			return nil
		}
		if err != nil {
			return xerr.NewCliErr(fmt.Sprintf("read backup record failed, err:%s", err.Error()))
		}
		if err := formatRecord([]model.LsBackup{*bak}); err != nil {
			return err
		}
		return nil
	}

	// show all backup record
	backupList, err := ls.ReadAll()
	if len(backupList) == 0 {
		fmt.Println("Didn't find any backup record.")
		return nil
	}
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("read backup record failed, err:%s", err.Error()))
	}

	if err := formatRecord(backupList); err != nil {
		return err
	}

	return nil
}
func formatRecord(backups []model.LsBackup, mode ...string) error {
	var m string

	if len(mode) == 0 {
		m = "json"
	} else {
		m = mode[0]
	}
	switch m {
	case "json":
		return formatRecordJson(backups)
	case "table":
		// TODO format record table
		return nil
	default:
		return formatRecordJson(backups)
	}
}

func formatRecordJson(backups []model.LsBackup) error {
	var ds []string
	for _, backup := range backups {
		data, err := json.MarshalIndent(backup, "", "\t")
		if err != nil {
			logging.Error(err.Error())
			return xerr.NewCliErr("format backup record failed, indent failed.")
		}
		ds = append(ds, string(data))
	}
	fmt.Println(ds)
	return nil
}

func init() {
	Show.PersistentFlags().StringVarP(&CSN, "csn", "", "", "commit sequence number")
	Show.PersistentFlags().StringVarP(&RecordID, "id", "", "", "backup record id")
}
