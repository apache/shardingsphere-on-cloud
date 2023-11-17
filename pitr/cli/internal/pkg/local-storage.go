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
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"

	strutil "github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/stringutil"
)

type (
	localStorage struct {
		rootDir   string
		backupDir string
	}

	ILocalStorage interface {
		WriteByJSON(name string, contents *model.LsBackup) error
		GenFilename(extn Extension) string
		ReadAll() ([]*model.LsBackup, error)
		ReadByID(id string) (*model.LsBackup, error)
		ReadByCSN(csn string) (*model.LsBackup, error)
		ReadAllByCSN(csn string) ([]*model.LsBackup, error)
		DeleteByName(name string) error
		HideByName(name string) error
		DeleteByHidedName(name string) error
	}

	Extension string
)

const (
	ExtnJSON Extension = "JSON"
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

func DefaultRootDir() string {
	return fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".gs_pitr")
}

func (ls *localStorage) init() error {
	// root dir
	fi, err := os.Stat(ls.rootDir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(ls.rootDir, 0777); err != nil {
				return fmt.Errorf("create root dir failure. dir: %s, err: %s", ls.rootDir, err)
			}
		} else if os.IsExist(err) {
			if !fi.IsDir() {
				return fmt.Errorf("file has already exist. name: %s", ls.rootDir)
			}
		} else {
			return fmt.Errorf("failed to get file info. root dir: %s, err: %s", ls.rootDir, err)
		}
	}

	// backup dir
	fi, err = os.Stat(ls.backupDir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(ls.backupDir, 0777); err != nil {
				return fmt.Errorf("create backup dir failure. dir: %s, err: %s", ls.backupDir, err)
			}
		} else if os.IsExist(err) {
			if !fi.IsDir() {
				return fmt.Errorf("backup: file has already exist. name: %s", ls.backupDir)
			}
		} else {
			return fmt.Errorf("failed to get file info. backup dir: %s, err: %s", ls.backupDir, err)
		}
	}

	return nil
}

func (ls *localStorage) WriteByJSON(name string, contents *model.LsBackup) error {
	if !strings.HasSuffix(name, ".json") {
		return fmt.Errorf("wrong file extension. file name: %s", name)
	}

	data, err := json.MarshalIndent(contents, "", "  ")
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s", ls.backupDir, name)
	fi, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file failure. file path: %s", path)
	}

	_, err = fi.Write(data)
	if err != nil {
		return fmt.Errorf("write to file failure. err: %s, data: %s", err, data)
	}

	return nil
}

func (ls *localStorage) ReadAll() ([]*model.LsBackup, error) {
	entries, err := os.ReadDir(ls.backupDir)
	if err != nil {
		return nil, xerr.NewCliErr(fmt.Sprintf("read the dir[path:%s] failed. err: %s", ls.backupDir, err))
	}

	backups := make([]*model.LsBackup, 0, len(entries))

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		info, err := entry.Info()
		if errors.Is(err, os.ErrNotExist) {
			return nil, xerr.NewCliErr("The file does not exist or has changed")
		} else if err != nil {
			return nil, xerr.NewCliErr(fmt.Sprintf("Unknown err: get entry info failed. err: %s", err))
		}

		if !strings.HasSuffix(info.Name(), ".json") {
			continue
		}

		path := fmt.Sprintf("%s/%s", ls.backupDir, info.Name())
		file, err := os.ReadFile(path)
		if err != nil {
			return nil, xerr.NewCliErr(fmt.Sprintf("read file failed. err: %s", err))
		}

		b := &model.LsBackup{}
		if err := json.Unmarshal(file, b); err != nil {
			return nil, xerr.NewCliErr(fmt.Sprintf("invalid contents[filePath=%s]. err: %s", path, err))
		}
		b.Info.FileName = info.Name()
		backups = append(backups, b)
	}
	return backups, nil
}

func (ls *localStorage) ReadByCSN(csn string) (*model.LsBackup, error) {
	list, err := ls.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		if v.Info.CSN == csn {
			return v, nil
		}
	}
	return nil, xerr.NewCliErr(xerr.NotFound)
}

func (ls *localStorage) ReadAllByCSN(csn string) ([]*model.LsBackup, error) {
	baks := []*model.LsBackup{}
	list, err := ls.ReadAll()
	if err != nil {
		return baks, err
	}
	for _, v := range list {
		c := v
		if v.Info.CSN == csn {
			baks = append(baks, c)
		}
	}

	return baks, nil
}

func (ls *localStorage) ReadByID(id string) (*model.LsBackup, error) {
	list, err := ls.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		if v.Info.ID == id {
			return v, nil
		}
	}
	return nil, xerr.NewCliErr(xerr.NotFound)
}

/*
GenFilename gen a filename based on the file extension

	if extn is empty,return a postfix-free filename
	if extn=JSON,return the JSON filename like **.json
*/
func (ls *localStorage) GenFilename(extn Extension) string {
	prefix := time.Now().UTC().Format("20060102150405")
	suffix := strutil.Random(8)

	switch extn {
	case ExtnJSON:
		return fmt.Sprintf("%s_%s.json", prefix, suffix)
	default:
		return fmt.Sprintf("%s_%s", prefix, suffix)
	}
}

type mode int

const (
	normal mode = iota
	hided
)

func (ls *localStorage) deleteByName(name string, mode mode) error {
	path := fmt.Sprintf("%s/%s", ls.backupDir, name)
	if err := os.Remove(path); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("delete file failed. err: %s", err))
	}
	return nil
}

func (ls *localStorage) DeleteByName(name string) error {
	return ls.deleteByName(name, normal)
}

func (ls *localStorage) DeleteByHidedName(name string) error {
	return ls.deleteByName(fmt.Sprintf(".%s", name), hided)
}

func (ls *localStorage) HideByName(name string) error {
	path := fmt.Sprintf("%s/%s", ls.backupDir, name)
	hided := fmt.Sprintf("%s/.%s", ls.backupDir, name)
	if err := os.Rename(path, hided); err != nil {
		return xerr.NewCliErr(fmt.Sprintf("hide file failed. err: %s", err))
	}
	return nil
}
