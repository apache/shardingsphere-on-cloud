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
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/gsutil"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test ShardingSphere Proxy With Sqlmock", func() {
	var (
		db     *sql.DB
		dbmock sqlmock.Sqlmock
		err    error
		proxy  IShardingSphereProxy

		clusterInfo = &model.ClusterInfo{
			SnapshotInfo: &model.SnapshotInfo{
				Csn:        "100",
				CreateTime: "2023-05-16 18:12:20",
			},
			MetaData: model.MetaData{
				Databases: map[string]string{
					"sharding_db": "sharding_db",
				},
			},
		}
	)
	BeforeEach(func() {
		db, dbmock, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(db).ShouldNot(BeNil())
		Expect(dbmock).ShouldNot(BeNil())
		monkey.Patch(gsutil.Open, func(_, _, _, _ string, _ uint16) (*sql.DB, error) {
			return db, nil
		})

		proxy, err = NewShardingSphereProxy("root", "root", "opengauss", "localhost", 13308)
		Expect(err).To(BeNil())
	})
	AfterEach(func() {
		db.Close()
	})

	It("export metadata", func() {
		data, err := json.Marshal(clusterInfo)
		Expect(err).To(BeNil())
		encodedData := base64.StdEncoding.EncodeToString(data)

		dbmock.ExpectQuery(regexp.QuoteMeta("EXPORT METADATA;")).WillReturnRows(sqlmock.NewRows([]string{"id", "create_time", "data"}).AddRow("id", "2023-05-16", encodedData))
		clusterInfo, err := proxy.ExportMetaData()
		Expect(err).To(BeNil())
		Expect(clusterInfo).NotTo(BeNil())
		Expect(clusterInfo.SnapshotInfo.Csn).To(Equal("100"))
	})

	It("import metadata", func() {
		dbmock.ExpectExec(regexp.QuoteMeta("IMPORT METADATA")).WillReturnResult(sqlmock.NewResult(1, 1))
		Expect(proxy.ImportMetaData(clusterInfo)).To(BeNil())
	})

})

var _ = Describe("IShardingSphereProxy", func() {
	Context("NewShardingSphereProxy", func() {
		var (
			host     = "local"
			port     = uint16(13308)
			username = "root"
			password = "root"
			dbName   = "postgres"
		)

		It("Connecting shardingsphere proxy", func() {
			Skip("Manually exec:dependent environment")
			ss, err := NewShardingSphereProxy(username, password, dbName, host, port)
			Expect(err).To(BeNil())
			Expect(ss).NotTo(BeNil())
		})

		It("Export meta data", func() {
			Skip("Manually exec:dependent environment")
			ss, err := NewShardingSphereProxy(username, password, dbName, host, port)
			Expect(err).To(BeNil())
			Expect(ss).NotTo(BeNil())

			fmt.Println(ss.ExportMetaData())
		})

		It("Export storage node", func() {
			Skip("Manually exec:dependent environment")
			ss, err := NewShardingSphereProxy(username, password, dbName, host, port)
			Expect(err).To(BeNil())
			Expect(ss).NotTo(BeNil())

			fmt.Println(ss.ExportStorageNodes())

			ss, err = NewShardingSphereProxy(username, password, dbName, host, port)
			Expect(err).To(BeNil())
			Expect(ss).NotTo(BeNil())

			fmt.Println(ss.ExportStorageNodes())
		})

		It("Lock and unlock", func() {
			Skip("Manually exec:dependent environment")
			ss, err := NewShardingSphereProxy(username, password, dbName, host, port)
			Expect(err).To(BeNil())
			Expect(ss).NotTo(BeNil())

			fmt.Println(ss.LockForRestore())
			time.Sleep(time.Second * 5)
			fmt.Println(ss.Unlock())

			fmt.Println(ss.LockForBackup())
			time.Sleep(time.Second * 5)
			fmt.Println(ss.Unlock())
		})

	})
})
