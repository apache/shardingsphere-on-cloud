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

package strutil

import (
	"crypto/rand"
	"strconv"
)

const (
	charSet  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charSize = len(charSet)
)

const (
	digitSet  = "1234567890"
	digitSize = len(digitSet)
)

func Random(n uint) string {
	bs := make([]byte, n)
	for i := uint(0); i < n; i++ {
		_, _ = rand.Read(bs[i : i+1])
	}
	for k, v := range bs {
		bs[k] = charSet[v%byte(charSize)]
	}
	return string(bs)
}

func RandomInt(n uint) int64 {
	bs := make([]byte, n)
	_, _ = rand.Read(bs[0:1])
	for i := uint(1); i < n; i++ {
		_, _ = rand.Read(bs[i : i+1])
	}
	for k, v := range bs {
		bs[k] = digitSet[v%byte(digitSize)]
	}
	v, _ := strconv.ParseInt(string(bs), 10, 64)
	return v
}

func RandomUint(n uint) uint64 {
	return uint64(RandomInt(n))
}
