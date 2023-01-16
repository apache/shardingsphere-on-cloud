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

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	DefaultExtlibPath = "/opt/shardingsphere-proxy/ext-lib"
	imageName         = "apache/shardingsphere-proxy"
)

func ComputeNodeNewDeployment(cn *v1alpha1.ComputeNode) *v1.Deployment {
	deploy := ComputeNodeDefaultDeployment(cn.GetObjectMeta(), cn.GroupVersionKind())

	// basic information
	deploy.Name = cn.Name
	deploy.Namespace = cn.Namespace
	deploy.Labels = cn.Labels
	deploy.Spec.Selector = cn.Spec.Selector
	deploy.Spec.Replicas = &cn.Spec.Replicas
	deploy.Spec.Template.Labels = cn.Labels
	deploy.Spec.Template.Spec.Containers[0].Image = fmt.Sprintf("%s:%s", imageName, cn.Spec.ServerVersion)
	if deploy.Spec.Template.Spec.Containers[0].Ports == nil {
		deploy.Spec.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{}
	}
	for _, pb := range cn.Spec.PortBindings {
		deploy.Spec.Template.Spec.Containers[0].Ports = append(deploy.Spec.Template.Spec.Containers[0].Ports, corev1.ContainerPort{
			Name:          pb.Name,
			HostIP:        pb.HostIP,
			ContainerPort: pb.ContainerPort,
			Protocol:      pb.Protocol,
		})
	}

	// additional information
	deploy.Spec.Template.Spec.Containers[0].Resources = cn.Spec.Resources
	for _, v := range deploy.Spec.Template.Spec.Volumes {
		if v.Name == "shardingsphere-proxy-config" {
			v.ConfigMap.LocalObjectReference.Name = cn.Name
		}
	}

	if cn.Spec.Probes != nil {
		if cn.Spec.Probes.StartupProbe != nil {
			deploy.Spec.Template.Spec.Containers[0].StartupProbe = cn.Spec.Probes.StartupProbe.DeepCopy()
		}
		if cn.Spec.Probes.LivenessProbe != nil {
			deploy.Spec.Template.Spec.Containers[0].LivenessProbe = cn.Spec.Probes.LivenessProbe.DeepCopy()
		}
		if cn.Spec.Probes.ReadinessProbe != nil {
			deploy.Spec.Template.Spec.Containers[0].ReadinessProbe = cn.Spec.Probes.ReadinessProbe.DeepCopy()
		}
	}
	if len(cn.Spec.ImagePullSecrets) > 0 {
		deploy.Spec.Template.Spec.ImagePullSecrets = cn.Spec.ImagePullSecrets
	}
	if cn.Spec.StorageNodeConnector != nil {
		if cn.Spec.StorageNodeConnector.Type == v1alpha1.ConnectorTypeMySQL {
			// add or update initContainer
			if len(deploy.Spec.Template.Spec.InitContainers) > 0 {
				for idx, v := range deploy.Spec.Template.Spec.InitContainers[0].Env {
					if v.Name == "MYSQL_CONNECTOR_VERSION" {
						deploy.Spec.Template.Spec.InitContainers[0].Env[idx].Value = cn.Spec.StorageNodeConnector.Version
					}
				}
			} else {
				deploy.Spec.Template.Spec.InitContainers = []corev1.Container{
					{
						Name:    "boostrap",
						Image:   "busybox:1.35.0",
						Command: []string{"/bin/sh", "-c", download_script},
						Env: []corev1.EnvVar{
							{
								Name:  "MYSQL_CONNECTOR_VERSION",
								Value: cn.Spec.StorageNodeConnector.Version,
							},
						},
						VolumeMounts: []corev1.VolumeMount{
							{
								Name:      "mysql-connector-java",
								MountPath: DefaultExtlibPath,
							},
						},
					},
				}

				deploy.Spec.Template.Spec.Containers[0].VolumeMounts = append(deploy.Spec.Template.Spec.Containers[0].VolumeMounts, corev1.VolumeMount{
					Name:      "mysql-connector-java",
					SubPath:   fmt.Sprintf("mysql-connector-java-%s.jar", cn.Spec.StorageNodeConnector.Version),
					MountPath: fmt.Sprintf("%s/mysql-connector-java-%s.jar", DefaultExtlibPath, cn.Spec.StorageNodeConnector.Version),
				})

				deploy.Spec.Template.Spec.Volumes = append(deploy.Spec.Template.Spec.Volumes, corev1.Volume{
					Name: "mysql-connector-java",
					VolumeSource: corev1.VolumeSource{
						EmptyDir: &corev1.EmptyDirVolumeSource{},
					},
				})
			}
		}
	}

	return deploy
}

const download_script = `wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/${MYSQL_CONNECTOR_VERSION}/mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar;
wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/${MYSQL_CONNECTOR_VERSION}/mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar.md5;
if [ $(md5sum /mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar | cut -d ' ' -f1) = $(cat /mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar.md5) ];
then echo success;
else echo failed;exit 1;fi;mv /mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar /opt/shardingsphere-proxy/ext-lib`

func ComputeNodeDefaultDeployment(meta metav1.Object, gvk schema.GroupVersionKind) *v1.Deployment {
	defaultMaxUnavailable := intstr.FromInt(0)
	defaultMaxSurge := intstr.FromInt(3)
	defaultImage := "apache/shardingsphere-proxy:5.3.0"

	return &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "shardingsphere-proxy",
			Namespace: "default",
			Labels:    map[string]string{},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(meta, gvk),
			},
		},
		Spec: v1.DeploymentSpec{
			Strategy: v1.DeploymentStrategy{
				Type: v1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &v1.RollingUpdateDeployment{
					MaxUnavailable: &defaultMaxUnavailable,
					MaxSurge:       &defaultMaxSurge,
				},
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "shardingsphere-proxy",
							Image:           defaultImage,
							ImagePullPolicy: corev1.PullIfNotPresent,
							Ports: []corev1.ContainerPort{
								{
									Name:          "proxy",
									ContainerPort: 3307,
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "shardingsphere-proxy-config",
									MountPath: "/opt/shardingsphere-proxy/conf",
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "shardingsphere-proxy-config",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "shardingsphere-proxy-config",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func ComputeNodeUpdateDeployment(cn *v1alpha1.ComputeNode, cur *v1.Deployment) *v1.Deployment {
	exp := &v1.Deployment{}
	exp.ObjectMeta = cur.ObjectMeta
	exp.ObjectMeta.ResourceVersion = ""
	exp.Labels = cur.Labels
	exp.Annotations = cur.Annotations
	exp.Spec = ComputeNodeNewDeployment(cn).Spec
	return exp
}
