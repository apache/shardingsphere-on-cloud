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

package responder

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/xerror"
)

var (
	unknownErrCode = -1
	unknownErrMsg  = "Unknown error"

	successCode = 0
	successMsg  = "Success"
)

type resp struct {
	Code int    `json:"code" validate:"required"`
	Msg  string `json:"msg" validate:"required"`
	Data any    `json:"data"`
}

func Success(ctx *fiber.Ctx, data any) error {
	ctx.Status(http.StatusOK)
	return ctx.JSON(&resp{
		Code: successCode,
		Msg:  successMsg,
		Data: data,
	})
}

func Error(ctx *fiber.Ctx, e error) error {
	if e == nil {
		return xerror.New(unknownErrCode, unknownErrMsg)
	}
	ctx.Status(http.StatusOK)
	err, ok := xerror.FromError(e)
	if ok {
		return ctx.JSON(&resp{
			Code: err.Code,
			Msg:  err.Msg,
		})
	}
	return ctx.JSON(&resp{
		Code: unknownErrCode,
		Msg:  unknownErrMsg,
	})
}

func RawError(ctx *fiber.Ctx, e error) error {
	if e == nil {
		return xerror.New(unknownErrCode, unknownErrMsg)
	}
	ctx.Status(http.StatusOK)
	return ctx.JSON(&resp{
		Code: unknownErrCode,
		Msg:  e.Error(),
	})
}

func ErrorWithData(ctx *fiber.Ctx, e error, data any) error {
	if e == nil {
		return xerror.New(unknownErrCode, unknownErrMsg)
	}
	ctx.Status(http.StatusOK)

	if err, ok := xerror.FromError(e); ok {
		return ctx.JSON(&resp{
			Code: err.Code,
			Msg:  err.Msg,
			Data: data,
		})
	}
	return ctx.JSON(&resp{
		Code: unknownErrCode,
		Msg:  unknownErrMsg,
		Data: data,
	})
}

func NotFound(ctx *fiber.Ctx, msg string) error {
	ctx.Status(http.StatusNotFound)
	return ctx.JSON(&resp{
		Code: successCode,
		Msg:  msg,
	})
}
