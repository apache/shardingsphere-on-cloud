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
	appsv1 "k8s.io/api/apps/v1"
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

// Build returns a new Deployment
func (db builder) Build(ctx context.Context, cn *v1alpha1.ComputeNode) *appsv1.Deployment {
	return NewDeployment(cn)
}
