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
	"context"
	"fmt"
	"net/http"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/httputils"
)

type agentServer struct {
	addr string

	_apiBackup      string
	_apiRestore     string
	_apiShowDetail  string
	_apiShowList    string
	_apiDiskspace   string
	_apiHealthCheck string
}

type IAgentServer interface {
	CheckStatus(in *model.HealthCheckIn) error
	Backup(in *model.BackupIn) (string, error)
	DeleteBackup(in *model.DeleteBackupIn) error
	Restore(in *model.RestoreIn) error
	ShowDetail(in *model.ShowDetailIn) (*model.BackupInfo, error)
	ShowList(in *model.ShowListIn) ([]model.BackupInfo, error)
	ShowDiskSpace(in *model.DiskSpaceIn) (*model.DiskSpaceInfo, error)
}

var _ IAgentServer = (*agentServer)(nil)

func NewAgentServer(addr string) IAgentServer {
	return &agentServer{
		addr: addr,

		_apiBackup:      "/api/backup",
		_apiRestore:     "/api/restore",
		_apiShowDetail:  "/api/show",
		_apiShowList:    "/api/show/list",
		_apiDiskspace:   "/api/diskspace",
		_apiHealthCheck: "/api/healthz",
	}
}

type CommonOutResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// CheckStatus check agent server is alive
func (as *agentServer) CheckStatus(in *model.HealthCheckIn) error {
	url := fmt.Sprintf("%s%s", as.addr, as._apiHealthCheck)

	out := &CommonOutResp{}
	r := httputils.NewRequest(context.Background(), http.MethodPost, url)
	r.Body(in)

	if err := r.Send(out); err != nil {
		return xerr.NewUnknownErr(url, in, nil, err)
	}

	if out.Code != 0 {
		return xerr.NewAgentServerErr(out.Code, out.Msg)
	}

	return nil
}

func (as *agentServer) Backup(in *model.BackupIn) (string, error) {
	url := fmt.Sprintf("%s%s", as.addr, as._apiBackup)

	out := &model.BackupOutResp{}
	r := httputils.NewRequest(context.Background(), http.MethodPost, url)
	r.Body(in)

	if err := r.Send(out); err != nil {
		return "", xerr.NewUnknownErr(url, in, out, err)
	}

	if out.Code != 0 {
		return "", xerr.NewAgentServerErr(out.Code, out.Msg)
	}

	return out.Data.ID, nil
}

// nolint:dupl
func (as *agentServer) Restore(in *model.RestoreIn) error {
	url := fmt.Sprintf("%s%s", as.addr, as._apiRestore)

	out := &model.RestoreResp{}
	r := httputils.NewRequest(context.Background(), http.MethodPost, url)
	r.Body(in)

	if err := r.Send(out); err != nil {
		return xerr.NewUnknownErr(url, in, out, err)
	}

	if out.Code != 0 {
		return xerr.NewAgentServerErr(out.Code, out.Msg)
	}

	return nil
}

func (as *agentServer) ShowDetail(in *model.ShowDetailIn) (*model.BackupInfo, error) {
	url := fmt.Sprintf("%s%s", as.addr, as._apiShowDetail)

	out := &model.BackupDetailResp{}
	r := httputils.NewRequest(context.Background(), http.MethodPost, url)
	r.Body(in)

	if err := r.Send(out); err != nil {
		return nil, xerr.NewUnknownErr(url, in, out, err)
	}

	if out.Code != 0 {
		return nil, xerr.NewAgentServerErr(out.Code, out.Msg)
	}

	return &out.Data, nil
}

func (as *agentServer) ShowList(in *model.ShowListIn) ([]model.BackupInfo, error) {
	url := fmt.Sprintf("%s%s", as.addr, as._apiShowList)

	out := &model.BackupListResp{}
	r := httputils.NewRequest(context.Background(), http.MethodPost, url)
	r.Body(in)

	if err := r.Send(out); err != nil {
		return nil, xerr.NewUnknownErr(url, in, out, err)
	}

	if out.Code != 0 {
		return nil, xerr.NewAgentServerErr(out.Code, out.Msg)
	}

	return out.Data, nil
}

func (as *agentServer) ShowDiskSpace(in *model.DiskSpaceIn) (*model.DiskSpaceInfo, error) {
	url := fmt.Sprintf("%s%s", as.addr, as._apiDiskspace)

	out := &model.DiskSpaceInfo{}
	r := httputils.NewRequest(context.Background(), http.MethodPost, url)
	r.Body(in)

	if err := r.Send(out); err != nil {
		return nil, xerr.NewUnknownErr(url, in, out, err)
	}

	if out.Code != 0 {
		return nil, xerr.NewAgentServerErr(out.Code, out.Msg)
	}

	return out, nil
}

// nolint:dupl
func (as *agentServer) DeleteBackup(in *model.DeleteBackupIn) error {
	url := fmt.Sprintf("%s%s", as.addr, as._apiBackup)

	out := &model.DeleteBackupOut{}
	r := httputils.NewRequest(context.Background(), http.MethodDelete, url)
	r.Body(in)

	if err := r.Send(out); err != nil {
		return xerr.NewUnknownErr(url, in, out, err)
	}

	if out.Code != 0 {
		return xerr.NewAgentServerErr(out.Code, out.Msg)
	}

	return nil
}
