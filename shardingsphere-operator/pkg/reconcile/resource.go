/*
 *   Copyright © 2022，SphereEx Authors
 *   All Rights Reserved.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package reconcile

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sphere-ex.com/shardingsphere-operator/api/v1alpha1"
	"strconv"
	"strings"
)

const imageName = "apache/shardingsphere-proxy"

var logback = `<?xml version="1.0"?>
<configuration>
    <appender name="console" class="ch.qos.logback.core.ConsoleAppender">
        <encoder>
            <pattern>[%-5level] %d{yyyy-MM-dd HH:mm:ss.SSS} [%thread] %logger{36} - %msg%n</pattern>
        </encoder>
    </appender>
    <appender name="sqlConsole" class="ch.qos.logback.core.ConsoleAppender">
        <encoder>
            <pattern>[%-5level] %d{yyyy-MM-dd HH:mm:ss.SSS} [%thread] [%X{database}] [%X{user}] [%X{host}] %logger{36} - %msg%n</pattern>
        </encoder>
    </appender>
    
    <logger name="ShardingSphere-SQL" level="info" additivity="false">
        <appender-ref ref="sqlConsole" />
    </logger>
    <logger name="org.apache.shardingsphere" level="info" additivity="false">
        <appender-ref ref="console" />
    </logger>
    
    <logger name="com.zaxxer.hikari" level="error" />
    
    <logger name="com.atomikos" level="error" />
    
    <logger name="io.netty" level="error" />
    
    <root>
        <level value="info" />
        <appender-ref ref="console" />
    </root>
</configuration> 
`

func ConstructCascadingDeployment(proxy *v1alpha1.Proxy) *appsv1.Deployment {
	dp := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proxy.Name,
			Namespace: proxy.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(proxy.GetObjectMeta(), proxy.GroupVersionKind()),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RecreateDeploymentStrategyType,
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"apps": proxy.Name,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"apps": proxy.Name,
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            "proxy",
							Image:           fmt.Sprintf("%s:%s", imageName, proxy.Spec.Version),
							ImagePullPolicy: v1.PullIfNotPresent,
							Ports: []v1.ContainerPort{
								{
									ContainerPort: proxy.Spec.Port,
								},
							},
							Env: []v1.EnvVar{
								{
									Name:  "PORT",
									Value: strconv.FormatInt(int64(proxy.Spec.Port), 10),
								},
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
	dp.Spec.Template.Spec.Containers[0].Resources = *proxy.Spec.Resources
	dp.Spec.Template.Spec.Containers[0].LivenessProbe = proxy.Spec.LivenessProbe
	dp.Spec.Template.Spec.Containers[0].ReadinessProbe = proxy.Spec.ReadinessProbe
	dp.Spec.Template.Spec.Containers[0].StartupProbe = proxy.Spec.StartupProbe
	if len(proxy.Spec.ImagePullSecrets) > 0 {
		dp.Spec.Template.Spec.ImagePullSecrets = proxy.Spec.ImagePullSecrets
	}
	return processOptionalParameter(proxy, dp)
}

func ConstructCascadingService(proxy *v1alpha1.Proxy) *v1.Service {

	svc := v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proxy.Name,
			Namespace: proxy.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(proxy.GetObjectMeta(), proxy.GroupVersionKind()),
			},
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"apps": proxy.Name,
			},
			Type: proxy.Spec.ServiceType.Type,
			Ports: []v1.ServicePort{
				{
					Name:       "proxy-port",
					TargetPort: fromInt32(proxy.Spec.Port),
					Port:       proxy.Spec.Port,
				},
			},
		},
	}
	if proxy.Spec.ServiceType.Type == v1.ServiceTypeNodePort {
		svc.Spec.Ports[0].NodePort = proxy.Spec.ServiceType.NodePort
	}
	return &svc
}

func addInitContainer(dp *appsv1.Deployment, mysql *v1alpha1.MySQLDriver) {

	if len(dp.Spec.Template.Spec.InitContainers) == 0 {
		dp.Spec.Template.Spec.Containers[0].VolumeMounts = append(dp.Spec.Template.Spec.Containers[0].VolumeMounts, v1.VolumeMount{
			Name:      "mysql-connect-jar",
			MountPath: "/opt/shardingsphere-proxy/ext-lib",
		},
		)

		dp.Spec.Template.Spec.Volumes = append(dp.Spec.Template.Spec.Volumes, v1.Volume{
			Name: "mysql-connect-jar",
			VolumeSource: v1.VolumeSource{
				EmptyDir: &v1.EmptyDirVolumeSource{},
			},
		})
	}

	scriptStr := strings.Builder{}
	t1, _ := template.New("shell").Parse(`wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/{{ .Version }}/mysql-connector-java-{{ .Version }}.jar;
wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/{{ .Version }}/mysql-connector-java-{{ .Version }}.jar.md5;
if [ $(md5sum /mysql-connector-java-{{ .Version }}.jar | cut -d ' ' -f1) = $(cat /mysql-connector-java-{{ .Version }}.jar.md5) ];
then echo success;
else echo failed;exit 1;fi;mv /mysql-connector-java-{{ .Version }}.jar /opt/shardingsphere-proxy/ext-lib`)
	_ = t1.Execute(&scriptStr, mysql)
	dp.Spec.Template.Spec.InitContainers = []v1.Container{
		{
			Name:    "download-mysql-connect",
			Image:   "busybox:1.35.0",
			Command: []string{"/bin/sh", "-c", scriptStr.String()},
			VolumeMounts: []v1.VolumeMount{
				{
					Name:      "mysql-connect-jar",
					MountPath: "/opt/shardingsphere-proxy/ext-lib",
				},
			},
		},
	}

}

func processOptionalParameter(proxy *v1alpha1.Proxy, dp *appsv1.Deployment) *appsv1.Deployment {
	if proxy.Spec.MySQLDriver != nil {
		addInitContainer(dp, proxy.Spec.MySQLDriver)
	}
	return dp
}

// ConstructCascadingConfigmap Construct spec resources to Configmap
func ConstructCascadingConfigmap(proxyConfig *v1alpha1.ProxyConfig) *v1.ConfigMap {
	y := toYaml(proxyConfig)
	return &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proxyConfig.Name,
			Namespace: proxyConfig.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(proxyConfig.GetObjectMeta(), proxyConfig.GroupVersionKind()),
			},
		},
		Data: map[string]string{
			"server.yaml": y,
			"logback.xml": logback,
		},
	}

}

// ConstructHPA Create HPA if you need
func ConstructHPA(proxy *v1alpha1.Proxy) *autoscalingv2beta2.HorizontalPodAutoscaler {
	return &autoscalingv2beta2.HorizontalPodAutoscaler{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proxy.Name,
			Namespace: proxy.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(proxy.GetObjectMeta(), proxy.GroupVersionKind()),
			},
		},
		Spec: autoscalingv2beta2.HorizontalPodAutoscalerSpec{
			ScaleTargetRef: autoscalingv2beta2.CrossVersionObjectReference{
				Kind:       "Deployment",
				Name:       proxy.Name,
				APIVersion: appsv1.SchemeGroupVersion.String(),
			},
			MinReplicas: &proxy.Spec.AutomaticScaling.MinInstance,
			MaxReplicas: proxy.Spec.AutomaticScaling.MaxInstance,
			Metrics: []autoscalingv2beta2.MetricSpec{
				{
					Type: autoscalingv2beta2.ResourceMetricSourceType,
					Resource: &autoscalingv2beta2.ResourceMetricSource{
						Name: "cpu",
						Target: autoscalingv2beta2.MetricTarget{
							Type:               autoscalingv2beta2.UtilizationMetricType,
							AverageUtilization: &proxy.Spec.AutomaticScaling.Target,
						},
					},
				},
			},
			Behavior: &autoscalingv2beta2.HorizontalPodAutoscalerBehavior{
				ScaleUp: &autoscalingv2beta2.HPAScalingRules{
					StabilizationWindowSeconds: &proxy.Spec.AutomaticScaling.ScaleUpWindows,
				},
				ScaleDown: &autoscalingv2beta2.HPAScalingRules{
					StabilizationWindowSeconds: &proxy.Spec.AutomaticScaling.ScaleDownWindows,
					Policies: []autoscalingv2beta2.HPAScalingPolicy{
						{
							Type:          autoscalingv2beta2.PodsScalingPolicy,
							Value:         1,
							PeriodSeconds: 30,
						},
					},
				},
			},
		},
	}

}

// ToYaml Convert ProxyConfig spec content to yaml format
func toYaml(proxyConfig *v1alpha1.ProxyConfig) string {
	y, _ := yaml.Marshal(proxyConfig.Spec)
	return string(y)
}

// UpdateDeployment FIXME:merge UpdateDeployment and ConstructCascadingDeployment
func UpdateDeployment(proxy *v1alpha1.Proxy, runtimeDeployment *appsv1.Deployment) {
	runtimeDeployment.Spec.Template.Spec.Containers[0].Image = fmt.Sprintf("%s:%s", imageName, proxy.Spec.Version)
	if proxy.Spec.AutomaticScaling == nil {
		runtimeDeployment.Spec.Replicas = &proxy.Spec.Replicas
	}
	runtimeDeployment.Spec.Template.Spec.Volumes[0].ConfigMap.Name = proxy.Spec.ProxyConfigName
	runtimeDeployment.Spec.Template.Spec.Containers[0].Env[0].Value = strconv.FormatInt(int64(proxy.Spec.Port), 10)
	runtimeDeployment.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort = proxy.Spec.Port
	if proxy.Spec.MySQLDriver.Version != "" {
		addInitContainer(runtimeDeployment, proxy.Spec.MySQLDriver)
	}
	runtimeDeployment.Spec.Template.Spec.Containers[0].Resources = *proxy.Spec.Resources
	runtimeDeployment.Spec.Template.Spec.Containers[0].LivenessProbe = proxy.Spec.LivenessProbe
	runtimeDeployment.Spec.Template.Spec.Containers[0].ReadinessProbe = proxy.Spec.ReadinessProbe
	runtimeDeployment.Spec.Template.Spec.Containers[0].StartupProbe = proxy.Spec.StartupProbe

}

func UpdateService(proxy *v1alpha1.Proxy, runtimeService *v1.Service) {
	runtimeService.Spec.Type = proxy.Spec.ServiceType.Type
	runtimeService.Spec.Ports[0].Port = proxy.Spec.Port
	runtimeService.Spec.Ports[0].TargetPort = fromInt32(proxy.Spec.Port)
	if proxy.Spec.ServiceType.NodePort != 0 {
		runtimeService.Spec.Ports[0].NodePort = proxy.Spec.ServiceType.NodePort
	}
}

func UpdateHPA(proxy *v1alpha1.Proxy, runtimeHPA *autoscalingv2beta2.HorizontalPodAutoscaler) {
	runtimeHPA.Spec.Metrics[0].Resource.Target.AverageUtilization = &proxy.Spec.AutomaticScaling.Target
	runtimeHPA.Spec.Behavior.ScaleDown.StabilizationWindowSeconds = &proxy.Spec.AutomaticScaling.ScaleDownWindows
	runtimeHPA.Spec.Behavior.ScaleUp.StabilizationWindowSeconds = &proxy.Spec.AutomaticScaling.ScaleUpWindows
	runtimeHPA.Spec.MaxReplicas = proxy.Spec.AutomaticScaling.MaxInstance
	runtimeHPA.Spec.MinReplicas = &proxy.Spec.AutomaticScaling.MinInstance
}

func fromInt32(val int32) intstr.IntOrString {
	return intstr.IntOrString{Type: intstr.Int, IntVal: val}
}
