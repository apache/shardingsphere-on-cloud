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

package proxy

import (
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewHPA returns a new HorizontalPodAutoscaler
func NewHPA(ssproxy *v1alpha1.ShardingSphereProxy) *autoscalingv2beta2.HorizontalPodAutoscaler {
	return ConstructHPA(ssproxy)
}

// ConstructHPA Create HPA if you need
func ConstructHPA(proxy *v1alpha1.ShardingSphereProxy) *autoscalingv2beta2.HorizontalPodAutoscaler {
	var metrics = ConstructDefaultHPAMetric(&proxy.Spec.AutomaticScaling.Target)

	if len(proxy.Spec.AutomaticScaling.CustomMetrics) > 0 {
		metrics = proxy.Spec.AutomaticScaling.CustomMetrics
	}

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
			Metrics:     metrics,
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

// ConstructDefaultHPAMetric using CPU as default HPA metrics
func ConstructDefaultHPAMetric(target *int32) []autoscalingv2beta2.MetricSpec {
	return []autoscalingv2beta2.MetricSpec{
		{
			Type: autoscalingv2beta2.ResourceMetricSourceType,
			Resource: &autoscalingv2beta2.ResourceMetricSource{
				Name: "cpu",
				Target: autoscalingv2beta2.MetricTarget{
					Type:               autoscalingv2beta2.UtilizationMetricType,
					AverageUtilization: target,
				},
			},
		},
	}
}

// UpdateHPA update HorizontalPodAutoscaler with ShardingSphereProxy
func UpdateHPA(proxy *v1alpha1.ShardingSphereProxy, act *autoscalingv2beta2.HorizontalPodAutoscaler) *autoscalingv2beta2.HorizontalPodAutoscaler {
	act.Spec.Behavior.ScaleDown.StabilizationWindowSeconds = &proxy.Spec.AutomaticScaling.ScaleDownWindows
	act.Spec.Behavior.ScaleUp.StabilizationWindowSeconds = &proxy.Spec.AutomaticScaling.ScaleUpWindows
	act.Spec.MaxReplicas = proxy.Spec.AutomaticScaling.MaxInstance
	act.Spec.MinReplicas = &proxy.Spec.AutomaticScaling.MinInstance
	if len(proxy.Spec.AutomaticScaling.CustomMetrics) > 0 {
		act.Spec.Metrics = proxy.Spec.AutomaticScaling.CustomMetrics
	} else {
		// We need to reconstruct the default hpa metric when the user deletes the custom metric.
		act.Spec.Metrics = ConstructDefaultHPAMetric(&proxy.Spec.AutomaticScaling.Target)
	}

	exp := act.DeepCopy()
	return exp
}
