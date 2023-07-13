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

package vpa

import (
	"context"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	autoscalingv1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewVerticalPodAutoscalerClient creates a new VerticalPodAutoscaler
func NewVerticalPodAutoscalerClient(c client.Client) VerticalPodAutoscaler {
	return vpaClient{
		getter: getter{
			Client: c,
		},
		setter: setter{
			Client: c,
		},
	}
}

// VerticalPodAutoscaler interface contains setter and getter
type VerticalPodAutoscaler interface {
	Getter
	Setter
}

type vpaClient struct {
	getter
	setter
}

// Getter get VerticalPodAutoscaler from different parameters
type Getter interface {
	GetByNamespacedName(context.Context, types.NamespacedName) (*autoscalingv1.VerticalPodAutoscaler, error)
}

type getter struct {
	client.Client
}

// GetByNamespacedName returns Deployment from given namespaced name
func (dg getter) GetByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*autoscalingv1.VerticalPodAutoscaler, error) {
	hpa := &autoscalingv1.VerticalPodAutoscaler{}
	if err := dg.Client.Get(ctx, namespacedName, hpa); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	} else {
		return hpa, nil
	}
}

// Setter get VerticalPodAutoscaler from different parameters
type Setter interface {
	Create(context.Context, *autoscalingv1.VerticalPodAutoscaler) error
	Update(context.Context, *autoscalingv1.VerticalPodAutoscaler) error
}

type setter struct {
	client.Client
}

// Create creates VerticalPodAutoscaler
func (ds setter) Create(ctx context.Context, vp *autoscalingv1.VerticalPodAutoscaler) error {
	return ds.Client.Create(ctx, vp)
}

// Update updates VerticalPodAutoscaler
func (ds setter) Update(ctx context.Context, vp *autoscalingv1.VerticalPodAutoscaler) error {
	return ds.Client.Update(ctx, vp)
}
