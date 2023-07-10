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

package cloudnativepg

import (
	"strconv"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"

	cnpgv1 "github.com/cloudnative-pg/cloudnative-pg/api/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// nolint:gocognit
// NewCluster returns a new Cluster
func NewCluster(sn *v1alpha1.StorageNode, sp *v1alpha1.StorageProvider) *cnpgv1.Cluster {
	builder := NewClusterBuilder(sn.GetObjectMeta(), sn.GetObjectKind().GroupVersionKind())
	builder.SetName(sn.Name)
	builder.SetNamespace(sn.Namespace)

	if len(sp.Spec.Parameters["description"]) > 0 {
		builder.SetDescription(sp.Spec.Parameters["description"])
	}

	if len(sp.Spec.Parameters["imageName"]) > 0 {
		builder.SetImageName(sp.Spec.Parameters["imageName"])
	}

	if len(sp.Spec.Parameters["imagePullPolicy"]) > 0 {
		builder.SetImagePullPolicy(sp.Spec.Parameters["imagePullPolicy"])
	}

	if len(sp.Spec.Parameters["postgresUID"]) > 0 {
		pUID, _ := strconv.ParseInt(sp.Spec.Parameters["postgresUID"], 10, 64)
		builder.SetPostgresUID(pUID)
	}

	if len(sp.Spec.Parameters["postgresGID"]) > 0 {
		pGID, _ := strconv.ParseInt(sp.Spec.Parameters["postgresGID"], 10, 64)
		builder.SetPostgresGID(pGID)
	}

	if len(sp.Spec.Parameters["minSyncReplicas"]) > 0 {
		minSyncReplicas, _ := strconv.Atoi(sp.Spec.Parameters["minSyncReplicas"])
		builder.SetMinSyncReplicas(minSyncReplicas)
	}

	if len(sp.Spec.Parameters["maxSyncReplicas"]) > 0 {
		maxSyncReplicas, _ := strconv.Atoi(sp.Spec.Parameters["maxSyncReplicas"])
		builder.SetMaxSyncReplicas(maxSyncReplicas)
	}

	if len(sp.Spec.Parameters["replicaCluster.enabled"]) > 0 {
		enabledrc, _ := strconv.ParseBool(sp.Spec.Parameters["replicaCluster.enabled"])
		builder.SetReplicaClusterEnabled(enabledrc)
	}

	if len(sp.Spec.Parameters["replicaCluster.source"]) > 0 {
		builder.SetReplicaClusterSource(sp.Spec.Parameters["replicaCluster.source"])
	}

	if len(sp.Spec.Parameters["enableSuperuserAccess"]) > 0 {
		enablesa, _ := strconv.ParseBool(sp.Spec.Parameters["enableSuperuserAccess"])
		builder.SetEnableSuperuserAccess(&enablesa)
	}

	if len(sp.Spec.Parameters["superuserSecret"]) > 0 {
		builder.SetSuperuserSecret(sp.Spec.Parameters["superuserSecret"])
	}

	builder.SetCertificates(
		sp.Spec.Parameters["certificates.serverCASecret"],
		sp.Spec.Parameters["certificates.serverTLSSecret"],
		sp.Spec.Parameters["certificates.replicationTLSSecret"],
		sp.Spec.Parameters["certificates.clientCASecret"],
		[]string{sp.Spec.Parameters["certificates.ServerAltDNSNames"]})

	if len(sp.Spec.Parameters["maxStartDelay"]) > 0 {
		maxStartDelay, _ := strconv.ParseInt(sp.Spec.Parameters["maxStartDelay"], 10, 32)
		builder.SetMaxStartDelay(int32(maxStartDelay))
	}

	if len(sp.Spec.Parameters["maxStopDelay"]) > 0 {
		maxStopDelay, _ := strconv.ParseInt(sp.Spec.Parameters["maxStopDelay"], 10, 32)
		builder.SetMaxStopDelay(int32(maxStopDelay))
	}

	if len(sp.Spec.Parameters["maxSwitchoverDelay"]) > 0 {
		maxSwitchoverDelay, _ := strconv.ParseInt(sp.Spec.Parameters["maxSwitchoverDelay"], 10, 32)
		builder.SetMaxSwitchoverDelay(int32(maxSwitchoverDelay))
	}

	if len(sp.Spec.Parameters["failoverDelay"]) > 0 {
		failoverDelay, _ := strconv.ParseInt(sp.Spec.Parameters["failoverDelay"], 10, 32)
		builder.SetFailoverDelay(int32(failoverDelay))
	}

	if len(sp.Spec.Parameters["primaryUpdateStrategy"]) > 0 {
		builder.SetPrimaryUpdateStrategy(sp.Spec.Parameters["primaryUpdateStrategy"])
	}

	if len(sp.Spec.Parameters["primaryUpdateMethod"]) > 0 {
		builder.SetPrimaryUpdateMethod(sp.Spec.Parameters["primaryUpdateMethod"])
	}

	if len(sp.Spec.Parameters["backup.retentionPolicy"]) > 0 {
		builder.SetBackupRetentionPolicy(sp.Spec.Parameters["backup.retentionPolicy"])
	}

	if len(sp.Spec.Parameters["backup.target"]) > 0 {
		builder.SetBackupTarget(sp.Spec.Parameters["backup.target"])
	}

	if len(sp.Spec.Parameters["instances"]) > 0 {
		ins, _ := strconv.Atoi(sp.Spec.Parameters["instances"])
		builder.SetInstances(ins)
	}

	if len(sp.Spec.Parameters["storage.size"]) > 0 {
		builder.SetStorageSize(sp.Spec.Parameters["storage.size"])
	}

	if len(sp.Spec.Parameters["logLevel"]) > 0 {
		builder.SetLogLevel(sp.Spec.Parameters["logLevel"])
	}

	return builder.Build()
}

// NewClusterBuilder returns a ClusterBuilder
func NewClusterBuilder(meta metav1.Object, gvk schema.GroupVersionKind) ClusterBuilder {
	return &clusterBuilder{
		cluster: DefaultCluster(meta, gvk),
	}
}

type ClusterBuilder interface {
	SetName(name string) ClusterBuilder
	SetNamespace(namespace string) ClusterBuilder
	SetDescription(description string) ClusterBuilder
	SetInstances(n int) ClusterBuilder
	SetImageName(name string) ClusterBuilder
	SetImagePullPolicy(p string) ClusterBuilder
	SetPostgresUID(u int64) ClusterBuilder
	SetPostgresGID(g int64) ClusterBuilder
	SetMinSyncReplicas(n int) ClusterBuilder
	SetMaxSyncReplicas(n int) ClusterBuilder
	SetReplicaClusterEnabled(e bool) ClusterBuilder
	SetReplicaClusterSource(s string) ClusterBuilder
	SetSuperuserSecret(s string) ClusterBuilder
	SetEnableSuperuserAccess(e *bool) ClusterBuilder
	SetCertificates(serverCASecret, serverTLSSecret, replicationTLSSecret, clientCASecret string, serverAltDNSNames []string) ClusterBuilder
	SetStorageSize(s string) ClusterBuilder
	SetMaxStartDelay(d int32) ClusterBuilder
	SetMaxStopDelay(d int32) ClusterBuilder
	SetMaxSwitchoverDelay(d int32) ClusterBuilder
	SetFailoverDelay(d int32) ClusterBuilder
	SetPrimaryUpdateStrategy(s string) ClusterBuilder
	SetPrimaryUpdateMethod(m string) ClusterBuilder
	SetBackupRetentionPolicy(r string) ClusterBuilder
	SetBackupTarget(t string) ClusterBuilder
	SetLogLevel(l string) ClusterBuilder
	Build() *cnpgv1.Cluster
}

type clusterBuilder struct {
	cluster *cnpgv1.Cluster
}

// SetName sets the name of the cluster
func (b *clusterBuilder) SetName(name string) ClusterBuilder {
	b.cluster.Name = name
	return b
}

// SetNamespace sets the namespace of the cluster
func (b *clusterBuilder) SetNamespace(namespace string) ClusterBuilder {
	b.cluster.Namespace = namespace
	return b
}

// SetDescription sets the description of the cluster
func (b *clusterBuilder) SetDescription(description string) ClusterBuilder {
	b.cluster.Spec.Description = description
	return b
}

// SetImagePullPolicy sets the image pull policy, default is IfNotPresent
func (b *clusterBuilder) SetImagePullPolicy(p string) ClusterBuilder {
	b.cluster.Spec.ImagePullPolicy = corev1.PullPolicy(p)
	return b
}

// SetPostgresUID sets the UID of the `postgres` user inside the image, defaults to `26`
func (b *clusterBuilder) SetPostgresUID(u int64) ClusterBuilder {
	b.cluster.Spec.PostgresUID = u
	return b
}

// SetPostgresGID sets the GID of the `postgres` user inside the image, defaults to `26`
func (b *clusterBuilder) SetPostgresGID(g int64) ClusterBuilder {
	b.cluster.Spec.PostgresGID = g
	return b
}

// SetInstances sets the number of instances
func (b *clusterBuilder) SetInstances(n int) ClusterBuilder {
	b.cluster.Spec.Instances = n
	return b
}

// SetImageName sets the name of the container image
func (b *clusterBuilder) SetImageName(name string) ClusterBuilder {
	b.cluster.Spec.ImageName = name
	return b
}

// SetMinSyncReplicas sets the minimum number of synchronous replicas
func (b *clusterBuilder) SetMinSyncReplicas(n int) ClusterBuilder {
	b.cluster.Spec.MinSyncReplicas = n
	return b
}

// SetMaxSyncReplicas sets the maximum number of synchronous replicas
func (b *clusterBuilder) SetMaxSyncReplicas(n int) ClusterBuilder {
	b.cluster.Spec.MaxSyncReplicas = n
	return b
}

// SetReplicaClusterEnabled sets whether this cluster is a replica of an existing cluster
func (b *clusterBuilder) SetReplicaClusterEnabled(e bool) ClusterBuilder {
	b.cluster.Spec.ReplicaCluster.Enabled = e
	return b
}

// SetReplicaClusterSource sets the source cluster of this replica
func (b *clusterBuilder) SetReplicaClusterSource(s string) ClusterBuilder {
	b.cluster.Spec.ReplicaCluster.Source = s
	return b
}

// SetSuperuserSecret sets the secret containing the superuser password
func (b *clusterBuilder) SetSuperuserSecret(s string) ClusterBuilder {
	b.cluster.Spec.SuperuserSecret.Name = s
	return b
}

// SetEnableSuperuserAccess sets whether superuser access is enabled
func (b *clusterBuilder) SetEnableSuperuserAccess(e *bool) ClusterBuilder {
	b.cluster.Spec.EnableSuperuserAccess = e
	return b
}

// SetCertificates sets the configuration for the CA and related certificates
func (b *clusterBuilder) SetCertificates(serverCASecret, serverTLSSecret, replicationTLSSecret, clientCASecret string, serverAltDNSNames []string) ClusterBuilder {
	b.cluster.Spec.Certificates.ServerCASecret = serverCASecret
	b.cluster.Spec.Certificates.ServerTLSSecret = serverTLSSecret
	b.cluster.Spec.Certificates.ReplicationTLSSecret = replicationTLSSecret
	b.cluster.Spec.Certificates.ClientCASecret = clientCASecret
	b.cluster.Spec.Certificates.ServerAltDNSNames = serverAltDNSNames
	return b
}

// SetStorageSize sets the storage size of the cluster
func (b *clusterBuilder) SetStorageSize(s string) ClusterBuilder {
	b.cluster.Spec.StorageConfiguration.Size = s
	return b
}

// SetMaxStartDelay sets the time in seconds that is allowed for a PostgreSQL instance to
// successfully start up (default 30)
func (b *clusterBuilder) SetMaxStartDelay(d int32) ClusterBuilder {
	b.cluster.Spec.MaxStartDelay = d
	return b
}

// SetMaxStopDelay sets the time in seconds that is allowed for a PostgreSQL instance to
// gracefully shutdown (default 30)
func (b *clusterBuilder) SetMaxStopDelay(d int32) ClusterBuilder {
	b.cluster.Spec.MaxStopDelay = d
	return b
}

// SetMaxSwitchoverDelay sets the time in seconds that is allowed for a primary PostgreSQL instance
// to gracefully shutdown during a switchover
func (b *clusterBuilder) SetMaxSwitchoverDelay(d int32) ClusterBuilder {
	b.cluster.Spec.MaxSwitchoverDelay = d
	return b
}

// SetFailoverDelay sets the amount of time (in seconds) to wait before triggering a failover after
// the primary PostgreSQL instance in the cluster was detected to be unhealthy
func (b *clusterBuilder) SetFailoverDelay(d int32) ClusterBuilder {
	b.cluster.Spec.FailoverDelay = d
	return b
}

// SetPrimaryUpdateStrategy sets the primary update strategy of the cluster
func (b *clusterBuilder) SetPrimaryUpdateStrategy(s string) ClusterBuilder {
	b.cluster.Spec.PrimaryUpdateStrategy = cnpgv1.PrimaryUpdateStrategy(s)
	return b
}

// SetPrimaryUpdateMethod sets the primary update method of the cluster
func (b *clusterBuilder) SetPrimaryUpdateMethod(m string) ClusterBuilder {
	b.cluster.Spec.PrimaryUpdateMethod = cnpgv1.PrimaryUpdateMethod(m)
	return b
}

// SetBackupRetentionPolicy sets the backup retention policy of the cluster
func (b *clusterBuilder) SetBackupRetentionPolicy(r string) ClusterBuilder {
	b.cluster.Spec.Backup.RetentionPolicy = r
	return b
}

// SetBackupTarget sets the backup target of the cluster
func (b *clusterBuilder) SetBackupTarget(t string) ClusterBuilder {
	b.cluster.Spec.Backup.Target = cnpgv1.BackupTarget(t)
	return b
}

// SetLogLevel sets the instances' log level
func (b *clusterBuilder) SetLogLevel(l string) ClusterBuilder {
	b.cluster.Spec.LogLevel = l
	return b
}

// Build builds the cluster
func (b *clusterBuilder) Build() *cnpgv1.Cluster {
	return b.cluster
}

func DefaultCluster(meta metav1.Object, gvk schema.GroupVersionKind) *cnpgv1.Cluster {
	return &cnpgv1.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "shardingsphere-proxy-cnpg",
			Namespace: "default",
			Labels:    map[string]string{},
		},
		Spec: cnpgv1.ClusterSpec{
			Instances:             3,
			PrimaryUpdateStrategy: cnpgv1.PrimaryUpdateStrategyUnsupervised,
			StorageConfiguration: cnpgv1.StorageConfiguration{
				Size: "1Gi",
			},
		},
	}
}
