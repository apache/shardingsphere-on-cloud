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

import "github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"

type ShowIn struct {
	DbPort       uint16 `json:"db_port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DnBackupId   string `json:"dn_backup_id"`
	DnBackupPath string `json:"dn_backup_path"`
	Instance     string `json:"instance"`
}

func (in *ShowIn) Validate() error {
	if in == nil {
		return cons.Internal
	}

	if in.DbPort == 0 {
		return cons.InvalidDbPort
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
