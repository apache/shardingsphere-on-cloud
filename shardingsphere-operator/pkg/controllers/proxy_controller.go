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
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile"

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
	//miniReadyCount Minimum number of replicas that can be served
	miniReadyCount = 1
)

// ProxyReconciler reconciles a ShardingSphereProxy object
type ProxyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
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
	log := logger.FromContext(ctx)

	rt, err := r.getRuntimeShardingSphereProxy(ctx, req.NamespacedName)
	if apierrors.IsNotFound(err) {
		log.Info("Resource in work queue no longer exists!")
		return ctrl.Result{}, nil
	} else if err != nil {
		log.Error(err, "Error getting CRD resource")
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
	if res, err := r.reconcileDeployment(ctx, req.NamespacedName, rt); err != nil {
		log.Error(err, "Error reconcile Deployment")
		return res, err
	}

	if res, err := r.reconcileService(ctx, req.NamespacedName, rt); err != nil {
		log.Error(err, "Error reconcile Service")
		return res, err
	}
	if res, err := r.reconcilePodList(ctx, req.Namespace, req.Name, rt); err != nil {
		log.Error(err, "Error reconcile Pod list")
		return res, err
	}

	if res, err := r.reconcileHPA(ctx, req.NamespacedName, rt); err != nil {
		log.Error(err, "Error reconcile HPA")
		return res, err
	}

	return ctrl.Result{}, nil
}

func (r *ProxyReconciler) reconcileDeployment(ctx context.Context, namespacedName types.NamespacedName, ssproxy *v1alpha1.ShardingSphereProxy) (ctrl.Result, error) {
	deploy := &appsv1.Deployment{}

	var err error
	if err = r.Get(ctx, namespacedName, deploy); err != nil {
		if !apierrors.IsNotFound(err) {
			return ctrl.Result{}, err
		} else {
			exp := reconcile.NewDeployment(ssproxy)
			if err := r.Create(ctx, exp); err != nil {
				ssproxy.SetInitializationFailed()
				_ = r.Status().Update(ctx, ssproxy)
				return ctrl.Result{}, err
			}
		}
	} else {
		act := deploy.DeepCopy()
		exp := reconcile.UpdateDeployment(ssproxy, act)
		if err := r.Update(ctx, exp); err != nil {
			return ctrl.Result{Requeue: true}, err
		}
	}
	return ctrl.Result{}, nil
}

func (r *ProxyReconciler) reconcileHPA(ctx context.Context, namespacedName types.NamespacedName, ssproxy *v1alpha1.ShardingSphereProxy) (ctrl.Result, error) {
	hpa := &autoscalingv2beta2.HorizontalPodAutoscaler{}

	var err error
	if err = r.Get(ctx, namespacedName, hpa); err != nil {
		if !apierrors.IsNotFound(err) {
			return ctrl.Result{}, err
		} else {
			if ssproxy.Spec.AutomaticScaling != nil && ssproxy.Spec.AutomaticScaling.Enable {
				exp := reconcile.NewHPA(ssproxy)
				if err := r.Create(ctx, exp); err != nil {
					ssproxy.SetInitializationFailed()
					_ = r.Status().Update(ctx, ssproxy)
					return ctrl.Result{}, err
				}
			}
		}
	} else {
		if ssproxy.Spec.AutomaticScaling == nil || !ssproxy.Spec.AutomaticScaling.Enable {
			if err := r.Delete(ctx, hpa); err != nil {
				return ctrl.Result{}, err
			}
		} else {
			act := hpa.DeepCopy()
			exp := reconcile.UpdateHPA(ssproxy, act)
			if err := r.Update(ctx, exp); err != nil {
				return ctrl.Result{}, err
			}
		}
	}

	return ctrl.Result{}, nil
}

func (r *ProxyReconciler) reconcileService(ctx context.Context, namespacedName types.NamespacedName, ssproxy *v1alpha1.ShardingSphereProxy) (ctrl.Result, error) {
	service := &v1.Service{}

	var err error
	if err = r.Get(ctx, namespacedName, service); err != nil {
		if !apierrors.IsNotFound(err) {
			return ctrl.Result{}, err
		} else {
			exp := reconcile.NewService(ssproxy)
			if err := r.Create(ctx, exp); err != nil {
				ssproxy.SetInitializationFailed()
				_ = r.Status().Update(ctx, ssproxy)
				return ctrl.Result{}, err
			}
			ssproxy.SetInitialized()
			return ctrl.Result{RequeueAfter: WaitingForReady}, nil
		}
	} else {
		act := service.DeepCopy()
		reconcile.UpdateService(ssproxy, act)
		if err := r.Update(ctx, act); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *ProxyReconciler) reconcilePodList(ctx context.Context, namespace, name string, ssproxy *v1alpha1.ShardingSphereProxy) (ctrl.Result, error) {
	podList := &v1.PodList{}
	if err := r.List(ctx, podList, client.InNamespace(namespace), client.MatchingLabels(map[string]string{"apps": name})); err != nil {
		return ctrl.Result{}, err
	}

	result := ctrl.Result{}
	readyNodes := reconcile.CountingReadyPods(podList)
	if reconcile.IsRunning(podList) {
		if readyNodes < miniReadyCount {
			result.RequeueAfter = WaitingForReady
			if readyNodes != ssproxy.Status.ReadyNodes {
				ssproxy.SetPodStarted(readyNodes)
			}
		} else {
			if ssproxy.Status.Phase != v1alpha1.StatusReady {
				ssproxy.SetReady(readyNodes)
			} else if readyNodes != ssproxy.Spec.Replicas {
				ssproxy.UpdateReadyNodes(readyNodes)
			}
		}
	} else {
		// TODO: Waiting for pods to start exceeds the maximum number of retries
		ssproxy.SetPodNotStarted(readyNodes)
		result.RequeueAfter = WaitingForReady
	}

	// TODO: Compare Status with or without modification
	if err := r.Status().Update(ctx, ssproxy); err != nil {
		return result, err
	}

	return result, nil
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
