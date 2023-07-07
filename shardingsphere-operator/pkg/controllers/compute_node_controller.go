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
	"time"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/deployment"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/service"
	reconcile "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/computenode"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	computeNodeControllerName = "compute-node-controller"
	defaultRequeueTime        = 10 * time.Second
)

// ComputeNodeReconciler is a controller for the compute node
type ComputeNodeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger

	Deployment        deployment.Deployment
	DeploymentBuilder reconcile.Builder
	Service           service.Service
	ConfigMap         configmap.ConfigMap
}

// SetupWithManager sets up the controller with the Manager
func (r *ComputeNodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ComputeNode{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Pod{}).
		Owns(&corev1.Service{}).
		Owns(&corev1.ConfigMap{}).
		Complete(r)
}

// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=computenodes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=computenodes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// Reconcile handles main function of this controller
func (r *ComputeNodeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues(computeNodeControllerName, req.NamespacedName)

	cn := &v1alpha1.ComputeNode{}
	if err := r.Get(ctx, req.NamespacedName, cn); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{RequeueAfter: defaultRequeueTime}, nil
		}

		logger.Error(err, "Failed to get the compute node")
		return ctrl.Result{Requeue: true}, err
	}

	if err := r.reconcileStatus(ctx, cn); err != nil {
		logger.Error(err, "Failed to reconcile status")
	}

	errors := []error{}
	if err := r.reconcileDeployment(ctx, cn); err != nil {
		logger.Error(err, "Failed to reconcile deployement")
		errors = append(errors, err)
	}
	if err := r.reconcileService(ctx, cn); err != nil {
		logger.Error(err, "Failed to reconcile service")
		errors = append(errors, err)
	}
	if err := r.reconcileConfigMap(ctx, cn); err != nil {
		logger.Error(err, "Failed to reconcile configmap")
		errors = append(errors, err)
	}

	if len(errors) != 0 {
		return ctrl.Result{Requeue: true}, errors[0]
	}

	return ctrl.Result{RequeueAfter: defaultRequeueTime}, nil
}

func (r *ComputeNodeReconciler) reconcileDeployment(ctx context.Context, cn *v1alpha1.ComputeNode) error {
	deploy, err := r.getDeploymentByNamespacedName(ctx, types.NamespacedName{Namespace: cn.Namespace, Name: cn.Name})
	if err != nil {
		return err
	}
	if deploy != nil {
		return r.updateDeployment(ctx, cn, deploy)
	}
	return r.createDeployment(ctx, cn)
}

func (r *ComputeNodeReconciler) createDeployment(ctx context.Context, cn *v1alpha1.ComputeNode) error {
	deploy := r.DeploymentBuilder.Build(ctx, cn)
	err := r.Deployment.Create(ctx, deploy)
	if err != nil && apierrors.IsAlreadyExists(err) || err == nil {
		return nil
	}
	return err
}

func (r *ComputeNodeReconciler) updateDeployment(ctx context.Context, cn *v1alpha1.ComputeNode, deploy *appsv1.Deployment) error {
	exp := r.DeploymentBuilder.Build(ctx, cn)
	exp.ObjectMeta = deploy.ObjectMeta
	exp.Labels = deploy.Labels
	exp.Annotations = deploy.Annotations

	if !reflect.DeepEqual(deploy.Spec, exp.Spec) {
		return r.Deployment.Update(ctx, exp)
	}
	return nil
}

func (r *ComputeNodeReconciler) getDeploymentByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*appsv1.Deployment, error) {
	dp, err := r.Deployment.GetByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, err
	}
	return dp, nil
}

func (r *ComputeNodeReconciler) reconcileService(ctx context.Context, cn *v1alpha1.ComputeNode) error {
	svc, err := r.getServiceByNamespacedName(ctx, types.NamespacedName{Namespace: cn.Namespace, Name: cn.Name})
	if err != nil {
		return err
	}
	if svc != nil {
		return r.updateService(ctx, cn, svc)
	}
	return r.createService(ctx, cn)
}

func (r *ComputeNodeReconciler) createService(ctx context.Context, cn *v1alpha1.ComputeNode) error {
	svc := r.Service.Build(ctx, cn)
	err := r.Service.Create(ctx, svc)
	if err != nil && apierrors.IsAlreadyExists(err) || err == nil {
		return nil
	}
	return err
}

func (r *ComputeNodeReconciler) updateComputeNodePortBindings(ctx context.Context, cn *v1alpha1.ComputeNode) error {
	if rt, err := r.getRuntimeComputeNode(ctx, types.NamespacedName{
		Namespace: cn.Namespace,
		Name:      cn.Name,
	}); err == nil {
		rt.Spec.PortBindings = cn.Spec.PortBindings
		if err := r.Update(ctx, rt); err != nil {
			return err
		}
	}
	return nil
}

func (r *ComputeNodeReconciler) updateService(ctx context.Context, cn *v1alpha1.ComputeNode, s *corev1.Service) error {
	pbs := []v1alpha1.PortBinding{}
	copy(cn.Spec.PortBindings, pbs)
	switch cn.Spec.ServiceType {
	case corev1.ServiceTypeClusterIP:
		updateServiceClusterIP(cn.Spec.PortBindings)
		if !reflect.DeepEqual(cn.Spec.PortBindings, pbs) {
			return r.updateComputeNodePortBindings(ctx, cn)
		}
	case corev1.ServiceTypeExternalName:
		fallthrough
	case corev1.ServiceTypeLoadBalancer:
		fallthrough
	case corev1.ServiceTypeNodePort:
		updateServiceNodePort(cn.Spec.PortBindings, s.Spec.Ports)
		if !reflect.DeepEqual(cn.Spec.PortBindings, pbs) {
			return r.updateComputeNodePortBindings(ctx, cn)
		}
	}

	exp := r.Service.Build(ctx, cn)
	exp.ObjectMeta = s.ObjectMeta
	exp.Spec.ClusterIP = s.Spec.ClusterIP
	exp.Spec.ClusterIPs = s.Spec.ClusterIPs

	if cn.Spec.ServiceType == corev1.ServiceTypeNodePort {
		exp.Spec.Ports = updateNodePorts(cn.Spec.PortBindings, s.Spec.Ports)
	}

	if !reflect.DeepEqual(exp.Spec, s.Spec) {
		return r.Update(ctx, exp)
	}
	return nil
}

func updateServiceNodePort(pbs []v1alpha1.PortBinding, ports []corev1.ServicePort) {
	for idx := range ports {
		for i := range pbs {
			if ports[idx].Name == pbs[i].Name {
				if pbs[i].NodePort == 0 {
					pbs[i].NodePort = ports[idx].NodePort
				}
				break
			}
		}
	}
}

func updateServiceClusterIP(portBindings []v1alpha1.PortBinding) {
	for idx := range portBindings {
		if portBindings[idx].NodePort != 0 {
			portBindings[idx].NodePort = 0
			break
		}
	}
}

func updateNodePorts(portbindings []v1alpha1.PortBinding, svcports []corev1.ServicePort) []corev1.ServicePort {
	ports := []corev1.ServicePort{}
	for pb := range portbindings {
		for sp := range svcports {
			if portbindings[pb].Name == svcports[sp].Name {
				port := corev1.ServicePort{
					Name:       portbindings[pb].Name,
					TargetPort: intstr.FromInt(int(portbindings[pb].ContainerPort)),
					Port:       portbindings[pb].ServicePort,
					Protocol:   portbindings[pb].Protocol,
				}
				if svcports[sp].NodePort != 0 {
					port.NodePort = svcports[sp].NodePort
				}
				ports = append(ports, port)
			}
		}
	}
	return ports
}

func (r *ComputeNodeReconciler) getServiceByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*corev1.Service, error) {
	svc, err := r.Service.GetByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func (r *ComputeNodeReconciler) createConfigMap(ctx context.Context, cn *v1alpha1.ComputeNode) error {
	cm := r.ConfigMap.Build(ctx, cn)
	err := r.ConfigMap.Create(ctx, cm)
	if err != nil && apierrors.IsAlreadyExists(err) || err == nil {
		return nil
	}
	return err
}

func (r *ComputeNodeReconciler) updateConfigMap(ctx context.Context, cn *v1alpha1.ComputeNode, cm *corev1.ConfigMap) error {
	exp := r.ConfigMap.Build(ctx, cn)
	exp.ObjectMeta = cm.ObjectMeta
	exp.Labels = cm.Labels
	exp.Annotations = cm.Annotations
	if !reflect.DeepEqual(cm.Data, exp.Data) {
		return r.ConfigMap.Update(ctx, exp)
	}
	return nil
}

func (r *ComputeNodeReconciler) getConfigMapByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*corev1.ConfigMap, error) {
	cm, err := r.ConfigMap.GetByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, err
	}
	return cm, nil
}

func (r *ComputeNodeReconciler) reconcileConfigMap(ctx context.Context, cn *v1alpha1.ComputeNode) error {
	cm, err := r.getConfigMapByNamespacedName(ctx, types.NamespacedName{Namespace: cn.Namespace, Name: cn.Name})
	if err != nil {
		return err
	}
	if cm != nil {
		return r.updateConfigMap(ctx, cn, cm)
	}
	return r.createConfigMap(ctx, cn)
}

func (r *ComputeNodeReconciler) reconcileStatus(ctx context.Context, cn *v1alpha1.ComputeNode) error {
	podlist := &corev1.PodList{}
	if err := r.List(ctx, podlist, client.InNamespace(cn.Namespace), client.MatchingLabels(cn.Spec.Selector.MatchLabels)); err != nil {
		return err
	}

	service := &corev1.Service{}
	if err := r.Get(ctx, types.NamespacedName{
		Namespace: cn.Namespace,
		Name:      cn.Name,
	}, service); err != nil {
		if !apierrors.IsNotFound(err) {
			return err
		}
	}

	status := reconcileComputeNodeStatus(podlist, service, cn)
	rt, err := r.getRuntimeComputeNode(ctx, types.NamespacedName{
		Namespace: cn.Namespace,
		Name:      cn.Name,
	})
	if err != nil {
		return err
	}
	rt.Status = *status

	return r.Status().Update(ctx, rt)
}

func getReadyProxyInstances(podlist *corev1.PodList) int32 {
	var cnt int32

	findRunningPod := func(pod *corev1.Pod) {
		if pod.Status.Phase != corev1.PodRunning {
			return
		}

		if isTrueReadyPod(pod) {
			for j := range pod.Status.ContainerStatuses {
				if pod.Status.ContainerStatuses[j].Name == "shardingsphere-proxy" && pod.Status.ContainerStatuses[j].Ready {
					cnt++
				}
			}
		}
	}

	for idx := range podlist.Items {
		findRunningPod(&podlist.Items[idx])
	}
	return cnt
}

func isTrueReadyPod(pod *corev1.Pod) bool {
	for i := range pod.Status.Conditions {
		if pod.Status.Conditions[i].Type == corev1.PodReady && pod.Status.Conditions[i].Status == corev1.ConditionTrue {
			return true
		}
	}
	return false
}

func updateComputeNodeStatusCondition(conditions []v1alpha1.ComputeNodeCondition, conds []v1alpha1.ComputeNodeCondition) []v1alpha1.ComputeNodeCondition {
	for idx := range conds {
		var found bool
		for i := range conditions {
			conditions[i].LastUpdateTime = conds[idx].LastUpdateTime
			if conditions[i].Type == conds[idx].Type {
				found = true
				conditions[i].Type = conds[idx].Type
				conditions[i].Status = conds[idx].Status
				conditions[i].Message = conds[idx].Message
				conditions[i].Reason = conds[idx].Reason
			} else if conds[idx].Type == v1alpha1.ComputeNodeConditionUnknown || conditions[i].Type == v1alpha1.ComputeNodeConditionUnknown {
				conditions[i].Status = v1alpha1.ConditionStatusFalse
			} else {
				continue
			}
		}

		// check current conditions
		if len(conditions) == 0 || !found {
			conditions = append(conditions, conds[idx])
		}
	}

	return conditions
}

func reconcileComputeNodeStatus(podlist *corev1.PodList, svc *corev1.Service, cn *v1alpha1.ComputeNode) *v1alpha1.ComputeNodeStatus {
	conds := reconcile.GetConditionFromPods(podlist)

	cn.Status.Conditions = updateComputeNodeStatusCondition(cn.Status.Conditions, conds)

	ready := getReadyProxyInstances(podlist)
	cn.Status.Ready = fmt.Sprintf("%d/%d", ready, len(podlist.Items))
	cn.Status.Replicas = int32(len(podlist.Items))

	if ready > 0 {
		cn.Status.Phase = v1alpha1.ComputeNodeStatusReady
	} else {
		cn.Status.Phase = v1alpha1.ComputeNodeStatusNotReady
	}

	cn.Status.LoadBalancer.ClusterIP = svc.Spec.ClusterIP
	cn.Status.LoadBalancer.Ingress = svc.Status.LoadBalancer.Ingress
	return &cn.Status
}

func (r *ComputeNodeReconciler) getRuntimeComputeNode(ctx context.Context, namespacedName types.NamespacedName) (*v1alpha1.ComputeNode, error) {
	rt := &v1alpha1.ComputeNode{}
	err := r.Get(ctx, namespacedName, rt)
	return rt, err
}
