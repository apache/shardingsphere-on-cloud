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
	"reflect"
	"time"

	"bou.ke/monkey"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	mock_pkg "github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/mocks"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Restore", func() {
	var (
		// implement your own test case params
		user           string
		password       string
		host           string
		port           uint16
		backupRecordID string
	)
	Context("Restore", func() {
		proxy, _ := pkg.NewShardingSphereProxy(user, password, pkg.DefaultDBName, host, port)
		ls, _ := pkg.NewLocalStorage(pkg.DefaultRootDir())
		bak, _ := ls.ReadByID(backupRecordID)
		It("restore to ss proxy should be ok", func() {
			if user == "" || password == "" || host == "" || port == 0 || backupRecordID == "" {
				Skip("user or password or host or port or backupRecordID is empty")
			}
			databaseNamesExist = []string{"sharding_db"}
			Expect(restoreDataToSSProxy(proxy, bak)).To(BeNil())
		})
		It("check database if exists", func() {
			if user == "" || password == "" || host == "" || port == 0 || backupRecordID == "" {
				Skip("user or password or host or port or backupRecordID is empty")
			}
			Expect(checkDatabaseExist(proxy, bak)).To(Equal(xerr.NewCliErr("user abort")))
		})
	})
})

var _ = Describe("test restore", func() {
	var (
		proxy *mock_pkg.MockIShardingSphereProxy
		ls    *mock_pkg.MockILocalStorage
		as    *mock_pkg.MockIAgentServer
		bak   = &model.LsBackup{
			Info: &model.BackupMetaInfo{
				ID: "backup-id-1",
			},
			DnList: []*model.DataNode{
				{
					IP: "127.0.0.1",
				},
			},
			SsBackup: &model.SsBackup{
				Status: "",
				ClusterInfo: &model.ClusterInfo{
					MetaData: model.MetaData{
						Databases: map[string]string{},
						Props:     "",
						Rules:     "",
					},
					SnapshotInfo: nil,
				},
				StorageNodes: []*model.StorageNode{
					{
						IP: "127.0.0.1",
					},
				},
			},
		}
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		proxy = mock_pkg.NewMockIShardingSphereProxy(ctrl)
		as = mock_pkg.NewMockIAgentServer(ctrl)
		ls = mock_pkg.NewMockILocalStorage(ctrl)
		monkey.Patch(pkg.NewShardingSphereProxy, func(user, password, database, host string, port uint16) (pkg.IShardingSphereProxy, error) {
			return proxy, nil
		})
		monkey.Patch(pkg.NewLocalStorage, func(rootDir string) (pkg.ILocalStorage, error) {
			return ls, nil
		})
	})

	AfterEach(func() {
		ctrl.Finish()
		monkey.UnpatchAll()
	})

	It("check database if exists", func() {
		monkey.Patch(getUserApproveInTerminal, func(_ string) error { return nil })
		proxy.EXPECT().ExportMetaData()
		Expect(checkDatabaseExist(proxy, bak)).To(BeNil())
	})

	It("test exec restore main func", func() {
		// patch ReadByID of mock ls
		monkey.PatchInstanceMethod(reflect.TypeOf(ls), "ReadByID", func(_ *mock_pkg.MockILocalStorage, _ string) (*model.LsBackup, error) { return bak, nil })
		monkey.Patch(pkg.NewAgentServer, func(_ string) pkg.IAgentServer { return as })

		RecordID = "backup-id"
		proxy.EXPECT().ExportMetaData().Return(&model.ClusterInfo{}, nil)
		proxy.EXPECT().ImportMetaData(gomock.Any()).Return(nil)
		as.EXPECT().CheckStatus().Return(nil)
		as.EXPECT().Restore(gomock.Any()).Return(nil)
		Expect(restore()).To(BeNil())
	})

	// test getUserApproveInTerminal
	Context("test userApproveInTerminal", func() {
		// test user abort
		It("user abort", func() {
			// exec getUserApproveInTerminal
			Expect(getUserApproveInTerminal("")).To(Equal(xerr.NewCliErr("User abort")))
		})
	})

	Context("restore data to ss proxy", func() {

		It("no need to drop database", func() {
			proxy.EXPECT().ImportMetaData(gomock.Any())
			Expect(restoreDataToSSProxy(proxy, bak)).To(BeNil())
		})

		It("need to drop database first", func() {
			databaseNamesExist = []string{"sharding_db"}
			proxy.EXPECT().DropDatabase(gomock.Any())
			proxy.EXPECT().ImportMetaData(gomock.Any())
			Expect(restoreDataToSSProxy(proxy, bak)).To(BeNil())
		})
	})

	Context("test exec restore", func() {
		It("should be success", func() {
			ctrl := gomock.NewController(GinkgoT())
			as := mock_pkg.NewMockIAgentServer(ctrl)
			monkey.Patch(pkg.NewAgentServer, func(_ string) pkg.IAgentServer {
				return as
			})
			defer func() {
				ctrl.Finish()
				monkey.UnpatchAll()
			}()
			as.EXPECT().Restore(gomock.Any()).Do(func(_ *model.RestoreIn) {
				time.Sleep(3 * time.Second)
			}).Return(nil)
			Expect(execRestore(bak)).To(BeNil())
		})
	})
})
