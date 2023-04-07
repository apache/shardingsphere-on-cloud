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

package model

type (
	BackupIn struct {
		DBPort   uint16 `json:"db_port"`
		DBName   string `json:"db_name"`
		Username string `json:"username"`
		Password string `json:"password"`

		DnBackupPath string       `json:"dn_backup_path"`
		DnThreadsNum uint8        `json:"dn_threads_num"`
		DnBackupMode DBBackupMode `json:"dn_backup_mode"`
		Instance     string       `json:"instance"`
	}

	BackupOut struct {
		ID string `json:"backup_id"`
	}

	BackupOutResp struct {
		Code int       `json:"code" validate:"required"`
		Msg  string    `json:"msg" validate:"required"`
		Data BackupOut `json:"data"`
	}
)
