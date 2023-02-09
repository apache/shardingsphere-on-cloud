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
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func command(name string, args ...string) error {
	c := "-c"
	args = append([]string{c}, args...)

	cmd := exec.Command(name, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("can not obtain stdout pipe for command[args=%+v]:%s", args, err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("the command is err[args=%+v]:%s", args, err)
	}

	reader := bufio.NewReader(stdout)

	index := 1
	for {
		line, err := reader.ReadString('\n')
		if io.EOF == err {
			break
		} else if err != nil {
			return fmt.Errorf("read string is err[args=%+v]:%s", args, err)
		}

		fmt.Print(index, "\t", line)
		index++
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("cmd wait is err[args=%+v]:%s", args, err)
	}

	return nil
}
