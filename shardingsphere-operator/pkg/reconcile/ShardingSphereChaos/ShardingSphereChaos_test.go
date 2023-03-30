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
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	ss "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/ShardingSphereChaos"
	chaosV1AlphaV1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"time"
)

var _ = Describe("ShardingSphereChaos", func() {
	Context("check related resource created by ShardingSphereChaos Controller", func() {
		var ssChaos *v1alpha1.ShardingSphereChaos
		BeforeEach(func() {
			ssChaos = &v1alpha1.ShardingSphereChaos{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ssChaos",
					Namespace: "default",
					Labels: map[string]string{
						"app": "shardingsphere-proxy",
					},
				},
				Spec: v1alpha1.ShardingSphereChaosSpec{
					ChaosKind: v1alpha1.PodChaosKind,
					EmbedChaos: v1alpha1.EmbedChaos{
						PodChaos: &v1alpha1.PodChaosSpec{
							PodSelector: v1alpha1.PodSelector{
								Mode: v1alpha1.FixedMode,
								Selector: v1alpha1.PodSelectorSpec{
									GenericSelectorSpec: v1alpha1.GenericSelectorSpec{
										Namespaces: []string{"mesh-test"},
										LabelSelectors: map[string]string{
											"app.kubernetes.io/component": "zookeeper-new",
										},
									},
									Nodes: nil,
									Pods: map[string][]string{
										"mesh-test": {"zookeeper-new-2"},
									},
									NodeSelectors:     nil,
									PodPhaseSelectors: nil,
								},
							},
							Action: v1alpha1.PodKillAction,
						},
					},
				},
			}

			Expect(k8sClient.Create(ctx, ssChaos)).To(BeNil())
		})

		AfterEach(Expect(k8sClient.Delete(ctx, ssChaos)).To(BeNil()))

		if ssChaos.Spec.ChaosKind == v1alpha1.PodChaosKind {
			It("should create podChaos", func() {
				var podChaos *chaosV1AlphaV1.PodChaos
				nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
				Eventually(func() bool {
					return k8sClient.Get(ctx, nameSpacedName, podChaos) == nil
				}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			})
		}

		if ssChaos.Spec.ChaosKind == v1alpha1.NetworkChaosKind {
			It("should create networkChaos", func() {
				var networkChaos *chaosV1AlphaV1.NetworkChaos
				nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
				Eventually(func() bool {
					return k8sClient.Get(ctx, nameSpacedName, networkChaos) == nil
				}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			})
		}

		if ssChaos.Spec.ChaosKind == v1alpha1.WorkFlowKind {
			It("should create workflow", func() {
				var workflow *chaosV1AlphaV1.Workflow
				nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
				Eventually(func() bool {
					return k8sClient.Get(ctx, nameSpacedName, workflow) == nil
				}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			})
		}
		//todo: add injectReq test here
	})

	Context("verify metadata", func() {
		var ssChaos *v1alpha1.ShardingSphereChaos
		BeforeEach(func() {
			ssChaos = &v1alpha1.ShardingSphereChaos{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ssChaos",
					Namespace: "default",
					Labels: map[string]string{
						"app": "shardingsphere-proxy",
					},
				},
				Spec: v1alpha1.ShardingSphereChaosSpec{
					ChaosKind: v1alpha1.PodChaosKind,
					EmbedChaos: v1alpha1.EmbedChaos{
						PodChaos: &v1alpha1.PodChaosSpec{
							PodSelector: v1alpha1.PodSelector{
								Mode: v1alpha1.FixedMode,
								Selector: v1alpha1.PodSelectorSpec{
									GenericSelectorSpec: v1alpha1.GenericSelectorSpec{
										Namespaces: []string{"mesh-test"},
										LabelSelectors: map[string]string{
											"app.kubernetes.io/component": "zookeeper-new",
										},
									},
									Nodes: nil,
									Pods: map[string][]string{
										"mesh-test": {"zookeeper-new-2"},
									},
									NodeSelectors:     nil,
									PodPhaseSelectors: nil,
								},
							},
							Action: v1alpha1.PodKillAction,
						},
					},
				},
			}

			Expect(k8sClient.Create(ctx, ssChaos)).To(BeNil())
		})

		AfterEach(Expect(k8sClient.Delete(ctx, ssChaos)).To(BeNil()))

		if ssChaos.Spec.ChaosKind == v1alpha1.PodChaosKind {
			It("verify podChaos metadata", func() {
				var podChaos *chaosV1AlphaV1.PodChaos
				nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
				Eventually(func() bool {
					return k8sClient.Get(ctx, nameSpacedName, podChaos) == nil
				}, time.Second*10, time.Millisecond*250).Should(BeTrue())
				Expect(podChaos.Name).To(Equal(ssChaos.Name))
				Expect(podChaos.Namespace).To(Equal(ssChaos.Namespace))
				Expect(podChaos.Labels).To(Equal(ssChaos.Labels))
			})
		}

		if ssChaos.Spec.ChaosKind == v1alpha1.NetworkChaosKind {
			It("verify networkChaos metadata", func() {
				var networkChaos *chaosV1AlphaV1.NetworkChaos
				nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
				Eventually(func() bool {
					return k8sClient.Get(ctx, nameSpacedName, networkChaos) == nil
				}, time.Second*10, time.Millisecond*250).Should(BeTrue())
				Expect(networkChaos.Name).To(Equal(ssChaos.Name))
				Expect(networkChaos.Namespace).To(Equal(ssChaos.Namespace))
				Expect(networkChaos.Labels).To(Equal(ssChaos.Labels))
			})
		}

		if ssChaos.Spec.ChaosKind == v1alpha1.WorkFlowKind {
			It("verify workflow metadata", func() {
				var workflow *chaosV1AlphaV1.Workflow
				nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
				Eventually(func() bool {
					return k8sClient.Get(ctx, nameSpacedName, workflow) == nil
				}, time.Second*10, time.Millisecond*250).Should(BeTrue())
				Expect(workflow.Name).To(Equal(ssChaos.Name))
				Expect(workflow.Namespace).To(Equal(ssChaos.Namespace))
				Expect(workflow.Labels).To(Equal(ssChaos.Labels))
			})
		}

		//todo: add injectReq test here
	})

	//todo: add more chaos action tests

	Context("verify PodChaos spec", func() {
		var ssChaos *v1alpha1.ShardingSphereChaos

		BeforeEach(func() {
			ssChaos = &v1alpha1.ShardingSphereChaos{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ssChaos",
					Namespace: "default",
					Labels: map[string]string{
						"app": "shardingsphere-proxy",
					},
				},
				Spec: v1alpha1.ShardingSphereChaosSpec{
					ChaosKind: v1alpha1.PodChaosKind,
					EmbedChaos: v1alpha1.EmbedChaos{
						PodChaos: &v1alpha1.PodChaosSpec{
							PodSelector: v1alpha1.PodSelector{
								Mode: v1alpha1.FixedMode,
								Selector: v1alpha1.PodSelectorSpec{
									GenericSelectorSpec: v1alpha1.GenericSelectorSpec{
										Namespaces: []string{"mesh-test"},
										LabelSelectors: map[string]string{
											"app.kubernetes.io/component": "zookeeper-new",
										},
									},
									Nodes: nil,
									Pods: map[string][]string{
										"mesh-test": {"zookeeper-new-2"},
									},
									NodeSelectors:     nil,
									PodPhaseSelectors: nil,
								},
							},
							Action: v1alpha1.PodKillAction,
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, ssChaos)).To(BeNil())
		})

		AfterEach(Expect(k8sClient.Delete(ctx, ssChaos)).To(BeNil()))

		It("verify podChaos Spec", func() {
			var podChaos *chaosV1AlphaV1.PodChaos
			var ssChaosPodSelector *v1alpha1.PodSelector
			nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
			Eventually(func() bool {
				return k8sClient.Get(ctx, nameSpacedName, podChaos) == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			ssChaosPodSelector = &ssChaos.Spec.PodChaos.PodSelector
			Expect(podChaos.Spec.PodSelector.Mode).To(Equal(chaosV1AlphaV1.SelectorMode(ssChaos.Spec.PodChaos.Mode)))
			Expect(podChaos.Spec.PodSelector).To(Equal(ss.DeepCopyPodSelector(ssChaosPodSelector)))
			Expect(podChaos.Spec.GracePeriod).To(Equal(ssChaos.Spec.PodChaos.GracePeriod))
		})
	})

	Context("verify networkChaos spec", func() {
		var ssChaos *v1alpha1.ShardingSphereChaos
		var duration string = "1m"
		BeforeEach(func() {
			ssChaos = &v1alpha1.ShardingSphereChaos{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ssChaos",
					Namespace: "default",
					Labels: map[string]string{
						"app": "shardingsphere-proxy",
					},
				},
				Spec: v1alpha1.ShardingSphereChaosSpec{
					ChaosKind: v1alpha1.NetworkChaosKind,
					EmbedChaos: v1alpha1.EmbedChaos{
						NetworkChaos: &v1alpha1.NetworkChaosSpec{
							PodSelector: v1alpha1.PodSelector{
								Selector: v1alpha1.PodSelectorSpec{
									GenericSelectorSpec: v1alpha1.GenericSelectorSpec{
										Namespaces: []string{"mesh-test"},
										LabelSelectors: map[string]string{
											"app": "shardingsphere-proxy-apache-shardingsphere-proxy",
										},
									},
								},
								Mode: v1alpha1.AllMode,
							},
							Action:      v1alpha1.PartitionAction,
							Device:      "",
							Duration:    &duration,
							TcParameter: v1alpha1.TcParameter{},
							Direction:   v1alpha1.To,
							Target: &v1alpha1.PodSelector{
								Selector: v1alpha1.PodSelectorSpec{
									GenericSelectorSpec: v1alpha1.GenericSelectorSpec{
										Namespaces: []string{"mesh-test"},
										LabelSelectors: map[string]string{
											"app.kubernetes.io/name": "zookeeper",
										},
									},
								},
								Mode: v1alpha1.AllMode,
							},
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, ssChaos)).To(BeNil())
		})

		AfterEach(Expect(k8sClient.Delete(ctx, ssChaos)).To(BeNil()))

		It("verify netWorkChaos Spec", func() {
			var netWorkChaos *chaosV1AlphaV1.NetworkChaos
			var ssChaosPodSelector *v1alpha1.PodSelector
			nameSpacedName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
			Eventually(func() bool {
				return k8sClient.Get(ctx, nameSpacedName, netWorkChaos) == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			ssChaosPodSelector = &ssChaos.Spec.PodChaos.PodSelector
			Expect(netWorkChaos.Spec.PodSelector.Mode).To(Equal(chaosV1AlphaV1.SelectorMode(ssChaos.Spec.PodChaos.Mode)))
			Expect(netWorkChaos.Spec.PodSelector).To(Equal(ss.DeepCopyPodSelector(ssChaosPodSelector)))
			Expect(netWorkChaos.Spec.Action).To(Equal(chaosV1AlphaV1.NetworkChaosAction(ssChaos.Spec.NetworkChaos.Action)))
			Expect(*netWorkChaos.Spec.Duration).To(Equal(*ssChaos.Spec.NetworkChaos.Duration))
			Expect(netWorkChaos.Spec.Direction).To(Equal(ssChaos.Spec.NetworkChaos.Direction))
		})
	})

})