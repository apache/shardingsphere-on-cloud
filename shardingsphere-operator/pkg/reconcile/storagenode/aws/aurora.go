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
	"errors"
	"fmt"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
)

// CreateAuroraCluster creates aurora cluster
// ref: https://docs.aws.amazon.com/zh_cn/AmazonRDS/latest/APIReference/API_CreateDBInstance.html
func (c *RdsClient) CreateAuroraCluster(ctx context.Context, node *v1alpha1.StorageNode, params map[string]string) error {
	aurora := c.Aurora()

	// set required params
	aurora.SetDBInstanceClass(params["instanceClass"]).
		SetEngine(params["engine"]).
		SetDBClusterIdentifier(params["clusterIdentifier"])

	// set optional params
	if params["engineVersion"] != "" {
		aurora.SetEngineVersion(params["engineVersion"])
	}
	if params["masterUsername"] != "" {
		aurora.SetMasterUsername(params["masterUsername"])
	}
	if params["masterUserPassword"] != "" {
		aurora.SetMasterUserPassword(params["masterUserPassword"])
	}

	err := aurora.Create(ctx)
	return err
}

func (c *RdsClient) GetAuroraCluster(ctx context.Context, node *v1alpha1.StorageNode) (cluster *rds.DescCluster, err error) {
	identifier, ok := node.Annotations[v1alpha1.AnnotationsClusterIdentifier]
	if !ok {
		return nil, errors.New("cluster identifier is empty")
	}
	if node.Status.Cluster.Properties == nil || node.Status.Cluster.Properties["clusterIdentifier"] == "" {
		// cluster not created
		return nil, nil
	}

	aurora := c.Aurora()
	aurora.SetDBClusterIdentifier(identifier)
	return aurora.Describe(ctx)
}

func (c *RdsClient) DeleteAuroraCluster(ctx context.Context, node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) error {
	identifier, ok := node.Annotations[v1alpha1.AnnotationsClusterIdentifier]
	if !ok {
		return fmt.Errorf("cluster identifier is empty")
	}
	// get instances of aurora cluster
	filters := map[string][]string{
		"db-cluster-id": {identifier},
	}
	instances, err := c.GetInstancesByFilters(ctx, filters)
	if err != nil {
		return fmt.Errorf("get instances failed, %v", err)
	}
	// delete instance first
	for _, ins := range instances {
		if err := c.DeleteInstance(ctx, node, storageProvider); err != nil {
			return fmt.Errorf("delete instance=%s of aurora=%s failed, %v", ins.DBInstanceIdentifier, identifier, err)
		}
	}
	// delete cluster
	aurora := c.Aurora()
	aurora.SetDBClusterIdentifier(identifier)
	return aurora.Delete(ctx)
}
