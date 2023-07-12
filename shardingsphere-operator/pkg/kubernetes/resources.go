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

package kubernetes

import (
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaosmesh"
	cloudnativepg "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/cloudnative-pg"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/deployment"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/hpa"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/service"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Resources respresents the resources
type Resources interface {
	KubernetesResources
	ExtendedResources
}

// KubernetesResource represents a Kubernetes resource
type KubernetesResources interface {
	Deployment() deployment.Deployment
	Service() service.Service
	ConfigMap() configmap.ConfigMap
	HPA() hpa.HorizontalPodAutoscaler
	// Job()
}

// ExtendedResources represents a extended resource
type ExtendedResources interface {
	ChaosMesh() chaosmesh.Chaos
	CloudNativePG() cloudnativepg.CloudNativePG
}

// NewResources return a instance of Resources
func NewResources(c client.Client) Resources {
	return &resources{
		KubernetesResources: &kubernetes{
			deployment: deployment.NewDeploymentClient(c),
			service:    service.NewServiceClient(c),
			configmap:  configmap.NewConfigMapClient(c),
			hpa:        hpa.NewHorizontalPodAutoscalerClient(c),
		},
		ExtendedResources: &extended{
			chaosmesh:     chaosmesh.NewChaos(c),
			cloudnativepg: cloudnativepg.NewCloudNativePGClient(c),
		},
	}
}

type resources struct {
	KubernetesResources
	ExtendedResources
}

var _ KubernetesResources = &kubernetes{}

type kubernetes struct {
	deployment deployment.Deployment
	service    service.Service
	configmap  configmap.ConfigMap
	hpa        hpa.HorizontalPodAutoscaler
}

// Deployment returns a Kubernetes deployment
func (r *kubernetes) Deployment() deployment.Deployment {
	return r.deployment
}

// Service returns a Kubernetes service
func (r *kubernetes) Service() service.Service {
	return r.service
}

// ConfigMap returns a Kubernetes configmap
func (r *kubernetes) ConfigMap() configmap.ConfigMap {
	return r.configmap
}

// HPA returns a Kubernetes HPA
func (r *kubernetes) HPA() hpa.HorizontalPodAutoscaler {
	return r.hpa
}

var _ ExtendedResources = &extended{}

type extended struct {
	chaosmesh     chaosmesh.Chaos
	cloudnativepg cloudnativepg.CloudNativePG
}

// ChaosMesh returns a extended chaosmesh
func (r *extended) ChaosMesh() chaosmesh.Chaos {
	return r.chaosmesh
}

// CloudNativePG returns a extended cloudnativepg
func (r *extended) CloudNativePG() cloudnativepg.CloudNativePG {
	return r.cloudnativepg
}
