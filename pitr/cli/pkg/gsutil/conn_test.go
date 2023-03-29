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

package gsutil

import (
	"database/sql"

	"bou.ke/monkey"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("OpenGauss", func() {
	Context("Connection", func() {
		It("empty user", func() {
			og, err := Open("", "root", "postgres", "127.0.0.1", uint16(13308))
			Expect(err.Error()).To(Equal("user is empty"))
			Expect(og).To(BeNil())
		})

		It("empty password", func() {
			og, err := Open("root", "", "postgres", "127.0.0.1", uint16(13308))
			Expect(err.Error()).To(Equal("password is empty"))
			Expect(og).To(BeNil())
		})

		It("empty database", func() {
			og, err := Open("root", "root", "", "127.0.0.1", uint16(13308))
			Expect(err.Error()).To(Equal("db name is empty"))
			Expect(og).To(BeNil())
		})

		It("Open and ping", func() {
			monkey.Patch(sql.Open, func(driverName, dataSourceName string) (*sql.DB, error) {
				return &sql.DB{}, nil
			})
			defer monkey.UnpatchAll()
			og, err := Open("root", "root", "postgres", "127.0.0.1", uint16(13308))
			Expect(err).To(BeNil())
			Expect(og).NotTo(BeNil())
		})
	})
})
