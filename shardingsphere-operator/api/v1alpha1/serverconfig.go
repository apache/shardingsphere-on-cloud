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

package v1alpha1

// User is a slice about authorized host and password for compute node.
// Format:
// user:<username>@<hostname>,hostname is % or empty string means do not care about authorized host
// password:<password>
type User struct {
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}

// Privilege for storage node, the default value is ALL_PRIVILEGES_PERMITTED
type Privilege struct {
	Type string `json:"type" yaml:"type"`
}

// Auth  is used to set up initial user to login compute node, and authority data of storage node.
type Auth struct {
	Users []User `json:"users" yaml:"users"`
	// +optional
	Privilege *Privilege `json:"privilege,omitempty"`
}

// Props Apache ShardingSphere provides the way of property configuration to configure system level configuration.
type Props struct {
	// The max thread size of worker group to execute SQL. One ShardingSphereDataSource will use a independent thread pool, it does not share thread pool even different data source in same JVM.
	// +optional
	KernelExecutorSize int `json:"kernel-executor-size,omitempty" yaml:"kernel-executor-size,omitempty"`
	// Whether validate table meta data consistency when application startup or updated.
	// +optional
	CheckTableMetadataEnabled bool `json:"check-table-metadata-enabled,omitempty" yaml:"check-table-metadata-enabled,omitempty"`
	// ShardingSphereProxy backend query fetch size. A larger value may increase the memory usage of ShardingSphere ShardingSphereProxy. The default value is -1, which means set the minimum value for different JDBC drivers.
	// +optional
	ProxyBackendQueryFetchSize int `json:"proxy-backend-query-fetch-size,omitempty" yaml:"proxy-backend-query-fetch-size,omitempty"`
	// Whether validate duplicate table when application startup or updated.
	// +optional
	CheckDuplicateTableEnabled bool `json:"check-duplicate-table-enabled,omitempty" yaml:"check-duplicate-table-enabled,omitempty"`
	// ShardingSphereProxy frontend Netty executor size. The default value is 0, which means let Netty decide.
	// +optional
	ProxyFrontendExecutorSize int `json:"proxy-frontend-executor-size,omitempty" yaml:"proxy-frontend-executor-size,omitempty"`
	// Available options of proxy backend executor suitable: OLAP(default), OLTP. The OLTP option may reduce time cost of writing packets to client, but it may increase the latency of SQL execution and block other clients if client connections are more than proxy-frontend-executor-size, especially executing slow SQL.
	// +optional
	ProxyBackendExecutorSuitable string `json:"proxy-backend-executor-suitable,omitempty" yaml:"proxy-backend-executor-suitable,omitempty"`
	// +optional
	ProxyBackendDriverType string `json:"proxy-backend-driver-type,omitempty" yaml:"proxy-backend-driver-type,omitempty"`
	// +optional
	ProxyFrontendDatabaseProtocolType string `json:"proxy-frontend-database-protocol-type" yaml:"proxy-frontend-database-protocol-type,omitempty"`
}

type ClusterProps struct {
	// Namespace of registry center
	Namespace string `json:"namespace" yaml:"namespace"`
	//Server lists of registry center
	ServerLists string `json:"server-lists" yaml:"server-lists"`
	//RetryIntervalMilliseconds Milliseconds of retry interval. default: 500
	// +optional
	RetryIntervalMilliseconds int `json:"retryIntervalMilliseconds,omitempty" yaml:"retryIntervalMilliseconds,omitempty"`
	// MaxRetries Max retries of client connection. default: 3
	// +optional
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`
	// TimeToLiveSeconds Seconds of ephemeral data live.default: 60
	// +optional
	TimeToLiveSeconds int `json:"timeToLiveSeconds,omitempty" yaml:"timeToLiveSeconds,omitempty"`
	// OperationTimeoutMilliseconds Milliseconds of operation timeout. default: 500
	// +optional
	OperationTimeoutMilliseconds int `json:"operationTimeoutMilliseconds,omitempty" yaml:"operationTimeoutMilliseconds,omitempty"`
	// Password of login
	// +optional
	Digest string `json:"digest,omitempty" yaml:"digest,omitempty"`
}

type RepositoryConfig struct {

	// +kubebuilder:validation:Enum=ZooKeeper

	//Type of persist repository
	Type string `json:"type" yaml:"type"`
	//Properties of persist repository
	Props ClusterProps `json:"props" yaml:"props"`
}

// ClusterConfig needs to fill in the relevant configuration required by Cluster mode
type ClusterConfig struct {

	// +kubebuilder:validation:Enum=Cluster

	// Type of mode configuration. Values only support: Cluster
	Type string `json:"type" yaml:"type"`
	// Persist repository configuration
	Repository RepositoryConfig `json:"repository" yaml:"repository"`
	// Whether overwrite persistent configuration with local configuration
	// If it is false, it will not upload the local configuration to ZooKeeper.
	// If it is true, it will upload the local configuration to ZooKeeper and overwrite the original configuration of ZooKeeper.
	Overwrite bool `json:"overwrite" yaml:"overwrite"`
}
