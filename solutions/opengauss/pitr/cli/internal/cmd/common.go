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

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/logging"
	"github.com/jedib0t/go-pretty/v6/table"
)

func validate(ls pkg.ILocalStorage, csn, recordID string) ([]*model.LsBackup, error) {
	var (
		baks []*model.LsBackup
		err  error
	)
	if CSN != "" {
		baks, err = ls.ReadAllByCSN(csn)
		if err != nil {
			return baks, xerr.NewCliErr(fmt.Sprintf("read backup record by csn failed. err: %s", err))
		}
	}

	if RecordID != "" {
		b, err := ls.ReadByID(recordID)
		baks = append(baks, b)
		if err != nil {
			return baks, xerr.NewCliErr(fmt.Sprintf("read backup record by id failed. err: %s", err))
		}
	}

	if len(baks) == 0 {
		return baks, xerr.NewCliErr(fmt.Sprintf("backup record not found. err: %s", err))
	}
	if len(baks) > 1 {
		return baks, xerr.NewCliErr("multiple backup records found. please using ID to submit one specific record.")
	}

	return baks, nil
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

	// IMPORTANT: we don't support multiple agent server run on the same host
	asMap := make(map[string]bool)
	asDuplicate := false

	for _, node := range lsBackup.SsBackup.StorageNodes {
		sn := node
		as := pkg.NewAgentServer(fmt.Sprintf("%s:%d", convertLocalhost(sn.IP), AgentPort))
		in := &model.HealthCheckIn{
			DBPort:   sn.Port,
			DBName:   sn.Database,
			Username: sn.Username,
			Password: sn.Password,
		}
		if err := as.CheckStatus(in); err != nil {
			statusList = append(statusList, &model.AgentServerStatus{IP: sn.IP, Port: AgentPort, Status: fmt.Sprintf("Unavailable: %s", err)})
			available = false
		} else {
			statusList = append(statusList, &model.AgentServerStatus{IP: sn.IP, Port: AgentPort, Status: "Available"})
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Agent Server Status")
	t.AppendHeader(table.Row{"#", "Agent Server IP", "Agent Server Port", "Status"})

	for i, s := range statusList {
		t.AppendRow([]interface{}{i + 1, s.IP, s.Port, s.Status})
		t.AppendSeparator()
	}

	t.Render()

	for _, node := range lsBackup.SsBackup.StorageNodes {
		if _, ok := asMap[node.IP]; ok {
			asDuplicate = true
			break
		}
		asMap[node.IP] = true
	}

	if asDuplicate {
		logging.Error("IMPORTANT!: we don't support multiple agent server run on the same host.\n")
		return false
	}

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
	return nil
}
