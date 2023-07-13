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

package vpa

import (
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/metadata"
	v1 "k8s.io/api/autoscaling/v1"
	autoscalingv1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
)

// VerticalPodAutoscalerBuilder is a builder for VPA
type VerticalPodAutoscalerBuilder interface {
	metadata.MetadataBuilder

	SetTargetRef(ref v1.CrossVersionObjectReference) VerticalPodAutoscalerBuilder
	SetUpdatePolicy(pu *autoscalingv1.PodUpdatePolicy) VerticalPodAutoscalerBuilder
	SetResourcePolicy(pr *autoscalingv1.PodResourcePolicy) VerticalPodAutoscalerBuilder
	SetRecommenders(rs []*autoscalingv1.VerticalPodAutoscalerRecommenderSelector) VerticalPodAutoscalerBuilder

	BuildVPA() *autoscalingv1.VerticalPodAutoscaler
}

// NewVerticalPodAutoscalerBuilder returns a VerticalPodAutoscalerBuilder for VPA
func NewVerticalPodAutoscalerBuilder() VerticalPodAutoscalerBuilder {
	return &vpaBuilder{}
}

type vpaBuilder struct {
	vpa *autoscalingv1.VerticalPodAutoscaler
	metadata.MetadataBuilder
}

// SetTargetRef set the scale target
func (v *vpaBuilder) SetTargetRef(ref v1.CrossVersionObjectReference) VerticalPodAutoscalerBuilder {
	v.vpa.Spec.TargetRef = &ref
	return v
}

// SetUpdatePolicy set the rules on how changes are applied to the pods
func (v *vpaBuilder) SetUpdatePolicy(pu *autoscalingv1.PodUpdatePolicy) VerticalPodAutoscalerBuilder {
	v.vpa.Spec.UpdatePolicy = pu
	return v
}

// SetResourcePolicy set how the autoscaler computes recommended resources
func (v *vpaBuilder) SetResourcePolicy(pr *autoscalingv1.PodResourcePolicy) VerticalPodAutoscalerBuilder {
	v.vpa.Spec.ResourcePolicy = pr
	return v
}

// SetRecommenders set the recommenders
func (v *vpaBuilder) SetRecommenders(rs []*autoscalingv1.VerticalPodAutoscalerRecommenderSelector) VerticalPodAutoscalerBuilder {
	v.vpa.Spec.Recommenders = rs
	return v
}

// BuildVPA returns a VPA
func (v *vpaBuilder) BuildVPA() *autoscalingv1.VerticalPodAutoscaler {
	v.vpa.ObjectMeta = *v.BuildMetadata()
	return v.vpa
}
