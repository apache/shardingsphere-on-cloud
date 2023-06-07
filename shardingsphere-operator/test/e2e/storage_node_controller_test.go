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

package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"time"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/controllers"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode/aws"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	dbmesh_rds "github.com/database-mesh/golang-sdk/aws/client/rds"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("StorageNode Controller Suite Test For AWS RDS Instance", func() {
	storageProviderName := "test-storage-provider"
	instanceIdentifier := "test-instance-identifier"

	BeforeEach(func() {
		StorageProvider := &v1alpha1.StorageProvider{
			ObjectMeta: metav1.ObjectMeta{
				Name: storageProviderName,
			},
			Spec: v1alpha1.StorageProviderSpec{
				Provisioner: v1alpha1.ProvisionerAWSRDSInstance,
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

		Expect(k8sClient.Create(ctx, StorageProvider)).Should(Succeed())
	})

	AfterEach(func() {
		StorageProvider := &v1alpha1.StorageProvider{}
		Expect(k8sClient.Get(ctx, client.ObjectKey{Name: storageProviderName}, StorageProvider)).Should(Succeed())
		Expect(k8sClient.Delete(ctx, StorageProvider)).Should(Succeed())
	})

	Context("reconcile storageNode", func() {
		AfterEach(func() {
			monkey.UnpatchAll()
		})

		It("should create success", func() {
			// mock get instance func returns success
			monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "GetInstance", func(_ *aws.RdsClient, _ context.Context, _ *v1alpha1.StorageNode) (*dbmesh_rds.DescInstance, error) {
				return &dbmesh_rds.DescInstance{
					DBInstanceStatus: dbmesh_rds.DBInstanceStatusAvailable,
					Endpoint: dbmesh_rds.Endpoint{
						Address: "127.0.0.1",
						Port:    3306,
					},
				}, nil
			})
			// mock delete instance func returns success
			monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "DeleteInstance", func(_ *aws.RdsClient, _ context.Context, _ *v1alpha1.StorageNode, _ *v1alpha1.StorageProvider) error {
				return nil
			})

			nodeName := "test-storage-node-ready"
			node := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      nodeName,
					Namespace: "default",
					Annotations: map[string]string{
						v1alpha1.AnnotationsInstanceIdentifier: instanceIdentifier,
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: storageProviderName,
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
						v1alpha1.AnnotationsInstanceIdentifier: instanceIdentifier,
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: storageProviderName,
				},
			}
			Expect(k8sClient.Create(ctx, node)).Should(Succeed())

			getNode := &v1alpha1.StorageNode{}
			Expect(k8sClient.Get(ctx, client.ObjectKey{Name: nodeName, Namespace: "default"}, getNode)).Should(Succeed())

			// delete storage node
			Expect(k8sClient.Delete(ctx, getNode)).Should(Succeed())
			Eventually(func() bool {
				newSN := &v1alpha1.StorageNode{}
				err := k8sClient.Get(ctx, client.ObjectKey{Name: nodeName, Namespace: "default"}, newSN)
				return err != nil
			}, 10*time.Second, 1*time.Second).Should(BeTrue())
		})

		It("should register and unregister storage unit success", func() {
			// mock mysql
			db, dbmock, err := sqlmock.New()
			Expect(err).Should(Succeed())
			Expect(dbmock).ShouldNot(BeNil())
			// mock rds DescribeDBInstances func returns success
			g := monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "GetInstance", func(_ *aws.RdsClient, _ context.Context, _ *v1alpha1.StorageNode) (*dbmesh_rds.DescInstance, error) {
				return &dbmesh_rds.DescInstance{
					DBInstanceStatus: dbmesh_rds.DBInstanceStatusAvailable,
					Endpoint: dbmesh_rds.Endpoint{
						Address: "127.0.0.1",
						Port:    3306,
					},
				}, nil
			})
			monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "DeleteInstance", func(_ *aws.RdsClient, _ context.Context, _ *v1alpha1.StorageNode, _ *v1alpha1.StorageProvider) error {
				return nil
			})
			monkey.Patch(sql.Open, func(_ string, _ string) (*sql.DB, error) {
				return db, nil
			})
			monkey.PatchInstanceMethod(reflect.TypeOf(db), "Close", func(_ *sql.DB) error {
				return nil
			})

			cn := &v1alpha1.ComputeNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-compute-node",
					Namespace: "default",
					Labels: map[string]string{
						"app": "test-app",
					},
				},
				Spec: v1alpha1.ComputeNodeSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app": "test-app",
						},
					},
					PortBindings: []v1alpha1.PortBinding{
						{
							Name:          "http",
							ContainerPort: 3307,
							Protocol:      "TCP",
							ServicePort:   3307,
						},
					},
					Bootstrap: v1alpha1.BootstrapConfig{
						ServerConfig: v1alpha1.ServerConfig{
							Authority: v1alpha1.ComputeNodeAuthority{
								Users: []v1alpha1.ComputeNodeUser{
									{
										User:     "test-user",
										Password: "test-password",
									},
								},
								Privilege: v1alpha1.ComputeNodePrivilege{
									Type: v1alpha1.AllPermitted,
								},
							},
							Props: map[string]string{
								"proxy-frontend-database-protocol-type": "MySQL",
							},
							Mode: v1alpha1.ComputeNodeServerMode{
								Repository: v1alpha1.Repository{
									Type:  "ZooKeeper",
									Props: nil,
								},
								Type: "Zookeeper",
							},
						},
					},
				},
			}

			nodeName := "test-storage-node-register"
			node := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      nodeName,
					Namespace: "default",
					Annotations: map[string]string{
						v1alpha1.AnnotationsInstanceIdentifier:              instanceIdentifier,
						controllers.AnnotationKeyRegisterStorageUnitEnabled: "true",
						v1alpha1.AnnotationsInstanceDBName:                  "test-db-name",
						controllers.AnnotationKeyComputeNodeName:            "test-compute-node",
						controllers.AnnotationKeyLogicDatabaseName:          "test-logic-db-name",
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: storageProviderName,
				},
			}

			Expect(k8sClient.Create(ctx, cn)).Should(Succeed())
			Expect(k8sClient.Create(ctx, node)).Should(Succeed())

			dbmock.ExpectExec(regexp.QuoteMeta("CREATE DATABASE IF NOT EXISTS")).WillReturnResult(sqlmock.NewResult(1, 1))
			dbmock.ExpectExec(regexp.QuoteMeta("USE")).WillReturnResult(sqlmock.NewResult(1, 1))
			dbmock.ExpectExec(regexp.QuoteMeta("REGISTER STORAGE UNIT IF NOT EXISTS")).WillReturnResult(sqlmock.NewResult(0, 0))

			Eventually(func() v1alpha1.StorageNodePhaseStatus {
				newSN := &v1alpha1.StorageNode{}
				Expect(k8sClient.Get(ctx, client.ObjectKey{Name: nodeName, Namespace: "default"}, newSN)).Should(Succeed())
				return newSN.Status.Phase
			}, time.Second*10, time.Millisecond*250).Should(Equal(v1alpha1.StorageNodePhaseReady))

			Eventually(func() bool {
				newSN := &v1alpha1.StorageNode{}
				Expect(k8sClient.Get(ctx, client.ObjectKey{Name: nodeName, Namespace: "default"}, newSN)).Should(Succeed())
				return newSN.Status.Registered
			}, 20, 2).Should(BeTrue())

			// delete storage node
			Expect(k8sClient.Delete(ctx, node)).Should(Succeed())

			dbmock.ExpectExec(regexp.QuoteMeta("USE")).WillReturnResult(sqlmock.NewResult(1, 1))
			dbmock.ExpectQuery(regexp.QuoteMeta("SHOW RULES USED STORAGE UNIT")).WillReturnRows(sqlmock.NewRows([]string{"type", "name"}).AddRow("sharding", "t_order"))
			dbmock.ExpectExec("DROP SHARDING TABLE RULE").WillReturnResult(sqlmock.NewResult(1, 1))
			dbmock.ExpectExec(regexp.QuoteMeta("UNREGISTER STORAGE UNIT")).WillReturnResult(sqlmock.NewResult(0, 0))
			Eventually(func() v1alpha1.StorageNodePhaseStatus {
				newSN := &v1alpha1.StorageNode{}
				Expect(k8sClient.Get(ctx, client.ObjectKey{Name: nodeName, Namespace: "default"}, newSN)).Should(Succeed())
				return newSN.Status.Phase
			}, 20, 2).Should(Equal(v1alpha1.StorageNodePhaseDeleting))

			g.Unpatch()
			monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "GetInstance", func(_ *aws.RdsClient, _ context.Context, _ *v1alpha1.StorageNode) (*dbmesh_rds.DescInstance, error) {
				return nil, nil
			})

			Eventually(func() bool {
				newSN := &v1alpha1.StorageNode{}
				err := k8sClient.Get(ctx, client.ObjectKey{Name: nodeName, Namespace: "default"}, newSN)
				return apierrors.IsNotFound(err)
			}, 20, 2).Should(BeTrue())
		})
	})
})

var _ = Describe("StorageNode Controller Suite Test For AWS Aurora Cluster", func() {
	storageProviderName := "test-storage-provider"
	clusterIdentifier := "test-aurora-cluster-identifier"

	BeforeEach(func() {
		provider := &v1alpha1.StorageProvider{
			ObjectMeta: metav1.ObjectMeta{
				Name: storageProviderName,
			},
			Spec: v1alpha1.StorageProviderSpec{
				Provisioner: v1alpha1.ProvisionerAWSAurora,
				Parameters: map[string]string{
					"engine":             "aurora-mysql",
					"engineVersion":      "5.7",
					"instanceClass":      "db.t2.small",
					"masterUsername":     "test-user",
					"masterUserPassword": "test-password",
				},
			},
		}
		Expect(k8sClient.Create(ctx, provider)).Should(Succeed())
	})

	AfterEach(func() {
		monkey.UnpatchAll()

		StorageProvider := &v1alpha1.StorageProvider{}
		Expect(k8sClient.Get(ctx, client.ObjectKey{Name: storageProviderName}, StorageProvider)).Should(Succeed())
		Expect(k8sClient.Delete(ctx, StorageProvider)).Should(Succeed())
	})

	Context("When Creat StorageNode with Aurora Cluster Not Exist", func() {
		It("Should Success", func() {
			snName := "test-storage-node-creating"
			// monkey patch
			monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "GetAuroraCluster", func(_ *aws.RdsClient, _ context.Context, _ *v1alpha1.StorageNode) (*dbmesh_rds.DescCluster, error) {
				return &dbmesh_rds.DescCluster{
					DBClusterIdentifier: clusterIdentifier,
					Status:              dbmesh_rds.DBClusterStatusCreating,
					PrimaryEndpoint:     "test-primary-endpoint",
					ReaderEndpoint:      "test-reader-endpoint",
					Port:                3306,
				}, nil
			})
			monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "GetInstancesByFilters", func(_ *aws.RdsClient, _ context.Context, _ map[string][]string) ([]*dbmesh_rds.DescInstance, error) {
				return []*dbmesh_rds.DescInstance{
					{
						DBInstanceIdentifier: fmt.Sprintf("%s-insatnce-0", clusterIdentifier),
						DBInstanceStatus:     dbmesh_rds.DBInstanceStatusCreating,
						Endpoint:             dbmesh_rds.Endpoint{Address: "test-instance-0-endpoint", Port: 3306},
					},
				}, nil
			})

			sn := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      snName,
					Namespace: "default",
					Annotations: map[string]string{
						v1alpha1.AnnotationsClusterIdentifier: clusterIdentifier,
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: storageProviderName,
					Replicas:            2,
				},
			}
			Expect(k8sClient.Create(ctx, sn)).Should(Succeed())

			Eventually(func() string {
				newSN := &v1alpha1.StorageNode{}
				Expect(k8sClient.Get(ctx, client.ObjectKey{Name: snName, Namespace: "default"}, newSN)).Should(Succeed())
				return newSN.Status.Cluster.Status
			}, time.Second*10, time.Millisecond*250).Should(Equal(dbmesh_rds.DBClusterStatusCreating))
		})

		It("should success when cluster is available", func() {
			snName := "test-storage-node-available"
			monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "GetAuroraCluster", func(_ *aws.RdsClient, _ context.Context, _ *v1alpha1.StorageNode) (*dbmesh_rds.DescCluster, error) {
				return &dbmesh_rds.DescCluster{
					DBClusterIdentifier: clusterIdentifier,
					Status:              dbmesh_rds.DBClusterStatusAvailable,
					PrimaryEndpoint:     "test-primary-endpoint",
					ReaderEndpoint:      "test-reader-endpoint",
					Port:                3306,
				}, nil
			})
			monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "GetInstancesByFilters", func(_ *aws.RdsClient, _ context.Context, _ map[string][]string) ([]*dbmesh_rds.DescInstance, error) {
				return []*dbmesh_rds.DescInstance{
					{
						DBInstanceIdentifier: fmt.Sprintf("%s-insatnce-0", clusterIdentifier),
						DBInstanceStatus:     dbmesh_rds.DBInstanceStatusAvailable,
						Endpoint:             dbmesh_rds.Endpoint{Address: "test-instance-0-endpoint", Port: 3306},
					},
					{
						DBInstanceIdentifier: fmt.Sprintf("%s-insatnce-1", clusterIdentifier),
						DBInstanceStatus:     dbmesh_rds.DBInstanceStatusAvailable,
						Endpoint:             dbmesh_rds.Endpoint{Address: "test-instance-1-endpoint", Port: 3306},
					},
				}, nil
			})
			monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "DeleteAuroraCluster", func(_ *aws.RdsClient, _ context.Context, _ *v1alpha1.StorageNode, _ *v1alpha1.StorageProvider) error {
				return nil
			})

			sn := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      snName,
					Namespace: "default",
					Annotations: map[string]string{
						v1alpha1.AnnotationsClusterIdentifier: clusterIdentifier,
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: storageProviderName,
					Replicas:            2,
				},
			}
			Expect(k8sClient.Create(ctx, sn)).Should(Succeed())

			newSN := &v1alpha1.StorageNode{}
			Eventually(func() v1alpha1.StorageNodePhaseStatus {
				Expect(k8sClient.Get(ctx, client.ObjectKey{Name: snName, Namespace: "default"}, newSN)).Should(Succeed())
				return newSN.Status.Phase
			}, time.Second*10, time.Millisecond*250).Should(Equal(v1alpha1.StorageNodePhaseReady))

			Expect(newSN.Status.Instances).Should(HaveLen(2))
		})
	})
})
