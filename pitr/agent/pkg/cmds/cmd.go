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
	"os"
	"os/exec"
	"strings"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/logging"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/syncutils"
)

type Output struct {
	LineNo  uint32 // Start 1
	Message string
	Error   error
}

// AsyncExec Async exec a command
func AsyncExec(name string, args ...string) (chan *Output, error) {
	args = loadArgs(args...)

	c := "-c"
	args = append([]string{c}, args...)

	cmd := exec.Command(name, args...)
	logging.Debug(cmd.String())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("can not obtain stdout pipe for command[args=%+v]:%s", args, err)
	}
	if err = cmd.Start(); err != nil {
		return nil, fmt.Errorf("the command is err[args=%+v]:%s", args, err)
	}

	var (
		scanner = bufio.NewScanner(stdout)
		output  = make(chan *Output)
		index   = uint32(1)
	)
	go func() {
		if err = syncutils.NewRecoverFuncWithErrRet("", func() error {
			for scanner.Scan() {
				output <- &Output{
					LineNo:  index,
					Message: scanner.Text(),
					Error:   err,
				}
				index++
			}

			if err = scanner.Err(); err != nil {
				output <- &Output{
					LineNo: index,
					Error:  err,
				}
			}

			if err = cmd.Wait(); err != nil {
				if ee, ok := err.(*exec.ExitError); ok {
					output <- &Output{
						Error: fmt.Errorf("exec failure[ee=%s],wrap=%w", ee, cons.CmdOperateFailed),
					}
				} else {
					output <- &Output{
						Error: fmt.Errorf("cmd.Wait return err=%s,wrap=%w", err, cons.Internal),
					}
				}
			}
			return nil
		})(); err != nil {
			// only panic err
			output <- &Output{
				Error: err,
			}
		}
		close(output)
	}()

	return output, nil
}

// Exec exec a command
func Exec(name string, args ...string) (string, error) {
	args = loadArgs(args...)

	c := "-c"
	args = append([]string{c}, args...)

	cmd := exec.Command(name, args...)
	logging.Debug(cmd.String())

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("can not obtain stdout pipe for command[args=%+v]:%s", args, err)
	}
	if err = cmd.Start(); err != nil {
		return "", fmt.Errorf("the command is err[args=%+v]:%s", args, err)
	}

	reader, err := io.ReadAll(stdout)
	if err != nil {
		return "", fmt.Errorf("io.ReadAll return err=%w", err)
	}

	if err = cmd.Wait(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("exec failure[ee=%s,stdout=%s],wrap=%w", ee, string(reader), cons.CmdOperateFailed)
		}
		return "", cons.Internal
	}
	return string(reader), nil
}

// loadArgs if env is set, replace the command.
func loadArgs(args ...string) []string {
	if len(args) > 0 {
		arg := args[0]

		// check gs_probackup
		if v, ok := os.LookupEnv("gs_probackup"); ok && strings.HasPrefix(arg, "gs_probackup") {
			args[0] = strings.Replace(arg, "gs_probackup", v, 1)
		}

		// check gs_ctrl
		if v, ok := os.LookupEnv("gs_ctrl"); ok && strings.HasPrefix(arg, "gs_ctrl") {
			args[0] = strings.Replace(arg, "gs_ctrl", v, 1)
		}
	}
	return args
}
