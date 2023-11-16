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
	"errors"
	"time"

	"bou.ke/monkey"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	mock_pkg "github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/mocks"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/timeutil"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Show", func() {
	Context("when show backup history", func() {
		It("format output", func() {
			bak := &model.LsBackup{
				Info: &model.BackupMetaInfo{
					ID:        "back-id",
					CSN:       "csn",
					StartTime: timeutil.Now().String(),
					EndTime:   timeutil.Now().Add(time.Second * 10).String(),
				},
				DnList: []*model.DataNode{
					{
						IP:        "127.0.0.1",
						Port:      3306,
						Status:    "Running",
						BackupID:  "back-id",
						StartTime: timeutil.Now().String(),
						EndTime:   timeutil.Now().Add(time.Second * 10).String(),
					},
					{
						IP:        "127.0.0.2",
						Port:      3306,
						Status:    "Completed",
						BackupID:  "back-id",
						StartTime: timeutil.Now().String(),
						EndTime:   timeutil.Now().Add(time.Second * 10).String(),
					},
				},
				SsBackup: &model.SsBackup{
					Status: "Running",
					ClusterInfo: &model.ClusterInfo{
						MetaData: model.MetaData{
							Databases: map[string]string{
								"db1": "db1",
							},
							Props: "props",
							Rules: "rules",
						},
						SnapshotInfo: &model.SnapshotInfo{},
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
			err := formatRecord([]*model.LsBackup{bak, bak})
			Expect(err).To(BeNil())
		})
	})

	Context("search backup history", func() {
		var (
			ls *mock_pkg.MockILocalStorage
		)

		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			ls = mock_pkg.NewMockILocalStorage(ctrl)

			monkey.Patch(pkg.NewLocalStorage, func(root string) (pkg.ILocalStorage, error) {
				return ls, nil
			})
		})

		AfterEach(func() {
			ctrl.Finish()
			monkey.UnpatchAll()
		})

		It("no record", func() {
			CSN = ""
			RecordID = ""
			ls.EXPECT().ReadAll().Return([]*model.LsBackup{}, nil)
			Expect(show()).To(BeNil())
		})

		It("get by csn", func() {
			CSN = "csn"
			RecordID = ""
			ls.EXPECT().ReadByCSN(gomock.Any()).Return(&model.LsBackup{}, nil)
			Expect(show()).To(BeNil())
		})
		It("get by csn failed", func() {
			CSN = "csn"
			RecordID = ""
			ls.EXPECT().ReadByCSN(gomock.Any()).Return(nil, errors.New("error"))
			Expect(show()).NotTo(BeNil())
		})
		It("get by csn empty", func() {
			CSN = "csn"
			RecordID = ""
			ls.EXPECT().ReadByCSN(gomock.Any()).Return(nil, nil)
			Expect(show()).To(BeNil())
		})

		It("get by id", func() {
			CSN = ""
			RecordID = "record-id"
			ls.EXPECT().ReadByID(gomock.Any()).Return(&model.LsBackup{}, nil)
			Expect(show()).To(BeNil())
		})
		It("get by id failed", func() {
			CSN = ""
			RecordID = "record-id"
			ls.EXPECT().ReadByID(gomock.Any()).Return(nil, errors.New("error"))
			Expect(show()).NotTo(BeNil())
		})
		It("get by id empty", func() {
			CSN = ""
			RecordID = "record-id"
			ls.EXPECT().ReadByID(gomock.Any()).Return(nil, nil)
			Expect(show()).To(BeNil())
		})
	})

})
