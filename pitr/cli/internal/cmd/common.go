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

package cmd

import (
	"fmt"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
)

func validate(ls pkg.ILocalStorage, CSN, RecordID string) (*model.LsBackup, error) {
	var (
		bak *model.LsBackup
		err error
	)
	if CSN != "" {
		bak, err = ls.ReadByCSN(CSN)
		if err != nil {
			return bak, xerr.NewCliErr(fmt.Sprintf("read backup record by csn failed. err: %s", err))
		}
	}

	if RecordID != "" {
		bak, err = ls.ReadByID(RecordID)
		if err != nil {
			return bak, xerr.NewCliErr(fmt.Sprintf("read backup record by id failed. err: %s", err))
		}
	}
	return bak, nil
}
