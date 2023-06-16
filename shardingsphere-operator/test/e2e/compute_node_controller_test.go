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
	"fmt"
	"time"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

/*
	func Test_GetReadyProxyInstances(t *testing.T) {
		// create sample PodList
		podlist := v1.PodList{
			Items: []v1.Pod{
				{
					Status: v1.PodStatus{
						Phase: v1.PodRunning,
						Conditions: []v1.PodCondition{
							{
								Type:   v1.PodReady,
								Status: v1.ConditionTrue,
							},
						},
						ContainerStatuses: []v1.ContainerStatus{
							{
								Name:  "shardingsphere-proxy",
								Ready: true,
							},
						},
					},
				},
				{
					Status: v1.PodStatus{
						Phase: v1.PodRunning,
						Conditions: []v1.PodCondition{
							{
								Type:   v1.PodReady,
								Status: v1.ConditionTrue,
							},
						},
						ContainerStatuses: []v1.ContainerStatus{
							{
								Name:  "another-container",
								Ready: true,
							},
						},
					},
				},
				{
					Status: v1.PodStatus{
						Phase: v1.PodRunning,
						Conditions: []v1.PodCondition{
							{
								Type:   v1.PodReady,
								Status: v1.ConditionFalse,
							},
						},
						ContainerStatuses: []v1.ContainerStatus{
							{
								Name:  "shardingsphere-proxy",
								Ready: false,
							},
						},
					},
				},
				{
					Status: v1.PodStatus{
						Phase: v1.PodPending,
						Conditions: []v1.PodCondition{
							{
								Type:   v1.PodReady,
								Status: v1.ConditionTrue,
							},
						},
						ContainerStatuses: []v1.ContainerStatus{
							{
								Name:  "shardingsphere-proxy",
								Ready: true,
							},
						},
					},
				},
			},
		}

		// expected result is 1 because only one pod has a ready shardingsphere-proxy container
		expected := int32(1)

		// call the function to get the actual result
		actual := getReadyProxyInstances(&podlist)

		// compare the expected and actual results
		if actual != expected {
			t.Errorf("getReadyInstances returned %d, expected %d", actual, expected)
		}
	}
*/

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

	Context("verify deployment spec", func() {
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

		It("verify deployment spec basic", func() {
			createdDeploy := &appsv1.Deployment{}
			namespacedName := types.NamespacedName{Name: cn.Name, Namespace: cn.Namespace}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, createdDeploy)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			Expect(createdDeploy.Spec.Selector.MatchLabels).To(Equal(cn.Spec.Selector.MatchLabels))
			Expect(*createdDeploy.Spec.Replicas).To(Equal(cn.Spec.Replicas))
		})
		//TODO: Add more tests for DeploymentSpec
		It("verify deployment spec volume", func() {})
		It("verify deployment spec template spec init containers", func() {})
		It("verify deployment spec template spec containers", func() {})
		It("verify deployment spec template spec volume mounts", func() {})
		It("verify deployment spec template spec ports", func() {})
		It("verify deployment spec template spec image tag", func() {})
	})

	Context("verify service spec", func() {
		//TODO: Add BeforeEach and AfterEach here
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

		It("verify service spec selectors", func() {
			createdService := &corev1.Service{}
			namespacedName := types.NamespacedName{Name: cn.Name, Namespace: cn.Namespace}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, createdService)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			Expect(createdService.Spec.Selector).To(Equal(cn.Spec.Selector.MatchLabels))
		})
		//TODO: Add more tests for ServiceSpec
		It("verify service spec ports", func() {
			//e.g. Expect().To(Equal())
		})
	})

	Context("verify configmap spec", func() {
		//TODO: Add BeforeEach and AfterEach here
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
		It("verify configmap spec serverconfig", func() {
			createdConfigMap := &corev1.ConfigMap{}
			namespacedName := types.NamespacedName{Name: cn.Name, Namespace: cn.Namespace}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, createdConfigMap)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			//e.g. Expect().To(Equal())
		})

		It("verify configmap spec logback", func() {
			createdConfigMap := &corev1.ConfigMap{}
			namespacedName := types.NamespacedName{Name: cn.Name, Namespace: cn.Namespace}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, createdConfigMap)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())
			//TODO: Add more tests for ConfigMapSpec
			//e.g. Expect().To(Equal())
		})
	})

})

var _ = Describe("GetNamespacedByName", func() {
	Context("Assert Get ConfigMap ", func() {
		var cn *v1alpha1.ComputeNode
		BeforeEach(func() {
			cn = &v1alpha1.ComputeNode{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ComputeNode",
					APIVersion: fmt.Sprintf("%s/%s", v1alpha1.GroupVersion.Group, v1alpha1.GroupVersion.Version),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "default",
					Labels: map[string]string{
						"test_key": "test_value",
					},
				},
				Spec: v1alpha1.ComputeNodeSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"test_key": "test_value",
						},
					},
					PortBindings: []v1alpha1.PortBinding{
						{
							Name:          "port",
							ContainerPort: 3307,
							ServicePort:   3307,
						},
					},
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
						AgentConfig: v1alpha1.AgentConfig{
							Plugins: &v1alpha1.AgentPlugin{
								Logging: &v1alpha1.PluginLogging{
									File: v1alpha1.LoggingFile{
										Props: v1alpha1.Properties{
											"test_logging_key": "test_logging_value",
										},
									},
								},
								Metrics: &v1alpha1.PluginMetrics{
									Prometheus: v1alpha1.Prometheus{
										Host: "test_host",
										Port: 1234,
										Props: v1alpha1.Properties{
											"test_metrics_key": "test_metrics_value",
										},
									},
								},
								Tracing: &v1alpha1.PluginTracing{
									OpenTracing: v1alpha1.OpenTracing{
										Props: v1alpha1.Properties{
											"test_opentracing_key": "test_opentracing_value",
										},
									},
									OpenTelemetry: v1alpha1.OpenTelemetry{
										Props: v1alpha1.Properties{
											"test_opentelemetry_key": "test_opentelemetry_value",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, cn)).To(BeNil())
		})
		AfterEach(func() {
			Expect(k8sClient.Delete(ctx, cn)).To(BeNil())
		})

		It("get configmap should be equal", func() {
			factory := configmap.NewConfigMapFactory(cn)
			gvk := schema.GroupVersionKind{
				Group:   "shardingsphere.apache.org",
				Version: "v1alpha1",
				Kind:    "ComputeNode",
			}

			meta := cn.GetObjectMeta()
			cm := factory.NewConfigMapBuilder(meta, gvk).Build()

			expect := &corev1.ConfigMap{}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, types.NamespacedName{
					Namespace: cn.Namespace,
					Name:      cn.Name,
				}, expect)
				return err == nil
			}, time.Second*10, time.Millisecond*250).Should(BeTrue())

			Expect(expect).To(Not(BeNil()))
			Expect(err).To(BeNil())
			Expect(expect.Name).To(Equal(cm.Name))
			Expect(expect.Namespace).To(Equal(cm.Namespace))
			Expect(expect.Data).To(Equal(cm.Data))
		})
	})
})

var _ = Describe("Default ConfigMap", func() {
	var (
		expect = &corev1.ConfigMap{}
		cn     = &v1alpha1.ComputeNode{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ComputeNode",
				APIVersion: fmt.Sprintf("%s/%s", v1alpha1.GroupVersion.Group, v1alpha1.GroupVersion.Version),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test_name",
				Namespace: "test_namespace",
				Labels: map[string]string{
					"test_key": "test_value",
				},
			},
		}
	)

	BeforeEach(func() {
		expect.Name = "test_name"
		expect.Namespace = "test_namespace"
		expect.Labels = map[string]string{
			"test_key": "test_value",
		}
		expect.Data = map[string]string{}
		expect.Data[configmap.ConfigDataKeyForLogback] = configmap.DefaultLogback
		expect.Data[configmap.ConfigDataKeyForServer] = configmap.DefaultServerConfig
		expect.Data[configmap.ConfigDataKeyForAgent] = ""
	})

	Context("Assert ObjectMeta", func() {
		c := configmap.NewConfigMapClient(nil)
		cm := c.Build(ctx, cn)

		It("name should be equal", func() {
			Expect(expect.Name).To(Equal(cm.Name))
		})
		It("namespace should be equal", func() {
			Expect(expect.Namespace).To(Equal(cm.Namespace))
		})
		It("labels should be equal", func() {
			Expect(expect.Labels).To(Equal(cm.Labels))
		})
	})

	Context("Assert Default Spec Data", func() {
		c := configmap.NewConfigMapClient(nil)
		cm := c.Build(ctx, cn)

		It("default server config should be equal", func() {
			Expect(expect.Data[configmap.ConfigDataKeyForServer]).To(Equal(cm.Data[configmap.ConfigDataKeyForServer]))
		})
		It("default logback should be equal", func() {
			Expect(expect.Data[configmap.ConfigDataKeyForLogback]).To(Equal(cm.Data[configmap.ConfigDataKeyForLogback]))
		})
		It("default agent config should be equal", func() {
			Expect(expect.Data[configmap.ConfigDataKeyForAgent]).To(Equal(cm.Data[configmap.ConfigDataKeyForAgent]))
		})
	})

	Context("Assert Update Spec Data", func() {
		cn.TypeMeta = metav1.TypeMeta{
			Kind:       "ComputeNode",
			APIVersion: fmt.Sprintf("%s/%s", v1alpha1.GroupVersion.Group, v1alpha1.GroupVersion.Version),
		}
		cn.ObjectMeta = metav1.ObjectMeta{
			Name:      "test_name",
			Namespace: "test_namespace",
			Labels: map[string]string{
				"test_key": "test_value",
			},
			Annotations: map[string]string{
				"test_anno_key": "test_anno_value",
			},
		}
		cn.Spec.Bootstrap = v1alpha1.BootstrapConfig{
			ServerConfig: v1alpha1.ServerConfig{
				Authority: v1alpha1.ComputeNodeAuthority{
					Users: []v1alpha1.ComputeNodeUser{
						{
							User:     "test_user@%",
							Password: "test_password",
						},
					},
					Privilege: v1alpha1.ComputeNodePrivilege{
						Type: v1alpha1.AllPermitted,
					},
				},
				Mode: v1alpha1.ComputeNodeServerMode{
					Type: v1alpha1.ModeTypeCluster,
					Repository: v1alpha1.Repository{
						Type: v1alpha1.RepositoryTypeZookeeper,
						Props: v1alpha1.Properties{
							"test_repo_key": "test_repo_value",
						},
					},
				},
				Props: v1alpha1.Properties{
					"test_prop_key": "test_prop_value",
				},
			},
		}

		c := configmap.NewConfigMapClient(nil)
		cm := c.Build(ctx, cn)
		cm = configmap.UpdateComputeNodeConfigMap(cn, cm)
		cfg := &v1alpha1.ServerConfig{}
		err := yaml.Unmarshal([]byte(cm.Data[configmap.ConfigDataKeyForServer]), &cfg)
		if err != nil {
			fmt.Printf("Err: %s\n", err)
		}

		It("server config should be equal", func() {
			Expect(cfg.Authority).To(Equal(cn.Spec.Bootstrap.ServerConfig.Authority))
			Expect(cfg.Mode).To(Equal(cn.Spec.Bootstrap.ServerConfig.Mode))
			Expect(cfg.Props).To(Equal(cn.Spec.Bootstrap.ServerConfig.Props))
		})
		It("default logback should be equal", func() {
			Expect(expect.Data[configmap.ConfigDataKeyForLogback]).To(Equal(cm.Data[configmap.ConfigDataKeyForLogback]))
		})
		It("default agent config should be equal", func() {
			Expect(expect.Data[configmap.ConfigDataKeyForAgent]).To(Equal(cm.Data[configmap.ConfigDataKeyForAgent]))
		})
	})
})

var _ = Describe("Standalone Server Config", func() {
	Context("Assert Simple Service Config Data", func() {
		cn := &v1alpha1.ComputeNode{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ComputeNode",
				APIVersion: fmt.Sprintf("%s/%s", v1alpha1.GroupVersion.Group, v1alpha1.GroupVersion.Version),
			},
			Spec: v1alpha1.ComputeNodeSpec{
				Bootstrap: v1alpha1.BootstrapConfig{
					ServerConfig: v1alpha1.ServerConfig{
						Mode: v1alpha1.ComputeNodeServerMode{
							Type: v1alpha1.ModeTypeStandalone,
						},
					},
				},
			},
		}

		expect := &v1alpha1.ServerConfig{}
		c := configmap.NewConfigMapClient(nil)
		cm := c.Build(ctx, cn)
		err := yaml.Unmarshal([]byte(cm.Data[configmap.ConfigDataKeyForServer]), &expect)
		if err != nil {
			fmt.Printf("Err: %s\n", err)
		}

		It("server config authority should be equal", func() {
			Expect(expect.Authority).To(Equal(cn.Spec.Bootstrap.ServerConfig.Authority))
		})
		It("server config mode should be equal", func() {
			Expect(expect.Mode).To(Equal(cn.Spec.Bootstrap.ServerConfig.Mode))
		})
		It("server config props should be equal", func() {
			Expect(expect.Props).To(Equal(cn.Spec.Bootstrap.ServerConfig.Props))
		})
	})

	Context("Assert Full Service Config Data", func() {
		cn := &v1alpha1.ComputeNode{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ComputeNode",
				APIVersion: fmt.Sprintf("%s/%s", v1alpha1.GroupVersion.Group, v1alpha1.GroupVersion.Version),
			},
			Spec: v1alpha1.ComputeNodeSpec{
				Bootstrap: v1alpha1.BootstrapConfig{
					ServerConfig: v1alpha1.ServerConfig{
						Authority: v1alpha1.ComputeNodeAuthority{
							Users: []v1alpha1.ComputeNodeUser{
								{
									User:     "test_user@%",
									Password: "test_password",
								},
							},
							Privilege: v1alpha1.ComputeNodePrivilege{
								Type: v1alpha1.AllPermitted,
							},
						},
						Mode: v1alpha1.ComputeNodeServerMode{
							Type: v1alpha1.ModeTypeStandalone,
						},
						Props: v1alpha1.Properties{
							"test_prop_key": "test_prop_value",
						},
					},
				},
			},
		}

		expect := &v1alpha1.ServerConfig{}
		c := configmap.NewConfigMapClient(nil)
		cm := c.Build(ctx, cn)
		err := yaml.Unmarshal([]byte(cm.Data[configmap.ConfigDataKeyForServer]), &expect)
		if err != nil {
			fmt.Printf("Err: %s\n", err)
		}
		It("server config authority should be equal", func() {
			Expect(expect.Authority).To(Equal(cn.Spec.Bootstrap.ServerConfig.Authority))
		})
		It("server config mode should be equal", func() {
			Expect(expect.Mode).To(Equal(cn.Spec.Bootstrap.ServerConfig.Mode))
		})
		It("server config props should be equal", func() {
			Expect(expect.Props).To(Equal(cn.Spec.Bootstrap.ServerConfig.Props))
		})
	})
})

var _ = Describe("Cluster Server Config", func() {
	var (
		expect = &v1alpha1.ServerConfig{}
		cn     = &v1alpha1.ComputeNode{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ComputeNode",
				APIVersion: fmt.Sprintf("%s/%s", v1alpha1.GroupVersion.Group, v1alpha1.GroupVersion.Version),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test_name",
				Namespace: "test_namespace",
				Labels: map[string]string{
					"test_key": "test_value",
				},
			},
			Spec: v1alpha1.ComputeNodeSpec{
				Bootstrap: v1alpha1.BootstrapConfig{
					ServerConfig: v1alpha1.ServerConfig{
						Authority: v1alpha1.ComputeNodeAuthority{
							Users: []v1alpha1.ComputeNodeUser{
								{
									User:     "test_user@%",
									Password: "test_password",
								},
							},
							Privilege: v1alpha1.ComputeNodePrivilege{
								Type: v1alpha1.AllPermitted,
							},
						},
						Mode: v1alpha1.ComputeNodeServerMode{
							Type: v1alpha1.ModeTypeCluster,
							Repository: v1alpha1.Repository{
								Type: v1alpha1.RepositoryTypeZookeeper,
								Props: v1alpha1.Properties{
									"test_repo_key": "test_repo_value",
								},
							},
						},
						Props: v1alpha1.Properties{
							"test_prop_key": "test_prop_value",
						},
					},
				},
			},
		}
	)

	BeforeEach(func() {
		c := configmap.NewConfigMapClient(nil)
		cm := c.Build(ctx, cn)

		err := yaml.Unmarshal([]byte(cm.Data[configmap.ConfigDataKeyForServer]), &expect)
		if err != nil {
			fmt.Printf("Err: %s\n", err)
		}
	})

	Context("Assert Service Config Data", func() {
		It("server config authority should be equal", func() {
			Expect(expect.Authority).To(Equal(cn.Spec.Bootstrap.ServerConfig.Authority))
		})
		It("server config mode should be equal", func() {
			Expect(expect.Mode).To(Equal(cn.Spec.Bootstrap.ServerConfig.Mode))
		})
		It("server config props should be equal", func() {
			Expect(expect.Props).To(Equal(cn.Spec.Bootstrap.ServerConfig.Props))
		})
	})
})

var _ = Describe("Logback Config", func() {
	Context("Assert Logback Config Data From Annotations", func() {
		var (
			expect = ""
			cn     = &v1alpha1.ComputeNode{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ComputeNode",
					APIVersion: fmt.Sprintf("%s/%s", v1alpha1.GroupVersion.Group, v1alpha1.GroupVersion.Version),
				},
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						configmap.AnnoLogbackConfig: "test_logback_value",
					},
				},
				Spec: v1alpha1.ComputeNodeSpec{
					Bootstrap: v1alpha1.BootstrapConfig{
						LogbackConfig: configmap.DefaultLogback,
					},
				},
			}
		)

		c := configmap.NewConfigMapClient(nil)
		cm := c.Build(ctx, cn)
		expect = "test_logback_value"

		It("Logback config should be equal", func() {
			Expect(expect).To(Equal(cm.Data[configmap.ConfigDataKeyForLogback]))
		})
	})

	Context("Assert Logback Config Data", func() {
		var (
			expect = ""
			cn     = &v1alpha1.ComputeNode{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ComputeNode",
					APIVersion: fmt.Sprintf("%s/%s", v1alpha1.GroupVersion.Group, v1alpha1.GroupVersion.Version),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test_name",
					Namespace: "test_namespace",
					Labels: map[string]string{
						"test_key": "test_value",
					},
				},
				Spec: v1alpha1.ComputeNodeSpec{
					Bootstrap: v1alpha1.BootstrapConfig{
						LogbackConfig: configmap.DefaultLogback,
					},
				},
			}
		)

		c := configmap.NewConfigMapClient(nil)
		cm := c.Build(ctx, cn)
		expect = configmap.DefaultLogback

		It("Logback config should be equal", func() {
			Expect(expect).To(Equal(cm.Data[configmap.ConfigDataKeyForLogback]))
		})
	})
})

var _ = Describe("Agent Config", func() {
	Context("Assert Full Agent Config Data", func() {
		var (
			expect = &v1alpha1.AgentConfig{}
			cn     = &v1alpha1.ComputeNode{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ComputeNode",
					APIVersion: fmt.Sprintf("%s/%s", v1alpha1.GroupVersion.Group, v1alpha1.GroupVersion.Version),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test_name",
					Namespace: "test_namespace",
					Labels: map[string]string{
						"test_key": "test_value",
					},
				},
				Spec: v1alpha1.ComputeNodeSpec{
					Bootstrap: v1alpha1.BootstrapConfig{
						AgentConfig: v1alpha1.AgentConfig{
							Plugins: &v1alpha1.AgentPlugin{
								Logging: &v1alpha1.PluginLogging{
									File: v1alpha1.LoggingFile{
										Props: v1alpha1.Properties{
											"test_logging_key": "test_logging_value",
										},
									},
								},
								Metrics: &v1alpha1.PluginMetrics{
									Prometheus: v1alpha1.Prometheus{
										Host: "test_host",
										Port: 1234,
										Props: v1alpha1.Properties{
											"test_metrics_key": "test_metrics_value",
										},
									},
								},
								Tracing: &v1alpha1.PluginTracing{
									OpenTracing: v1alpha1.OpenTracing{
										Props: v1alpha1.Properties{
											"test_opentracing_key": "test_opentracing_value",
										},
									},
									OpenTelemetry: v1alpha1.OpenTelemetry{
										Props: v1alpha1.Properties{
											"test_opentelemetry_key": "test_opentelemetry_value",
										},
									},
								},
							},
						},
					},
				},
			}
		)

		c := configmap.NewConfigMapClient(nil)
		cm := c.Build(ctx, cn)

		err := yaml.Unmarshal([]byte(cm.Data[configmap.ConfigDataKeyForAgent]), &expect)
		if err != nil {
			fmt.Printf("Err: %s\n", err)
		}

		It("agent config plugins should be equal", func() {
			Expect(expect.Plugins).To(Equal(cn.Spec.Bootstrap.AgentConfig.Plugins))
		})
		It("agent config logging should be equal", func() {
			Expect(expect.Plugins.Logging).To(Equal(cn.Spec.Bootstrap.AgentConfig.Plugins.Logging))
		})
		It("agent config metrics should be equal", func() {
			Expect(expect.Plugins.Metrics).To(Equal(cn.Spec.Bootstrap.AgentConfig.Plugins.Metrics))
		})
		It("agent config tracing should be equal", func() {
			Expect(expect.Plugins.Tracing).To(Equal(cn.Spec.Bootstrap.AgentConfig.Plugins.Tracing))
		})
	})
})
