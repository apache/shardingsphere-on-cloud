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

package reconcile

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func Test_ConstructCascadingDeployment(t *testing.T) {
	cases := []struct {
		proxy   *v1alpha1.ShardingSphereProxy
		exp     *appsv1.Deployment
		message string
	}{
		{
			exp:     &appsv1.Deployment{},
			message: "Nil ShardingSphereProxy definition should lead to empty Deployment",
		},
		{
			proxy:   &v1alpha1.ShardingSphereProxy{},
			exp:     &appsv1.Deployment{},
			message: "Empty ShardingSphereProxy definition should lead to empty Deployment",
		},
		{
			proxy: &v1alpha1.ShardingSphereProxy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "testname",
					Namespace: "testnamespace",
					Annotations: map[string]string{
						AnnoRollingUpdateMaxSurge:       "3",
						AnnoRollingUpdateMaxUnavailable: "1",
					},
				},
				Spec: v1alpha1.ProxySpec{
					Version: "5.1.2",
					// ServiceType: ServiceTypeNodePort,
					Replicas:         3,
					AutomaticScaling: &v1alpha1.AutomaticScaling{},
					ImagePullSecrets: []v1.LocalObjectReference{},
					ProxyConfigName:  "shardingsphere-proxy-config",
					Port:             3307,
					MySQLDriver:      &v1alpha1.MySQLDriver{},
					Resources:        v1.ResourceRequirements{},
					LivenessProbe: &v1.Probe{
						ProbeHandler: v1.ProbeHandler{
							TCPSocket: &v1.TCPSocketAction{},
						},
						InitialDelaySeconds: 30,
						TimeoutSeconds:      3,
						PeriodSeconds:       5,
						SuccessThreshold:    1,
						FailureThreshold:    3,
					},
					ReadinessProbe: &v1.Probe{
						ProbeHandler: v1.ProbeHandler{
							TCPSocket: &v1.TCPSocketAction{},
						},
						InitialDelaySeconds: 30,
						TimeoutSeconds:      3,
						PeriodSeconds:       5,
						SuccessThreshold:    1,
						FailureThreshold:    3,
					},
					StartupProbe: &v1.Probe{
						ProbeHandler: v1.ProbeHandler{
							TCPSocket: &v1.TCPSocketAction{},
						},
						InitialDelaySeconds: 30,
						TimeoutSeconds:      3,
						PeriodSeconds:       5,
						SuccessThreshold:    1,
						FailureThreshold:    3,
					},
				},
			},
			exp: &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "testname",
					Namespace: "testnamespace",
				},
				Spec: appsv1.DeploymentSpec{
					Strategy: appsv1.DeploymentStrategy{
						Type: appsv1.RollingUpdateDeploymentStrategyType,
						RollingUpdate: &appsv1.RollingUpdateDeployment{
							MaxUnavailable: func(v int) *intstr.IntOrString { p := intstr.FromInt(v); return &p }(1),
							MaxSurge:       func(v int) *intstr.IntOrString { p := intstr.FromInt(v); return &p }(3),
						},
					},
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"apps": "testname",
						},
					},
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								"apps": "testname",
							},
						},
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								{
									Name:            "proxy",
									Image:           fmt.Sprintf("%s:%s", "apache/shardingsphere-proxy", "5.1.2"),
									ImagePullPolicy: v1.PullIfNotPresent,
									Ports: []v1.ContainerPort{
										{
											ContainerPort: 3307,
										},
									},
									Env: []v1.EnvVar{
										{
											Name:  "PORT",
											Value: strconv.FormatInt(int64(3307), 10),
										},
									},
									LivenessProbe: &v1.Probe{
										ProbeHandler: v1.ProbeHandler{
											TCPSocket: &v1.TCPSocketAction{},
										},
										InitialDelaySeconds: 30,
										TimeoutSeconds:      3,
										PeriodSeconds:       5,
										SuccessThreshold:    1,
										FailureThreshold:    3,
									},
									ReadinessProbe: &v1.Probe{
										ProbeHandler: v1.ProbeHandler{
											TCPSocket: &v1.TCPSocketAction{},
										},
										InitialDelaySeconds: 30,
										TimeoutSeconds:      3,
										PeriodSeconds:       5,
										SuccessThreshold:    1,
										FailureThreshold:    3,
									},
									StartupProbe: &v1.Probe{
										ProbeHandler: v1.ProbeHandler{
											TCPSocket: &v1.TCPSocketAction{},
										},
										InitialDelaySeconds: 30,
										TimeoutSeconds:      3,
										PeriodSeconds:       5,
										SuccessThreshold:    1,
										FailureThreshold:    3,
									},
									VolumeMounts: []v1.VolumeMount{
										{
											Name:      "config",
											MountPath: "/opt/shardingsphere-proxy/conf",
										},
									},
								},
							},
							Volumes: []v1.Volume{
								{
									Name: "config",
									VolumeSource: v1.VolumeSource{
										ConfigMap: &v1.ConfigMapVolumeSource{
											LocalObjectReference: v1.LocalObjectReference{
												Name: "shardingsphere-proxy-config",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			message: "Normal ShardingSphereProxy definition should lead to normal Deployment",
		},
	}

	for _, c := range cases {
		act := ConstructCascadingDeployment(c.proxy)
		assert.Equal(t, c.exp.ObjectMeta.Name, act.ObjectMeta.Name, c.message)
		assert.Equal(t, c.exp.ObjectMeta.Namespace, act.ObjectMeta.Namespace, c.message)
		// assert.EqualValues(t, c.exp.Spec, act.Spec, c.message)
		if c.proxy != nil {
			if c.proxy.Spec.AutomaticScaling != nil {
				assert.Equal(t, c.exp.Spec.Replicas, act.Spec.Replicas, c.message)
			}

			if len(c.exp.Spec.Template.Spec.Containers) > 0 {
				assert.EqualValues(t, c.exp.Spec.Template.Spec.Containers[0].Resources, act.Spec.Template.Spec.Containers[0].Resources, c.message)
			}

			if c.proxy.Spec.LivenessProbe != nil {
				assert.EqualValues(t, c.exp.Spec.Template.Spec.Containers[0].LivenessProbe, act.Spec.Template.Spec.Containers[0].LivenessProbe, c.message)
			}
			if c.proxy.Spec.ReadinessProbe != nil {
				assert.EqualValues(t, c.exp.Spec.Template.Spec.Containers[0].ReadinessProbe, act.Spec.Template.Spec.Containers[0].ReadinessProbe, c.message)
			}
			if c.proxy.Spec.StartupProbe != nil {
				assert.EqualValues(t, c.exp.Spec.Template.Spec.Containers[0].StartupProbe, act.Spec.Template.Spec.Containers[0].StartupProbe, c.message)
			}
			if len(c.proxy.Spec.ImagePullSecrets) > 0 {
				assert.EqualValues(t, c.exp.Spec.Template.Spec.ImagePullSecrets, act.Spec.Template.Spec.ImagePullSecrets, c.message)
			}

			if c.proxy.Annotations[AnnoRollingUpdateMaxSurge] != "" {
				assert.Equal(t, c.exp.Spec.Strategy.RollingUpdate.MaxSurge.StrVal, act.Spec.Strategy.RollingUpdate.MaxSurge.StrVal, c.message)
			} else {
				if c.exp.Name != "" && act.Spec.Strategy.RollingUpdate != nil {
					assert.Equal(t, 1, act.Spec.Strategy.RollingUpdate.MaxSurge.IntVal, c.message)
				}
			}

			if c.proxy.Annotations[AnnoRollingUpdateMaxUnavailable] != "" {
				assert.Equal(t, c.exp.Spec.Strategy.RollingUpdate.MaxUnavailable.StrVal, act.Spec.Strategy.RollingUpdate.MaxUnavailable.StrVal, c.message)
			} else {
				if c.exp.Name != "" && act.Spec.Strategy.RollingUpdate != nil {
					assert.Equal(t, 0, act.Spec.Strategy.RollingUpdate.MaxUnavailable.StrVal, c.message)
				}
			}
		}
	}
}

func Test_addInitContaienr(t *testing.T) {
	cases := []struct {
		deploy  *appsv1.Deployment
		mysql   *v1alpha1.MySQLDriver
		message string
	}{
		{
			deploy: &appsv1.Deployment{
				Spec: appsv1.DeploymentSpec{
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							InitContainers: []v1.Container{},
							Containers: []v1.Container{
								{
									VolumeMounts: []v1.VolumeMount{},
								},
							},
							Volumes: []v1.Volume{},
						},
					},
				},
			},
			mysql: &v1alpha1.MySQLDriver{
				Version: "5.1.47",
			},
			message: "Add InitContainer for MySQL Driver",
		},
	}

	for _, c := range cases {
		addInitContainer(c.deploy, c.mysql)
		assert.Equal(t, c.deploy.Spec.Template.Spec.InitContainers[0].Name, "download-mysql-connect", c.message)
	}
}

func Test_processOptionalParameter(t *testing.T) {
	/*
	   cases := []struct{

	   }{
	       {

	       },
	   }

	   for _, c := range cases {

	   }
	*/
}
func Test_UpdateDeployment(t *testing.T) {
	var rep int32 = 3
	cases := []struct {
		proxy   *v1alpha1.ShardingSphereProxy
		deploy  *appsv1.Deployment
		message string
	}{
		{
			proxy: &v1alpha1.ShardingSphereProxy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "testname",
					Namespace: "testnamespace",
				},
				Spec: v1alpha1.ProxySpec{
					Version: "5.1.2",
					// ServiceType: ServiceTypeNodePort,
					Replicas:         3,
					AutomaticScaling: &v1alpha1.AutomaticScaling{},
					ImagePullSecrets: []v1.LocalObjectReference{},
					ProxyConfigName:  "shardingsphere-proxy-config",
					Port:             3307,
					MySQLDriver:      &v1alpha1.MySQLDriver{},
					Resources:        v1.ResourceRequirements{},
					LivenessProbe: &v1.Probe{
						ProbeHandler: v1.ProbeHandler{
							TCPSocket: &v1.TCPSocketAction{},
						},
						InitialDelaySeconds: 30,
						TimeoutSeconds:      3,
						PeriodSeconds:       5,
						SuccessThreshold:    1,
						FailureThreshold:    3,
					},
					ReadinessProbe: &v1.Probe{
						ProbeHandler: v1.ProbeHandler{
							TCPSocket: &v1.TCPSocketAction{},
						},
						InitialDelaySeconds: 30,
						TimeoutSeconds:      3,
						PeriodSeconds:       5,
						SuccessThreshold:    1,
						FailureThreshold:    3,
					},
					StartupProbe: &v1.Probe{
						ProbeHandler: v1.ProbeHandler{
							TCPSocket: &v1.TCPSocketAction{},
						},
						InitialDelaySeconds: 30,
						TimeoutSeconds:      3,
						PeriodSeconds:       5,
						SuccessThreshold:    1,
						FailureThreshold:    3,
					},
				},
			},
			deploy: &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "testname",
					Namespace: "testnamespace",
				},
				Spec: appsv1.DeploymentSpec{
					Strategy: appsv1.DeploymentStrategy{
						Type: appsv1.RecreateDeploymentStrategyType,
					},
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"apps": "testname",
						},
					},
					Replicas: &rep,
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								"apps": "testname",
							},
						},
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								{
									Name:            "proxy",
									Image:           fmt.Sprintf("%s:%s", "apache/shardingsphere-proxy", "5.2.0"),
									ImagePullPolicy: v1.PullIfNotPresent,
									Ports: []v1.ContainerPort{
										{
											ContainerPort: 3307,
										},
									},
									Env: []v1.EnvVar{
										{
											Name:  "PORT",
											Value: strconv.FormatInt(int64(3307), 10),
										},
									},
									LivenessProbe: &v1.Probe{
										ProbeHandler: v1.ProbeHandler{
											TCPSocket: &v1.TCPSocketAction{},
										},
										InitialDelaySeconds: 30,
										TimeoutSeconds:      3,
										PeriodSeconds:       5,
										SuccessThreshold:    1,
										FailureThreshold:    3,
									},
									ReadinessProbe: &v1.Probe{
										ProbeHandler: v1.ProbeHandler{
											TCPSocket: &v1.TCPSocketAction{},
										},
										InitialDelaySeconds: 30,
										TimeoutSeconds:      3,
										PeriodSeconds:       5,
										SuccessThreshold:    1,
										FailureThreshold:    3,
									},
									StartupProbe: &v1.Probe{
										ProbeHandler: v1.ProbeHandler{
											TCPSocket: &v1.TCPSocketAction{},
										},
										InitialDelaySeconds: 30,
										TimeoutSeconds:      3,
										PeriodSeconds:       5,
										SuccessThreshold:    1,
										FailureThreshold:    3,
									},
									VolumeMounts: []v1.VolumeMount{
										{
											Name:      "config",
											MountPath: "/opt/shardingsphere-proxy/conf",
										},
									},
									Resources: v1.ResourceRequirements{},
								},
							},
							Volumes: []v1.Volume{
								{
									Name: "config",
									VolumeSource: v1.VolumeSource{
										ConfigMap: &v1.ConfigMapVolumeSource{
											LocalObjectReference: v1.LocalObjectReference{
												Name: "shardingsphere-proxy-config",
											},
										},
									},
								},
							},
						},
					},
				},
			},

			message: "Deployment should be updated",
		},
	}

	for _, c := range cases {
		exp := UpdateDeployment(c.proxy, c.deploy)
		assert.Equal(t, fmt.Sprintf("%s:%s", imageName, c.proxy.Spec.Version), exp.Spec.Template.Spec.Containers[0].Image, c.message)
		assert.Equal(t, c.proxy.Spec.Replicas, *exp.Spec.Replicas, c.message)
		assert.Equal(t, c.proxy.Spec.ProxyConfigName, exp.Spec.Template.Spec.Volumes[0].ConfigMap.Name, c.message)
		assert.Equal(t, c.proxy.Spec.Port, exp.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort, c.message)
		assert.EqualValues(t, c.proxy.Spec.Resources, exp.Spec.Template.Spec.Containers[0].Resources, c.message)
		assert.EqualValues(t, c.proxy.Spec.LivenessProbe, exp.Spec.Template.Spec.Containers[0].LivenessProbe, c.message)
		assert.EqualValues(t, c.proxy.Spec.ReadinessProbe, exp.Spec.Template.Spec.Containers[0].ReadinessProbe, c.message)
		assert.EqualValues(t, c.proxy.Spec.StartupProbe, exp.Spec.Template.Spec.Containers[0].StartupProbe, c.message)
	}
}
