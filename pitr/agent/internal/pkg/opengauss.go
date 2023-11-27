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

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/cmds"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/gsutil"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/logging"

	"github.com/dlclark/regexp2"
)

type (
	openGauss struct {
		shell      string
		pgData     string
		pgDataTemp string
		log        logging.ILog
	}

	IOpenGauss interface {
		AsyncBackup(backupPath, instanceName, backupMode string, threadsNum uint8, dbPort uint16) (string, error)
		ShowBackup(backupPath, instanceName, backupID string) (*model.Backup, error)
		Init(backupPath string) error
		AddInstance(backupPath, instance string) error
		DelInstance(backupPath, instance string) error
		DelBackup(backupPath, instance, backupID string) error
		Start() error
		Stop() error
		Status() (string, error)
		Restore(backupPath, instance, backupID string, threadsNum uint8) error
		ShowBackupList(backupPath, instanceName string) ([]*model.Backup, error)
		Auth(user, password, dbName string, dbPort uint16) error
		CheckSchema(user, password, dbName string, dbPort uint16, schema string) error
		MvTempToPgData() error
		MvPgDataToTemp() error
		CleanPgDataTemp() error
	}
)

var _ IOpenGauss = (*openGauss)(nil)

func NewOpenGauss(shell, pgData string, log logging.ILog) IOpenGauss {
	dirs := strings.Split(pgData, "/")
	dirs = append(dirs[0:len(dirs)-1], "temp")

	return &openGauss{
		shell:      shell,
		pgData:     pgData,
		pgDataTemp: strings.Join(dirs, "/"),
		log:        log,
	}
}

const (
	_backupFmt    = "gs_probackup backup --backup-path=%s --instance=%s --backup-mode=%s --pgdata=%s --threads=%d --pgport %d 2>&1"
	_showFmt      = "gs_probackup show --instance=%s --backup-path=%s --backup-id=%s --format=json 2>&1"
	_delBackupFmt = "gs_probackup delete --backup-path=%s --instance=%s --backup-id=%s 2>&1"
	_restoreFmt   = "gs_probackup restore --backup-path=%s --instance=%s --backup-id=%s --pgdata=%s --threads=%d 2>&1"

	_initFmt  = "gs_probackup init --backup-path=%s 2>&1"
	_rmDirFmt = "rm -r %s"

	_addInstanceFmt = "gs_probackup add-instance --backup-path=%s --instance=%s --pgdata=%s 2>&1"
	_delInstanceFmt = "gs_probackup del-instance --backup-path=%s --instance=%s 2>&1"

	_startOpenGaussFmt = "gs_ctl start --pgdata=%s"
	_stopOpenGaussFmt  = "gs_ctl stop --pgdata=%s"
	_statusGaussFmt    = "gs_ctl status --pgdata=%s"

	_showListFmt = "gs_probackup show --instance=%s --backup-path=%s --format=json 2>&1"

	_mvFmt = "mv %s %s"

	_CmdErrorFmt = "cmds.Exec[shell=%s,cmd=%s] return err wrap: %s"
)

func (og *openGauss) AsyncBackup(backupPath, instanceName, backupMode string, threadsNum uint8, dbPort uint16) (string, error) {
	var (
		bid string
		err error
	)
	cmd := fmt.Sprintf(_backupFmt, backupPath, instanceName, backupMode, og.pgData, threadsNum, dbPort)
	outputs, err := cmds.AsyncExec(og.shell, cmd)
	if err != nil {
		return "", fmt.Errorf("cmds.AsyncExec[shell=%s,cmd=%s] return err wrap: %w", og.shell, cmd, cons.CmdAsyncBackupFailed)
	}

	for output := range outputs {
		og.log.
			Field("backup_path", backupPath).
			Field("instance", instanceName).
			Field("backup_mode", backupMode).
			Field("pgdata", og.pgData).
			Debug(fmt.Sprintf("AsyncBackup output[lineNo=%d,msg=%s,err=%v]", output.LineNo, output.Message, output.Error))

		if output.Error != nil {
			og.log.Error(fmt.Sprintf("output.Error[%s] is not nil", output.Error))
			return "", output.Error
		}

		if strings.Contains(output.Message, "INFO: Backup start") {
			bid, err = og.getBackupID(output.Message)
			if err != nil {
				og.log.Error(fmt.Sprintf("og.getBackupID[source=%s] return err wrap: %s", output.Message, err))
				return "", err
			}
		}
	}
	return bid, nil //nolint
}

//nolint:dupl
func (og *openGauss) ShowBackup(backupPath, instanceName, backupID string) (*model.Backup, error) {
	cmd := fmt.Sprintf(_showFmt, instanceName, backupPath, backupID)
	list, err := og.showbackup(cmd, instanceName)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	return nil, err
}

func (og *openGauss) DelBackup(backupPath, instanceName, backupID string) error {
	cmd := fmt.Sprintf(_delBackupFmt, backupPath, instanceName, backupID)
	_, err := cmds.Exec(og.shell, cmd)
	if err != nil {
		og.log.Error(fmt.Sprintf(_CmdErrorFmt, og.shell, cmd, err))
		return err
	}
	return nil
}

func (og *openGauss) Init(backupPath string) error {
	cmd := fmt.Sprintf(_initFmt, backupPath)

	output, err := cmds.Exec(og.shell, cmd)
	og.log.Debug(fmt.Sprintf("Init output[msg=%s,err=%v]", output, err))

	if errors.Is(err, cons.CmdOperateFailed) {
		og.log.Error(fmt.Sprintf("init backup path failure,err: %s, wrap: %s", err, cons.BackupPathAlreadyExist))
		return err
	}
	if err != nil {
		og.log.Error(fmt.Sprintf(_CmdErrorFmt, og.shell, cmd, err))
		return err
	}

	return nil
}

func (og *openGauss) deinit(backupPath string) error {
	if !strings.HasPrefix(backupPath, "/home/omm/") {
		return cons.NoPermission
	}

	cmd := fmt.Sprintf(_rmDirFmt, backupPath)
	if _, err := cmds.Exec(og.shell, cmd); err != nil {
		og.log.Error(fmt.Sprintf(_CmdErrorFmt, og.shell, cmd, err))
		return err
	}
	return nil
}

func (og *openGauss) AddInstance(backupPath, instance string) error {
	cmd := fmt.Sprintf(_addInstanceFmt, backupPath, instance, og.pgData)

	output, err := cmds.Exec(og.shell, cmd)
	og.log.Debug(fmt.Sprintf("AddInstance[output=%s,err=%v]", output, err))

	if errors.Is(err, cons.CmdOperateFailed) {
		og.log.Error(fmt.Sprintf("add instance failure[output=%s], err: %s, wrap: %s", output, err, cons.InstanceAlreadyExist))
		return fmt.Errorf("add instance failure[output=%s], err: %s, wrap: %w", output, err, cons.InstanceAlreadyExist)
	}
	if err != nil {
		og.log.Error(fmt.Sprintf(_CmdErrorFmt, og.shell, cmd, cons.CmdAddInstanceFailed))
		return err
	}

	return nil
}

func (og *openGauss) DelInstance(backupPath, instance string) error {
	cmd := fmt.Sprintf(_delInstanceFmt, backupPath, instance)
	output, err := cmds.Exec(og.shell, cmd)
	og.log.Debug(fmt.Sprintf("DelInstance[output=%s,err=%v]", output, err))

	if errors.Is(err, cons.CmdOperateFailed) {
		og.log.Error(fmt.Sprintf("delete instance failure[output=%s], err: %s, wrap: %s", output, err, cons.InstanceNotExist))
		return err
	}
	if err != nil {
		og.log.Error(fmt.Sprintf(_CmdErrorFmt, og.shell, cmd, cons.CmdDelInstanceFailed))
		return err
	}
	return nil
}

//nolint:dupl
func (og *openGauss) Start() error {
	cmd := fmt.Sprintf(_startOpenGaussFmt, og.pgData)
	output, err := cmds.Exec(og.shell, cmd)
	og.log.Debug(fmt.Sprintf("Start openGauss[output=%s]", output))

	if errors.Is(err, cons.CmdOperateFailed) {
		og.log.Error(fmt.Sprintf("start openGauss failure[output=%s], err: %s, wrap: %s", output, err, cons.StartOpenGaussFailed))
		return err
	}
	if err != nil {
		og.log.Error(fmt.Sprintf("cmds.Exec[shell=%s,cmd=%s] output=%s return err wrap: %s", og.shell, cmd, output, cons.CmdStartOpenGaussFailed))
		return err
	}

	return nil
}

//nolint:dupl
func (og *openGauss) Stop() error {
	cmd := fmt.Sprintf(_stopOpenGaussFmt, og.pgData)
	output, err := cmds.Exec(og.shell, cmd)
	og.log.Debug(fmt.Sprintf("Stop openGauss[output=%s]", output))

	if errors.Is(err, cons.CmdOperateFailed) {
		og.log.Error(fmt.Sprintf("stop openGauss failure[output=%s], err: %s, wrap: %s", output, err, cons.StopOpenGaussFailed))
		return err
	}
	if err != nil {
		og.log.Error(fmt.Sprintf("cmds.Exec[shell=%s,cmd=%s] output=%s return err wrap: %s", og.shell, cmd, output, cons.CmdStopOpenGaussFailed))
		return err
	}

	return nil
}

/*
Status return openGauss server status:

		`Running`  return "Runnging",nil
	    `Stopped`   return "Stopped",nil

The others are abnormal states,return "" and error.
*/
func (og *openGauss) Status() (string, error) {
	cmd := fmt.Sprintf(_statusGaussFmt, og.pgData)
	output, err := cmds.Exec(og.shell, cmd)
	og.log.Debug(fmt.Sprintf("Status openGauss[output=%s]", output))

	if errors.Is(err, cons.CmdOperateFailed) {
		if strings.Contains(err.Error(), "no server running") {
			return "Stopped", nil
		}
		og.log.Error(fmt.Sprintf("get openGauss status failure[output=%s], err: %s, wrap: %s", output, err, cons.StopOpenGaussFailed))
		return "", err
	}
	if err != nil {
		og.log.Error(fmt.Sprintf("cmds.Exec[shell=%s,cmd=%s] output=%s return err wrap: %s", og.shell, cmd, output, cons.CmdStatusOpenGaussFailed))
		return "", err
	}

	if strings.Contains(output, "server is running") {
		return "Running", nil
	}

	return "", cons.UnknownOgStatus
}

// Restore TODO:Dependent environments require integration testing
func (og *openGauss) Restore(backupPath, instance, backupID string, threadsNum uint8) error {
	cmd := fmt.Sprintf(_restoreFmt, backupPath, instance, backupID, og.pgData, threadsNum)
	outputs, err := cmds.AsyncExec(og.shell, cmd)
	for output := range outputs {
		og.log.
			//nolint:exhaustive
			Fields(map[logging.FieldKey]string{
				"backup_path": backupPath,
				"instance":    instance,
				"backup_id":   backupID,
			}).
			Debug(fmt.Sprintf("Restore openGauss[lineNo=%d,msg=%s]", output.LineNo, output.Message))

		if err != nil {
			og.log.Error(fmt.Sprintf("cmds.AsyncExec[output=%s] return err: %s, wrap: %s", output.Message, output.Error, cons.RestoreFailed))
			return err
		}

		if output.Error != nil {
			og.log.Error(fmt.Sprintf("cmds.AsyncExec outputs: Error[%s] is not nil, wrap: %s", output.Error, cons.RestoreFailed))
			return output.Error
		}
	}
	return nil
}

//nolint:dupl
func (og *openGauss) showbackup(cmd, instanceName string) ([]*model.Backup, error) {
	output, err := cmds.Exec(og.shell, cmd)
	if err != nil {
		og.log.Error(fmt.Sprintf(_CmdErrorFmt, og.shell, cmd, cons.CmdShowBackupFailed))
		return nil, err
	}

	var list []*model.BackupList
	if err = json.Unmarshal([]byte(output), &list); err != nil {
		og.log.Error(fmt.Sprintf("json.Unmarshal[output=%s] return err: %s, wrap: %s", output, err, cons.JSONUnmarshalFailed))
		return nil, err
	}

	for _, ins := range list {
		if ins.Instance == instanceName {
			if len(ins.List) == 0 {
				og.log.Error(fmt.Sprintf("instance[name=%s], backupList[v=%+v], err wrap: %s", ins.Instance, list, cons.DataNotFound))
				return nil, err
			}

			return ins.List, nil
		}
	}

	og.log.Error(fmt.Sprintf("backupList[v=%+v], err wrap: %s", list, cons.DataNotFound))
	return nil, err
}

func (og *openGauss) ShowBackupList(backupPath, instanceName string) ([]*model.Backup, error) {
	cmd := fmt.Sprintf(_showListFmt, instanceName, backupPath)
	return og.showbackup(cmd, instanceName)
}

//nolint:unused
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
	if err != nil {
		og.log.Error(fmt.Sprintf("unmatch any backup id[msg=%s], err: %s", msg, err))
		return "", err
	}
	if match == nil || match.Length == 0 {
		og.log.Error(fmt.Sprintf("unmatch any backup id,match.lenght is 0, err wrap: %s", cons.UnmatchBackupID))
		return "", err
	}

	return match.String(), err
}

func (og *openGauss) Auth(user, password, dbName string, dbPort uint16) error {
	if strings.Trim(user, " ") == "" ||
		strings.Trim(password, " ") == "" ||
		strings.Trim(dbName, " ") == "" ||
		dbPort == 0 {
		return fmt.Errorf("invalid inputs[user=%s,password=%s,dbName=%s,dbPort=%d], err wrap: %w", user, password, dbName, dbPort, cons.MissingDBInformation)
	}

	_og, err := gsutil.Open(user, password, dbName, dbPort)
	if err != nil {
		og.log.Error(fmt.Sprintf("gsutil.Open failure,err=%s", err))
		return err
	}

	if err := _og.Ping(); err != nil {
		og.log.Error(fmt.Sprintf("ping openGauss fail[user=%s,pw length=%d,dbName=%s], err wrap: %s", user, len(password), dbName, err))
		return err
	}

	return nil
}

func (og *openGauss) MvPgDataToTemp() error {
	cmd := fmt.Sprintf(_mvFmt, og.pgData, og.pgDataTemp)
	_, err := cmds.Exec(og.shell, cmd)
	if errors.Is(err, cons.CmdOperateFailed) {
		og.log.Error(fmt.Sprintf("mv pgdata to temp dir failure, err: %s, wrap: %s", err, cons.MvPgDataToTempFailed))
		return err
	}
	if err != nil {
		og.log.Error(fmt.Sprintf(_CmdErrorFmt, og.shell, cmd, err))
		return err
	}

	return nil
}

func (og *openGauss) MvTempToPgData() error {
	cmd := fmt.Sprintf(_mvFmt, og.pgDataTemp, og.pgData)
	_, err := cmds.Exec(og.shell, cmd)
	if errors.Is(err, cons.CmdOperateFailed) {
		og.log.Error(fmt.Sprintf("mv temp to pgdata dir failure, err: %s, wrap: %s", err, cons.MvTempToPgDataFailed))
		return err
	}
	if err != nil {
		og.log.Error(fmt.Sprintf(_CmdErrorFmt, og.shell, cmd, err))
		return err
	}
	return nil
}

func (og *openGauss) CleanPgDataTemp() error {
	cmd := fmt.Sprintf(_rmDirFmt, og.pgDataTemp)
	_, err := cmds.Exec(og.shell, cmd)
	if errors.Is(err, cons.CmdOperateFailed) {
		og.log.Error(fmt.Sprintf("clean pgdata temp dir failure, err: %s, wrap: %s", err, cons.CleanPgDataTempFailed))
		return err
	}
	if err != nil {
		og.log.Error(fmt.Sprintf(_CmdErrorFmt, og.shell, cmd, err))
		return err
	}
	return nil
}

func (og *openGauss) CheckSchema(user, password, dbName string, dbPort uint16, schema string) error {
	_og, err := gsutil.Open(user, password, dbName, dbPort)
	if err != nil {
		og.log.Error(fmt.Sprintf("gsutil.Open failure, err wrap: %s", err))
		return err
	}

	if err := _og.CheckSchema(schema); err != nil {
		og.log.Error(fmt.Sprintf("check openGauss schema fail[user=%s,dbName=%s, schema=%s], err wrap: %s", user, dbName, schema, err))
		return err
	}
	return nil
}
