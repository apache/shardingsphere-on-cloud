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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// ShardingSphereChaosList contains a list of ShardingSphereChaos
type ShardingSphereChaosList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShardingSphereChaos `json:"items"`
}

// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name=Age,type=date
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// ShardingSphereChaos defines a chaos test case for the ShardingSphere Proxy cluster
type ShardingSphereChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ShardingSphereChaosSpec `json:"spec,omitempty"`

	Status ShardingSphereChaosStatus `json:"status,omitempty"`
}

// ShardingSphereChaosSpec defines the desired state of ShardingSphereChaos
type ShardingSphereChaosSpec struct {
	//InjectJob batchV1Beta1.JobTemplateSpec `json:"InjectJob,omitempty"`

	ChaosKind ChaosKind `json:"chaosKind,omitempty"`

	EmbedChaos `json:",inline"`

	//Verify batchV1Beta1.JobTemplateSpec `json:"Verify,omitempty"`
}

type ChaosKind string

const (
	NetworkChaosKind ChaosKind = "networkChaos"

	PodChaosKind ChaosKind = "podChaos"
)

type EmbedChaos struct {
	// +optional
	NetworkChaos *NetworkChaosSpec `json:"networkChaos,omitempty"`
	// +optional
	PodChaos *PodChaosSpec `json:"podChaos,omitempty"`
}

type DeploymentCondition string

const (
	Creating     DeploymentCondition = "Creating"
	AllRecovered DeploymentCondition = "AllRecovered"
	Paused       DeploymentCondition = "Paused"
	AllInjected  DeploymentCondition = "AllInjected"
)

type Jobschedule string

const (
	JobCreating Jobschedule = "JobCreating"
	JobFailed   Jobschedule = "JobFailed"
	JobFinish   Jobschedule = "JobFinish"
)

// ShardingSphereChaosStatus defines the actual state of ShardingSphereChaos
type ShardingSphereChaosStatus struct {
	ChaosCondition DeploymentCondition `json:"deploymentCondition"`
	InjectStatus   Jobschedule         `json:"InjectStatus"`
	VerifyStatus   Jobschedule         `json:"VerifyStatus"`
}

// pod chaos

type PodChaosAction string

var (
	PodFailureAction    PodChaosAction = "podFailure"
	ContainerKillAction PodChaosAction = "ContainerKill"
)

type PodChaosSpec struct {
	PodSelector `json:",inline"`
	Action      PodChaosAction `json:"action"`
	// +optional
	PodActionParam *PodActionParam `json:",inline"`
}

type PodActionParam struct {
	//+optional
	PodFailure *PodFailureActionParams `json:"podFailure,omitempty"`
	//+optional
	ContainerKill *ContainerKillActionParams `json:"containerKill,omitempty"`
}

type PodFailureActionParams struct {
	// +optional
	Duration string `json:"duration,omitempty"`
}

type ContainerKillActionParams struct {
	// +optional
	ContainerNames []string `json:"containerNames,omitempty"`
}

//network chaos

type NetworkChaosSpec struct {
	Source PodSelector `json:",inline"`

	// +optional
	Duration *string `json:"duration,omitempty"`

	//+optional
	Direction Direction `json:"direction,omitempty"`

	// +optional
	Target *PodSelector `json:"target,omitempty"`

	Action NetworkChaosAction `json:"action"`

	// +optional
	NetWorkParams *NetWorkParams `json:",inline"`
}

type NetWorkParams struct {
	// +optional
	Delay *DelayActionParams `json:"delay,omitempty"`
	// +optional
	Loss *LossActionParams `json:"loss,omitempty"`
	// +optional
	Duplicate *DuplicateActionParams `json:"duplicate,omitempty"`
	// +optional
	Corrupt *CorruptActionParams `json:"corrupt,omitempty"`
}

type DelayActionParams struct {
	// +optional
	Latency string `json:"latency,omitempty"`
	// +optional
	Correlation string `json:"correlation,omitempty"`
	// +optional
	Jitter string `json:"jitter,omitempty"`
}

type LossActionParams struct {
	// +optional
	Loss string `json:"loss,omitempty"`
	// +optional
	Correlation string `json:"correlation,omitempty"`
}

type DuplicateActionParams struct {
	// +optional
	Duplicate string `json:"duplicate,omitempty"`
	// +optional
	Correlation string `json:"correlation,omitempty"`
}

type CorruptActionParams struct {
	// +optional
	Corrupt string `json:"corrupt,omitempty"`
	// +optional
	Correlation string `json:"correlation,omitempty"`
}

type NetworkChaosAction string

const (
	DelayAction NetworkChaosAction = "delay"

	LossAction NetworkChaosAction = "loss"

	DuplicateAction NetworkChaosAction = "duplicate"

	CorruptAction NetworkChaosAction = "corrupt"

	PartitionAction NetworkChaosAction = "partition"
)

type Direction string

const (
	To Direction = "to"

	From Direction = "from"

	Both Direction = "both"
)

//selector

type PodSelector struct {
	// +optional
	Namespaces []string `json:"namespaces,omitempty"`

	// +optional
	LabelSelectors map[string]string `json:"labelSelectors,omitempty"`

	// +optional
	AnnotationSelectors map[string]string `json:"annotationSelectors,omitempty"`

	// +optional
	Nodes []string `json:"nodes,omitempty"`

	// +optional
	Pods map[string][]string `json:"pods,omitempty"`

	// +optional
	NodeSelectors map[string]string `json:"nodeSelectors,omitempty"`

	ExpressionSelectors []metav1.LabelSelectorRequirement `json:"expressionSelectors,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ShardingSphereChaos{}, &ShardingSphereChaosList{})
}
