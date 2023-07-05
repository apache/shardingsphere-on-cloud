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

package metadata

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MetadataBuilder represents a metadata
type MetadataBuilder interface {
	SetName(name string) MetadataBuilder
	SetNamespace(namespace string) MetadataBuilder
	SetLabels(labels map[string]string) MetadataBuilder
	SetAnnotations(annos map[string]string) MetadataBuilder
	SetOwnerReferences(refs []metav1.OwnerReference) MetadataBuilder
	SetFinalizers(fs []string) MetadataBuilder
	BuildMetadata() *metav1.ObjectMeta
}

type metadataBuilder struct {
	metadata *metav1.ObjectMeta
}

// NewMetadataBuilder creates a new metadata builder
func NewMetadataBuilder() MetadataBuilder {
	return &metadataBuilder{
		metadata: &metav1.ObjectMeta{},
	}
}

// SetName sets the name of the metadata
func (b *metadataBuilder) SetName(name string) MetadataBuilder {
	b.metadata.Name = name
	return b
}

// SetNamespace sets the namespace of the metadata
func (b *metadataBuilder) SetNamespace(namespace string) MetadataBuilder {
	b.metadata.Namespace = namespace
	return b
}

// SetLabels sets the labels of the metadata
func (b *metadataBuilder) SetLabels(labels map[string]string) MetadataBuilder {
	if b.metadata.Labels == nil {
		b.metadata.Labels = map[string]string{}
	}
	b.metadata.Labels = labels
	return b
}

// SetAnnotations set the annotations of the metadata
func (b *metadataBuilder) SetAnnotations(annos map[string]string) MetadataBuilder {
	if b.metadata.Annotations == nil {
		b.metadata.Annotations = map[string]string{}
	}
	b.metadata.Annotations = annos
	return b
}

// SetOwnerReferences sets the owner references of the metadata
func (b *metadataBuilder) SetOwnerReferences(refs []metav1.OwnerReference) MetadataBuilder {
	if b.metadata.OwnerReferences == nil {
		b.metadata.OwnerReferences = []metav1.OwnerReference{}
	}
	b.metadata.OwnerReferences = refs
	return b
}

// SetFinalizer sets the finalizer of the metadata
func (b *metadataBuilder) SetFinalizers(fs []string) MetadataBuilder {
	if b.metadata.Finalizers == nil {
		b.metadata.Finalizers = []string{}
	}
	b.metadata.Finalizers = fs
	return b
}

// Build returns the metadata
func (b *metadataBuilder) BuildMetadata() *metav1.ObjectMeta {
	return b.metadata
}
