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
	dbmeshv1alpha1 "github.com/database-mesh/golang-sdk/kubernetes/api/v1alpha1"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

const (
	defaultTestNamespace          = "test-namespace"
	defaultTestDBClass            = "test-database-class"
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
	Expect(dbmeshv1alpha1.AddToScheme(scheme)).To(Succeed())
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

var _ = Describe("StorageNode Controller Mock Test", func() {
	BeforeEach(func() {
		// mock aws rds client
		mockCtrl = gomock.NewController(GinkgoT())
		mockAws = mock_aws.NewMockIRdsClient(mockCtrl)

		monkey.Patch(aws.NewRdsClient, func(rds dbmesh_rds.RDS) aws.IRdsClient {
			return mockAws
		})

		// create default resource
		dbClass := &dbmeshv1alpha1.DatabaseClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: defaultTestDBClass,
			},
			Spec: dbmeshv1alpha1.DatabaseClassSpec{
				Provisioner: dbmeshv1alpha1.ProvisionerAWSRDSInstance,
			},
		}

		storageNode := &v1alpha1.StorageNode{
			ObjectMeta: metav1.ObjectMeta{
				Name:      defaultTestStorageNode,
				Namespace: defaultTestNamespace,
				Annotations: map[string]string{
					dbmeshv1alpha1.AnnotationsInstanceIdentifier: defaultTestInstanceIdentifier,
				},
			},
			Spec: v1alpha1.StorageNodeSpec{
				DatabaseClassName: defaultTestDBClass,
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
		Expect(fakeClient.Delete(ctx, &dbmeshv1alpha1.DatabaseClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: defaultTestDBClass,
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
					DatabaseClassName: defaultTestDBClass,
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
					DatabaseClassName: "no-database",
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
				DBInstanceStatus: v1alpha1.StorageNodeInstanceStatusCreating,
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
			Expect(newSN.Status.Instances[0].Status).To(Equal(v1alpha1.StorageNodeInstanceStatusCreating))
		})

		It("should reconcile successfully with Available Instance", func() {
			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      defaultTestStorageNode,
					Namespace: defaultTestNamespace,
				},
			}

			rdsInstance := &dbmesh_rds.DescInstance{
				DBInstanceStatus: v1alpha1.StorageNodeInstanceStatusAvailable,
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
			Expect(newSN.Status.Instances[0].Status).To(Equal(v1alpha1.StorageNodeInstanceStatusReady))
		})
	})

	Context("reconcile storage node in Ready status when it's been deleted", func() {
		var (
			rdsInstanceAvailable = dbmesh_rds.DescInstance{
				DBInstanceIdentifier: defaultTestInstanceIdentifier,
				DBInstanceStatus:     v1alpha1.StorageNodeInstanceStatusAvailable,
				Endpoint: dbmesh_rds.Endpoint{
					Address: "127.0.0.1",
					Port:    3306,
				},
			}
			instanceInDeleting = dbmesh_rds.DescInstance{
				DBInstanceIdentifier: defaultTestInstanceIdentifier,
				DBInstanceStatus:     v1alpha1.StorageNodeInstanceStatusDeleting,
				Endpoint: dbmesh_rds.Endpoint{
					Address: "127.0.0.1",
					Port:    3306,
				},
			}
		)
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
						dbmeshv1alpha1.AnnotationsInstanceIdentifier: defaultTestInstanceIdentifier,
					},
				},
				Spec: v1alpha1.StorageNodeSpec{DatabaseClassName: defaultTestDBClass},
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
			mockAws.EXPECT().GetInstance(gomock.Any(), gomock.Any()).Return(&instanceInDeleting, nil)
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
					Finalizers: []string{
						FinalizerName,
					},
					DeletionTimestamp: &deleteTime,
				},
				Spec: v1alpha1.StorageNodeSpec{
					DatabaseClassName: defaultTestDBClass,
				},
				Status: v1alpha1.StorageNodeStatus{
					Phase: v1alpha1.StorageNodePhaseDeleting,
					Instances: []v1alpha1.InstanceStatus{
						{
							Status: v1alpha1.StorageNodeInstanceStatusDeleting,
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
					DatabaseClassName: defaultTestDBClass,
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
						AnnotationKeyRegisterStorageUnitEnabled:  "true",
						dbmeshv1alpha1.AnnotationsInstanceDBName: "test_db",
						AnnotationKeyComputeNodeNamespace:        defaultTestNamespace,
						AnnotationKeyComputeNodeName:             cnName,
						AnnotationKeyLogicDatabaseName:           "sharding_db",
					},
				},
				Spec: v1alpha1.StorageNodeSpec{
					DatabaseClassName: defaultTestDBClass,
				},
			}
			ins := &dbmesh_rds.DescInstance{
				DBInstanceIdentifier: "ins-test-register-storage-node",
				DBInstanceStatus:     v1alpha1.StorageNodeInstanceStatusAvailable,
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
			mockSS.EXPECT().RegisterStorageUnit(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
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
						AnnotationKeyComputeNodeName:      "test-get-shardingsphere-server",
						AnnotationKeyComputeNodeNamespace: defaultTestNamespace,
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
						AnnotationKeyComputeNodeName:             testName,
						AnnotationKeyComputeNodeNamespace:        defaultTestNamespace,
						AnnotationKeyRegisterStorageUnitEnabled:  "true",
						AnnotationKeyLogicDatabaseName:           testName,
						dbmeshv1alpha1.AnnotationsInstanceDBName: testName,
					},
				},
				Status: v1alpha1.StorageNodeStatus{
					Phase: v1alpha1.StorageNodePhaseReady,
					Instances: []v1alpha1.InstanceStatus{
						{
							Status:   v1alpha1.StorageNodeInstanceStatusAvailable,
							Endpoint: v1alpha1.Endpoint{},
						},
					},
				},
			}

			mockSS.EXPECT().CreateDatabase(gomock.Any()).Return(nil)
			mockSS.EXPECT().Close().Return(nil)
			mockSS.EXPECT().RegisterStorageUnit(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

			Expect(reconciler.registerStorageUnit(ctx, sn)).To(BeNil())
			Expect(sn.Status.Registered).To(BeTrue())
		})
	})
})
