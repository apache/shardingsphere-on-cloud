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

package autoscaler

import (
	"context"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/hpa"

	autoscalingv2 "k8s.io/api/autoscaling/v2"
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// BuildHorizontalPodAutoScaler returns a new HPA
func (b builder) BuildHorizontalPodAutoScaler(ctx context.Context, meta metav1.ObjectMeta, gvk schema.GroupVersionKind, policy *v1alpha1.ScalingPolicy) *autoscalingv2beta2.HorizontalPodAutoscaler {
	blder := hpa.NewHorizontalPodAutoScalerBuilder()
	blder.SetName(meta.Name).SetNamespace(meta.Namespace).SetLabels(meta.Labels).SetAnnotations(meta.Annotations).SetOwnerReferences([]metav1.OwnerReference{
		*metav1.NewControllerRef(meta.GetObjectMeta(), gvk),
	})

	blder.SetScaleTargetRef(autoscalingv2.CrossVersionObjectReference{
		Kind:       "ComputeNode",
		Name:       policy.TargetSelector.ObjectRef.Name,
		APIVersion: "shardingsphere.apache.org/v1alpha1",
	})

	blder.SetMinReplicas(policy.Horizontal.MinReplicas)
	blder.SetMaxReplicas(policy.Horizontal.MaxReplicas)
	if policy.Horizontal.Metrics != nil {
		blder.SetMetrics(policy.Horizontal.Metrics)
	}

	var (
		up, down *autoscalingv2.HPAScalingRules
	)

	if policy.Horizontal.ScaleUpRules != nil {
		up = policy.Horizontal.ScaleUpRules
	}
	if policy.Horizontal.ScaleDownRules != nil {
		down = policy.Horizontal.ScaleDownRules
	}

	if up != nil || down != nil {
		blder.SetBehavior(&autoscalingv2.HorizontalPodAutoscalerBehavior{
			ScaleUp:   up,
			ScaleDown: down,
		})
	}

	return blder.BuildHPA()
}
