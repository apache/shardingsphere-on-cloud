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
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/vpa"
	v1 "k8s.io/api/autoscaling/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	autoscalingv1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
)

func (b builder) BuildVerticalPodAutoscaler(ctx context.Context, meta *metav1.ObjectMeta, gvk schema.GroupVersionKind, policy *v1alpha1.ScalingPolicy) *autoscalingv1.VerticalPodAutoscaler {
	blder := vpa.NewVerticalPodAutoscalerBuilder()
	blder.SetName(meta.Name).SetNamespace(meta.Namespace).SetLabels(meta.Labels).SetAnnotations(meta.Annotations).SetOwnerReferences([]metav1.OwnerReference{
		*metav1.NewControllerRef(meta.GetObjectMeta(), gvk),
	})

	blder.SetTargetRef(v1.CrossVersionObjectReference{
		Kind:       "ComputeNode",
		Name:       policy.TargetSelector.ObjectRef.Name,
		APIVersion: "shardingsphere.apache.org/v1alpha1",
	})

	if policy.Vertical.UpdatePolicy != nil {
		blder.SetUpdatePolicy(policy.Vertical.UpdatePolicy)
	}
	if policy.Vertical.ResourcePolicy != nil {
		blder.SetResourcePolicy(policy.Vertical.ResourcePolicy)
	}
	if policy.Vertical.Recommenders != nil {
		var r []*autoscalingv1.VerticalPodAutoscalerRecommenderSelector
		for _, recommender := range policy.Vertical.Recommenders {
			recommenderCopy := recommender
			r = append(r, &recommenderCopy)
		}
		blder.SetRecommenders(r)
	}

	return blder.BuildVPA()
}
