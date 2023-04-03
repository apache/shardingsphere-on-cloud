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

package ShardingSphereChaos_test

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
					ChaosKind: v1alpha1.PodChaosKind,
					EmbedChaos: v1alpha1.EmbedChaos{
						PodChaos: &v1alpha1.PodChaosSpec{
							PodSelector: v1alpha1.PodSelector{
								Namespaces: []string{"mesh-test"},
								LabelSelectors: map[string]string{
									"app.kubernetes.io/component": "zookeeper-new",
								},
							},
							Action: v1alpha1.PodFailureAction,
							PodActionParam: &v1alpha1.PodActionParam{
								PodFailure: &v1alpha1.PodFailureActionParams{
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

//todo: need fix schema
//import (
//	"fmt"
//	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
//	chaosV1AlphaV1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
//	. "github.com/onsi/ginkgo/v2"
//	. "github.com/onsi/gomega"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/apimachinery/pkg/types"
//	"time"
//)
//
//var _ = Describe("ShardingSphereChaos", func() {
//	Context("check related resource created by ShardingSphereChaos Controller", func() {
//		var ssChaos *v1alpha1.ShardingSphereChaos
//		BeforeEach(func() {
//			ssChaos = &v1alpha1.ShardingSphereChaos{
//				ObjectMeta: metav1.ObjectMeta{
//					Name:      "test.sschaos",
//					Namespace: "default",
//					Labels: map[string]string{
//						"app": "shardingsphere-proxy",
//					},
//					Annotations: map[string]string{
//						"spec/mode": "all",
//					},
//				},
//				Spec: v1alpha1.ShardingSphereChaosSpec{
//					ChaosKind: v1alpha1.PodChaosKind,
//					EmbedChaos: v1alpha1.EmbedChaos{
//						PodChaos: &v1alpha1.PodChaosSpec{
//							PodSelector: v1alpha1.PodSelector{
//								Namespaces: []string{"mesh-test"},
//								LabelSelectors: map[string]string{
//									"app.kubernetes.io/component": "zookeeper-new",
//								},
//							},
//							Action: v1alpha1.PodFailureAction,
//							PodActionParam: &v1alpha1.PodActionParam{
//								PodFailure: &v1alpha1.PodFailureActionParams{
//									Duration: "5m",
//								},
//							},
//						},
//					},
//				},
//			}
//			Expect(k8sClient.Create(ctx, ssChaos)).To(BeNil())
//		})
//
//		AfterEach(func() {
//			Expect(k8sClient.Delete(ctx, ssChaos)).To(BeNil())
//		})
//
//		It("should create podChaos", func() {
//			var podChaos chaosV1AlphaV1.PodChaos
//			nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
//			Eventually(func() bool {
//				err = k8sClient.Get(ctx, nameSpacedName, &podChaos)
//				fmt.Println(err)
//				return err == nil
//			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
//		})
//
//	})
//
//	//Context("verify metadata", func() {
//	//	var ssChaos *v1alpha1.ShardingSphereChaos
//	//	BeforeEach(func() {
//	//		ssChaos = &v1alpha1.ShardingSphereChaos{
//	//			ObjectMeta: metav1.ObjectMeta{
//	//				Name:      "test-ssChaos",
//	//				Namespace: "default",
//	//				Labels: map[string]string{
//	//					"app": "shardingsphere-proxy",
//	//				},
//	//				Annotations: map[string]string{
//	//					"spec/mode": "all",
//	//				},
//	//			},
//	//			Spec: v1alpha1.ShardingSphereChaosSpec{
//	//				ChaosKind: v1alpha1.PodChaosKind,
//	//				EmbedChaos: v1alpha1.EmbedChaos{
//	//					PodChaos: &v1alpha1.PodChaosSpec{
//	//						PodSelector: v1alpha1.PodSelector{
//	//							Namespaces: []string{"mesh-test"},
//	//							LabelSelectors: map[string]string{
//	//								"app.kubernetes.io/component": "zookeeper-new",
//	//							},
//	//						},
//	//						Action: v1alpha1.PodFailureAction,
//	//						PodActionParam: &v1alpha1.PodActionParam{
//	//							PodFailure: &v1alpha1.PodFailureActionParams{
//	//								Duration: "5m",
//	//							},
//	//						},
//	//					},
//	//				},
//	//			},
//	//		}
//	//		Expect(k8sClient.Create(ctx, ssChaos)).To(BeNil())
//	//
//	//	})
//	//
//	//	AfterEach(func() {
//	//		Expect(k8sClient.Delete(ctx, ssChaos)).To(BeNil())
//	//	})
//	//
//	//	if ssChaos.Spec.ChaosKind == v1alpha1.PodChaosKind {
//	//		It("verify podChaos metadata", func() {
//	//			var podChaos *chaosV1AlphaV1.PodChaos
//	//			nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
//	//			Eventually(func() bool {
//	//				return k8sClient.Get(ctx, nameSpacedName, podChaos) == nil
//	//			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
//	//			Expect(podChaos.Name).To(Equal(ssChaos.Name))
//	//			Expect(podChaos.Namespace).To(Equal(ssChaos.Namespace))
//	//			Expect(podChaos.Labels).To(Equal(ssChaos.Labels))
//	//		})
//	//	}
//	//
//	//	if ssChaos.Spec.ChaosKind == v1alpha1.NetworkChaosKind {
//	//		It("verify networkChaos metadata", func() {
//	//			var networkChaos *chaosV1AlphaV1.NetworkChaos
//	//			nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
//	//			Eventually(func() bool {
//	//				return k8sClient.Get(ctx, nameSpacedName, networkChaos) == nil
//	//			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
//	//			Expect(networkChaos.Name).To(Equal(ssChaos.Name))
//	//			Expect(networkChaos.Namespace).To(Equal(ssChaos.Namespace))
//	//			Expect(networkChaos.Labels).To(Equal(ssChaos.Labels))
//	//		})
//	//	}
//	//
//	//	//todo: add injectReq test here
//	//})
//	//
//	////todo: add more chaos action tests
//	//
//	//Context("verify PodChaos spec", func() {
//	//	var ssChaos *v1alpha1.ShardingSphereChaos
//	//
//	//	BeforeEach(func() {
//	//		ssChaos = &v1alpha1.ShardingSphereChaos{
//	//			ObjectMeta: metav1.ObjectMeta{
//	//				Name:      "test-ssChaos",
//	//				Namespace: "default",
//	//				Labels: map[string]string{
//	//					"app": "shardingsphere-proxy",
//	//				},
//	//				Annotations: map[string]string{
//	//					"spec/mode": "all",
//	//				},
//	//			},
//	//			Spec: v1alpha1.ShardingSphereChaosSpec{
//	//				ChaosKind: v1alpha1.PodChaosKind,
//	//				EmbedChaos: v1alpha1.EmbedChaos{
//	//					PodChaos: &v1alpha1.PodChaosSpec{
//	//						PodSelector: v1alpha1.PodSelector{
//	//							Namespaces: []string{"mesh-test"},
//	//							LabelSelectors: map[string]string{
//	//								"app.kubernetes.io/component": "zookeeper-new",
//	//							},
//	//						},
//	//						Action: v1alpha1.PodFailureAction,
//	//						PodActionParam: &v1alpha1.PodActionParam{
//	//							PodFailure: &v1alpha1.PodFailureActionParams{
//	//								Duration: "5m",
//	//							},
//	//						},
//	//					},
//	//				},
//	//			},
//	//		}
//	//		Expect(k8sClient.Create(ctx, ssChaos)).To(BeNil())
//	//	})
//	//
//	//	AfterEach(func() {
//	//		Expect(k8sClient.Delete(ctx, ssChaos)).To(BeNil())
//	//	})
//	//
//	//	It("verify podChaos Spec", func() {
//	//		var podChaos *chaosV1AlphaV1.PodChaos
//	//		var ssChaosPodSelector *v1alpha1.PodSelector
//	//		nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
//	//		Eventually(func() bool {
//	//			return k8sClient.Get(ctx, nameSpacedName, podChaos) == nil
//	//		}, time.Second*10, time.Millisecond*250).Should(BeTrue())
//	//		ssChaosPodSelector = &ssChaos.Spec.PodChaos.PodSelector
//	//		Expect(podChaos.Spec.PodSelector.Mode).To(Equal(chaosV1AlphaV1.SelectorMode(ssChaos.Annotations["spec/mode"])))
//	//		Expect(podChaos.Spec.PodSelector.Selector.Pods).To(Equal(ssChaosPodSelector.Pods))
//	//		Expect(podChaos.Spec.PodSelector.Selector.LabelSelectors).To(Equal(ssChaosPodSelector.LabelSelectors))
//	//		Expect(podChaos.Spec.Action).To(Equal(ssChaos.Spec.PodChaos.Action))
//	//	})
//	//})
//	//
//	//Context("verify networkChaos spec", func() {
//	//	var ssChaos *v1alpha1.ShardingSphereChaos
//	//	var duration = "1m"
//	//	BeforeEach(func() {
//	//		ssChaos = &v1alpha1.ShardingSphereChaos{
//	//			ObjectMeta: metav1.ObjectMeta{
//	//				Name:      "test-ssChaos",
//	//				Namespace: "default",
//	//				Labels: map[string]string{
//	//					"app": "shardingsphere-proxy",
//	//				},
//	//				Annotations: map[string]string{
//	//					"spec/mode":        "all",
//	//					"spec/target/mode": "all",
//	//				},
//	//			},
//	//			Spec: v1alpha1.ShardingSphereChaosSpec{
//	//				ChaosKind: v1alpha1.NetworkChaosKind,
//	//				EmbedChaos: v1alpha1.EmbedChaos{
//	//					NetworkChaos: &v1alpha1.NetworkChaosSpec{
//	//						Source: v1alpha1.PodSelector{
//	//							Namespaces: []string{"mesh-test"},
//	//							LabelSelectors: map[string]string{
//	//								"app": "shardingsphere-proxy-apache-shardingsphere-proxy",
//	//							},
//	//						},
//	//						Action:    v1alpha1.PartitionAction,
//	//						Duration:  &duration,
//	//						Direction: v1alpha1.To,
//	//						Target: &v1alpha1.PodSelector{
//	//							Namespaces: []string{"mesh-test"},
//	//							LabelSelectors: map[string]string{
//	//								"app.kubernetes.io/name": "zookeeper",
//	//							},
//	//						},
//	//					},
//	//				},
//	//			},
//	//		}
//	//		Expect(k8sClient.Create(ctx, ssChaos)).To(BeNil())
//	//	})
//	//
//	//	AfterEach(func() {
//	//		Expect(k8sClient.Delete(ctx, ssChaos)).To(BeNil())
//	//	})
//	//
//	//	It("verify netWorkChaos Spec", func() {
//	//		var netWorkChaos *chaosV1AlphaV1.NetworkChaos
//	//		var ssChaosPodSelector *v1alpha1.PodSelector
//	//		nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
//	//		Eventually(func() bool {
//	//			return k8sClient.Get(ctx, nameSpacedName, netWorkChaos) == nil
//	//		}, time.Second*10, time.Millisecond*250).Should(BeTrue())
//	//		ssChaosPodSelector = &ssChaos.Spec.PodChaos.PodSelector
//	//		Expect(netWorkChaos.Spec.PodSelector.Mode).To(Equal(chaosV1AlphaV1.SelectorMode(ssChaos.Annotations["spec/mode"])))
//	//		Expect(netWorkChaos.Spec.Target.Mode).To(Equal(chaosV1AlphaV1.SelectorMode(ssChaos.Annotations["spec/target/mode"])))
//	//		Expect(netWorkChaos.Spec.PodSelector.Selector.Namespaces).To(Equal(ssChaosPodSelector.Namespaces))
//	//		Expect(netWorkChaos.Spec.Target.Selector.Namespaces).To(Equal(ssChaos.Spec.NetworkChaos.Target.Namespaces))
//	//		Expect(netWorkChaos.Spec.PodSelector.Selector.LabelSelectors).To(Equal(ssChaos.Spec.NetworkChaos.Source.LabelSelectors))
//	//		Expect(netWorkChaos.Spec.Target.Selector.LabelSelectors).To(Equal(ssChaos.Spec.NetworkChaos.Target.LabelSelectors))
//	//		Expect(netWorkChaos.Spec.Action).To(Equal(chaosV1AlphaV1.NetworkChaosAction(ssChaos.Spec.NetworkChaos.Action)))
//	//		Expect(netWorkChaos.Spec.PodSelector.Value).To(Equal(ssChaos.Annotations["spec/value"]))
//	//		Expect(netWorkChaos.Spec.Target.Value).To(Equal(ssChaos.Annotations["spec/target/value"]))
//	//		Expect(*netWorkChaos.Spec.Duration).To(Equal(*ssChaos.Spec.NetworkChaos.Duration))
//	//	})
//	//})
//})
