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
 *
 */

package storagenode

import (
	"context"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
)

type DatabaseCluster struct {
	Status          string   `json:"status"`
	PrimaryEndpoint string   `json:"primaryEndpoint"`
	ReaderEndpoints []string `json:"readerEndpoints"`
}

type DatabaseInstance struct {
	Endpoint v1alpha1.Endpoint `json:"endpoint"`
}

type IDBClusterClient interface {
	// IsValid validate the client parameters
	IsValid(node *v1alpha1.StorageNode) error
	GetCluster(ctx context.Context, node *v1alpha1.StorageNode) (cluster *DatabaseCluster, err error)
	CreateCluster(ctx context.Context, node *v1alpha1.StorageNode, params map[string]string) (cluster *DatabaseCluster, err error)
	DeleteCluster(ctx context.Context, node *v1alpha1.StorageNode) error
	GetInstances(ctx context.Context, cluster *DatabaseCluster) ([]*DatabaseInstance, error)
	CreateInstance(ctx context.Context) error
	DeleteInstance(ctx context.Context) error
}
