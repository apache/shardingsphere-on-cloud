/*
 * Copyright (c) 2022.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
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
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logger "sigs.k8s.io/controller-runtime/pkg/log"
	shardingspherev1alpha1 "sphere-ex.com/shardingsphere-operator/api/v1alpha1"
	"sphere-ex.com/shardingsphere-operator/pkg/reconcile"
	"time"
)

const (
	//WaitingForReady 时间选择参考 kubelet 重启时间
	WaitingForReady = 10 * time.Second
	Threshold       = int32(5)
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
	// TODO: 超过错误重试阈值
	// 错误重试应该以次数为限制，超过阈值限制将直接转化为失败最终态，不再继续重新排队
	// 调谐代码中错误尝试阈值实现方式暂时不完善

	// TODO: 错误需不需要进行重新排队？
	log := logger.FromContext(ctx)

	run := &shardingspherev1alpha1.Proxy{}
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
		// TODO: 对比 Deployment 有没有更改
		reconcile.UpdateDeployment(run, runtimeDeployment)
		err = r.Update(ctx, runtimeDeployment)
		if err != nil {
			log.Error(err, "Error updating cascaded deployment")
			// 重试是为了处理出现冲突这个错误
			// TODO: 单独为冲突错误进行重新排队处理
			return ctrl.Result{Requeue: true}, err
		}
	}

	// TODO: service 是否需要纠正
	runtimeService := &v1.Service{}
	err = r.Get(ctx, req.NamespacedName, runtimeService)
	cascadingService := reconcile.ConstructCascadingService(run)
	if apierrors.IsNotFound(err) {
		err = r.Create(ctx, cascadingService)
		if err != nil {
			run.SetInitializationFailed()
			_ = r.Status().Update(ctx, run)
			log.Error(err, "Error creating cascaded service")
			return ctrl.Result{}, err
		}
	} else if err != nil {
		log.Error(err, "Error getting cascaded service")
		return ctrl.Result{}, err
	}

	result := ctrl.Result{}
	podList := &v1.PodList{}
	err = r.List(ctx, podList, client.InNamespace(req.Namespace), client.MatchingLabels(map[string]string{"apps": req.Name}))
	if err != nil {
		log.Error(err, "Error listing cascaded pod")
		return ctrl.Result{}, err
	}
	if reconcile.IsRunning(podList) {
		readyNodes := reconcile.ReadyCount(podList)
		if readyNodes != run.Spec.Replicas {
			restartTimes := reconcile.RestartCount(podList)
			if restartTimes > Threshold {
				run.SetFailed(readyNodes)
				_ = r.Status().Update(ctx, run)
				log.Error(nil, "The times of restarts exceeds the threshold")
				return result, nil
			}
			result.RequeueAfter = (time.Duration(restartTimes) + 1) * WaitingForReady
			if readyNodes != run.Status.ReadyNodes {
				run.SetPodStarted(readyNodes)
			}
		} else {
			if run.Status.Phase != shardingspherev1alpha1.StatusReady {
				log.Info("Status is now ready!")
				run.SetReady(readyNodes)
			}
		}
	} else {
		run.SetInitialized()
		result.RequeueAfter = WaitingForReady
	}
	// TODO: 对比 Status 有没有修改
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
		For(&shardingspherev1alpha1.Proxy{}).
		Owns(&appsv1.Deployment{}).
		Owns(&v1.Pod{}).
		Complete(r)
}
