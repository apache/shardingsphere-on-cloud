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
	"time"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	reconcile "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/proxy"
	"github.com/go-logr/logr"

	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	logger "sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	//WaitingForReady Time selection reference kubelet restart time
	WaitingForReady = 10 * time.Second

	proxyControllerName = "proxy_controller"
)

// ProxyReconciler reconciles a ShardingSphereProxy object
type ProxyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

//+kubebuilder:rbac:groups=shardingsphere.apache.org,resources=proxies,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=shardingsphere.apache.org,resources=proxies/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=shardingsphere.apache.org,resources=proxies/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=deployment/status,verbs=get;list
//+kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=pods/status,verbs=get;list;watch;
//+kubebuilder:rbac:groups=autoscaling,resources=horizontalpodautoscalers,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.

func (r *ProxyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues(proxyControllerName, req.NamespacedName)

	rt, err := r.getRuntimeShardingSphereProxy(ctx, req.NamespacedName)
	if apierrors.IsNotFound(err) {
		logger.Info("Resource in work queue no longer exists!")
		return ctrl.Result{}, nil
	} else if err != nil {
		logger.Error(err, "Error getting CRD resource")
		return ctrl.Result{}, err
	}

	return r.reconcile(ctx, req, rt)
}

func (r *ProxyReconciler) getRuntimeShardingSphereProxy(ctx context.Context, namespacedName types.NamespacedName) (*v1alpha1.ShardingSphereProxy, error) {
	rt := &v1alpha1.ShardingSphereProxy{}
	err := r.Get(ctx, namespacedName, rt)
	return rt, err
}

func (r *ProxyReconciler) reconcile(ctx context.Context, req ctrl.Request, rt *v1alpha1.ShardingSphereProxy) (ctrl.Result, error) {
	log := logger.FromContext(ctx)
	if res, err := r.reconcileDeployment(ctx, req.NamespacedName); err != nil {
		log.Error(err, "Error reconcile Deployment")
		return res, err
	}

	if res, err := r.reconcileService(ctx, req.NamespacedName); err != nil {
		log.Error(err, "Error reconcile Service")
		return res, err
	}
	if res, err := r.reconcilePodList(ctx, req.Namespace, req.Name); err != nil {
		log.Error(err, "Error reconcile Pod list")
		return res, err
	}

	if res, err := r.reconcileHPA(ctx, req.NamespacedName); err != nil {
		log.Error(err, "Error reconcile HPA")
		return res, err
	}

	return ctrl.Result{RequeueAfter: WaitingForReady}, nil
}

func (r *ProxyReconciler) reconcileDeployment(ctx context.Context, namespacedName types.NamespacedName) (ctrl.Result, error) {
	proxy, err := r.getRuntimeShardingSphereProxy(ctx, namespacedName)
	if err != nil {
		return ctrl.Result{}, err
	}

	deploy := &appsv1.Deployment{}
	err = r.Get(ctx, namespacedName, deploy)

	if apierrors.IsNotFound(err) {
		exp := reconcile.NewDeployment(proxy)
		if err := r.Create(ctx, exp); err != nil {
			return ctrl.Result{}, err
		}
	}

	if err != nil {
		return ctrl.Result{}, err
	}

	act := deploy.DeepCopy()
	exp := reconcile.UpdateDeployment(proxy, act)

	if err := r.Update(ctx, exp); err != nil {
		return ctrl.Result{Requeue: true}, err
	}
	return ctrl.Result{}, nil
}

func (r *ProxyReconciler) reconcileHPA(ctx context.Context, namespacedName types.NamespacedName) (ctrl.Result, error) {
	proxy, err := r.getRuntimeShardingSphereProxy(ctx, namespacedName)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Get the HPA
	hpa := &autoscalingv2beta2.HorizontalPodAutoscaler{}
	err = r.Get(ctx, namespacedName, hpa)

	// If the HPA doesn't exist, create it
	if apierrors.IsNotFound(err) {
		if proxy.Spec.AutomaticScaling != nil && proxy.Spec.AutomaticScaling.Enable {
			exp := reconcile.NewHPA(proxy)
			if err := r.Create(ctx, exp); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if err != nil {
		return ctrl.Result{}, err
	}

	// If the HPA exists, but we don't want it, delete it
	if proxy.Spec.AutomaticScaling == nil || !proxy.Spec.AutomaticScaling.Enable {
		if err := r.Delete(ctx, hpa); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	// If the HPA exists and we want it, update it
	act := hpa.DeepCopy()
	exp := reconcile.UpdateHPA(proxy, act)
	if err := r.Update(ctx, exp); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ProxyReconciler) reconcileService(ctx context.Context, namespacedName types.NamespacedName) (ctrl.Result, error) {
	ssproxy, err := r.getRuntimeShardingSphereProxy(ctx, namespacedName)
	if err != nil {
		return ctrl.Result{}, err
	}

	service := &v1.Service{}
	err = r.Get(ctx, namespacedName, service)

	if apierrors.IsNotFound(err) {
		exp := reconcile.NewService(ssproxy)
		if err := r.Create(ctx, exp); err != nil {
			return ctrl.Result{}, err
		}
	}

	if err != nil {
		return ctrl.Result{}, err
	}

	act := service.DeepCopy()
	exp := reconcile.UpdateService(ssproxy, act)
	if err := r.Update(ctx, exp); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ProxyReconciler) reconcilePodList(ctx context.Context, namespace, name string) (ctrl.Result, error) {
	podList := &v1.PodList{}
	if err := r.List(ctx, podList, client.InNamespace(namespace), client.MatchingLabels(map[string]string{"apps": name})); err != nil {
		return ctrl.Result{}, err
	}

	result := ctrl.Result{}

	rt, err := r.getRuntimeShardingSphereProxy(ctx, types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	})
	if err != nil {
		return ctrl.Result{}, err
	}

	rt.Status = reconcile.ReconcileStatus(podList, rt)

	// TODO: Compare Status with or without modification
	if err := r.Status().Update(ctx, rt); err != nil {
		return result, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProxyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ShardingSphereProxy{}).
		Owns(&appsv1.Deployment{}).
		Owns(&v1.Service{}).
		Owns(&v1.Pod{}).
		Complete(r)
}
