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
)

type (
	BackupIn struct {
		DBPort   uint16 `json:"db_port"`
		DBName   string `json:"db_name"`
		Username string `json:"username"`
		Password string `json:"password"`

		DnBackupPath string `json:"dn_backup_path"`
		DnThreadsNum uint8  `json:"dn_threads_num"`
		DnBackupMode string `json:"dn_backup_mode"`
		Instance     string `json:"instance"`
	}

	BackupOut struct {
		ID string `json:"backup_id"`
	}

	DeleteBackupIn struct {
		DBPort   uint16 `json:"db_port"`
		DBName   string `json:"db_name"`
		Username string `json:"username"`
		Password string `json:"password"`

		DnBackupPath string `json:"dn_backup_path"`
		Instance     string `json:"instance"`
		BackupID     string `json:"backup_id"`
	}
)

func (in *BackupIn) Validate() error {
	if in == nil {
		return cons.Internal
	}

	if in.DBPort == 0 {
		return cons.InvalidDBPort
	}

	if in.DBName == "" {
		return cons.MissingDBName
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

	if in.DnThreadsNum == 0 {
		return cons.InvalidDnThreadsNum
	}

	if in.DnBackupMode == "" {
		return cons.MissingDnBackupMode
	}

	if in.DnBackupMode != cons.DBBackModeFull && in.DnBackupMode != cons.DBBackModePTrack {
		return cons.InvalidDnBackupMode
	}

	if in.Instance == "" {
		return cons.MissingInstance
	}
	return nil
}

// nolint:dupl
func (in *DeleteBackupIn) Validate() error {
	if in == nil {
		return cons.Internal
	}

	if in.DBPort == 0 {
		return cons.InvalidDBPort
	}

	if in.DBName == "" {
		return cons.MissingDBName
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

	if in.BackupID == "" {
		return cons.MissingBackupID
	}

	if in.Instance == "" {
		return cons.MissingInstance
	}
	return nil
}
