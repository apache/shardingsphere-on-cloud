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
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
)

type HorizontalPodAutoscalerBuilder interface {
	metadata.MetadataBuilder

	SetScaleTargetRef(ref autoscalingv2beta2.CrossVersionObjectReference) HorizontalPodAutoscalerBuilder
	SetMinReplicas(n int32) HorizontalPodAutoscalerBuilder
	SetMaxReplicas(n int32) HorizontalPodAutoscalerBuilder
	SetMetrics(specs []autoscalingv2beta2.MetricSpec) HorizontalPodAutoscalerBuilder
	SetBehavior(bh *autoscalingv2beta2.HorizontalPodAutoscalerBehavior) HorizontalPodAutoscalerBuilder

	BuildHPA() *autoscalingv2beta2.HorizontalPodAutoscaler
}

func NewHorizontalPodAutoScalerBuilder() HorizontalPodAutoscalerBuilder {
	return &hpaBuilder{
		hpa:             &autoscalingv2beta2.HorizontalPodAutoscaler{},
		MetadataBuilder: metadata.NewMetadataBuilder(),
	}
}

type hpaBuilder struct {
	hpa *autoscalingv2beta2.HorizontalPodAutoscaler
	metadata.MetadataBuilder
}

func (b *hpaBuilder) SetScaleTargetRef(ref autoscalingv2beta2.CrossVersionObjectReference) HorizontalPodAutoscalerBuilder {
	b.hpa.Spec.ScaleTargetRef = ref
	return b
}

func (b *hpaBuilder) SetMinReplicas(n int32) HorizontalPodAutoscalerBuilder {
	b.hpa.Spec.MinReplicas = &n
	return b
}

func (b *hpaBuilder) SetMaxReplicas(n int32) HorizontalPodAutoscalerBuilder {
	b.hpa.Spec.MaxReplicas = n
	return b
}

func (b *hpaBuilder) SetMetrics(specs []autoscalingv2beta2.MetricSpec) HorizontalPodAutoscalerBuilder {
	b.hpa.Spec.Metrics = specs
	return b
}

func (b *hpaBuilder) SetBehavior(bh *autoscalingv2beta2.HorizontalPodAutoscalerBehavior) HorizontalPodAutoscalerBuilder {
	b.hpa.Spec.Behavior = bh
	return b
}

func (b *hpaBuilder) BuildHPA() *autoscalingv2beta2.HorizontalPodAutoscaler {
	return b.hpa
}
