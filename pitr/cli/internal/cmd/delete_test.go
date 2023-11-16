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
	"reflect"
	"time"

	"bou.ke/monkey"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	mock_pkg "github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/mocks"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("test delete", func() {
	var (
		proxy *mock_pkg.MockIShardingSphereProxy
		ls    *mock_pkg.MockILocalStorage
		as    *mock_pkg.MockIAgentServer
		bak   = &model.LsBackup{
			Info: &model.BackupMetaInfo{
				ID:       "backup-id-1",
				FileName: "backup",
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
		monkey.Patch(getUserApproveInTerminal, func(_ string) error {
			return nil
		})
	})

	AfterEach(func() {
		ctrl.Finish()
		monkey.UnpatchAll()
	})

	It("test exec delete main func", func() {
		// patch ReadByID of mock ls
		monkey.PatchInstanceMethod(reflect.TypeOf(ls), "ReadByID", func(_ *mock_pkg.MockILocalStorage, _ string) (*model.LsBackup, error) { return bak, nil })
		monkey.Patch(pkg.NewAgentServer, func(_ string) pkg.IAgentServer { return as })

		RecordID = "backup-id"
		as.EXPECT().CheckStatus(gomock.Any()).Return(nil)
		ls.EXPECT().HideByName(bak.Info.FileName).Return(nil)
		as.EXPECT().DeleteBackup(gomock.Any()).Return(nil)
		ls.EXPECT().DeleteByHidedName(bak.Info.FileName).Return(nil)
		Expect(deleteRecord()).To(BeNil())
	})

	Context("test exec delete", func() {
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
			as.EXPECT().DeleteBackup(gomock.Any()).Do(func(_ *model.DeleteBackupIn) {
				time.Sleep(3 * time.Second)
			}).Return(nil)
			Expect(_execDelete(bak)).To(BeNil())
		})
	})
})
