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
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/deployment"
	reconcile "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/computenode"
	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	computeNodeControllerName = "compute-node-controller"
	defaultRequeueTime        = 10 * time.Second
)

type ComputeNodeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger

	Deployment deployment.Deployment
	// Service
	// ConfigMap
}

// SetupWithManager sets up the controller with the Manager.
func (r *ComputeNodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ComputeNode{}).
		Owns(&appsv1.Deployment{}).
		Owns(&v1.Service{}).
		Owns(&v1.ConfigMap{}).
		Complete(r)
}

func (r *ComputeNodeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues(computeNodeControllerName, req.NamespacedName)

	cn := &v1alpha1.ComputeNode{}
	if err := r.Get(ctx, req.NamespacedName, cn); err != nil {
		if apierrors.IsNotFound(err) {
			logger.Error(err, "computenode not found")
			return ctrl.Result{RequeueAfter: defaultRequeueTime}, nil
		} else {
			logger.Error(err, "get computenode")
			return ctrl.Result{Requeue: true}, err
		}
	}

	errors := []error{}
	if err := r.reconcileDeployment(ctx, cn); err != nil {
		logger.Error(err, "reconcile deployment")
		errors = append(errors, err)
	}
	if err := r.reconcileService(ctx, cn); err != nil {
		logger.Error(err, "reconcile service")
		errors = append(errors, err)
	}
	if err := r.reconcileConfigMap(ctx, cn); err != nil {
		logger.Error(err, "reconcile configmap")
		errors = append(errors, err)
	}
	if err := r.reconcileStatus(ctx, cn); err != nil {
		logger.Error(err, "reconcile pod list")
		errors = append(errors, err)
	}

	if len(errors) != 0 {
		return ctrl.Result{Requeue: true}, errors[0]
	}

	return ctrl.Result{RequeueAfter: defaultRequeueTime}, nil
}

func (r *ComputeNodeReconciler) reconcileDeployment(ctx context.Context, cn *v1alpha1.ComputeNode) error {
	deploy, found, err := r.getDeploymentByNamespacedName(ctx, types.NamespacedName{Namespace: cn.Namespace, Name: cn.Name})
	if found {
		if err := r.updateDeployment(ctx, cn, deploy); err != nil {
			return err
		}
	} else {
		if err != nil {
			return err
		} else {
			if err := r.createDeployment(ctx, cn); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *ComputeNodeReconciler) createDeployment(ctx context.Context, cn *v1alpha1.ComputeNode) error {
	deploy := reconcile.NewDeployment(cn)
	if err := r.Create(ctx, deploy); err != nil {
		return err
	}
	return nil
}

func (r *ComputeNodeReconciler) updateDeployment(ctx context.Context, cn *v1alpha1.ComputeNode, deploy *appsv1.Deployment) error {
	exp := reconcile.UpdateDeployment(cn, deploy)
	if err := r.Update(ctx, exp); err != nil {
		return err
	}
	return nil
}

func (r *ComputeNodeReconciler) getDeploymentByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*appsv1.Deployment, bool, error) {
	dp, err := r.Deployment.GetByNamespacedName(ctx, namespacedName)
	// found
	if dp != nil {
		return dp, true, nil
	}
	// error
	if err != nil {
		return nil, false, err
	} else {
		// not found
		return nil, false, nil
	}
}

func (r *ComputeNodeReconciler) reconcileService(ctx context.Context, cn *v1alpha1.ComputeNode) error {
	cur := &v1.Service{}
	if err := r.Get(ctx, types.NamespacedName{
		Namespace: cn.Namespace,
		Name:      cn.Name,
	}, cur); err != nil {
		if apierrors.IsNotFound(err) {
			// create
			exp := reconcile.ComputeNodeNewService(cn)
			if err := r.Create(ctx, exp); err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	}
	// update
	exp := reconcile.ComputeNodeUpdateService(cn, cur)
	if err := r.Update(ctx, exp); err != nil {
		return err
	}

	return nil
}

func (r *ComputeNodeReconciler) reconcileConfigMap(ctx context.Context, cn *v1alpha1.ComputeNode) error {
	cur := &v1.ConfigMap{}
	if err := r.Get(ctx, types.NamespacedName{
		Namespace: cn.Namespace,
		Name:      cn.Name,
	}, cur); err != nil {
		if apierrors.IsNotFound(err) {
			// create
			exp := reconcile.ComputeNodeNewConfigMap(cn)
			if err := r.Create(ctx, exp); err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	}

	// update
	//FIXME: need to rolling update Deployment if ConfigMap indeed updated
	exp := reconcile.ComputeNodeUpdateConfigMap(cn, cur)
	if err := r.Update(ctx, exp); err != nil {
		return err
	}

	return nil
}

func (r *ComputeNodeReconciler) reconcileStatus(ctx context.Context, cn *v1alpha1.ComputeNode) error {
	podList := &v1.PodList{}
	if err := r.List(ctx, podList, client.InNamespace(cn.Namespace), client.MatchingLabels(cn.Labels)); err != nil {
		return err
	}

	service := &v1.Service{}
	if err := r.Get(ctx, types.NamespacedName{
		Namespace: cn.Namespace,
		Name:      cn.Name,
	}, service); err != nil {
		return err
	}

	rt, err := r.getRuntimeComputeNode(ctx, types.NamespacedName{
		Namespace: cn.Namespace,
		Name:      cn.Name,
	})
	if err != nil {
		return err
	}

	rt.Status = reconcileComputeNodeStatus(*podList, *service, *rt)

	// TODO: Compare Status with or without modification
	if err := r.Status().Update(ctx, rt); err != nil {
		return err
	}

	return nil
}

func getReadyInstances(podlist v1.PodList) int32 {
	var cnt int32
	for _, p := range podlist.Items {
		if p.Status.Phase == v1.PodRunning {
			for _, c := range p.Status.Conditions {
				if c.Type == v1.PodReady && c.Status == v1.ConditionTrue {
					for _, con := range p.Status.ContainerStatuses {
						if con.Name == "shardingsphere-proxy" && con.Ready {
							cnt++
						}
					}
				}
			}
		}
	}
	return cnt
}

func newConditions(conditions []v1alpha1.ComputeNodeCondition, cond v1alpha1.ComputeNodeCondition) []v1alpha1.ComputeNodeCondition {
	if conditions == nil {
		conditions = []v1alpha1.ComputeNodeCondition{}
	}
	if cond.Type == "" {
		return conditions
	}

	found := false
	for idx, _ := range conditions {
		if conditions[idx].Type == cond.Type {
			conditions[idx].LastUpdateTime = cond.LastUpdateTime
			conditions[idx].Status = cond.Status
			found = true
			break
		}
	}

	if !found {
		conditions = append(conditions, cond)
	}

	return conditions
}

func updateReadyConditions(conditions []v1alpha1.ComputeNodeCondition, cond v1alpha1.ComputeNodeCondition) []v1alpha1.ComputeNodeCondition {
	return newConditions(conditions, cond)
}

func updateNotReadyConditions(conditions []v1alpha1.ComputeNodeCondition, cond v1alpha1.ComputeNodeCondition) []v1alpha1.ComputeNodeCondition {
	cur := newConditions(conditions, cond)

	for idx, _ := range cur {
		if cur[idx].Type == v1alpha1.ComputeNodeConditionReady {
			cur[idx].LastUpdateTime = metav1.Now()
			cur[idx].Status = v1alpha1.ConditionStatusFalse
		}
	}

	return cur
}

func clusterCondition(podlist v1.PodList) v1alpha1.ComputeNodeCondition {
	cond := v1alpha1.ComputeNodeCondition{}
	if len(podlist.Items) == 0 {
		return cond
	}

	condStarted := v1alpha1.ComputeNodeCondition{
		Type:           v1alpha1.ComputeNodeConditionStarted,
		Status:         v1alpha1.ConditionStatusTrue,
		LastUpdateTime: metav1.Now(),
	}
	condUnknown := v1alpha1.ComputeNodeCondition{
		Type:           v1alpha1.ComputeNodeConditionUnknown,
		Status:         v1alpha1.ConditionStatusTrue,
		LastUpdateTime: metav1.Now(),
	}
	condDeployed := v1alpha1.ComputeNodeCondition{
		Type:           v1alpha1.ComputeNodeConditionDeployed,
		Status:         v1alpha1.ConditionStatusTrue,
		LastUpdateTime: metav1.Now(),
	}
	condFailed := v1alpha1.ComputeNodeCondition{
		Type:           v1alpha1.ComputeNodeConditionFailed,
		Status:         v1alpha1.ConditionStatusTrue,
		LastUpdateTime: metav1.Now(),
	}

	//FIXME: do not capture ConditionStarted in some cases
	for _, p := range podlist.Items {
		switch p.Status.Phase {
		case v1.PodRunning:
			return condStarted
		case v1.PodUnknown:
			return condUnknown
		case v1.PodPending:
			return condDeployed
		case v1.PodFailed:
			return condFailed
		}
	}
	return cond
}

func reconcileComputeNodeStatus(podlist v1.PodList, svc v1.Service, rt v1alpha1.ComputeNode) v1alpha1.ComputeNodeStatus {
	readyInstances := getReadyInstances(podlist)

	rt.Status.ReadyInstances = readyInstances
	if rt.Spec.Replicas == 0 {
		rt.Status.Phase = v1alpha1.ComputeNodeStatusNotReady
	} else {
		if readyInstances < miniReadyCount {
			rt.Status.Phase = v1alpha1.ComputeNodeStatusNotReady
		} else {
			rt.Status.Phase = v1alpha1.ComputeNodeStatusReady
		}
	}

	if rt.Status.Phase == v1alpha1.ComputeNodeStatusReady {
		rt.Status.Conditions = updateReadyConditions(rt.Status.Conditions, v1alpha1.ComputeNodeCondition{
			Type:           v1alpha1.ComputeNodeConditionReady,
			Status:         v1alpha1.ConditionStatusTrue,
			LastUpdateTime: metav1.Now(),
		})
	} else {
		cond := clusterCondition(podlist)
		rt.Status.Conditions = updateNotReadyConditions(rt.Status.Conditions, cond)
	}

	rt.Status.LoadBalancer.ClusterIP = svc.Spec.ClusterIP
	rt.Status.LoadBalancer.Ingress = svc.Status.LoadBalancer.Ingress

	return rt.Status
}

func (r *ComputeNodeReconciler) getRuntimeComputeNode(ctx context.Context, namespacedName types.NamespacedName) (*v1alpha1.ComputeNode, error) {
	rt := &v1alpha1.ComputeNode{}
	err := r.Get(ctx, namespacedName, rt)
	return rt, err
}
