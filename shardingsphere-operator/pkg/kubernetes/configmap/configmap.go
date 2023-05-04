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
	"context"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewConfigMapClient returns a new ConfigMap client
func NewConfigMapClient(c client.Client) ConfigMap {
	return configmapClient{
		builder: builder{},
		getter: getter{
			Client: c,
		},
		setter: setter{
			Client: c,
		},
	}
}

// ConfigMap interface contains setter and getter
type ConfigMap interface {
	Builder
	Getter
	Setter
}

// Getter get ConfigMap from different parameters
type Getter interface {
	GetByNamespacedName(context.Context, types.NamespacedName) (*corev1.ConfigMap, error)
}

// Setter set ConfigMap from different parameters
type Setter interface {
	Create(context.Context, *corev1.ConfigMap) error
	Update(context.Context, *corev1.ConfigMap) error
}

// Builder build ConfigMap from given ComputeNode
type Builder interface {
	// Build(context.Context, *v1alpha1.ComputeNode) *corev1.ConfigMap
	Build(context.Context, runtime.Object) *corev1.ConfigMap
}

type configmapClient struct {
	builder
	getter
	setter
}

type getter struct {
	client.Client
}

// GetByNamespacedName returns ConfigMap from given namespaced name
func (cg getter) GetByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*corev1.ConfigMap, error) {
	cm := &corev1.ConfigMap{}
	if err := cg.Get(ctx, namespacedName, cm); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	} else {
		return cm, err
	}
}

type setter struct {
	client.Client
}

// Create creates ConfigMap
func (cs setter) Create(ctx context.Context, cm *corev1.ConfigMap) error {
	return cs.Client.Create(ctx, cm)
}

// Update updates ConfigMap
func (cs setter) Update(ctx context.Context, cm *corev1.ConfigMap) error {
	return cs.Client.Update(ctx, cm)
}

type builder struct{}

// Build returns a ConfigMap
func (b builder) Build(ctx context.Context, obj runtime.Object) *corev1.ConfigMap {
	return NewConfigMap(obj)
}

// func (b builder) Build(ctx context.Context, cn *v1alpha1.ComputeNode) *corev1.ConfigMap {
// 	return NewComputeNodeConfigMap(cn)
// }
