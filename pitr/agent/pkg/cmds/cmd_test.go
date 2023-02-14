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
	"fmt"
	"testing"
)

const (
	sh = "/bin/sh"
)

var backup = "gs_probackup backup -B /home/omm/data --instance=ins-default-0 -b full -D /data/opengauss/3.1.1/data/single_node/  2>&1"
var ping = "ping www.baidu.com"

func TestCommand(t *testing.T) {
	output, err := Commands(sh, backup)
	if err != nil {
		t.Fatal(err)
	}

	for {
		select {
		case out, ok := <-output:
			if ok {
				if out.Error != nil {
					fmt.Println(out.LineNo, "\t", out.Error.Error())
				} else {
					fmt.Println(out.LineNo, "\t", out.Message)
				}
			} else {
				return
			}
		}
	}
}
