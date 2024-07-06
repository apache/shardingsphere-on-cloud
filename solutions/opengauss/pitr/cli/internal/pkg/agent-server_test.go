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

package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/httputils"

	"bou.ke/monkey"
	mock_httputils "github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/httputils/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAgentServer_Backup(t *testing.T) {
	t.SkipNow()
	//Note:just for test api,you need map you own host.
	as := NewAgentServer("http://agent-server:18080")

	backupID, err := as.Backup(&model.BackupIn{
		DBPort:       5432,
		DBName:       "omm",
		Username:     "og",
		Password:     "1234567890@SphereEx",
		DnBackupPath: "/home/omm/data",
		DnThreadsNum: 1,
		DnBackupMode: "FULL",
		Instance:     "ins-default-0",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(backupID)
}

func TestAgentServer_Restore(t *testing.T) {
	t.SkipNow()
	//Note:just for test api,you need map you own host.
	as := NewAgentServer("http://agent-server:18080")

	err := as.Restore(&model.RestoreIn{
		DBPort:       5432,
		DBName:       "omm",
		Username:     "og",
		Password:     "1234567890@SphereEx",
		DnBackupPath: "/home/omm/data",
		Instance:     "ins-default-0",
		DnBackupID:   "RR3FIC",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Success")
}

func TestAgentServer_ShowDetail(t *testing.T) {
	t.SkipNow()
	//Note:just for test api,you need map you own host.
	as := NewAgentServer("http://agent-server:18080")

	backupInfo, err := as.ShowDetail(&model.ShowDetailIn{
		DBPort:       5432,
		DBName:       "omm",
		Username:     "og",
		Password:     "1234567890@SphereEx",
		DnBackupPath: "/home/omm/data",
		Instance:     "ins-default-0",
		DnBackupID:   "RR3FIC",
	})
	if err != nil {
		panic(err)
	}

	indent, err := json.MarshalIndent(backupInfo, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(indent))
}

func TestAgentServer_ShowList(t *testing.T) {
	t.SkipNow()
	//Note:just for test api,you need map you own host.
	as := NewAgentServer("http://agent-server:18080")

	list, err := as.ShowList(&model.ShowListIn{
		DBPort:       5432,
		DBName:       "omm",
		Username:     "og",
		Password:     "1234567890@SphereEx",
		DnBackupPath: "/home/omm/data",
		Instance:     "ins-default-0",
	})
	if err != nil {
		panic(err)
	}

	indent, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(indent))
}

var _ = Describe("AgentServer", func() {
	var (
		ctrl *gomock.Controller
		req  *mock_httputils.MockIreq
		as   IAgentServer
	)
	BeforeEach(func() {
		as = NewAgentServer("http://agent-server:18080")
		ctrl = gomock.NewController(GinkgoT())
		req = mock_httputils.NewMockIreq(ctrl)
		req.EXPECT().Body(gomock.Any())
		monkey.Patch(httputils.NewRequest, func(c context.Context, method, url string) httputils.Ireq {
			return req
		})
	})
	AfterEach(func() {
		ctrl.Finish()
		monkey.UnpatchAll()
	})

	Context("backup", func() {
		It("backup failed", func() {
			req.EXPECT().Send(gomock.Any()).Return(fmt.Errorf("error"))
			_, err := as.Backup(&model.BackupIn{})
			Expect(err).ShouldNot(BeNil())
		})

		It("backup success", func() {
			req.EXPECT().Send(gomock.Any()).Do(func(i *model.BackupOutResp) {
				i.Data.ID = "backup-id"
			}).Return(nil)
			as := NewAgentServer("http://agent-server:18080")
			resp, err := as.Backup(&model.BackupIn{})
			Expect(err).Should(BeNil())
			Expect(resp).Should(Equal("backup-id"))
		})
	})

	Context("restore", func() {
		It("restore failed", func() {
			req.EXPECT().Send(gomock.Any()).Return(fmt.Errorf("error"))
			err := as.Restore(&model.RestoreIn{})
			Expect(err).ShouldNot(BeNil())
		})
		// restore success
		It("restore success", func() {
			req.EXPECT().Send(gomock.Any()).Return(nil)
			err := as.Restore(&model.RestoreIn{})
			Expect(err).Should(BeNil())
		})

	})

	Context("show detail", func() {
		It("show detail failed", func() {
			req.EXPECT().Send(gomock.Any()).Return(fmt.Errorf("error"))
			_, err := as.ShowDetail(&model.ShowDetailIn{})
			Expect(err).ShouldNot(BeNil())
		})
		// show detail success
		It("show detail success", func() {
			req.EXPECT().Send(gomock.Any()).Do(func(i *model.BackupDetailResp) {
				i.Data = model.BackupInfo{}
			}).Return(nil)
			resp, err := as.ShowDetail(&model.ShowDetailIn{})
			Expect(err).Should(BeNil())
			Expect(resp).Should(Equal(&model.BackupInfo{}))
		})
	})

	Context("show list", func() {
		It("show list failed", func() {
			req.EXPECT().Send(gomock.Any()).Return(fmt.Errorf("error"))
			_, err := as.ShowList(&model.ShowListIn{})
			Expect(err).ShouldNot(BeNil())
		})
		// show list success
		It("show list success", func() {
			req.EXPECT().Send(gomock.Any()).Do(func(i *model.BackupListResp) {
				i.Data = []model.BackupInfo{}
			}).Return(nil)
			resp, err := as.ShowList(&model.ShowListIn{})
			Expect(err).Should(BeNil())
			Expect(resp).Should(Equal([]model.BackupInfo{}))
		})
	})
})
