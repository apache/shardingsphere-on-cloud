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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// StorageProviderBinding grants a namespace access to a cluster-scoped StorageProvider.
// Only cluster administrators should be allowed to create or modify this resource.
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=spb
type StorageProviderBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec StorageProviderBindingSpec `json:"spec,omitempty"`
}

// StorageProviderBindingSpec defines the StorageProvider granted to a namespace.
type StorageProviderBindingSpec struct {
	// +kubebuilder:validation:Required
	StorageProviderName string `json:"storageProviderName"`
}

// +kubebuilder:object:root=true
// StorageProviderBindingList contains a list of StorageProviderBinding resources.
type StorageProviderBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageProviderBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageProviderBinding{}, &StorageProviderBindingList{})
}
