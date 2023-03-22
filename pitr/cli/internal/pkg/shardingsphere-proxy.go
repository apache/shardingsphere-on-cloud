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
	"encoding/json"
	"fmt"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/gsutil"
)

type (
	shardingSphereProxy struct {
		db *sql.DB
	}

	IShardingSphereProxy interface {
		ExportMetaData() (*model.ClusterInfo, error)
		ExportStorageNodes() ([]*model.StorageNode, error)
		LockForRestore() error
		LockForBackup() error
		Unlock() error
		ImportMetaData(in *model.ClusterInfo) error
	}
)

const (
	DefaultDbName = "postgres"
)

func NewShardingSphereProxy(user, password, dbName, host string, port uint16) (IShardingSphereProxy, error) {
	db, err := gsutil.Open(user, password, dbName, host, port)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		efmt := "db ping fail[host=%s,port=%d,user=%s,pwLen=%d,dbName=%s],err=%s"
		return nil, fmt.Errorf(efmt, host, port, user, len(password), dbName, err)
	}
	return &shardingSphereProxy{db: db}, nil
}

/*
LockForBackup 停写，同时锁 CSN，备份场景使用
*/
func (ss *shardingSphereProxy) LockForBackup() error {
	_, err := ss.db.Exec(`LOCK CLUSTER WITH LOCK_STRATEGY(TYPE(NAME="WRITE", PROPERTIES("lock_csn"=true)));`)
	if err != nil {
		return xerr.NewCliErr("ss lock for backup failure")
	}
	return nil
}

/*
LockForRestore 停读写，不需要锁 CSN，恢复场景使用
*/
func (ss *shardingSphereProxy) LockForRestore() error {
	_, err := ss.db.Exec(`LOCK CLUSTER WITH LOCK_STRATEGY(TYPE(NAME="READ_WRITE"))`)
	if err != nil {
		return xerr.NewCliErr("ss lock for restore failure")
	}
	return nil
}

func (ss *shardingSphereProxy) Unlock() error {
	_, err := ss.db.Exec("UNLOCK CLUSTER;")
	if err != nil {
		return xerr.NewCliErr("ss unlock failure")
	}
	return nil
}

/*
ExportMetaData 导出 SS 元数据

+-----------------------------+-------------------------+----------------------------------------+
| id                          | create_time             | data                                   |
+-------------------------------------------------------+----------------------------------------+
| 734bb036-b15d-4af0-be87-237 | 2023-01-01 12:00:00 897 | {"meta_data":{},"snapshot_info":{}}    |
+-------------------------------------------------------+----------------------------------------+
*/
func (ss *shardingSphereProxy) ExportMetaData() (*model.ClusterInfo, error) {
	query, err := ss.db.Query(`EXPORT METADATA`)
	if err != nil {
		return nil, xerr.NewCliErr(fmt.Sprintf("export meta data failure,err=%s", err))
	}
	var (
		id         string
		createTime string
		data       string
	)
	if query.Next() {
		if err = query.Scan(&id, &createTime, &data); err != nil {
			return nil, xerr.NewCliErr(fmt.Sprintf("query scan failure,err=%s", err))
		}
		if err = query.Close(); err != nil {
			return nil, xerr.NewCliErr(fmt.Sprintf("query close failure,err=%s", err))
		}
	}
	out := model.ClusterInfo{}
	if err = json.Unmarshal([]byte(data), &out); err != nil {
		return nil, fmt.Errorf("json unmarshal return err=%s", err)
	}
	return &out, nil
}

/*
ExportStorageNodes 导出存储节点数据

+-----------------------------+-------------------------+----------------------------------------+
| id                          | create_time             | data                                   |
+-------------------------------------------------------+----------------------------------------+
| 734bb036-b15d-4af0-be87-237 | 2023-01-01 12:00:00 897 | {"storage_nodes":{"sharding_db":[]}}   |
+-------------------------------------------------------+----------------------------------------+
*/
func (ss *shardingSphereProxy) ExportStorageNodes() ([]*model.StorageNode, error) {
	query, err := ss.db.Query(`EXPORT STORAGE NODES;`)
	if err != nil {
		return nil, xerr.NewCliErr(fmt.Sprintf("export storage nodes failure,err=%s", err))
	}
	var (
		id         string
		createTime string
		data       string
	)
	if query.Next() {
		if err = query.Scan(&id, &createTime, &data); err != nil {
			return nil, xerr.NewCliErr(fmt.Sprintf("query scan failure,err=%s", err))
		}

		if err = query.Close(); err != nil {
			return nil, xerr.NewCliErr(fmt.Sprintf("query close failure,err=%s", err))
		}
	}
	out := &model.StorageNodesInfo{}
	if err = json.Unmarshal([]byte(data), &out); err != nil {
		return nil, fmt.Errorf("json unmarshal return err=%s", err)
	}
	return out.StorageNodes.List, nil
}

// ImportMetaData 备份数据恢复
func (ss *shardingSphereProxy) ImportMetaData(in *model.ClusterInfo) error {
	if in == nil {
		return xerr.NewCliErr("import meta data is nil")
	}
	marshal, err := json.Marshal(in)
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("json marshal,invalid data[in=%+v]", in))
	}

	_, err = ss.db.Exec(fmt.Sprintf(`IMPORT METADATA "%s";`, marshal))
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("import metadata failure,err=%s", err))
	}

	return nil
}
