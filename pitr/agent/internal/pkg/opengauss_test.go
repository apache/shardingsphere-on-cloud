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
	Context("AsyncBackup", func() {
		It("One backup", func() {
			og := &openGauss{
				shell: "/bin/sh",
			}
			backupID, err := og.AsyncBackup(
				"/home/omm/data",
				"ins-default-0",
				"full",
				"/data/opengauss/3.1.1/data/single_node/",
			)
            
			Expect(err).To(BeNil())
			Expect(backupID).NotTo(BeEmpty())
			fmt.Println(fmt.Sprintf("BackupID:%s", backupID))
			time.Sleep(time.Second * 10)
		})
	})
})
