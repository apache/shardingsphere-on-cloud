package ShardingSphereChaos

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
import (
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	chaosv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewPodChaos(chao *v1alpha1.ShardingSphereChaos) *chaosv1alpha1.PodChaos {
	podChao := chao.Spec.PodChaos
	meshPodChao := &chaosv1alpha1.PodChaos{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      make(map[string]string),
			Annotations: make(map[string]string),
			Name:        chao.Name,
			Namespace:   chao.Namespace,
		},
	}
	meshPodChao.Spec = *deepCopyPodChaoSpec(podChao)
	return meshPodChao
}

func NewNetworkPodChaos(chao *v1alpha1.ShardingSphereChaos) *chaosv1alpha1.NetworkChaos {
	networkChao := chao.Spec.NetworkChaos
	meshNetworkChao := &chaosv1alpha1.NetworkChaos{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      make(map[string]string),
			Annotations: make(map[string]string),
			Name:        chao.Name,
			Namespace:   chao.Namespace,
		},
	}
	meshNetworkChao.Spec = *deepCopyNetworkSpec(networkChao)

	return meshNetworkChao
}

func NewWorkflow(chao *v1alpha1.ShardingSphereChaos) *chaosv1alpha1.Workflow {
	workflow := chao.Spec.Workflow
	meshWorkflow := &chaosv1alpha1.Workflow{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      make(map[string]string),
			Annotations: make(map[string]string),
			Name:        chao.Name,
			Namespace:   chao.Namespace,
		},
		Spec: chaosv1alpha1.WorkflowSpec{
			Entry:     workflow.Entry,
			Templates: DeepCopyWorkTemplates(workflow.Templates),
		},
	}

	return meshWorkflow
}

func UpdateNetworkChaos(ssChaos *v1alpha1.ShardingSphereChaos, cur *chaosv1alpha1.NetworkChaos) *chaosv1alpha1.NetworkChaos {
	exp := &chaosv1alpha1.NetworkChaos{}
	exp.ObjectMeta = cur.ObjectMeta
	exp.ObjectMeta.ResourceVersion = ""
	exp.Labels = cur.Labels
	exp.Annotations = cur.Annotations
	exp.Spec = NewNetworkPodChaos(ssChaos).Spec
	return exp
}

func UpdatePodChaos(ssChaos *v1alpha1.ShardingSphereChaos, cur *chaosv1alpha1.PodChaos) *chaosv1alpha1.PodChaos {
	exp := &chaosv1alpha1.PodChaos{}
	exp.ObjectMeta = cur.ObjectMeta
	exp.ObjectMeta.ResourceVersion = ""
	exp.Labels = cur.Labels
	exp.Annotations = cur.Annotations
	exp.Spec = NewPodChaos(ssChaos).Spec
	return exp
}

func UpdateWorkflow(ssChaos *v1alpha1.ShardingSphereChaos, cur *chaosv1alpha1.Workflow) *chaosv1alpha1.Workflow {
	exp := &chaosv1alpha1.Workflow{}
	exp.ObjectMeta = cur.ObjectMeta
	exp.ObjectMeta.ResourceVersion = ""
	exp.Labels = cur.Labels
	exp.Annotations = cur.Annotations
	exp.Spec = NewWorkflow(ssChaos).Spec
	return exp
}

func DeepCopyWorkTemplates(templates []v1alpha1.WorkFlowTemplate) (meshTemplates []chaosv1alpha1.Template) {

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

func deepCopyPodChaoSpec(podChao *v1alpha1.PodChaosSpec) *chaosv1alpha1.PodChaosSpec {
	spec := chaosv1alpha1.PodChaosSpec{
		ContainerSelector: *deepCopyContainerSelector(&podChao.PodSelector),
		Action:            chaosv1alpha1.PodChaosAction(podChao.Action),
		Duration:          podChao.Duration,
		GracePeriod:       podChao.GracePeriod,
	}

	return &spec
}

func deepCopyNetworkSpec(networkChao *v1alpha1.NetworkChaosSpec) *chaosv1alpha1.NetworkChaosSpec {
	spec := chaosv1alpha1.NetworkChaosSpec{
		PodSelector:     *DeepCopyPodSelector(&networkChao.PodSelector),
		Action:          chaosv1alpha1.NetworkChaosAction(networkChao.Action),
		Device:          networkChao.Device,
		Duration:        networkChao.Duration,
		TcParameter:     *deepCopyTcParameter(&networkChao.TcParameter),
		Direction:       chaosv1alpha1.Direction(networkChao.Direction),
		Target:          DeepCopyPodSelector(networkChao.Target),
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

func deepCopyContainerSelector(selector *v1alpha1.PodSelector) *chaosv1alpha1.ContainerSelector {
	return &chaosv1alpha1.ContainerSelector{
		PodSelector:    *DeepCopyPodSelector(selector),
		ContainerNames: []string{},
	}
}

func DeepCopyPodSelector(selector *v1alpha1.PodSelector) *chaosv1alpha1.PodSelector {
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
