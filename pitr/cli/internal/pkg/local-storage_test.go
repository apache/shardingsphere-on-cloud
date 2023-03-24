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
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ILocalStorage", func() {
	AfterEach(func() {
		root := fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".gs_pitr")
		_ = os.RemoveAll(root)
	})
	Context("localStorage", func() {
		It("New:Initialize the cli program directory.", func() {
			root := fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".gs_pitr")
			ls, err := NewLocalStorage(root)
			Expect(err).To(BeNil())
			Expect(ls).NotTo(BeNil())
		})

		It("ReadALL", func() {
			root := fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".gs_pitr")
			ls, err := NewLocalStorage(root)
			Expect(err).To(BeNil())
			Expect(ls).NotTo(BeNil())

			list, err := ls.ReadAll()
			Expect(err).To(BeNil())
			fmt.Println(fmt.Sprintf("%+v", list))
			for _, v := range list {
				fmt.Println(fmt.Sprintf("%+v", v.Info))
			}
		})

		It("ReadByCSN", func() {
			root := fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".gs_pitr")
			ls, err := NewLocalStorage(root)
			Expect(err).To(BeNil())
			Expect(ls).NotTo(BeNil())

			_, err = ls.ReadByCSN("e19b6935-c437-4cf0-b820-3275bd2727a2")
			Expect(err.Error()).To(Equal("Not found"))
		})

		It("ReadByID", func() {
			root := fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".gs_pitr")
			ls, err := NewLocalStorage(root)
			Expect(err).To(BeNil())
			Expect(ls).NotTo(BeNil())

			_, err = ls.ReadByID("66785e18-b8d3-42f4-9967-a4119be15cea")
			Expect(err.Error()).To(Equal("Not found"))
		})

		It("GenFilename and WriteByJSON", func() {
			root := fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".gs_pitr")
			ls, err := NewLocalStorage(root)
			Expect(err).To(BeNil())
			Expect(ls).NotTo(BeNil())

			filename := ls.GenFilename(ExtnJSON)
			Expect(filename).NotTo(BeEmpty())
			backupID := uuid.New().String()

			contents := model.LsBackup{
				Info: &model.BackupMetaInfo{
					ID:        backupID,
					CSN:       uuid.New().String(),
					StartTime: time.Now().Unix(),
					EndTime:   time.Now().Add(time.Minute).Unix(),
				},
				DnList: []*model.DataNode{
					{
						IP:        "1.1.1.1",
						Port:      5432,
						Status:    "Completed",
						BackupID:  "SK08DAK1",
						StartTime: time.Now().Unix(),
						EndTime:   time.Now().Unix(),
					},
					{
						IP:        "1.1.1.2",
						Port:      5432,
						Status:    "Completed",
						BackupID:  "SK08DAK2",
						StartTime: time.Now().Unix(),
						EndTime:   time.Now().Unix(),
					},
				},
				SsBackup: &model.SsBackup{
					Status: "Completed",
					ClusterInfo: &model.ClusterInfo{
						MetaData: model.MetaData{
							Databases: map[string]string{
								"db1": "db1",
							},
							Props: "Props",
							Rules: "Rules",
						},
						SnapshotInfo: &model.SnapshotInfo{},
					},
					StorageNodes: nil,
				},
			}
			err = ls.WriteByJSON(filename, &contents)
			Expect(err).To(BeNil())
			bak, err := ls.ReadByID(backupID)
			Expect(err).To(BeNil())
			Expect(bak).NotTo(BeNil())
		})
	})
})
