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
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type req struct {
	ctx    context.Context
	method string
	header map[string]string
	url    string
	body   any
	query  map[string]string
}

type Ireq interface {
	Header(h map[string]string)
	Body(b any)
	Query(m map[string]string)
	Send(body any) error
}

func NewRequest(ctx context.Context, method, url string) Ireq {
	if !strings.HasPrefix(url, "http") {
		url = fmt.Sprintf("https://%s", url)
	}
	r := &req{
		ctx:    ctx,
		method: method,
		url:    url,
	}
	return r
}

func (r *req) Header(h map[string]string) {
	r.header = h
}

func (r *req) Body(b any) {
	r.body = b
}

func (r *req) Query(m map[string]string) {
	r.query = m
}

func (r *req) Send(body any) error {
	var (
		bs  []byte
		err error
	)

	if r.body != nil {
		bs, err = json.Marshal(r.body)
		if err != nil {
			return fmt.Errorf("json.Marshal return err=%w", err)
		}
	}

	_req, err := http.NewRequestWithContext(r.ctx, r.method, r.url, bytes.NewReader(bs))
	if err != nil {
		return fmt.Errorf("new request failure,err=%w", err)
	}

	// set header
	_req = r.setReqHeader(_req)

	for k, v := range r.query {
		values := _req.URL.Query()
		values.Add(k, v)
		_req.URL.RawQuery = values.Encode()
	}

	tr := &http.Transport{
		//nolint:gosec
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}
	resp, err := c.Do(_req)
	if err != nil {
		return fmt.Errorf("http request err=%w", err)
	}

	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("invalid response,err=%w", err)
	}
	if body != nil {
		if err = json.Unmarshal(all, body); err != nil {
			return fmt.Errorf("json unmarshal return err=%w", err)
		}
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response status code is not 200, code=%d", resp.StatusCode)
	}

	return nil
}

func (r *req) setReqHeader(req *http.Request) *http.Request {
	for k, v := range r.header {
		req.Header.Set(k, v)
	}

	if req.Header.Get("x-request-id") == "" {
		req.Header.Set("x-request-id", uuid.New().String())
	}

	// set default header if method is post
	if r.method == http.MethodPost || r.method == http.MethodDelete {
		if req.Header.Get("Content-Type") == "" {
			req.Header.Set("Content-Type", "application/json")
		}

	}
	return req
}
