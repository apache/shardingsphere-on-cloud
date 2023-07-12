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

package hpa

import (
	"context"

	autoscalingv2 "k8s.io/api/autoscaling/v2"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewHorizontalPodAutoscalerClient creates a new HorizontalPodAutoscaler
func NewHorizontalPodAutoscalerClient(c client.Client) HorizontalPodAutoscaler {
	return hpaClient{
		getter: getter{
			Client: c,
		},
		setter: setter{
			Client: c,
		},
	}
}

// HorizontalPodAutoscaler interface contains setter and getter
type HorizontalPodAutoscaler interface {
	Getter
	Setter
}

type hpaClient struct {
	getter
	setter
}

// Getter get HorizontalPodAutoscaler from different parameters
type Getter interface {
	GetByNamespacedName(context.Context, types.NamespacedName) (*autoscalingv2.HorizontalPodAutoscaler, error)
}

type getter struct {
	client.Client
}

// GetByNamespacedName returns Deployment from given namespaced name
func (dg getter) GetByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*autoscalingv2.HorizontalPodAutoscaler, error) {
	hpa := &autoscalingv2.HorizontalPodAutoscaler{}
	if err := dg.Client.Get(ctx, namespacedName, hpa); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	} else {
		return hpa, nil
	}
}

// Setter get HorizontalPodAutoscaler from different parameters
type Setter interface {
	Create(context.Context, *autoscalingv2.HorizontalPodAutoscaler) error
	Update(context.Context, *autoscalingv2.HorizontalPodAutoscaler) error
}

type setter struct {
	client.Client
}

// Create creates HorizontalPodAutoscaler
func (ds setter) Create(ctx context.Context, dp *autoscalingv2.HorizontalPodAutoscaler) error {
	return ds.Client.Create(ctx, dp)
}

// Update updates HorizontalPodAutoscaler
func (ds setter) Update(ctx context.Context, dp *autoscalingv2.HorizontalPodAutoscaler) error {
	return ds.Client.Update(ctx, dp)
}
