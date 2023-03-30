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
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	reconcile "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/ShardingSphereChaos"
	chaosv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	batchV1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ShardingSphereChaosReconciler is a controller for the ShardingSphereChaos
type ShardingSphereChaosReconciler struct { //
	client.Client
	Scheme *runtime.Scheme
	//todo: add job definition
	//Job    job.Job
}

// Reconcile handles main function of this controller
func (r *ShardingSphereChaosReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.WithValues("ShardingSphereChaos", req.NamespacedName)

	var ssChaos v1alpha1.ShardingSphereChaos
	if err := r.Get(ctx, req.NamespacedName, &ssChaos); err != nil {
		logger.Error(err, "unable to fetch ShardingSphereChaos source")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if !ssChaos.ObjectMeta.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, nil
	}

	if err := r.reconcileChaosMesh(ctx, &ssChaos); err != nil {
		logger.Error(err, " unable to reconcile chaos")
		return ctrl.Result{}, err
	}

	//todo: add inject reconcile and check status here

	return ctrl.Result{}, nil
}

func (r *ShardingSphereChaosReconciler) reconcileChaosMesh(ctx context.Context, ssChao *v1alpha1.ShardingSphereChaos) error {
	constructAndCreateChao := func(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos) error {

		var CreateErr error
		switch ssChaos.Spec.ChaosKind {
		case v1alpha1.NetworkChaosKind:
			CreateErr = r.CreateNetworkChaos(ctx, ssChaos)
		case v1alpha1.PodChaosKind:
			CreateErr = r.CreatePodChaos(ctx, ssChaos)
		case v1alpha1.WorkFlowKind:
			CreateErr = r.CreateWorkFlow(ctx, ssChaos)
		}

		return CreateErr
	}

	if err := constructAndCreateChao(ctx, ssChao); err != nil {
		return err
	}

	return nil
}

func (r *ShardingSphereChaosReconciler) CreateWorkFlow(ctx context.Context, chao *v1alpha1.ShardingSphereChaos) error {
	workflow := reconcile.NewWorkflow(chao)
	if err := r.Create(ctx, workflow); err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	}

	return nil
}

func (r *ShardingSphereChaosReconciler) CreateNetworkChaos(ctx context.Context, chao *v1alpha1.ShardingSphereChaos) error {
	networkChaos := reconcile.NewNetworkPodChaos(chao)
	if err := r.Create(ctx, networkChaos); err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	}

	return nil
}

func (r *ShardingSphereChaosReconciler) CreatePodChaos(ctx context.Context, chao *v1alpha1.ShardingSphereChaos) error {
	podChao := reconcile.NewPodChaos(chao)
	if err := r.Create(ctx, podChao); err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ShardingSphereChaosReconciler) SetupWithManager(mgr ctrl.Manager) error {

	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ShardingSphereChaos{}).
		Owns(&chaosv1alpha1.PodChaos{}).
		Owns(&chaosv1alpha1.NetworkChaos{}).
		Owns(&chaosv1alpha1.Workflow{}).
		Owns(&batchV1.Job{}).
		Complete(r)
}
