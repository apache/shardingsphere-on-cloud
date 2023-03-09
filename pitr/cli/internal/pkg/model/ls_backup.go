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
	// LsBackup LocalStorageBackup
	LsBackup struct {
		Info     *BackupMetaInfo `json:"info"`
		DnList   []DataNode      `json:"dn_list"`
		SsBackup *SsBackup       `json:"ss_backup"`
	}

	BackupMetaInfo struct {
		ID        string `json:"id"`
		CSN       string `json:"csn"`
		StartTime int64  `json:"start_time"` // Unix time
		Endtime   int64  `json:"end_time"`   // Unix time
	}

	DataNode struct {
		IP        string `json:"ip"`
		Port      string `json:"port"`
		Status    string `json:"status"`
		BackupID  string `json:"backup_id"`
		StartTime int64  `json:"start_time"` // Unix time
		Endtime   int64  `json:"end_time"`   // Unix time
	}
)

type (
	SsBackup struct {
		Status       string         `json:"status"`
		ClusterInfo  ClusterInfo    `json:"cluster_info"`
		StorageNodes []StorageNodes `json:"storage_nodes"`
	}

	ClusterInfo struct {
		MetaData     MetaData     `json:"meta_data"`
		SnapshotInfo SnapshotInfo `json:"snapshot_info"`
	}

	StorageNodes struct {
		IP       string `json:"ip"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
		Remark   string `json:"remark"`
	}

	MetaData struct {
		Databases Databases `json:"databases"`
		Props     string    `json:"props"`
		Rules     string    `json:"rules"`
	}

	Databases struct {
		ShardingDb string `json:"sharding_db"`
		AnotherDb  string `json:"another_db"`
	}

	SnapshotInfo struct {
		Csn        string `json:"csn"`
		CreateTime string `json:"create_time"`
	}
)
