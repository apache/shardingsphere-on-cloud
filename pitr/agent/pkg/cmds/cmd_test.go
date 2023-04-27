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

package cmds

import (
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	sh = "/bin/sh"
)

var _ = Describe("Commands", func() {
	Context("AsyncExec", func() {
		It("ping", func() {
			output, err := AsyncExec(sh, "ping 127.0.0.1")
			Expect(err).To(BeNil())

			for i := uint32(0); i < 10; i++ {
				select {
				case out, ok := <-output:
					Expect(out.LineNo).To(Equal(i + 1))
					Expect(out.Error).To(BeNil())
					Expect(ok).To(Equal(true))
				}
			}
		})
	})

	Context("Exec", func() {
		It("echo", func() {
			output, err := Exec(sh, "sleep 1;echo 10;sleep 1;echo 20;")
			Expect(err).To(BeNil())
			Expect(output).To(Equal("10\n20\n"))
		})
	})

	Context("test loadArgs", func() {
		It("test reload gs_probackup", func() {
			newGS := "/bin/new/gs_probackup"
			err := os.Setenv("gs_probackup", newGS)
			Expect(err).To(BeNil())

			defer os.Unsetenv("gs_probackup")

			args := loadArgs("gs_probackup", "backup", "-B", "/tmp", "-b", "FULL", "-D", "/tmp")
			Expect(args[0]).To(Equal(newGS))

			args = loadArgs("gs_probackup backup -B /tmp -b FULL -D /tmp")
			Expect(strings.HasPrefix(args[0], newGS)).To(Equal(true))
		})
	})
})
