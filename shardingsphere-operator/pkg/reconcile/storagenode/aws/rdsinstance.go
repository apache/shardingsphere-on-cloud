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
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"

	"github.com/database-mesh/golang-sdk/aws/client/rds"
	"github.com/database-mesh/golang-sdk/pkg/random"
)

// nolint:gocognit
func validCreateInstanceParams(node *v1alpha1.StorageNode, paramsptr *map[string]string) error {
	requiredParams := map[string]string{
		"engine":             "engine is empty",
		"engineVersion":      "engine version is empty",
		"instanceClass":      "instance class is empty",
		"masterUsername":     "master username is empty",
		"masterUserPassword": "master user password is empty",
		"allocatedStorage":   "allocated storage is empty",
	}
	params := *paramsptr

	for k, v := range requiredParams {
		if val, ok := params[k]; !ok || val == "" {
			if k == "masterUserPassword" || k == "masterUsername" {
				generator := generate(k)
				params[k] = generator
			} else {
				return errors.New(v)
			}
		}
	}

	if username, ok := params["masterUsername"]; ok {
		validatedUsername, err := validateusername(username)
		if err != nil {
			return err
		}
		params["masterUsername"] = validatedUsername
	}

	// validate instance identifier.
	if val, ok := node.Annotations[v1alpha1.AnnotationsInstanceIdentifier]; !ok || val == "" {
		return errors.New("instance identifier is empty")
	}

	// TODO set options to generate password and write back to storage node annos.
	// TODO set options to set master username by user.
	// validate master user password length. must be greater than 8. from aws doc.
	lp := len(params["masterUserPassword"])
	if lp < 8 || lp > 41 {
		return errors.New("master user password length should be greater than 8")
	} else {
		node.Annotations[v1alpha1.AnnotationsMasterUserPassword] = params["masterUserPassword"]
	}

	return nil
}

func validateusername(str string) (string, error) {
	if l := len(str); l > 16 {
		return "", errors.New("username length should be less than 16")
	}

	pattern := "^[a-zA-Z0-9_-]+$"
	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		return "", fmt.Errorf("validateusername error: %s", err.Error())
	}
	if !matched {
		return "", errors.New("username contains invalid characters")
	}
	str = regexp.MustCompile(`_{2,}`).ReplaceAllString(str, "_")
	str = regexp.MustCompile(`-{2,}`).ReplaceAllString(str, "-")
	return str, nil
}

func generate(gentype string) string {
	switch gentype {
	case "masterUsername":
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(16)))
		if err != nil {
			return ""
		}
		length := int(nBig.Int64()) + 1
		return random.StringN(length)
	case "masterUserPassword":
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(34)))
		if err != nil {
			return ""
		}
		length := int(nBig.Int64()) + 8
		chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%^&*()_+-={}[]:;?"
		return random.StringCustom(length, []byte(chars))
	default:
		return ""
	}
}

func (c *RdsClient) CreateInstance(ctx context.Context, node *v1alpha1.StorageNode, params map[string]string) error {
	// validate params
	if err := validCreateInstanceParams(node, &params); err != nil {
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
		SetDBInstanceIdentifier(node.Annotations[v1alpha1.AnnotationsInstanceIdentifier]).
		SetMasterUsername(params["masterUsername"]).
		SetMasterUserPassword(params["masterUserPassword"]).
		SetAllocatedStorage(int32(storage))
	// set database name if needed.
	if v, ok := params[node.Annotations[v1alpha1.AnnotationsInstanceDBName]]; ok {
		instance.SetDBName(v)
	}
	return instance.Create(ctx)
}

func (c *RdsClient) GetInstance(ctx context.Context, node *v1alpha1.StorageNode) (*rds.DescInstance, error) {
	identifier, ok := node.Annotations[v1alpha1.AnnotationsInstanceIdentifier]
	if !ok {
		return nil, errors.New("instance identifier is empty")
	}
	instance := c.Instance()
	instance.SetDBInstanceIdentifier(identifier)
	return instance.Describe(ctx)
}

func (c *RdsClient) GetInstanceByIdentifier(ctx context.Context, identifier string) (*rds.DescInstance, error) {
	instance := c.Instance()
	instance.SetDBInstanceIdentifier(identifier)
	return instance.Describe(ctx)
}

// DeleteInstance delete rds instance.
// aws rds instance status doc: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/accessing-monitoring.html
func (c *RdsClient) DeleteInstance(ctx context.Context, node *v1alpha1.StorageNode, storageProvider *v1alpha1.StorageProvider) error {
	// TODO add more test case.
	/* TODO set options to skip final snapshot and backup stuff depends on database class ClaimPolicy.
	"error": "operation error RDS: DeleteDBInstance,
	https response error StatusCode: 400,
	RequestID: ae094e3c-d8f1-49ba-aed1-cb0618b3641d,
	api error InvalidParameterCombination:
	FinalDBSnapshotIdentifier is required unless SkipFinalSnapshot is specified."
	*/

	identifier, ok := node.Annotations[v1alpha1.AnnotationsInstanceIdentifier]
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

	if ins == nil || ins.DBInstanceStatus == rds.DBInstanceStatusDeleting {
		return nil
	}

	var isDeleteBackup, isSkipFinalSnapshot bool
	switch storageProvider.Spec.ReclaimPolicy {
	case v1alpha1.StorageReclaimPolicyDeleteWithFinalSnapshot:
		isDeleteBackup, isSkipFinalSnapshot = true, false
	case v1alpha1.StorageReclaimPolicyDelete:
		isDeleteBackup, isSkipFinalSnapshot = true, true
	case v1alpha1.StorageReclaimPolicyRetain:
		isDeleteBackup, isSkipFinalSnapshot = false, true
	}

	instance.SetDeleteAutomateBackups(isDeleteBackup)
	instance.SetSkipFinalSnapshot(isSkipFinalSnapshot)

	return instance.Delete(ctx)
}
