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
	"bou.ke/monkey"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	mock_pkg "github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/mocks"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/promptutil"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Root", func() {
	Context("when check disk space", func() {
		var (
			as *mock_pkg.MockIAgentServer
		)
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			as = mock_pkg.NewMockIAgentServer(ctrl)

			monkey.Patch(pkg.NewAgentServer, func(_ string) pkg.IAgentServer {
				return as
			})
			monkey.Patch(promptutil.GetUserApproveInTerminal, func(_ string) error {
				return nil
			})
		})
		AfterEach(func() {
			defer monkey.UnpatchAll()
		})

		It("should print help message", func() {
			bak := &model.LsBackup{
				SsBackup: &model.SsBackup{
					StorageNodes: []*model.StorageNode{
						{
							IP: "127.0.0.1",
						},
					},
				},
			}
			BackupPath = "/tmp"
			as.EXPECT().ShowDiskSpace(gomock.Any()).Return(&model.DiskSpaceInfo{
				Code: 0,
				Msg:  "",
				Data: "文件系统        容量  已用  可用 已用% 挂载点\n/dev/vda1        20G   13G  7.8G   62% /",
			}, nil)
			err := checkDiskSpace(bak)
			Expect(err).To(BeNil())
		})
	})
})
