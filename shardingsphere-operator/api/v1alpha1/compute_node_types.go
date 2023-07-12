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

// +kubebuilder:object:root=true
// ComputeNodeList contains a list of ComputeNode
type ComputeNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNode `json:"items"`
}

// +kubebuilder:printcolumn:JSONPath=".status.ready",name=Ready,type=string
// +kubebuilder:printcolumn:JSONPath=".status.phase",name=Status,type=string
// +kubebuilder:printcolumn:JSONPath=".status.loadBalancer.clusterIP",name="Cluster-IP",type=string
// +kubebuilder:printcolumn:JSONPath=".spec.portBindings[*].servicePort",name="ServicePorts",type=integer
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name=Age,type=date
// +kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas,selectorpath=.status.selector
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// ComputeNode is the Schema for the ShardingSphere Proxy API
type ComputeNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ComputeNodeSpec `json:"spec,omitempty"`
	// +optional
	Status ComputeNodeStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type PrivilegeType string

const (
	AllPermitted PrivilegeType = "ALL_PERMITTED"
)

// ComputeNodePrivilege for storage node, the default value is ALL_PERMITTED
type ComputeNodePrivilege struct {
	Type PrivilegeType `json:"type,omitempty" yaml:"type,omitempty"`
}

// ComputeNodeUser is a slice about authorized host and password for compute node.
// Format:
// user:<username>@<hostname>,hostname is % or empty string means do not care about authorized host
// password:<password>
type ComputeNodeUser struct {
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}

// ComputeNodeAuthority  is used to set up initial user to login compute node, and authority data of storage node.
type ComputeNodeAuthority struct {
	// +optional
	Users []ComputeNodeUser `json:"users,omitempty" yaml:"users,omitempty"`
	// +optional
	Privilege ComputeNodePrivilege `json:"privilege,omitempty" yaml:"privilege,omitempty"`
}

type RepositoryType string

const (
	RepositoryTypeZookeeper RepositoryType = "ZooKeeper"
	RepositoryTypeEtcd      RepositoryType = "Etcd"
)

// Repository is the metadata persistent store for ShardingSphere
type Repository struct {
	// +kubebuilder:validation:Enum=ZooKeeper;Etcd
	// type of metadata repository
	Type RepositoryType `json:"type,omitempty" yaml:"type,omitempty"`
	// properties of metadata repository
	// +optional
	Props Properties `json:"props,omitempty" yaml:"props,omitempty"`
}

type ModeType string

const (
	ModeTypeCluster    ModeType = "Cluster"
	ModeTypeStandalone ModeType = "Standalone"
)

// ComputeNodeServerMode is the mode for ShardingSphere Proxy
type ComputeNodeServerMode struct {
	// +optional
	Repository Repository `json:"repository,omitempty" yaml:"repository,omitempty"`
	// +optional
	Type ModeType `json:"type,omitempty" yaml:"type,omitempty"`
}

// ServerConfig defines the bootstrap config for a ShardingSphere Proxy
type ServerConfig struct {
	// +optional
	Authority ComputeNodeAuthority `json:"authority,omitempty" yaml:"authority,omitempty"`
	// +optional
	Mode ComputeNodeServerMode `json:"mode,omitempty" yaml:"mode,omitempty"`
	// +optional
	Props Properties `json:"props,omitempty" yaml:"props,omitempty"`
}

// LogbackConfig contains contents of the expected logback.xml
type LogbackConfig string

// +kubebuilder:pruning:PreserveUnknownFields
type Properties map[string]string

type LoggingFile struct {
	Props Properties `json:"props,omitempty" yaml:"props,omitempty"`
}

// PluginLogging defines the plugin for logging
type PluginLogging struct {
	File LoggingFile `json:"file,omitempty" yaml:"File,omitempty"`
}

type Prometheus struct {
	Host  string     `json:"host" yaml:"host"`
	Port  int32      `json:"port" yaml:"port"`
	Props Properties `json:"props,omitempty" yaml:"props,omitempty"`
}

// PluginMetrics defines the plugin for metrics
type PluginMetrics struct {
	Prometheus Prometheus `json:"prometheus,omitempty" yaml:"Prometheus,omitempty"`
}

type OpenTelemetry struct {
	Props Properties `json:"props,omitempty" yaml:"props,omitempty"`
}

type OpenTracing struct {
	Props Properties `json:"props,omitempty" yaml:"props,omitempty" `
}

// PluginTracing defines the plugin for tracing
type PluginTracing struct {
	// +optional
	OpenTracing OpenTracing `json:"openTracing,omitempty" yaml:"OpenTracing,omitempty"`
	// +optional
	OpenTelemetry OpenTelemetry `json:"openTelemetry,omitempty" yaml:"OpenTelemetry,omitempty"`
}

// AgentPlugin defines a set of plugins for ShardingSphere Agent
type AgentPlugin struct {
	// +optional
	Logging *PluginLogging `json:"logging,omitempty" yaml:"logging,omitempty"`
	// +optional
	Metrics *PluginMetrics `json:"metrics,omitempty" yaml:"metrics,omitempty"`
	// +optional
	Tracing *PluginTracing `json:"tracing,omitempty" yaml:"tracing,omitempty"`
}

// AgentConfig defines the config for ShardingSphere-Agent, renderred as agent.yaml
type AgentConfig struct {
	Plugins *AgentPlugin `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

// ServiceType defines the Service in Kubernetes of ShardingSphere-Proxy
type Service struct {
	Ports []corev1.ServicePort `json:"ports,omitempty" yaml:"ports,omitempty"`
	// +kubebuilder:validation:Enum=ClusterIP;NodePort;LoadBalancer;ExternalName
	Type corev1.ServiceType `json:"type" yaml:"type"`
}

// ProxyProbe defines the probe actions for LivenesProbe, ReadinessProbe and StartupProbe
type ProxyProbe struct {
	// Probes are not allowed for ephemeral containers.
	// +optional
	LivenessProbe *corev1.Probe `json:"livenessProbe,omitempty" yaml:"livenessProbe,omitempty"`
	// Probes are not allowed for ephemeral containers.
	// +optional
	ReadinessProbe *corev1.Probe `json:"readinessProbe,omitempty" yaml:"readinessProbe,omitempty"`
	// Probes are not allowed for ephemeral containers.
	// +optional
	StartupProbe *corev1.Probe `json:"startupProbe,omitempty" yaml:"startupProbe,omitempty"`
}

// ConnectorType defines the frontend protocol for ShardingSphere Proxy
type ConnectorType string

const (
	ConnectorTypeMySQL      ConnectorType = "mysql"
	ConnectorTypePostgreSQL ConnectorType = "postgresql"
)

// MySQLDriver Defines the mysql-driven version in ShardingSphere-proxy
type StorageNodeConnector struct {
	Type ConnectorType `json:"type" yaml:"type"`
	// +kubebuilder:validation:Pattern=`^([1-9]\d|[1-9])(\.([1-9]\d|\d)){2}$`
	// mysql-driven version,must be x.y.z
	Version string `json:"version" yaml:"version"`
}

// BootstrapConfig is used for any ShardingSphere Proxy startup
type BootstrapConfig struct {
	// +optional
	ServerConfig ServerConfig `json:"serverConfig,omitempty" yaml:"serverConfig,omitempty"`
	// +optional
	LogbackConfig LogbackConfig `json:"logbackConfig,omitempty" yaml:"logbackConfig,omitempty"`
	// +optional
	AgentConfig AgentConfig `json:"agentConfig,omitempty" yaml:"agentConfig,omitempty"`
}

type PortBinding struct {
	// If specified, this must be an IANA_SVC_NAME and unique within the pod. Each
	// named port in a pod must have a unique name. Name for the port that can be
	// referred to by services.
	// +optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Number of port to expose on the pod's IP address.
	// This must be a valid port number, 0 < x < 65536.
	ContainerPort int32 `json:"containerPort" yaml:"containerPort"`
	// Protocol for port. Must be UDP, TCP, or SCTP.
	// Defaults to "TCP".
	// +optional
	// +default="TCP"
	Protocol corev1.Protocol `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	// What host IP to bind the external port to.
	// +optional
	HostIP string `json:"hostIP,omitempty" yaml:"hostIP,omitempty"`

	// The port that will be exposed by this service.
	ServicePort int32 `json:"servicePort" yaml:"servicePort"`

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
	NodePort int32 `json:"nodePort,omitempty" yaml:"nodePort,omitempty"`
}

// ProxySpec defines the desired state of ShardingSphereProxy
type ComputeNodeSpec struct {
	StorageNodeConnector *StorageNodeConnector `json:"storageNodeConnector,omitempty" yaml:"storageNodeConnector,omitempty"`
	// version  is the version of ShardingSphere-Proxy
	ServerVersion string `json:"serverVersion,omitempty" yaml:"serverVersion,omitempty"`

	// replicas is the expected number of replicas of ShardingSphere-Proxy
	// +optional
	Replicas int32 `json:"replicas" yaml:"replicas"`
	// selector defines a set of label selectors
	Selector *metav1.LabelSelector `json:"selector" yaml:"selector"`

	// +optional
	Probes *ProxyProbe `json:"probes,omitempty" yaml:"probes,omitempty"`
	// +optional
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty" yaml:"imagePullSecrets,omitempty"`
	// +optional
	Env []corev1.EnvVar `json:"env,omitempty" yaml:"env,omitempty"`
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty" yaml:"resources,omitempty"`
	// +optional
	PortBindings []PortBinding `json:"portBindings,omitempty" yaml:"portBindings,omitempty"`

	// +kubebuilder:validation:Enum=ClusterIP;NodePort;LoadBalancer;ExternalName
	// +optional
	ServiceType corev1.ServiceType `json:"serviceType,omitempty" yaml:"serviceType,omitempty"`

	// +optional
	Bootstrap BootstrapConfig `json:"bootstrap,omitempty" yaml:"bootstrap,omitempty"`
}

// ComputeNodeStatus defines the observed state of ShardingSphere Proxy
type ComputeNodeStatus struct {
	Selector string `json:"selector" yaml:"selector"`

	Replicas int32 `json:"replicas" yaml:"replicas"`

	Ready string `json:"ready,omitempty" yaml:"ready,omitempty"`
	// The generation observed by the deployment controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" yaml:"observedGeneration,omitempty"`

	// ShardingSphere-Proxy phase are a brief summary of the ShardingSphere-Proxy life cycle
	// There are two possible phase values:
	// Ready: ShardingSphere-Proxy can already provide external services
	// NotReady: ShardingSphere-Proxy cannot provide external services
	// +optional
	Phase ComputeNodePhaseStatus `json:"phase,omitempty" yaml:"phase,omitempty"`

	// Conditions The conditions array, the reason and message fields
	// +optional
	Conditions []ComputeNodeCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	// LoadBalancer contains the current status of the load-balancer,
	// if one is present.
	// +optional
	LoadBalancer LoadBalancerStatus `json:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty"`
}

// LoadBalancerStatus represents the status of service endpoints
type LoadBalancerStatus struct {
	// +optional
	ClusterIP string `json:"clusterIP,omitempty" yaml:"clusterIP,omitempty"`

	// Ingress is a list containing ingress points for the load-balancer.
	// Traffic intended for the service should be sent to these ingress points.
	// +optional
	Ingress []corev1.LoadBalancerIngress `json:"ingress,omitempty" yaml:"ingress,omitempty"`
}

// ComputeNodePhase represents a brief summary of the compute node
type ComputeNodePhaseStatus string

const (
	// ComputeNodeStatusReady indicates that at least one compute node is ready for connections
	ComputeNodeStatusReady ComputeNodePhaseStatus = "Ready"
	// ComputeNodeStatusNotReady indicates that no compute node is ready
	ComputeNodeStatusNotReady ComputeNodePhaseStatus = "NotReady"
	// ComputeNodeStatusUnknown indicates that cannot determine the status of compute node at present
	ComputeNodeStatusUnknown ComputeNodePhaseStatus = "Unknown"
)

// ComputeNodeConditioType represents the type of a compute node condition during the startup process of ShardingSphere-Proxy
type ComputeNodeConditionType string

const (
	// ComputeNodeConditionPending indicates that at least one pod is in pending phase
	ComputeNodeConditionPending ComputeNodeConditionType = "Pending"
	// ComputeNodeConditionInitialized indicates that at least one pod is scheduled
	ComputeNodeConditionDeployed ComputeNodeConditionType = "Deployed"
	// ComputeNodeConditionInitialized indicates that at least one pod is initialized
	ComputeNodeConditionInitialized ComputeNodeConditionType = "Initialized"
	// ComputeNodeConditionInitialized indicates that at least one pod is started
	ComputeNodeConditionStarted ComputeNodeConditionType = "Started"
	// ComputeNodeConditionInitialized indicates that at least one pod is ready
	ComputeNodeConditionReady ComputeNodeConditionType = "Ready"
	// ComputeNodeConditionInitialized indicates that at least one pod is unknown
	ComputeNodeConditionUnknown ComputeNodeConditionType = "Unknown"
	// ComputeNodeConditionInitialized indicates that at least one pod is failed
	ComputeNodeConditionFailed ComputeNodeConditionType = "Failed"
	// ComputeNodeConditionInitialized indicates that at least one pod is succeed
	ComputeNodeConditionSucceed ComputeNodeConditionType = "Succeed"
)

// ConditionStatus represents the validation status of a condition
type ConditionStatus string

const (
	ConditionStatusTrue    ConditionStatus = "True"
	ConditionStatusFalse   ConditionStatus = "False"
	ConditionStatusUnknown ConditionStatus = "Unknown"
)

// ComputeNodeCondition defines a condition template
type ComputeNodeCondition struct {
	Type               ComputeNodeConditionType `json:"type" yaml:"type"`
	Status             ConditionStatus          `json:"status" yaml:"status"`
	LastTransitionTime metav1.Time              `json:"lastTransitionTime,omitempty" yaml:"lastTransitionTime,omitempty"`
	LastUpdateTime     metav1.Time              `json:"lastUpdateTime,omitempty" yaml:"lastUpdateTime,omitempty"`
	Reason             string                   `json:"reason" yaml:"reason"`
	Message            string                   `json:"message" yaml:"message"`
}

func init() {
	SchemeBuilder.Register(&ComputeNode{}, &ComputeNodeList{})
}
