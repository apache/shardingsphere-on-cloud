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

package handler

import (
	"fmt"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/pkg"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/responder"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/handler/view"

	"github.com/gofiber/fiber/v2"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"
)

func Backup(ctx *fiber.Ctx) error {
	in := &view.BackupIn{}

	if err := ctx.BodyParser(in); err != nil {
		return fmt.Errorf("body parse err=%s,wrap=%w", err, cons.BodyParseFailed)
	}

	if err := in.Validate(); err != nil {
		return fmt.Errorf("invalid parameter,err=%w", err)
	}

	if err := pkg.OG.Auth(in.Username, in.Password, in.DbName, in.DbPort); err != nil {
		efmt := "pkg.OG.Auth failure[un=%s,pw.len=%d,db=%s],err=%w"
		return fmt.Errorf(efmt, in.Username, len(in.Password), in.DbName, err)
	}

	backupID, err := pkg.OG.AsyncBackup(in.DnBackupPath, in.Instance, in.DnBackupMode, 1)
	if err != nil {
		efmt := "pkg.OG.AsyncBackup[path=%s,instance=%s,mode=%s] failure,err=%w"
		return fmt.Errorf(efmt, in.DnBackupPath, in.Instance, in.DnBackupMode, err)
	}

	return responder.Success(ctx, view.BackupOut{
		ID: backupID,
	})
}
