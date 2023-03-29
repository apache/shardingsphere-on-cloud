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
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/spf13/cobra"
)

var (
	// Host ss-proxy host
	Host string
	// Port ss-proxy port
	Port uint16
	// Username ss-proxy username
	Username string
	// Password ss-proxy password
	Password string
	// AgentPort agent-server port
	AgentPort uint16
	// BackupPath openGauss data backup path
	BackupPath string
	// BackupModeStr openGauss data backup mode string (FULL or PTRACK)
	BackupModeStr string
	// BackupMode openGauss data backup mode (FULL or PTRACK)
	BackupMode model.DBBackupMode
	// ThreadsNum openGauss data backup task thread num
	ThreadsNum uint8
	// CSN openGauss data backup commit sequence number
	CSN string
	// RecordID openGauss data backup record id
	RecordID string
)

var RootCmd = &cobra.Command{
	Use:   "gs_pitr",
	Short: "PITR tools for openGauss",

	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
		HiddenDefaultCmd:  true,
	},
}
