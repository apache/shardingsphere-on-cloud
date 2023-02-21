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

package logging

type ILog interface {
	Field(k FieldKey, v string) ILog
	Fields(m map[FieldKey]string) ILog

	Debug(s string)
	Info(s string)
	Warn(s string)
	Error(s string)
	Panic(s string)
}

type FieldKey string

func (f FieldKey) String() string {
	return string(f)
}

const (
	ErrorKey   FieldKey = "error"
	RequestID  FieldKey = "requestID"
	Stack      FieldKey = "stack"
	Duration   FieldKey = "duration"
	Path       FieldKey = "path"       // original routing path
	RequestUri FieldKey = "requestUri" // http requesting uri
	HttpMethod FieldKey = "httpMethod"
	HttpStatus FieldKey = "httpStatus"
)
