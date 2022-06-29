/*
 *   Copyright © 2022，Beijing Sifei Software Technology Co., LTD.
 *   All Rights Reserved.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package controllers

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logger "sigs.k8s.io/controller-runtime/pkg/log"
	"sphere-ex.com/shardingsphere-operator/api/v1alpha1"
	"sphere-ex.com/shardingsphere-operator/pkg/reconcile"
	"time"
)

const (
	//WaitingForReady Time selection reference kubelet restart time
	WaitingForReady = 10 * time.Second
)

// ProxyReconciler reconciles a Proxy object
type ProxyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=shardingsphere.sphere-ex.com,resources=proxies,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=shardingsphere.sphere-ex.com,resources=proxies/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=shardingsphere.sphere-ex.com,resources=proxies/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=deployment/status,verbs=get;list
//+kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=pods/status,verbs=get;list;watch;

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.

func (r *ProxyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.FromContext(ctx)

	run := &v1alpha1.Proxy{}
	err := r.Get(ctx, req.NamespacedName, run)
	if apierrors.IsNotFound(err) {
		log.Info("Resource in work queue no longer exists!")
		return ctrl.Result{}, nil
	} else if err != nil {
		log.Error(err, "Error getting  CRD resource")
		return ctrl.Result{}, err
	}

	runtimeDeployment := &appsv1.Deployment{}
	err = r.Get(ctx, req.NamespacedName, runtimeDeployment)
	if apierrors.IsNotFound(err) {
		cascadingDeployment := reconcile.ConstructCascadingDeployment(run)
		err = r.Create(ctx, cascadingDeployment)
		if err != nil {
			run.SetInitializationFailed()
			_ = r.Status().Update(ctx, run)
			log.Error(err, "Error creating cascaded deployment")
			return ctrl.Result{}, err
		}
	} else if err != nil {
		log.Error(err, "Error getting cascaded deployment")
		return ctrl.Result{}, err
	} else {
		originDeployment := runtimeDeployment.DeepCopy()
		reconcile.UpdateDeployment(run, originDeployment)
		err = r.Update(ctx, originDeployment)
		if err != nil {
			log.Error(err, "Error updating cascaded deployment")
			return ctrl.Result{Requeue: true}, err
		}
	}

	runtimeService := &v1.Service{}
	err = r.Get(ctx, req.NamespacedName, runtimeService)
	if apierrors.IsNotFound(err) {
		cascadingService := reconcile.ConstructCascadingService(run)
		err = r.Create(ctx, cascadingService)
		if err != nil {
			run.SetInitializationFailed()
			_ = r.Status().Update(ctx, run)
			log.Error(err, "Error creating cascaded service")
			return ctrl.Result{}, err
		}
		run.SetInitialized()
		return ctrl.Result{RequeueAfter: WaitingForReady}, nil
	} else if err != nil {
		log.Error(err, "Error getting cascaded service")
		return ctrl.Result{}, err
	} else {
		originService := runtimeService.DeepCopy()
		reconcile.UpdateService(run, originService)
		err = r.Update(ctx, originService)
		if err != nil {
			log.Error(err, "Error updating cascaded service")
			return ctrl.Result{}, err
		}
	}

	podList := &v1.PodList{}
	err = r.List(ctx, podList, client.InNamespace(req.Namespace), client.MatchingLabels(map[string]string{"apps": req.Name}))
	if err != nil {
		log.Error(err, "Error listing cascaded pod")
		return ctrl.Result{}, err
	}

	result := ctrl.Result{}
	readyNodes := reconcile.CountingReadyPods(podList)
	if reconcile.IsRunning(podList) {
		if readyNodes != run.Spec.Replicas {
			result.RequeueAfter = WaitingForReady
			if readyNodes != run.Status.ReadyNodes {
				run.SetPodStarted(readyNodes)
			}
		} else {
			if run.Status.Phase != v1alpha1.StatusReady {
				log.Info("Status is now ready!")
				run.SetReady(readyNodes)
			}
		}
	} else {
		// TODO: Waiting for pods to start exceeds the maximum number of retries
		run.SetPodNotStarted(readyNodes)
		result.RequeueAfter = WaitingForReady
	}

	// TODO: Compare Status with or without modification
	err = r.Status().Update(ctx, run)
	if err != nil {
		log.Error(err, "Error updating status")
		return result, err
	}
	log.Info("RuntimeCRD status ", "status", run.Status)
	return result, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProxyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Proxy{}).
		Owns(&appsv1.Deployment{}).
		Owns(&v1.Service{}).
		Owns(&v1.Pod{}).
		Complete(r)
}
