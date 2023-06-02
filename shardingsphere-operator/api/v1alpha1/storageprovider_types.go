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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	AnnotationsVPCSecurityGroupIds = "storageproviders.shardingsphere.apache.org/vpc-security-group-ids"
	AnnotationsSubnetGroupName     = "storageproviders.shardingsphere.apache.org/vpc-subnet-group-name"
	AnnotationsAvailabilityZones   = "storageproviders.shardingsphere.apache.org/availability-zones"
	AnnotationsClusterIdentifier   = "storageproviders.shardingsphere.apache.org/cluster-identifier"
	AnnotationsInstanceIdentifier  = "storageproviders.shardingsphere.apache.org/instance-identifier"
	AnnotationsInstanceDBName      = "storageproviders.shardingsphere.apache.org/instance-db-name"
	AnnotationsSnapshotIdentifier  = "storageproviders.shardingsphere.apache.org/snapshot-identifier"
	AnnotationsMasterUsername      = "storageproviders.shardingsphere.apache.org/master-username"
	AnnotationsMasterUserPassword  = "storageproviders.shardingsphere.apache.org/master-user-password"

	ProvisionerAWSRDSInstance = "storageproviders.shardingsphere.apache.org/aws-rds-instance"
	ProvisionerAWSRDSCluster  = "storageproviders.shardingsphere.apache.org/aws-rds-cluster"
	ProvisionerAWSAurora      = "storageproviders.shardingsphere.apache.org/aws-aurora"
)

// StorageReclaimPolicy defines the reclaim policy for storage
type StorageReclaimPolicy string

const (
	// StorageReclaimPolicyDeleteWithFinalSnapshot The database will be deleted with a final snapshot reserved.
	StorageReclaimPolicyDeleteWithFinalSnapshot StorageReclaimPolicy = "DeleteWithFinalSnapshot"
	// StorageReclaimPolicyDelete The database will be deleted.
	StorageReclaimPolicyDelete StorageReclaimPolicy = "Delete"
	// StorageReclaimPolicyRetain The database will be retained.
	// The default policy is Retain.
	StorageReclaimPolicyRetain StorageReclaimPolicy = "Retain"
)

// StorageProviderSpec defines the desired state of StorageProvider
type StorageProviderSpec struct {
	Provisioner string            `json:"provisioner"`
	Parameters  map[string]string `json:"parameters"`

	//+kubebuilder:validation:Optional
	//+kubebuilder:validation:Enum=DeleteWithFinalSnapshot;Delete;Retain
	//+kubebuilder:default:=Retain
	//+optional
	ReclaimPolicy StorageReclaimPolicy `json:"reclaimPolicy,omitempty"`
}

// StorageProviderStatus defines the observed state of StorageProvider
type StorageProviderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster,shortName=sp

// StorageProvider is the Schema for the storageproviders API
type StorageProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status StorageProviderStatus `json:"status,omitempty"`
	Spec   StorageProviderSpec   `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// StorageProviderList contains a list of StorageProvider
type StorageProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageProvider `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageProvider{}, &StorageProviderList{})
}
