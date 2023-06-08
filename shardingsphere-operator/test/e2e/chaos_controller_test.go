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
	"database/sql"
	"regexp"

	mockChaos "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaosmesh/mocks"
	"github.com/golang/mock/gomock"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

type compare struct {
	phase           v1alpha1.ChaosPhase
	conditionStatus []metav1.ConditionStatus
}

var _ = Describe("ShardingSphereChaos", func() {
	var (
		/*
			testNamespacedName = types.NamespacedName{
				Namespace: "default",
				Name:      "testsschaos",
			}
			duration = "5s"
		*/
		db *sql.DB
	)

	BeforeEach(func() {
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
		mockchaosStub(mockchaos)
		mockDBStub(dbmock)
	})

	AfterEach(func() {
		monkey.UnpatchAll()
		db.Close()
	})

	Context("reconcile Chaos", func() {
		/*
			var desireStatus = compare{
				phase:           v1alpha1.AfterChaos,
				conditionStatus: []metav1.ConditionStatus{metav1.ConditionTrue, metav1.ConditionTrue, metav1.ConditionTrue, metav1.ConditionTrue},
			}
				It("should create successfully", func() {
					ssChaos := &v1alpha1.Chaos{
						ObjectMeta: metav1.ObjectMeta{
							Name:      testNamespacedName.Name,
							Namespace: testNamespacedName.Namespace,
						},
						Spec: v1alpha1.ChaosSpec{
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
								Duration: metav1.Duration{Duration: 10 * time.Second},
								ReqTime:  metav1.Duration{Duration: 5 * time.Second},
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
						Status: v1alpha1.ChaosStatus{},
					}

					Expect(k8sClient.Create(ctx, ssChaos)).Should(Succeed())

					Eventually(func() compare {
						var chaos v1alpha1.ShardingSphereChaos
						Expect(k8sClient.Get(ctx, testNamespacedName, &chaos)).Should(Succeed())
						now := compare{
							phase:           chaos.Status.Phase,
							conditionStatus: make([]metav1.ConditionStatus, 0),
						}
						for i := range chaos.Status.Conditions {
							now.conditionStatus = append(now.conditionStatus, chaos.Status.Conditions[i].Status)
						}

						return now
					}, 25*time.Second, 1*time.Second).Should(Equal(desireStatus))

					Expect(k8sClient.Delete(ctx, ssChaos)).Should(Succeed())
				})
		*/

	})

})
