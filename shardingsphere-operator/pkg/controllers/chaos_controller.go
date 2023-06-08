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
	"reflect"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaosmesh"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/job"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/pressure"
	sschaos "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/shardingspherechaos"

	"github.com/go-logr/logr"
	batchV1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const (
	ChaosControllerName = "chaos-controller"
	ChaosFinalizerName  = "shardingsphere.apache.org/finalizer"
)

// ChaosReconciler is a controller for the Chaos
type ChaosReconciler struct {
	client.Client

	Scheme    *runtime.Scheme
	Log       logr.Logger
	Events    record.EventRecorder
	ClientSet *clientset.Clientset

	Chaos chaosmesh.Chaos

	Job       job.Job
	ExecCtrls []*ExecCtrl
	ConfigMap configmap.ConfigMap
}

// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=chaos,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=chaos/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=chaos/finalizers,verbs=update

// Reconcile handles main function of this controller
func (r *ChaosReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues(ChaosControllerName, req.NamespacedName)

	ssChaos, err := r.getRuntimeChaos(ctx, req.NamespacedName)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("start reconcile chaos")
	if ssChaos.ObjectMeta.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(ssChaos, ChaosFinalizerName) {
			controllerutil.AddFinalizer(ssChaos, ChaosFinalizerName)
			if err := r.Update(ctx, ssChaos); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else if controllerutil.ContainsFinalizer(ssChaos, ChaosFinalizerName) {
		return r.finalize(ctx, ssChaos)
	}

	var errors []error
	if err := r.reconcileChaos(ctx, ssChaos); err != nil {
		errors = append(errors, err)
		logger.Error(err, "reconcile chaos error")
	}

	if err := r.reconcileStatus(ctx, ssChaos); err != nil {
		errors = append(errors, err)
		logger.Error(err, "failed to update status")
	}

	if len(errors) > 0 {
		return ctrl.Result{Requeue: true}, err
	}
	return ctrl.Result{RequeueAfter: defaultRequeueTime}, nil
}

func (r *ChaosReconciler) reconcileChaos(ctx context.Context, chaos *v1alpha1.Chaos) error {
	logger := r.Log.WithValues("reconcile chaos", fmt.Sprintf("%s/%s", chaos.Namespace, chaos.Name))

	namespacedName := types.NamespacedName{
		Namespace: chaos.Namespace,
		Name:      chaos.Name,
	}

	if chaos.Spec.EmbedChaos.PodChaos != nil {
		if err := r.reconcilePodChaos(ctx, chaos, namespacedName); err != nil {
			logger.Error(err, "reconcile pod chaos error")
			return err
		}
	}

	if chaos.Spec.EmbedChaos.NetworkChaos != nil {
		if err := r.reconcileNetworkChaos(ctx, chaos, namespacedName); err != nil {
			logger.Error(err, "reconcile network chaos error")
			return err
		}
	}

	return nil
}

func (r *ChaosReconciler) reconcileStatus(ctx context.Context, chaos *v1alpha1.Chaos) error {
	cur := chaos.Status.DeepCopy()

	if err := r.updateChaosCondition(ctx, chaos); err != nil {
		return err
	}

	if reflect.DeepEqual(cur, chaos.Status) {
		return nil
	}

	return r.Status().Update(ctx, chaos)
}

func (r *ChaosReconciler) updateChaosCondition(ctx context.Context, chaos *v1alpha1.Chaos) error {
	namespacedName := types.NamespacedName{
		Namespace: chaos.Namespace,
		Name:      chaos.Name,
	}

	if chaos.Spec.EmbedChaos.PodChaos != nil {
		switch chaos.Spec.EmbedChaos.PodChaos.Action {
		case v1alpha1.CPUStress:
			fallthrough
		case v1alpha1.MemoryStress:
			sc, err := r.Chaos.GetStressChaosByNamespacedName(ctx, namespacedName)
			if err != nil {
				return err
			}
			chaos.Status.ChaosCondition = chaosmesh.ConvertChaosStatus(ctx, chaos, sc)
		case v1alpha1.PodFailure:
			fallthrough
		case v1alpha1.PodKill:
			fallthrough
		case v1alpha1.ContainerKill:
			pc, err := r.Chaos.GetPodChaosByNamespacedName(ctx, namespacedName)
			if err != nil {
				return err
			}
			chaos.Status.ChaosCondition = chaosmesh.ConvertChaosStatus(ctx, chaos, pc)
		}

	}

	if chaos.Spec.EmbedChaos.NetworkChaos != nil {
		nc, err := r.Chaos.GetNetworkChaosByNamespacedName(ctx, namespacedName)
		if err != nil {
			return err
		}
		chaos.Status.ChaosCondition = chaosmesh.ConvertChaosStatus(ctx, chaos, nc)
	}

	return nil
}

type ExecCtrl struct {
	cancel   context.CancelFunc
	pressure *pressure.Pressure
}

func makeExecName(namespacedName types.NamespacedName, execType string) string {
	return fmt.Sprintf("%s-%s-%s", namespacedName.Namespace, namespacedName.Name, execType)
}

func (r *ChaosReconciler) getRuntimeChaos(ctx context.Context, name types.NamespacedName) (*v1alpha1.Chaos, error) {
	var rt = &v1alpha1.Chaos{}
	err := r.Get(ctx, name, rt)
	return rt, err
}

// nolint:nestif
func (r *ChaosReconciler) finalize(ctx context.Context, ssChaos *v1alpha1.Chaos) (ctrl.Result, error) {
	namespacedName := types.NamespacedName{
		Namespace: ssChaos.Namespace,
		Name:      ssChaos.Name,
	}
	r.deleteExec(namespacedName)
	if err := r.deleteExternalResources(ctx, ssChaos); err != nil {
		return ctrl.Result{}, err
	}

	controllerutil.RemoveFinalizer(ssChaos, ChaosFinalizerName)
	if err := r.Update(ctx, ssChaos); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ChaosReconciler) deleteExternalResources(ctx context.Context, chao *v1alpha1.Chaos) error {
	nameSpacedName := types.NamespacedName{Namespace: chao.Namespace, Name: chao.Name}
	if chao.Spec.EmbedChaos.PodChaos != nil {
		switch chao.Spec.EmbedChaos.PodChaos.Action {
		case v1alpha1.CPUStress:
			fallthrough
		case v1alpha1.MemoryStress:
			if err := r.deleteStressChaos(ctx, nameSpacedName); err != nil {
				return err
			}
		case v1alpha1.PodFailure:
			fallthrough
		case v1alpha1.PodKill:
			fallthrough
		case v1alpha1.ContainerKill:
			if err := r.deletePodChaos(ctx, nameSpacedName); err != nil {
				return err
			}
		}
		return nil
	}

	if chao.Spec.EmbedChaos.NetworkChaos != nil {
		if err := r.deleteNetworkChaos(ctx, nameSpacedName); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (r *ChaosReconciler) deleteExec(namespacedName types.NamespacedName) {
	steady, chaos := makeExecName(namespacedName, string(sschaos.InSteady)), makeExecName(namespacedName, string(sschaos.InChaos))
	execR := make([]*ExecCtrl, 0, len(r.ExecCtrls))
	for i := range r.ExecCtrls {
		exec := r.ExecCtrls[i].pressure
		if exec.Name == steady || exec.Name == chaos {
			r.ExecCtrls[i].cancel()
			continue
		}
		execR = append(execR, r.ExecCtrls[i])
	}
	r.ExecCtrls = execR
}

func (r *ChaosReconciler) deletePodChaos(ctx context.Context, namespacedName types.NamespacedName) error {
	podchao, err := r.getPodChaosByNamespacedName(ctx, namespacedName)
	if err != nil {
		return err
	}
	if podchao != nil {
		if err := r.Chaos.DeletePodChaos(ctx, podchao); err != nil {
			return err
		}
	}

	return nil
}

func (r *ChaosReconciler) deleteNetworkChaos(ctx context.Context, namespacedName types.NamespacedName) error {
	networkchao, err := r.getNetworkChaosByNamespacedName(ctx, namespacedName)
	if err != nil {
		return err
	}
	if networkchao != nil {
		if err := r.Chaos.DeleteNetworkChaos(ctx, networkchao); err != nil {
			return err
		}
	}

	return nil
}

func (r *ChaosReconciler) deleteStressChaos(ctx context.Context, namespacedName types.NamespacedName) error {
	sc, err := r.getStressChaosByNamespacedName(ctx, namespacedName)
	if err != nil {
		return err
	}
	if sc != nil {
		if err := r.Chaos.DeleteStressChaos(ctx, sc); err != nil {
			return err
		}
	}

	return nil
}

func (r *ChaosReconciler) reconcilePodChaos(ctx context.Context, chaos *v1alpha1.Chaos, namespacedName types.NamespacedName) error {
	switch chaos.Spec.EmbedChaos.PodChaos.Action {
	case v1alpha1.PodFailure:
		fallthrough
	case v1alpha1.ContainerKill:
		fallthrough
	case v1alpha1.PodKill:
		pc, err := r.getPodChaosByNamespacedName(ctx, namespacedName)
		if err != nil {
			return err
		}
		if pc != nil {
			return r.updatePodChaos(ctx, chaos, pc)
		}

		return r.createPodChaos(ctx, chaos)
	case v1alpha1.CPUStress:
		fallthrough
	case v1alpha1.MemoryStress:
		sc, err := r.getStressChaosByNamespacedName(ctx, namespacedName)
		if err != nil {
			return err
		}
		if sc != nil {
			return r.updateStressChaos(ctx, chaos, sc)
		}

		return r.createStressChaos(ctx, chaos)
	}
	return nil
}

func (r *ChaosReconciler) getPodChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (chaosmesh.PodChaos, error) {
	pc, err := r.Chaos.GetPodChaosByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, err
	}
	return pc, nil
}

func (r *ChaosReconciler) createPodChaos(ctx context.Context, chaos *v1alpha1.Chaos) error {
	err := r.Chaos.CreatePodChaos(ctx, chaos)
	if err != nil {
		return err
	}
	r.Events.Event(chaos, "Normal", "Created", fmt.Sprintf("PodChaos %s", " is created successfully"))
	return nil
}

func (r *ChaosReconciler) updatePodChaos(ctx context.Context, chaos *v1alpha1.Chaos, podChaos chaosmesh.PodChaos) error {
	err := r.Chaos.UpdatePodChaos(ctx, podChaos, chaos)
	if err != nil {
		return err
	}

	return nil
}

func (r *ChaosReconciler) reconcileNetworkChaos(ctx context.Context, chaos *v1alpha1.Chaos, namespacedName types.NamespacedName) error {
	nc, err := r.getNetworkChaosByNamespacedName(ctx, namespacedName)
	if err != nil {
		return err
	}
	if nc != nil {
		return r.updateNetWorkChaos(ctx, chaos, nc)
	}

	return r.createNetworkChaos(ctx, chaos)
}

func (r *ChaosReconciler) updateNetWorkChaos(ctx context.Context, chaos *v1alpha1.Chaos, networkChaos chaosmesh.NetworkChaos) error {
	err := r.Chaos.UpdateNetworkChaos(ctx, networkChaos, chaos)
	if err != nil {
		return err
	}

	return nil
}

func (r *ChaosReconciler) createNetworkChaos(ctx context.Context, chaos *v1alpha1.Chaos) error {
	err := r.Chaos.CreateNetworkChaos(ctx, chaos)
	if err != nil {
		return err
	}

	r.Events.Event(chaos, "Normal", "created", fmt.Sprintf("NetworkChaos %s", "  is created successfully"))
	return nil
}

func (r *ChaosReconciler) getNetworkChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (chaosmesh.NetworkChaos, error) {
	nc, err := r.Chaos.GetNetworkChaosByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, err
	}
	return nc, nil
}

func (r *ChaosReconciler) getStressChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (chaosmesh.StressChaos, error) {
	pc, err := r.Chaos.GetStressChaosByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, err
	}
	return pc, nil
}

func (r *ChaosReconciler) createStressChaos(ctx context.Context, chaos *v1alpha1.Chaos) error {
	err := r.Chaos.CreateStressChaos(ctx, chaos)
	if err != nil {
		return err
	}
	r.Events.Event(chaos, "Normal", "Created", fmt.Sprintf("StressChaos %s", " is created successfully"))
	return nil
}

func (r *ChaosReconciler) updateStressChaos(ctx context.Context, chaos *v1alpha1.Chaos, stress chaosmesh.StressChaos) error {
	err := r.Chaos.UpdateStressChaos(ctx, stress, chaos)
	if err != nil {
		return err
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ChaosReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Chaos{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&batchV1.Job{}).
		Complete(r)
}
