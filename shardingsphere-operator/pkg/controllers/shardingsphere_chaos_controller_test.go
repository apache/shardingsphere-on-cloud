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
	"database/sql"
	"regexp"
	"time"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	mockChaos "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaosmesh/mocks"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/pressure"
	reconcile "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/shardingspherechaos"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

func mockchaosStub(chaos *mockChaos.MockChaos) {
	chaos.EXPECT().NewNetworkChaos(gomock.Any(), gomock.Any()).Return(gomock.Any()).AnyTimes()
	chaos.EXPECT().NewPodChaos(gomock.Any(), gomock.Any()).Return(gomock.Any()).AnyTimes()

	chaos.EXPECT().CreatePodChaos(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	chaos.EXPECT().CreateNetworkChaos(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	chaos.EXPECT().DeleteNetworkChaos(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	chaos.EXPECT().DeletePodChaos(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	chaos.EXPECT().UpdatePodChaos(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	chaos.EXPECT().UpdateNetworkChaos(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	chaos.EXPECT().GetNetworkChaosByNamespacedName(gomock.Any(), gomock.Any()).Return(gomock.Any(), nil).AnyTimes()
	chaos.EXPECT().GetPodChaosByNamespacedName(gomock.Any(), gomock.Any()).Return(gomock.Any(), nil).AnyTimes()
}

func mockDBStub(mock sqlmock.Sqlmock) {
	mock.ExpectExec(regexp.QuoteMeta("REGISTER STORAGE UNIT")).WillReturnResult(sqlmock.NewResult(1, 1))
}

var _ = Describe("shardingsphere mock test", func() {
	var (
		testNamespacedName = types.NamespacedName{
			Namespace: "test-ssChaos-namespace",
			Name:      "test-ssChaos-name",
		}
		duration = "30s"
	)
	var (
		ctx        = context.TODO()
		fakeClient client.Client
		reconciler *ShardingSphereChaosReconciler
		mockCtrl   *gomock.Controller
		mockchaos  *mockChaos.MockChaos
		db         *sql.DB
	)

	BeforeEach(func() {
		scheme := runtime.NewScheme()
		Expect(clientgoscheme.AddToScheme(scheme)).To(Succeed())
		Expect(v1alpha1.AddToScheme(scheme)).To(Succeed())
		fakeClient = fake.NewClientBuilder().WithScheme(scheme).Build()
		mockCtrl = gomock.NewController(GinkgoT())

		mockchaos = mockChaos.NewMockChaos(mockCtrl)
		mockchaosStub(mockchaos)

		reconciler = &ShardingSphereChaosReconciler{
			Client:    fakeClient,
			Scheme:    scheme,
			Log:       logf.Log,
			Events:    record.NewFakeRecorder(100),
			Chaos:     mockchaos,
			ExecCtrls: make([]*ExecCtrl, 0),
			ConfigMap: configmap.NewConfigMapClient(fakeClient),
		}

		var (
			dbmock sqlmock.Sqlmock
			err    error
		)

		db, dbmock, err = sqlmock.New()
		Expect(err).To(BeNil())
		Expect(db).NotTo(BeNil())
		Expect(dbmock).NotTo(BeNil())

		monkey.Patch(sql.Open, func(driverName, dataSourceName string) (*sql.DB, error) {
			return db, nil
		})

		mockDBStub(dbmock)
	})

	AfterEach(func() {
		monkey.UnpatchAll()
		db.Close()
	})

	Context("create shardingsphere chaos", func() {
		It("should create successfully", func() {
			ssChaos := &v1alpha1.ShardingSphereChaos{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testNamespacedName.Name,
					Namespace: testNamespacedName.Namespace,
				},
				Spec: v1alpha1.ShardingSphereChaosSpec{
					EmbedChaos: v1alpha1.EmbedChaos{
						PodChaos: &v1alpha1.PodChaosSpec{
							PodSelector: v1alpha1.PodSelector{
								LabelSelectors: map[string]string{
									"app.kubernetes.io/component": "zookeeper",
								},
							},
							Action: v1alpha1.PodFailure,
							Params: v1alpha1.PodChaosParams{
								PodFailure: &v1alpha1.PodFailureParams{
									Duration: &duration,
								},
							},
						},
					},
					PressureCfg: v1alpha1.PressureCfg{
						SsHost:   "127.0.0.1:3306/ds_1",
						Duration: metav1.Duration{Duration: 30 * time.Second},
						ReqTime:  metav1.Duration{Duration: 30 * time.Second},
						DistSQLs: []v1alpha1.DistSQL{
							{
								SQL:  "REGISTER STORAGE UNIT ?()",
								Args: []string{"ds_1"},
							},
						},
						ConcurrentNum: 2,
						ReqNum:        5,
					},
				},
				Status: v1alpha1.ShardingSphereChaosStatus{},
			}

			Expect(fakeClient.Create(ctx, ssChaos)).Should(Succeed())
			chaos := &v1alpha1.ShardingSphereChaos{}
			_, err := reconciler.Reconcile(ctx, ctrl.Request{NamespacedName: testNamespacedName})
			Expect(err).To(BeNil())
			Expect(fakeClient.Get(ctx, testNamespacedName, chaos)).Should(Succeed())
			Expect(fakeClient.Delete(ctx, chaos)).Should(Succeed())
		})
	})

	Context("reconcile ssChaos in BeforeSteady phase", func() {
		It("chaos should be nil,execRecorder should be steady", func() {
			ssChaos := &v1alpha1.ShardingSphereChaos{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testNamespacedName.Name,
					Namespace: testNamespacedName.Namespace,
				},
				Spec: v1alpha1.ShardingSphereChaosSpec{
					EmbedChaos: v1alpha1.EmbedChaos{
						PodChaos: &v1alpha1.PodChaosSpec{
							PodSelector: v1alpha1.PodSelector{
								LabelSelectors: map[string]string{
									"app.kubernetes.io/component": "zookeeper",
								},
							},
							Action: v1alpha1.PodFailure,
							Params: v1alpha1.PodChaosParams{
								PodFailure: &v1alpha1.PodFailureParams{
									Duration: &duration,
								},
							},
						},
					},
					PressureCfg: v1alpha1.PressureCfg{
						SsHost:   "127.0.0.1:3306/ds_1",
						Duration: metav1.Duration{Duration: 30 * time.Second},
						ReqTime:  metav1.Duration{Duration: 30 * time.Second},
						DistSQLs: []v1alpha1.DistSQL{
							{
								SQL:  "REGISTER STORAGE UNIT ?()",
								Args: []string{"ds_1"},
							},
						},
						ConcurrentNum: 2,
						ReqNum:        5,
					},
				},
				Status: v1alpha1.ShardingSphereChaosStatus{},
			}

			Expect(fakeClient.Create(ctx, ssChaos)).Should(Succeed())
			for i := 0; i < 5; i++ {
				_, err := reconciler.Reconcile(ctx, ctrl.Request{NamespacedName: testNamespacedName})
				Expect(err).To(BeNil())
			}

			var inSteadyChaos v1alpha1.ShardingSphereChaos
			Expect(fakeClient.Get(ctx, testNamespacedName, &inSteadyChaos)).Should(Succeed())
			Expect(inSteadyChaos.Status.Phase).To(Equal(v1alpha1.BeforeSteady))

			Expect(len(reconciler.ExecCtrls)).To(Equal(1))
			Expect(fakeClient.Delete(ctx, &inSteadyChaos)).Should(Succeed())
		})
	})

	Context("reconcile ssChaos in BeforeChaos", func() {
		It("phase should in beforeChaos,execRecorder should gt 2", func() {
			ssChaos := &v1alpha1.ShardingSphereChaos{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testNamespacedName.Name,
					Namespace: testNamespacedName.Namespace,
				},
				Spec: v1alpha1.ShardingSphereChaosSpec{
					EmbedChaos: v1alpha1.EmbedChaos{
						PodChaos: &v1alpha1.PodChaosSpec{
							PodSelector: v1alpha1.PodSelector{
								LabelSelectors: map[string]string{
									"app.kubernetes.io/component": "zookeeper",
								},
							},
							Action: v1alpha1.PodFailure,
							Params: v1alpha1.PodChaosParams{
								PodFailure: &v1alpha1.PodFailureParams{
									Duration: &duration,
								},
							},
						},
					},
					PressureCfg: v1alpha1.PressureCfg{
						SsHost:   "127.0.0.1:3306/ds_1",
						Duration: metav1.Duration{Duration: 30 * time.Second},
						ReqTime:  metav1.Duration{Duration: 30 * time.Second},
						DistSQLs: []v1alpha1.DistSQL{
							{
								SQL:  "REGISTER STORAGE UNIT ?()",
								Args: []string{"ds_1"},
							},
						},
						ConcurrentNum: 2,
						ReqNum:        5,
					},
				},
				Status: v1alpha1.ShardingSphereChaosStatus{
					Phase: v1alpha1.BeforeChaos,
				},
			}

			Expect(fakeClient.Create(ctx, ssChaos)).Should(Succeed())
			var chao v1alpha1.ShardingSphereChaos
			Expect(fakeClient.Get(ctx, testNamespacedName, &chao)).Should(Succeed())
			steadyExec := pressure.NewPressure(reconcile.MakeJobName(ssChaos.Name, reconcile.InSteady), ssChaos.Spec.PressureCfg.DistSQLs)
			steadyExec.Active = false
			execCtx, cancel := context.WithCancel(ctx)
			execCtrl := ExecCtrl{
				cancel:   cancel,
				pressure: steadyExec,
				ctx:      execCtx,
			}
			reconciler.ExecCtrls = append(reconciler.ExecCtrls, &execCtrl)

			for i := 0; i < 10; i++ {
				_, err := reconciler.Reconcile(ctx, ctrl.Request{NamespacedName: testNamespacedName})
				Expect(err).To(BeNil())
			}
			Expect(len(reconciler.ExecCtrls)).To(Equal(2))

			var inChaosChaos v1alpha1.ShardingSphereChaos
			Expect(fakeClient.Get(ctx, testNamespacedName, &inChaosChaos)).Should(Succeed())
			Expect(inChaosChaos.Status.Phase).To(Equal(v1alpha1.BeforeChaos))
		})
	})
})
