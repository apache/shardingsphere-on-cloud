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
 *
 */

package e2e

import (
	"context"
	"reflect"
	"time"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode/aws"
	dbmesh_rds "github.com/database-mesh/golang-sdk/aws/client/rds"

	"bou.ke/monkey"
	dbmeshv1alpha1 "github.com/database-mesh/golang-sdk/kubernetes/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("StorageNode Controller Suite Test", func() {
	var databaseClassName = "test-database-class"

	BeforeEach(func() {
		databaseClass := &dbmeshv1alpha1.DatabaseClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: databaseClassName,
			},
			Spec: dbmeshv1alpha1.DatabaseClassSpec{
				Provisioner: dbmeshv1alpha1.ProvisionerAWSRDSInstance,
				Parameters: map[string]string{
					"engine":             "mysql",
					"engineVersion":      "5.7",
					"instanceClass":      "db.t3.micro",
					"allocatedStorage":   "20",
					"masterUsername":     "root",
					"masterUserPassword": "root123456",
				},
			},
		}

		Expect(k8sClient.Create(ctx, databaseClass)).Should(Succeed())
	})

	AfterEach(func() {
		databaseClass := &dbmeshv1alpha1.DatabaseClass{}
		Expect(k8sClient.Get(ctx, client.ObjectKey{Name: databaseClassName}, databaseClass)).Should(Succeed())
		Expect(k8sClient.Delete(ctx, databaseClass)).Should(Succeed())
	})

	Context("reconcile storageNode", func() {
		AfterEach(func() {
			monkey.UnpatchAll()
		})

		It("should create success", func() {
			// mock get instance func returns success
			monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "GetInstance", func(_ *aws.RdsClient, _ context.Context, _ *v1alpha1.StorageNode) (*dbmesh_rds.DescInstance, error) {
				return &dbmesh_rds.DescInstance{
					DBInstanceStatus: v1alpha1.StorageNodeInstanceStatusAvailable,
					Endpoint: dbmesh_rds.Endpoint{
						Address: "127.0.0.1",
						Port:    3306,
					},
				}, nil
			})
			// mock delete instance func returns success
			monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "DeleteInstance", func(_ *aws.RdsClient, _ context.Context, _ *v1alpha1.StorageNode, _ *dbmeshv1alpha1.DatabaseClass) error {
				return nil
			})

			nodeName := "test-storage-node-ready"
			node := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      nodeName,
					Namespace: "default",
					Annotations: map[string]string{
						dbmeshv1alpha1.AnnotationsInstanceIdentifier: "test-instance-identifier",
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					DatabaseClassName: databaseClassName,
				},
			}

			// create resource
			Expect(k8sClient.Create(ctx, node)).Should(Succeed())

			// check storage node status
			Eventually(func() v1alpha1.StorageNodePhaseStatus {
				newSN := &v1alpha1.StorageNode{}
				Expect(k8sClient.Get(ctx, client.ObjectKey{Name: nodeName, Namespace: "default"}, newSN)).Should(Succeed())
				return newSN.Status.Phase
			}, 10*time.Second, 1*time.Second).Should(Equal(v1alpha1.StorageNodePhaseReady))

			// delete resource
			Expect(k8sClient.Delete(ctx, node)).Should(Succeed())
		})

		It("should delete success", func() {
			nodeName := "test-storage-node-delete"
			node := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      nodeName,
					Namespace: "default",
					Annotations: map[string]string{
						dbmeshv1alpha1.AnnotationsInstanceIdentifier: "test-instance-identifier",
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					DatabaseClassName: databaseClassName,
				},
			}
			Expect(k8sClient.Create(ctx, node)).Should(Succeed())

			getNode := &v1alpha1.StorageNode{}
			Expect(k8sClient.Get(ctx, client.ObjectKey{Name: nodeName, Namespace: "default"}, getNode)).Should(Succeed())

			// delete storage node
		})
	})
})
