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
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

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
