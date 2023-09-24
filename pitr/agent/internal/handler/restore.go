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

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/handler/view"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/pkg"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/responder"

	"github.com/gofiber/fiber/v2"
)

func Restore(ctx *fiber.Ctx) (err error) {
	in := &view.RestoreIn{}

	if err = ctx.BodyParser(in); err != nil {
		err = fmt.Errorf("body parse err: %s, wrap: %w", err, cons.BodyParseFailed)
		return
	}

	if err = in.Validate(); err != nil {
		err = fmt.Errorf("invalid parameter, err wrap: %w", err)
		return
	}

	if err = pkg.OG.Auth(in.Username, in.Password, in.DBName, in.DBPort); err != nil {
		efmt := "pkg.OG.Auth failure[un=%s,pw.len=%d,db=%s], err wrap: %w"
		err = fmt.Errorf(efmt, in.Username, len(in.Password), in.DBName, err)
		return
	}

	// stop openGauss
	if err = pkg.OG.Stop(); err != nil {
		err = fmt.Errorf("stop openGauss failure, err wrap: %w", err)
		return
	}

	defer func() {
		if err != nil {
			startErr := pkg.OG.Start()
			if startErr != nil {
				err = fmt.Errorf("pkg.OG.Start() return err: %s, wrap: %w", startErr, err)
				return
			}
		}
	}()

	// move pgdata to temp
	if err = pkg.OG.MvPgDataToTemp(); err != nil {
		err = fmt.Errorf("pkg.OG.MvPgDataToTemp return err wrap: %w", err)
		return
	}

	var status = "restoring"
	defer func() {
		if status != "restore success" {
			mvErr := pkg.OG.MvTempToPgData()
			err = fmt.Errorf("resotre failre[err=%s], pkg.OG.MvTempToPgData return err wrap: %w", err, mvErr)
		}
	}()

	// restore data from backup
	if err = pkg.OG.Restore(in.DnBackupPath, in.Instance, in.DnBackupID); err != nil {
		efmt := "pkg.OG.Restore failure[path=%s,instance=%s,backupID=%s], err wrap: %w"
		err = fmt.Errorf(efmt, in.DnBackupPath, in.Instance, in.DnBackupID, err)
		status = "restore failure"
		return
	}
	status = "restore success"

	// clean temp
	if err = pkg.OG.CleanPgDataTemp(); err != nil {
		err = fmt.Errorf("pkg.OG.CleanPgDataTemp return err wrap: %w", err)
		return
	}

	if err = pkg.OG.Start(); err != nil {
		err = fmt.Errorf("pkg.OG.Start return err wrap: %w", err)
		return
	}

	return responder.Success(ctx, nil)
}
