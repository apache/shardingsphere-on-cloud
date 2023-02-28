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

package main

import (
    "flag"
    "fmt"
    "github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/handler"
    "github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/handler/middleware"
    "github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/pkg"
    "github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/logging"
    "github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/responder"
    "github.com/gofiber/fiber/v2"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
    "os/signal"
    "syscall"
)

const (
	debugLogLevel = "debug"
)

var (
	app *fiber.App
	log logging.ILog
)

var (
	logLevel string
	port     string
	pgData   string
	tlsCrt   string
	tlsKey   string
)

func init() {
	// 参数通过 flag 输入
	flag.StringVar(&logLevel, "log-level", "info", "optional:log level,option values:info or debug")
	flag.StringVar(&port, "port", "443", "HTTP service port")

	flag.StringVar(&tlsCrt, "tls-crt", "", "Require:TLS certificate file path")
	flag.StringVar(&tlsKey, "tls-key", "", "Require:TLS key file path")

	flag.StringVar(&pgData, "pgdata", "", "Optional:Get the value from cli flags or env")
}

func main() {
	flag.Parse()

	shell := os.Getenv("SHELL")
	if shell == "" {
		panic(fmt.Errorf("shell does not exist"))
	}
	pkg.Init(shell)

	if pgData == "" {
		pgData = os.Getenv("PGDATA")
		if pgData == "" {
			panic(fmt.Errorf("PGDATA:no database directory specified and environment variable PGDATA unset"))
		}
	}
	// todo tag-1
	//	if strings.Trim(tlsCrt, " ") == "" || strings.Trim(tlsKey, " ") == "" {
	//		panic(fmt.Errorf("lack of HTTPs certificate"))
	//	}

	var level = zapcore.InfoLevel
	if logLevel == debugLogLevel {
		level = zapcore.DebugLevel
	}

	prodConfig := zap.NewProductionConfig()
	prodConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	prodConfig.Level = zap.NewAtomicLevelAt(level)
	logger, err := prodConfig.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.FatalLevel),
	)
	if err != nil {
		panic(fmt.Errorf("an unknown error occured in the zap-log"))
	}

	log = logging.Init(logger)
	app = fiber.New()

	go func() {
		if err := Serve(port); err != nil {
			panic(err)
		}
	}()
	log.Info("app startup successfully.")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if app != nil {
		if err := app.Shutdown(); err != nil {
			log.Field(logging.ErrorKey, err.Error()).Error("http app closed failure")
		}
	}
	log.Info("app windup successfully.")
	log.Info("app has exited...")
}

// Serve run a http server on the specified port.
func Serve(port string) error {
	app.Use(
		middleware.Recover(logging.Log()),
		middleware.UniformErrResp(logging.Log()),
	)

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return responder.Success(ctx, "pong")
	})

	app.Use(
		middleware.AccessLog(logging.Log()),
		middleware.Logger(logging.Log()),
	)

	app.Route("/api", func(r fiber.Router) {
		r.Use(middleware.RequestIDChecker())

		r.Post("/backup", handler.Backup)
		r.Post("/restore", handler.Restore)
		r.Post("/show", handler.Show)
		r.Post("/show/list", handler.ShowList)
	})

	// 404
	app.Use(func(ctx *fiber.Ctx) error {
		return responder.NotFound(ctx, "API not found")
	})

	return app.Listen(":18080")
	// todo tag-1
	//	return app.ListenTLS(fmt.Sprintf(":%s", port), tlsCrt, tlsKey)
}
