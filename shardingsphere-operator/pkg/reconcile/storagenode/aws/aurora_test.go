/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
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

	"github.com/database-mesh/golang-sdk/aws"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ctx = context.Background()

var _ = Describe("Aurora", func() {
	Context("Test valid create aurora params", func() {
		It("should be success", func() {
			sn := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-valid-create-aurora-params",
					Namespace: "test-namespace",
					Annotations: map[string]string{
						v1alpha1.AnnotationsClusterIdentifier: "test",
					},
				},
			}
			params := map[string]string{
				"engine":             "aurora-mysql",
				"engineVersion":      "5.7",
				"instanceClass":      "db.t2.small",
				"clusterIdentifier":  "",
				"masterUsername":     "root",
				"masterUserPassword": "root123456",
			}
			err := validateCreateAuroraParams(sn, &params)
			Expect(err).To(BeNil())
		})
	})
})

var _ = Describe("Test For AWS Aurora Manually", func() {
	var (
		region    string
		accessKey string
		secretKey string
	)
	Context("Test create aurora cluster with 2 replicas", func() {
		It("should be success", func() {
			if region == "" || accessKey == "" || secretKey == "" {
				Skip("Skip test create aurora cluster")
			}

			sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
			rdsClient := rds.NewService(sess[region])
			awsClient := NewRdsClient(rdsClient)

			sn := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-create-aurora-cluster",
					Namespace: "test-namespace",
					Annotations: map[string]string{
						v1alpha1.AnnotationsClusterIdentifier: "test-create-aurora-cluster-identifier",
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					Replicas: 2,
				},
			}
			params := map[string]string{
				"engine":             "aurora-mysql",
				"engineVersion":      "5.7",
				"instanceClass":      "db.t2.small",
				"clusterIdentifier":  "",
				"masterUsername":     "root",
				"masterUserPassword": "root123456",
			}

			Expect(awsClient.CreateAuroraCluster(ctx, sn, params)).Should(Succeed())
		})
	})
	Context("Test Get Aurora Cluster", func() {
		It("should be success", func() {
			if region == "" || accessKey == "" || secretKey == "" {
				Skip("Skip test create aurora cluster")
			}
			sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
			rdsClient := rds.NewService(sess[region])
			awsClient := NewRdsClient(rdsClient)
			sn := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-get-aurora-cluster",
					Namespace: "test-namespace",
					Annotations: map[string]string{
						v1alpha1.AnnotationsClusterIdentifier: "test-create-aurora-cluster-identifier",
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					Replicas: 2,
				},
			}

			ac, err := awsClient.GetAuroraCluster(ctx, sn)
			Expect(err).To(BeNil())
			Expect(len(ac.DBClusterMembers)).To(Equal(2))
		})
	})

	Context("Test Delete Aurora Cluster", func() {
		It("should be success", func() {
			if region == "" || accessKey == "" || secretKey == "" {
				Skip("Skip test create aurora cluster")
			}
			sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
			rdsClient := rds.NewService(sess[region])
			awsClient := NewRdsClient(rdsClient)

			sn := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-get-aurora-cluster",
					Namespace: "test-namespace",
					Annotations: map[string]string{
						v1alpha1.AnnotationsClusterIdentifier: "test-create-aurora-cluster-identifier",
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: "test-get-aurora-cluster",
				},
			}

			storageProvider := &v1alpha1.StorageProvider{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-get-aurora-cluster",
				},
				Spec: v1alpha1.StorageProviderSpec{
					Provisioner:   v1alpha1.ProvisionerAWSAurora,
					ReclaimPolicy: v1alpha1.StorageReclaimPolicyDelete,
				},
			}
			Expect(awsClient.DeleteAuroraCluster(ctx, sn, storageProvider)).Should(Succeed())
		})
	})
})
