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

	"github.com/gofiber/fiber/v2"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/handler/view"
)

func Restore(ctx *fiber.Ctx) (err error) {
	in := &view.RestoreIn{}

	if err = ctx.BodyParser(in); err != nil {
		err = fmt.Errorf("body parse err=%s,wrap=%w", err, cons.BodyParseFailed)
		return
	}

	if err = in.Validate(); err != nil {
		err = fmt.Errorf("invalid parameter,err=%w", err)
		return
	}

	if err = pkg.OG.Auth(in.Username, in.Password, in.DbName, in.DbPort); err != nil {
		efmt := "pkg.OG.Auth failure[un=%s,pw.len=%d,db=%s],err=%w"
		err = fmt.Errorf(efmt, in.Username, len(in.Password), in.DbName, err)
		return
	}

	// stop openGauss
	if err = pkg.OG.Stop(); err != nil {
		err = fmt.Errorf("stop openGauss failure,err=%w", err)
		return
	}
	defer func() {
		if err != nil {
			err2 := pkg.OG.Start()
			if err2 != nil {
				err = fmt.Errorf("pkg.OG.Start() return err=%s,wrap=%w", err2, err)
				return
			}
		}
	}()

	if err = pkg.OG.MvPgDataToTemp(); err != nil {
		err = fmt.Errorf("pkg.OG.MvPgDataToTemp return err=%w", err)
		return
	}

	if err = pkg.OG.Restore(in.DnBackupPath, in.Instance, in.DnBackupId); err != nil {
		efmt := "pkg.OG.Restore failure[path=%s,instance=%s,backupID=%s],err=%w"
		err = fmt.Errorf(efmt, in.DnBackupPath, in.Instance, in.DnBackupId, err)

		err2 := pkg.OG.MvTempToPgData()
		err = fmt.Errorf("resotre failre[err=%s],pkg.OG.MvTempToPgData return err=%w", err, err2)

		return
	}

	if err = pkg.OG.CleanPgDataTemp(); err != nil {
		err = fmt.Errorf("pkg.OG.CleanPgDataTemp return err=%w", err)
		return
	}

	if err = pkg.OG.Start(); err != nil {
		err = fmt.Errorf("pkg.OG.Start return err=%w", err)
		return
	}

	if err = responder.Success(ctx, nil); err != nil {
		err = fmt.Errorf("responder failure,err=%s,wrap=%w", err, cons.Internal)
		return nil
	}
	return
}
