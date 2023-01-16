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

package proxyconfig

import (
	"testing"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func Test_ToYAML(t *testing.T) {
	cases := []struct {
		proxyCfg *v1alpha1.ShardingSphereProxyServerConfig
		exp      string
		message  string
	}{
		{
			proxyCfg: &v1alpha1.ShardingSphereProxyServerConfig{},
			exp: `mode:
  type: ""
  repository:
    type: ""
    props:
      namespace: ""
      server-lists: ""
authority:
  users: []
  privilege: null
`,
			message: "Empty proxy server config should generate yaml with empty value",
		},
		{
			proxyCfg: &v1alpha1.ShardingSphereProxyServerConfig{
				Spec: v1alpha1.ProxyConfigSpec{
					ClusterConfig: v1alpha1.ClusterConfig{
						Type: "Cluster",
						Repository: v1alpha1.RepositoryConfig{
							Type: "Zookeeper",
							Props: v1alpha1.ClusterProps{
								Namespace:                    "cluster-sharding-mode",
								ServerLists:                  "host1:2181,host2:2181",
								RetryIntervalMilliseconds:    500,
								MaxRetries:                   3,
								TimeToLiveSeconds:            60,
								OperationTimeoutMilliseconds: 500,
							},
						},
					},
					Authority: v1alpha1.Auth{
						Users: []v1alpha1.User{
							{
								User:     "John",
								Password: "123456",
							},
						},
						Privilege: &v1alpha1.Privilege{
							Type: "Admin",
						},
					},
					Props: &v1alpha1.Props{
						KernelExecutorSize:                16,
						CheckTableMetadataEnabled:         false,
						ProxyBackendQueryFetchSize:        -1,
						CheckDuplicateTableEnabled:        false,
						ProxyFrontendExecutorSize:         0,
						ProxyBackendExecutorSuitable:      "OLAP",
						ProxyBackendDriverType:            "JDBC",
						ProxyFrontendDatabaseProtocolType: "",
					},
				},
			},
			exp: `mode:
  type: Cluster
  repository:
    type: Zookeeper
    props:
      namespace: cluster-sharding-mode
      server-lists: host1:2181,host2:2181
      retryIntervalMilliseconds: 500
      maxRetries: 3
      timeToLiveSeconds: 60
      operationTimeoutMilliseconds: 500
authority:
  users:
  - user: John
    password: "123456"
  privilege:
    type: Admin
props:
  kernel-executor-size: 16
  proxy-backend-query-fetch-size: -1
  proxy-backend-executor-suitable: OLAP
  proxy-backend-driver-type: JDBC
`,
			message: "Should generate yaml with the same value as ShardingSphereProxyServerConfig.Spec",
		},
	}

	for _, c := range cases {
		act := toYaml(c.proxyCfg)
		assert.Equal(t, c.exp, act, c.message)
	}
}
