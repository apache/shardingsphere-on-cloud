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

	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewDeployment creates a new Deployment
func NewDeployment(c client.Client) Deployment {
	return deploymentClient{
		deploymentGetter: deploymentGetter{
			Client: c,
		},
		deploymentSetter: deploymentSetter{
			Client: c,
		},
	}
}

// Deployment interface contains setter and getter
type Deployment interface {
	DeploymentGetter
	DeploymentSetter
}

type deploymentClient struct {
	deploymentGetter
	deploymentSetter
}

// DeploymentGetter get Deployment from different parameters
type DeploymentGetter interface {
	GetByNamespacedName(context.Context, types.NamespacedName) (*appsv1.Deployment, error)
}

type deploymentGetter struct {
	client.Client
}

// GetByNamespacedName returns Deployment from given namespaced name
func (dg deploymentGetter) GetByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*appsv1.Deployment, error) {
	dp := &appsv1.Deployment{}
	if err := dg.Get(ctx, namespacedName, dp); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	} else {
		return dp, nil
	}
}

// DeploymentMapGetter get Deployment from different parameters
type DeploymentSetter interface {
}

type deploymentSetter struct {
	client.Client
}
