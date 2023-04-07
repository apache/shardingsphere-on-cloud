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
	Send(body any) (int, error)
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

func (r *req) Send(body any) (int, error) {
	var (
		bs  []byte
		err error
	)

	if r.body != nil {
		bs, err = json.Marshal(r.body)
		if err != nil {
			return -1, fmt.Errorf("json.Marshal return err=%w", err)
		}
	}

	_req, err := http.NewRequestWithContext(r.ctx, r.method, r.url, bytes.NewReader(bs))
	if err != nil {
		return -1, fmt.Errorf("new request failure,err=%w", err)
	}

	for k, v := range r.header {
		_req.Header.Set(k, v)
	}

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
		return -1, fmt.Errorf("http request err=%w", err)
	}

	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, fmt.Errorf("invalid response,err=%w", err)
	}
	if body != nil {
		if err = json.Unmarshal(all, body); err != nil {
			return -1, fmt.Errorf("json unmarshal return err=%w", err)
		}
	}

	return resp.StatusCode, nil
}
