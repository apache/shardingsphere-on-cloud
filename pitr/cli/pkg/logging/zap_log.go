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
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	fields map[FieldKey]string
	logger *zap.Logger
}

func NewLog(l *zap.Logger) ILog {
	return &ZapLogger{logger: l}
}

func (z *ZapLogger) Field(k FieldKey, v string) ILog {
	m := map[FieldKey]string{k: v}
	for k, v := range z.fields {
		m[k] = v
	}
	return &ZapLogger{logger: z.logger, fields: m}
}

func (z *ZapLogger) Fields(m map[FieldKey]string) ILog {
	if m == nil {
		m = map[FieldKey]string{}
	}
	for k, v := range z.fields {
		m[k] = v
	}
	return &ZapLogger{logger: z.logger, fields: m}
}

func (z *ZapLogger) Debug(s string) {
	z.logger.Debug(s, fields(z)...)
}

func (z *ZapLogger) Info(s string) {
	z.logger.Info(s, fields(z)...)
}

func (z *ZapLogger) Warn(s string) {
	z.logger.Warn(s, fields(z)...)
}

func (z *ZapLogger) Error(s string) {
	z.logger.Error(s, fields(z)...)
}

func (z *ZapLogger) Panic(s string) {
	z.logger.Panic(s, fields(z)...)
}

func fields(z *ZapLogger) []zap.Field {
	fields := make([]zap.Field, 0, len(z.fields))
	for k, v := range z.fields {
		fields = append(fields, zap.Field{
			Key:    k.String(),
			Type:   zapcore.StringType,
			String: v,
		})
	}
	return fields
}
