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

package controllers_test

import (
	"context"

	"bou.ke/monkey"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/controllers"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode/awsaurora"
	mock_storagenode "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/storagenode/mocks"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
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

var _ = Describe("StorageNode Controller", func() {
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

		reconciler = &controllers.StorageNodeReconciler{
			Client:   fakeClient,
			Log:      logf.Log,
			Recorder: recorder,
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

			Expect(fakeClient.Create(context.Background(), storageNode)).Should(Succeed())
			sn := &v1alpha1.StorageNode{}
			Expect(fakeClient.Get(context.Background(), client.ObjectKey{Name: "test-storage-node", Namespace: "test-namespace"}, sn)).Should(Succeed())
			Expect(fakeClient.Delete(context.Background(), storageNode)).Should(Succeed())
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
			Expect(fakeClient.Create(context.Background(), storageNode)).Should(Succeed())
			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      "test-storage-node",
					Namespace: "test-namespace",
				},
			}
			_, err := reconciler.Reconcile(context.Background(), req)
			Expect(client.IgnoreNotFound(err)).Should(Succeed())
			Expect(fakeClient.Delete(context.Background(), storageNode)).Should(Succeed())
		})
	})

	Context("create storage node with exist databaseClass", func() {
		// test create node with exist databaseClass and success
		It("should reconcile successfully", func() {
			var mockCtrl *gomock.Controller
			var cc *mock_storagenode.MockIDBClusterClient

			mockCtrl = gomock.NewController(GinkgoT())
			cc = mock_storagenode.NewMockIDBClusterClient(mockCtrl)

			monkey.Patch(awsaurora.New, func(_ rds.RDS) storagenode.IDBClusterClient {
				return cc
			})

			defer func() {
				mockCtrl.Finish()
			}()

			dbClass := &dbmeshv1alpha1.DatabaseClass{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-database-class",
				},
				Spec: dbmeshv1alpha1.DatabaseClassSpec{
					Provisioner: "aws-aurora",
				},
			}
			Expect(fakeClient.Create(context.Background(), dbClass)).Should(Succeed())

			storageNode := &v1alpha1.StorageNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-storage-node",
					Namespace: "test-namespace",
				},
				Spec: v1alpha1.StorageNodeSpec{
					DatabaseClassName: "test-database-class",
				},
			}
			Expect(fakeClient.Create(context.Background(), storageNode)).Should(Succeed())

			req := ctrl.Request{
				NamespacedName: client.ObjectKey{
					Name:      "test-storage-node",
					Namespace: "test-namespace",
				},
			}

			// cluster is not exist and create a new cluster
			cc.EXPECT().IsValid(gomock.Any()).Return(nil).AnyTimes()
			cc.EXPECT().GetCluster(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
			cc.EXPECT().CreateCluster(gomock.Any(), gomock.Any(), gomock.Any()).Return(&storagenode.DatabaseCluster{}, nil).AnyTimes()
			_, err := reconciler.Reconcile(context.Background(), req)
			Expect(err).To(BeNil())

			newSN := &v1alpha1.StorageNode{}
			Expect(fakeClient.Get(context.Background(), client.ObjectKey{Name: "test-storage-node", Namespace: "test-namespace"}, newSN)).Should(Succeed())
		})
	})
})
