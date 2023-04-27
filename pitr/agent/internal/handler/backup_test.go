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

package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/pkg"
	mock_pkg "github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/pkg/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Backup", func() {
	Context("delete backup", func() {
		var mockOG *mock_pkg.MockIOpenGauss
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockOG = mock_pkg.NewMockIOpenGauss(ctrl)
			pkg.OG = mockOG
		})
		AfterEach(func() {
			ctrl.Finish()
		})

		It("delete failed with lack of required params", func() {
			requestBody := `{"db_port":3306,"db_name":"test_db"}`
			req := httptest.NewRequest(http.MethodDelete, "/api/backup", strings.NewReader(requestBody))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(500))
		})

		It("delete success", func() {
			requestBody := `{
				"db_port": 3306,
				"db_name": "test_db",
				"username": "user",
				"password": "password",
				"dn_backup_path": "/tmp",
				"backup_id": "backup_id",
				"instance": "instance"
			}`

			mockOG.EXPECT().Auth(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockOG.EXPECT().DelBackup(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

			req := httptest.NewRequest(http.MethodDelete, "/api/backup", strings.NewReader(requestBody))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})
	})
})
