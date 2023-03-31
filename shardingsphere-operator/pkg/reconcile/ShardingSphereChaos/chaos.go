package ShardingSphereChaos

import (
	"context"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaos"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ChaosHandler interface {
	NewPodChaos(ssChao *v1alpha1.ShardingSphereChaos) chaos.PodChaos
	NewNetworkPodChaos(ssChao *v1alpha1.ShardingSphereChaos) chaos.NetworkChaos
	UpdateNetworkChaos(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos, r client.Client, cur chaos.NetworkChaos) error
	UpdatePodChaos(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos, r client.Client, cur chaos.PodChaos) error
	CreatePodChaos(ctx context.Context, r client.Client, podChao chaos.PodChaos) error
	CreateNetworkChaos(ctx context.Context, r client.Client, networkChao chaos.NetworkChaos) error
}

var ChaosHandle ChaosHandler

func init() {
	//todo: replace to config
	ChaosHandle = &chaosMeshHandler{}
}
