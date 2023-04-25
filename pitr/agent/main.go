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
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/handler"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/handler/middleware"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/pkg"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/logging"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/responder"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
)

const (
	debugLogLevel = "debug"
)

var (
	app *fiber.App
	log logging.ILog
)

var (
	logLevel      string
	port          string
	pgData        string
	tlsCrt        string
	tlsKey        string
	envSourceFile string
)

func init() {
	// 参数通过 flag 输入
	flag.StringVar(&logLevel, "log-level", "info", "optional:log level,option values:info or debug")
	flag.StringVar(&port, "port", "443", "HTTP service port")

	flag.StringVar(&tlsCrt, "tls-crt", "", "Require:TLS certificate file path")
	flag.StringVar(&tlsKey, "tls-key", "", "Require:TLS key file path")

	flag.StringVar(&pgData, "pgdata", "", "Optional:Get the value from cli flags or env")

	flag.StringVar(&envSourceFile, "env-source-file", "", "Optional:env source file path")
}

func main() {
	flag.Parse()

	if envSourceFile != "" {
		err := godotenv.Load(envSourceFile)
		if err != nil {
			panic(fmt.Errorf("load env source file error:%s", err.Error()))
		}
	}

	shell := os.Getenv("SHELL")
	if shell == "" {
		panic(fmt.Errorf("shell does not exist"))
	}

	if pgData == "" {
		pgData = os.Getenv("PGDATA")
		if pgData == "" {
			panic(fmt.Errorf("PGDATA:no database directory specified and environment variable PGDATA unset"))
		}
	}

	if _, err := os.Stat(pgData); os.IsNotExist(err) {
		panic(fmt.Errorf("PGDATA:%s the database directory does not exist", pgData))
	}

	pgData = strings.Trim(pgData, " ")
	if strings.HasSuffix(pgData, "/") {
		dirs := strings.Split(pgData, "/")
		dirs = dirs[0 : len(dirs)-1]
		pgData = strings.Join(dirs, "/")
	}

	if strings.Trim(tlsCrt, " ") == "" || strings.Trim(tlsKey, " ") == "" {
		panic(fmt.Errorf("lack of HTTPs certificate"))
	}

	if _, err := os.Stat(tlsCrt); os.IsNotExist(err) {
		panic(fmt.Errorf("TLS certificate file does not exist"))
	}
	if _, err := os.Stat(tlsKey); os.IsNotExist(err) {
		panic(fmt.Errorf("TLS key file does not exist"))
	}

	var level = zapcore.InfoLevel
	if logLevel == debugLogLevel {
		level = zapcore.DebugLevel
	}

	log = logging.Init(level)
	pkg.Init(shell, pgData, log)

	SetupApp()

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

func SetupApp() {
	app = fiber.New()

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
		r.Post("/diskspace", handler.DiskSpace)
	})

	// 404
	app.Use(func(ctx *fiber.Ctx) error {
		return responder.NotFound(ctx, "API not found")
	})
}

// Serve run a http server on the specified port.
func Serve(port string) error {

	//return app.Listen(":18080")
	return app.ListenTLS(fmt.Sprintf(":%s", port), tlsCrt, tlsKey)
}
