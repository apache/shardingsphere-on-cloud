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

package pkg

import (
	"database/sql"
	"fmt"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/gsutil"
)

type (
	shardingSphere struct {
		db *sql.DB
	}

	IShardingSphere interface{}
)

const (
	DefaultDbName = "postgres"
)

func NewShardingSphereProxy(user, password, dbName, host string, port uint16) (IShardingSphere, error) {
	db, err := gsutil.Open(user, password, dbName, host, port)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		efmt := "db ping fail[host=%s,port=%d,user=%s,pwLen=%d,dbName=%s],err=%s"
		return nil, fmt.Errorf(efmt, host, port, user, len(password), dbName, err)
	}
	return &shardingSphere{db: db}, nil
}
