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

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/handler/view"

	"github.com/gofiber/fiber/v2"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"
)

func Show(ctx *fiber.Ctx) error {
	in := &view.ShowIn{}

	if err := ctx.BodyParser(in); err != nil {
		return fmt.Errorf("body parse err=%s,wrap=%w", err, cons.BodyParseFailed)
	}

	if err := in.Validate(); err != nil {
		return err
	}

	return ctx.JSON(in)
}
