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

package syncutils

import (
	"fmt"
)

func NewRecoverFuncWithErrRet(msg string, fn func() error) func() (err error) {
	return func() error {
		defer func() {
			r := recover()
			if r != nil {
				if err, ok := r.(error); ok {
					err = fmt.Errorf("NewRecoverFuncWithErrRet[msg=%s],err=%s", msg, err) //nolint
				} else {
					err = fmt.Errorf("NewRecoverFuncWithErrRet[msg=%s],recover msg=%+v", msg, r) //nolint
				}
			}
		}()

		return fn()
	}
}
