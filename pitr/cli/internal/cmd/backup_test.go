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
	"context"
	"errors"

	"bou.ke/monkey"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	mock_pkg "github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/mocks"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/httputils"
	mock_httputils "github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/httputils/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Backup", func() {

	Context("do check", func() {
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
			as.EXPECT().ShowDetail(gomock.Any()).Return(nil, errors.New("timeout"))
			status, err := doCheck(as, sn, "", 0)
			Expect(err).To(HaveOccurred())
			Expect(status).To(Equal(model.SsBackupStatusCheckError))
		})

		It("mock agent server and return failed status", func() {
			as.EXPECT().ShowDetail(gomock.Any()).Return(&model.BackupInfo{Status: model.SsBackupStatusFailed}, nil)
			status, err := doCheck(as, sn, "", 0)
			Expect(err).ToNot(HaveOccurred())
			Expect(status).To(Equal(model.SsBackupStatusFailed))
		})

		It("mock agent server and return completed status", func() {
			as.EXPECT().ShowDetail(gomock.Any()).Return(&model.BackupInfo{Status: model.SsBackupStatusCompleted}, nil)
			status, err := doCheck(as, sn, "", 0)
			Expect(err).ToNot(HaveOccurred())
			Expect(status).To(Equal(model.SsBackupStatusCompleted))
		})

		It("mock agent server and return check err first time and then success", func() {
			as.EXPECT().ShowDetail(gomock.Any()).Return(nil, errors.New("timeout"))
			as.EXPECT().ShowDetail(gomock.Any()).Return(&model.BackupInfo{Status: model.SsBackupStatusCompleted}, nil)
			status, err := doCheck(as, sn, "", 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(status).To(Equal(model.SsBackupStatusCompleted))
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
		var (
			as *mock_pkg.MockIAgentServer
		)
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			as = mock_pkg.NewMockIAgentServer(ctrl)
		})
		AfterEach(func() {
			ctrl.Finish()
		})
		bak := &model.LsBackup{
			DnList: nil,
			SsBackup: &model.SsBackup{
				Status:       "Running",
				StorageNodes: []*model.StorageNode{},
			},
		}

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

	Context("check backup status", func() {
		var (
			as       *mock_pkg.MockIAgentServer
			lsbackup *model.LsBackup
		)
		BeforeEach(func() {
			lsbackup = &model.LsBackup{
				DnList: []*model.DataNode{
					{
						IP:   "127.0.0.1",
						Port: 3306,
					},
					{
						IP:   "127.0.0.2",
						Port: 3307,
					},
				},
				SsBackup: &model.SsBackup{
					Status: "Running",
					StorageNodes: []*model.StorageNode{
						{
							IP:   "127.0.0.1",
							Port: 3306,
						},
						{
							IP:   "127.0.0.2",
							Port: 3307,
						},
					},
				},
				Info: &model.BackupMetaInfo{},
			}

			ctrl = gomock.NewController(GinkgoT())
			as = mock_pkg.NewMockIAgentServer(ctrl)

			monkey.Patch(pkg.NewAgentServer, func(_ string) pkg.IAgentServer {
				return as
			})
		})
		AfterEach(func() {
			ctrl.Finish()
			monkey.UnpatchAll()
		})

		It("check error 1", func() {
			as.EXPECT().ShowDetail(gomock.Any()).Return(nil, errors.New("timeout")).AnyTimes()
			Expect(checkBackupStatus(lsbackup)).To(Equal(model.SsBackupStatusFailed))
		})

		It("check error 2", func() {
			as.EXPECT().ShowDetail(gomock.Any()).Return(nil, errors.New("timeout")).Times(1)
			as.EXPECT().ShowDetail(gomock.Any()).Return(&model.BackupInfo{Status: model.SsBackupStatusFailed}, nil).AnyTimes()
			Expect(checkBackupStatus(lsbackup)).To(Equal(model.SsBackupStatusFailed))
		})

		It("check error 3", func() {
			as.EXPECT().ShowDetail(gomock.Any()).Return(nil, errors.New("timeout")).Times(2)
			as.EXPECT().ShowDetail(gomock.Any()).Return(&model.BackupInfo{Status: model.SsBackupStatusCompleted}, nil).AnyTimes()
			Expect(checkBackupStatus(lsbackup)).To(Equal(model.SsBackupStatusCompleted))
		})

		It("check failed", func() {
			as.EXPECT().ShowDetail(gomock.Any()).Return(&model.BackupInfo{Status: model.SsBackupStatusFailed}, nil).AnyTimes()
			Expect(checkBackupStatus(lsbackup)).To(Equal(model.SsBackupStatusFailed))
		})

		It("check success", func() {
			as.EXPECT().ShowDetail(gomock.Any()).Return(&model.BackupInfo{Status: model.SsBackupStatusCompleted}, nil).AnyTimes()
			Expect(checkBackupStatus(lsbackup)).To(Equal(model.SsBackupStatusCompleted))
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
			monkey.Patch(getUserApproveInTerminal, func(_ string) error {
				return nil
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

	Context("test check agent server status", func() {
		var mockCtrl *gomock.Controller
		var mockIreq *mock_httputils.MockIreq

		ls := &model.LsBackup{
			Info:   nil,
			DnList: nil,
			SsBackup: &model.SsBackup{
				StorageNodes: []*model.StorageNode{
					{
						IP:   "127.0.0.1",
						Port: 3306,
					},
					{
						IP:   "127.0.0.2",
						Port: 3307,
					},
				},
			},
		}
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockIreq = mock_httputils.NewMockIreq(mockCtrl)

			monkey.Patch(httputils.NewRequest, func(ctx context.Context, method, url string) httputils.Ireq {
				return mockIreq
			})
		})

		AfterEach(func() {
			monkey.UnpatchAll()
			mockCtrl.Finish()
		})

		It("agent server is not running", func() {
			mockIreq.EXPECT().Send(gomock.Any()).Return(errors.New("error")).AnyTimes()
			mockIreq.EXPECT().Header(gomock.Any()).AnyTimes()
			Expect(checkAgentServerStatus(ls)).To(BeFalse())
		})

		It("agent server are running", func() {
			mockIreq.EXPECT().Send(gomock.Any()).Return(nil).AnyTimes()
			mockIreq.EXPECT().Header(gomock.Any()).AnyTimes()
			Expect(checkAgentServerStatus(ls)).To(BeTrue())
		})

		It("one agent server is not running", func() {
			mockIreq.EXPECT().Send(gomock.Any()).Return(errors.New("failed"))
			mockIreq.EXPECT().Send(gomock.Any()).Return(nil).AnyTimes()
			mockIreq.EXPECT().Header(gomock.Any()).AnyTimes()
			Expect(checkAgentServerStatus(ls)).To(BeFalse())
		})
	})

	Context("test delete backup data", func() {
		bak := &model.LsBackup{
			Info: nil,
			DnList: []*model.DataNode{
				{
					IP:   "test.delete.backup",
					Port: 3306,
				},
			},
			SsBackup: &model.SsBackup{
				StorageNodes: []*model.StorageNode{
					{
						IP:   "test.delete.backup",
						Port: 3306,
					},
				},
			},
		}
		It("should delete failed", func() {
			deleteBackupFiles(bak)
		})

		It("should delete success", func() {
			ctrl := gomock.NewController(GinkgoT())
			as := mock_pkg.NewMockIAgentServer(ctrl)
			monkey.Patch(pkg.NewAgentServer, func(addr string) pkg.IAgentServer {
				return as
			})

			defer monkey.UnpatchAll()
			defer ctrl.Finish()
			as.EXPECT().DeleteBackup(gomock.Any()).Return(nil)
			deleteBackupFiles(bak)
		})
	})
})
