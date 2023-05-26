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

package controllers

import (
	"context"

	shardingspherev1alpha1 "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/go-logr/logr"

	reconcile "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/proxyconfig"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	proxyConfigControllerName = "proxyconfig_controller"
)

// ProxyConfigReconciler reconciles a ShardingSphereProxyServerConfig object
type ProxyConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=shardingsphereproxyserverconfigs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=shardingsphereproxyserverconfigs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=shardingsphereproxyserverconfigs/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ShardingSphereProxyServerConfig object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *ProxyConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues(proxyConfigControllerName, req.NamespacedName)

	run := &shardingspherev1alpha1.ShardingSphereProxyServerConfig{}
	err := r.Get(ctx, req.NamespacedName, run)
	if apierrors.IsNotFound(err) {
		logger.Info("Resource in work queue no longer exists!")
		return ctrl.Result{}, nil
	} else if err != nil {
		logger.Error(err, "Error getting CRD resource")
		return ctrl.Result{}, err
	}

	cm := &v1.ConfigMap{}
	configmap := reconcile.ConstructCascadingConfigmap(run)
	err = r.Get(ctx, req.NamespacedName, cm)
	if apierrors.IsNotFound(err) {
		logger.Info("Creating cascaded configmap")
		err = r.Create(ctx, configmap)
		if err != nil {
			logger.Error(err, "Error creating cascaded configmap")
			return ctrl.Result{}, err
		}
		run.SetMetadataRepository(run.Spec.ClusterConfig.Repository.Type)
		err = r.Status().Update(ctx, run)
		if err != nil {
			logger.Error(err, "Error updating CRD resource status")
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	} else if err != nil {
		logger.Error(err, "Error getting cascaded configmap")
		return ctrl.Result{}, err
	}

	if !equality.Semantic.DeepEqual(configmap.Data, cm.Data) {
		cm = configmap
		logger.Info("Update or correct the configmap")
		err = r.Update(ctx, configmap)
		if err != nil {
			logger.Error(err, "Error updating cascaded configmap")
			// TODO: Error handling for conflict errors alone
			return ctrl.Result{Requeue: true}, err
		}
	}
	if run.Status.MetadataRepository != run.Spec.ClusterConfig.Repository.Type || run.Status.MetadataRepository == "" {
		run.SetMetadataRepository(run.Spec.ClusterConfig.Repository.Type)
		err = r.Status().Update(ctx, run)
		if err != nil {
			logger.Error(err, "Error updating CRD resource status")
			return ctrl.Result{}, err
		}
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProxyConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&shardingspherev1alpha1.ShardingSphereProxyServerConfig{}).
		Owns(&v1.ConfigMap{}).
		Complete(r)
}
