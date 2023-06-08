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
// ChaosList contains a list of Chaos
type ChaosList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Chaos `json:"items"`
}

// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name=Age,type=date
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// Chaos defines a chaos test case for the ShardingSphere Proxy cluster
type Chaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ChaosSpec   `json:"spec,omitempty"`
	Status            ChaosStatus `json:"status,omitempty"`
}

// ChaosSpec defines the desired state of Chaos
type ChaosSpec struct {
	EmbedChaos `json:",inline"`

	// +optional
	InjectJob *JobSpec `json:"injectJob,omitempty" yaml:"injectJob,omitempty"`
	// +optional
	PressureCfg *PressureCfg `json:"pressureCfg,omitempty" yaml:"pressureCfg,omitempty"`
}

type PressureCfg struct {
	ZkHost        string          `json:"zkHost,omitempty"`
	SsHost        string          `json:"ssHost"`
	Duration      metav1.Duration `json:"duration"`
	ReqTime       metav1.Duration `json:"reqTime"`
	DistSQLs      []DistSQL       `json:"distSQLs,omitempty"`
	ConcurrentNum int             `json:"concurrentNum"`
	ReqNum        int             `json:"reqNum"`
}

type DistSQL struct {
	SQL  string   `json:"sql"`
	Args []string `json:"args,omitempty"`
}

type Script string

// JobSpec specifies the config of job to create
type JobSpec struct {
	// +optional
	Experimental Script `json:"experimental,omitempty" yaml:"experimental,omitempty"`
	// +optional
	Pressure Script `json:"pressure,omitempty" yaml:"pressure,omitempty"`
	// +optional
	Verify Script `json:"verify,omitempty" yaml:"verify,omitempty"`
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
	AllRecovered ChaosCondition = "AllRecovered"
	Paused       ChaosCondition = "Paused"
	AllInjected  ChaosCondition = "AllInjected"
	NoTarget     ChaosCondition = "NoTarget"
	Unknown      ChaosCondition = "Unknown"
)

// ChaosStatus defines the actual state of Chaos
type ChaosStatus struct {
	// +optional
	ChaosCondition ChaosCondition `json:"chaosCondition,omitempty" yaml:"chaosCondition,omitempty"`
	// +optional
	Phase ChaosPhase `json:"phase,omitempty" yaml:"phase,omitempty"`
	// +optional
	// Result Result `json:"result,omitempty" yaml:"result,omitempty"`
	// +optional
	Conditions []*metav1.Condition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}

// Result represents the result of the Chaos
type Result struct {
	Steady Msg `json:"steady"`
	Chaos  Msg `json:"chaos"`
}

type Metrics string

type Msg struct {
	Metrics        Metrics `json:"metrics"`
	Result         string  `json:"result"`
	Duration       string  `json:"duration"`
	FailureDetails string  `json:"failureDetails"`
}

type ChaosPhase string

var (
	BeforeSteady ChaosPhase = "BeforeSteady"
	AfterSteady  ChaosPhase = "AfterSteady"
	BeforeChaos  ChaosPhase = "BeforeChaos"
	AfterChaos   ChaosPhase = "AfterChaos"
)

// PodChaosAction Specify the action type of pod Chaos
type PodChaosAction string

var (
	PodFailure    PodChaosAction = "PodFailure"
	ContainerKill PodChaosAction = "ContainerKill"
	PodKill       PodChaosAction = "PodKill"
	CPUStress     PodChaosAction = "CPUStress"
	MemoryStress  PodChaosAction = "MemoryStress"
)

// PodChaosSpec Fields that need to be configured for pod type chaos
type PodChaosSpec struct {
	PodSelector `json:"selector,omitempty"`
	Action      PodChaosAction `json:"action"`

	// +optional
	Params PodChaosParams `json:"params,omitempty"`
}

// PodActionParams Optional parameters for pod type configuration
type PodChaosParams struct {
	// +optional
	PodFailure *PodFailureParams `json:"podFailure,omitempty" yaml:"PodFailure,omitempty"`
	// +optional
	ContainerKill *ContainerKillParams `json:"containerKill,omitempty" yaml:"containerKill,omitempty"`
	//+optional
	CPUStress *CPUStressParams `json:"cpuStress,omitempty" yaml:"cpuStress,omitempty"`
	//+optional
	MemoryStress *MemoryStressParams `json:"memoryStress,omitempty" yaml:"memoryStress,omitempty"`
	// +optional
	PodKill *PodKillParams `json:"podKill,omitempty" yaml:"podKill,omitempty"`
}

type PodFailureParams struct {
	// +optional
	Duration *string `json:"duration,omitempty"`
}

type ContainerKillParams struct {
	// +optional
	ContainerNames []string `json:"containerNames,omitempty"`
}

type CPUStressParams struct {
	Duration string `json:"duration"`
	//+optional
	Cores int `json:"cores,omitempty"`
	//+optional
	Load int `json:"load,omitempty"`
}

type MemoryStressParams struct {
	Duration string `json:"duration"`
	//+optional
	Workers int `json:"workers,omitempty"`
	//+optional
	Consumption string `json:"consumption,omitempty"`
}

type PodKillParams struct {
	// +optional
	GracePeriod int64 `json:"gracePeriod,omitempty"`
}

// NetworkChaosSpec Fields that need to be configured for network type chaos
type NetworkChaosSpec struct {
	Source PodSelector  `json:"source,omitempty"`
	Target *PodSelector `json:"target,omitempty"`

	// +optional
	Action NetworkChaosAction `json:"action"`
	// +optional
	Duration *string `json:"duration,omitempty"`
	// +optional
	Direction Direction `json:"direction,omitempty"`
	// +optional
	Params NetworkChaosParams `json:"params,omitempty"`
}

// NetworkParams Optional parameters for network type configuration
type NetworkChaosParams struct {
	// +optional
	Delay *DelayParams `json:"delay,omitempty"`
	// +optional
	Loss *LossParams `json:"loss,omitempty"`
	// +optional
	Duplication *DuplicationParams `json:"duplicate,omitempty"`
	// +optional
	Corruption *CorruptionParams `json:"corrupt,omitempty"`
}

type DelayParams struct {
	// +optional
	Latency string `json:"latency,omitempty"`
	// +optional
	Jitter string `json:"jitter,omitempty"`
}

type LossParams struct {
	// +optional
	Loss string `json:"loss,omitempty"`
}

type DuplicationParams struct {
	// +optional
	Duplication string `json:"duplicate,omitempty"`
}

type CorruptionParams struct {
	// +optional
	Corruption string `json:"corrupt,omitempty"`
}

// NetworkChaosAction specify the action type of network Chaos
type NetworkChaosAction string

const (
	Delay       NetworkChaosAction = "Delay"
	Loss        NetworkChaosAction = "Loss"
	Duplication NetworkChaosAction = "Duplication"
	Corruption  NetworkChaosAction = "Corruption"
	Partition   NetworkChaosAction = "Partition"
	Bandwidth   NetworkChaosAction = "Bandwidth"
)

// Direction specifies the direction of action of network chaos
type Direction string

const (
	To   Direction = "to"
	From Direction = "from"
	Both Direction = "both"
)

// PodSelector used to select the target of the specified chaos
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
	// +optional
	ExpressionSelectors []metav1.LabelSelectorRequirement `json:"expressionSelectors,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Chaos{}, &ChaosList{})
}
