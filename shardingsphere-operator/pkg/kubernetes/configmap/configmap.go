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
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewConfigMap(c client.Client) ConfigMap {
	return configmapClient{
		configmapGetter: configmapGetter{
			Client: c,
		},
		configmapSetter: configmapSetter{
			Client: c,
		},
	}
}

type ConfigMap interface {
	ConfigMapGetter
	ConfigMapSetter
}

type ConfigMapGetter interface {
	GetByNamespacedName(context.Context, types.NamespacedName) (*corev1.ConfigMap, error)
}

type ConfigMapSetter interface {
}

type configmapClient struct {
	configmapGetter
	configmapSetter
}

type configmapGetter struct {
	client.Client
}

func (cg configmapGetter) GetByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*corev1.ConfigMap, error) {
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

type configmapSetter struct {
	client.Client
}
