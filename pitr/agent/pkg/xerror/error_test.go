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

package xerror

import (
	"errors"
	"fmt"
	"testing"
)

func TestFromError(t *testing.T) {
	var e error
	_, b := FromError(e)
	if b {
		t.Fail()
	}

	e = errors.New("error type")
	_, b = FromError(e)
	if b {
		t.Fail()
	}

	daoErr := fmt.Errorf("xerror:%w", errors.New("not found"))
	serviceErr := fmt.Errorf("xerror:%s,wrap=%w", daoErr, &xerr{Code: 999})
	_, b = FromError(serviceErr)
	if !b {
		t.Fail()
	}

	_, b = FromError(serviceErr)
	if !b {
		t.Fail()
	}
}
