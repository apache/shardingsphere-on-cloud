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
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/metadata"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/pod"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeploymentBuilder returns a deployment builder
type DeploymentBuilder interface {
	metadata.MetadataBuilder
	pod.PodSpecBuilder

	SetSelectors(selectors *metav1.LabelSelector) DeploymentBuilder
	SetReplicas(r *int32) DeploymentBuilder
	SetRollingUpdateStrategy(maxUnavailable, maxSurge int) DeploymentBuilder

	SetPodTemplateMetadata(obj *metav1.ObjectMeta) DeploymentBuilder
	SetPodTemplateSpec(tpl *corev1.PodTemplateSpec) DeploymentBuilder

	BuildDeployment() *appsv1.Deployment
}

// NewDeploymentBuilder creates a new DeploymentBuilder
func NewDeploymentBuilder() DeploymentBuilder {
	return &deploymentBuilder{
		deployment:      &appsv1.Deployment{},
		PodSpecBuilder:  pod.NewPodBuilder(),
		MetadataBuilder: metadata.NewMetadataBuilder(),
	}
}

type deploymentBuilder struct {
	deployment *appsv1.Deployment
	pod.PodSpecBuilder
	metadata.MetadataBuilder
}

// SetLabelsAndSelectors sets labels and selectors to Deployment labels, spec.selectors
// and spec.template.labels
func (d *deploymentBuilder) SetSelectors(selectors *metav1.LabelSelector) DeploymentBuilder {
	d.deployment.Spec.Selector = selectors

	return d
}

// SetReplicas sets Deployment replicas
func (d *deploymentBuilder) SetReplicas(r *int32) DeploymentBuilder {
	d.deployment.Spec.Replicas = r
	return d
}

func (d *deploymentBuilder) SetRollingUpdateStrategy(maxUnavailable, maxSurge int) DeploymentBuilder {
	if d.deployment.Spec.Strategy.RollingUpdate == nil {
		d.deployment.Spec.Strategy.RollingUpdate = &appsv1.RollingUpdateDeployment{}
	}

	ms := intstr.FromInt(maxSurge)
	mu := intstr.FromInt(maxUnavailable)
	d.deployment.Spec.Strategy.RollingUpdate.MaxSurge = &ms
	d.deployment.Spec.Strategy.RollingUpdate.MaxUnavailable = &mu

	return d
}

// SetPodTemplateMetadata sets Deployment PodTemplateMetadata for ShardingSphereProxy Pod
func (d *deploymentBuilder) SetPodTemplateMetadata(obj *metav1.ObjectMeta) DeploymentBuilder {
	d.deployment.Spec.Template.ObjectMeta = *obj
	return d
}

// SetPodTemplateSpec sets Deployment PodTemplateSpec for ShardingSphereProxy Pod
func (d *deploymentBuilder) SetPodTemplateSpec(tpl *corev1.PodTemplateSpec) DeploymentBuilder {
	d.deployment.Spec.Template = *tpl
	return d
}

// SetPodTemplateAnnotations sets annotations for ShardingSphereProxy Pod
func (d *deploymentBuilder) SetPodTemplateAnnotations(annotations map[string]string) DeploymentBuilder {
	d.deployment.Spec.Template.Annotations = annotations
	return d
}

// SetPodTemplateLabels sets labels for ShardingSphereProxy Pod
func (d *deploymentBuilder) SetPodTemplateLabels(labels map[string]string) DeploymentBuilder {
	d.deployment.Spec.Template.Labels = labels
	return d
}

// Build returns a Deployment
func (d *deploymentBuilder) BuildDeployment() *appsv1.Deployment {
	d.deployment.ObjectMeta = *d.MetadataBuilder.BuildMetadata()
	return d.deployment
}

// SharedVolumeAndMountBuilder build a Volume which could be mounted by different containers
type SharedVolumeAndMountBuilder interface {
	SetName(name string) SharedVolumeAndMountBuilder
	SetMountPath(idx int, path string) SharedVolumeAndMountBuilder
	SetSubPath(idx int, subpath string) SharedVolumeAndMountBuilder
	SetVolumeMountSize(size int) SharedVolumeAndMountBuilder
	SetVolumeSourceEmptyDir() SharedVolumeAndMountBuilder
	SetVolumeSourceConfigMap(name string, kps ...corev1.KeyToPath) SharedVolumeAndMountBuilder
	Build() (*corev1.Volume, []*corev1.VolumeMount)
}

type sharedVolumeAndMountBuilder struct {
	volume       *corev1.Volume
	volumeMounts []*corev1.VolumeMount
}

// NewSharedVolumeAndMountBuilder returns a new SharedVolumeAndMountBuilder
func NewSharedVolumeAndMountBuilder() SharedVolumeAndMountBuilder {
	return &sharedVolumeAndMountBuilder{
		volume:       &corev1.Volume{},
		volumeMounts: []*corev1.VolumeMount{},
	}
}

// SetName sets Volume and VolumeMounts name
func (b *sharedVolumeAndMountBuilder) SetName(name string) SharedVolumeAndMountBuilder {
	b.volume.Name = name
	for vm := range b.volumeMounts {
		b.volumeMounts[vm].Name = name
	}
	return b
}

// SetVolumeMountSize sets size of VolumeMounts
func (b *sharedVolumeAndMountBuilder) SetVolumeMountSize(size int) SharedVolumeAndMountBuilder {
	if len(b.volumeMounts) != size {
		vms := make([]*corev1.VolumeMount, size)
		for vm := range b.volumeMounts {
			vms[vm] = b.volumeMounts[vm].DeepCopy()
		}
		b.volumeMounts = vms
	}

	for vm := range b.volumeMounts {
		if b.volumeMounts[vm] == nil {
			b.volumeMounts[vm] = &corev1.VolumeMount{}
		}
	}
	return b
}

// SetMountPath sets mountPath of a Volume
func (b *sharedVolumeAndMountBuilder) SetMountPath(idx int, path string) SharedVolumeAndMountBuilder {
	if b.volumeMounts[idx] == nil {
		b.volumeMounts[idx] = &corev1.VolumeMount{}
	}
	b.volumeMounts[idx].MountPath = path
	return b
}

// SetSubPath sets subPath of a Volume
func (b *sharedVolumeAndMountBuilder) SetSubPath(idx int, subpath string) SharedVolumeAndMountBuilder {
	if b.volumeMounts[idx] == nil {
		b.volumeMounts[idx] = &corev1.VolumeMount{}
	}
	b.volumeMounts[idx].SubPath = subpath
	return b
}

// SetVolumeSourceEmptyDir sets a EmptyDir as Volume
func (b *sharedVolumeAndMountBuilder) SetVolumeSourceEmptyDir() SharedVolumeAndMountBuilder {
	if b.volume.EmptyDir == nil {
		b.volume.EmptyDir = &corev1.EmptyDirVolumeSource{}
	}
	return b
}

// SetVolumeSourceConfigMap sets a ConfigMap as Volume
func (b *sharedVolumeAndMountBuilder) SetVolumeSourceConfigMap(name string, kps ...corev1.KeyToPath) SharedVolumeAndMountBuilder {
	if b.volume.ConfigMap == nil {
		b.volume.ConfigMap = &corev1.ConfigMapVolumeSource{}
	}
	b.volume.ConfigMap.LocalObjectReference.Name = name

	if len(kps) > 0 {
		b.volume.ConfigMap.Items = kps
	}
	return b
}

// Build creates a new volume and volumeMounts
func (b *sharedVolumeAndMountBuilder) Build() (*corev1.Volume, []*corev1.VolumeMount) {
	return b.volume, b.volumeMounts
}

// VolumeAndMountBuilder build a Volume and related VolumeMount
type VolumeAndMountBuilder interface {
	SetName(string) VolumeAndMountBuilder
	Build() (*corev1.Volume, *corev1.VolumeMount)
}

// NewVolumeAndMountBuilder returns a VolumeAndMountBuilder
func NewVolumeAndMountBuilder() VolumeAndMountBuilder {
	return &volumeAndMountBuilder{
		volume:      &corev1.Volume{},
		volumemount: &corev1.VolumeMount{},
	}
}

type volumeAndMountBuilder struct {
	volume      *corev1.Volume
	volumemount *corev1.VolumeMount
}

// SetName sets Volume and VolumeMount name
func (b *volumeAndMountBuilder) SetName(name string) VolumeAndMountBuilder {
	b.volume.Name = name
	b.volumemount.Name = name
	return b
}

// SetMountPath sets mountPath of VolumeMount
func (b *volumeAndMountBuilder) SetMountPath(path string) VolumeAndMountBuilder {
	b.volumemount.MountPath = path
	return b
}

// SetSubPath sets subPath of VolumeMount
func (b *volumeAndMountBuilder) SetSubPath(subpath string) VolumeAndMountBuilder {
	b.volumemount.SubPath = subpath
	return b
}

// SetVolumeSourceEmptyDir sets EmptyDir as a Volume
func (b *volumeAndMountBuilder) SetVolumeSourceEmptyDir() VolumeAndMountBuilder {
	b.volume.EmptyDir = &corev1.EmptyDirVolumeSource{}
	return b
}

// SetVolumeSourceConfigMap sets ConfigMap as a Volume
func (b *volumeAndMountBuilder) SetVolumeSourceConfigMap(name string) VolumeAndMountBuilder {
	b.volume.ConfigMap.LocalObjectReference.Name = name
	return b
}

// Build builds a Volume and VolumeMount
func (b *volumeAndMountBuilder) Build() (*corev1.Volume, *corev1.VolumeMount) {
	return b.volume, b.volumemount
}
