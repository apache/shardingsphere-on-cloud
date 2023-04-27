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

package handler_test

import (
	"os"
	"testing"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/handler"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/logging"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/responder"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var (
	app  *fiber.App
	ctrl *gomock.Controller
)

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suite")
}

var _ = BeforeSuite(func() {
	// init log
	logging.Init(zap.DebugLevel)

	Expect(os.Setenv("SHELL", "/bin/bash")).To(Succeed())

	// init app
	app = fiber.New()
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return responder.Success(ctx, "pong")
	})

	app.Route("/api", func(r fiber.Router) {
		r.Post("/diskspace", handler.DiskSpace)
		r.Delete("/backup", handler.DeleteBackup)
	})
})
