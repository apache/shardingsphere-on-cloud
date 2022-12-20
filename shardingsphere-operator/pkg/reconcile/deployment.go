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
	"reflect"
	"strconv"
	"strings"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewDeployment(ssproxy *v1alpha1.ShardingSphereProxy) *v1.Deployment {
	return ConstructCascadingDeployment(ssproxy)
}
func ConstructCascadingDeployment(proxy *v1alpha1.ShardingSphereProxy) *v1.Deployment {
	if proxy == nil || reflect.DeepEqual(proxy, &v1alpha1.ShardingSphereProxy{}) {
		return &v1.Deployment{}
	}

	dp := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proxy.Name,
			Namespace: proxy.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(proxy.GetObjectMeta(), proxy.GroupVersionKind()),
			},
		},
		Spec: v1.DeploymentSpec{
			Strategy: v1.DeploymentStrategy{
				Type: v1.RecreateDeploymentStrategyType,
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"apps": proxy.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"apps": proxy.Name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "proxy",
							Image:           fmt.Sprintf("%s:%s", imageName, proxy.Spec.Version),
							ImagePullPolicy: corev1.PullIfNotPresent,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: proxy.Spec.Port,
								},
							},
							Env: []corev1.EnvVar{
								{
									Name:  "PORT",
									Value: strconv.FormatInt(int64(proxy.Spec.Port), 10),
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "config",
									MountPath: "/opt/shardingsphere-proxy/conf",
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "config",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: proxy.Spec.ProxyConfigName,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	if proxy.Spec.AutomaticScaling == nil {
		dp.Spec.Replicas = &proxy.Spec.Replicas
	}

	dp.Spec.Template.Spec.Containers[0].Resources = proxy.Spec.Resources

	if proxy.Spec.LivenessProbe != nil {
		dp.Spec.Template.Spec.Containers[0].LivenessProbe = proxy.Spec.LivenessProbe
	}
	if proxy.Spec.ReadinessProbe != nil {
		dp.Spec.Template.Spec.Containers[0].ReadinessProbe = proxy.Spec.ReadinessProbe
	}
	if proxy.Spec.StartupProbe != nil {
		dp.Spec.Template.Spec.Containers[0].StartupProbe = proxy.Spec.StartupProbe
	}
	if len(proxy.Spec.ImagePullSecrets) > 0 {
		dp.Spec.Template.Spec.ImagePullSecrets = proxy.Spec.ImagePullSecrets
	}
	return processOptionalParameter(proxy, dp)
}

func processOptionalParameter(proxy *v1alpha1.ShardingSphereProxy, dp *v1.Deployment) *v1.Deployment {
	if proxy.Spec.MySQLDriver != nil {
		addInitContainer(dp, proxy.Spec.MySQLDriver)
	}
	return dp
}

const script = `wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/${VERSION}/mysql-connector-java-${VERSION}.jar;
wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/${VERSION}/mysql-connector-java-${VERSION}.jar.md5;
if [ $(md5sum /mysql-connector-java-${VERSION}.jar | cut -d ' ' -f1) = $(cat /mysql-connector-java-${VERSION}.jar.md5) ];
then echo success;
else echo failed;exit 1;fi;mv /mysql-connector-java-${VERSION}.jar /opt/shardingsphere-proxy/ext-lib`

func addInitContainer(dp *v1.Deployment, mysql *v1alpha1.MySQLDriver) {
	if len(dp.Spec.Template.Spec.InitContainers) == 0 {
		dp.Spec.Template.Spec.Containers[0].VolumeMounts = append(dp.Spec.Template.Spec.Containers[0].VolumeMounts, corev1.VolumeMount{
			Name:      "mysql-connect-jar",
			MountPath: "/opt/shardingsphere-proxy/ext-lib",
		})

		dp.Spec.Template.Spec.Volumes = append(dp.Spec.Template.Spec.Volumes, corev1.Volume{
			Name: "mysql-connect-jar",
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		})
	}

	/*
			scriptStr := strings.Builder{}
			t1, _ := template.New("shell").Parse(`wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/{{ .Version }}/mysql-connector-java-{{ .Version }}.jar;
		wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/{{ .Version }}/mysql-connector-java-{{ .Version }}.jar.md5;
		if [ $(md5sum /mysql-connector-java-{{ .Version }}.jar | cut -d ' ' -f1) = $(cat /mysql-connector-java-{{ .Version }}.jar.md5) ];
		then echo success;
		else echo failed;exit 1;fi;mv /mysql-connector-java-{{ .Version }}.jar /opt/shardingsphere-proxy/ext-lib`)
			_ = t1.Execute(&scriptStr, mysql)
	*/

	dp.Spec.Template.Spec.InitContainers = []corev1.Container{
		{
			Name:  "download-mysql-connect",
			Image: "busybox:1.35.0",
			// Command: []string{"/bin/sh", "-c", scriptStr.String()},
			Command: []string{"/bin/sh", "-c", script},
			Env: []corev1.EnvVar{
				{
					Name:  "VERSION",
					Value: mysql.Version,
				},
			},
			VolumeMounts: []corev1.VolumeMount{
				{
					Name:      "mysql-connect-jar",
					MountPath: "/opt/shardingsphere-proxy/ext-lib",
				},
			},
		},
	}

}

// UpdateDeployment FIXME:merge UpdateDeployment and ConstructCascadingDeployment
func UpdateDeployment(proxy *v1alpha1.ShardingSphereProxy, act *v1.Deployment) (*v1.Deployment, bool) {
	exp := act.DeepCopy()
	var diff bool

	if proxy.Spec.AutomaticScaling == nil || !proxy.Spec.AutomaticScaling.Enable {
		exp.Spec.Replicas, diff = updateReplicas(proxy, act)
	}

<<<<<<< HEAD
	act.Spec.Template.Spec.Containers[0].Image = fmt.Sprintf("%s:%s", imageName, proxy.Spec.Version)
	act.Spec.Template.Spec.Containers[0].Env[0].Value = strconv.FormatInt(int64(proxy.Spec.Port), 10)
	act.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort = proxy.Spec.Port

	act.Spec.Template.Spec.Containers[0].Resources = proxy.Spec.Resources
	act.Spec.Template.Spec.Containers[0].LivenessProbe = proxy.Spec.LivenessProbe
	act.Spec.Template.Spec.Containers[0].ReadinessProbe = proxy.Spec.ReadinessProbe
	act.Spec.Template.Spec.Containers[0].StartupProbe = proxy.Spec.StartupProbe
=======
	exp.Spec.Template, diff = updatePodTemplateSpec(proxy, act)
	return exp, diff
}
>>>>>>> refactor: update deployment reconcile

func updateReplicas(proxy *v1alpha1.ShardingSphereProxy, act *v1.Deployment) (*int32, bool) {
	if *act.Spec.Replicas != proxy.Spec.Replicas {
		return &proxy.Spec.Replicas, true
	}
	return act.Spec.Replicas, false
}

func updatePodTemplateSpec(proxy *v1alpha1.ShardingSphereProxy, act *v1.Deployment) (corev1.PodTemplateSpec, bool) {
	exp := act.Spec.Template.DeepCopy()

	SSProxyContainer, diff := updateSSProxyContainer(proxy, act)
	if diff {
		for i, _ := range exp.Spec.Containers {
			if exp.Spec.Containers[i].Name == "proxy" {
				exp.Spec.Containers[i] = *SSProxyContainer
			}
		}
	}

	initContainer, diff := updateInitContainer(proxy, act)
	if diff {
		for i, _ := range exp.Spec.InitContainers {
			if exp.Spec.InitContainers[i].Name == "download-mysql-connect" {
				exp.Spec.InitContainers[i] = *initContainer
			}
		}
	}

	configName, diff := updateConfigName(proxy, act)
	if diff {
		exp.Spec.Volumes[0].ConfigMap.Name = configName
	}

	return *exp, diff
}

func updateConfigName(proxy *v1alpha1.ShardingSphereProxy, act *v1.Deployment) (string, bool) {
	if act.Spec.Template.Spec.Volumes[0].ConfigMap.Name != proxy.Spec.ProxyConfigName {
		return proxy.Spec.ProxyConfigName, true
	}
	return act.Spec.Template.Spec.Volumes[0].ConfigMap.Name, false
}

func updateInitContainer(proxy *v1alpha1.ShardingSphereProxy, act *v1.Deployment) (*corev1.Container, bool) {
	var exp *corev1.Container
	var diff bool

	for _, c := range act.Spec.Template.Spec.InitContainers {
		if c.Name == "download-mysql-connect" {
			for _, env := range c.Env {
				if env.Name == "VERSION" {
					if env.Value != proxy.Spec.MySQLDriver.Version {
						diff = true
						env.Value = proxy.Spec.MySQLDriver.Version
					}
				}
			}
			exp = c.DeepCopy()
		}
	}

	return exp, diff
}

func updateSSProxyContainer(proxy *v1alpha1.ShardingSphereProxy, act *v1.Deployment) (*corev1.Container, bool) {
	var diff bool
	var exp *corev1.Container

	for _, c := range act.Spec.Template.Spec.Containers {
		if c.Name == "proxy" {
			exp = c.DeepCopy()

			tag := strings.Split(c.Image, ":")[1]
			if tag != proxy.Spec.Version {
				diff = true
				exp.Image = fmt.Sprintf("%s:%s", imageName, proxy.Spec.Version)
			}

			if proxy.Spec.Resources != nil && !reflect.DeepEqual(c.Resources, *proxy.Spec.Resources) {
				diff = true
				exp.Resources = *proxy.Spec.Resources
			}

			if proxy.Spec.LivenessProbe != nil && !reflect.DeepEqual(c.LivenessProbe, *proxy.Spec.LivenessProbe) {
				diff = true
				exp.LivenessProbe = proxy.Spec.LivenessProbe
			}

			if proxy.Spec.ReadinessProbe != nil && !reflect.DeepEqual(c.ReadinessProbe, *proxy.Spec.ReadinessProbe) {
				diff = true
				exp.ReadinessProbe = proxy.Spec.ReadinessProbe
			}

			if proxy.Spec.StartupProbe != nil && !reflect.DeepEqual(c.StartupProbe, *proxy.Spec.StartupProbe) {
				diff = true
				exp.StartupProbe = proxy.Spec.StartupProbe
			}

			for _, e := range c.Env {
				if e.Name == "PORT" {
					proxyPort := strconv.FormatInt(int64(proxy.Spec.Port), 10)
					if e.Value != proxyPort {
						diff = true
						e.Value = proxyPort
						exp.Ports[0].ContainerPort = proxy.Spec.Port
					}
				}
			}
		}
	}
	return exp, diff
}
