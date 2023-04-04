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
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewChaos(client client.Client) Chaos {
	return ChaosClient{
		ChaosGetter: chaosMeshGetter{client},
		ChaosSetter: chaosMeshSetter{
			sschaos.NewChaosMeshHandler(client),
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

type ChaosGetter interface {
	GetPodChaosByNamespacedName(context.Context, types.NamespacedName) (sschaos.PodChaos, error)
	GetNetworkChaosByNamespacedName(context.Context, types.NamespacedName) (sschaos.NetworkChaos, error)
}

type ChaosSetter interface {
	sschaos.ChaosHandler
}
