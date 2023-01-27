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

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewService(c client.Client) Service {
	return serviceClient{
		serviceGetter: serviceGetter{
			Client: c,
		},
		serviceSetter: serviceSetter{
			Client: c,
		},
	}
}

type Service interface {
	ServiceGetter
	ServiceSetter
}

type ServiceGetter interface {
	GetByNamespacedName(context.Context, types.NamespacedName) (*corev1.Service, error)
}

type ServiceSetter interface{}

type serviceClient struct {
	serviceGetter
	serviceSetter
}

type serviceGetter struct {
	client.Client
}

type serviceSetter struct {
	client.Client
}

func (sg serviceGetter) GetByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*corev1.Service, error) {
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
