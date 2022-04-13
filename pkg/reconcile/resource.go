package reconcile

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	sphereexcomv1alpha1 "sphere-ex.com/shardingsphere-operator/api/v1alpha1"
)

func ConstructCascadingDeployment(proxy *sphereexcomv1alpha1.Proxy) *appsv1.Deployment {
	if proxy.Spec.Port == 0 {
		proxy.Spec.Port = 3307
	}
	dp := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proxy.Name,
			Namespace: proxy.Namespace,
			Annotations: map[string]string{
				"UpdateTime": metav1.Now().Format(metav1.RFC3339Micro),
			},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(proxy.GetObjectMeta(), proxy.GroupVersionKind()),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RecreateDeploymentStrategyType,
			},
			Replicas: &proxy.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"apps": "proxy-" + proxy.Name,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"apps": "proxy-" + proxy.Name,
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            "proxy",
							Image:           fmt.Sprintf("apache/shardingsphere-proxy:%s", proxy.Spec.Version),
							ImagePullPolicy: v1.PullIfNotPresent,
							Ports: []v1.ContainerPort{
								{
									ContainerPort: proxy.Spec.Port,
								},
							},
							Env: []v1.EnvVar{
								{
									Name:  "PORT",
									Value: string(proxy.Spec.Port),
								},
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "config",
									MountPath: "/opt/shardingsphere-proxy/conf",
								},
								{
									Name:      "mysql-connect-jar",
									MountPath: "/opt/shardingsphere-proxy/ext-lib",
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
						{
							Name: "mysql-connect-jar",
							VolumeSource: v1.VolumeSource{
								EmptyDir: &v1.EmptyDirVolumeSource{},
							},
						},
					},
				},
			},
		},
	}
	return &dp
}

func ConstructCascadingService(proxy *sphereexcomv1alpha1.Proxy) *v1.Service {

	svc := v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proxy.Name + "-svc",
			Namespace: proxy.Namespace,
			Annotations: map[string]string{
				"UpdateTime": metav1.Now().Format(metav1.RFC3339Micro),
			},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(proxy.GetObjectMeta(), proxy.GroupVersionKind()),
			},
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"apps": "proxy-" + proxy.Name,
			},
			Type: proxy.Spec.ServiceType.Type,
			Ports: []v1.ServicePort{
				{
					Name: "proxy-port",
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: proxy.Spec.Port,
					},
					Port: proxy.Spec.Port,
				},
			},
		},
	}
	if proxy.Spec.ServiceType.Type == v1.ServiceTypeNodePort {
		svc.Spec.Ports[0].NodePort = proxy.Spec.ServiceType.NodePort
	}
	return &svc
}
