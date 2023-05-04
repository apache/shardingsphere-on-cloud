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

package configmap

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ConfigMapFactory interface {
	NewConfigMapBuilder() ConfigMapBuilder
}

func NewConfigMapFactory(obj runtime.Object) ConfigMapFactory {
	return &configmapFactory{
		// gvk: gvk,
		obj: obj,
	}
}

type configmapFactory struct {
	// gvk schema.GroupVersionKind
	obj runtime.Object
}

func (c *configmapFactory) NewConfigMapBuilder() ConfigMapBuilder {
	gvk := c.obj.GetObjectKind().GroupVersionKind()

	if gvk.Group == "shardingsphere.apache.org" {
		if gvk.Kind == "ComputeNode" && gvk.Version == "v1alpha1" {
			return &computeNodeConfigMapBuilder{
				obj: c.obj,
			}
		}

		if gvk.Kind == "ShardingSphereChaos" && gvk.Version == "v1alpha1" {
			return &shardingsphereChaosConfigMapBuilder{
				obj: c.obj,
			}
		}
	}

	return &configMapBuilder{}
}

// ConfigMapBuilder generic configmap interface
type ConfigMapBuilder interface {
	SetName(name string) ConfigMapBuilder
	SetNamespace(namespace string) ConfigMapBuilder
	SetLabels(labels map[string]string) ConfigMapBuilder
	SetAnnotations(annos map[string]string) ConfigMapBuilder
	SetBinaryData(binaryData map[string][]byte) ConfigMapBuilder
	SetData(data map[string]string) ConfigMapBuilder
	Build() *v1.ConfigMap
}

// configMapBuilder common configmap implementation
type configMapBuilder struct {
	configmap *v1.ConfigMap
}

// NewConfigMapBuilder Create a new common configmap builder
func NewConfigMapBuilder(configmap *v1.ConfigMap) ConfigMapBuilder {
	return &configMapBuilder{configmap}
}

// SetName set the ConfigMap name
func (c *configMapBuilder) SetName(name string) ConfigMapBuilder {
	c.configmap.Name = name
	return c
}

// SetNamespace set the ConfigMap namespace
func (c *configMapBuilder) SetNamespace(namespace string) ConfigMapBuilder {
	c.configmap.Namespace = namespace
	return c
}

// SetLabels set the ConfigMap labels
func (c *configMapBuilder) SetLabels(labels map[string]string) ConfigMapBuilder {
	c.configmap.Labels = labels
	return c
}

// SetAnnotations set the ConfigMap annotations
func (c *configMapBuilder) SetAnnotations(annos map[string]string) ConfigMapBuilder {
	c.configmap.Annotations = annos
	return c
}

func (c *configMapBuilder) SetData(data map[string]string) ConfigMapBuilder {
	c.configmap.Data = data
	return c
}

func (c *configMapBuilder) SetBinaryData(binary map[string][]byte) ConfigMapBuilder {
	c.configmap.BinaryData = binary
	return c
}

// Build returns a ConfigMap
func (c *configMapBuilder) Build() *v1.ConfigMap {
	return c.configmap
}
