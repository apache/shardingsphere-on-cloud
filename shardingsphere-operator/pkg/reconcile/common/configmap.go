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

package common

/*

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ConfigMapFactory interface {
	NewConfigMapBuilderFromGVK(gvk schema.GroupVersionKind) ConfigMapBuilder
}

type configmapFactory struct{}

func (c configmapFactory) NewConfigMapBuilderFromGVK(gvk schema.GroupVersionKind) ConfigMapBuilder {
	if gvk.Group == "shardingsphere.apache.org" {
		if gvk.Kind == "ComputeNode" && gvk.Version == "v1alpha1" {
			return &commonConfigMapBuilder{}
		}
	}

		// if gvk.Group == "shardingsphere.apache.org" && gvk.Kind == "ShardingSphereChaos" && gvk.Version == "v1alpha1" {
		// 	return shardingsphereChaosConfigMapBuilder{}
		// }

	return nil
}

// type computeNodeConfigMapBuilder struct{}

// func (c computeNodeConfigMapFactory) NewConfigMapBuilderFromGVK()

// type shardingsphereChaosConfigMapFactory struct{}

// ConfigMapBuilder generic configmap interface
type ConfigMapBuilder interface {
	SetName(name string) ConfigMapBuilder
	SetNamespace(namespace string) ConfigMapBuilder
	SetLabels(labels map[string]string) ConfigMapBuilder
	SetAnnotations(annos map[string]string) ConfigMapBuilder
	Build() *v1.ConfigMap
}

// commonConfigMapBuilder common configmap implementation
type commonConfigMapBuilder struct {
	configmap *v1.ConfigMap
}

// NewCommonConfigMapBuilder Create a new common configmap builder
func NewCommonConfigMapBuilder(configmap *v1.ConfigMap) ConfigMapBuilder {
	return &commonConfigMapBuilder{configmap}
}

// SetName set the ConfigMap name
func (c *commonConfigMapBuilder) SetName(name string) ConfigMapBuilder {
	c.configmap.Name = name
	return c
}

// SetNamespace set the ConfigMap namespace
func (c *commonConfigMapBuilder) SetNamespace(namespace string) ConfigMapBuilder {
	c.configmap.Namespace = namespace
	return c
}

// SetLabels set the ConfigMap labels
func (c *commonConfigMapBuilder) SetLabels(labels map[string]string) ConfigMapBuilder {
	c.configmap.Labels = labels
	return c
}

// SetAnnotations set the ConfigMap annotations
func (c *commonConfigMapBuilder) SetAnnotations(annos map[string]string) ConfigMapBuilder {
	c.configmap.Annotations = annos
	return c
}

// Build returns a ConfigMap
func (c *commonConfigMapBuilder) Build() *v1.ConfigMap {
	return c.configmap
}

*/
