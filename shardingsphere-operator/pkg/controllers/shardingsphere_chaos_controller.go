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
	sschaosv1alpha1 "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaos"
	reconcile "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/shardingspherechaos"
	chaosv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/go-logr/logr"
	batchV1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ShardingSphereChaosReconciler is a controller for the ShardingSphereChaos
type ShardingSphereChaosReconciler struct { //
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
	Chaos  chaos.Chaos

	//todo: add job definition
	//Job    job.Job
}

// Reconcile handles main function of this controller
func (r *ShardingSphereChaosReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues("ShardingSphereChaos", req.NamespacedName)

	var ssChaos sschaosv1alpha1.ShardingSphereChaos
	if err := r.Get(ctx, req.NamespacedName, &ssChaos); err != nil {
		logger.Error(err, "unable to fetch ShardingSphereChaos source")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if !ssChaos.ObjectMeta.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, nil
	}

	logger.Info("start reconcile chaos")
	if err := r.reconcileChaos(ctx, &ssChaos); err != nil {
		logger.Error(err, " unable to reconcile chaos")
		return ctrl.Result{}, err
	}

	//todo: add inject reconcile and check status here

	if err := r.reconcileStatus(ctx, &ssChaos); err != nil {
		logger.Error(err, "failed to update status")
	}

	return ctrl.Result{}, nil
}

func (r *ShardingSphereChaosReconciler) reconcileChaos(ctx context.Context, ssChao *sschaosv1alpha1.ShardingSphereChaos) error {
	logger := r.Log.WithValues("reconcile chaos", ssChao.Name)
	namespaceName := types.NamespacedName{Namespace: ssChao.Namespace, Name: ssChao.Name}
	switch ssChao.Spec.ChaosKind {
	case sschaosv1alpha1.PodChaosKind:
		chao, isExist, err := r.getPodChaosByNamespacedName(ctx, namespaceName)
		if err != nil {
			logger.Error(err, "pod chaos err")
			return err
		}
		if isExist {
			return r.updatePodChaos(ctx, ssChao, chao)
		}

		return r.CreatePodChaos(ctx, ssChao)
	case sschaosv1alpha1.NetworkChaosKind:
		chao, isExist, err := r.getNetworkChaosByNamespacedName(ctx, namespaceName)
		if err != nil {
			logger.Error(err, "network chao err")
			return err
		}
		if isExist {
			return r.updateNetWorkChaos(ctx, ssChao, chao)
		}
		return r.CreateNetworkChaos(ctx, ssChao)
	}
	return nil
}

func (r *ShardingSphereChaosReconciler) reconcileStatus(ctx context.Context, ssChaos *sschaosv1alpha1.ShardingSphereChaos) error {
	var (
		chaoCondition  sschaosv1alpha1.ChaosCondition
		namespacedName = types.NamespacedName{
			Namespace: ssChaos.Namespace,
			Name:      ssChaos.Name,
		}
	)
	if ssChaos.Spec.ChaosKind == sschaosv1alpha1.PodChaosKind {
		chao, err := r.Chaos.GetPodChaosByNamespacedName(ctx, namespacedName)
		if err != nil {
			return err
		}
		chaoCondition = r.Chaos.ConvertChaosStatus(ctx, ssChaos, chao)
	}

	if ssChaos.Spec.ChaosKind == sschaosv1alpha1.NetworkChaosKind {
		chao, err := r.Chaos.GetNetworkChaosByNamespacedName(ctx, namespacedName)
		if err != nil {
			return err
		}
		chaoCondition = r.Chaos.ConvertChaosStatus(ctx, ssChaos, chao)
	}

	var rt *sschaosv1alpha1.ShardingSphereChaos
	if err := r.Get(ctx, namespacedName, rt); err != nil {
		return err
	}
	ssChaos.Status.ChaosCondition = chaoCondition
	rt.Status = ssChaos.Status
	return r.Status().Update(ctx, rt)
}

func (r *ShardingSphereChaosReconciler) getNetworkChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (reconcile.NetworkChaos, bool, error) {
	nc, err := r.Chaos.GetNetworkChaosByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, false, err
	}
	if nc == nil {
		return nil, false, nil
	}
	return nc, true, nil
}

func (r *ShardingSphereChaosReconciler) getPodChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (reconcile.PodChaos, bool, error) {
	pc, err := r.Chaos.GetPodChaosByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, false, err
	}
	if pc == nil {
		return nil, false, nil
	}
	return pc, true, nil
}

func (r *ShardingSphereChaosReconciler) updatePodChaos(ctx context.Context, chao *sschaosv1alpha1.ShardingSphereChaos, podChaos reconcile.PodChaos) error {
	return r.Chaos.UpdatePodChaos(ctx, chao, podChaos)
}

func (r *ShardingSphereChaosReconciler) CreatePodChaos(ctx context.Context, chao *sschaosv1alpha1.ShardingSphereChaos) error {
	podChaos, err := r.Chaos.NewPodChaos(chao)
	if err != nil {
		return err
	}
	return r.Chaos.CreatePodChaos(ctx, podChaos)
}

func (r *ShardingSphereChaosReconciler) updateNetWorkChaos(ctx context.Context, chao *sschaosv1alpha1.ShardingSphereChaos, netWorkChaos reconcile.NetworkChaos) error {
	return r.Chaos.UpdateNetworkChaos(ctx, chao, netWorkChaos)
}

func (r *ShardingSphereChaosReconciler) CreateNetworkChaos(ctx context.Context, chao *sschaosv1alpha1.ShardingSphereChaos) error {
	networkChaos, err := r.Chaos.NewNetworkPodChaos(chao)
	if err != nil {
		return err
	}
	return r.Chaos.CreateNetworkChaos(ctx, networkChaos)
}

// SetupWithManager sets up the controller with the Manager.
func (r *ShardingSphereChaosReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sschaosv1alpha1.ShardingSphereChaos{}).
		Owns(&chaosv1alpha1.PodChaos{}).
		Owns(&chaosv1alpha1.NetworkChaos{}).
		Owns(&batchV1.Job{}).
		Complete(r)
}
