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
	"reflect"
	"testing"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func Test_ComputeNodeUpdateService(t *testing.T) {
	cases := []struct {
		id  int
		cn  *v1alpha1.ComputeNode
		cur *corev1.Service
		exp *corev1.Service
		msg string
	}{
		{
			id: 1,
			cn: &v1alpha1.ComputeNode{
				Spec: v1alpha1.ComputeNodeSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{},
					},
					ServiceType: corev1.ServiceTypeClusterIP,
					PortBindings: []v1alpha1.PortBinding{
						{
							Name:          "proxy",
							ContainerPort: 3307,
							ServicePort:   3307,
						},
					},
				},
			},
			cur: &corev1.Service{
				Spec: corev1.ServiceSpec{
					Selector: map[string]string{},
					Type:     corev1.ServiceTypeClusterIP,
					Ports: []corev1.ServicePort{
						{
							Name:       "proxy",
							TargetPort: intstr.FromInt(3307),
							Port:       3307,
						},
					},
				},
			},
			exp: &corev1.Service{
				Spec: corev1.ServiceSpec{
					Selector: map[string]string{},
					Type:     corev1.ServiceTypeClusterIP,
					Ports: []corev1.ServicePort{
						{
							Name:       "proxy",
							TargetPort: intstr.FromInt(3307),
							Port:       3307,
						},
					},
				},
			},
			msg: "update clusterip case",
		},
		{
			id: 2,
			cn: &v1alpha1.ComputeNode{
				Spec: v1alpha1.ComputeNodeSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{},
					},
					ServiceType: corev1.ServiceTypeNodePort,
					PortBindings: []v1alpha1.PortBinding{
						{
							Name:          "proxy",
							ContainerPort: 3307,
							ServicePort:   3307,
							NodePort:      30000,
						},
					},
				},
			},
			cur: &corev1.Service{
				Spec: corev1.ServiceSpec{
					Selector: map[string]string{},
					Type:     corev1.ServiceTypeNodePort,
					Ports: []corev1.ServicePort{
						{
							Name:       "proxy",
							TargetPort: intstr.FromInt(3307),
							Port:       3307,
							NodePort:   30000,
						},
					},
				},
			},
			exp: &corev1.Service{
				Spec: corev1.ServiceSpec{
					Selector: map[string]string{},
					Type:     corev1.ServiceTypeNodePort,
					Ports: []corev1.ServicePort{
						{
							Name:       "proxy",
							TargetPort: intstr.FromInt(3307),
							Port:       3307,
							NodePort:   30000,
						},
					},
				},
			},
			msg: "update nodeport case",
		},
		{
			id: 3,
			cn: &v1alpha1.ComputeNode{
				Spec: v1alpha1.ComputeNodeSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{},
					},
					ServiceType: corev1.ServiceTypeNodePort,
					PortBindings: []v1alpha1.PortBinding{
						{
							Name:          "proxy",
							ContainerPort: 3307,
							ServicePort:   3307,
						},
					},
				},
			},
			cur: &corev1.Service{
				Spec: corev1.ServiceSpec{
					Selector: map[string]string{},
					Type:     corev1.ServiceTypeClusterIP,
					Ports: []corev1.ServicePort{
						{
							Name:       "proxy",
							TargetPort: intstr.FromInt(3307),
							Port:       3307,
						},
					},
				},
			},
			exp: &corev1.Service{
				Spec: corev1.ServiceSpec{
					Selector: map[string]string{},
					Type:     corev1.ServiceTypeNodePort,
					Ports: []corev1.ServicePort{
						{
							Name:       "proxy",
							TargetPort: intstr.FromInt(3307),
							Port:       3307,
						},
					},
				},
			},
			msg: "update clusterip to nodeport case",
		},
		{
			id: 4,
			cn: &v1alpha1.ComputeNode{
				Spec: v1alpha1.ComputeNodeSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{},
					},
					ServiceType: corev1.ServiceTypeClusterIP,
					PortBindings: []v1alpha1.PortBinding{
						{
							Name:          "proxy",
							ContainerPort: 3307,
							ServicePort:   3307,
						},
					},
				},
			},
			cur: &corev1.Service{
				Spec: corev1.ServiceSpec{
					Selector: map[string]string{},
					Type:     corev1.ServiceTypeNodePort,
					Ports: []corev1.ServicePort{
						{
							Name:       "proxy",
							TargetPort: intstr.FromInt(3307),
							Port:       3307,
							NodePort:   30000,
						},
					},
				},
			},
			exp: &corev1.Service{
				Spec: corev1.ServiceSpec{
					Selector: map[string]string{},
					Type:     corev1.ServiceTypeClusterIP,
					Ports: []corev1.ServicePort{
						{
							Name:       "proxy",
							TargetPort: intstr.FromInt(3307),
							Port:       3307,
						},
					},
				},
			},
			msg: "update nodeport to clusterip case",
		},
	}

	for _, c := range cases {
		act := UpdateService(c.cn, c.cur)
		assert.Equal(t, act, c.exp, fmt.Sprintf("%d: %s\n", c.id, c.msg))
	}
}

func TestUpdateService(t *testing.T) {
	cn := &v1alpha1.ComputeNode{}
	cn.Spec.Selector = &metav1.LabelSelector{MatchLabels: map[string]string{"key": "val"}}

	cur := &corev1.Service{}
	// Setting up some test data for cur instance
	cur.ObjectMeta = metav1.ObjectMeta{Name: "test-service", Namespace: "test-ns"}
	cur.Labels = map[string]string{"key": "val"}
	cur.Annotations = map[string]string{"anno": "val"}
	cur.Spec = corev1.ServiceSpec{ClusterIP: "10.96.0.1", Ports: []corev1.ServicePort{{Name: "test-port", Port: 80}}}

	exp := &corev1.Service{}
	// Setting up some expected data for exp instance
	exp.ObjectMeta = cur.ObjectMeta
	exp.Labels = cur.Labels
	exp.Annotations = cur.Annotations
	exp.Spec = NewService(cn).Spec
	exp.Spec.Type = corev1.ServiceTypeNodePort
	exp.Spec.ClusterIP = cur.Spec.ClusterIP
	exp.Spec.ClusterIPs = cur.Spec.ClusterIPs
	exp.Spec.Ports = []corev1.ServicePort{{Name: "test-port", Port: 8080, TargetPort: intstr.FromInt(8080), NodePort: 30000}}

	// Make the necessary changes to expected and cur instances manually
	cn.Spec.ServiceType = corev1.ServiceTypeNodePort
	cn.Spec.PortBindings = []v1alpha1.PortBinding{{Name: "test-port", ContainerPort: 8080, ServicePort: 8080, NodePort: 30000}}
	cur.Spec.Ports[0] = corev1.ServicePort{Name: "test-port", Port: 80, NodePort: 30000}

	result := UpdateService(cn, cur)

	// Check if the updates in the result match the expected result
	if !reflect.DeepEqual(result, exp) {
		t.Errorf("UpdateService() returned:\n %+v\n, expected:\n %+v\n", result, exp)
	}
}
