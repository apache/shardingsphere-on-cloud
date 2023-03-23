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
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
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
		proxy, _ := pkg.NewShardingSphereProxy(user, password, pkg.DefaultDbName, host, port)
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
