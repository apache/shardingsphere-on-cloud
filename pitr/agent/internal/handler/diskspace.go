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
	"os"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/handler/view"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/cmds"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/responder"

	"github.com/gofiber/fiber/v2"
)

func DiskSpace(ctx *fiber.Ctx) error {
	in := &view.DiskSpaceIn{}

	if err := ctx.BodyParser(in); err != nil {
		return fmt.Errorf("body parse err: %s, wrap: %w", err, cons.BodyParseFailed)
	}

	if err := in.Validate(); err != nil {
		return fmt.Errorf("invalid parameter, err wrap: %w", err)
	}

	if err := os.MkdirAll(in.DiskPath, 0755); err != nil {
		return fmt.Errorf("mkdir [%s] failure, err wrap: %w", in.DiskPath, err)
	}

	// show disk space
	cmd := fmt.Sprintf("df -h %s", in.DiskPath)
	output, err := cmds.Exec(os.Getenv("SHELL"), cmd)
	if err != nil {
		return fmt.Errorf("exec cmd [%s] failure, err wrap: %w", cmd, err)
	}

	return responder.Success(ctx, output)
}
