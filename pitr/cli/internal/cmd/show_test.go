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
 *
 */

package cmd

import (
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Show", func() {
	Context("when show backup history", func() {
		It("format output", func() {
			bak := &model.LsBackup{
				Info: &model.BackupMetaInfo{
					ID:        "back-id",
					CSN:       "csn",
					StartTime: time.Now().Unix(),
					EndTime:   time.Now().Add(time.Second * 10).Unix(),
				},
				DnList: []*model.DataNode{
					{
						IP:        "127.0.0.1",
						Port:      3306,
						Status:    "Running",
						BackupID:  "back-id",
						StartTime: time.Now().Unix(),
						EndTime:   time.Now().Add(time.Second * 10).Unix(),
					},
					{
						IP:        "127.0.0.2",
						Port:      3306,
						Status:    "Completed",
						BackupID:  "back-id",
						StartTime: time.Now().Unix(),
						EndTime:   time.Now().Add(time.Second * 10).Unix(),
					},
				},
				SsBackup: &model.SsBackup{
					Status: "Running",
					ClusterInfo: &model.ClusterInfo{
						MetaData: model.MetaData{
							Databases: model.Databases{
								ShardingDb: "sharding_db",
								AnotherDb:  "",
							},
							Props: "props",
							Rules: "rules",
						},
						SnapshotInfo: model.SnapshotInfo{},
					},
					StorageNodes: []*model.StorageNode{
						{
							IP:       "127.0.0.1",
							Port:     3306,
							Username: "username",
							Password: "password",
							Database: "database",
							Remark:   "remark",
						},
						{
							IP:       "127.0.0.2",
							Port:     3306,
							Username: "username",
							Password: "password",
							Database: "database",
							Remark:   "remark",
						},
					},
				},
			}
			err := formatRecord([]model.LsBackup{*bak, *bak})
			Expect(err).To(BeNil())
		})
	})
	Context("search backup history", func() {
		It("no record", func() {
			Expect(show()).To(BeNil())
		})
	})

})
