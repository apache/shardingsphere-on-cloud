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

package pkg

import (
	"fmt"
	"time"

	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("OpenGauss,requires opengauss environment", func() {
	Context("AsyncBackup & ShowBackupDetail", func() {
		It("One backup and show", func() {
			og := &openGauss{
				shell: "/bin/sh",
			}

			var (
				data     = "/home/omm/data"
				instance = "ins-default-0"
			)

			backupID, err := og.AsyncBackup(
				data,
				instance,
				"full",
				"/data/opengauss/3.1.1/data/single_node/",
			)

			Expect(err).To(BeNil())
			Expect(backupID).NotTo(BeEmpty())
			fmt.Println(fmt.Sprintf("BackupID:%s", backupID))

			// timeout 60s
			for i := 0; i < 60; i++ {
				backup, err := og.ShowBackupDetail(
					data,
					instance,
					backupID,
				)

				Expect(err).To(BeNil())
				Expect(backup).NotTo(BeNil())
				Expect(backup.ID).To(Equal(backupID))
				if backup.Status == "OK" {
					goto End
				}
				time.Sleep(time.Second)
			}
			Fail("Timeout[60s]")
		End:
		})
	})
})
