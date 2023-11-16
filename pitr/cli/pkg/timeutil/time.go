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

package timeutil

import "time"

const (
	unifiedTimeFormat = "2006-01-02 15:04:05"
)

type atime struct {
	time.Time
}

func Now() atime {
	return atime{time.Now()}
}

func (t atime) Add(d time.Duration) atime {
	_ = t.Time.Add(d)
	return t
}

func (t atime) String() string {
	return UnifiedTimeFormat(t.Time)
}

func (t atime) Unit() int64 {
	return UnixTimestampFormat(t.Time)
}

func Init() string {
	return ""
}

func UnifiedTimeFormat(t time.Time) string {
	return t.Format(unifiedTimeFormat)
}

func UnixTimestampFormat(t time.Time) int64 {
	return t.Unix()
}
