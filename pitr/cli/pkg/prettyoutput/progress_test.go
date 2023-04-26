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

package prettyoutput_test

import (
	"time"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/prettyoutput"
	"github.com/jedib0t/go-pretty/v6/progress"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Progress", func() {
	It("test style", func() {
		pw := prettyoutput.NewPW(5)
		go pw.Render()

		for i := 0; i < 5; i++ {
			go func() {
				tracker := &progress.Tracker{
					Message: "test",
					Total:   10,
					Units:   progress.UnitsDefault,
				}

				pw.AppendTracker(tracker)
				ticker := time.Tick(time.Millisecond * 100)
				for !tracker.IsDone() {
					for range ticker {
						tracker.Increment(1)
					}
				}
			}()
		}

		time.Sleep(time.Millisecond * 100)
		for pw.IsRenderInProgress() {
			time.Sleep(time.Millisecond * 100)
		}
	})
})
