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

package shardingspherechaos_test

import (
	"context"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/controllers"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaos"
	chaosV1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ = Describe("ShardingSphereChaos", func() {

	Context("check related resource created by ShardingSphereChaos Controller", func() {
		var (
			fakeClient client.Client
			reconciler *controllers.ShardingSphereChaosReconciler
			ssChaos    *v1alpha1.ShardingSphereChaos
			ctx        context.Context
			name       = "test.sschaos"
			namespace  = "default"
			req        = reconcile.Request{
				NamespacedName: client.ObjectKey{
					Namespace: name,
					Name:      namespace,
				},
			}
		)
		BeforeEach(func() {
			scheme := runtime.NewScheme()
			Expect(corev1.AddToScheme(scheme)).To(Succeed())
			Expect(appsv1.AddToScheme(scheme)).To(Succeed())
			Expect(batchv1.AddToScheme(scheme)).To(Succeed())
			Expect(chaosV1alpha1.AddToScheme(scheme)).To(Succeed())
			Expect(v1alpha1.AddToScheme(scheme)).To(Succeed())
			fakeClient = fake.NewClientBuilder().WithScheme(scheme).Build()
			logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))
			reconciler = &controllers.ShardingSphereChaosReconciler{
				Client: fakeClient,
				Scheme: scheme,
				Log:    logf.Log,
				Chaos:  chaos.NewChaos(fakeClient),
			}
			ctx = context.Background()
			ssChaos = &v1alpha1.ShardingSphereChaos{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: namespace,
					Labels: map[string]string{
						"app": "shardingsphere-proxy",
					},
					Annotations: map[string]string{
						"spec/mode": "all",
					},
				},
				Spec: v1alpha1.ShardingSphereChaosSpec{
					EmbedChaos: v1alpha1.EmbedChaos{
						PodChaos: &v1alpha1.PodChaosSpec{
							PodSelector: v1alpha1.PodSelector{
								Namespaces: []string{"mesh-test"},
								LabelSelectors: map[string]string{
									"app.kubernetes.io/component": "zookeeper-new",
								},
							},
							Action: v1alpha1.PodFailureAction,
							PodActionParam: v1alpha1.PodActionParam{
								PodFailure: v1alpha1.PodFailureActionParams{
									Duration: "5m",
								},
							},
						},
					},
				},
			}
			Expect(fakeClient.Create(ctx, ssChaos)).To(Succeed())
			_, err := reconciler.Reconcile(context.Background(), req)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			Expect(fakeClient.Delete(ctx, ssChaos)).To(Succeed())
			_, err := reconciler.Reconcile(context.Background(), req)
			Expect(err).NotTo(HaveOccurred())
		})

		//It("should create podChaos", func() {
		//	var podChaos chaosV1alpha1.PodChaos
		//	nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
		//	Eventually(func() bool {
		//		return Expect(fakeClient.Get(ctx, nameSpacedName, &podChaos)).To(Succeed())
		//	}, time.Second*10, time.Millisecond*250).Should(BeTrue())
		//})

	})

})
