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

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/logging"
	"github.com/jedib0t/go-pretty/v6/table"
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

func getUserApproveInTerminal(prompt string) error {
	logging.Warn(fmt.Sprintf("\n%s", prompt))
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

func convertLocalhost(ip string) string {
	if ip == "127.0.0.1" {
		return Host
	}
	return ip
}

func checkAgentServerStatus(lsBackup *model.LsBackup) bool {

	statusList := make([]*model.AgentServerStatus, 0)

	// all agent server are available
	available := true

	for _, node := range lsBackup.SsBackup.StorageNodes {
		sn := node
		as := pkg.NewAgentServer(fmt.Sprintf("%s:%d", convertLocalhost(sn.IP), AgentPort))
		if err := as.CheckStatus(); err != nil {
			statusList = append(statusList, &model.AgentServerStatus{IP: sn.IP, Status: "Unavailable"})
			available = false
		} else {
			statusList = append(statusList, &model.AgentServerStatus{IP: sn.IP, Status: "Available"})
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Agent Server Status")
	t.AppendHeader(table.Row{"#", "Agent Server IP", "Status"})

	for i, s := range statusList {
		t.AppendRow([]interface{}{i + 1, s.IP, s.Status})
		t.AppendSeparator()
	}

	t.Render()

	return available
}

func checkDiskSpace(lsBackup *model.LsBackup) error {
	var (
		diskspaceList = make([]*model.DiskSpaceStatus, 0)
	)
	for _, sn := range lsBackup.SsBackup.StorageNodes {
		var data string
		as := pkg.NewAgentServer(fmt.Sprintf("%s:%d", convertLocalhost(sn.IP), AgentPort))
		in := &model.DiskSpaceIn{
			DiskPath: BackupPath,
		}

		out, err := as.ShowDiskSpace(in)

		if err != nil {
			data = "Check disk space failed."
		} else {
			data = out.Data
		}

		diskspaceList = append(diskspaceList, &model.DiskSpaceStatus{
			IP:              sn.IP,
			Path:            BackupPath,
			DiskSpaceStatus: data,
		})
	}

	// print diskspace result formatted
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Disk Space Status")
	t.AppendHeader(table.Row{"#", "Data Node IP", "Disk Path", "Disk Space Status"})
	for i, ds := range diskspaceList {
		t.AppendRow([]interface{}{i + 1, ds.IP, ds.Path, ds.DiskSpaceStatus})
		t.AppendSeparator()
	}
	t.Render()
	prompt := fmt.Sprintf(
		"Please Check All Nodes Disk Space, Make Sure Have Enough Space To Backup Or Restore Data.\n" +
			"Are you sure to continue? (Y/N)")
	return getUserApproveInTerminal(prompt)
}
