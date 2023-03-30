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
	batchV1Beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
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
	InjectReqs batchV1Beta1.JobTemplateSpec `json:"injectReqs,omitempty"`

	ChaosKind ChaosKind `json:"chaosKind,omitempty"`

	EmbedChaos `json:".inline"`

	Verify batchV1Beta1.JobTemplateSpec `json:"Verify,omitempty"`
}

type ChaosKind string

const (
	WorkFlowKind ChaosKind = "workflow"

	NetworkChaosKind ChaosKind = "networkChaos"

	PodChaosKind ChaosKind = "podChaos"
)

type EmbedChaos struct {
	// +optional
	NetworkChaos *NetworkChaosSpec `json:"networkChaos,omitempty"`
	// +optional
	PodChaos *PodChaosSpec `json:"podChaos,omitempty"`
	// +optional
	Workflow *WorkflowSpec `json:"workflow,omitempty"`
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
	JobFinish   Jobschedule = "JobFinish"
)

// ShardingSphereChaosStatus defines the actual state of ShardingSphereChaos
type ShardingSphereChaosStatus struct {
	ChaosCondition DeploymentCondition `json:"deploymentCondition"`
	InjectStatus   Jobschedule         `json:"InjectStatus"`
	VerifyStatus   Jobschedule         `json:"VerifyStatus"`
}

// pod chaos

type PodChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PodChaosSpec `json:"spec"`
}

type PodChaosSpec struct {
	PodSelector `json:",inline"`

	// Action defines the specific pod chaos action.
	// Supported action: pod-kill / pod-failure / container-kill
	// Default action: pod-kill
	// +kubebuilder:validation:Enum=pod-kill;pod-failure;container-kill
	Action PodChaosAction `json:"action"`

	// Duration represents the duration of the chaos action.
	// It is required when the action is `PodFailureAction`.
	// A duration string is a possibly signed sequence of
	// decimal numbers, each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// +optional
	Duration *string `json:"duration,omitempty" webhook:"Duration"`

	// GracePeriod is used in pod-kill action. It represents the duration in seconds before the pod should be deleted.
	// Value must be non-negative integer. The default value is zero that indicates delete immediately.
	// +optional
	// +kubebuilder:validation:Minimum=0
	GracePeriod int64 `json:"gracePeriod"`
}

type PodChaosAction string

const (
	// PodKillAction represents the chaos action of killing pods.
	PodKillAction PodChaosAction = "pod-kill"
	// PodFailureAction represents the chaos action of injecting errors to pods.
	// This action will cause the pod to not be created for a while.
	PodFailureAction PodChaosAction = "pod-failure"
	// ContainerKillAction represents the chaos action of killing the container
	ContainerKillAction PodChaosAction = "container-kill"
)

//network chaos

type NetworkChaosSpec struct {
	PodSelector `json:",inline"`

	// Action defines the specific network chaos action.
	// Supported action: partition, netem, delay, loss, duplicate, corrupt
	// Default action: delay
	// +kubebuilder:validation:Enum=netem;delay;loss;duplicate;corrupt;partition;bandwidth
	Action NetworkChaosAction `json:"action"`

	// Device represents the network device to be affected.
	// +optional
	Device string `json:"device,omitempty"`

	// Duration represents the duration of the chaos action
	Duration *string `json:"duration,omitempty" webhook:"Duration"`

	// TcParameter represents the traffic control definition
	TcParameter `json:",inline"`

	// Direction represents the direction, this applies on netem and network partition action
	// +optional
	// +kubebuilder:validation:Enum=to;from;both
	// +kubebuilder:default=to
	Direction Direction `json:"direction,omitempty"`

	// Target represents network target, this applies on netem and network partition action
	// +optional
	Target *PodSelector `json:"target,omitempty" webhook:",nilable"`

	// TargetDevice represents the network device to be affected in target scope.
	// +optional
	TargetDevice string `json:"targetDevice,omitempty"`

	// ExternalTargets represents network targets outside k8s
	// +optional
	ExternalTargets []string `json:"externalTargets,omitempty"`
}

type NetworkChaosAction string

const (
	// NetemAction is a combination of several chaos actions i.e. delay, loss, duplicate, corrupt.
	// When using this action multiple specs are merged into one Netem RPC and sends to chaos daemon.
	NetemAction NetworkChaosAction = "netem"

	// DelayAction represents the chaos action of adding delay on pods.
	DelayAction NetworkChaosAction = "delay"

	// LossAction represents the chaos action of losing packets on pods.
	LossAction NetworkChaosAction = "loss"

	// DuplicateAction represents the chaos action of duplicating packets on pods.
	DuplicateAction NetworkChaosAction = "duplicate"

	// CorruptAction represents the chaos action of corrupting packets on pods.
	CorruptAction NetworkChaosAction = "corrupt"

	// PartitionAction represents the chaos action of network partition of pods.
	PartitionAction NetworkChaosAction = "partition"

	// BandwidthAction represents the chaos action of network bandwidth of pods.
	BandwidthAction NetworkChaosAction = "bandwidth"
)

type Direction string

const (
	// To represents network packet from source to target
	To Direction = "to"

	// From represents network packet to source from target
	From Direction = "from"

	// Both represents both directions
	Both Direction = "both"
)

// DelaySpec defines detail of a delay action
type DelaySpec struct {
	Latency string `json:"latency" webhook:"Duration"`
	// +optional
	Correlation string `json:"correlation,omitempty" default:"0" webhook:"FloatStr"`
	// +optional
	Jitter string `json:"jitter,omitempty" default:"0ms" webhook:"Duration"`
	// +optional
	Reorder *ReorderSpec `json:"reorder,omitempty"`
}

// LossSpec defines detail of a loss action
type LossSpec struct {
	Loss string `json:"loss" webhook:"FloatStr"`
	// +optional
	Correlation string `json:"correlation,omitempty" default:"0" webhook:"FloatStr"`
}

// DuplicateSpec defines detail of a duplicate action
type DuplicateSpec struct {
	Duplicate string `json:"duplicate" webhook:"FloatStr"`
	// +optional
	Correlation string `json:"correlation,omitempty" default:"0" webhook:"FloatStr"`
}

// CorruptSpec defines detail of a corrupt action
type CorruptSpec struct {
	Corrupt string `json:"corrupt" webhook:"FloatStr"`
	// +optional
	Correlation string `json:"correlation,omitempty" default:"0" webhook:"FloatStr"`
}

// BandwidthSpec defines detail of bandwidth limit.
type BandwidthSpec struct {
	// Rate is the speed knob. Allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.
	Rate string `json:"rate" webhook:"Rate"`
	// Limit is the number of bytes that can be queued waiting for tokens to become available.
	// +kubebuilder:validation:Minimum=1
	Limit uint32 `json:"limit"`
	// Buffer is the maximum amount of bytes that tokens can be available for instantaneously.
	// +kubebuilder:validation:Minimum=1
	Buffer uint32 `json:"buffer"`
	// Peakrate is the maximum depletion rate of the bucket.
	// The peakrate does not need to be set, it is only necessary
	// if perfect millisecond timescale shaping is required.
	// +optional
	// +kubebuilder:validation:Minimum=0
	Peakrate *uint64 `json:"peakrate,omitempty"`
	// Minburst specifies the size of the peakrate bucket. For perfect
	// accuracy, should be set to the MTU of the interface.  If a
	// peakrate is needed, but some burstiness is acceptable, this
	// size can be raised. A 3000 byte minburst allows around 3mbit/s
	// of peakrate, given 1000 byte packets.
	// +optional
	// +kubebuilder:validation:Minimum=0
	Minburst *uint32 `json:"minburst,omitempty"`
}

// ReorderSpec defines details of packet reorder.
type ReorderSpec struct {
	Reorder string `json:"reorder" webhook:"FloatStr"`
	// +optional
	Correlation string `json:"correlation,omitempty" default:"0" webhook:"FloatStr"`
	Gap         int    `json:"gap"`
}

// TcParameter represents the parameters for a traffic control chaos
type TcParameter struct {
	// Delay represents the detail about delay action
	// +optional
	Delay *DelaySpec `json:"delay,omitempty"`

	// Loss represents the detail about loss action
	// +optional
	Loss *LossSpec `json:"loss,omitempty"`

	// DuplicateSpec represents the detail about loss action
	// +optional
	Duplicate *DuplicateSpec `json:"duplicate,omitempty"`

	// Corrupt represents the detail about corrupt action
	// +optional
	Corrupt *CorruptSpec `json:"corrupt,omitempty"`

	// Bandwidth represents the detail about bandwidth control action
	// +optional
	Bandwidth *BandwidthSpec `json:"bandwidth,omitempty"`
}

//workflow

type TemplateType string

const (
	TypeTask     TemplateType = "Task"
	TypeSerial   TemplateType = "Serial"
	TypeParallel TemplateType = "Parallel"
	TypeSuspend  TemplateType = "Suspend"
	TypeSchedule TemplateType = "Schedule"
)

type WorkflowSpec struct {
	//工作流入口点
	Entry string `json:"entry"`

	Templates []WorkFlowTemplate `json:"templates"`
}

type WorkFlowTemplate struct {
	Name string       `json:"name"`
	Type TemplateType `json:"templateType"`
	// +optional
	Deadline *string `json:"deadline,omitempty"`
	// Task describes the behavior of the custom task. Only used when Type is TypeTask.
	// +optional
	Task *Task `json:"task,omitempty"`
	// Children describes the children steps of serial or parallel node. Only used when Type is TypeSerial or TypeParallel.
	// +optional
	Children []string `json:"children,omitempty"`
	// ConditionalBranches describes the conditional branches of custom tasks. Only used when Type is TypeTask.
	// +optional
	ConditionalBranches []ConditionalBranch `json:"conditionalBranches,omitempty"`
	// EmbedChaos describe the chaos to be injected with chaos nodes. Only used when Type is Type<Something>Chaos.
	// +optional
	*EmbedChaos `json:",inline"`
	// Schedule describe the Schedule(describing scheduled chaos) to be injected with chaos nodes. Only used when Type is TypeSchedule.
	// +optional
	Schedule *ChaosOnlyScheduleSpec `json:"schedule,omitempty"`
}

type ConditionalBranch struct {
	// Target is the name of other template, if expression is evaluated as true, this template will be spawned.
	Target string `json:"target"`
	// Expression is the expression for this conditional branch, expected type of result is boolean. If expression is empty, this branch will always be selected/the template will be spawned.
	// +optional
	Expression string `json:"expression,omitempty"`
}

type ScheduleTemplateType string

type ConcurrencyPolicy string

var (
	ForbidConcurrent ConcurrencyPolicy = "Forbid"
	AllowConcurrent  ConcurrencyPolicy = "Allow"
)

// ChaosOnlyScheduleSpec is very similar with ScheduleSpec, but it could not schedule Workflow
// because we could not resolve nested CRD now
type ChaosOnlyScheduleSpec struct {
	Schedule string `json:"schedule"`

	// +optional
	// +nullable
	// +kubebuilder:validation:Minimum=0
	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds"`

	// +optional
	// +kubebuilder:validation:Enum=Forbid;Allow
	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrencyPolicy"`

	// +optional
	// +kubebuilder:validation:Minimum=1
	HistoryLimit int `json:"historyLimit,omitempty"`

	// TODO: use a custom type, as `TemplateType` contains other possible values
	Type ScheduleTemplateType `json:"type"`

	EmbedChaos `json:",inline"`
}

type Task struct {
	// Container is the main container image to run in the pod
	Container *corev1.Container `json:"container,omitempty"`

	// Volumes is a list of volumes that can be mounted by containers in a template.
	// +patchStrategy=merge
	// +patchMergeKey=name
	Volumes []corev1.Volume `json:"volumes,omitempty" patchStrategy:"merge" patchMergeKey:"name"`

	// TODO: maybe we could specify parameters in other ways, like loading context from file
}

//selector

// LabelSelectorRequirements is list of LabelSelectorRequirement
type LabelSelectorRequirements []metav1.LabelSelectorRequirement

// SelectorMode represents the mode to run chaos action.
type SelectorMode string

const (
	// OneMode represents that the system will do the chaos action on one object selected randomly.
	OneMode SelectorMode = "one"
	// AllMode represents that the system will do the chaos action on all objects
	// regardless of status (not ready or not running pods includes).
	// Use this label carefully.
	AllMode SelectorMode = "all"
	// FixedMode represents that the system will do the chaos action on a specific number of running objects.
	FixedMode SelectorMode = "fixed"
	// FixedPercentMode to specify a fixed % that can be inject chaos action.
	FixedPercentMode SelectorMode = "fixed-percent"
	// RandomMaxPercentMode to specify a maximum % that can be inject chaos action.
	RandomMaxPercentMode SelectorMode = "random-max-percent"
)

// GenericSelectorSpec defines some selectors to select objects.
type GenericSelectorSpec struct {
	// Namespaces is a set of namespace to which objects belong.
	// +optional
	Namespaces []string `json:"namespaces,omitempty"`

	// Map of string keys and values that can be used to select objects.
	// A selector based on fields.
	// +optional
	FieldSelectors map[string]string `json:"fieldSelectors,omitempty"`

	// Map of string keys and values that can be used to select objects.
	// A selector based on labels.
	// +optional
	LabelSelectors map[string]string `json:"labelSelectors,omitempty"`

	// a slice of label selector expressions that can be used to select objects.
	// A list of selectors based on set-based label expressions.
	// +optional
	ExpressionSelectors LabelSelectorRequirements `json:"expressionSelectors,omitempty" swaggerignore:"true"`

	// Map of string keys and values that can be used to select objects.
	// A selector based on annotations.
	// +optional
	AnnotationSelectors map[string]string `json:"annotationSelectors,omitempty"`
}

// PodSelectorSpec defines the some selectors to select objects.
// If the all selectors are empty, all objects will be used in chaos experiment.
type PodSelectorSpec struct {
	GenericSelectorSpec `json:",inline"`

	// Nodes is a set of node name and objects must belong to these nodes.
	// +optional
	Nodes []string `json:"nodes,omitempty"`

	// Pods is a map of string keys and a set values that used to select pods.
	// The key defines the namespace which pods belong,
	// and the each values is a set of pod names.
	// +optional
	Pods map[string][]string `json:"pods,omitempty"`

	// Map of string keys and values that can be used to select nodes.
	// Selector which must match a node's labels,
	// and objects must belong to these selected nodes.
	// +optional
	NodeSelectors map[string]string `json:"nodeSelectors,omitempty"`

	// PodPhaseSelectors is a set of condition of a pod at the current time.
	// supported value: Pending / Running / Succeeded / Failed / Unknown
	// +optional
	PodPhaseSelectors []string `json:"podPhaseSelectors,omitempty"`
}

func (in *PodSelectorSpec) DefaultNamespace(namespace string) {
	if len(in.Namespaces) == 0 {
		in.Namespaces = []string{namespace}
	}
}

type PodSelector struct {
	// Selector is used to select pods that are used to inject chaos action.
	Selector PodSelectorSpec `json:"selector"`

	// Mode defines the mode to run chaos action.
	// Supported mode: one / all / fixed / fixed-percent / random-max-percent
	// +kubebuilder:validation:Enum=one;all;fixed;fixed-percent;random-max-percent
	Mode SelectorMode `json:"mode"`

	// Value is required when the mode is set to `FixedMode` / `FixedPercentMode` / `RandomMaxPercentMode`.
	// If `FixedMode`, provide an integer of pods to do chaos action.
	// If `FixedPercentMode`, provide a number from 0-100 to specify the percent of pods the server can do chaos action.
	// IF `RandomMaxPercentMode`,  provide a number from 0-100 to specify the max percent of pods to do chaos action
	// +optional
	Value string `json:"value,omitempty"`
}

type ContainerSelector struct {
	PodSelector `json:",inline"`

	// ContainerNames indicates list of the name of affected container.
	// If not set, the first container will be injected
	// +optional
	ContainerNames []string `json:"containerNames,omitempty"`
}
