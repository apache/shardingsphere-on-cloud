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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// NewCluster returns a new Cluster
func NewCluster(sn *v1alpha1.StorageNode, sp *v1alpha1.StorageProvider) *cnpgv1.Cluster {
	builder := NewClusterBuilder(sn.GetObjectMeta(), sn.GetObjectKind().GroupVersionKind())
	builder.SetName(sn.Name)
	builder.SetNamespace(sn.Namespace)

	if len(sp.Spec.Parameters["imageName"]) > 0 {
		builder.SetImageName(sp.Spec.Parameters["imageName"])
	}

	if len(sp.Spec.Parameters["backup.retentionPolicy"]) > 0 {
		builder.SetBackupRetentionPolicy(sp.Spec.Parameters["backup.retentionPolicy"])
	}

	if len(sp.Spec.Parameters["backup.target"]) > 0 {
		builder.SetBackupTarget(cnpgv1.BackupTarget(sp.Spec.Parameters["backup.target"]))
	}

	if len(sp.Spec.Parameters["instances"]) > 0 {
		ins, _ := strconv.Atoi(sp.Spec.Parameters["instances"])
		builder.SetInstances(ins)
	}

	if len(sp.Spec.Parameters["storage.size"]) > 0 {
		builder.SetStorageSize(sp.Spec.Parameters["storage.size"])
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
	SetInstances(n int) ClusterBuilder
	SetImageName(name string) ClusterBuilder
	SetStorageSize(s string) ClusterBuilder
	SetBackupRetentionPolicy(r string) ClusterBuilder
	SetBackupTarget(t cnpgv1.BackupTarget) ClusterBuilder
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

// SetStorageSize sets the storage size of the cluster
func (b *clusterBuilder) SetStorageSize(s string) ClusterBuilder {
	b.cluster.Spec.StorageConfiguration.Size = s
	return b
}

// SetBackupRetentionPolicy sets the backup retention policy of the cluster
func (b *clusterBuilder) SetBackupRetentionPolicy(r string) ClusterBuilder {
	b.cluster.Spec.Backup.RetentionPolicy = r
	return b
}

// SetBackupTarget sets the backup target of the cluster
func (b *clusterBuilder) SetBackupTarget(t cnpgv1.BackupTarget) ClusterBuilder {
	b.cluster.Spec.Backup.Target = t
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
