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

package service

import (
	"context"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewService creates a new Service
func NewServiceClient(c client.Client) Service {
	return serviceClient{
		builder: builder{},
		getter: getter{
			Client: c,
		},
		setter: setter{
			Client: c,
		},
	}
}

// Service interface contains setter and getter
type Service interface {
	Builder
	Getter
	Setter
}

// Getter get Service from different parameters
type Getter interface {
	GetByNamespacedName(context.Context, types.NamespacedName) (*corev1.Service, error)
}

// Setter set Service from different parameters
type Setter interface {
	Create(context.Context, *corev1.Service) error
	Update(context.Context, *corev1.Service) error
}

// Builder builds a Service
type Builder interface {
	Build(ctx context.Context, cn *v1alpha1.ComputeNode) *corev1.Service
}

type serviceClient struct {
	builder
	getter
	setter
}

type getter struct {
	client.Client
}

// GetByNamespacedName returns a service by its namespaced name
func (sg getter) GetByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*corev1.Service, error) {
	svc := &corev1.Service{}
	if err := sg.Get(ctx, namespacedName, svc); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	} else {
		return svc, nil
	}
}

type setter struct {
	client.Client
}

// Create creates a service
func (sg setter) Create(ctx context.Context, svc *corev1.Service) error {
	return sg.Client.Create(ctx, svc)
}

// Update updates a service
func (sg setter) Update(ctx context.Context, svc *corev1.Service) error {
	return sg.Client.Update(ctx, svc)
}

type builder struct{}

// Build builds a service
func (b builder) Build(ctx context.Context, cn *v1alpha1.ComputeNode) *corev1.Service {
	return NewService(cn)
}
