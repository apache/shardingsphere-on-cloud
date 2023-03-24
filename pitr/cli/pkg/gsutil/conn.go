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
	"strings"

	_ "gitee.com/opengauss/openGauss-connector-go-pq"
)

func Open(user, password, dbName, host string, port uint16) (*sql.DB, error) {
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
	db, err := sql.Open("opengauss", fmt.Sprintf(connStr, host, port, user, password, dbName))
	if err != nil {
		efmt := "sql:open fail[host=%s,port=%d,user=%s,pwLen=%d,dbName=%s],err=%s"
		return nil, fmt.Errorf(efmt, host, port, user, len(password), dbName, err)
	}

	return db, nil
}
