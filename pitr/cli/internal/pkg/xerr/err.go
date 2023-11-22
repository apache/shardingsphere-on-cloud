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

package xerr

import "fmt"

type err struct {
	msg string
}

const (
	postErrFmt    = "httputils.NewRequest[url=%s, body=%v, out=%v] return err=%s, wrap=%w"
	httpErrFmt    = "httputils.NewRequest[method=%s, url=%s, body=%v, out=%v] return err=%s, wrap=%w"
	httpRawErrFmt = "err=%s"

	Unknown           = "Unknown error"
	InvalidHTTPStatus = "Invalid http status"
	NotFound          = "Not found"
	Http              = "Http request error"
)

func (e *err) Error() string {
	return e.msg
}

func NewCliErr(msg string) error {
	return &err{
		msg: msg,
	}
}

func NewUnknownErr(url string, in, out interface{}, err error) error {
	return fmt.Errorf(postErrFmt, url, in, out, err, NewCliErr(Unknown))
}

func NewHttpRequestErr(method, url string, in, out interface{}, err error) error {
	return fmt.Errorf(httpErrFmt, method, url, in, out, err, NewCliErr(Http))
}

func NewHttpRawRequestErr(err error) error {
	return fmt.Errorf(httpRawErrFmt, err)
}

func NewAgentServerErr(code int, msg string) error {
	return &err{
		msg: fmt.Sprintf("agent server err[code=%d, msg=%s]", code, msg),
	}
}
