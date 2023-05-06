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

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
	dbmeshv1alpha1 "github.com/database-mesh/golang-sdk/kubernetes/api/v1alpha1"
)

func validCreateInstanceParams(node *v1alpha1.StorageNode, params map[string]string) error {
	requiredParams := map[string]string{
		"engine":             "engine is empty",
		"engineVersion":      "engine version is empty",
		"instanceClass":      "instance class is empty",
		"masterUsername":     "master username is empty",
		"masterUserPassword": "master user password is empty",
		"allocatedStorage":   "allocated storage is empty",
	}

	for k, v := range requiredParams {
		if val, ok := params[k]; !ok || val == "" {
			return errors.New(v)
		}
	}

	// validate instance identifier.
	if val, ok := node.Annotations[dbmeshv1alpha1.AnnotationsInstanceIdentifier]; !ok || val == "" {
		return errors.New("instance identifier is empty")
	}

	// TODO set options to generate password and write back to storage node annos.
	// TODO set options to set master username by user.
	// validate master user password length. must be greater than 8. from aws doc.
	if len(params["masterUserPassword"]) < 8 {
		return errors.New("master user password length should be greater than 8")
	}

	return nil
}

func (c *RdsClient) CreateInstance(ctx context.Context, node *v1alpha1.StorageNode, params map[string]string) error {
	// validate params
	if err := validCreateInstanceParams(node, params); err != nil {
		return err
	}

	storage, err := strconv.ParseInt(params["allocatedStorage"], 10, 64)
	if err != nil {
		return fmt.Errorf("allocated storage is not a number: %s", err.Error())
	}

	instance := c.Instance()
	instance.SetEngine(params["engine"]).
		SetEngineVersion(params["engineVersion"]).
		SetDBInstanceClass(params["instanceClass"]).
		SetDBInstanceIdentifier(node.Annotations[dbmeshv1alpha1.AnnotationsInstanceIdentifier]).
		SetMasterUsername(params["masterUsername"]).
		SetMasterUserPassword(params["masterUserPassword"]).
		SetAllocatedStorage(int32(storage))
	return instance.Create(ctx)
}

func (c *RdsClient) GetInstance(ctx context.Context, node *v1alpha1.StorageNode) (*rds.DescInstance, error) {
	identifier, ok := node.Annotations[dbmeshv1alpha1.AnnotationsInstanceIdentifier]
	if !ok {
		return nil, errors.New("instance identifier is empty")
	}
	instance := c.Instance()
	instance.SetDBInstanceIdentifier(identifier)
	return instance.Describe(ctx)
}

// DeleteInstance delete rds instance.
// aws rds instance status doc: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/accessing-monitoring.html
func (c *RdsClient) DeleteInstance(ctx context.Context, node *v1alpha1.StorageNode, databaseClass *dbmeshv1alpha1.DatabaseClass) error {
	// TODO add more test case.
	/* TODO set options to skip final snapshot and backup stuff depends on database class ClaimPolicy.
	"error": "operation error RDS: DeleteDBInstance,
	https response error StatusCode: 400,
	RequestID: ae094e3c-d8f1-49ba-aed1-cb0618b3641d,
	api error InvalidParameterCombination:
	FinalDBSnapshotIdentifier is required unless SkipFinalSnapshot is specified."
	*/

	identifier, ok := node.Annotations[dbmeshv1alpha1.AnnotationsInstanceIdentifier]
	if !ok {
		return errors.New("instance identifier is empty")
	}

	instance := c.Instance()
	instance.SetDBInstanceIdentifier(identifier)

	// check instance status first
	ins, err := c.GetInstance(ctx, node)
	if err != nil {
		return err
	}
	if ins == nil || ins.DBInstanceStatus == "deleting" {
		return nil
	}

	instance.SetDeleteAutomateBackups(true)
	instance.SetSkipFinalSnapshot(true)
	return instance.Delete(ctx)
}
