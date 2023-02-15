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

package computenode_test

import (
	"time"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("ComputeNodeController", func() {
	Context("check related resource created by compute node controller", func() {
		var cn *v1alpha1.ComputeNode
		BeforeEach(func() {
			cn = &v1alpha1.ComputeNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-name",
					Namespace: "default",
					Labels: map[string]string{
						"app": "shardingsphere-proxy",
					},
				},
				Spec: v1alpha1.ComputeNodeSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app": "shardingsphere-proxy",
						},
					},
					Replicas: 2,
					Bootstrap: v1alpha1.BootstrapConfig{
						ServerConfig: v1alpha1.ServerConfig{
							Mode: v1alpha1.ComputeNodeServerMode{
								Repository: v1alpha1.Repository{
									Type: v1alpha1.RepositoryTypeZookeeper,
								},
							},
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
					StorageNodeConnector: &v1alpha1.StorageNodeConnector{
						Type:    v1alpha1.ConnectorTypeMySQL,
						Version: "5.1.47",
					},
					PortBindings: []v1alpha1.PortBinding{
						{
							Name:          "port",
							ContainerPort: 3307,
							ServicePort:   3307,
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, cn)).To(BeNil())
		})
		AfterEach(func() {
			Expect(k8sClient.Delete(ctx, cn)).To(BeNil())
		})

		// Integration tests using It blocks are written here.
		It("should create deployment", func() {
			createdDeploy := &appsv1.Deployment{}
			namespacedName := types.NamespacedName{Name: cn.Name, Namespace: cn.Namespace}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, createdDeploy)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
		})

		It("should create service", func() {
			createdService := &corev1.Service{}
			namespacedName := types.NamespacedName{Name: cn.Name, Namespace: cn.Namespace}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, createdService)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
		})

		It("should create configmap", func() {
			createdConfigmap := &corev1.ConfigMap{}
			namespacedName := types.NamespacedName{Name: cn.Name, Namespace: cn.Namespace}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, createdConfigmap)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
		})

	})

	Context("verify metadata", func() {
		var cn *v1alpha1.ComputeNode
		BeforeEach(func() {
			cn = &v1alpha1.ComputeNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-name",
					Namespace: "default",
					Labels: map[string]string{
						"app": "shardingsphere-proxy",
					},
					Annotations: map[string]string{
						"anno-key": "anno-value",
					},
				},
				Spec: v1alpha1.ComputeNodeSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app": "shardingsphere-proxy",
						},
					},
					Replicas: 2,
					Bootstrap: v1alpha1.BootstrapConfig{
						ServerConfig: v1alpha1.ServerConfig{
							Mode: v1alpha1.ComputeNodeServerMode{
								Repository: v1alpha1.Repository{
									Type: v1alpha1.RepositoryTypeZookeeper,
								},
							},
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
					StorageNodeConnector: &v1alpha1.StorageNodeConnector{
						Type:    v1alpha1.ConnectorTypeMySQL,
						Version: "5.1.47",
					},
					PortBindings: []v1alpha1.PortBinding{
						{
							Name:          "port",
							ContainerPort: 3307,
							ServicePort:   3307,
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, cn)).To(BeNil())
		})
		AfterEach(func() {
			Expect(k8sClient.Delete(ctx, cn)).To(BeNil())
		})

		It("verify deployment metadata", func() {
			createdDeploy := &appsv1.Deployment{}
			namespacedName := types.NamespacedName{Name: cn.Name, Namespace: cn.Namespace}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, createdDeploy)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			Expect(createdDeploy.Name).To(Equal(cn.Name))
			Expect(createdDeploy.Namespace).To(Equal(cn.Namespace))
			Expect(createdDeploy.Labels).To(Equal(cn.Labels))
		})

		It("verify service metadata", func() {
			createdService := &corev1.Service{}
			namespacedName := types.NamespacedName{Name: cn.Name, Namespace: cn.Namespace}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, createdService)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			Expect(createdService.Name).To(Equal(cn.Name))
			Expect(createdService.Namespace).To(Equal(cn.Namespace))
			Expect(createdService.Labels).To(Equal(cn.Labels))
		})

		It("verify configmap metadata", func() {
			createdConfigMap := &corev1.ConfigMap{}
			namespacedName := types.NamespacedName{Name: cn.Name, Namespace: cn.Namespace}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, createdConfigMap)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			Expect(createdConfigMap.Name).To(Equal(cn.Name))
			Expect(createdConfigMap.Namespace).To(Equal(cn.Namespace))
			Expect(createdConfigMap.Labels).To(Equal(cn.Labels))
		})
	})

	Context("verify spec", func() {
		var cn *v1alpha1.ComputeNode
		BeforeEach(func() {
			cn = &v1alpha1.ComputeNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-name",
					Namespace: "default",
					Labels: map[string]string{
						"app": "shardingsphere-proxy",
					},
				},
				Spec: v1alpha1.ComputeNodeSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app": "shardingsphere-proxy",
						},
					},
					Replicas: 2,
					Bootstrap: v1alpha1.BootstrapConfig{
						ServerConfig: v1alpha1.ServerConfig{
							Mode: v1alpha1.ComputeNodeServerMode{
								Repository: v1alpha1.Repository{
									Type: v1alpha1.RepositoryTypeZookeeper,
								},
							},
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
					StorageNodeConnector: &v1alpha1.StorageNodeConnector{
						Type:    v1alpha1.ConnectorTypeMySQL,
						Version: "5.1.47",
					},
					PortBindings: []v1alpha1.PortBinding{
						{
							Name:          "port",
							ContainerPort: 3307,
							ServicePort:   3307,
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, cn)).To(BeNil())
		})
		AfterEach(func() {
			Expect(k8sClient.Delete(ctx, cn)).To(BeNil())
		})

		It("verify deployment spec", func() {
			createdDeploy := &appsv1.Deployment{}
			namespacedName := types.NamespacedName{Name: cn.Name, Namespace: cn.Namespace}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, createdDeploy)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			Expect(createdDeploy.Spec.Selector.MatchLabels).To(Equal(cn.Spec.Selector.MatchLabels))
			Expect(*createdDeploy.Spec.Replicas).To(Equal(cn.Spec.Replicas))
			//TODO: Add more tests for DeploymentSpec
		})

		It("verify service spec", func() {
			createdService := &corev1.Service{}
			namespacedName := types.NamespacedName{Name: cn.Name, Namespace: cn.Namespace}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, createdService)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			//TODO: Add more tests for ServiceSpec
		})

		It("verify configmap spec", func() {
			createdConfigMap := &corev1.ConfigMap{}
			namespacedName := types.NamespacedName{Name: cn.Name, Namespace: cn.Namespace}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, createdConfigMap)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			//TODO: Add more tests for ConfigMapSpec
		})

	})
})
