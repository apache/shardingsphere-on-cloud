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
	"strings"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
)

func validateCreateAuroraParams(node *v1alpha1.StorageNode, paramsPtr *map[string]string) error {
	requiredParams := map[string]string{
		"instanceClass":      "instance class is empty",
		"engine":             "engine is empty",
		"engineVersion":      "engine version is empty",
		"clusterIdentifier":  "cluster identifier is empty",
		"masterUsername":     "master username is empty",
		"masterUserPassword": "master user password is empty",
	}
	params := *paramsPtr
	if v, ok := node.Annotations[v1alpha1.AnnotationsClusterIdentifier]; !ok || v == "" {
		return errors.New("cluster identifier is empty")
	} else {
		params["clusterIdentifier"] = v
	}

	if len(params["clusterIdentifier"]) > 50 {
		return errors.New("cluster identifier is too long, max length is 50")
	}

	for k, v := range requiredParams {
		if val, ok := params[k]; !ok || val == "" {
			return fmt.Errorf(v)
		}
	}
	return nil
}

// CreateAuroraCluster creates aurora cluster
// ref: https://docs.aws.amazon.com/zh_cn/AmazonRDS/latest/APIReference/API_CreateDBInstance.html
func (c *RdsClient) CreateAuroraCluster(ctx context.Context, node *v1alpha1.StorageNode, params map[string]string) error {
	if err := validateCreateAuroraParams(node, &params); err != nil {
		return err
	}

	aurora := c.Aurora()

	// set required params
	aurora.SetDBInstanceClass(params["instanceClass"]).
		SetEngine(params["engine"]).
		SetEngineVersion(params["engineVersion"]).
		SetDBClusterIdentifier(params["clusterIdentifier"]).
		SetMasterUsername(params["masterUsername"]).
		SetMasterUserPassword(params["masterUserPassword"]).
		SetInstanceNumber(node.Spec.Replicas)

	if v, ok := node.Annotations[v1alpha1.AnnotationsInstanceDBName]; ok {
		aurora.SetDBName(v)
	}

	if v, ok := params["publicAccessible"]; ok && v == "false" {
		aurora.SetPublicAccessible(false)
	} else {
		aurora.SetPublicAccessible(true)
	}

	if params["vpcSecurityGroupIds"] != "" {
		aurora.SetVpcSecurityGroupIds(strings.Split(params["vpcSecurityGroupIds"], ","))
	}

	if err := aurora.Create(ctx); err != nil {
		return fmt.Errorf("create aurora cluster failed, %v", err)
	}

	return nil
}

func (c *RdsClient) GetAuroraCluster(ctx context.Context, node *v1alpha1.StorageNode) (cluster *rds.DescCluster, err error) {
	identifier, ok := node.Annotations[v1alpha1.AnnotationsClusterIdentifier]
	if !ok {
		return nil, errors.New("cluster identifier is empty")
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

	aurora := c.Aurora()
	aurora.SetDBClusterIdentifier(identifier)

	switch storageProvider.Spec.ReclaimPolicy {
	case v1alpha1.StorageReclaimPolicyDelete:
		aurora.SetDeleteAutomateBackups(true).SetSkipFinalSnapshot(true)
	case v1alpha1.StorageReclaimPolicyRetain:
		aurora.SetDeleteAutomateBackups(false).SetSkipFinalSnapshot(true)
	case v1alpha1.StorageReclaimPolicyDeleteWithFinalSnapshot:
		aurora.SetDeleteAutomateBackups(true).SetSkipFinalSnapshot(false)
	}

	return aurora.Delete(ctx)
}
