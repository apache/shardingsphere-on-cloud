/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// User TODO: description
type User struct {
	UserName string `json:"userName" yaml:"user"`
	PassWord string `json:"passWord" yaml:"-"`
	// +optional
	HostName string `json:"hostName,omitempty" yaml:"-"`
}

// Provider TODO: description
type Provider struct {
	Type string `json:"type" yaml:"type"`
}

// Auth TODO: description
type Auth struct {
	Users []User `json:"users" yaml:"users"`
	// +optional
	Provider Provider `json:"provider,omitempty"`
}

// Props TODO: description
type Props struct {
	// +optional
	KernelExecutorSize string `json:"kernel-executor-size,omitempty" yaml:"kernel-executor-size"`
	// +optional
	CheckTableMetadataEnabled string `json:"check-table-metadata-enabled,omitempty" yaml:"check-table-metadata-enabled"`
	// +optional
	ProxyBackendQueryFetchSize string `json:"proxy-backend-query-fetch-size,omitempty" yaml:"proxy-backend-query-fetch-size"`
	// +optional
	CheckDuplicateTableEnabled string `json:"check-duplicate-table-enabled,omitempty" yaml:"check-duplicate-table-enabled"`
	// +optional
	ProxyFrontendExecutorSize string `json:"proxy-frontend-executor-size,omitempty" yaml:"proxy-frontend-executor-size"`
	// +optional
	ProxyBackendExecutorSuitable string `json:"proxy-backend-executor-suitable,omitempty" yaml:"proxy-backend-executor-suitable"`
}

// ProxyConfigSpec defines the desired state of ProxyConfig
type ProxyConfigSpec struct {
	AUTHORITY Auth  `json:"AUTHORITY" yaml:"AUTHORITY"`
	Props     Props `json:"Props" yaml:"props"`
}

// ProxyConfigStatus defines the observed state of ProxyConfig
type ProxyConfigStatus struct {
	// crd change or create timestamp
	UpdateTime metav1.Time `json:"updateTime"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ProxyConfig is the Schema for the proxyconfigs API
type ProxyConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProxyConfigSpec   `json:"spec,omitempty"`
	Status ProxyConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:printcolumn:JSONPath=".status.updateTime",name=UpdateTime,type=string
//+kubebuilder:object:root=true

// ProxyConfigList contains a list of ProxyConfig
type ProxyConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProxyConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProxyConfig{}, &ProxyConfigList{})
}
