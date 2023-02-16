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
	"encoding/json"
	"fmt"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/pkg/model"

	"github.com/dlclark/regexp2"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/cmds"
)

type openGauss struct {
	shell string
}

const (
	_backupFmt = "gs_probackup backup --backup-path=%s --instance=%s --backup-mode=%s --pgdata=%s 2>&1"
	_showFmt   = "gs_probackup show --instance=%s --backup-path=%s --backup-id=%s --format=json 2>&1 "
)

func (og *openGauss) AsyncBackup(backupPath, instanceName, backupMode, pgData string) (string, error) {
	cmd := fmt.Sprintf(_backupFmt, backupPath, instanceName, backupMode, pgData)
	outputs, err := cmds.AsyncExec(og.shell, cmd)
	if err != nil {
		return "", fmt.Errorf("cmds.AsyncExec[shell=%s,cmd=%s] return err=%w", og.shell, cmd, err)
	}

	for output := range outputs {
		if output.Error != nil {
			return "", fmt.Errorf("output.Error[%w] is not nil", output.Error)
		}

		// get the backup id from the first line
		bid, err := og.getBackupID(output.Message)
		if err != nil {
			return "", fmt.Errorf("og.getBackupID[source=%s] return err=%w", output.Message, err)
		}
		//ignore other output
		go og.ignore(outputs)
		return bid, nil
	}
	return "", fmt.Errorf("unknow err")
}

func (og *openGauss) ShowBackupDetail(backupPath, instanceName, backupID string) (*model.Backup, error) {
	cmd := fmt.Sprintf(_showFmt, instanceName, backupPath, backupID)
	output, err := cmds.Exec(og.shell, cmd)
	if err != nil {
		return nil, fmt.Errorf("cmds.Exec[shell=%s,cmd=%s] return err=%w", og.shell, cmd, err)
	}

	var list []model.BackupList
	if err = json.Unmarshal([]byte(output), &list); err != nil {
		return nil, fmt.Errorf("json.Unmarshal[output=%s] return err=%s,wrap=%w", output, err, cons.Internal)
	}

	for _, ins := range list {
		if ins.Instance == instanceName {
			if len(ins.List) == 0 {
				return nil, fmt.Errorf("instance[name=%s],backupList[v=%+v],err=%w", ins.Instance, list, cons.DataNotFound)
			}

			return &ins.List[0], nil
		}
	}

	return nil, fmt.Errorf("backupList[v=%+v],err=%w", list, cons.DataNotFound)
}

func (og *openGauss) ignore(outputs chan *cmds.Output) {
	defer func() {
		_ = recover()
	}()

	for range outputs {
		//ignore all
	}
	//outputs closed
}

func (og *openGauss) getBackupID(msg string) (string, error) {
	re := regexp2.MustCompile("(?<=backup ID:\\s+)\\w+(?=,)", 0)
	match, err := re.FindStringMatch(msg)
	return match.String(), err
}
