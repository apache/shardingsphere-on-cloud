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

package aws

import (
	"context"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
	dbmeshv1alpha1 "github.com/database-mesh/golang-sdk/kubernetes/api/v1alpha1"
)

func (c *RdsClient) CreateAuroraCluster(ctx context.Context, node *v1alpha1.StorageNode, params map[string]string) error {
	aurora := c.Aurora()
	err := aurora.Create(ctx)
	return err
}

func (c *RdsClient) GetAuroraCluster(ctx context.Context, node *v1alpha1.StorageNode) (cluster *rds.DescCluster, err error) {
	if node.Status.Cluster.Properties == nil || node.Status.Cluster.Properties["clusterIdentifier"] == "" {
		// cluster not created
		return nil, nil
	}

	aurora := c.Aurora()
	return aurora.Describe(ctx)
}

func (c *RdsClient) DeleteAuroraCluster(ctx context.Context, node *v1alpha1.StorageNode, databaseClass *dbmeshv1alpha1.DatabaseClass) error {
	aurora := c.Aurora()
	return aurora.Delete(ctx)
}
