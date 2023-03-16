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
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/model"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("IShardingSphereProxy", func() {
	Context("NewShardingSphereProxy", func() {
		var (
			host     = "local"
			port     = uint16(13308)
			username = "root"
			password = "root"
			dbName   = "postgres"
		)

		It("Connecting shardingsphere proxy", func() {
			Skip("Manually exec:dependent environment")
			ss, err := NewShardingSphereProxy(username, password, dbName, host, port)
			Expect(err).To(BeNil())
			Expect(ss).NotTo(BeNil())
		})

		It("Export meta data", func() {
			Skip("Manually exec:dependent environment")
			ss, err := NewShardingSphereProxy(username, password, dbName, host, port)
			Expect(err).To(BeNil())
			Expect(ss).NotTo(BeNil())

			fmt.Println(ss.ExportMetaData())
		})

		It("Export storage node", func() {
			Skip("Manually exec:dependent environment")
			ss, err := NewShardingSphereProxy(username, password, dbName, host, port)
			Expect(err).To(BeNil())
			Expect(ss).NotTo(BeNil())

			fmt.Println(ss.ExportStorageNodes())

			ss, err = NewShardingSphereProxy(username, password, dbName, host, port)
			Expect(err).To(BeNil())
			Expect(ss).NotTo(BeNil())

			fmt.Println(ss.ExportStorageNodes())
		})

		It("Lock and unlock", func() {
			Skip("Manually exec:dependent environment")
			ss, err := NewShardingSphereProxy(username, password, dbName, host, port)
			Expect(err).To(BeNil())
			Expect(ss).NotTo(BeNil())

			fmt.Println(ss.LockForRestore())
			time.Sleep(time.Second * 5)
			fmt.Println(ss.Unlock())

			fmt.Println(ss.LockForBackup())
			time.Sleep(time.Second * 5)
			fmt.Println(ss.Unlock())
		})

	})
})

var (
	// implement with your own env
	u  string
	p  string
	h  string
	pt uint16
	db string
)

func Test_shardingSphereProxy_Unlock(t *testing.T) {
	tests := []struct {
		name string

		wantErr bool
	}{
		{
			name:    "test unlock after lock",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss, _ := NewShardingSphereProxy(u, p, db, h, pt)
			if err := ss.LockForBackup(); (err != nil) != tt.wantErr {
				t.Errorf("Lock() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := ss.Unlock(); (err != nil) != tt.wantErr {
				t.Errorf("Unlock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_shardingSphereProxy_ExportMetaData(t *testing.T) {
	tests := []struct {
		name    string
		want    *model.ClusterInfo
		wantErr bool
	}{
		{
			name:    "test export metadata",
			wantErr: false,
			want:    &model.ClusterInfo{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss, _ := NewShardingSphereProxy(u, p, db, h, pt)
			_, err := ss.ExportMetaData()
			if (err != nil) != tt.wantErr {
				t.Errorf("ExportMetaData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("ExportMetaData() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func Test_shardingSphereProxy_ExportStorageNodes(t *testing.T) {
	tests := []struct {
		name    string
		want    []*model.StorageNode
		wantErr bool
	}{
		{
			name:    "test export storage nodes",
			want:    []*model.StorageNode{},
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss, _ := NewShardingSphereProxy(u, p, db, h, pt)
			_, err := ss.ExportStorageNodes()
			if (err != nil) != tt.wantErr {
				t.Errorf("ExportStorageNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("ExportStorageNodes() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
