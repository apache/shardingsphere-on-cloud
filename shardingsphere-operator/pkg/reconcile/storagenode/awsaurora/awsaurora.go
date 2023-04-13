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

package awsaurora

import (
	"context"
	"errors"
	"strings"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
)

const (
	AnnoUsername            = "shardingsphere.apache.org/storagenode-aws-aurora-username"
	AnnoPassword            = "shardingsphere.apache.org/storagenode-aws-aurora-password"
	AnnoVpcSecurityGroupIds = "shardingsphere.apache.org/storagenode-aws-aurora-vpcSecurityGroupIds"
	AnnoSubnetGroupName     = "shardingsphere.apache.org/storagenode-aws-aurora-subnetGroupName"
	AnnoDBClusterIdentifier = "shardingsphere.apache.org/storagenode-aws-aurora-dbClusterIdentifier"
)

type AwsAurora struct {
	aurora rds.Aurora
}

func New(rds rds.RDS) storagenode.IDBClusterClient {
	return &AwsAurora{aurora: rds.Aurora()}
}

// IsValid validate the parameters in annotation for aws aurora
// required parameters:
// - username
// - password
// - vpcSecurityGroupIds
// - subnetGroupName
// - dbClusterIdentifier
func (a *AwsAurora) IsValid(node *v1alpha1.StorageNode) error {
	var errMsg []string
	if node.Annotations[AnnoUsername] == "" {
		errMsg = append(errMsg, "username is required")
	}
	if node.Annotations[AnnoPassword] == "" {
		errMsg = append(errMsg, "password is required")
	}
	if node.Annotations[AnnoVpcSecurityGroupIds] == "" {
		errMsg = append(errMsg, "vpcSecurityGroupIds is required")
	}
	if node.Annotations[AnnoSubnetGroupName] == "" {
		errMsg = append(errMsg, "subnetGroupName is required")
	}
	if node.Annotations[AnnoDBClusterIdentifier] == "" {
		errMsg = append(errMsg, "dbClusterIdentifier is required")
	}

	if len(errMsg) > 0 {
		return errors.New(strings.Join(errMsg, "\n"))
	}
	return nil
}

func (a *AwsAurora) GetCluster(ctx context.Context, node *v1alpha1.StorageNode) (cluster *storagenode.DatabaseCluster, err error) {
	if node.Status.Cluster.Properties == nil || node.Status.Cluster.Properties["clusterIdentifier"] == "" {
		// cluster not created
		return nil, nil
	}

	//cluster, err := a.aurora.DescribeDBClusters(ctx, node.Status.Cluster.Properties["ClusterIdentifier"])
	//if err != nil {
	//	return nil, err
	//}

	return cluster, nil
}

func (a *AwsAurora) CreateCluster(ctx context.Context, node *v1alpha1.StorageNode, params map[string]string) (cluster *storagenode.DatabaseCluster, err error) {
	err = a.aurora.Create(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (a *AwsAurora) DeleteCluster(ctx context.Context, node *v1alpha1.StorageNode) error {
	//TODO implement me
	panic("implement me")
}

func (a *AwsAurora) GetInstances(ctx context.Context, cluster *storagenode.DatabaseCluster) ([]*storagenode.DatabaseInstance, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AwsAurora) CreateInstance(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a *AwsAurora) DeleteInstance(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
