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

package httputils

import (
	"bytes"
	"context"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test req", func() {
	Context("Test set header", func() {
		It("should set header for post", func() {
			bs := []byte("test")
			r := NewRequest(context.Background(), "POST", "http://localhost:8080")
			_req, err := http.NewRequestWithContext(context.Background(), "POST", "http://localhost:8080", bytes.NewReader(bs))
			Expect(err).To(BeNil())

			// check header
			_req = r.(*req).setReqHeader(_req)
			Expect(_req.Header.Get("Content-Type")).To(Equal("application/json"))
			Expect(_req.Header.Get("x-request-id")).ToNot(BeEmpty())
		})
	})
})
