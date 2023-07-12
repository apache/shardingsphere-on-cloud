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
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	vpav1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
)

// +kubebuilder:object:root=true
// AutoScalerList contains a list of AutoScaler
type AutoScalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoScaler `json:"items"`
}

// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name=Age,type=date
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// AutoScaler defines a autoscaler for ShardingSphere given
type AutoScaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoScalerSpec   `json:"spec,omitempty"`
	Status            AutoScalerStatus `json:"status,omitempty"`
}

// AutoScalerSpec defines a list of scaling policies
type AutoScalerSpec struct {
	// PolicyGroup declares a set of scaling policies, including horizontal and vertical scaling
	// Allow ComputeNode to enable HPA and VPA at the same time, without guaranteeing the result
	// Haven't verified the configuration of HPA and VPA of StorageNode yet
	PolicyGroup []ScalingPolicy `json:"policyGroup,omitempty" yaml:"policyGroup,omiempty"`
}

// ScalingPolicy defines a policy for scaling
type ScalingPolicy struct {
	// TargetSelector is used to select the auto-scaling target
	// Support native CrossVersionObjectReference and Selector
	// The first version plans to support CrossVersionObjectReference first
	TargetSelector *ObjectRefSelector `json:"targetSelector,omitempty" yaml:"targetSelector,omitempty"`

	// Provider is the provider of the scaling mechanism, and the optional values are:
	// - Empty: default value, which means provided by ShardingSphere Operator
	// - KubernetesHPA: Indicates the use of Kubernetes native HPA
	// - KubernetesVPA: Indicates the use of Kubernetes community VPA
	// - Other: Indicates a controller using a third-party controller
	// +optional
	Provider string `json:"provider,omitempty" yaml:"provider,omitempty"`

	// HorizontalScaling contains the necessary parameters for HPA scaling
	// Does not contain StorageNode related configuration
	// +optional
	Horizontal *HorizontalScaling `json:"horizontal,omitempty" yaml:"horizontal,omitempty"`

	// VerticalScaling contains the necessary parameters for VPA scaling
	// Does not contain StorageNode related configuration
	// +optional
	Vertical *VerticalScaling `json:"vertical,omitempty" yaml:"vertical,omitempty"`
}

// ObjectRefSelector defines a selector for objects
type ObjectRefSelector struct {
	// +optional
	ObjectRef autoscalingv2.CrossVersionObjectReference `json:"objectRef,omitempty" yaml:"objectRef,omitempty"`
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty" yaml:"selector,omitempty"`
}

// The following configuration items are basically the same as HPA configuration,
// please refer to the corresponding documentation for descriptiontype
type HorizontalScaling struct {
	// maxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up.
	// It cannot be less that minReplicas.
	// +optional
	MaxReplicas int32 `json:"maxReplicas,omitempty" yaml:"maxReplicas,omitempty"`
	// minReplicas is the lower limit for the number of replicas to which the autoscaler
	// can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the
	// alpha feature gate HPAScaleToZero is enabled and at least one Object or External
	// metric is configured.  Scaling is active as long as at least one metric value is
	// available.
	// +optional
	MinReplicas int32 `json:"minReplicas,omitempty" yaml:"minReplicas,omitempty"`

	// scaleUp is scaling policy for scaling Up.
	// If not set, the default value is the higher of:
	//   * increase no more than 4 pods per 60 seconds
	//   * double the number of pods per 60 seconds
	// No stabilization is used.
	// +optional
	ScaleUpRules *autoscalingv2.HPAScalingRules `json:"scaleUpRules,omitempty" yaml:"scaleUpRules,omitempty"`
	// scaleDown is scaling policy for scaling Down.
	// If not set, the default value is to allow to scale down to minReplicas pods, with a
	// 300 second stabilization window (i.e., the highest recommendation for
	// the last 300sec is used).
	// +optional
	ScaleDownRules *autoscalingv2.HPAScalingRules `json:"scaleDownRules,omitempty" yaml:"scaleDownRules,omitempty"`

	// metrics contains the specifications for which to use to calculate the
	// desired replica count (the maximum replica count across all metrics will
	// be used).  The desired replica count is calculated multiplying the
	// ratio between the target value and the current value by the current
	// number of pods.  Ergo, metrics used must decrease as the pod count is
	// increased, and vice-versa.  See the individual metric source types for
	// more information about how each type of metric must respond.
	// If not set, the default metric will be set to 80% average CPU utilization.
	// +optional
	Metrics []autoscalingv2.MetricSpec `json:"metrics,omitempty" yaml:"metrics,omitempty"`
}

// The following configuration items are basically the same as the VPA configuration,
// please refer to the corresponding documentation for instructions
type VerticalScaling struct {
	// Describes the rules on how changes are applied to the pods.
	// If not specified, all fields in the `PodUpdatePolicy` are set to their
	// default values.
	// +optional
	UpdatePolicy *vpav1.PodUpdatePolicy `json:"updatePolicy,omitempty" yaml:"updatePolicy,omitempty"`

	// Controls how the autoscaler computes recommended resources.
	// The resource policy may be used to set constraints on the recommendations
	// for individual containers. If not specified, the autoscaler computes recommended
	// resources for all containers in the pod, without additional constraints.
	// +optional
	ResourcePolicy *vpav1.PodResourcePolicy `json:"resourcePolicy,omitempty" yaml:"resourcePolicy,omitempty"`

	// Recommender responsible for generating recommendation for this object.
	// List should be empty (then the default recommender will generate the
	// recommendation) or contain exactly one recommender.
	// +optional
	Recommenders []vpav1.VerticalPodAutoscalerRecommenderSelector `json:"recommenders,omitempty" yaml:"recommenders,omitempty"`
}

// AutoScalerStatus defines the status of a autoscaler
type AutoScalerStatus struct {
	// +optional
	Conditions []AutoScalerCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}

// AutoScalerCondition defiens the condition of a autoscaler
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

func init() {
	SchemeBuilder.Register(&AutoScaler{}, &AutoScalerList{})
}
