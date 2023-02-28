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

package view

import (
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/pkg/model"
)

type (
	ShowIn struct {
		DbPort       uint16 `json:"db_port"`
		DbName       string `json:"db_name"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		DnBackupId   string `json:"dn_backup_id"`
		DnBackupPath string `json:"dn_backup_path"`
		Instance     string `json:"instance"`
	}

	BackupInfo struct {
		Id        string `json:"dn_backup_id"`
		Path      string `json:"dn_backup_path"`
		Mode      string `json:"db_backup_mode"`
		Instance  string `json:"instance"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		Status    string `json:"status"`
	}
)

func (in *ShowIn) Validate() error {
	if in == nil {
		return cons.Internal
	}

	if in.DbPort == 0 {
		return cons.InvalidDbPort
	}

	if in.DbName == "" {
		return cons.MissingDbName
	}

	if in.Username == "" {
		return cons.MissingUsername
	}

	if in.Password == "" {
		return cons.MissingPassword
	}

	if in.DnBackupPath == "" {
		return cons.MissingDnBackupPath
	}

	if in.DnBackupId == "" {
		return cons.MissingDnBackupId
	}

	if in.Instance == "" {
		return cons.MissingInstance
	}
	return nil
}

func NewBackupInfo(data *model.Backup, path, instance string) *BackupInfo {
	if data == nil {
		return nil
	}
	return &BackupInfo{
		Id:        data.ID,
		Path:      path,
		Mode:      data.BackupMode,
		Instance:  instance,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
		Status:    statusTrans(data.Status),
	}
}

func NewBackupInfoList(list []model.Backup, path, instance string) []BackupInfo {
	if len(list) == 0 {
		return []BackupInfo{}
	}
	ret := make([]BackupInfo, 0, len(list))
	for _, v := range list {
		ret = append(ret, BackupInfo{
			Id:        v.ID,
			Path:      path,
			Mode:      v.BackupMode,
			Instance:  instance,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
			Status:    statusTrans(v.Status),
		})
	}
	return ret
}

func statusTrans(status string) string {
	switch status {
	case "OK":
		return "Completed"
	case "ERROR":
		return "Failed"
	case "RUNNING":
		return "Running"
	default:
		return "Other"
	}
}
