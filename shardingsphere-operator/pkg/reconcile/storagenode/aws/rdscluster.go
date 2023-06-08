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
	"strconv"
	"strings"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
)

func validateCreateRDSClusterParams(node *v1alpha1.StorageNode, paramsPtr *map[string]string) error {
	requiredParams := map[string]string{
		"instanceClass":      "instance class is empty",
		"engine":             "engine is empty",
		"engineVersion":      "engine version is empty",
		"clusterIdentifier":  "cluster identifier is empty",
		"masterUsername":     "master username is empty",
		"masterUserPassword": "master user password is empty",
		"allocatedStorage":   "allocated storage is empty",
		"iops":               "iops is empty",
		"storageType":        "storage type is empty",
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

	// valid mysql engine version
	if params["engine"] == "mysql" {
		version := strings.Split(params["engineVersion"], ".")[0]
		if version != "8" {
			return fmt.Errorf("mysql engine version is not supported, only support 8.x")
		}
	}

	if params["storageType"] != "io1" {
		return fmt.Errorf("storage type is not supported, only support io1")
	}

	return nil
}

func getAllocatedStorage(allocatedStorageStr string) (allocatedStorage int, err error) {
	allocatedStorage, err = strconv.Atoi(allocatedStorageStr)
	if err != nil {
		return 0, fmt.Errorf("allocated storage is not a number, %v", err)
	}
	if allocatedStorage < 100 || allocatedStorage > 65536 {
		return 0, fmt.Errorf("allocated storage is out of range, min is 100, max is 65536")
	}
	return allocatedStorage, nil
}

func getIOPS(iopsStr string) (iops int, err error) {
	iops, err = strconv.Atoi(iopsStr)
	if err != nil {
		return 0, fmt.Errorf("iops is not a number, %v", err)
	}
	if iops < 1000 || iops > 256000 {
		return 0, fmt.Errorf("iops is out of range, min is 1000, max is 256000")
	}
	return iops, nil
}

// CreateRDSCluster creates rds cluster
// ref: https://docs.aws.amazon.com/zh_cn/AmazonRDS/latest/APIReference/API_CreateDBInstance.html
func (c *RdsClient) CreateRDSCluster(ctx context.Context, node *v1alpha1.StorageNode, params map[string]string) error {
	if err := validateCreateRDSClusterParams(node, &params); err != nil {
		return err
	}

	cc := c.Cluster()

	allocatedStorage, err := getAllocatedStorage(params["allocatedStorage"])
	if err != nil {
		return err
	}

	iops, err := getIOPS(params["iops"])
	if err != nil {
		return err
	}

	if iops > allocatedStorage*50 || iops*2 < allocatedStorage {
		return fmt.Errorf("the IOPS to GiB ratio must be between 0.5 and 50, current iops is %d, allocated storage is %d", iops, allocatedStorage)
	}

	cc.SetEngine(params["engine"]).
		SetEngineVersion(params["engineVersion"]).
		SetDBClusterIdentifier(params["clusterIdentifier"]).
		SetMasterUsername(params["masterUsername"]).
		SetMasterUserPassword(params["masterUserPassword"]).
		SetAllocatedStorage(int32(allocatedStorage)).
		SetDBClusterInstanceClass(params["instanceClass"]).
		SetIOPS(int32(iops)).
		SetStorageType(params["storageType"])

	if v, ok := node.Annotations[v1alpha1.AnnotationsInstanceDBName]; ok && v != "" {
		cc.SetDatabaseName(v)
	}

	if v, ok := params["publicAccessible"]; ok && v == "false" {
		cc.SetPublicAccessible(false)
	} else {
		cc.SetPublicAccessible(true)
	}

	if params["vpcSecurityGroupIds"] != "" {
		cc.SetVpcSecurityGroupIds(strings.Split(params["vpcSecurityGroupIds"], ","))
	}

	if err := cc.Create(ctx); err != nil {
		return fmt.Errorf("create rds cluster failed, %v", err)
	}

	return nil
}

func (c *RdsClient) GetRDSCluster(ctx context.Context, node *v1alpha1.StorageNode) (cluster *rds.DescCluster, err error) {
	identifier, ok := node.Annotations[v1alpha1.AnnotationsClusterIdentifier]
	if !ok {
		return nil, errors.New("cluster identifier is empty")
	}

	cc := c.Cluster()
	cc.SetDBClusterIdentifier(identifier)
	return cc.Describe(ctx)
}

func (c *RdsClient) DeleteRDSCluster(ctx context.Context, node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) error {
	identifier, ok := node.Annotations[v1alpha1.AnnotationsClusterIdentifier]
	if !ok {
		return fmt.Errorf("cluster identifier is empty")
	}

	cc := c.Cluster()
	cc.SetDBClusterIdentifier(identifier)

	cluster, err := cc.Describe(ctx)
	if err != nil {
		return fmt.Errorf("describe rds cluster failed, %v", err)
	}
	if cluster == nil || cluster.Status == string(rds.DBClusterStatusDeleting) {
		return nil
	}

	switch storageProvider.Spec.ReclaimPolicy {
	case v1alpha1.StorageReclaimPolicyDelete:
		cc.SetSkipFinalSnapshot(true)
	case v1alpha1.StorageReclaimPolicyDeleteWithFinalSnapshot:
		if v, ok := node.Annotations[v1alpha1.AnnotationsFinalSnapshotIdentifier]; !ok || v == "" {
			return fmt.Errorf("final snapshot identifier is empty")
		}
		if cluster.Status != string(rds.DBClusterStatusAvailable) {
			return fmt.Errorf("rds cluster is not available, can not delete with final snapshot")
		}
		cc.SetFinalDBSnapshotIdentifier(node.Annotations[v1alpha1.AnnotationsFinalSnapshotIdentifier])
		cc.SetSkipFinalSnapshot(false)
	case v1alpha1.StorageReclaimPolicyRetain:
		return fmt.Errorf("rds cluster does not support retain policy")
	}

	return cc.Delete(ctx)
}
