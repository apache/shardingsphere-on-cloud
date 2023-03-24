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

	"bou.ke/monkey"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	mock_pkg "github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/mocks"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var ctrl *gomock.Controller

var _ = Describe("Backup", func() {
	Context("check status", func() {
		var (
			as *mock_pkg.MockIAgentServer
			sn = &model.StorageNode{
				IP: "127.0.0.1",
			}
		)
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			as = mock_pkg.NewMockIAgentServer(ctrl)
		})
		AfterEach(func() {
			ctrl.Finish()
		})

		It("agent server return err", func() {
			as.EXPECT().ShowDetail(&model.ShowDetailIn{Instance: defaultInstance}).Return(nil, errors.New("timeout"))
			Expect(checkStatus(as, sn, "", "", 0)).To(Equal(model.SsBackupStatusCheckError))
		})

		It("mock agent server and return failed status", func() {
			as.EXPECT().ShowDetail(&model.ShowDetailIn{Instance: defaultInstance}).Return(&model.BackupInfo{Status: model.SsBackupStatusFailed}, nil)
			Expect(checkStatus(as, sn, "", "", 0)).To(Equal(model.SsBackupStatusFailed))
		})

		It("mock agent server and return completed status", func() {
			as.EXPECT().ShowDetail(&model.ShowDetailIn{Instance: defaultInstance}).Return(&model.BackupInfo{Status: model.SsBackupStatusCompleted}, nil)
			Expect(checkStatus(as, sn, "", "", 0)).To(Equal(model.SsBackupStatusCompleted))
		})

		It("mock agent server and return timeout error first time and then retry 1 time return completed status", func() {
			as.EXPECT().ShowDetail(&model.ShowDetailIn{Instance: defaultInstance}).Times(1).Return(nil, errors.New("timeout"))
			as.EXPECT().ShowDetail(&model.ShowDetailIn{Instance: defaultInstance}).Return(&model.BackupInfo{Status: model.SsBackupStatusCompleted}, nil)
			Expect(checkStatus(as, sn, "", "", 1)).To(Equal(model.SsBackupStatusCompleted))
		})
	})

	Context("export data", func() {
		var (
			proxy *mock_pkg.MockIShardingSphereProxy
			ls    *mock_pkg.MockILocalStorage
		)

		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			proxy = mock_pkg.NewMockIShardingSphereProxy(ctrl)
			ls = mock_pkg.NewMockILocalStorage(ctrl)
		})

		AfterEach(func() {
			ctrl.Finish()
		})

		It("export data", func() {
			// mock proxy export metadata
			proxy.EXPECT().ExportMetaData().Return(&model.ClusterInfo{}, nil)
			// mock proxy export node storage data
			proxy.EXPECT().ExportStorageNodes().Return([]*model.StorageNode{}, nil)
			// mock ls generate filename
			ls.EXPECT().GenFilename(pkg.ExtnJSON).Return("mock.json")
			// mock ls write by json
			ls.EXPECT().WriteByJSON("mock.json", gomock.Any()).Return(nil)

			bk, err := exportData(proxy, ls)
			Expect(err).To(BeNil())
			Expect(bk.Info.CSN).To(Equal(""))
		})
	})
	Context("exec backup", func() {

		var as *mock_pkg.MockIAgentServer
		bak := &model.LsBackup{
			DnList: nil,
			SsBackup: &model.SsBackup{
				Status:       "Running",
				StorageNodes: []*model.StorageNode{},
			},
		}
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			as = mock_pkg.NewMockIAgentServer(ctrl)
		})
		AfterEach(func() {
			ctrl.Finish()
		})
		It("exec backup empty storage nodes", func() {
			Expect(execBackup(bak)).To(BeNil())
		})
		It("exec backup 2 storage nodes", func() {
			bak.SsBackup.StorageNodes = []*model.StorageNode{
				{
					IP:       "127.0.0.1",
					Port:     80,
					Username: "",
					Password: "",
					Database: "",
					Remark:   "",
				},
				{
					IP:       "127.0.0.2",
					Port:     443,
					Username: "",
					Password: "",
					Database: "",
					Remark:   "",
				},
			}
			as.EXPECT().Backup(gomock.Any()).Return("", nil)
			dnCh := make(chan *model.DataNode, 10)
			Expect(_execBackup(as, bak.SsBackup.StorageNodes[0], dnCh)).To(BeNil())
			Expect(len(dnCh)).To(Equal(1))
			as.EXPECT().Backup(gomock.Any()).Return("", xerr.NewCliErr("backup failed"))
			Expect(_execBackup(as, bak.SsBackup.StorageNodes[0], dnCh)).ToNot(BeNil())
			close(dnCh)
			Expect(len(dnCh)).To(Equal(1))

		})
	})

	Context("exec backup", func() {
		It("exec backup", func() {
			var (
				as       *mock_pkg.MockIAgentServer
				node     = &model.StorageNode{}
				failSnCh = make(chan *model.StorageNode, 10)
				dnCh     = make(chan *model.DataNode, 10)
			)
			as = mock_pkg.NewMockIAgentServer(ctrl)

			defer close(dnCh)
			defer ctrl.Finish()
			as.EXPECT().Backup(gomock.Any()).Return("backup-id", nil)
			Expect(_execBackup(as, node, dnCh)).To(BeNil())
			Expect(len(failSnCh)).To(Equal(0))
			Expect(len(dnCh)).To(Equal(1))
		})
	})
})

var _ = Describe("test backup manually", func() {
	var (
		// implement with your own dev
		u  string
		p  string
		db string
		h  string
		pt uint16
	)
	Context("test manually", func() {
		It("unlock after lock", func() {
			if u == "" || p == "" || db == "" || h == "" || pt == 0 {
				Skip("need to set u, p, db, h, pt first")
			}
			proxy, _ := pkg.NewShardingSphereProxy(u, p, db, h, pt)
			Expect(proxy.LockForBackup()).To(BeNil())
			Expect(proxy.Unlock()).To(BeNil())
		})

		It("export data in dev", func() {
			if u == "" || p == "" || db == "" || h == "" || pt == 0 {
				Skip("need to set u, p, db, h, pt first")
			}
			proxy, _ := pkg.NewShardingSphereProxy(u, p, db, h, pt)
			ls, _ := pkg.NewLocalStorage("./")

			Expect(proxy.LockForBackup()).To(BeNil())
			defer func() {
				Expect(proxy.Unlock()).To(BeNil())
			}()

			bk, err := exportData(proxy, ls)

			Expect(err).To(BeNil())
			Expect(bk.Info).NotTo(BeNil())
		})

		It("test all", func() {
			if u == "" || p == "" || db == "" || h == "" || pt == 0 {
				Skip("need to set u, p, db, h, pt first")
			}
			proxy, _ := pkg.NewShardingSphereProxy(u, p, db, h, pt)
			ls, _ := pkg.NewLocalStorage(pkg.DefaultRootDir())
			bak, err := exportData(proxy, ls)
			Expect(err).To(BeNil())
			Expect(bak.Info).NotTo(BeNil())

			AgentPort = 18080
			BackupPath = "/home/omm/data"
			ThreadsNum = 1

			err = execBackup(bak)
			Expect(err).To(BeNil())
			Expect(bak.SsBackup.Status).To(Equal(model.SsBackupStatusRunning))

			err = ls.WriteByJSON(filename, bak)
			Expect(err).To(BeNil())

			Expect(checkBackupStatus(bak)).To(Equal(model.SsBackupStatusCompleted))

			err = ls.WriteByJSON(filename, bak)
			Expect(err).To(BeNil())
		})
	})
})

var _ = Describe("test backup mock", func() {
	Context("test by mock", func() {
		var (
			proxy *mock_pkg.MockIShardingSphereProxy
			ls    *mock_pkg.MockILocalStorage
			as    *mock_pkg.MockIAgentServer
		)
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			proxy = mock_pkg.NewMockIShardingSphereProxy(ctrl)
			ls = mock_pkg.NewMockILocalStorage(ctrl)
			as = mock_pkg.NewMockIAgentServer(ctrl)

			monkey.Patch(pkg.NewShardingSphereProxy, func(u, p, db, h string, pt uint16) (pkg.IShardingSphereProxy, error) {
				return proxy, nil
			})
			monkey.Patch(pkg.NewLocalStorage, func(rootDir string) (pkg.ILocalStorage, error) {
				return ls, nil
			})
			monkey.Patch(pkg.NewAgentServer, func(addr string) pkg.IAgentServer {
				return as
			})
		})
		AfterEach(func() {
			monkey.UnpatchAll()
			ctrl.Finish()
		})

		It("test backup empty", func() {
			proxy.EXPECT().LockForBackup().Return(nil)
			proxy.EXPECT().ExportMetaData().Return(&model.ClusterInfo{}, nil)
			proxy.EXPECT().ExportStorageNodes().Return([]*model.StorageNode{}, nil)
			proxy.EXPECT().Unlock().Return(nil)
			ls.EXPECT().GenFilename(gomock.Any()).Return("filename")
			ls.EXPECT().WriteByJSON(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			Expect(backup()).To(BeNil())
		})
	})
})
