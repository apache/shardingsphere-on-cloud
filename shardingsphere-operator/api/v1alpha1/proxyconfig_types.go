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

// ProxyConfigSpec defines the desired state of ShardingSphereProxyServerConfig
type ProxyConfigSpec struct {
	ClusterConfig ClusterConfig `json:"mode" yaml:"mode"`
	Authority     Auth          `json:"authority" yaml:"authority"`
	// +optional
	Props *Props `json:"props,omitempty" yaml:"props,omitempty"`
}

// ProxyConfigStatus defines the observed state of ShardingSphereProxyServerConfig
type ProxyConfigStatus struct {
	MetadataRepository string `json:"metadataRepository"`
}

//+kubebuilder:printcolumn:JSONPath=".status.metadataRepository",name=MetadataRepository,type=string
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ShardingSphereProxyServerConfig is the Schema for the proxyconfigs API
type ShardingSphereProxyServerConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProxyConfigSpec   `json:"spec,omitempty"`
	Status ProxyConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ShardingSphereProxyServerConfigList contains a list of ShardingSphereProxyServerConfig
type ShardingSphereProxyServerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShardingSphereProxyServerConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ShardingSphereProxyServerConfig{}, &ShardingSphereProxyServerConfigList{})
}

func (in *ShardingSphereProxyServerConfig) SetMetadataRepository(metadataType string) {
	in.Status.MetadataRepository = metadataType
}
