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
	"reflect"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes"
	reconcile "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/autoscaler"

	"github.com/go-logr/logr"
	autoscalingv2beta2 "k8s.io/api/autoscaling/v2"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	autoscalingv1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	autoScalerControllerName = "autoscaler-controller"
)

// AutoScalerReconciler is a controller for the shardingsphere cluster
type AutoScalerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger

	Builder   reconcile.Builder
	Resources kubernetes.Resources
}

// SetupWithManager sets up the controller with the Manager
func (r *AutoScalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.AutoScaler{}).
		Owns(&autoscalingv2beta2.HorizontalPodAutoscaler{}).
		Complete(r)
}

// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=computenodes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=autoscaling/v2,resources=horizontalpodautoscaler,verbs=get;list;watch;create;update;patch;delete
// Reconcile handles main function of this controller
func (r *AutoScalerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues(autoScalerControllerName, req.NamespacedName)

	as := &v1alpha1.AutoScaler{}
	if err := r.Get(ctx, req.NamespacedName, as); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{RequeueAfter: defaultRequeueTime}, nil
		}

		logger.Error(err, "Failed to get the autoscaler")
		return ctrl.Result{Requeue: true}, err
	}

	if err := r.reconcileAutoScaler(ctx, as); err != nil {
		logger.Error(err, "Failed to reconcile autoscaler")
	}

	return ctrl.Result{RequeueAfter: defaultRequeueTime}, nil
}

func (r *AutoScalerReconciler) reconcileAutoScaler(ctx context.Context, as *v1alpha1.AutoScaler) error {
	gvk := as.GroupVersionKind()

	for i := range as.Spec.PolicyGroup {
		pg := as.Spec.PolicyGroup[i]
		if pg.Provider == "KubernetesHPA" && pg.Horizontal != nil {
			if err := r.reconcileHPA(ctx, &as.ObjectMeta, gvk, &pg); err != nil {
				return err
			}
		}
		if pg.Provider == "KubernetesVPA" && pg.Vertical != nil {
			if err := r.reconcileVPA(ctx, &as.ObjectMeta, gvk, &pg); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *AutoScalerReconciler) reconcileHPA(ctx context.Context, meta *metav1.ObjectMeta, gvk schema.GroupVersionKind, policy *v1alpha1.ScalingPolicy) error {
	hpa, err := r.getHPAByNamespacedName(ctx, types.NamespacedName{Namespace: meta.Namespace, Name: meta.Name})
	if err != nil {
		return err
	}
	if hpa != nil {
		return r.updateHPA(ctx, meta, gvk, policy, hpa)
	}
	return r.createHPA(ctx, meta, gvk, policy)
}

func (r *AutoScalerReconciler) getHPAByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*autoscalingv2beta2.HorizontalPodAutoscaler, error) {
	return r.Resources.HPA().GetByNamespacedName(ctx, namespacedName)
}

// nolint:dupl
func (r *AutoScalerReconciler) updateHPA(ctx context.Context, meta *metav1.ObjectMeta, gvk schema.GroupVersionKind, policy *v1alpha1.ScalingPolicy, hpa *autoscalingv2beta2.HorizontalPodAutoscaler) error {
	exp := r.Builder.BuildHorizontalPodAutoScaler(ctx, meta, gvk, policy)
	exp.ObjectMeta = hpa.ObjectMeta
	exp.Labels = hpa.Labels
	exp.Annotations = hpa.Annotations

	if !reflect.DeepEqual(hpa.Spec, exp.Spec) {
		return r.Resources.HPA().Update(ctx, hpa)
	}
	return nil
}

func (r *AutoScalerReconciler) createHPA(ctx context.Context, meta *metav1.ObjectMeta, gvk schema.GroupVersionKind, policy *v1alpha1.ScalingPolicy) error {
	hpa := r.Builder.BuildHorizontalPodAutoScaler(ctx, meta, gvk, policy)
	err := r.Resources.HPA().Create(ctx, hpa)
	if err != nil && apierrors.IsAlreadyExists(err) || err == nil {
		return nil
	}
	return err
}

func (r *AutoScalerReconciler) reconcileVPA(ctx context.Context, meta *metav1.ObjectMeta, gvk schema.GroupVersionKind, policy *v1alpha1.ScalingPolicy) error {
	vpa, err := r.getVPAByNamespacedName(ctx, types.NamespacedName{Namespace: meta.Namespace, Name: meta.Name})
	if err != nil {
		return err
	}
	if vpa != nil {
		return r.updateVPA(ctx, meta, gvk, policy, vpa)
	}
	return r.createVPA(ctx, meta, gvk, policy)
}

func (r *AutoScalerReconciler) getVPAByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*autoscalingv1.VerticalPodAutoscaler, error) {
	return r.Resources.VPA().GetByNamespacedName(ctx, namespacedName)
}

// nolint:dupl
func (r *AutoScalerReconciler) updateVPA(ctx context.Context, meta *metav1.ObjectMeta, gvk schema.GroupVersionKind, policy *v1alpha1.ScalingPolicy, vpa *autoscalingv1.VerticalPodAutoscaler) error {
	exp := r.Builder.BuildVerticalPodAutoscaler(ctx, meta, gvk, policy)
	exp.ObjectMeta = vpa.ObjectMeta
	exp.Labels = vpa.Labels
	exp.Annotations = vpa.Annotations

	if !reflect.DeepEqual(vpa.Spec, exp.Spec) {
		return r.Resources.VPA().Update(ctx, vpa)
	}
	return nil
}

// nolint:dupl
func (r *AutoScalerReconciler) createVPA(ctx context.Context, meta *metav1.ObjectMeta, gvk schema.GroupVersionKind, policy *v1alpha1.ScalingPolicy) error {
	vpa := r.Builder.BuildVerticalPodAutoscaler(ctx, meta, gvk, policy)
	err := r.Resources.VPA().Create(ctx, vpa)
	if err != nil && apierrors.IsAlreadyExists(err) || err == nil {
		return nil
	}
	return err
}
