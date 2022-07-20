/*
 *   Copyright © 2022，SphereEx Authors
 *   All Rights Reserved.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

//ServiceType defines the Service in Kubernetes of ShardingSphere-Proxy
type ServiceType struct {
	// +kubebuilder:validation:Enum=ClusterIP;NodePort;LoadBalancer;ExternalName

	Type v1.ServiceType `json:"type"`
	// +kubebuilder:validation:Minimum=0

	// The port on each node on which this service is exposed when type is
	// NodePort or LoadBalancer.  Usually assigned by the system. If a value is
	// specified, in-range, and not in use it will be used, otherwise the
	// operation will fail.  If not specified, a port will be allocated if this
	// Service requires one.  If this field is specified when creating a
	// Service which does not need it, creation will fail. This field will be
	// wiped when updating a Service to no longer need it (e.g. changing type
	// from NodePort to ClusterIP).
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	// +optional
	NodePort int32 `json:"nodePort"`
}

//MySQLDriver Defines the mysql-driven version in ShardingSphere-proxy
type MySQLDriver struct {
	// +kubebuilder:validation:Pattern=`^([1-9]\d|[1-9])(\.([1-9]\d|\d)){2}$`
	// mysql-driven version,must be x.y.z
	Version string `json:"version"`
}

// AutomaticScaling HPA configuration
type AutomaticScaling struct {
	// +optional
	Enable bool `json:"enable,omitempty"`
	// +optional
	ScaleUpWindows int32 `json:"scaleUpWindows,omitempty"`
	// +optional
	ScaleDownWindows int32 `json:"scaleDownWindows,omitempty"`
	// +optional
	Target int32 `json:"target,omitempty"`
	// +optional
	MaxInstance int32 `json:"maxInstance,omitempty"`
	// +optional
	MinInstance int32 `json:"minInstance,omitempty"`
}

// ProxySpec defines the desired state of Proxy
type ProxySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Version  is the version of ShardingSphere-Proxy
	Version     string      `json:"version"`
	ServiceType ServiceType `json:"serviceType"`
	//Replicas is the expected number of replicas of ShardingSphere-Proxy
	Replicas int32 `json:"replicas"`
	// +optional
	AutomaticScaling *AutomaticScaling `json:"automaticScaling,omitempty"`
	// +optional
	ImagePullSecrets []v1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	// +kubebuilder:validation:MinLength=0
	// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`

	//ProxyConfigName is the name of the ProxyConfig CRD
	ProxyConfigName string `json:"proxyConfigName"`

	// +kubebuilder:validation:Minimum=0
	//Port is ShardingSphere-Proxy startup port
	Port int32 `json:"port"`
	// +optional
	MySQLDriver *MySQLDriver `json:"mySQLDriver,omitempty"`
	// +optional
	Resources *v1.ResourceRequirements `json:"resources,omitempty"`
	// Probes are not allowed for ephemeral containers.
	// +optional
	LivenessProbe *v1.Probe `json:"livenessProbe,omitempty"`
	// Probes are not allowed for ephemeral containers.
	// +optional
	ReadinessProbe *v1.Probe `json:"readinessProbe,omitempty" `
	// Probes are not allowed for ephemeral containers.
	// +optional
	StartupProbe *v1.Probe `json:"startupProbe,omitempty"`
}

//+kubebuilder:printcolumn:JSONPath=".status.readyNodes",name=ReadyNodes,type=integer
//+kubebuilder:printcolumn:JSONPath=".status.phase",name=Phase,type=string
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Proxy is the Schema for the proxies API
type Proxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProxySpec   `json:"spec,omitempty"`
	Status ProxyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ProxyList contains a list of Proxy
type ProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Proxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Proxy{}, &ProxyList{})
}
