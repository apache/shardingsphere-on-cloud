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
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaosmesh"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/job"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/pressure"
	sschaos "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/shardingspherechaos"

	"github.com/go-logr/logr"
	batchV1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const (
	ShardingSphereChaosControllerName = "shardingsphere-chaos-controller"
	SSChaosFinalizeName               = "shardingsphere.apache.org/finalizer"
)

// ShardingSphereChaosReconciler is a controller for the ShardingSphereChaos
type ShardingSphereChaosReconciler struct {
	client.Client

	Scheme    *runtime.Scheme
	Log       logr.Logger
	Events    record.EventRecorder
	ClientSet *clientset.Clientset

	Chaos     chaosmesh.Chaos
	Job       job.Job
	ExecCtrls []*ExecCtrl
	ConfigMap configmap.ConfigMap
}

type ExecCtrl struct {
	cancel   context.CancelFunc
	pressure *pressure.Pressure
	ctx      context.Context
}

// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=shardingspherechaos,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=shardingspherechaos/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=shardingsphere.apache.org,resources=shardingspherechaos/finalizers,verbs=update

// Reconcile handles main function of this controller
func (r *ShardingSphereChaosReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues(ShardingSphereChaosControllerName, req.NamespacedName)

	ssChaos, err := r.getRuntimeChaos(ctx, req.NamespacedName)

	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("start reconcile shardingspherechaos")
	if ssChaos.ObjectMeta.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(ssChaos, SSChaosFinalizeName) {
			controllerutil.AddFinalizer(ssChaos, SSChaosFinalizeName)
			if err := r.Update(ctx, ssChaos); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else if controllerutil.ContainsFinalizer(ssChaos, SSChaosFinalizeName) {
		return r.finalize(ctx, ssChaos)
	}

	var errors []error
	if err := r.reconcileChaos(ctx, ssChaos); err != nil {
		errors = append(errors, err)

		logger.Error(err, "reconcile shardingspherechaos error")
		r.Events.Event(ssChaos, "Warning", "shardingspherechaos error", err.Error())
	}

	if err := r.reconcileConfigMap(ctx, ssChaos); err != nil {
		errors = append(errors, err)

		logger.Error(err, "reconcile configmap error")
		r.Events.Event(ssChaos, "Warning", "configmap error", err.Error())
	}

	if err := r.reconcilePressure(ctx, ssChaos); err != nil {
		errors = append(errors, err)
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

func (r *ShardingSphereChaosReconciler) reconcilePressure(ctx context.Context, chao *v1alpha1.ShardingSphereChaos) error {

	if chao.Status.Phase == "" {
		return nil
	}
	exec := r.getNeedExec(chao)

	//if exec in this phase do not exist,create it
	if exec == nil {
		exec := pressure.NewPressure(getExecName(chao), chao.Spec.PressureCfg.DistSQLs)

		//we need to set active to true,prevent it start after we start reconcile status
		exec.Active = true
		execCtx, cancel := context.WithCancel(ctx)
		execCtrl := &ExecCtrl{
			cancel:   cancel,
			pressure: exec,
			ctx:      execCtx,
		}

		go exec.Run(execCtx, &chao.Spec.PressureCfg)
		r.ExecCtrls = append(r.ExecCtrls, execCtrl)
	}

	return nil
}

func (r *ShardingSphereChaosReconciler) reconcileStatus(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos) error {

	cur := chaos.Status.DeepCopy()

	setDefaultStatus(chaos)

	updateCondition(chaos)

	r.updatePhaseExec(chaos)

	if err := r.updateChaosCondition(ctx, chaos); err != nil {
		return err
	}

	if reflect.DeepEqual(cur, chaos.Status) {

		return nil
	}

	return r.Status().Update(ctx, chaos)
}

func updateCondition(chaos *v1alpha1.ShardingSphereChaos) {
	phase := chaos.Status.Phase

	for i := range chaos.Status.Conditions {
		condition := chaos.Status.Conditions[i]
		if string(phase) == condition.Type {
			if condition.Status == v1alpha1.ConditionStatusFalse {
				condition.Status = v1alpha1.ConditionStatusTrue
				condition.LastTransitionTime = metav1.Time{Time: time.Now()}
			}
			return
		}
	}
}

func (r *ShardingSphereChaosReconciler) updatePhaseExec(chaos *v1alpha1.ShardingSphereChaos) {
	exec := r.getNeedExec(chaos)

	//because the goroutine asynchronous,we cant check it start immediately or not
	if exec == nil || exec.Active {
		return
	}

	//todo: judge error

	msg := generateMsgFromExec(exec)
	var nextPhase v1alpha1.ChaosPhase
	//when exec finished, update phase
	switch chaos.Status.Phase {
	case v1alpha1.BeforeSteady:
		nextPhase = v1alpha1.AfterSteady
	case v1alpha1.AfterSteady:
		chaos.Status.Result.Steady = *msg
		//todo: add metrics

		nextPhase = v1alpha1.BeforeChaos
	case v1alpha1.BeforeChaos:
		chaos.Status.Result.Chaos = *msg
		//todo: add metrics

		nextPhase = v1alpha1.AfterChaos
	//case v1alpha1.AfterChaos:
	//	//todo: check result here
	//	return
	default:
		return
	}

	chaos.Status.Phase = nextPhase
}

func generateMsgFromExec(exec *pressure.Pressure) *v1alpha1.Msg {
	//todo: wait to change result compute way

	msg := v1alpha1.Msg{
		Result:   fmt.Sprintf("%d/%d", exec.Result.Success, exec.Result.Total),
		Duration: exec.Result.Duration.String(),
	}
	if exec.Err != nil {
		msg.FailureDetails = exec.Err.Error()
	}

	return &msg
}

func getExecName(chao *v1alpha1.ShardingSphereChaos) string {
	var execName string
	nameSpacedName := types.NamespacedName{Namespace: chao.Namespace, Name: chao.Name}

	if chao.Status.Phase == v1alpha1.BeforeSteady || chao.Status.Phase == v1alpha1.AfterSteady {
		execName = makeExecName(nameSpacedName, string(sschaos.InSteady))
	}
	if chao.Status.Phase == v1alpha1.BeforeChaos || chao.Status.Phase == v1alpha1.AfterChaos {
		execName = makeExecName(nameSpacedName, string(sschaos.InChaos))
	}

	return execName
}

func makeExecName(namespacedName types.NamespacedName, execType string) string {
	return fmt.Sprintf("%s-%s-%s", namespacedName.Namespace, namespacedName.Name, execType)
}

func (r *ShardingSphereChaosReconciler) getNeedExec(chao *v1alpha1.ShardingSphereChaos) *pressure.Pressure {
	jobName := getExecName(chao)

	//if pressure do not exist,run it
	for i := range r.ExecCtrls {
		if r.ExecCtrls[i].pressure.Name == jobName {
			return r.ExecCtrls[i].pressure
		}
	}

	return nil
}

func (r *ShardingSphereChaosReconciler) getRuntimeChaos(ctx context.Context, name types.NamespacedName) (*v1alpha1.ShardingSphereChaos, error) {
	var rt = &v1alpha1.ShardingSphereChaos{}
	err := r.Get(ctx, name, rt)
	return rt, err
}

// nolint:nestif
func (r *ShardingSphereChaosReconciler) finalize(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos) (ctrl.Result, error) {
	namespacedName := types.NamespacedName{
		Namespace: ssChaos.Namespace,
		Name:      ssChaos.Name,
	}
	r.deleteExec(namespacedName)
	if err := r.deleteExternalResources(ctx, ssChaos); err != nil {
		return ctrl.Result{}, err
	}

	controllerutil.RemoveFinalizer(ssChaos, SSChaosFinalizeName)
	if err := r.Update(ctx, ssChaos); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ShardingSphereChaosReconciler) deleteExternalResources(ctx context.Context, chao *v1alpha1.ShardingSphereChaos) error {
	nameSpacedName := types.NamespacedName{Namespace: chao.Namespace, Name: chao.Name}
	if chao.Spec.EmbedChaos.PodChaos != nil {
		if err := r.deletePodChaos(ctx, nameSpacedName); err != nil {
			return err
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

func (r *ShardingSphereChaosReconciler) deleteExec(namespacedName types.NamespacedName) {
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

func (r *ShardingSphereChaosReconciler) deletePodChaos(ctx context.Context, namespacedName types.NamespacedName) error {
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

func (r *ShardingSphereChaosReconciler) deleteNetworkChaos(ctx context.Context, namespacedName types.NamespacedName) error {
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

func (r *ShardingSphereChaosReconciler) reconcileChaos(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos) error {
	logger := r.Log.WithValues("reconcile shardingspherechaos", fmt.Sprintf("%s/%s", chaos.Namespace, chaos.Name))

	if chaos.Status.Phase == "" || chaos.Status.Phase == v1alpha1.BeforeSteady || chaos.Status.Phase == v1alpha1.AfterSteady {
		return nil
	}
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

func (r *ShardingSphereChaosReconciler) reconcilePodChaos(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos, namespacedName types.NamespacedName) error {
	pc, err := r.getPodChaosByNamespacedName(ctx, namespacedName)
	if err != nil {
		return err
	}
	if pc != nil {
		return r.updatePodChaos(ctx, chaos, pc)
	}

	return r.createPodChaos(ctx, chaos)
}

func (r *ShardingSphereChaosReconciler) getPodChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (chaosmesh.PodChaos, error) {
	pc, err := r.Chaos.GetPodChaosByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, err
	}
	return pc, nil
}

func (r *ShardingSphereChaosReconciler) createPodChaos(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos) error {
	err := r.Chaos.CreatePodChaos(ctx, chaos)
	if err != nil {
		return err
	}
	r.Events.Event(chaos, "Normal", "Created", fmt.Sprintf("PodChaos %s", " is created successfully"))
	return nil
}

func (r *ShardingSphereChaosReconciler) updatePodChaos(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos, podChaos chaosmesh.PodChaos) error {
	err := r.Chaos.UpdatePodChaos(ctx, podChaos, chaos)
	if err != nil {
		return err
	}

	return nil
}

func (r *ShardingSphereChaosReconciler) reconcileNetworkChaos(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos, namespacedName types.NamespacedName) error {
	nc, err := r.getNetworkChaosByNamespacedName(ctx, namespacedName)
	if err != nil {
		return err
	}
	if nc != nil {
		return r.updateNetWorkChaos(ctx, chaos, nc)
	}

	return r.createNetworkChaos(ctx, chaos)
}

func (r *ShardingSphereChaosReconciler) updateNetWorkChaos(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos, networkChaos chaosmesh.NetworkChaos) error {
	err := r.Chaos.UpdateNetworkChaos(ctx, networkChaos, chaos)
	if err != nil {
		return err
	}

	return nil
}

func (r *ShardingSphereChaosReconciler) createNetworkChaos(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos) error {
	err := r.Chaos.CreateNetworkChaos(ctx, chaos)
	if err != nil {
		return err
	}

	r.Events.Event(chaos, "Normal", "created", fmt.Sprintf("networkChaos %s", "  is created successfully"))
	return nil
}

func (r *ShardingSphereChaosReconciler) reconcileConfigMap(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos) error {
	namespaceName := types.NamespacedName{
		Namespace: chaos.Namespace,
		Name:      chaos.Name,
	}

	cm, err := r.getConfigMapByNamespacedName(ctx, namespaceName)
	if err != nil {
		return err
	}

	if cm != nil {
		if err := r.updateConfigMap(ctx, chaos, cm); err != nil {
			fmt.Printf("update configmap error: %s\n", err)
			return err
		}
	}

	if err = r.createConfigMap(ctx, chaos); err != nil {
		fmt.Printf("create configmap error: %s\n", err)
		return err
	}

	return nil
}

func setDefaultStatus(chaos *v1alpha1.ShardingSphereChaos) {

	if chaos.Status.Phase == "" {
		chaos.Status.Phase = v1alpha1.BeforeSteady
	}

	if len(chaos.Status.Conditions) == 0 {
		chaos.Status.Conditions = []*metav1.Condition{
			{
				Type:               string(v1alpha1.BeforeSteady),
				Status:             metav1.ConditionTrue,
				LastTransitionTime: metav1.Time{Time: time.Now()},
				Reason:             "InSteadyExperiment",
			},
			{
				Type:               string(v1alpha1.AfterSteady),
				Status:             metav1.ConditionFalse,
				LastTransitionTime: metav1.Time{Time: time.Now()},
				Reason:             "AfterSteadyExperiment",
			},
			{
				Type:               string(v1alpha1.BeforeChaos),
				Status:             metav1.ConditionFalse,
				LastTransitionTime: metav1.Time{Time: time.Now()},
				Reason:             "InChaoExperiment",
			},
			{
				Type:               string(v1alpha1.AfterChaos),
				Status:             metav1.ConditionFalse,
				LastTransitionTime: metav1.Time{Time: time.Now()},
				Reason:             "AfterChaosExperiment",
			},
		}
	}
}

func (r *ShardingSphereChaosReconciler) updateChaosCondition(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos) error {
	namespacedName := types.NamespacedName{
		Namespace: chaos.Namespace,
		Name:      chaos.Name,
	}

	if chaos.Spec.EmbedChaos.PodChaos != nil {
		pc, err := r.Chaos.GetPodChaosByNamespacedName(ctx, namespacedName)
		if err != nil {
			return err
		}
		chaos.Status.ChaosCondition = chaosmesh.ConvertChaosStatus(ctx, chaos, pc)
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

func (r *ShardingSphereChaosReconciler) getNetworkChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (chaosmesh.NetworkChaos, error) {
	nc, err := r.Chaos.GetNetworkChaosByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, err
	}
	return nc, nil
}

func (r *ShardingSphereChaosReconciler) getConfigMapByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*corev1.ConfigMap, error) {
	config, err := r.ConfigMap.GetByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (r *ShardingSphereChaosReconciler) updateConfigMap(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos, cur *corev1.ConfigMap) error {
	// exp := sschaos.UpdateShardingSphereChaosConfigMap(chao, cur)
	exp := r.ConfigMap.Build(ctx, chaos)
	exp.ObjectMeta = cur.ObjectMeta
	exp.ObjectMeta.ResourceVersion = ""
	exp.Labels = cur.Labels
	exp.Annotations = cur.Annotations
	return r.ConfigMap.Update(ctx, exp)
}

func (r *ShardingSphereChaosReconciler) createConfigMap(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos) error {
	cm := r.ConfigMap.Build(ctx, chaos)
	if err := ctrl.SetControllerReference(chaos, cm, r.Scheme); err != nil {
		return err
	}

	err := r.Create(ctx, cm)
	if err != nil && apierrors.IsAlreadyExists(err) {
		return nil
	}

	return err
}

// SetupWithManager sets up the controller with the Manager.
func (r *ShardingSphereChaosReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ShardingSphereChaos{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&batchV1.Job{}).
		Complete(r)
}
