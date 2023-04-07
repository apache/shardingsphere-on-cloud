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

package shardingspherechaos

import (
	"context"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
)

type ChaosHandler interface {
	NewPodChaos(ssChao *v1alpha1.ShardingSphereChaos) (PodChaos, error)
	NewNetworkPodChaos(ssChao *v1alpha1.ShardingSphereChaos) (NetworkChaos, error)
	UpdateNetworkChaos(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos, cur NetworkChaos) error
	UpdatePodChaos(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos, cur PodChaos) error
	CreatePodChaos(ctx context.Context, podChao PodChaos) error
	CreateNetworkChaos(ctx context.Context, networkChao NetworkChaos) error
	ConvertChaosStatus(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos, chaos GenericChaos) *v1alpha1.ChaosCondition
}

type GenericChaos interface {
}

type PodChaos interface {
	GenericChaos
}

type NetworkChaos interface {
	GenericChaos
}
