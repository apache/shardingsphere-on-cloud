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

func HealthCheck(ctx *fiber.Ctx) error {
	in := &view.HealthCheckIn{}
	if err := ctx.BodyParser(in); err != nil {
		return fmt.Errorf("body parse err: %s, wrap: %w", err, cons.BodyParseFailed)
	}

	if err := pkg.OG.Auth(in.Username, in.Password, in.DBName, in.DBPort); err != nil {
		efmt := "pkg.OG.Auth failure[un=%s,pw.len=%d,db=%s], err wrap: %w"
		return fmt.Errorf(efmt, in.Username, len(in.Password), in.DBName, err)
	}

	// check schema if needed
	if in.Schema != "" {
		if err := pkg.OG.CheckSchema(in.Username, in.Password, in.DBName, in.DBPort, in.Schema); err != nil {
			return fmt.Errorf("pkg.OG.CheckSchema return err: %s, err wrap: %w", err, err)
		}
	}

	return responder.Success(ctx, "")
}
