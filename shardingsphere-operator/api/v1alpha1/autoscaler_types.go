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

package v1alpha1

import (
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	vpav1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
)

// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name=Age,type=date
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type AutoScaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoScalerSpec   `json:"spec,omitempty"`
	Status            AutoScalerStatus `json:"status,omitempty"`
}

type AutoScalerSpec struct {
	PolicyGroup []ScalingPolicy `json:"policyGroup,omitempty" yaml:"policyGroup,omiempty"`
	Provider    string          `json:"provider,omitempty" yaml:"provider,omitempty"`
}

type ScalingPolicy struct {
	TargetSelector *ObjectRefSelector `json:"targetSelector,omitempty" yaml:"targetSelector,omitempty"`
	Horizontal     *HorizontalScaling `json:"horizontal,omitempty" yaml:"horizontal,omitempty"`
	Vertical       *VerticalScaling   `json:"vertical,omitempey" yaml:"vertical,omitempty"`
}

type ObjectRefSelector struct {
	ObjectRef autoscalingv2beta2.CrossVersionObjectReference `json:"objectRef,omitempty" yaml:"objectRef,omitempty"`
	Selector  *metav1.LabelSelector                          `json:"selector,omitempty" yaml:"selector,omitempty"`
}

type HorizontalScaling struct {
	MaxReplicas    int32                               `json:"maxReplicas,omitempty" yaml:"maxReplicas,omitempty"`
	MinReplicas    int32                               `json:"minReplicas,omitempty" yaml:"minReplicas,omitempty"`
	ScaleUpRules   *autoscalingv2beta2.HPAScalingRules `json:"scaleUpRules,omitempty" yaml:"scaleUpRules,omitempty"`
	ScaleDownRules *autoscalingv2beta2.HPAScalingRules `json:"scaleDownRules,omitempty" yaml:"scaleDownRules,omitempty"`
	Metrics        *autoscalingv2beta2.MetricSpec      `json:"metrics,omitempty" yaml:"metrics,omitempty"`
}

type VerticalScaling struct {
	UpdatePolicy   *vpav1.PodUpdatePolicy                           `json:"updatePolicy,omitempty" yaml:"updatePolicy,omitempty"`
	ResourcePolicy *vpav1.PodResourcePolicy                         `json:"resourcePolicy,omitempty" yaml:"resourcePolicy,omitempty"`
	Recommenders   []vpav1.VerticalPodAutoscalerRecommenderSelector `json:"recommenders,omitempty" yaml:"recommenders,omitempty"`
}

type AutoScalerStatus struct {
	Conditions []AutoScalerCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}

type AutoScalerCondition struct {
	Type               AutoScalerConditionType `json:"type" protobuf:"bytes,1,name=type"`
	Status             v1.ConditionStatus      `json:"status" protobuf:"bytes,2,name=status"`
	LastTransitionTime metav1.Time             `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
	Reason             string                  `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	Message            string                  `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}

type AutoScalerConditionType string

const (
	ScalingReady AutoScalerConditionType = "ScalingReady"
)
