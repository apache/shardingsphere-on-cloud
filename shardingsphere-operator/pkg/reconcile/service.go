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
	"reflect"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func ComputeNodeNewService(cn *v1alpha1.ComputeNode) *v1.Service {
	svc := ComputeNodeDefaultService(cn.GetObjectMeta(), cn.GroupVersionKind())
	svc.Name = cn.Name
	svc.Namespace = cn.Namespace
	svc.Labels = cn.Labels
	// svc.Spec.Selector = cn.Labels
	svc.Spec.Selector = cn.Spec.Selector.MatchLabels
	// svc.Spec.Type = cn.Spec.Service.Type
	svc.Spec.Type = cn.Spec.ServiceType
	// svc.Spec.Ports = cn.Spec.Service.Ports

	if svc.Spec.Ports == nil {
		svc.Spec.Ports = []corev1.ServicePort{}
	}
	for _, pb := range cn.Spec.PortBindings {
		svc.Spec.Ports = append(svc.Spec.Ports, corev1.ServicePort{
			Name:       pb.Name,
			TargetPort: intstr.FromInt(int(pb.ContainerPort)),
			Port:       pb.ServicePort,
			NodePort:   pb.NodePort,
			Protocol:   pb.Protocol,
		})
	}

	return svc
}

func ComputeNodeDefaultService(meta metav1.Object, gvk schema.GroupVersionKind) *v1.Service {
	return &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "shardingsphere-proxy",
			Namespace: "default",
			Labels:    map[string]string{},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(meta, gvk),
			},
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{},
			Type:     v1.ServiceTypeClusterIP,
			Ports:    []v1.ServicePort{},
		},
	}
}

func ComputeNodeUpdateService(cn *v1alpha1.ComputeNode, cur *v1.Service) *v1.Service {
	exp := &v1.Service{}
	exp.ObjectMeta = cur.ObjectMeta
	exp.Labels = cur.Labels
	exp.Annotations = cur.Annotations
	exp.Spec = ComputeNodeNewService(cn).Spec
	exp.Spec.ClusterIP = cur.Spec.ClusterIP
	exp.Spec.ClusterIPs = cur.Spec.ClusterIPs
	return exp
}

func NewService(ssproxy *v1alpha1.ShardingSphereProxy) *v1.Service {
	return ConstructCascadingService(ssproxy)
}

func ConstructCascadingService(proxy *v1alpha1.ShardingSphereProxy) *v1.Service {
	if proxy == nil || reflect.DeepEqual(proxy, &v1alpha1.ShardingSphereProxy{}) {
		return &v1.Service{}
	}

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

func UpdateService(proxy *v1alpha1.ShardingSphereProxy, runtimeService *v1.Service) *v1.Service {
	exp := &v1.Service{}
	runtimeService.Spec.Type = proxy.Spec.ServiceType.Type
	runtimeService.Spec.Ports[0].Port = proxy.Spec.Port
	runtimeService.Spec.Ports[0].TargetPort = fromInt32(proxy.Spec.Port)
	if proxy.Spec.ServiceType.NodePort != 0 {
		runtimeService.Spec.Ports[0].NodePort = proxy.Spec.ServiceType.NodePort
	}
	exp = runtimeService.DeepCopy()
	return exp
}
