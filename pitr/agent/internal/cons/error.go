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

package cons

import (
	"github.com/apache/shardingsphere-on-cloud/pitr/agent/pkg/xerror"
)

var (
	Internal               = xerror.New(10000, "Internal error.")
	InvalidHttpHeader      = xerror.New(10001, "Invalid http header.")
	DataNotFound           = xerror.New(10002, "Data not found.")
	CmdOperateFailed       = xerror.New(10003, "Command operate failed.")
	BackupPathAlreadyExist = xerror.New(10004, "The backup path already exists.")
	NoPermission           = xerror.New(10005, "No permission to operate.")
)
