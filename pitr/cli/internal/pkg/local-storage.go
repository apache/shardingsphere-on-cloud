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
	"fmt"
	"os"
)

type (
	localStorage struct{}

	ILocalStorage interface {
		Init(dirName string) error
	}
)

func (ls *localStorage) Init(root string) error {
	// root dir
	fi, err := os.Stat(root)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(root, 0777); err != nil {
				return fmt.Errorf("create root dir failure,dir=%s,err=%s", root, err)
			}
		} else if os.IsExist(err) {
			if !fi.IsDir() {
				return fmt.Errorf("file has already exist,name=%s", root)
			}
		} else {
			return fmt.Errorf("failed to get file info,root dir=%s,err=%s", root, err)
		}
	}

	// backup dir
	backup := fmt.Sprintf("%s/%s", root, "backup")
	fi, err = os.Stat(backup)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(backup, 0777); err != nil {
				return fmt.Errorf("create root dir failure,dir=%s,err=%s", backup, err)
			}
		} else if os.IsExist(err) {
			if !fi.IsDir() {
				return fmt.Errorf("file has already exist,name=%s", backup)
			}
		} else {
			return fmt.Errorf("failed to get file info,root dir=%s,err=%s", backup, err)
		}
	}

	return nil
}
