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
	"fmt"
	"k8s.io/client-go/tools/record"
	"time"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	v1 "k8s.io/api/core/v1"

	sschaosv1alpha1 "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaos"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/job"
	reconcile "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/shardingspherechaos"
	chaosv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/go-logr/logr"
	batchV1 "k8s.io/api/batch/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ShardingSphereChaosControllerName = "shardingsphere-chaos-controller"
	ssChaosDefaultEnqueueTime         = 10 * time.Second
	defaultCreatedMessage             = " is created successfully"
	defaultUpdateMessage              = "new changes updated"
)

// ShardingSphereChaosReconciler is a controller for the ShardingSphereChaos
type ShardingSphereChaosReconciler struct { //
	client.Client
	Scheme    *runtime.Scheme
	Log       logr.Logger
	Chaos     chaos.Chaos
	Job       job.Job
	ConfigMap configmap.ConfigMap
	Events    record.EventRecorder
}

// Reconcile handles main function of this controller
func (r *ShardingSphereChaosReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues(ShardingSphereChaosControllerName, req.NamespacedName)

	ssChaos, err := r.getRuntimeSSChaos(ctx, req.NamespacedName)
	if err != nil {
		return ctrl.Result{}, err
	}

	if !ssChaos.ObjectMeta.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, nil
	}
	logger.Info("start reconcile chaos")
	if err := r.reconcileChaos(ctx, ssChaos); err != nil {
		if err == reconcile.ErrChangedSpec {
			errHandle := r.handleChaosChange(ctx, req.NamespacedName)
			return ctrl.Result{}, errHandle
		}
		logger.Error(err, " unable to reconcile chaos")
		r.Events.Event(ssChaos, "Warning", "chaos err", err.Error())
		return ctrl.Result{}, err
	}
	if err := r.reconcileConfigMap(ctx, ssChaos); err != nil {
		logger.Error(err, "unable to reconcile configmap")
		r.Events.Event(ssChaos, "Warning", "configmap err", err.Error())
		return ctrl.Result{}, err
	}
	if err := r.reconcileJob(ctx, ssChaos); err != nil {
		logger.Error(err, "unable to reconcile job")
		r.Events.Event(ssChaos, "Warning", "job err", err.Error())
		return ctrl.Result{}, err
	}
	if err := r.reconcileStatus(ctx, req.NamespacedName); err != nil {
		r.Events.Event(ssChaos, "Warning", "update status error", err.Error())
		logger.Error(err, "failed to update status")
	}

	return ctrl.Result{RequeueAfter: ssChaosDefaultEnqueueTime}, nil
}

func (r *ShardingSphereChaosReconciler) handleChaosChange(ctx context.Context, name types.NamespacedName) error {

	ssChaos, err := r.getRuntimeSSChaos(ctx, name)
	if err != nil {
		return err
	}
	if ssChaos.Status.Phase != sschaosv1alpha1.PhaseBeforeExperiment {
		ssChaos.Status.Phase = sschaosv1alpha1.PhaseAfterExperiment
		if err := r.Status().Update(ctx, ssChaos); err != nil {
			return err
		}
	}
	return nil
}

func (r *ShardingSphereChaosReconciler) getRuntimeSSChaos(ctx context.Context, name types.NamespacedName) (*sschaosv1alpha1.ShardingSphereChaos, error) {
	var rt = &sschaosv1alpha1.ShardingSphereChaos{}
	err := r.Get(ctx, name, rt)
	return rt, client.IgnoreNotFound(err)
}

func (r *ShardingSphereChaosReconciler) reconcileChaos(ctx context.Context, ssChao *sschaosv1alpha1.ShardingSphereChaos) error {
	logger := r.Log.WithValues("reconcile chaos", ssChao.Name)
	if ssChao.Status.Phase == sschaosv1alpha1.PhaseBeforeExperiment || ssChao.Status.Phase == "" {
		return nil
	}
	namespaceName := types.NamespacedName{Namespace: ssChao.Namespace, Name: ssChao.Name}
	if ssChao.Spec.EmbedChaos.PodChaos != nil {
		chao, isExist, err := r.getPodChaosByNamespacedName(ctx, namespaceName)
		if err != nil {
			logger.Error(err, "pod chaos err")
			return err
		}
		if isExist {
			return r.updatePodChaos(ctx, ssChao, chao)
		}
		return r.CreatePodChaos(ctx, ssChao)
	} else if ssChao.Spec.EmbedChaos.NetworkChaos != nil {
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

func (r *ShardingSphereChaosReconciler) reconcileConfigMap(ctx context.Context, ssChaos *sschaosv1alpha1.ShardingSphereChaos) error {
	logger := r.Log.WithValues("reconcile configmap", ssChaos.Name)
	namespaceName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
	rConfigmap, isExist, err := r.getConfigMapByNamespacedName(ctx, namespaceName)
	if err != nil {
		logger.Error(err, "get configmap error")
		return err
	}

	if isExist {
		return r.updateConfigMap(ctx, ssChaos, rConfigmap)
	}

	return r.CreateConfigMap(ctx, ssChaos)
}

func (r *ShardingSphereChaosReconciler) reconcileJob(ctx context.Context, ssChaos *sschaosv1alpha1.ShardingSphereChaos) error {
	logger := r.Log.WithValues("reconcile job", ssChaos.Name)
	namespaceName := types.NamespacedName{Namespace: ssChaos.Namespace, Name: ssChaos.Name}
	rJob, isExist, err := r.getJobByNamespacedName(ctx, namespaceName)
	if err != nil {
		logger.Error(err, "get job err")
		return err
	}
	var nowInjectRequirement reconcile.InjectRequirement
	if ssChaos.Status.Phase == "" || ssChaos.Status.Phase == sschaosv1alpha1.PhaseBeforeExperiment || ssChaos.Status.Phase == sschaosv1alpha1.PhaseAfterExperiment {
		nowInjectRequirement = reconcile.Experimental
	}
	if ssChaos.Status.Phase == sschaosv1alpha1.PhaseInChaos || ssChaos.Status.Phase == sschaosv1alpha1.PhaseRecoveredChaos {
		nowInjectRequirement = reconcile.Pressure
	}
	if isExist {
		return r.updateJob(ctx, nowInjectRequirement, ssChaos, rJob)
	}

	return r.createJob(ctx, nowInjectRequirement, ssChaos)
}

func (r *ShardingSphereChaosReconciler) reconcileStatus(ctx context.Context, namespacedName types.NamespacedName) error {
	ssChaos, err := r.getRuntimeSSChaos(ctx, namespacedName)
	if err != nil {
		return err
	}
	if ssChaos.Status.Phase == "" {
		ssChaos.Status.Phase = sschaosv1alpha1.PhaseBeforeExperiment
	}
	rJob := &batchV1.Job{}
	if err := r.Get(ctx, namespacedName, rJob); err != nil {
		return err
	}

	if ssChaos.Status.Phase == sschaosv1alpha1.PhaseBeforeExperiment && rJob.Status.Succeeded == 1 {
		ssChaos.Status.Phase = sschaosv1alpha1.PhaseAfterExperiment
	}
	if ssChaos.Status.Phase != sschaosv1alpha1.PhaseBeforeExperiment {
		if ssChaos.Spec.EmbedChaos.PodChaos != nil {
			chao, err := r.Chaos.GetPodChaosByNamespacedName(ctx, namespacedName)
			if err != nil {
				return err
			}
			ssChaos.Status.ChaosCondition = r.Chaos.ConvertChaosStatus(ctx, ssChaos, chao)
		} else if ssChaos.Spec.EmbedChaos.NetworkChaos != nil {
			chao, err := r.Chaos.GetNetworkChaosByNamespacedName(ctx, namespacedName)
			if err != nil {
				return err
			}
			ssChaos.Status.ChaosCondition = r.Chaos.ConvertChaosStatus(ctx, ssChaos, chao)
		}

		if ssChaos.Status.ChaosCondition == sschaosv1alpha1.AllInjected && ssChaos.Status.Phase == sschaosv1alpha1.PhaseAfterExperiment {
			ssChaos.Status.Phase = sschaosv1alpha1.PhaseInChaos
		}

		if ssChaos.Status.ChaosCondition == sschaosv1alpha1.AllRecovered && ssChaos.Status.Phase == sschaosv1alpha1.PhaseInChaos {
			ssChaos.Status.Phase = sschaosv1alpha1.PhaseRecoveredChaos
		}
	}

	rt, err := r.getRuntimeSSChaos(ctx, namespacedName)
	if err != nil {
		return err
	}
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

func (r *ShardingSphereChaosReconciler) getConfigMapByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*v1.ConfigMap, bool, error) {
	config, err := r.ConfigMap.GetByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, false, err
	}
	if config == nil {
		return nil, false, nil
	}

	return config, true, nil
}

func (r *ShardingSphereChaosReconciler) getJobByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*batchV1.Job, bool, error) {
	injectJob, err := r.Job.GetByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, false, err
	}
	if injectJob == nil {
		return nil, false, nil
	}

	return injectJob, true, nil
}

func (r *ShardingSphereChaosReconciler) updateConfigMap(ctx context.Context, chao *sschaosv1alpha1.ShardingSphereChaos, cur *v1.ConfigMap) error {
	exp := reconcile.UpdateConfigMap(chao, cur)
	if exp == nil {
		return nil
	}
	return r.Update(ctx, exp)
}

func (r *ShardingSphereChaosReconciler) CreateConfigMap(ctx context.Context, chao *sschaosv1alpha1.ShardingSphereChaos) error {
	rConfigMap := reconcile.NewSSConfigMap(chao)
	if err := ctrl.SetControllerReference(chao, rConfigMap, r.Scheme); err != nil {
		return err
	}
	err := r.Create(ctx, rConfigMap)
	if err == nil && apierrors.IsAlreadyExists(err) {
		return nil
	}

	return err
}

func (r *ShardingSphereChaosReconciler) updateJob(ctx context.Context, requirement reconcile.InjectRequirement, chao *sschaosv1alpha1.ShardingSphereChaos, cur *batchV1.Job) error {
	exp, err := reconcile.UpdateJob(chao, requirement, cur)
	if err != nil {
		return err
	}
	if exp != nil {
		if err := r.Delete(ctx, cur); err != nil {
			return err
		}
		if err := ctrl.SetControllerReference(chao, exp, r.Scheme); err != nil {
			return err
		}
		if err := r.Create(ctx, exp); err != nil {
			return err
		}
	}
	return nil
}

// todo:
func (r *ShardingSphereChaosReconciler) createJob(ctx context.Context, requirement reconcile.InjectRequirement, chao *sschaosv1alpha1.ShardingSphereChaos) error {
	injectJob, err := reconcile.NewJob(chao, requirement)
	if err := ctrl.SetControllerReference(chao, injectJob, r.Scheme); err != nil {
		return err
	}
	if err != nil {
		return err
	}
	err = r.Create(ctx, injectJob)
	if err == nil && apierrors.IsAlreadyExists(err) {
		return nil
	}

	return err
}

func (r *ShardingSphereChaosReconciler) updatePodChaos(ctx context.Context, chao *sschaosv1alpha1.ShardingSphereChaos, podChaos reconcile.PodChaos) error {
	err := r.Chaos.UpdatePodChaos(ctx, chao, podChaos)
	if err != nil {
		if err == reconcile.ErrNotChanged {
			return nil
		}
		return err
	}
	r.Events.Event(chao, "Normal", "applied", fmt.Sprintf("podChaos %s", defaultUpdateMessage))
	return reconcile.ErrChangedSpec
}

func (r *ShardingSphereChaosReconciler) CreatePodChaos(ctx context.Context, chao *sschaosv1alpha1.ShardingSphereChaos) error {
	podChaos, err := r.Chaos.NewPodChaos(chao)
	if err != nil {
		return err
	}
	err = r.Chaos.CreatePodChaos(ctx, podChaos)
	if err != nil {
		return err
	}
	fmt.Println("phase", chao.Status.Phase)
	r.Events.Event(chao, "Normal", "created", fmt.Sprintf("podChaos %s", defaultCreatedMessage))
	return nil
}

func (r *ShardingSphereChaosReconciler) updateNetWorkChaos(ctx context.Context, chao *sschaosv1alpha1.ShardingSphereChaos, netWorkChaos reconcile.NetworkChaos) error {
	err := r.Chaos.UpdateNetworkChaos(ctx, chao, netWorkChaos)
	if err != nil {
		if err == reconcile.ErrNotChanged {
			return nil
		}
		return err
	}
	r.Events.Event(chao, "Normal", "applied", fmt.Sprintf("networkChaos %s", defaultUpdateMessage))
	return reconcile.ErrChangedSpec
}

func (r *ShardingSphereChaosReconciler) CreateNetworkChaos(ctx context.Context, chao *sschaosv1alpha1.ShardingSphereChaos) error {
	networkChaos, err := r.Chaos.NewNetworkPodChaos(chao)
	if err != nil {
		return err
	}
	err = r.Chaos.CreateNetworkChaos(ctx, networkChaos)
	if err != nil {
		return err
	}

	r.Events.Event(chao, "Normal", "created", fmt.Sprintf("networkChaos %s", defaultCreatedMessage))
	return nil
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
