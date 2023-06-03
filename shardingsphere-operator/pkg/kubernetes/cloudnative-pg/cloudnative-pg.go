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
	"context"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"

	cnpgv1 "github.com/cloudnative-pg/cloudnative-pg/api/v1"
	dbmeshv1alpha1 "github.com/database-mesh/golang-sdk/kubernetes/api/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CloudNativePG interface contains setter and getter
type CloudNativePG interface {
	Builder
	Getter
	Setter
}

type cloudnativePGClient struct {
	builder
	getter
	setter
}

// Builder build CloudNativePG Cluster from different parameters
type Builder interface {
	NewCluster(context.Context, *v1alpha1.StorageNode, *dbmeshv1alpha1.DatabaseClass) cnpgv1.Cluster
}

// Getter get CloudNativePG Cluster from different parameters
type Getter interface {
	GetCluster(context.Context, types.NamespacedName) (cnpgv1.Cluster, error)
}

// Setter set CloudNativePG Cluster from different parameters
type Setter interface {
	CreateCluster(context.Context, *cnpgv1.Cluster) error
	UpdateCluster(context.Context, *cnpgv1.Cluster) error
	DeleteCluster(context.Context, *cnpgv1.Cluster) error
}

type getter struct {
	client.Client
}

// GetClusterByNamespacedName returns a ClusterNativePG Cluster
func (cg getter) GetClusterByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*cnpgv1.Cluster, error) {
	c := &cnpgv1.Cluster{}
	if err := cg.Get(ctx, namespacedName, c); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return c, nil
}

type builder struct{}

// NewCluster builds a new CloudNative PG Cluster
func (blder builder) NewCluster(ctx context.Context, sn *v1alpha1.StorageNode, dbcalss *dbmeshv1alpha1.DatabaseClass) *cnpgv1.Cluster {
	c, _ := NewCluster()
	return c
}

type setter struct {
	client.Client
}

// CreateCluster creates a new CloudNative PG Cluster
func (cs setter) CreateCluster(ctx context.Context, cluster *cnpgv1.Cluster) error {
	return cs.Client.Create(ctx, cluster)
}

// UpdateCluster updates a existing CloudNative PG Cluster
func (cs setter) UpdateCluster(ctx context.Context, cluster *cnpgv1.Cluster) error {
	return cs.Client.Update(ctx, cluster)
}

// DeleteCluster deletes a existing CloudNative PG Cluster
func (cs setter) DeletePodChaos(ctx context.Context, cluster *cnpgv1.Cluster) error {
	return cs.Client.Delete(ctx, cluster)
}
