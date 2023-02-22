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
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	strutil "github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/stringutil"
)

type (
	localStorage struct {
		rootDir   string
		backupDir string
	}

	ILocalStorage interface {
		init() error
		WriteByJSON(name string, contents anyStruct) error
		GenFilename(extn extension) string
	}

	anyStruct any

	extension string
)

const (
	ExtnJSON extension = "JOSN"
)

func NewLocalStorage(root string) (ILocalStorage, error) {
	ls := &localStorage{
		rootDir:   root,
		backupDir: fmt.Sprintf("%s/%s", root, "backup"),
	}

	if err := ls.init(); err != nil {
		return nil, err
	}

	return ls, nil
}

func (ls *localStorage) init() error {
	// root dir
	fi, err := os.Stat(ls.rootDir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(ls.rootDir, 0777); err != nil {
				return fmt.Errorf("create root dir failure,dir=%s,err=%s", ls.rootDir, err)
			}
		} else if os.IsExist(err) {
			if !fi.IsDir() {
				return fmt.Errorf("file has already exist,name=%s", ls.rootDir)
			}
		} else {
			return fmt.Errorf("failed to get file info,root dir=%s,err=%s", ls.rootDir, err)
		}
	}

	// backup dir
	fi, err = os.Stat(ls.backupDir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(ls.backupDir, 0777); err != nil {
				return fmt.Errorf("create backup dir failure,dir=%s,err=%s", ls.backupDir, err)
			}
		} else if os.IsExist(err) {
			if !fi.IsDir() {
				return fmt.Errorf("backup:file has already exist,name=%s", ls.backupDir)
			}
		} else {
			return fmt.Errorf("failed to get file info,backup dir=%s,err=%s", ls.backupDir, err)
		}
	}

	return nil
}

func (ls *localStorage) WriteByJSON(name string, contents anyStruct) error {
	if !strings.HasSuffix(name, ".json") {
		return fmt.Errorf("wrong file extension,file name is %s", name)
	}

	data, err := json.MarshalIndent(contents, "", "  ")
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s", ls.backupDir, name)
	fi, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file failure,file path is %s", path)
	}

	_, err = fi.Write(data)
	if err != nil {
		return fmt.Errorf("write to file failure,err=%s,data is %s", err, data)
	}

	return nil
}

/*
GenFilename gen a filename based on the file extension

	if extn is empty,return a postfix-free filename
	if extn=JSON,return the JSON filename like **.json
*/
func (ls *localStorage) GenFilename(extn extension) string {
	prefix := time.Now().UTC().Format("20060102150405")
	suffix := strutil.Random(8)

	switch extn {
	case ExtnJSON:
		return fmt.Sprintf("%s_%s.json", prefix, suffix)
	default:
		return fmt.Sprintf("%s_%s", prefix, suffix)
	}
}
