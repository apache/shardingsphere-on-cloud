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
	InvalidHTTPHeader      = xerror.New(10001, "Invalid http header.")
	DataNotFound           = xerror.New(10002, "Data not found.")
	CmdOperateFailed       = xerror.New(10003, "Command operate failed.")
	BackupPathAlreadyExist = xerror.New(10004, "The backup path already exists.")
	NoPermission           = xerror.New(10005, "No permission to operate.")
	InstanceAlreadyExist   = xerror.New(10006, "The instance already exist.")
	InstanceNotExist       = xerror.New(10007, "The instance not exist.")
	StartOpenGaussFailed   = xerror.New(10008, "Failed to start opengauss.")
	StopOpenGaussFailed    = xerror.New(10009, "Failed to stop opengauss.")
	RestoreFailed          = xerror.New(10010, "Failed to restore opengauss.")
	InvalidDBPort          = xerror.New(10011, "Invalid dn port.")
	MissingUsername        = xerror.New(10012, "Missing username")
	MissingPassword        = xerror.New(10013, "Missing password.")
	MissingDnBackupPath    = xerror.New(10014, "Missing dn backup path.")
	InvalidDnThreadsNum    = xerror.New(10015, "Invalid dn threads num.")
	MissingDnBackupMode    = xerror.New(10016, "Missing dn backup mode.")
	InvalidDnBackupMode    = xerror.New(10017, "Invalid dn backup mode.")
	MissingInstance        = xerror.New(10018, "Missing instance.")
	MissingDnBackupID      = xerror.New(10019, "Missing dn backup id.")
	BodyParseFailed        = xerror.New(10020, "Invalid http request body.")
	MissingDBName          = xerror.New(10021, "Missing db name.")
	DBConnectionFailed     = xerror.New(10022, "Database connection failed.")
	UnmatchBackupID        = xerror.New(10023, "Unmatch any backup id.")
	InvalidPgDataDir       = xerror.New(10024, "Invalid PGDATA dir.")
	UnknownOgStatus        = xerror.New(10025, "Unknown openGauss status.")
	MvPgDataToTempFailed   = xerror.New(10026, "Move pgdata dir to temp failed.")
	MvTempToPgDataFailed   = xerror.New(10027, "Move temp dir to pgdata failed.")
	CleanPgDataTempFailed  = xerror.New(10028, "Clean pgdata temp dir failed.")
	MissingDiskPath        = xerror.New(10029, "Missing disk path.")
)
