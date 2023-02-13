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
	"github.com/gofiber/fiber/v2"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/logging"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/responder"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/xerror"
)

func UniformErrResp(log logging.ILog) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		if err == nil {
			return nil
		}
		log.Fields(map[logging.FieldKey]string{
			logging.ErrorKey:  err.Error(),
			logging.RequestID: ctx.Get(cons.RequestID),
		}).Error("UniformErrResp:an error occurred")
		if e, b := xerror.FromError(err); b {
			return responder.Error(ctx, e)
		}
		return responder.Error(ctx, cons.Internal)
	}
}
