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

package chaos

import (
	"context"

	sschaos "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/shardingspherechaos"
	chaosV1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type chaosMeshGetter struct {
	client.Client
}

func (cg chaosMeshGetter) GetPodChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (sschaos.PodChaos, error) {
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

func (cg chaosMeshGetter) GetNetworkChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (sschaos.NetworkChaos, error) {
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

type chaosMeshSetter struct {
	sschaos.ChaosHandler
}
