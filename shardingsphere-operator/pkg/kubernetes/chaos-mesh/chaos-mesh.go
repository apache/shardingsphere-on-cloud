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

package chaos_mesh

import (
	"context"
	chaosV1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewJob creates a new Job
func NewChaos(c client.Client) Chaos {
	return ChaosClient{
		ChaosGetter: chaosGetter{
			Client: c,
		},
		ChaosSetter: chaosSetter{
			Client: c,
		},
	}
}

// Chaos interface contains setter and getter
type Chaos interface {
	ChaosGetter
	ChaosSetter
}

type ChaosClient struct {
	ChaosGetter
	ChaosSetter
}

// JobGetter get Job from different parameters
type ChaosGetter interface {
	GetPodChaosByNamespacedName(context.Context, types.NamespacedName) (*chaosV1alpha1.PodChaos, error)
	GetNetworkChaosByNamespacedName(context.Context, types.NamespacedName) (*chaosV1alpha1.NetworkChaos, error)
	GetWorkflowByNamespacedName(context.Context, types.NamespacedName) (*chaosV1alpha1.Workflow, error)
}

type chaosGetter struct {
	client.Client
}

func (cg chaosGetter) GetPodChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*chaosV1alpha1.PodChaos, error) {
	dp := &chaosV1alpha1.PodChaos{}
	if err := cg.Get(ctx, namespacedName, dp); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	} else {
		return dp, nil
	}
}

func (cg chaosGetter) GetNetworkChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*chaosV1alpha1.NetworkChaos, error) {
	dp := &chaosV1alpha1.NetworkChaos{}
	if err := cg.Get(ctx, namespacedName, dp); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	} else {
		return dp, nil
	}
}

func (cg chaosGetter) GetWorkflowByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*chaosV1alpha1.Workflow, error) {
	dp := &chaosV1alpha1.Workflow{}
	if err := cg.Get(ctx, namespacedName, dp); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	} else {
		return dp, nil
	}
}

// JobMapGetter get Job from different parameters
type ChaosSetter interface {
}

type chaosSetter struct {
	client.Client
}
