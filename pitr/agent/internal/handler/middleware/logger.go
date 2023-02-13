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

package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/logging"
)

func Logger(log logging.ILog) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var (
			start = time.Now()
		)
		err := ctx.Next()
		m := map[logging.FieldKey]string{
			logging.Duration:   fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
			logging.Path:       ctx.Route().Path,
			logging.RequestUri: string(ctx.Request().RequestURI()),
			logging.RequestID:  ctx.Get(cons.RequestID),
			logging.HttpStatus: fmt.Sprintf("%d", ctx.Response().StatusCode()),
			logging.HttpMethod: ctx.Method(),
		}
		if err != nil {
			m[logging.ErrorKey] = err.Error()
		}
		log.Fields(m).Info("logger-middleware")
		return err
	}
}

// AccessLog logging Access log.
func AccessLog(log logging.ILog) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Fields(map[logging.FieldKey]string{
			logging.Path:       ctx.Route().Path,
			logging.RequestUri: string(ctx.Request().RequestURI()),
			logging.RequestID:  ctx.Get(cons.RequestID),
			logging.HttpMethod: ctx.Method(),
		}).Info("Access log")
		return ctx.Next()
	}
}
