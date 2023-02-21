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
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/cmd"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/logging"
)

func main() {
	prodConfig := zap.NewProductionConfig()
	prodConfig.Encoding = "console"
	prodConfig.DisableCaller = true
	prodConfig.DisableStacktrace = true
	prodConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	prodConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	prodConfig.EncoderConfig.ConsoleSeparator = "  "

	logger, err := prodConfig.Build(
		zap.WithCaller(false),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.FatalLevel),
	)
	if err != nil {
		panic(fmt.Errorf("an unknown error occured in the zap-log"))
	}
	logging.Init(logger)

	if err := cmd.Root.Execute(); err != nil {
		logging.Error(fmt.Sprintf("Err: %s", err))
	}
}
