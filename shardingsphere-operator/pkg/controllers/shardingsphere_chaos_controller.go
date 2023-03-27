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
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	chaosv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ShardingSphereChaosReconciler is a controller for the ShardingSphereChaos
type ShardingSphereChaosReconciler struct { //
	client.Client
	Scheme *runtime.Scheme
}

// Reconcile handles main function of this controller
func (r *ShardingSphereChaosReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.WithValues("ShardingSphereChaos", req.NamespacedName)

	var ssChaos v1alpha1.ShardingSphereChaos
	if err := r.Get(ctx, req.NamespacedName, &ssChaos); err != nil {
		logger.Error(err, "unable to fetch ShardingSphereChaos source")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	if !ssChaos.ObjectMeta.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, nil
	}

	//todo: add inject check status here

	constructAndCreateChao := func(ctx context.Context, template *v1alpha1.ShardingSphereChaos, req ctrl.Request) error {

		var CreateErr error
		switch template.Spec.ChaosKind {
		case v1alpha1.NetworkChaosKind:
			CreateErr = r.CreateNetworkChaos(ctx, template, template.Name, template.Namespace)
		case v1alpha1.PodChaosKind:
			CreateErr = r.CreatePodChaos(ctx, template, template.Name, template.Namespace)
		case v1alpha1.WorkFlowKind:
			CreateErr = r.CreateWorkFlow(ctx, template, template.Name, template.Namespace)
		}

		return CreateErr
	}

	if err := constructAndCreateChao(ctx, &ssChaos, req); err != nil && !apierrors.IsAlreadyExists(err) {
		return ctrl.Result{}, err
	}

	//todo:update sschaos status to judge inject or verify 

	return ctrl.Result{}, nil
}

func (r *ShardingSphereChaosReconciler) CreateWorkFlow(ctx context.Context, chao *v1alpha1.ShardingSphereChaos, name string, namespace string) error {
	workflow := chao.Spec.Workflow
	if workflow == nil {
		return fmt.Errorf("workflow not defined in spec")
	}
	meshWorkflow := &chaosv1alpha1.Workflow{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      make(map[string]string),
			Annotations: make(map[string]string),
			Name:        name,
			Namespace:   namespace,
		},
		Spec: chaosv1alpha1.WorkflowSpec{
			Entry:     workflow.Entry,
			Templates: deepCopyWorkTemplates(workflow.Templates),
		},
	}

	if err := ctrl.SetControllerReference(chao, meshWorkflow, r.Scheme); err != nil {
		return err
	}
	return r.Create(ctx, meshWorkflow)
}

func deepCopyWorkTemplates(templates []v1alpha1.WorkFlowTemplate) (meshTemplates []chaosv1alpha1.Template) {

	for _, v := range templates {
		meshTemplate := &chaosv1alpha1.Template{
			Name:                v.Name,
			Type:                chaosv1alpha1.TemplateType(v.Type),
			Deadline:            v.Deadline,
			Task:                deepCopyWorkFlowTask(v.Task),
			Children:            v.Children,
			ConditionalBranches: deepCopyConditionalBranches(v.ConditionalBranches),
			EmbedChaos:          deepCopyEmbedChaos(v.EmbedChaos),
			Schedule:            deepCopyScheduleSpec(v.Schedule),
		}
		meshTemplates = append(meshTemplates, *meshTemplate)
	}

	return
}

func deepCopyWorkFlowTask(task *v1alpha1.Task) *chaosv1alpha1.Task {
	meshTask := &chaosv1alpha1.Task{
		Container: task.Container,
		Volumes:   task.Volumes,
	}

	return meshTask
}

func deepCopyScheduleSpec(schedule *v1alpha1.ChaosOnlyScheduleSpec) *chaosv1alpha1.ChaosOnlyScheduleSpec {
	meshSchedule := &chaosv1alpha1.ChaosOnlyScheduleSpec{
		Schedule:                schedule.Schedule,
		StartingDeadlineSeconds: schedule.StartingDeadlineSeconds,
		ConcurrencyPolicy:       chaosv1alpha1.ConcurrencyPolicy(schedule.ConcurrencyPolicy),
		HistoryLimit:            schedule.HistoryLimit,
		Type:                    chaosv1alpha1.ScheduleTemplateType(schedule.Type),
		EmbedChaos:              *deepCopyEmbedChaos(&schedule.EmbedChaos),
	}

	return meshSchedule
}

func deepCopyConditionalBranches(branches []v1alpha1.ConditionalBranch) (meshBranches []chaosv1alpha1.ConditionalBranch) {
	for i := range branches {
		meshBranch := chaosv1alpha1.ConditionalBranch{
			Target:     branches[i].Target,
			Expression: branches[i].Expression,
		}
		meshBranches = append(meshBranches, meshBranch)
	}

	return
}

func deepCopyEmbedChaos(chaos *v1alpha1.EmbedChaos) *chaosv1alpha1.EmbedChaos {
	var meshEmbedChaos *chaosv1alpha1.EmbedChaos

	if chaos.NetworkChaos != nil {
		meshEmbedChaos.NetworkChaos = deepCopyNetworkSpec(chaos.NetworkChaos)
	}

	if chaos.PodChaos != nil {
		meshEmbedChaos.PodChaos = deepCopyPodChaoSpec(chaos.PodChaos)
	}

	return meshEmbedChaos
}

func (r *ShardingSphereChaosReconciler) CreatePodChaos(ctx context.Context, chao *v1alpha1.ShardingSphereChaos, name string, namespace string) error {
	podChao := chao.Spec.PodChaos
	if podChao == nil {
		return fmt.Errorf("podchaos not defined in spec")
	}
	meshPodChao := &chaosv1alpha1.PodChaos{

		ObjectMeta: metav1.ObjectMeta{
			Labels:      make(map[string]string),
			Annotations: make(map[string]string),
			Name:        name,
			Namespace:   namespace,
		},
	}
	meshPodChao.Spec = *deepCopyPodChaoSpec(podChao)
	if err := ctrl.SetControllerReference(chao, meshPodChao, r.Scheme); err != nil {
		return err
	}
	return r.Create(ctx, meshPodChao)
}

func deepCopyPodChaoSpec(podChao *v1alpha1.PodChaosSpec) *chaosv1alpha1.PodChaosSpec {
	spec := chaosv1alpha1.PodChaosSpec{
		ContainerSelector: *deepCopyContainerSelector(&podChao.ContainerSelector),
		Action:            chaosv1alpha1.PodChaosAction(podChao.Action),
		Duration:          podChao.Duration,
		GracePeriod:       podChao.GracePeriod,
	}

	return &spec
}

func (r *ShardingSphereChaosReconciler) CreateNetworkChaos(ctx context.Context, chao *v1alpha1.ShardingSphereChaos, name string, namespace string) error {
	networkChao := chao.Spec.NetworkChaos
	if networkChao == nil {
		return fmt.Errorf("networkchao not defined in spec")
	}
	meshNetworkChao := &chaosv1alpha1.NetworkChaos{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      make(map[string]string),
			Annotations: make(map[string]string),
			Name:        name,
			Namespace:   namespace,
		},
	}
	meshNetworkChao.Spec = *deepCopyNetworkSpec(networkChao)
	if err := ctrl.SetControllerReference(chao, meshNetworkChao, r.Scheme); err != nil {
		return err
	}
	return r.Create(ctx, meshNetworkChao)
}

func deepCopyNetworkSpec(networkChao *v1alpha1.NetworkChaosSpec) *chaosv1alpha1.NetworkChaosSpec {
	spec := chaosv1alpha1.NetworkChaosSpec{
		PodSelector:     *deepCopyPodSelector(&networkChao.PodSelector),
		Action:          chaosv1alpha1.NetworkChaosAction(networkChao.Action),
		Device:          networkChao.Device,
		Duration:        networkChao.Duration,
		TcParameter:     *deepCopyTcParameter(&networkChao.TcParameter),
		Direction:       chaosv1alpha1.Direction(networkChao.Direction),
		Target:          deepCopyPodSelector(networkChao.Target),
		TargetDevice:    networkChao.TargetDevice,
		ExternalTargets: networkChao.ExternalTargets,
	}

	return &spec
}

func deepCopyTcParameter(tcParam *v1alpha1.TcParameter) *chaosv1alpha1.TcParameter {
	chaoTcParam := &chaosv1alpha1.TcParameter{}
	if tcParam.Delay != nil {
		chaoTcParam.Delay = &chaosv1alpha1.DelaySpec{
			Latency:     tcParam.Delay.Latency,
			Correlation: tcParam.Delay.Correlation,
			Jitter:      tcParam.Delay.Jitter,
			Reorder:     deepCopyRecorderSpec(tcParam.Delay.Reorder),
		}
	}

	if tcParam.Corrupt != nil {
		chaoTcParam.Corrupt = &chaosv1alpha1.CorruptSpec{
			Corrupt:     tcParam.Corrupt.Corrupt,
			Correlation: tcParam.Corrupt.Correlation,
		}
	}

	if tcParam.Bandwidth != nil {
		chaoTcParam.Bandwidth = &chaosv1alpha1.BandwidthSpec{
			Rate:     tcParam.Bandwidth.Rate,
			Limit:    tcParam.Bandwidth.Limit,
			Buffer:   tcParam.Bandwidth.Buffer,
			Peakrate: tcParam.Bandwidth.Peakrate,
			Minburst: tcParam.Bandwidth.Minburst,
		}
	}

	if tcParam.Loss != nil {
		chaoTcParam.Loss = &chaosv1alpha1.LossSpec{
			Loss:        tcParam.Loss.Loss,
			Correlation: tcParam.Loss.Correlation,
		}
	}

	if tcParam.Duplicate != nil {
		chaoTcParam.Duplicate = &chaosv1alpha1.DuplicateSpec{
			Duplicate:   tcParam.Duplicate.Duplicate,
			Correlation: tcParam.Duplicate.Correlation,
		}
	}

	return chaoTcParam
}

func deepCopyRecorderSpec(recorder *v1alpha1.ReorderSpec) *chaosv1alpha1.ReorderSpec {
	return &chaosv1alpha1.ReorderSpec{
		Reorder:     recorder.Reorder,
		Correlation: recorder.Correlation,
		Gap:         recorder.Gap,
	}
}

func deepCopyContainerSelector(selector *v1alpha1.ContainerSelector) *chaosv1alpha1.ContainerSelector {
	return &chaosv1alpha1.ContainerSelector{
		PodSelector:    *deepCopyPodSelector(&selector.PodSelector),
		ContainerNames: selector.ContainerNames,
	}
}

func deepCopyPodSelector(selector *v1alpha1.PodSelector) *chaosv1alpha1.PodSelector {
	return &chaosv1alpha1.PodSelector{
		Selector: chaosv1alpha1.PodSelectorSpec{
			GenericSelectorSpec: chaosv1alpha1.GenericSelectorSpec{
				Namespaces:          selector.Selector.GenericSelectorSpec.Namespaces,
				FieldSelectors:      selector.Selector.GenericSelectorSpec.FieldSelectors,
				LabelSelectors:      selector.Selector.GenericSelectorSpec.LabelSelectors,
				ExpressionSelectors: chaosv1alpha1.LabelSelectorRequirements(selector.Selector.GenericSelectorSpec.ExpressionSelectors),
				AnnotationSelectors: selector.Selector.GenericSelectorSpec.AnnotationSelectors,
			},
			Nodes:             selector.Selector.Nodes,
			Pods:              selector.Selector.Pods,
			NodeSelectors:     selector.Selector.NodeSelectors,
			PodPhaseSelectors: selector.Selector.PodPhaseSelectors,
		},
		Mode:  chaosv1alpha1.SelectorMode(selector.Mode),
		Value: selector.Value,
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *ShardingSphereChaosReconciler) SetupWithManager(mgr ctrl.Manager) error {

	//todo: and inject and verify crd(job) here
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ShardingSphereChaos{}).
		Owns(&chaosv1alpha1.PodChaos{}).
		Owns(&chaosv1alpha1.NetworkChaos{}).
		Owns(&chaosv1alpha1.Workflow{}).
		Complete(r)
}
