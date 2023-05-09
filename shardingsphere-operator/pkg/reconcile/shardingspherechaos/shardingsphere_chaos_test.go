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
	"fmt"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"math/rand"
)

var _ = Describe("ShardingSphereChaos", func() {
	var d = "5m"

	Context("check related resource created by ShardingSphereChaos Controller", func() {
		var (
			ssChaos   *v1alpha1.ShardingSphereChaos
			name      = fmt.Sprintf("%s-%d", "test.sschaos-", rand.Int31())
			namespace = "default"
			ctx       = context.Background()
		)
		BeforeEach(func() {
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
							Action: v1alpha1.PodFailure,
							Params: v1alpha1.PodChaosParams{
								PodFailure: &v1alpha1.PodFailureParams{
									Duration: &d,
								},
							},
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, ssChaos)).To(BeNil())
		})

		AfterEach(func() {
			Expect(k8sClient.Delete(ctx, ssChaos)).To(BeNil())
		})

		//It("should create configmap", func() {
		//	configmap := &corev1.ConfigMap{}
		//	namespacedName := types.NamespacedName{Name: name, Namespace: namespace}
		//	Eventually(func() bool {
		//		err := k8sClient.Get(ctx, namespacedName, configmap)
		//		return err == nil
		//	}, time.Second*10, time.Millisecond*250).Should(BeTrue())
		//})

	})

})
