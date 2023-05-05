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
	"errors"
	"fmt"
	"strings"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	sschaos "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaosmesh"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/job"
	reconcile "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/shardingspherechaos"

	"github.com/go-logr/logr"
	batchV1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/retry"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ShardingSphereChaosControllerName = "shardingsphere-chaos-controller"
	VerifyJobCheck                    = "Verify"
	SSChaosFinalizeName               = "shardingsphere.apache.org/finalizer"
)

type JobCondition string

var (
	CompleteJob JobCondition = "complete"
	FailureJob  JobCondition = "failure"
	SuspendJob  JobCondition = "suspend"
	ActiveJob   JobCondition = "active"

	ErrNoPod = errors.New("no pod in list")
)

// ShardingSphereChaosReconciler is a controller for the ShardingSphereChaos
type ShardingSphereChaosReconciler struct {
	client.Client

	Scheme    *runtime.Scheme
	Log       logr.Logger
	ClientSet *clientset.Clientset
	Events    record.EventRecorder

	Chaos     sschaos.Chaos
	Job       job.Job
	ConfigMap configmap.ConfigMap
}

// Reconcile handles main function of this controller
func (r *ShardingSphereChaosReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues(ShardingSphereChaosControllerName, req.NamespacedName)

	ssChaos, err := r.getRuntimeChaos(ctx, req.NamespacedName)

	if err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{RequeueAfter: defaultRequeueTime}, nil
		}

		logger.Error(err, "failed to get the shardingsphere chaos")
		return ctrl.Result{Requeue: true}, err
	}

	if err := r.finalize(ctx, ssChaos); err != nil {
		return ctrl.Result{Requeue: true}, err
	}

	logger.Info("start reconcile chaos")

	//TODO: consider merge these events
	if err := r.reconcileChaos(ctx, ssChaos); err != nil {
		logger.Error(err, "reconcile shardingspherechaos error")
		r.Events.Event(ssChaos, "Warning", "shardingspherechaos error", err.Error())
	}

	if err := r.reconcileConfigMap(ctx, ssChaos); err != nil {
		logger.Error(err, "reconcile configmap error")
		r.Events.Event(ssChaos, "Warning", "configmap error", err.Error())
	}

	if err := r.reconcileJob(ctx, ssChaos); err != nil {
		logger.Error(err, "reconcile job error")
		r.Events.Event(ssChaos, "Warning", "job error", err.Error())
	}

	if err := r.reconcileStatus(ctx, ssChaos); err != nil {
		logger.Error(err, "failed to update status")
		r.Events.Event(ssChaos, "Warning", "update status error", err.Error())
	}

	return ctrl.Result{RequeueAfter: defaultRequeueTime}, nil
}

func (r *ShardingSphereChaosReconciler) getRuntimeChaos(ctx context.Context, name types.NamespacedName) (*v1alpha1.ShardingSphereChaos, error) {
	var rt = &v1alpha1.ShardingSphereChaos{}
	err := r.Get(ctx, name, rt)
	return rt, err
}

// nolint:nestif
func (r *ShardingSphereChaosReconciler) finalize(ctx context.Context, chao *v1alpha1.ShardingSphereChaos) error {
	if chao.ObjectMeta.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(chao, SSChaosFinalizeName) {
			controllerutil.AddFinalizer(chao, SSChaosFinalizeName)
			if err := r.Update(ctx, chao); err != nil {
				return err
			}
		}
	} else if controllerutil.ContainsFinalizer(chao, SSChaosFinalizeName) {
		if err := r.deleteExternalResources(ctx, chao); err != nil {
			return err
		}

		controllerutil.RemoveFinalizer(chao, SSChaosFinalizeName)
		if err := r.Update(ctx, chao); err != nil {
			return err
		}
	}

	return nil
}

func (r *ShardingSphereChaosReconciler) deleteExternalResources(ctx context.Context, chao *v1alpha1.ShardingSphereChaos) error {
	nameSpacedName := types.NamespacedName{Namespace: chao.Namespace, Name: chao.Name}
	if chao.Spec.EmbedChaos.PodChaos != nil {
		podchao, err := r.getPodChaosByNamespacedName(ctx, nameSpacedName)
		if err != nil {
			return err
		}
		if podchao != nil {
			if err := r.Chaos.DeletePodChaos(ctx, podchao); err != nil {
				return err
			}
		}
	}

	if chao.Spec.EmbedChaos.NetworkChaos != nil {
		networkchao, err := r.getNetworkChaosByNamespacedName(ctx, nameSpacedName)
		if err != nil {
			return err
		}
		if networkchao != nil {
			if err := r.Chaos.DeleteNetworkChaos(ctx, networkchao); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *ShardingSphereChaosReconciler) reconcileChaos(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos) error {
	logger := r.Log.WithValues("reconcile shardingspherechaos", fmt.Sprintf("%s/%s", chaos.Namespace, chaos.Name))

	if len(chaos.Status.Phase) == 0 || chaos.Status.Phase == v1alpha1.BeforeExperiment {
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

func (r *ShardingSphereChaosReconciler) getPodChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (sschaos.PodChaos, error) {
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

func (r *ShardingSphereChaosReconciler) updatePodChaos(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos, podChaos sschaos.PodChaos) error {
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

func (r *ShardingSphereChaosReconciler) updateNetWorkChaos(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos, networkChaos sschaos.NetworkChaos) error {
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
		// return r.updateConfigMap(ctx, chaos, cm)
	}

	if err = r.createConfigMap(ctx, chaos); err != nil {
		fmt.Printf("create configmap error: %s\n", err)
		return err
	}

	return nil
}

func (r *ShardingSphereChaosReconciler) reconcileJob(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos) error {
	var nowInjectRequirement reconcile.InjectRequirement

	switch chaos.Status.Phase {
	case v1alpha1.InjectedChaos:
		nowInjectRequirement = reconcile.Pressure
	case v1alpha1.RecoveredChaos:
		nowInjectRequirement = reconcile.Verify
	case v1alpha1.BeforeExperiment, v1alpha1.AfterExperiment:
		fallthrough
	default:
		nowInjectRequirement = reconcile.Experimental
	}

	namespaceName := types.NamespacedName{
		Namespace: chaos.Namespace,
		Name:      reconcile.MakeJobName(chaos.Name, nowInjectRequirement),
	}

	job, err := r.getJobByNamespacedName(ctx, namespaceName)
	if err != nil {
		return err
	}

	if job != nil {
		return r.updateJob(ctx, nowInjectRequirement, chaos, job)
	}

	return r.createJob(ctx, nowInjectRequirement, chaos)
}

func (r *ShardingSphereChaosReconciler) reconcileStatus(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos) error {
	namespacedName := types.NamespacedName{
		Name:      chaos.Name,
		Namespace: chaos.Namespace,
	}
	chaos, err := r.getRuntimeChaos(ctx, namespacedName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil
		}
		return err
	}
	r.setDefaultStatus(chaos)

	req := getInjectRequirement(chaos)
	job, err := r.getJobByNamespacedName(ctx, types.NamespacedName{Namespace: chaos.Namespace, Name: reconcile.MakeJobName(chaos.Name, req)})
	if err != nil || job == nil {
		return err
	}

	if chaos.Status.Phase == v1alpha1.BeforeExperiment && job.Status.Succeeded == 1 {
		chaos.Status.Phase = v1alpha1.AfterExperiment
	}

	condition := getJobCondition(job.Status.Conditions)
	if condition == FailureJob {
		r.Events.Event(chaos, "Warning", "failed", fmt.Sprintf("job: %s", job.Name))
	}

	if chaos.Status.Phase == v1alpha1.RecoveredChaos {
		if err := r.updateRecoveredJob(ctx, chaos, job); err != nil {
			r.Events.Event(chaos, "Warning", "getPodLog", err.Error())
			return err
		}
	}

	if err := r.updatePhaseStart(ctx, chaos); err != nil {
		return err
	}

	// sts := setRtStatus(chaos)
	rt, err := r.getRuntimeChaos(ctx, namespacedName)
	if err != nil {
		return err
	}
	rt.Status = chaos.Status

	return r.Status().Update(ctx, rt)
}

func (r *ShardingSphereChaosReconciler) setDefaultStatus(chaos *v1alpha1.ShardingSphereChaos) {
	if chaos.Status.Phase == "" {
		chaos.Status.Phase = v1alpha1.BeforeExperiment
	}
	if chaos.Status.Results == nil {
		chaos.Status.Results = []v1alpha1.Result{}
	}
}

// getInjectRequirement to get the coming job requirement
// * BeforeExperiment: it hasn't been started, could start a new experiment
// * AfterExperiment: it has been finished, could start a new experiment
// * InjectChaos: it has been started, could start some pressure
// * recoveredChaos: it has been recovered, could start to verify

func getInjectRequirement(ssChaos *v1alpha1.ShardingSphereChaos) reconcile.InjectRequirement {
	var jobName reconcile.InjectRequirement

	if ssChaos.Status.Phase == v1alpha1.BeforeExperiment || ssChaos.Status.Phase == v1alpha1.AfterExperiment {
		jobName = reconcile.Experimental
	}

	if ssChaos.Status.Phase == v1alpha1.InjectedChaos {
		jobName = reconcile.Pressure
	}

	if ssChaos.Status.Phase == v1alpha1.RecoveredChaos {
		jobName = reconcile.Verify
	}

	return jobName
}

func (r *ShardingSphereChaosReconciler) getJobByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*batchV1.Job, error) {
	job, err := r.Job.GetByNamespacedName(ctx, namespacedName)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func getJobCondition(conditions []batchV1.JobCondition) JobCondition {
	var ret = ActiveJob
	for i := range conditions {
		p := &conditions[i]
		switch {
		case p.Type == batchV1.JobComplete && p.Status == corev1.ConditionTrue:
			ret = CompleteJob
		case p.Type == batchV1.JobFailed && p.Status == corev1.ConditionTrue:
			ret = FailureJob
		case p.Type == batchV1.JobSuspended && p.Status == corev1.ConditionTrue:
			ret = SuspendJob
		case p.Type == batchV1.JobFailureTarget:
			ret = FailureJob
		}

	}
	return ret
}

func isRecoveredJobType(rJob *batchV1.Job, requirement reconcile.InjectRequirement) bool {
	for i := range rJob.Spec.Template.Spec.Containers[0].Args {
		r := rJob.Spec.Template.Spec.Containers[0].Args[i]
		if strings.Contains(r, string(requirement)) {
			return true
		}
	}
	return false
}

func (r *ShardingSphereChaosReconciler) updateRecoveredJob(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos, rJob *batchV1.Job) error {
	if !isRecoveredJobType(rJob, reconcile.Verify) {
		return nil
	}

	for i := range ssChaos.Status.Results {
		if strings.HasPrefix(ssChaos.Status.Results[i].Detail.Message, VerifyJobCheck) {
			return nil
		}
	}

	logOpts := &corev1.PodLogOptions{}
	pod, err := r.getPodHaveLog(ctx, rJob)
	if err != nil || pod == nil {
		return err
	}
	podNamespacedName := types.NamespacedName{
		Namespace: pod.Namespace,
		Name:      pod.Name,
	}
	condition := getJobCondition(rJob.Status.Conditions)
	result := &v1alpha1.Result{}

	if condition == CompleteJob {
		log, err := r.getPodLog(ctx, podNamespacedName, logOpts)
		if err != nil {
			return err
		}
		if ssChaos.Spec.Expect.Verify == "" || ssChaos.Spec.Expect.Verify == log {
			result.Success = true
			result.Detail = v1alpha1.Detail{
				Time:    metav1.Time{Time: time.Now()},
				Message: fmt.Sprintf("%s: job succeeded", VerifyJobCheck),
			}
		} else {
			result.Success = false
			result.Detail = v1alpha1.Detail{
				Time:    metav1.Time{Time: time.Now()},
				Message: fmt.Sprintf("%s: %s", VerifyJobCheck, log),
			}
		}
		ssChaos.Status.Results = updateResult(ssChaos.Status.Results, *result, VerifyJobCheck)
	}

	if condition == FailureJob {
		log, err := r.getPodLog(ctx, podNamespacedName, logOpts)
		if err != nil {
			return err
		}
		result.Success = false
		result.Detail = v1alpha1.Detail{
			Time:    metav1.Time{Time: time.Now()},
			Message: fmt.Sprintf("%s: %s", VerifyJobCheck, log),
		}
		ssChaos.Status.Results = updateResult(ssChaos.Status.Results, *result, VerifyJobCheck)
	}

	return nil
}

func (r *ShardingSphereChaosReconciler) getPodHaveLog(ctx context.Context, rJob *batchV1.Job) (*corev1.Pod, error) {
	pods := &corev1.PodList{}

	if err := r.List(ctx, pods, client.MatchingLabels{"controller-uid": rJob.Spec.Template.Labels["controller-uid"]}); err != nil {
		return nil, err
	}
	if pods.Items == nil {
		return nil, nil
	}
	//FIXME: get the first pod
	var pod *corev1.Pod
	for i := range pods.Items {
		pod = &pods.Items[i]
		break
	}
	return pod, nil
}

// FIXME: this will broke when the job count is more than one
func updateResult(results []v1alpha1.Result, r v1alpha1.Result, check string) []v1alpha1.Result {
	for i := range results {
		msg := results[i].Detail.Message
		if strings.HasPrefix(msg, check) && strings.HasPrefix(r.Detail.Message, check) {
			results[i] = r
			return results
		}
	}
	results = append(results, r)
	return results
}

// FIXME: this will broke if the log is too long
func (r *ShardingSphereChaosReconciler) getPodLog(ctx context.Context, namespacedName types.NamespacedName, options *corev1.PodLogOptions) (string, error) {
	req := r.ClientSet.CoreV1().Pods(namespacedName.Namespace).GetLogs(namespacedName.Name, options)
	res := req.Do(ctx)
	if res.Error() != nil {
		return "", res.Error()
	}
	var ret []byte
	ret, err := res.Raw()
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func (r *ShardingSphereChaosReconciler) updatePhaseStart(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos) error {
	if ssChaos.Status.Phase != v1alpha1.BeforeExperiment {
		if err := r.updateChaosCondition(ctx, ssChaos); err != nil {
			return err
		}

		if ssChaos.Status.ChaosCondition == v1alpha1.AllInjected && ssChaos.Status.Phase == v1alpha1.AfterExperiment {
			ssChaos.Status.Phase = v1alpha1.InjectedChaos
		}

		if ssChaos.Status.ChaosCondition == v1alpha1.AllRecovered && ssChaos.Status.Phase == v1alpha1.InjectedChaos {
			ssChaos.Status.Phase = v1alpha1.RecoveredChaos
		}
	}

	return nil
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
		chaos.Status.ChaosCondition = sschaos.ConvertChaosStatus(ctx, chaos, pc)
	}

	if chaos.Spec.EmbedChaos.NetworkChaos != nil {
		nc, err := r.Chaos.GetNetworkChaosByNamespacedName(ctx, namespacedName)
		if err != nil {
			return err
		}
		chaos.Status.ChaosCondition = sschaos.ConvertChaosStatus(ctx, chaos, nc)
	}

	return nil
}

func (r *ShardingSphereChaosReconciler) getNetworkChaosByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (sschaos.NetworkChaos, error) {
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
	// exp := reconcile.UpdateShardingSphereChaosConfigMap(chao, cur)
	exp := r.ConfigMap.Build(ctx, chaos)
	exp.ObjectMeta = cur.ObjectMeta
	exp.ObjectMeta.ResourceVersion = ""
	exp.Labels = cur.Labels
	exp.Annotations = cur.Annotations
	return r.ConfigMap.Update(ctx, exp)

	// return r.Update(ctx, exp)
}

func (r *ShardingSphereChaosReconciler) createConfigMap(ctx context.Context, chaos *v1alpha1.ShardingSphereChaos) error {
	/*
		cm := reconcile.NewSSConfigMap(chaos)
		if err := ctrl.SetControllerReference(chaos, cm, r.Scheme); err != nil {
			return err
		}
	*/

	cm := r.ConfigMap.Build(ctx, chaos)

	err := r.Create(ctx, cm)
	if err != nil && apierrors.IsAlreadyExists(err) {
		return nil
	}

	return err
}

// TODO: consider a new job name pattern
func (r *ShardingSphereChaosReconciler) updateJob(ctx context.Context, requirement reconcile.InjectRequirement, chao *v1alpha1.ShardingSphereChaos, cur *batchV1.Job) error {
	isEqual, err := reconcile.IsJobChanged(chao, requirement, cur)
	if err != nil {
		return err
	}
	if !isEqual {
		if err := r.Delete(ctx, cur); err != nil && !apierrors.IsNotFound(err) {
			return err
		}
	}
	return nil
}

func (r *ShardingSphereChaosReconciler) createJob(ctx context.Context, requirement reconcile.InjectRequirement, chao *v1alpha1.ShardingSphereChaos) error {
	injectJob, err := reconcile.NewJob(chao, requirement)
	if err != nil {
		return err
	}
	if err := ctrl.SetControllerReference(chao, injectJob, r.Scheme); err != nil {
		return err
	}

	err = r.Create(ctx, injectJob)
	if err != nil {
		return client.IgnoreAlreadyExists(err)
	}
	//FIXME: consider remove the following L620-L676, perhaps don't need to pay much attention to the pod
	rJob := &batchV1.Job{}
	backoff := wait.Backoff{
		Steps:    6,
		Duration: 500 * time.Millisecond,
		Factor:   5.0,
		Jitter:   0.1,
	}

	if err := retry.OnError(
		backoff,
		func(e error) bool {
			return true
		},
		func() error {
			return r.Get(ctx, types.NamespacedName{Namespace: chao.Namespace, Name: reconcile.MakeJobName(chao.Name, requirement)}, rJob)
		},
	); err != nil {
		return err
	}

	podList := &corev1.PodList{}
	if err := retry.OnError(backoff, func(e error) bool {
		return e != nil
	}, func() error {
		if err := r.List(ctx, podList, client.MatchingLabels{"controller-uid": rJob.Spec.Template.Labels["controller-uid"]}); err != nil {
			return err
		}
		if len(podList.Items) == 0 {
			return ErrNoPod
		}
		return nil
	}); err != nil {
		return err
	}

	for i := range podList.Items {
		rPod := &podList.Items[i]
		if err := ctrl.SetControllerReference(rJob, rPod, r.Scheme); err != nil {
			return err
		}

		exp := rPod.DeepCopy()
		updateBackoff := wait.Backoff{
			Steps:    5,
			Duration: 30 * time.Millisecond,
			Factor:   5.0,
			Jitter:   0.1,
		}
		if err := retry.RetryOnConflict(updateBackoff, func() error {
			if err := r.Update(ctx, exp); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}
	}

	r.Events.Event(chao, "Normal", "Created", fmt.Sprintf("%s job created", requirement))
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ShardingSphereChaosReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ShardingSphereChaos{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&batchV1.Job{}).
		Complete(r)
}
