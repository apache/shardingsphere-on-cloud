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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StorageNodePhaseStatus string

const (
	StorageNodePhaseReady          StorageNodePhaseStatus = "Ready"
	StorageNodePhaseNotReady       StorageNodePhaseStatus = "NotReady"
	StorageNodePhaseDeleting       StorageNodePhaseStatus = "Deleting"
	StorageNodePhaseDeleteComplete StorageNodePhaseStatus = "DeleteComplete"
)

type StorageNodeConditionType string

// StorageNodeConditionType shows some states during the startup process of storage node.
const (
	// StorageNodeConditionTypeAvailable means the all the instances and the cluster are ready, and the storage node is ready to provide external services.
	StorageNodeConditionTypeAvailable StorageNodeConditionType = "Available"
	// StorageNodeConditionTypeClusterReady means the cluster is ready, does not mean the instances are all ready.
	StorageNodeConditionTypeClusterReady StorageNodeConditionType = "ClusterReady"
)

type StorageNodeConditions []*StorageNodeCondition

// StorageNodeCondition contains details for the current condition of this StorageNode.
type StorageNodeCondition struct {
	Type           StorageNodeConditionType `json:"type"`
	Status         corev1.ConditionStatus   `json:"status"`
	LastUpdateTime metav1.Time              `json:"lastUpdateTime,omitempty"`
	Reason         string                   `json:"reason"`
	Message        string                   `json:"message"`
}

// ClusterStatus is the status of a database cluster, including the primary endpoint, reader endpoints, and other properties.
// Properties are some additional information about the cluster, like 'arn, identifier, credentials, etc.'
type ClusterStatus struct {
	Status          string   `json:"status"`
	PrimaryEndpoint Endpoint `json:"primaryEndpoint"`
	// +optional
	ReaderEndpoints []Endpoint `json:"readerEndpoints"`
	// +optional
	Properties map[string]string `json:"properties"`
}

type CredentialType struct {
	BasicCredential `json:"basic_credential"`
}

type InstanceStatus struct {
	Status   string   `json:"status"`
	Endpoint Endpoint `json:"primaryEndpoint"`
	// +optional
	Properties map[string]string `json:"properties"`
}

type BasicCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Endpoint struct {
	Address string `json:"address"`
	Port    int32  `json:"port"`
}

// +kubebuilder:object:root=true
// StorageNodeList contains a list of StorageNode
type StorageNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageNode `json:"items"`
}

// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name=Age,type=date
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// StorageNode is the Schema for the ShardingSphere storage unit
type StorageNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec StorageNodeSpec `json:"spec,omitempty"`
	// +optional
	Status StorageNodeStatus `json:"status,omitempty"`
}

// StorageNodeSpec defines the desired state of a set of storage units
type StorageNodeSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:DatabaseClass defined by: https://github.com/database-mesh/golang-sdk/blob/main/kubernetes/api/v1alpha1/databaseclass.go
	DatabaseClassName string `json:"databaseClassName"`
	// +optional
	Schema string `json:"schema"`
}

// StorageNodeStatus defines the actual state of a set of storage units
type StorageNodeStatus struct {
	// The generation observed by the StorageNode controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Phase is a brief summary of the StorageNode life cycle
	// There are two possible phase values:
	// Ready: StorageNode can already provide external services
	// NotReady: StorageNode cannot provide external services
	// +optional
	Phase StorageNodePhaseStatus `json:"phase"`

	// Conditions The conditions array, the reason and message fields
	// +optional
	Conditions StorageNodeConditions `json:"conditions"`

	// Cluster contains the current status of the StorageNode cluster
	// +optional
	Cluster ClusterStatus `json:"cluster,omitempty"`

	// Instance contains the current status of the StorageNode instance
	Instances []InstanceStatus `json:"instances,omitempty"`
}

const (
	StorageNodeInstanceStatusAvailable = "available"
	StorageNodeInstanceStatusBackingup = "backingup"
	StorageNodeInstanceStatusCreating  = "creating"
	StorageNodeInstanceStatusDeleting  = "deleting"
	StorageNodeInstanceStatusFailed    = "failed"
	StorageNodeInstanceStatusModifying = "modifying"
	StorageNodeInstanceStatusRebooting = "rebooting"
	StorageNodeInstanceStatusRenaming  = "renaming"
	StorageNodeInstanceStatusStarting  = "starting"
	StorageNodeInstanceStatusStopped   = "stopped"
	StorageNodeInstanceStatusStopping  = "stopping"
)

// AddCondition adds the given condition to the StorageNodeConditions.
func (c *StorageNodeConditions) AddCondition(condition *StorageNodeCondition) {
	*c = append(*c, condition)
}

// UpsertCondition updates the given condition in the StorageNodeConditions.
func (c *StorageNodeConditions) UpsertCondition(condition *StorageNodeCondition) {
	for i, existing := range *c {
		if existing.Type == condition.Type {
			(*c)[i] = condition
			return
		}
	}
	c.AddCondition(condition)
}

// RemoveCondition removes the given condition from the StorageNodeConditions.
func (c *StorageNodeConditions) RemoveCondition(conditionType StorageNodeConditionType) {
	var newConditions []*StorageNodeCondition
	for _, existing := range *c {
		if existing.Type != conditionType {
			newConditions = append(newConditions, existing)
		}
	}
	*c = newConditions
}

func init() {
	SchemeBuilder.Register(&StorageNode{}, &StorageNodeList{})
}
