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

package logging

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var l *ZapLogger

func Log() ILog {
	return l
}

func Init(level zapcore.Level) ILog {
	prodConfig := zap.NewProductionConfig()
	prodConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	prodConfig.Level = zap.NewAtomicLevelAt(level)
	logger, err := prodConfig.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.FatalLevel),
	)
	if err != nil {
		panic(fmt.Errorf("an unknown error occurred in the zap-log"))
	}
	l = &ZapLogger{logger: logger}
	return l
}

func Field(k FieldKey, v string) ILog {
	//nolint:exhaustive
	m := map[FieldKey]string{k: v}
	for k, v := range l.fields {
		m[k] = v
	}
	return &ZapLogger{logger: l.logger, fields: m}
}

func Fields(m map[FieldKey]string) ILog {
	if m == nil {
		m = map[FieldKey]string{}
	}
	for k, v := range l.fields {
		m[k] = v
	}
	return &ZapLogger{logger: l.logger, fields: m}
}

func Debug(s string) {
	l.logger.Debug(s, fields(l)...)
}

func Info(s string) {
	l.logger.Info(s, fields(l)...)
}

func Warn(s string) {
	l.logger.Warn(s, fields(l)...)
}

func Error(s string) {
	l.logger.Error(s, fields(l)...)
}

func Panic(s string) {
	l.logger.Panic(s, fields(l)...)
}
