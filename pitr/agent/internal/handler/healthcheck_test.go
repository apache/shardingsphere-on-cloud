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

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Health Check", func() {
	It("should failed with empty body", func() {
		req := httptest.NewRequest(http.MethodPost, "/api/healthz", nil)
		resp, err := app.Test(req)
		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
	})

	It("should success", func() {
		body := `{"username": "root", "password": "root", "db_name": "test", "db_port": 3306}`
		req := httptest.NewRequest(http.MethodPost, "/api/healthz", strings.NewReader(body))
		resp, err := app.Test(req)
		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
	})
})
