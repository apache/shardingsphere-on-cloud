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

package controllers

import (
	"context"
	"time"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/service"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode/aws"
	mock_aws "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode/aws/mocks"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/shardingsphere"
	mock_shardingsphere "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/shardingsphere/mocks"

	"bou.ke/monkey"
	dbmesh_aws "github.com/database-mesh/golang-sdk/aws"
	dbmesh_rds "github.com/database-mesh/golang-sdk/aws/client/rds"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

const (
	defaultTestNamespace          = "test-namespace"
	defaultTestStorageProvider    = "test-storage-provider"
	defaultTestStorageNode        = "test-storage-node"
	defaultTestInstanceIdentifier = "test-database-instance"
)

var (
	ctx        = context.Background()
	fakeClient client.Client
	reconciler *StorageNodeReconciler
	mockCtrl   *gomock.Controller
	mockAws    *mock_aws.MockIRdsClient
	mockSS     *mock_shardingsphere.MockIServer
)

func fakeStorageNodeReconciler() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	scheme := runtime.NewScheme()
	Expect(v1alpha1.AddToScheme(scheme)).To(Succeed())
	Expect(corev1.AddToScheme(scheme)).To(Succeed())
	fakeClient = fake.NewClientBuilder().WithScheme(scheme).Build()

	sess := dbmesh_aws.NewSessions().SetCredential("AwsRegion", "AwsAccessKeyID", "AwsSecretAccessKey").Build()
	reconciler = &StorageNodeReconciler{
		Client:   fakeClient,
		Log:      logf.Log,
		Recorder: record.NewFakeRecorder(100),
		AwsRDS:   dbmesh_rds.NewService(sess["AwsRegion"]),
		Service:  service.NewServiceClient(fakeClient),
	}
}

var _ = BeforeEach(func() {
	fakeStorageNodeReconciler()
})

var _ = Describe("StorageNode Controller Mock Test For AWS Rds Instance", func() {
	BeforeEach(func() {
		// mock aws rds client
		mockCtrl = gomock.NewController(GinkgoT())
		mockAws = mock_aws.NewMockIRdsClient(mockCtrl)

		monkey.Patch(aws.NewRdsClient, func(rds dbmesh_rds.RDS) aws.IRdsClient {
			return mockAws
		})

		// create default resource
		dbClass := &v1alpha1.StorageProvider{
			ObjectMeta: metav1.ObjectMeta{
				Name: defaultTestStorageProvider,
			},
			Spec: v1alpha1.StorageProviderSpec{
				Provisioner: v1alpha1.ProvisionerAWSRDSInstance,
			},
		}

		storageNode := &v1alpha1.StorageNode{
			ObjectMeta: metav1.ObjectMeta{
				Name:      defaultTestStorageNode,
				Namespace: defaultTestNamespace,
				Annotations: map[string]string{
					v1alpha1.AnnotationsInstanceIdentifier: defaultTestInstanceIdentifier,
				},
			},
			Spec: v1alpha1.StorageNodeSpec{
				StorageProviderName: defaultTestStorageProvider,
			},
		}
		Expect(fakeClient.Create(ctx, dbClass)).Should(Succeed())
		Expect(fakeClient.Create(ctx, storageNode)).Should(Succeed())
	})

	AfterEach(func() {
		// delete default resource
		Expect(fakeClient.Delete(ctx, &v1alpha1.StorageNode{
			ObjectMeta: metav1.ObjectMeta{
				Name:      defaultTestStorageNode,
				Namespace: defaultTestNamespace,
			},
		})).Should(Succeed())
		Expect(fakeClient.Delete(ctx, &v1alpha1.StorageProvider{
			ObjectMeta: metav1.ObjectMeta{
				Name: defaultTestStorageProvider,
			},
		})).Should(Succeed())

		mockCtrl.Finish()
		monkey.UnpatchAll()
	})

	Context("create storage node", func() {
		It("should create storage node successfully", func() {
			storageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-storage-node-1",
					Namespace: defaultTestNamespace,
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: defaultTestStorageProvider,
				},
				Status: v1alpha1.StorageNodeStatus{},
			}
			Expect(fakeClient.Create(ctx, storageNode)).Should(Succeed())
			sn := &v1alpha1.StorageNode{}
			Expect(fakeClient.Get(ctx, client.ObjectKey{Name: "test-storage-node", Namespace: "test-namespace"}, sn)).Should(Succeed())
			Expect(fakeClient.Delete(ctx, storageNode)).Should(Succeed())
		})
	})

	Context("create storage node with unknown StorageClassName", func() {
		It("should create storage node successfully", func() {
			storageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-storage-node-2",
					Namespace: defaultTestNamespace,
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: "no-database",
				},
				Status: v1alpha1.StorageNodeStatus{},
			}
			Expect(fakeClient.Create(ctx, storageNode)).Should(Succeed())
			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      "test-storage-node-2",
					Namespace: defaultTestNamespace,
				},
			}
			_, err := reconciler.Reconcile(ctx, req)
			Expect(client.IgnoreNotFound(err)).Should(Succeed())
			Expect(fakeClient.Delete(ctx, storageNode)).Should(Succeed())
		})
	})

	Context("reconcile storageNode", func() {
		It("should reconcile successfully with Creating Instance", func() {
			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      defaultTestStorageNode,
					Namespace: defaultTestNamespace,
				},
			}

			rdsInstance := &dbmesh_rds.DescInstance{
				DBInstanceStatus: dbmesh_rds.DBInstanceStatusCreating,
				Endpoint: dbmesh_rds.Endpoint{
					Address: "127.0.0.1",
					Port:    3306,
				},
			}

			// mock aws rds client
			mockAws.EXPECT().GetInstance(gomock.Any(), gomock.Any()).Return(rdsInstance, nil).AnyTimes()

			_, err := reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())

			newSN := &v1alpha1.StorageNode{}
			Expect(fakeClient.Get(ctx, client.ObjectKey{Name: "test-storage-node", Namespace: "test-namespace"}, newSN)).Should(Succeed())
			Expect(newSN.Status.Phase).To(Equal(v1alpha1.StorageNodePhaseNotReady))
			Expect(newSN.Status.Instances).To(HaveLen(1))
			Expect(newSN.Status.Instances[0].Status).To(Equal(string(dbmesh_rds.DBInstanceStatusCreating)))
		})

		It("should reconcile successfully with Available Instance", func() {
			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      defaultTestStorageNode,
					Namespace: defaultTestNamespace,
				},
			}

			rdsInstance := &dbmesh_rds.DescInstance{
				DBInstanceStatus: dbmesh_rds.DBInstanceStatusAvailable,
				Endpoint: dbmesh_rds.Endpoint{
					Address: "127.0.0.1",
					Port:    3306,
				},
			}

			// mock aws rds client
			mockAws.EXPECT().GetInstance(gomock.Any(), gomock.Any()).Return(rdsInstance, nil)
			_, err := reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())

			newSN := &v1alpha1.StorageNode{}
			Expect(fakeClient.Get(ctx, client.ObjectKey{Name: defaultTestStorageNode, Namespace: defaultTestNamespace}, newSN)).Should(Succeed())

			Expect(newSN.Status.Phase).To(Equal(v1alpha1.StorageNodePhaseReady))
			Expect(newSN.Status.Instances).To(HaveLen(1))
			Expect(newSN.Status.Instances[0].Status).To(Equal(string(dbmesh_rds.DBInstanceStatusAvailable)))
		})
	})

	Context("reconcile storage node in Ready status when it's been deleted", func() {
		rdsInstanceAvailable := dbmesh_rds.DescInstance{
			DBInstanceIdentifier: defaultTestInstanceIdentifier,
			DBInstanceStatus:     dbmesh_rds.DBInstanceStatusAvailable,
			Endpoint: dbmesh_rds.Endpoint{
				Address: "127.0.0.1",
				Port:    3306,
			},
		}
		It("should be successful when instance is in available status", func() {
			deletingStorageNode := "test-deleting-storage-node"
			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      deletingStorageNode,
					Namespace: defaultTestNamespace,
				},
			}
			readyStorageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      deletingStorageNode,
					Namespace: defaultTestNamespace,
					Annotations: map[string]string{
						v1alpha1.AnnotationsInstanceIdentifier: defaultTestInstanceIdentifier,
					},
				},
				Spec: v1alpha1.StorageNodeSpec{StorageProviderName: defaultTestStorageProvider},
			}
			Expect(fakeClient.Create(ctx, readyStorageNode)).Should(Succeed())
			// mock aws rds client, get instance and return available status
			mockAws.EXPECT().GetInstance(gomock.Any(), gomock.Any()).Return(&rdsInstanceAvailable, nil)
			// reconcile storage node, add instance and set status to ready
			_, err := reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())

			// delete storage node
			Expect(fakeClient.Delete(ctx, &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      deletingStorageNode,
					Namespace: defaultTestNamespace,
				},
			})).Should(Succeed())

			// mock aws rds client, delete instance
			mockAws.EXPECT().GetInstance(gomock.Any(), gomock.Any()).Return(&rdsInstanceAvailable, nil)
			mockAws.EXPECT().DeleteInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

			_, err = reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())

			// storage node status should be deleting
			deletingSN := &v1alpha1.StorageNode{}
			Expect(fakeClient.Get(ctx, client.ObjectKey{Name: deletingStorageNode, Namespace: defaultTestNamespace}, deletingSN)).Should(Succeed())
			Expect(deletingSN.Status.Phase).To(Equal(v1alpha1.StorageNodePhaseDeleting))
		})

		It("should be successful when instance is in deleting status", func() {
			deletedStorageNodeName := "test-deleted-storage-node"
			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      deletedStorageNodeName,
					Namespace: defaultTestNamespace,
				},
			}
			deleteTime := metav1.NewTime(time.Now())
			storageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      deletedStorageNodeName,
					Namespace: defaultTestNamespace,
					Annotations: map[string]string{
						v1alpha1.AnnotationsInstanceIdentifier: defaultTestInstanceIdentifier,
					},
					Finalizers: []string{
						FinalizerName,
					},
					DeletionTimestamp: &deleteTime,
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: defaultTestStorageProvider,
				},
				Status: v1alpha1.StorageNodeStatus{
					Phase: v1alpha1.StorageNodePhaseDeleting,
					Instances: []v1alpha1.InstanceStatus{
						{
							Status: string(dbmesh_rds.DBInstanceStatusDeleting),
							Endpoint: v1alpha1.Endpoint{
								Address: "127.0.0.1",
								Port:    3306,
							},
						},
					},
				},
			}
			Expect(fakeClient.Create(ctx, storageNode)).Should(Succeed())
			// mock aws rds client, get nil instance
			mockAws.EXPECT().GetInstance(gomock.Any(), gomock.Any()).Return(nil, nil)
			_, err := reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())

			deletedCompleteSN := &v1alpha1.StorageNode{}
			Expect(fakeClient.Get(ctx, client.ObjectKey{Name: deletedStorageNodeName, Namespace: defaultTestNamespace}, deletedCompleteSN)).Should(Succeed())
			Expect(deletedCompleteSN.Status.Phase).To(Equal(v1alpha1.StorageNodePhaseDeleteComplete))
		})

		It("should be successful when storage node is delete complete status", func() {
			deletedCompletedStorageNodeName := "test-delete-completed-storage-node"
			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      deletedCompletedStorageNodeName,
					Namespace: defaultTestNamespace,
				},
			}
			deleteTime := metav1.NewTime(time.Now())
			storageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      deletedCompletedStorageNodeName,
					Namespace: defaultTestNamespace,
					Finalizers: []string{
						FinalizerName,
					},
					DeletionTimestamp: &deleteTime,
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: defaultTestStorageProvider,
				},
				Status: v1alpha1.StorageNodeStatus{
					Phase: v1alpha1.StorageNodePhaseDeleteComplete,
				},
			}
			Expect(fakeClient.Create(ctx, storageNode)).Should(Succeed())

			_, err := reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())
			finalSN := &v1alpha1.StorageNode{}
			err = fakeClient.Get(ctx, client.ObjectKey{Name: deletedCompletedStorageNodeName, Namespace: defaultTestNamespace}, finalSN)
			Expect(apierrors.IsNotFound(err)).To(BeTrue())
		})
	})

	Context("Test register storage node", func() {
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockSS = mock_shardingsphere.NewMockIServer(mockCtrl)
			monkey.Patch(shardingsphere.NewServer, func(_, _ string, _ uint, _, _ string) (shardingsphere.IServer, error) {
				return mockSS, nil
			})
		})
		AfterEach(func() {
			mockCtrl.Finish()
		})
		It("should be successful when storage node is not registered", func() {
			nodeName := "test register storage node"
			cnName := "test-compute-node"
			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      nodeName,
					Namespace: defaultTestNamespace,
				},
			}
			storageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      nodeName,
					Namespace: defaultTestNamespace,
					Annotations: map[string]string{
						AnnotationKeyRegisterStorageUnitEnabled: "true",
						v1alpha1.AnnotationsInstanceDBName:      "test_db",
						AnnotationKeyComputeNodeName:            cnName,
						AnnotationKeyLogicDatabaseName:          "sharding_db",
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: defaultTestStorageProvider,
				},
			}
			ins := &dbmesh_rds.DescInstance{
				DBInstanceIdentifier: "ins-test-register-storage-node",
				DBInstanceStatus:     dbmesh_rds.DBInstanceStatusAvailable,
				Endpoint: dbmesh_rds.Endpoint{
					Address: "127.0.0.1",
					Port:    3306,
				},
			}
			cn := &v1alpha1.ComputeNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      cnName,
					Namespace: defaultTestNamespace,
				},
				Spec: v1alpha1.ComputeNodeSpec{
					Bootstrap: v1alpha1.BootstrapConfig{
						ServerConfig: v1alpha1.ServerConfig{
							Authority: v1alpha1.ComputeNodeAuthority{
								Users: []v1alpha1.ComputeNodeUser{
									{
										User:     "root@%",
										Password: "root",
									},
								},
								Privilege: v1alpha1.ComputeNodePrivilege{
									Type: v1alpha1.AllPermitted,
								},
							},
							Props: map[string]string{
								ShardingSphereProtocolType: "MySQL",
							},
						},
					},
				},
			}

			svc := &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      cnName,
					Namespace: defaultTestNamespace,
				},
				Spec: corev1.ServiceSpec{
					Ports: []corev1.ServicePort{
						{
							Name: "shardingsphere-proxy",
							Port: 3306,
						},
					},
					ClusterIP: "127.0.0.1",
					Type:      corev1.ServiceTypeClusterIP,
				},
			}

			Expect(fakeClient.Create(ctx, storageNode)).Should(Succeed())
			Expect(fakeClient.Create(ctx, cn)).Should(Succeed())
			Expect(fakeClient.Create(ctx, svc)).Should(Succeed())

			// mock aws rds client, get available instance
			mockAws.EXPECT().GetInstance(gomock.Any(), gomock.Any()).Return(ins, nil).AnyTimes()
			// mock shardingsphere create database
			mockSS.EXPECT().CreateDatabase(gomock.Any()).Return(nil)
			// mock shardingsphere register storage unit
			mockSS.EXPECT().RegisterStorageUnit(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			// mock shardingsphere close connection
			mockSS.EXPECT().Close()

			// reconcile storage node status to Ready
			_, err := reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())
			// reconcile to register storage node
			_, err = reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())

			registeredSN := &v1alpha1.StorageNode{}
			Expect(fakeClient.Get(ctx, client.ObjectKey{Name: nodeName, Namespace: defaultTestNamespace}, registeredSN)).Should(Succeed())
			Expect(registeredSN.Status.Registered).To(BeTrue())
		})
	})

	Context("Test getShardingsphereServer", func() {
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockSS = mock_shardingsphere.NewMockIServer(mockCtrl)
			monkey.Patch(shardingsphere.NewServer, func(_, _ string, _ uint, _, _ string) (shardingsphere.IServer, error) {
				return mockSS, nil
			})
		})
		AfterEach(func() {
			mockCtrl.Finish()
			monkey.UnpatchAll()
		})
		It("should be successful when get shardingsphere server", func() {
			cn := &v1alpha1.ComputeNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-get-shardingsphere-server",
					Namespace: defaultTestNamespace,
				},
				Spec: v1alpha1.ComputeNodeSpec{
					Bootstrap: v1alpha1.BootstrapConfig{
						ServerConfig: v1alpha1.ServerConfig{
							Authority: v1alpha1.ComputeNodeAuthority{
								Users: []v1alpha1.ComputeNodeUser{
									{
										User:     "root",
										Password: "root",
									},
								},
							},
						},
					},
				},
			}
			svc := &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-get-shardingsphere-server",
					Namespace: defaultTestNamespace,
				},
				Spec: corev1.ServiceSpec{
					Ports: []corev1.ServicePort{
						{
							Name:     "http",
							Protocol: "TCP",
							Port:     3307,
						},
					},
				},
			}
			Expect(fakeClient.Create(ctx, cn)).Should(Succeed())
			Expect(fakeClient.Create(ctx, svc)).Should(Succeed())

			sn := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-get-shardingsphere-server",
					Namespace: defaultTestNamespace,
					Annotations: map[string]string{
						AnnotationKeyComputeNodeName: "test-get-shardingsphere-server",
					},
				},
			}

			ss, err := reconciler.getShardingsphereServer(ctx, sn)
			Expect(err).To(BeNil())
			Expect(ss).ToNot(BeNil())
		})
	})

	Context("Test registerStorageUnit", func() {
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockSS = mock_shardingsphere.NewMockIServer(mockCtrl)
			monkey.Patch(shardingsphere.NewServer, func(_, _ string, _ uint, _, _ string) (shardingsphere.IServer, error) {
				return mockSS, nil
			})
		})
		AfterEach(func() {
			mockCtrl.Finish()
			monkey.UnpatchAll()
		})
		It("should be successful when register storage unit", func() {
			testName := "test-register-storage-unit"
			cn := &v1alpha1.ComputeNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testName,
					Namespace: defaultTestNamespace,
				},
				Spec: v1alpha1.ComputeNodeSpec{
					Bootstrap: v1alpha1.BootstrapConfig{
						ServerConfig: v1alpha1.ServerConfig{
							Authority: v1alpha1.ComputeNodeAuthority{
								Users: []v1alpha1.ComputeNodeUser{
									{
										User:     "root",
										Password: "root",
									},
								},
							},
						},
					},
				},
			}
			svc := &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testName,
					Namespace: defaultTestNamespace,
				},
				Spec: corev1.ServiceSpec{
					Ports: []corev1.ServicePort{
						{
							Name:     "http",
							Protocol: "TCP",
							Port:     3307,
						},
					},
				},
			}
			Expect(fakeClient.Create(ctx, cn)).Should(Succeed())
			Expect(fakeClient.Create(ctx, svc)).Should(Succeed())

			sn := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testName,
					Namespace: defaultTestNamespace,
					Annotations: map[string]string{
						AnnotationKeyComputeNodeName:            testName,
						AnnotationKeyRegisterStorageUnitEnabled: "true",
						AnnotationKeyLogicDatabaseName:          testName,
						v1alpha1.AnnotationsInstanceDBName:      testName,
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: defaultTestStorageProvider,
				},
				Status: v1alpha1.StorageNodeStatus{
					Phase: v1alpha1.StorageNodePhaseReady,
					Instances: []v1alpha1.InstanceStatus{
						{
							Status:   string(dbmesh_rds.DBInstanceStatusAvailable),
							Endpoint: v1alpha1.Endpoint{},
						},
					},
				},
			}

			storageProvider := &v1alpha1.StorageProvider{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultTestStorageProvider,
				},
				Spec: v1alpha1.StorageProviderSpec{
					Provisioner: v1alpha1.ProvisionerAWSRDSInstance,
					Parameters: map[string]string{
						"masterUsername":     testName,
						"masterUserPassword": testName,
					},
				},
			}

			mockSS.EXPECT().CreateDatabase(gomock.Any()).Return(nil)
			mockSS.EXPECT().Close().Return(nil)
			mockSS.EXPECT().RegisterStorageUnit(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

			Expect(reconciler.registerStorageUnit(ctx, sn, storageProvider)).To(BeNil())
			Expect(sn.Status.Registered).To(BeTrue())
		})

		Context("Test unregisterStorageUnit", func() {
			BeforeEach(func() {
				mockCtrl = gomock.NewController(GinkgoT())
				mockSS = mock_shardingsphere.NewMockIServer(mockCtrl)
				monkey.Patch(shardingsphere.NewServer, func(_, _ string, _ uint, _, _ string) (shardingsphere.IServer, error) {
					return mockSS, nil
				})
			})
			AfterEach(func() {
				mockCtrl.Finish()
				monkey.UnpatchAll()
			})
			It("should be successful when unregister storage unit", func() {
				testName := "test-unregister-storage-unit"

				cn := &v1alpha1.ComputeNode{
					ObjectMeta: metav1.ObjectMeta{
						Name:      testName,
						Namespace: defaultTestNamespace,
					},
					Spec: v1alpha1.ComputeNodeSpec{
						Bootstrap: v1alpha1.BootstrapConfig{
							ServerConfig: v1alpha1.ServerConfig{
								Authority: v1alpha1.ComputeNodeAuthority{
									Users: []v1alpha1.ComputeNodeUser{
										{
											User:     "root",
											Password: "root",
										},
									},
								},
							},
						},
					},
				}
				svc := &corev1.Service{
					ObjectMeta: metav1.ObjectMeta{
						Name:      testName,
						Namespace: defaultTestNamespace,
					},
					Spec: corev1.ServiceSpec{
						Ports: []corev1.ServicePort{
							{
								Name:     "http",
								Protocol: "TCP",
								Port:     3307,
							},
						},
					},
				}
				Expect(fakeClient.Create(ctx, cn)).Should(Succeed())
				Expect(fakeClient.Create(ctx, svc)).Should(Succeed())

				sn := &v1alpha1.StorageNode{
					ObjectMeta: metav1.ObjectMeta{
						Name:      testName,
						Namespace: defaultTestNamespace,
						Annotations: map[string]string{
							AnnotationKeyLogicDatabaseName:     testName,
							v1alpha1.AnnotationsInstanceDBName: testName,
							AnnotationKeyComputeNodeName:       testName,
						},
					},
					Status: v1alpha1.StorageNodeStatus{
						Registered: true,
					},
				}
				Expect(fakeClient.Create(ctx, sn)).Should(Succeed())

				mockSS.EXPECT().UnRegisterStorageUnit(gomock.Any(), gomock.Any()).Return(nil)
				mockSS.EXPECT().Close().Return(nil)
				Expect(reconciler.unregisterStorageUnit(ctx, sn)).To(BeNil())
			})
		})
	})
})

var _ = Describe("StorageNode Controller Mock Test For AWS Aurora", func() {
	var provider *v1alpha1.StorageProvider
	BeforeEach(func() {
		provider = &v1alpha1.StorageProvider{
			ObjectMeta: metav1.ObjectMeta{
				Name: "aws-aurora",
			},
			Spec: v1alpha1.StorageProviderSpec{
				Provisioner: v1alpha1.ProvisionerAWSAurora,
				Parameters: map[string]string{
					"engine":             "aurora-mysql",
					"engineVersion":      "5.7",
					"masterUsername":     "root",
					"masterUserPassword": "root",
				},
			},
		}
		Expect(fakeClient.Create(ctx, provider)).Should(Succeed())

		// mock aws client
		// mock aws rds client
		mockCtrl = gomock.NewController(GinkgoT())
		mockAws = mock_aws.NewMockIRdsClient(mockCtrl)
		monkey.Patch(aws.NewRdsClient, func(rds dbmesh_rds.RDS) aws.IRdsClient {
			return mockAws
		})
		mockSS = mock_shardingsphere.NewMockIServer(mockCtrl)
		monkey.Patch(shardingsphere.NewServer, func(_, _ string, _ uint, _, _ string) (shardingsphere.IServer, error) {
			return mockSS, nil
		})
	})

	AfterEach(func() {
		mockCtrl.Finish()
		monkey.UnpatchAll()
	})

	Context("reconcile storage node", func() {
		It("should success when aws aurora cluster is not exits", func() {
			name := "test-aws-aurora-not-exists"
			namespacedName := types.NamespacedName{
				Name:      name,
				Namespace: defaultTestNamespace,
			}
			storageNode := v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: defaultTestNamespace,
					Annotations: map[string]string{
						v1alpha1.AnnotationsClusterIdentifier: "test-aws-aurora",
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: "aws-aurora",
				},
			}
			Expect(fakeClient.Create(ctx, &storageNode)).Should(Succeed())

			descCluster := &dbmesh_rds.DescCluster{
				DBClusterIdentifier: "test-aws-aurora",
				PrimaryEndpoint:     "test-aws-aurora.cluster-xxxxxx.us-east-1.rds.amazonaws.com",
				ReaderEndpoint:      "test-aws-aurora.cluster-ro-xxxxxx.us-east-1.rds.amazonaws.com",
				Port:                3306,
				Status:              dbmesh_rds.DBClusterStatusAvailable,
			}
			descInstance := &dbmesh_rds.DescInstance{
				DBInstanceIdentifier: "test-aws-aurora-1",
				DBClusterIdentifier:  "test-aws-aurora",
				Endpoint: dbmesh_rds.Endpoint{
					Address: "test-aws-aurora-1.cluster-xxxxxx.us-east-1.rds.amazonaws.com",
					Port:    3306,
				},
				DBInstanceStatus: dbmesh_rds.DBInstanceStatusAvailable,
			}

			// mock aws aurora cluster is not exist
			mockAws.EXPECT().GetAuroraCluster(gomock.Any(), gomock.Any()).Return(nil, nil).Times(1)
			// mock create aws aurora cluster
			mockAws.EXPECT().CreateAuroraCluster(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			// mock aws aurora cluster is created
			mockAws.EXPECT().GetAuroraCluster(gomock.Any(), gomock.Any()).Return(descCluster, nil).Times(1)
			// mock aws instance is created
			mockAws.EXPECT().GetInstancesByFilters(gomock.Any(), gomock.Any()).Return([]*dbmesh_rds.DescInstance{descInstance}, nil).Times(1)

			req := ctrl.Request{NamespacedName: namespacedName}
			_, err := reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())
			sn := &v1alpha1.StorageNode{}
			Expect(fakeClient.Get(ctx, namespacedName, sn)).Should(Succeed())
			Expect(sn.Status.Phase).To(Equal(v1alpha1.StorageNodePhaseReady))
		})

		It("should success when storage node been delete", func() {
			name := "test-aws-aurora-deleted"
			namespacedName := types.NamespacedName{
				Name:      name,
				Namespace: defaultTestNamespace,
			}
			req := ctrl.Request{NamespacedName: namespacedName}
			storageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: defaultTestNamespace,
					Annotations: map[string]string{
						v1alpha1.AnnotationsClusterIdentifier: "test-aws-aurora",
					},
					Finalizers: []string{FinalizerName},
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: "aws-aurora",
				},
				Status: v1alpha1.StorageNodeStatus{
					Phase: v1alpha1.StorageNodePhaseReady,
					Cluster: v1alpha1.ClusterStatus{
						Status:          dbmesh_rds.DBClusterStatusAvailable,
						PrimaryEndpoint: v1alpha1.Endpoint{Address: "test-aws-aurora.cluster-xxxxxx.us-east-1.rds.amazonaws.com", Port: 3306},
						ReaderEndpoints: []v1alpha1.Endpoint{{Address: "test-aws-aurora.cluster-ro-xxxxxx.us-east-1.rds.amazonaws.com", Port: 3306}},
					},
					Instances: []v1alpha1.InstanceStatus{
						{
							Status:   string(dbmesh_rds.DBInstanceStatusAvailable),
							Endpoint: v1alpha1.Endpoint{Address: "test-aws-aurora-1.cluster-xxxxxx.us-east-1.rds.amazonaws.com", Port: 3306},
						},
					},
				},
			}

			Expect(fakeClient.Create(ctx, storageNode)).Should(Succeed())

			descCluster := &dbmesh_rds.DescCluster{
				DBClusterIdentifier: "test-aws-aurora",
				PrimaryEndpoint:     "test-aws-aurora.cluster-xxxxxx.us-east-1.rds.amazonaws.com",
				ReaderEndpoint:      "test-aws-aurora.cluster-ro-xxxxxx.us-east-1.rds.amazonaws.com",
				Port:                3306,
				Status:              dbmesh_rds.DBClusterStatusAvailable,
			}

			descInstance := &dbmesh_rds.DescInstance{
				DBInstanceIdentifier: "test-aws-aurora-1",
				DBClusterIdentifier:  "test-aws-aurora",
				Endpoint: dbmesh_rds.Endpoint{
					Address: "test-aws-aurora-1.cluster-xxxxxx.us-east-1.rds.amazonaws.com",
					Port:    3306,
				},
				DBInstanceStatus: dbmesh_rds.DBInstanceStatusAvailable,
			}

			Expect(fakeClient.Delete(ctx, storageNode)).Should(Succeed())

			// mock aws aurora is exists
			mockAws.EXPECT().GetAuroraCluster(gomock.Any(), gomock.Any()).Return(descCluster, nil).Times(1)
			// mock get instances of aws aurora
			mockAws.EXPECT().GetInstancesByFilters(gomock.Any(), gomock.Any()).Return([]*dbmesh_rds.DescInstance{descInstance}, nil).Times(1)
			// mock delete aws aurora cluster
			mockAws.EXPECT().DeleteAuroraCluster(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

			_, err := reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())

			Expect(fakeClient.Get(ctx, namespacedName, storageNode)).Should(Succeed())
			Expect(storageNode.DeletionTimestamp).NotTo(BeNil())
			Expect(storageNode.Status.Phase).To(Equal(v1alpha1.StorageNodePhaseDeleting))
		})

		It("should be success when storage node is deleting", func() {
			name := "test-aws-aurora-deleting"
			namespacedName := types.NamespacedName{
				Name:      name,
				Namespace: defaultTestNamespace,
			}
			req := ctrl.Request{NamespacedName: namespacedName}
			deletionTimestamp := metav1.Now()
			storageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: defaultTestNamespace,
					Annotations: map[string]string{
						v1alpha1.AnnotationsClusterIdentifier: "test-aws-aurora",
					},
					Finalizers:        []string{FinalizerName},
					DeletionTimestamp: &deletionTimestamp,
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: "aws-aurora",
				},
				Status: v1alpha1.StorageNodeStatus{
					Phase: v1alpha1.StorageNodePhaseDeleting,
					Cluster: v1alpha1.ClusterStatus{
						Status:          dbmesh_rds.DBClusterStatusDeleting,
						PrimaryEndpoint: v1alpha1.Endpoint{Address: "test-aws-aurora.cluster-xxxxxx.us-east-1.rds.amazonaws.com", Port: 3306},
						ReaderEndpoints: []v1alpha1.Endpoint{{Address: "test-aws-aurora.cluster-ro-xxxxxx.us-east-1.rds.amazonaws.com", Port: 3306}},
					},
					Instances: []v1alpha1.InstanceStatus{
						{
							Status:   string(dbmesh_rds.DBInstanceStatusDeleting),
							Endpoint: v1alpha1.Endpoint{Address: "test-aws-aurora-1.cluster-xxxxxx.us-east-1.rds.amazonaws.com", Port: 3306},
						},
					},
				},
			}
			Expect(fakeClient.Create(ctx, storageNode)).Should(Succeed())

			// mock aws aurora is not exists
			mockAws.EXPECT().GetAuroraCluster(gomock.Any(), gomock.Any()).Return(nil, nil).Times(1)
			// mock get instances of aws aurora is not exists
			mockAws.EXPECT().GetInstancesByFilters(gomock.Any(), gomock.Any()).Return(nil, nil).Times(1)

			_, err := reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())
			Expect(fakeClient.Get(ctx, namespacedName, storageNode)).Should(Succeed())
			Expect(storageNode.Status.Phase).To(Equal(v1alpha1.StorageNodePhaseDeleteComplete))
		})

		It("should be success when storage node is delete completed", func() {
			name := "test-aws-aurora-delete-completed"
			namespacedName := types.NamespacedName{
				Name:      name,
				Namespace: defaultTestNamespace,
			}
			req := ctrl.Request{NamespacedName: namespacedName}
			deletionTimestamp := metav1.Now()
			storageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: defaultTestNamespace,
					Annotations: map[string]string{
						v1alpha1.AnnotationsClusterIdentifier: "test-aws-aurora",
					},
					Finalizers:        []string{FinalizerName},
					DeletionTimestamp: &deletionTimestamp,
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: "aws-aurora",
				},
				Status: v1alpha1.StorageNodeStatus{
					Phase: v1alpha1.StorageNodePhaseDeleteComplete,
				},
			}
			Expect(fakeClient.Create(ctx, storageNode)).Should(Succeed())

			_, err := reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())
			err = fakeClient.Get(ctx, namespacedName, storageNode)
			Expect(apierrors.IsNotFound(err)).To(BeTrue())
		})

		It("should be success when storage node is ready for register", func() {
			name := "test-aws-aurora-ready-for-register"
			namespacedName := types.NamespacedName{
				Name:      name,
				Namespace: defaultTestNamespace,
			}
			req := ctrl.Request{NamespacedName: namespacedName}
			storageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: defaultTestNamespace,
					Annotations: map[string]string{
						v1alpha1.AnnotationsClusterIdentifier:   "test-aws-aurora",
						AnnotationKeyRegisterStorageUnitEnabled: "true",
						AnnotationKeyLogicDatabaseName:          "test-logic-db",
						v1alpha1.AnnotationsInstanceDBName:      "test-instance-db",
						AnnotationKeyComputeNodeName:            "test-compute-node",
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: "aws-aurora",
				},
				Status: v1alpha1.StorageNodeStatus{
					Phase: v1alpha1.StorageNodePhaseReady,
					Cluster: v1alpha1.ClusterStatus{
						Status:          dbmesh_rds.DBClusterStatusAvailable,
						PrimaryEndpoint: v1alpha1.Endpoint{Address: "test-aws-aurora.cluster-xxxxxx.us-east-1.rds.amazonaws.com", Port: 3306},
						ReaderEndpoints: []v1alpha1.Endpoint{{Address: "test-aws-aurora.cluster-ro-xxxxxx.us-east-1.rds.amazonaws.com", Port: 3306}},
					},
					Instances: []v1alpha1.InstanceStatus{
						{
							Status:   string(dbmesh_rds.DBInstanceStatusAvailable),
							Endpoint: v1alpha1.Endpoint{Address: "test-aws-aurora-1.cluster-xxxxxx.us-east-1.rds.amazonaws.com", Port: 3306},
						},
					},
				},
			}
			cn := &v1alpha1.ComputeNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-compute-node",
					Namespace: defaultTestNamespace,
				},
				Spec: v1alpha1.ComputeNodeSpec{
					Bootstrap: v1alpha1.BootstrapConfig{
						ServerConfig: v1alpha1.ServerConfig{
							Authority: v1alpha1.ComputeNodeAuthority{
								Users: []v1alpha1.ComputeNodeUser{{User: "test-user", Password: "test-password"}},
							},
							Mode:  v1alpha1.ComputeNodeServerMode{},
							Props: nil,
						},
					},
				},
			}
			svc := &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-compute-node",
					Namespace: defaultTestNamespace,
				},
				Spec: corev1.ServiceSpec{
					Ports: []corev1.ServicePort{
						{Name: "http", Protocol: "TCP", Port: 3307},
					},
				},
			}
			Expect(fakeClient.Create(ctx, storageNode)).Should(Succeed())
			Expect(fakeClient.Create(ctx, cn)).Should(Succeed())
			Expect(fakeClient.Create(ctx, svc)).Should(Succeed())

			// mock aws aurora is available
			mockAws.EXPECT().GetAuroraCluster(gomock.Any(), gomock.Any()).Return(&dbmesh_rds.DescCluster{
				DBClusterIdentifier: "test-aws-aurora",
				PrimaryEndpoint:     "test-aws-aurora.cluster-xxxxxx.us-east-1.rds.amazonaws.com",
				Port:                int32(3306),
				Status:              dbmesh_rds.DBClusterStatusAvailable,
			}, nil).Times(1)
			// mock get instances of aws aurora are available
			mockAws.EXPECT().GetInstancesByFilters(gomock.Any(), gomock.Any()).Return([]*dbmesh_rds.DescInstance{
				{
					DBInstanceIdentifier: "test-aws-aurora-instance-0",
					DBInstanceStatus:     dbmesh_rds.DBInstanceStatusAvailable,
					Endpoint:             dbmesh_rds.Endpoint{Address: "test-aws-aurora-1.cluster-xxxxxx.us-east-1.rds.amazonaws.com", Port: 3306},
				},
			}, nil).Times(1)

			host, port, username, password := getDatasourceInfoFromCluster(storageNode, provider)

			// mock shardingsphere
			mockSS.EXPECT().CreateDatabase(gomock.Any()).Return(nil).Times(1)
			mockSS.EXPECT().RegisterStorageUnit("test-logic-db", getDSName(storageNode), host, uint(port), "test-instance-db", username, password).Return(nil).Times(1)
			mockSS.EXPECT().Close().Return(nil).Times(1)

			_, err := reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())

			err = fakeClient.Get(ctx, namespacedName, storageNode)
			Expect(storageNode.Status.Registered).To(BeTrue())
		})

		It("should be success unregistered when storage node is deleting", func() {
			snName := "test-aws-aurora-unregistered"
			namespacedName := types.NamespacedName{
				Name:      snName,
				Namespace: defaultTestNamespace,
			}
			req := ctrl.Request{NamespacedName: namespacedName}

			storageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      snName,
					Namespace: defaultTestNamespace,
					Annotations: map[string]string{
						v1alpha1.AnnotationsClusterIdentifier:   "test-aws-aurora",
						AnnotationKeyRegisterStorageUnitEnabled: "true",
						AnnotationKeyLogicDatabaseName:          "test-logic-db",
						v1alpha1.AnnotationsInstanceDBName:      "test-instance-db",
						AnnotationKeyComputeNodeName:            "test-compute-node",
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					StorageProviderName: "aws-aurora",
					Replicas:            2,
				},
				Status: v1alpha1.StorageNodeStatus{
					Phase:      v1alpha1.StorageNodePhaseReady,
					Registered: true,
				},
			}
			cn := &v1alpha1.ComputeNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-compute-node",
					Namespace: defaultTestNamespace,
				},
				Spec: v1alpha1.ComputeNodeSpec{
					Bootstrap: v1alpha1.BootstrapConfig{
						ServerConfig: v1alpha1.ServerConfig{
							Authority: v1alpha1.ComputeNodeAuthority{
								Users: []v1alpha1.ComputeNodeUser{{User: "test-user", Password: "test-password"}},
							},
						},
					},
				},
			}
			svc := &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-compute-node",
					Namespace: defaultTestNamespace,
				},
				Spec: corev1.ServiceSpec{
					Ports: []corev1.ServicePort{
						{Name: "http", Protocol: "TCP", Port: 3307},
					},
				},
			}
			Expect(fakeClient.Create(ctx, storageNode)).Should(Succeed())
			Expect(fakeClient.Create(ctx, cn)).Should(Succeed())
			Expect(fakeClient.Create(ctx, svc)).Should(Succeed())

			// mock aws aurora is available
			mockAws.EXPECT().GetAuroraCluster(gomock.Any(), gomock.Any()).Return(&dbmesh_rds.DescCluster{
				DBClusterIdentifier: "test-aws-aurora",
				PrimaryEndpoint:     "test-aws-aurora.cluster-xxxxxx.us-east-1.rds.amazonaws.com",
				Port:                int32(3306),
				Status:              dbmesh_rds.DBClusterStatusAvailable,
				DBClusterMembers: []dbmesh_rds.ClusterMember{
					{DBInstanceIdentifier: "test-aws-aurora-instance-0"},
					{DBInstanceIdentifier: "test-aws-aurora-instance-1"},
				},
			}, nil).Times(2)
			// mock get instances of aws aurora are available
			mockAws.EXPECT().GetInstancesByFilters(gomock.Any(), gomock.Any()).Return([]*dbmesh_rds.DescInstance{
				{
					DBInstanceIdentifier: "test-aws-aurora-instance-0",
					DBInstanceStatus:     dbmesh_rds.DBInstanceStatusAvailable,
					Endpoint:             dbmesh_rds.Endpoint{Address: "test-aws-aurora-1.cluster-xxxxxx.us-east-1.rds.amazonaws.com", Port: 3306},
				},
				{
					DBInstanceIdentifier: "test-aws-aurora-instance-1",
					DBInstanceStatus:     dbmesh_rds.DBInstanceStatusAvailable,
					Endpoint:             dbmesh_rds.Endpoint{Address: "test-aws-aurora-2.cluster-xxxxxx.us-east-1.rds.amazonaws.com", Port: 3306},
				},
			}, nil).Times(2)

			_, err := reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())

			Expect(fakeClient.Delete(ctx, storageNode)).Should(Succeed())
			// mock shardingsphere
			mockSS.EXPECT().UnRegisterStorageUnit("test-logic-db", getDSName(storageNode)).Return(nil).Times(1)
			mockSS.EXPECT().Close().Return(nil).Times(1)

			// mock delete aws aurora
			mockAws.EXPECT().DeleteAuroraCluster(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

			_, err = reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())

			err = fakeClient.Get(ctx, namespacedName, storageNode)
			Expect(storageNode.Status.Registered).To(BeFalse())
		})
	})

})
