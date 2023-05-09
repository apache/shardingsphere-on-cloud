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

	"bou.ke/monkey"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/controllers"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode/aws"
	mock_aws "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode/aws/mocks"
	dbmesh_aws "github.com/database-mesh/golang-sdk/aws"
	dbmesh_rds "github.com/database-mesh/golang-sdk/aws/client/rds"
	dbmeshv1alpha1 "github.com/database-mesh/golang-sdk/kubernetes/api/v1alpha1"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var _ = Describe("StorageNode Controller Mock Test", func() {
	var fakeClient client.Client
	var reconciler *controllers.StorageNodeReconciler
	BeforeEach(func() {
		logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

		scheme := runtime.NewScheme()
		Expect(dbmeshv1alpha1.AddToScheme(scheme)).To(Succeed())
		Expect(v1alpha1.AddToScheme(scheme)).To(Succeed())
		fakeClient = fake.NewClientBuilder().WithScheme(scheme).Build()

		eventBroadcaster := record.NewBroadcaster()
		recorder := eventBroadcaster.NewRecorder(
			scheme,
			corev1.EventSource{
				Component: "test-storage-node-controller",
			},
		)

		sess := dbmesh_aws.NewSessions().SetCredential("AwsRegion", "AwsAccessKeyID", "AwsSecretAccessKey").Build()
		reconciler = &controllers.StorageNodeReconciler{
			Client:   fakeClient,
			Log:      logf.Log,
			Recorder: recorder,
			AwsRDS:   dbmesh_rds.NewService(sess["AwsRegion"]),
		}
	})

	Context("create storage node", func() {
		It("should create storage node successfully", func() {
			storageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-storage-node",
					Namespace: "test-namespace",
				},
				Spec: v1alpha1.StorageNodeSpec{
					DatabaseClassName: "test-database-class",
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
					Name:      "test-storage-node",
					Namespace: "test-namespace",
				},
				Spec: v1alpha1.StorageNodeSpec{
					DatabaseClassName: "no-database",
				},
				Status: v1alpha1.StorageNodeStatus{},
			}
			Expect(fakeClient.Create(ctx, storageNode)).Should(Succeed())
			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      "test-storage-node",
					Namespace: "test-namespace",
				},
			}
			_, err := reconciler.Reconcile(ctx, req)
			Expect(client.IgnoreNotFound(err)).Should(Succeed())
			Expect(fakeClient.Delete(ctx, storageNode)).Should(Succeed())
		})
	})

	Context("reconcile storageNode with exist databaseClass", func() {
		var mockCtrl *gomock.Controller
		var mockAws *mock_aws.MockIRdsClient
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockAws = mock_aws.NewMockIRdsClient(mockCtrl)

			monkey.Patch(aws.NewRdsClient, func(rds dbmesh_rds.RDS) aws.IRdsClient {
				return mockAws
			})

			// create databaseClass
			dbClass := &dbmeshv1alpha1.DatabaseClass{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-database-class",
				},
				Spec: dbmeshv1alpha1.DatabaseClassSpec{
					Provisioner: dbmeshv1alpha1.ProvisionerAWSRDSInstance,
				},
			}
			Expect(fakeClient.Create(ctx, dbClass)).Should(Succeed())

			// create storageNode
			storageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-storage-node",
					Namespace: "test-namespace",
				},
				Spec: v1alpha1.StorageNodeSpec{
					DatabaseClassName: "test-database-class",
				},
			}

			Expect(fakeClient.Create(ctx, storageNode)).Should(Succeed())
		})

		AfterEach(func() {
			mockCtrl.Finish()
			monkey.UnpatchAll()
		})

		It("should reconcile successfully with Creating Instance", func() {
			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      "test-storage-node",
					Namespace: "test-namespace",
				},
			}

			rdsInstance := &dbmesh_rds.DescInstance{
				DBInstanceStatus: "creating",
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
			Expect(newSN.Status.Instances[0].Status).To(Equal("creating"))
		})

		It("should reconcile successfully with Available Instance", func() {
			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      "test-storage-node",
					Namespace: "test-namespace",
				},
			}

			rdsInstance := &dbmesh_rds.DescInstance{
				DBInstanceStatus: "available",
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
			Expect(fakeClient.Get(ctx, client.ObjectKey{Name: "test-storage-node", Namespace: "test-namespace"}, newSN)).Should(Succeed())

			Expect(newSN.Status.Phase).To(Equal(v1alpha1.StorageNodePhaseReady))
			Expect(newSN.Status.Instances).To(HaveLen(1))
			Expect(newSN.Status.Instances[0].Status).To(Equal("Ready"))
		})

		It("should reconcile successfully when storage node be deleted", func() {
			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      "test-storage-node",
					Namespace: "test-namespace",
				},
			}

			rdsInstance := &dbmesh_rds.DescInstance{
				DBInstanceStatus: "available",
				Endpoint: dbmesh_rds.Endpoint{
					Address: "127.0.0.1",
					Port:    3306,
				},
			}

			// mock aws rds client, get instance
			mockAws.EXPECT().GetInstance(gomock.Any(), gomock.Any()).Return(rdsInstance, nil).AnyTimes()
			// reconcile storage node, add instance and set status to ready
			_, err := reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())

			// delete storage node
			sn := &v1alpha1.StorageNode{}
			Expect(fakeClient.Get(ctx, client.ObjectKey{Name: "test-storage-node", Namespace: "test-namespace"}, sn)).Should(Succeed())
			Expect(fakeClient.Delete(ctx, sn)).Should(Succeed())

			// mock aws rds client, delete instance
			mockAws.EXPECT().DeleteInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			_, err = reconciler.Reconcile(ctx, req)
			Expect(err).To(BeNil())
		})
	})
})

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
		BeforeEach(func() {

			// mock get instance func returns Available status
			monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "GetInstance", func(_ *aws.RdsClient, _ context.Context, _ *v1alpha1.StorageNode) (*dbmesh_rds.DescInstance, error) {
				return &dbmesh_rds.DescInstance{
					DBInstanceStatus: "available",
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
		})

		AfterEach(func() {
			monkey.UnpatchAll()
		})

		It("should be success", func() {

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

		Context("reconcile storageNode with Creating instance", func() {
			BeforeEach(func() {
				// mock get instance func returns creating status
				monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "GetInstance", func(_ *aws.RdsClient, _ context.Context, _ *v1alpha1.StorageNode) (*dbmesh_rds.DescInstance, error) {
					return &dbmesh_rds.DescInstance{
						DBInstanceStatus: "creating",
						Endpoint: dbmesh_rds.Endpoint{
							Address: "127.0.0.1",
							Port:    3306,
						},
					}, nil
				})
				// mock delete instance func return success
				monkey.PatchInstanceMethod(reflect.TypeOf(&aws.RdsClient{}), "DeleteInstance", func(_ *aws.RdsClient, _ context.Context, _ *v1alpha1.StorageNode, _ *dbmeshv1alpha1.DatabaseClass) error {
					return nil
				})
			})

			AfterEach(func() {
				monkey.UnpatchAll()
			})

			It("should be success", func() {
				nodeName := "test-storage-node-creating"
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
				}, 10*time.Second, 1*time.Second).Should(Equal(v1alpha1.StorageNodePhaseNotReady))

				// delete resource
				Expect(k8sClient.Delete(ctx, node)).Should(Succeed())
			})
		})
	})
})
