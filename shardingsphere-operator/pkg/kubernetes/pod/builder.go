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

package pod

import (
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/container"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/metadata"
	corev1 "k8s.io/api/core/v1"
)

// PodBuilder represents the configuration of a pod
// nolint:unused
type PodBuilder interface {
	metadata.MetadataBuilder
	PodSpecBuilder
	BuildPod() *corev1.Pod
}

// nolint:unused
type podBuilder struct {
	pod *corev1.Pod
	metadata.MetadataBuilder
	PodSpecBuilder
}

// NewPodBuilder returns a new pod builder
func NewPodBuilder() PodBuilder {
	return &podBuilder{
		pod:             DefaultPod(),
		MetadataBuilder: metadata.NewMetadataBuilder(),
		PodSpecBuilder:  NewPodSpecBuilder(),
	}
}

// BuildPod creates a new pod
func (b *podBuilder) BuildPod() *corev1.Pod {
	b.pod.ObjectMeta = *b.BuildMetadata()
	b.pod.Spec = *b.BuildPodSpec()
	return b.pod
}

// DefaultPod returns a default pod
func DefaultPod() *corev1.Pod {
	return &corev1.Pod{
		Spec: corev1.PodSpec{
			InitContainers: []corev1.Container{},
			Containers:     []corev1.Container{},
			Volumes:        []corev1.Volume{},
		},
	}
}

// PodSpecBuilder build PodSpec
type PodSpecBuilder interface {
	SetVolumes(vs []corev1.Volume) PodSpecBuilder
	UpdateVolumeByName(v *corev1.Volume) PodSpecBuilder
	AppendVolumes(vs []corev1.Volume) PodSpecBuilder

	SetInitContainers(cs []corev1.Container) PodSpecBuilder
	UpdateInitContainerByName(c *corev1.Container) PodSpecBuilder
	FindInitContainerByName(name string) container.ContainerBuilder
	AppendInitContainers(cs []corev1.Container) PodSpecBuilder

	SetContainers(cs []corev1.Container) PodSpecBuilder
	UpdateContainerByName(c *corev1.Container) PodSpecBuilder
	FindContainerByName(name string) container.ContainerBuilder
	AppendContainers(cs []corev1.Container) PodSpecBuilder

	SetTerminationGracePeriodSeconds(secs *int64) PodSpecBuilder
	SetImagePullSecrets(secs []corev1.LocalObjectReference) PodSpecBuilder

	BuildPodSpec() *corev1.PodSpec
}

type podSpecBuilder struct {
	spec *corev1.PodSpec
}

// NewPodSpecBuilder returns a PodSpecBuilder
func NewPodSpecBuilder() PodSpecBuilder {
	return &podSpecBuilder{
		spec: &corev1.PodSpec{
			InitContainers: []corev1.Container{},
			Containers:     []corev1.Container{},
			Volumes:        []corev1.Volume{},
		},
	}
}

// SetVolumes sets the volumes
// nolint:unused
func (b *podSpecBuilder) SetVolumes(vs []corev1.Volume) PodSpecBuilder {
	if b.spec.Volumes == nil {
		b.spec.Volumes = []corev1.Volume{}
	}

	b.spec.Volumes = vs
	return b
}

// AppendVolumes append volumes to the container
// nolint:unused
func (b *podSpecBuilder) AppendVolumes(vs []corev1.Volume) PodSpecBuilder {
	if b.spec.Volumes == nil {
		b.spec.Volumes = []corev1.Volume{}
	}

	b.spec.Volumes = append(b.spec.Volumes, vs...)
	return b
}

// UpdateVolumeByName updates the volume by name
func (b *podSpecBuilder) UpdateVolumeByName(vol *corev1.Volume) PodSpecBuilder {
	if b.spec.Volumes == nil {
		b.spec.Volumes = []corev1.Volume{*vol}
	} else {
		for idx := range b.spec.Volumes {
			if b.spec.Volumes[idx].Name == vol.Name {
				b.spec.Volumes[idx] = *vol
				return b
			}
		}
		b.spec.Volumes = append(b.spec.Volumes, *vol)
	}
	return b
}

// SetInintContainers sets the int containers to the container
// nolint:unused
func (b *podSpecBuilder) SetInitContainers(cs []corev1.Container) PodSpecBuilder {
	if b.spec.InitContainers == nil {
		b.spec.InitContainers = cs
	}
	b.spec.InitContainers = cs
	return b
}

// AppendInitContainers append init containers to the container
// nolint:unused
func (b *podSpecBuilder) AppendInitContainers(cs []corev1.Container) PodSpecBuilder {
	if b.spec.InitContainers == nil {
		b.spec.InitContainers = cs
	}
	b.spec.InitContainers = append(b.spec.InitContainers, cs...)
	return b
}

// FindInitContainerByName returns a builder for this init container
func (b *podSpecBuilder) FindInitContainerByName(name string) container.ContainerBuilder {
	for idx := range b.spec.InitContainers {
		if b.spec.InitContainers[idx].Name == name {
			return container.NewContainerBuilderFromContainer(&b.spec.InitContainers[idx])
		}
	}
	return nil
}

// UpdateInitContainerByName will add or update the container with the specified name
func (b *podSpecBuilder) UpdateInitContainerByName(c *corev1.Container) PodSpecBuilder {
	if b.spec.InitContainers == nil {
		b.spec.InitContainers = []corev1.Container{*c}
	} else {
		for idx := range b.spec.InitContainers {
			if b.spec.InitContainers[idx].Name == c.Name {
				b.spec.InitContainers[idx] = *c
				return b
			}
		}
		b.spec.InitContainers = append(b.spec.InitContainers, *c)
	}
	return b
}

// SetContainer set the container to the pod
// nolint:unused
func (b *podSpecBuilder) SetContainers(cs []corev1.Container) PodSpecBuilder {
	if b.spec.Containers == nil {
		b.spec.Containers = cs
	}
	b.spec.Containers = cs
	return b
}

// AppendContainers appends containers to the pod
// nolint:unused
func (b *podSpecBuilder) AppendContainers(cs []corev1.Container) PodSpecBuilder {
	if b.spec.Containers == nil {
		b.spec.Containers = cs
	}
	b.spec.Containers = append(b.spec.Containers, cs...)
	return b
}

// FindContainerByName returns a builder for this container
func (b *podSpecBuilder) FindContainerByName(name string) container.ContainerBuilder {
	for idx := range b.spec.Containers {
		if b.spec.Containers[idx].Name == name {
			return container.NewContainerBuilderFromContainer(&b.spec.Containers[idx])
		}
	}
	return nil
}

// UpdateContainerByName will add or update the container with the specified name
func (b *podSpecBuilder) UpdateContainerByName(c *corev1.Container) PodSpecBuilder {
	if b.spec.Containers == nil {
		b.spec.Containers = []corev1.Container{*c}
	} else {
		for idx := range b.spec.Containers {
			if b.spec.Containers[idx].Name == c.Name {
				b.spec.Containers[idx] = *c
				return b
			}
		}
		b.spec.Containers = append(b.spec.Containers, *c)
	}
	return b
}

// SetImagePullSecrets sets the image pull secrets
// nolint:unused
func (b *podSpecBuilder) SetImagePullSecrets(secs []corev1.LocalObjectReference) PodSpecBuilder {
	if b.spec.ImagePullSecrets == nil {
		b.spec.ImagePullSecrets = []corev1.LocalObjectReference{}
	}

	b.spec.ImagePullSecrets = secs
	return b
}

// SetTerminationGracePeriodSeconds sets the grace period
// nolint:unused
func (b *podSpecBuilder) SetTerminationGracePeriodSeconds(secs *int64) PodSpecBuilder {
	b.spec.TerminationGracePeriodSeconds = secs
	return b
}

// Build returns a PodSpec
func (b *podSpecBuilder) BuildPodSpec() *corev1.PodSpec {
	return b.spec
}
