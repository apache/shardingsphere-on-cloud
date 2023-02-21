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

	"github.com/spf13/cobra"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/logging"
)

const (
	backupRecordID = "id"
	csn            = "csn"
)

var Restore = &cobra.Command{
	Use:   "restore",
	Short: "Restore a database cluster ",
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

		csn, err := cmd.Flags().GetString(csn)
		if err != nil {
			logging.Error(err.Error())
		}
		logging.Info(fmt.Sprintf("flags:csn:%s", csn))

		id, err := cmd.Flags().GetString(backupRecordID)
		if err != nil {
			logging.Error(err.Error())
		}
		logging.Info(fmt.Sprintf("flags:id:%s", id))

		logging.Info("Restore...")
	},
}

func init() {
	Restore.PersistentFlags().StringP(csn, "", "", "commit sequence number")
	Restore.PersistentFlags().StringP(backupRecordID, "", "", "backup record id")
}
