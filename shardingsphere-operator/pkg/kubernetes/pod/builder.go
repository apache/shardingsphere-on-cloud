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
	// appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodBuilder represents the configuration of a pod
// nolint:unused
type PodBuilder interface {
	SetName(name string) PodBuilder
	SetNamespace(namespace string) PodBuilder
	SetLabels(labels map[string]string) PodBuilder
	SetAnnotations(annos map[string]string) PodBuilder

	SetVolumes(vs []corev1.Volume) PodBuilder
	AppendVolumes(vs []corev1.Volume) PodBuilder

	SetInitContainers(cs []corev1.Container) PodBuilder
	AppendInitContainers(cs []corev1.Container) PodBuilder

	SetContainers(cs []corev1.Container) PodBuilder
	AppendContainers(cs []corev1.Container) PodBuilder

	SetTerminationGracePeriodSeconds(secs *int64) PodBuilder
	SetImagePullSecrets(secs []corev1.LocalObjectReference) PodBuilder
}

// nolint:unused
type podBuilder struct {
	pod *corev1.Pod
}

// SetName sets the name of the pod
// nolint:unused
func (b *podBuilder) SetName(name string) PodBuilder {
	b.pod.Name = name
	return b
}

// SetNamespace sets the namespace of the pod
// nolint:unused
func (b *podBuilder) SetNamespace(namespace string) PodBuilder {
	b.pod.Namespace = namespace
	return b
}

// SetLabels sets the labels of the pod
// nolint:unused
func (b *podBuilder) SetLabels(labels map[string]string) PodBuilder {
	if b.pod.Labels == nil {
		b.pod.Labels = map[string]string{}
	}
	b.pod.Labels = labels
	return b
}

// SetAnnotations set the annotations of the pod
// nolint:unused
func (b *podBuilder) SetAnnotations(annos map[string]string) PodBuilder {
	if b.pod.Annotations == nil {
		b.pod.Annotations = map[string]string{}
	}
	b.pod.Annotations = annos
	return b
}

// SetVolumes sets the volumes
// nolint:unused
func (b *podBuilder) SetVolumes(vs []corev1.Volume) PodBuilder {
	if b.pod.Spec.Volumes == nil {
		b.pod.Spec.Volumes = []corev1.Volume{}
	}

	b.pod.Spec.Volumes = vs
	return b
}

// AppendVolumes append volumes to the container
// nolint:unused
func (b *podBuilder) AppendVolumes(vs []corev1.Volume) PodBuilder {
	if b.pod.Spec.Volumes == nil {
		b.pod.Spec.Volumes = []corev1.Volume{}
	}

	b.pod.Spec.Volumes = append(b.pod.Spec.Volumes, vs...)
	return b
}

// SetInintContainers sets the int containers to the container
// nolint:unused
func (b *podBuilder) SetInitContainers(cs []corev1.Container) PodBuilder {
	if b.pod.Spec.InitContainers == nil {
		b.pod.Spec.InitContainers = cs
	}
	b.pod.Spec.InitContainers = cs
	return b
}

// AppendInitContainers append init containers to the container
// nolint:unused
func (b *podBuilder) AppendInitContainers(cs []corev1.Container) PodBuilder {
	if b.pod.Spec.InitContainers == nil {
		b.pod.Spec.InitContainers = cs
	}
	b.pod.Spec.InitContainers = append(b.pod.Spec.InitContainers, cs...)
	return b
}

// SetContainer set the container to the pod
// nolint:unused
func (b *podBuilder) SetContainers(cs []corev1.Container) PodBuilder {
	if b.pod.Spec.Containers == nil {
		b.pod.Spec.Containers = cs
	}
	b.pod.Spec.Containers = cs
	return b
}

// AppendContainers appends containers to the pod
// nolint:unused
func (b *podBuilder) AppendContainers(cs []corev1.Container) PodBuilder {
	if b.pod.Spec.Containers == nil {
		b.pod.Spec.Containers = cs
	}
	b.pod.Spec.Containers = append(b.pod.Spec.Containers, cs...)
	return b
}

// SetImagePullSecrets sets the image pull secrets
// nolint:unused
func (b *podBuilder) SetImagePullSecrets(secs []corev1.LocalObjectReference) PodBuilder {
	if b.pod.Spec.ImagePullSecrets == nil {
		b.pod.Spec.ImagePullSecrets = []corev1.LocalObjectReference{}
	}

	b.pod.Spec.ImagePullSecrets = secs
	return b
}

// SetTerminationGracePeriodSeconds sets the grace period
// nolint:unused
func (b *podBuilder) SetTerminationGracePeriodSeconds(secs *int64) PodBuilder {
	b.pod.Spec.TerminationGracePeriodSeconds = secs
	return b
}
