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

package computenode

import (
	"fmt"
	"testing"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func Test_NewDeployment(t *testing.T) {
	defaultMaxUnavailable := intstr.FromInt(0)
	defaultMaxSurge := intstr.FromInt(3)
	var defaultReplicas int32 = 2

	cases := []struct {
		id      int
		cn      *v1alpha1.ComputeNode
		exp     *v1.Deployment
		message string
	}{
		{
			id: 1,
			cn: &v1alpha1.ComputeNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-name",
					Namespace: "test-namespace",
					Labels: map[string]string{
						"k1": "v1",
						"k2": "v2",
					},
					Annotations: map[string]string{
						"anno1": "value1",
					},
				},
				Spec: v1alpha1.ComputeNodeSpec{
					StorageNodeConnector: &v1alpha1.StorageNodeConnector{
						Type:    v1alpha1.ConnectorTypeMySQL,
						Version: "5.1.47",
					},
					ServerVersion: "5.2.0",
					Replicas:      2,
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"ks1": "vs1",
						},
					},
					PortBindings: []v1alpha1.PortBinding{
						{
							Name:          "server",
							ContainerPort: 3307,
							ServicePort:   3307,
							Protocol:      corev1.ProtocolTCP,
						},
					},
					ServiceType: corev1.ServiceTypeClusterIP,
					Bootstrap: v1alpha1.BootstrapConfig{
						ServerConfig: v1alpha1.ServerConfig{
							Authority: v1alpha1.ComputeNodeAuthority{
								Privilege: v1alpha1.ComputeNodePrivilege{
									Type: v1alpha1.AllPermitted,
								},
								Users: []v1alpha1.ComputeNodeUser{
									{
										User:     "root%",
										Password: "root",
									},
								},
							},
							Mode: v1alpha1.ComputeNodeServerMode{
								Type: v1alpha1.ModeTypeCluster,
								Repository: v1alpha1.Repository{
									Type: v1alpha1.RepositoryTypeZookeeper,
									Props: v1alpha1.Properties{
										"p1": "v1",
									},
								},
							},
							Props: v1alpha1.Properties{
								"p2": "v2",
							},
						},
					},
				},
			},
			exp: &v1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-name",
					Namespace: "test-namespace",
					Labels: map[string]string{
						"k1": "v1",
						"k2": "v2",
					},
					Annotations: map[string]string{
						"anno1": "value1",
					},
				},
				Spec: v1.DeploymentSpec{
					Replicas: &defaultReplicas,
					Strategy: v1.DeploymentStrategy{
						Type: v1.RollingUpdateDeploymentStrategyType,
						RollingUpdate: &v1.RollingUpdateDeployment{
							MaxUnavailable: &defaultMaxUnavailable,
							MaxSurge:       &defaultMaxSurge,
						},
					},
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"ks1": "vs1",
						},
					},
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								"k1": "v1",
								"k2": "v2",
							},
							// Annotations: map[string]string{
							// 	"anno1": "value1",
							// },
						},
						Spec: corev1.PodSpec{
							InitContainers: []corev1.Container{
								{
									Name:    "download-mysql-jar",
									Image:   "busybox:1.35.0",
									Command: []string{"/bin/sh", "-c", downloadMysqlJarScript},
									Env: []corev1.EnvVar{
										{
											Name:  defaultMySQLDriverEnvName,
											Value: "5.1.47",
										},
									},
									VolumeMounts: []corev1.VolumeMount{
										{
											Name:      defaultMySQLDriverVolumeName,
											MountPath: defaultExtlibPath,
										},
									},
								},
							},
							Containers: []corev1.Container{
								{
									Name:  defaultContainerName,
									Image: fmt.Sprintf("%s:%s", defaultImageName, "5.2.0"),
									// ImagePullPolicy: corev1.PullIfNotPresent,
									Ports: []corev1.ContainerPort{
										{
											Name:          "server",
											ContainerPort: 3307,
											Protocol:      corev1.ProtocolTCP,
										},
									},
									Env: []corev1.EnvVar{
										{
											Name:  defaultMySQLDriverEnvName,
											Value: "5.1.47",
										},
									},
									VolumeMounts: []corev1.VolumeMount{
										{
											Name:      defaultConfigVolumeName,
											MountPath: defaultConfigVolumeMountPath,
										},
										{
											Name:      defaultMySQLDriverVolumeName,
											SubPath:   relativeMySQLDriverMountName("5.1.47"),
											MountPath: absoluteMySQLDriverMountName(defaultExtlibPath, "5.1.47"),
										},
									},
								},
							},
							Volumes: []corev1.Volume{
								{
									Name: defaultConfigVolumeName,
									VolumeSource: corev1.VolumeSource{
										ConfigMap: &corev1.ConfigMapVolumeSource{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "test-name",
											},
										},
									},
								},
								{
									Name: defaultMySQLDriverVolumeName,
									VolumeSource: corev1.VolumeSource{
										EmptyDir: &corev1.EmptyDirVolumeSource{},
									},
								},
							},
						},
					},
				},
			},
			message: "case 1",
		},
		{
			id: 2,
			cn: &v1alpha1.ComputeNode{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-java-agent",
					Namespace: "test-namespace",
					Labels: map[string]string{
						"k1": "v1",
					},
					Annotations: map[string]string{
						defaultAnnotationJavaAgentEnabled: "true",
					},
				},
				Spec: v1alpha1.ComputeNodeSpec{
					StorageNodeConnector: &v1alpha1.StorageNodeConnector{
						Type:    v1alpha1.ConnectorTypeMySQL,
						Version: "5.1.47",
					},
					ServerVersion: "5.3.1",
					Replicas:      2,
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"k1": "v1",
						},
					},
					PortBindings: []v1alpha1.PortBinding{
						{
							Name:          "server",
							ContainerPort: 3307,
							ServicePort:   3307,
							Protocol:      corev1.ProtocolTCP,
						},
					},
					ServiceType: corev1.ServiceTypeClusterIP,
					Bootstrap: v1alpha1.BootstrapConfig{
						AgentConfig: v1alpha1.AgentConfig{
							Plugins: v1alpha1.AgentPlugin{
								Metrics: &v1alpha1.PluginMetrics{
									Prometheus: v1alpha1.Prometheus{
										Host:  "localhost",
										Port:  9090,
										Props: map[string]string{},
									},
								},
							},
						},
					},
				},
			},
			exp: &v1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-java-agent",
					Namespace: "test-namespace",
					Labels: map[string]string{
						"k1": "v1",
					},
					Annotations: map[string]string{
						"anno1": "value1",
					},
				},
				Spec: v1.DeploymentSpec{
					Replicas: &defaultReplicas,
					Strategy: v1.DeploymentStrategy{
						Type: v1.RollingUpdateDeploymentStrategyType,
						RollingUpdate: &v1.RollingUpdateDeployment{
							MaxUnavailable: &defaultMaxUnavailable,
							MaxSurge:       &defaultMaxSurge,
						},
					},
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"k1": "v1",
						},
					},
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								"k1": "v1",
							},
						},
						Spec: corev1.PodSpec{
							InitContainers: []corev1.Container{
								{
									Name:    "download-mysql-jar",
									Image:   "busybox:1.35.0",
									Command: []string{"/bin/sh", "-c", downloadMysqlJarScript},
									Env: []corev1.EnvVar{
										{
											Name:  defaultMySQLDriverEnvName,
											Value: "5.1.47",
										},
									},
									VolumeMounts: []corev1.VolumeMount{
										{
											Name:      defaultMySQLDriverVolumeName,
											MountPath: defaultExtlibPath,
										},
									},
								},
								{
									Name:    "download-agent-bin-jar",
									Image:   "busybox:1.35.0",
									Command: []string{"/bin/sh", "-c", downloadAgentJarScript},
									Env: []corev1.EnvVar{
										{
											Name:  defaultAgentBinVersionEnvName,
											Value: "5.3.1",
										},
									},
									VolumeMounts: []corev1.VolumeMount{
										{
											Name:      defaultJavaAgentVolumeName,
											MountPath: defaultJavaAgentVolumeMountPath,
										},
									},
								},
							},
							Containers: []corev1.Container{
								{
									Name:  defaultContainerName,
									Image: fmt.Sprintf("%s:%s", defaultImageName, "5.3.1"),
									Ports: []corev1.ContainerPort{
										{
											Name:          "server",
											ContainerPort: 3307,
											Protocol:      corev1.ProtocolTCP,
										},
									},
									Env: []corev1.EnvVar{
										{
											Name:  defaultJavaToolOptionsName,
											Value: fmt.Sprintf(defaultJavaAgentEnvValue, "5.3.1"),
										},
									},
									VolumeMounts: []corev1.VolumeMount{
										{
											Name:      defaultConfigVolumeName,
											MountPath: defaultConfigVolumeMountPath,
										},
										{
											Name:      defaultMySQLDriverVolumeName,
											SubPath:   relativeMySQLDriverMountName("5.1.47"),
											MountPath: absoluteMySQLDriverMountName(defaultExtlibPath, "5.1.47"),
										},
										{
											Name:      defaultJavaAgentVolumeName,
											MountPath: defaultJavaAgentVolumeMountPath,
										},
										{
											Name:      defaultJavaAgentConfigVolumeName,
											MountPath: defaultJavaAgentConfigVolumeMountPath,
										},
									},
								},
							},
							Volumes: []corev1.Volume{
								{
									Name: defaultConfigVolumeName,
									VolumeSource: corev1.VolumeSource{
										ConfigMap: &corev1.ConfigMapVolumeSource{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "test-java-agent",
											},
										},
									},
								},
								{
									Name: defaultMySQLDriverVolumeName,
									VolumeSource: corev1.VolumeSource{
										EmptyDir: &corev1.EmptyDirVolumeSource{},
									},
								},
								{
									Name: defaultJavaAgentVolumeName,
									VolumeSource: corev1.VolumeSource{
										EmptyDir: &corev1.EmptyDirVolumeSource{},
									},
								},
								{
									Name: defaultJavaAgentConfigVolumeName,
									VolumeSource: corev1.VolumeSource{
										ConfigMap: &corev1.ConfigMapVolumeSource{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "test-java-agent",
											},
											Items: []corev1.KeyToPath{
												{
													Key:  ConfigDataKeyForAgent,
													Path: ConfigDataKeyForAgent,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			message: "case 2",
		},
	}

	for _, c := range cases {
		act := NewDeployment(c.cn)
		assertObjectMeta(t, c.exp.ObjectMeta, act.ObjectMeta)
		assertDeploymentSpec(t, c.exp.Spec, act.Spec)
	}
}

func assertObjectMeta(t *testing.T, exp, act metav1.ObjectMeta) bool {
	return assert.Equal(t, exp.Name, act.Name, "name should be equal") &&
		assert.Equal(t, exp.Namespace, act.Namespace, "namespace should be equal") &&
		assert.Equal(t, exp.Labels, act.Labels, "labels should be equal")
}

func assertDeploymentSpec(t *testing.T, exp, act v1.DeploymentSpec) bool {
	return assertRollingUpdateDeployment(t, *exp.Strategy.RollingUpdate, *act.Strategy.RollingUpdate) &&
		assert.Equal(t, exp.Selector, act.Selector, "selectors should be equal") &&
		assert.Equal(t, exp.Replicas, act.Replicas, "replicas should be equal") &&
		assertTemplateSpec(t, exp.Template, act.Template)
}

func assertRollingUpdateDeployment(t *testing.T, exp, act v1.RollingUpdateDeployment) bool {
	return assert.Equal(t, exp.MaxSurge, act.MaxSurge, "maxSurge should be equal") &&
		assert.Equal(t, exp.MaxUnavailable, act.MaxUnavailable, "maxUnavailable should be equal")
}

func assertTemplateSpec(t *testing.T, exp, act corev1.PodTemplateSpec) bool {
	return assertObjectMeta(t, exp.ObjectMeta, act.ObjectMeta) &&
		assertPodSpec(t, exp.Spec, act.Spec)
}

func assertPodSpec(t *testing.T, exp, act corev1.PodSpec) bool {
	return assert.ElementsMatch(t, exp.InitContainers, act.InitContainers, "init containers should be equal") &&
		assert.ElementsMatch(t, exp.Containers, act.Containers, "containers should be equal") &&
		assert.ElementsMatch(t, exp.Volumes, act.Volumes, "volumes should be equal")
}
