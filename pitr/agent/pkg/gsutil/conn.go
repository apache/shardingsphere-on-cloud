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

package gsutil

import (
	"database/sql"
	"fmt"
	_ "gitee.com/opengauss/openGauss-connector-go-pq"
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/internal/cons"
	"strings"
)

const defaultOGHost = "127.0.0.1"

type OpenGauss struct {
	db     *sql.DB
	user   string
	pwLen  int
	dbName string
}

func Open(user, password, dbName string, dbPort uint16) (*OpenGauss, error) {
	if strings.Trim(user, " ") == "" {
		return nil, fmt.Errorf("user is empty")
	}
	if strings.Trim(password, " ") == "" {
		return nil, fmt.Errorf("password is empty")
	}
	if strings.Trim(dbName, " ") == "" {
		return nil, fmt.Errorf("db name is empty")
	}

	connStr := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	db, err := sql.Open("opengauss", fmt.Sprintf(connStr, defaultOGHost, dbPort, user, password, dbName))
	if err != nil {
		efmt := "sql:open fail[user=%s,pwLen=%d,dbName=%s],err=%s,wrap=%w"
		return nil, fmt.Errorf(efmt, user, len(password), dbName, err, cons.DbConnectionFailed)
	}

	return &OpenGauss{
		db:     db,
		user:   user,
		pwLen:  len(password),
		dbName: dbName,
	}, nil
}

func (og *OpenGauss) Ping() error {
	if err := og.db.Ping(); err != nil {
		efmt := "db ping fail[user=%s,pwLen=%d,dbName=%s],err=%s,wrap=%w"
		return fmt.Errorf(efmt, og.user, og.pwLen, og.dbName, err, cons.DbConnectionFailed)
	}
	return nil
}
