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

package chaosmesh

import (
	"context"
	"reflect"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"

	chaosmeshv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewChaos creates a new Chaos
func NewChaos(client client.Client) Chaos {
	return chaosClient{
		builder: builder{},
		getter: getter{
			client,
		},
		setter: setter{
			client,
		},
	}
}

// Chaos interface contains setter and getter
type Chaos interface {
	Builder
	Getter
	Setter
}

type chaosClient struct {
	builder
	getter
	setter
}

// Builder build Chaos from different parameters
type Builder interface {
	NewPodChaos(context.Context, *v1alpha1.Chaos) PodChaos
	NewNetworkChaos(context.Context, *v1alpha1.Chaos) NetworkChaos
}

// Getter get Chaos from different parameters
type Getter interface {
	GetPodChaosByNamespacedName(context.Context, types.NamespacedName) (PodChaos, error)
	GetNetworkChaosByNamespacedName(context.Context, types.NamespacedName) (NetworkChaos, error)
}

// Setter set Chaos from different parameters
type Setter interface {
	CreatePodChaos(context.Context, *v1alpha1.Chaos) error
	UpdatePodChaos(context.Context, PodChaos, *v1alpha1.Chaos) error
	DeletePodChaos(context.Context, PodChaos) error

	CreateNetworkChaos(context.Context, *v1alpha1.Chaos) error
	UpdateNetworkChaos(context.Context, NetworkChaos, *v1alpha1.Chaos) error
	DeleteNetworkChaos(context.Context, NetworkChaos) error
}

type getter struct {
	client.Client
}

type PodChaos interface{}

func (cg getter) GetPodChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (PodChaos, error) {
	chaos := &chaosmeshv1alpha1.PodChaos{}
	if err := cg.Get(ctx, namespacedName, chaos); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	} else {
		return chaos, nil
	}
}

type NetworkChaos interface{}

func (cg getter) GetNetworkChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (NetworkChaos, error) {
	chaos := &chaosmeshv1alpha1.NetworkChaos{}
	if err := cg.Get(ctx, namespacedName, chaos); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	} else {
		return chaos, nil
	}
}

type builder struct{}

func (blder builder) NewPodChaos(ctx context.Context, sschaos *v1alpha1.Chaos) PodChaos {
	pc, _ := NewPodChaos(sschaos)
	return pc
}

func (blder builder) NewNetworkChaos(ctx context.Context, sschaos *v1alpha1.Chaos) NetworkChaos {
	nc, _ := NewNetworkChaos(sschaos)
	return nc
}

type setter struct {
	client.Client
}

// CreatePodChaos creates a new pod chaos
func (cs setter) CreatePodChaos(ctx context.Context, sschaos *v1alpha1.Chaos) error {
	pc, err := NewPodChaos(sschaos)
	if err != nil {
		return err
	}
	return cs.Client.Create(ctx, pc.(*chaosmeshv1alpha1.PodChaos))
}

// UpdatePodChaos updates a pod chaos
func (cs setter) UpdatePodChaos(ctx context.Context, podChaos PodChaos, sschaos *v1alpha1.Chaos) error {
	pc, err := NewPodChaos(sschaos)
	if err != nil {
		return err
	}
	s, ok := pc.(*chaosmeshv1alpha1.PodChaos)
	if !ok {
		return ErrConvert
	}
	t, ok := podChaos.(*chaosmeshv1alpha1.PodChaos)
	if !ok {
		return ErrConvert
	}
	if reflect.DeepEqual(s.Spec, t.Spec) {
		return nil
	}
	t.Spec = s.Spec

	return cs.Client.Update(ctx, t)
}

// DeletePodChaos deletes a pod chaos
func (cs setter) DeletePodChaos(ctx context.Context, chao PodChaos) error {
	podChao, ok := chao.(*chaosmeshv1alpha1.PodChaos)
	if !ok {
		return ErrConvert
	}
	if err := cs.Client.Delete(ctx, podChao); err != nil {
		return err
	}

	return nil
}

// CreateNetworkChaos creates a new network chaos
func (cs setter) CreateNetworkChaos(ctx context.Context, sschaos *v1alpha1.Chaos) error {
	nc, err := NewNetworkChaos(sschaos)
	if err != nil {
		return err
	}
	return cs.Client.Create(ctx, nc.(*chaosmeshv1alpha1.NetworkChaos))
}

// UpdateNetworkChaos updates a network chaos
func (cs setter) UpdateNetworkChaos(ctx context.Context, networkChaos NetworkChaos, sschaos *v1alpha1.Chaos) error {
	pc, err := NewNetworkChaos(sschaos)
	if err != nil {
		return err
	}
	s, ok := pc.(*chaosmeshv1alpha1.NetworkChaos)
	if !ok {
		return ErrConvert
	}
	t, ok := networkChaos.(*chaosmeshv1alpha1.NetworkChaos)
	if !ok {
		return ErrConvert
	}
	if reflect.DeepEqual(s.Spec, t.Spec) {
		return nil
	}
	t.Spec = s.Spec

	return cs.Client.Update(ctx, t)
}

func (cs setter) DeleteNetworkChaos(ctx context.Context, chao NetworkChaos) error {
	networkChaos, ok := chao.(*chaosmeshv1alpha1.NetworkChaos)
	if !ok {
		return ErrConvert
	}
	if err := cs.Client.Delete(ctx, networkChaos); err != nil {
		return err
	}

	return nil
}
