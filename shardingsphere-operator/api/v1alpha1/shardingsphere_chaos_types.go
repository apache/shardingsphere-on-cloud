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
	Spec              ShardingSphereChaosSpec   `json:"spec,omitempty"`
	Status            ShardingSphereChaosStatus `json:"status,omitempty"`
}

// ShardingSphereChaosSpec defines the desired state of ShardingSphereChaos
type ShardingSphereChaosSpec struct {
	InjectJob  JobSpec `json:"injectJob,omitempty"`
	EmbedChaos `json:",inline"`
}

// JobSpec Specifies the config of job to create
type JobSpec struct {
	// +optional
	Experimental string `json:"experimental,omitempty"`
	// +optional
	Pressure string `json:"pressure,omitempty"`
	// +optional
	Position string `json:"position,omitempty"`
}

type EmbedChaos struct {
	// +optional
	NetworkChaos *NetworkChaosSpec `json:"networkChaos,omitempty"`
	// +optional
	PodChaos *PodChaosSpec `json:"podChaos,omitempty"`
}

// ChaosCondition Show Chaos Progress
type ChaosCondition string

const (
	Creating     ChaosCondition = "Creating"
	AllRecovered ChaosCondition = "AllRecovered"
	Paused       ChaosCondition = "Paused"
	AllInjected  ChaosCondition = "AllInjected"
	NoTarget     ChaosCondition = "NoTarget"
	UnKnown      ChaosCondition = "UnKnown"
)

// ShardingSphereChaosStatus defines the actual state of ShardingSphereChaos
type ShardingSphereChaosStatus struct {
	ChaosCondition ChaosCondition `json:"chaosCondition"`
	Phase          Phase          `json:"phase"`
}

type Phase string

var (
	PhaseBeforeExperiment Phase = "before experiment"
	PhaseAfterExperiment  Phase = "after experiment"
	PhaseInChaos          Phase = "inject chaos"
	PhaseRecoveredChaos   Phase = "recover chaos"
)

// PodChaosAction Specify the action type of pod Chaos
type PodChaosAction string

var (
	PodFailureAction    PodChaosAction = "podFailure"
	ContainerKillAction PodChaosAction = "containerKill"
)

// PodChaosSpec Fields that need to be configured for pod type chaos
type PodChaosSpec struct {
	PodSelector `json:"selector,omitempty"`
	Action      PodChaosAction `json:"action"`
	//+optional
	PodActionParam PodActionParam `json:"params,omitempty"`
}

// PodActionParam Optional parameters for pod type configuration
type PodActionParam struct {
	//+optional
	PodFailure PodFailureActionParams `json:"podFailure,omitempty"`
	//+optional
	ContainerKill ContainerKillActionParams `json:"containerKill,omitempty"`
}

type PodFailureActionParams struct {
	// +optional
	Duration string `json:"duration,omitempty"`
}

type ContainerKillActionParams struct {
	// +optional
	ContainerNames []string `json:"containerNames,omitempty"`
}

// NetworkChaosSpec Fields that need to be configured for network type chaos
type NetworkChaosSpec struct {
	Source PodSelector `json:",inline"`
	// +optional
	Duration *string `json:"duration,omitempty"`
	//+optional
	Direction Direction `json:"direction,omitempty"`
	// +optional
	Target *PodSelector       `json:"target,omitempty"`
	Action NetworkChaosAction `json:"action"`
	// +optional
	Network *NetworkParams `json:"params,omitempty"`
}

// NetworkParams Optional parameters for network type configuration
type NetworkParams struct {
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

// NetworkChaosAction Specify the action type of network Chaos
type NetworkChaosAction string

const (
	DelayAction     NetworkChaosAction = "delay"
	LossAction      NetworkChaosAction = "loss"
	DuplicateAction NetworkChaosAction = "duplicate"
	CorruptAction   NetworkChaosAction = "corrupt"
	PartitionAction NetworkChaosAction = "partition"
)

// Direction Specifies the direction of action of network chaos
type Direction string

const (
	To   Direction = "to"
	From Direction = "from"
	Both Direction = "both"
)

// PodSelector Used to select the target of the specified chaos
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
	NodeSelectors       map[string]string                 `json:"nodeSelectors,omitempty"`
	ExpressionSelectors []metav1.LabelSelectorRequirement `json:"expressionSelectors,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ShardingSphereChaos{}, &ShardingSphereChaosList{})
}
