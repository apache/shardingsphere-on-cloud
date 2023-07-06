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
		pod:             DefauiltPod(),
		MetadataBuilder: metadata.NewMetadataBuilder(),
		PodSpecBuilder:  NewPodSpecBuilder(),
	}
}

func (b *podBuilder) BuildPod() *corev1.Pod {
	b.pod.ObjectMeta = *b.BuildMetadata()
	b.pod.Spec = *b.BuildPodSpec()
	return b.pod
}

func DefauiltPod() *corev1.Pod {
	return &corev1.Pod{
		Spec: corev1.PodSpec{
			InitContainers: []corev1.Container{},
			Containers:     []corev1.Container{},
			Volumes:        []corev1.Volume{},
		},
	}
}

type PodSpecBuilder interface {
	SetVolumes(vs []corev1.Volume) PodSpecBuilder
	AppendVolumes(vs []corev1.Volume) PodSpecBuilder

	SetInitContainers(cs []corev1.Container) PodSpecBuilder
	AppendInitContainers(cs []corev1.Container) PodSpecBuilder

	SetContainers(cs []corev1.Container) PodSpecBuilder
	AppendContainers(cs []corev1.Container) PodSpecBuilder

	SetTerminationGracePeriodSeconds(secs *int64) PodSpecBuilder
	SetImagePullSecrets(secs []corev1.LocalObjectReference) PodSpecBuilder

	BuildPodSpec() *corev1.PodSpec
}

type podSpecBuilder struct {
	spec *corev1.PodSpec
}

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
