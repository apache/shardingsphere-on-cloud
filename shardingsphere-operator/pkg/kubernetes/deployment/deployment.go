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

package deployment

import (
	"context"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/metadata"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/common"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewDeploymentClient creates a new Deployment
func NewDeploymentClient(c client.Client) Deployment {
	return deploymentClient{
		builder: builder{},
		getter: getter{
			Client: c,
		},
		setter: setter{
			Client: c,
		},
	}
}

// Deployment interface contains setter and getter
type Deployment interface {
	Builder
	Getter
	Setter
}

type deploymentClient struct {
	builder
	getter
	setter
}

// Getter get Deployment from different parameters
type Getter interface {
	GetByNamespacedName(context.Context, types.NamespacedName) (*appsv1.Deployment, error)
}

type getter struct {
	client.Client
}

// GetByNamespacedName returns Deployment from given namespaced name
func (dg getter) GetByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*appsv1.Deployment, error) {
	dp := &appsv1.Deployment{}
	if err := dg.Client.Get(ctx, namespacedName, dp); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	} else {
		return dp, nil
	}
}

// Setter get Deployment from different parameters
type Setter interface {
	Create(context.Context, *appsv1.Deployment) error
	Update(context.Context, *appsv1.Deployment) error
}

type setter struct {
	client.Client
}

// Create creates Deployment
func (ds setter) Create(ctx context.Context, dp *appsv1.Deployment) error {
	return ds.Client.Create(ctx, dp)
}

// Update updates Deployment
func (ds setter) Update(ctx context.Context, dp *appsv1.Deployment) error {
	return ds.Client.Update(ctx, dp)
}

// Builder build Deployment from given ComputeNode
type Builder interface {
	Build(context.Context, *v1alpha1.ComputeNode) *appsv1.Deployment
}

type builder struct{}

func (b builder) buildProbes(scb common.ContainerBuilder, cn *v1alpha1.ComputeNode) {
	if cn.Spec.Probes == nil {
		return
	}

	if cn.Spec.Probes.LivenessProbe != nil {
		scb.SetLivenessProbe(cn.Spec.Probes.LivenessProbe)
	}
	if cn.Spec.Probes.ReadinessProbe != nil {
		scb.SetReadinessProbe(cn.Spec.Probes.ReadinessProbe)
	}
	if cn.Spec.Probes.StartupProbe != nil {
		scb.SetStartupProbe(cn.Spec.Probes.StartupProbe)
	}
}

func (b builder) buildMetadata(ssbuilder ShardingSphereDeploymentBuilder, cn *v1alpha1.ComputeNode) {
	ssbuilder.SetName(cn.Name).
		SetNamespace(cn.Namespace).
		SetLabels(cn.Labels).
		SetAnnotations(cn.Annotations)
}

func getContainerPortsFromComputeNode(cn *v1alpha1.ComputeNode) []corev1.ContainerPort {
	ports := []corev1.ContainerPort{}
	for idx := range cn.Spec.PortBindings {
		ports = append(ports, corev1.ContainerPort{
			Name:          cn.Spec.PortBindings[idx].Name,
			HostIP:        cn.Spec.PortBindings[idx].HostIP,
			ContainerPort: cn.Spec.PortBindings[idx].ContainerPort,
			Protocol:      cn.Spec.PortBindings[idx].Protocol,
		})
	}
	return ports
}

func (b builder) buildSpec(ssbuilder ShardingSphereDeploymentBuilder, cn *v1alpha1.ComputeNode) {
	ssbuilder.SetSelectors(cn.Spec.Selector)
	ssbuilder.SetReplicas(&cn.Spec.Replicas)

	tpl := &corev1.PodTemplateSpec{}
	tm := metadata.NewMetadataBuilder()
	tm.SetLabels(cn.Labels)

	ports := getContainerPortsFromComputeNode(cn)

	scb := NewShardingSphereProxyContainerBuilder().
		SetVersion(cn.Spec.ServerVersion).
		SetPorts(ports).
		SetResources(cn.Spec.Resources)

	b.buildProbes(scb, cn)

	vcb := NewSharedVolumeAndMountBuilder().
		SetVolumeMountSize(1).
		SetName(defaultConfigVolumeName).
		SetVolumeSourceConfigMap(cn.Name).
		SetMountPath(0, defaultConfigVolumeMountPath)
	vc, vmc := vcb.Build()

	ssbuilder.AppendVolumes([]corev1.Volume{
		*vc,
	})
	scb.AppendVolumeMounts([]corev1.VolumeMount{*vmc[0]})

	sc := scb.BuildContainer()
	ssbuilder.AppendContainers([]corev1.Container{
		*sc,
	})

	if enabled, ok := cn.Annotations[DefaultAnnotationJavaAgentEnabled]; ok && enabled == "true" {
		ssbuilder.SetAgentBin(cn)
	}

	if cn.Spec.StorageNodeConnector != nil {
		if cn.Spec.StorageNodeConnector.Type == v1alpha1.ConnectorTypeMySQL {
			ssbuilder.SetMySQLConnector(cn)
		}
	}

	tpl.ObjectMeta = *tm.BuildMetadata()
	tpl.Spec = *ssbuilder.BuildPodSpec()
	ssbuilder.SetPodTemplateSpec(tpl)
}

// Build returns a new Deployment
func (b builder) Build(ctx context.Context, cn *v1alpha1.ComputeNode) *appsv1.Deployment {
	ssbuilder := NewShardingSphereDeploymentBuilder(cn.GetObjectMeta(), cn.GetObjectKind().GroupVersionKind())

	b.buildMetadata(ssbuilder, cn)
	b.buildSpec(ssbuilder, cn)

	return ssbuilder.BuildShardingSphereDeployment()
}
