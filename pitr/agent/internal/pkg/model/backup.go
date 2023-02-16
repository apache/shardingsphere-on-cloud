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
	Backup struct {
		ID                string `json:"id"`
		BackupMode        string `json:"backup-mode"`
		Wal               string `json:"wal"`
		CompressAlg       string `json:"compress-alg"`
		CompressLevel     int    `json:"compress-level"`
		FromReplica       string `json:"from-replica"`
		BlockSize         int    `json:"block-size"`
		XlogBlockSize     int    `json:"xlog-block-size"`
		ChecksumVersion   int    `json:"checksum-version"`
		ProgramVersion    string `json:"program-version"`
		ServerVersion     string `json:"server-version"`
		CurrentTli        int    `json:"current-tli"`
		ParentTli         int    `json:"parent-tli"`
		StartLsn          string `json:"start-lsn"`
		StopLsn           string `json:"stop-lsn"`
		StartTime         string `json:"start-time"`
		EndTime           string `json:"end-time"`
		RecoveryXid       int    `json:"recovery-xid"`
		RecoveryTime      string `json:"recovery-time"`
		RecoveryName      string `json:"recovery-name"`
		DataBytes         int    `json:"data-bytes"`
		WalBytes          int    `json:"wal-bytes"`
		UncompressedBytes int    `json:"uncompressed-bytes"`
		PgdataBytes       int    `json:"pgdata-bytes"`
		Status            string `json:"status"`
		ContentCrc        int64  `json:"content-crc"`
	}

	BackupList struct {
		Instance string   `json:"instance"`
		List     []Backup `json:"backups"`
	}
)
