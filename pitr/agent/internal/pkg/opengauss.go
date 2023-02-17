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
	"errors"
	"fmt"
	"strings"

	"github.com/dlclark/regexp2"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/cmds"
)

type openGauss struct {
	shell string
}

const (
	_backupFmt    = "gs_probackup backup --backup-path=%s --instance=%s --backup-mode=%s --pgdata=%s 2>&1"
	_showFmt      = "gs_probackup show --instance=%s --backup-path=%s --backup-id=%s --format=json 2>&1"
	_delBackupFmt = "gs_probackup delete --backup-path=%s --instance=%s --backup-id=%s 2>&1"

	_initFmt   = "gs_probackup init --backup-path=%s 2>&1"
	_deinitFmt = "rm -r %s"

	_addInstanceFmt = "gs_probackup add-instance --backup-path=%s --instance=%s --pgdata=%s 2>&1"
	_delInstanceFmt = "gs_probackup del-instance --backup-path=%s --instance=%s 2>&1"

	_startOpenGaussFmt = "gs_ctl start --pgdata=%s"
	_stopOpenGaussFmt  = "gs_ctl stop --pgdata=%s"
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

func (og *openGauss) delBackup(backupPath, instanceName, backupID string) error {
	cmd := fmt.Sprintf(_delBackupFmt, backupPath, instanceName, backupID)
	_, err := cmds.Exec(og.shell, cmd)
	if err != nil {
		return fmt.Errorf("cmds.Exec[shell=%s,cmd=%s] return err=%w", og.shell, cmd, err)
	}
	return nil
}

func (og *openGauss) Init(backupPath string) error {
	cmd := fmt.Sprintf(_initFmt, backupPath)
	_, err := cmds.Exec(og.shell, cmd)
	// already exist and it's not empty
	if errors.Is(err, cons.CmdOperateFailed) {
		return cons.BackupPathAlreadyExist
	}
	if err != nil {
		return fmt.Errorf("cmds.Exec[shell=%s,cmd=%s] return err=%w", og.shell, cmd, err)
	}
	return nil
}

func (og *openGauss) deinit(backupPath string) error {
	if !strings.HasPrefix(backupPath, "/home/omm/") {
		return cons.NoPermission
	}

	cmd := fmt.Sprintf(_deinitFmt, backupPath)
	if _, err := cmds.Exec(og.shell, cmd); err != nil {
		return fmt.Errorf("cmds.Exec[shell=%s,cmd=%s] return err=%w", og.shell, cmd, err)
	}
	return nil
}

func (og *openGauss) AddInstance(backupPath, instancee, pgData string) error {
	cmd := fmt.Sprintf(_addInstanceFmt, backupPath, instancee, pgData)
	_, err := cmds.Exec(og.shell, cmd)
	// already exist and it's not empty
	if errors.Is(err, cons.CmdOperateFailed) {
		return cons.InstanceAlreadyExist
	}
	if err != nil {
		return fmt.Errorf("cmds.Exec[shell=%s,cmd=%s] return err=%w", og.shell, cmd, err)
	}
	return nil
}

func (og *openGauss) DelInstance(backupPath, instancee string) error {
	cmd := fmt.Sprintf(_delInstanceFmt, backupPath, instancee)
	_, err := cmds.Exec(og.shell, cmd)
	// already exist and it's not empty
	if errors.Is(err, cons.CmdOperateFailed) {
		return cons.InstanceNotExist
	}
	if err != nil {
		return fmt.Errorf("cmds.Exec[shell=%s,cmd=%s] return err=%w", og.shell, cmd, err)
	}
	return nil
}

func (og *openGauss) Start(pgData string) error {
	cmd := fmt.Sprintf(_startOpenGaussFmt, pgData)
	_, err := cmds.Exec(og.shell, cmd)
	// already exist and it's not empty
	if errors.Is(err, cons.CmdOperateFailed) {
		return cons.StartOpenGaussFailed
	}
	if err != nil {
		return fmt.Errorf("cmds.Exec[shell=%s,cmd=%s] return err=%w", og.shell, cmd, err)
	}
	return nil
}

func (og *openGauss) Stop(pgData string) error {
	cmd := fmt.Sprintf(_stopOpenGaussFmt, pgData)
	_, err := cmds.Exec(og.shell, cmd)
	// already exist and it's not empty
	if errors.Is(err, cons.CmdOperateFailed) {
		return cons.StopOpenGaussFailed
	}
	if err != nil {
		return fmt.Errorf("cmds.Exec[shell=%s,cmd=%s] return err=%w", og.shell, cmd, err)
	}
	return nil
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
