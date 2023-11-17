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
)

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show backup history",
	Run: func(cmd *cobra.Command, args []string) {
		if CSN != "" && RecordID != "" {
			logging.Error("Please specify only one of csn and record id")
			return
		}

		if err := show(); err != nil {
			logging.Error(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(ShowCmd)
	ShowCmd.Flags().StringVarP(&CSN, "csn", "", "", "commit sequence number")
	ShowCmd.Flags().StringVarP(&RecordID, "id", "", "", "backup record id")
}

func show() error {
	ls, err := pkg.NewLocalStorage(pkg.DefaultRootDir())
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("connect to local storage failed. err: %s", err))
	}

	// show backup record by csn
	if CSN != "" {
		baks, err := ls.ReadAllByCSN(CSN)
		if err != nil {
			return xerr.NewCliErr(fmt.Sprintf("read backup record failed. err: %s", err))
		}
		if len(baks) == 0 {
			fmt.Printf("Didn't find backup record by csn: %s\n", CSN)
			return nil
		}

		if err := formatRecord(baks); err != nil {
			return err
		}
		return nil
	}

	// show backup record by id
	if RecordID != "" {
		bak, err := ls.ReadByID(RecordID)
		if err != nil {
			return xerr.NewCliErr(fmt.Sprintf("read backup record failed. err: %s", err))
		}
		if bak == nil {
			fmt.Printf("Didn't find backup record by record id: %s\n", RecordID)
			return nil
		}

		if err := formatRecord([]*model.LsBackup{bak}); err != nil {
			return err
		}
		return nil
	}

	// show all backup record
	backupList, err := ls.ReadAll()
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("read backup record failed. err: %s", err))
	}

	if len(backupList) == 0 {
		fmt.Printf("Didn't find any backup record.\n")
		return nil
	}

	if err := formatRecord(backupList); err != nil {
		return err
	}

	return nil
}
func formatRecord(backups []*model.LsBackup, mode ...string) error {
	var m string

	if len(mode) == 0 {
		m = "json"
	} else {
		m = mode[0]
	}
	switch m {
	case "json":
		return formatRecordJSON(backups)
	case "table":
		// TODO format record table
		return nil
	default:
		return formatRecordJSON(backups)
	}
}

func formatRecordJSON(backups []*model.LsBackup) error {
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
