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

package hpa

import (
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/metadata"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
)

// HorizontalPodAutoscalerBuilder is a builder for HPA
type HorizontalPodAutoscalerBuilder interface {
	metadata.MetadataBuilder

	SetScaleTargetRef(ref autoscalingv2.CrossVersionObjectReference) HorizontalPodAutoscalerBuilder
	SetMinReplicas(n int32) HorizontalPodAutoscalerBuilder
	SetMaxReplicas(n int32) HorizontalPodAutoscalerBuilder
	SetMetrics(specs []autoscalingv2.MetricSpec) HorizontalPodAutoscalerBuilder
	SetBehavior(bh *autoscalingv2.HorizontalPodAutoscalerBehavior) HorizontalPodAutoscalerBuilder

	BuildHPA() *autoscalingv2.HorizontalPodAutoscaler
}

// NewHorizontalPodAutoScalerBuilder returns a HorizontalPodAutoScalerBuilder for HPA
func NewHorizontalPodAutoScalerBuilder() HorizontalPodAutoscalerBuilder {
	return &hpaBuilder{}
}

type hpaBuilder struct {
	hpa *autoscalingv2.HorizontalPodAutoscaler
	metadata.MetadataBuilder
}

// SetScaleTargetRef set the scale target
func (b *hpaBuilder) SetScaleTargetRef(ref autoscalingv2.CrossVersionObjectReference) HorizontalPodAutoscalerBuilder {
	b.hpa.Spec.ScaleTargetRef = ref
	return b
}

// SetMinReplicas set the min replicas
func (b *hpaBuilder) SetMinReplicas(n int32) HorizontalPodAutoscalerBuilder {
	b.hpa.Spec.MinReplicas = &n
	return b
}

// SetMaxReplicas set the max replicas
func (b *hpaBuilder) SetMaxReplicas(n int32) HorizontalPodAutoscalerBuilder {
	b.hpa.Spec.MaxReplicas = n
	return b
}

// SetMetrics set the metrics
func (b *hpaBuilder) SetMetrics(specs []autoscalingv2.MetricSpec) HorizontalPodAutoscalerBuilder {
	b.hpa.Spec.Metrics = specs
	return b
}

// SetBehavior set the behavior
func (b *hpaBuilder) SetBehavior(bh *autoscalingv2.HorizontalPodAutoscalerBehavior) HorizontalPodAutoscalerBuilder {
	b.hpa.Spec.Behavior = bh
	return b
}

// BuildHPA returns a HPA
func (b *hpaBuilder) BuildHPA() *autoscalingv2.HorizontalPodAutoscaler {
	b.hpa.ObjectMeta = *b.BuildMetadata()
	return b.hpa
}
