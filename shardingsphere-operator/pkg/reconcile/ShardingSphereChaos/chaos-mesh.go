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

package ShardingSphereChaos

import (
	"context"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaos"
	chaosv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	podSelectorMode      = "spec/mode"
	podSelectorValue     = "spec/value"
	device               = "spec/device"
	targetDevice         = "spec/targetDevice"
	targetPodSelectMode  = "spec/target/mode"
	targetPodSelectValue = "spec/target/value"
)

type chaosMeshHandler struct{}

func (c *chaosMeshHandler) CreatePodChaos(ctx context.Context, r client.Client, chao chaos.PodChaos) error {
	podChao := chao.(*chaosv1alpha1.PodChaos)
	if err := r.Create(ctx, podChao); err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	}

	return nil
}

func (c *chaosMeshHandler) CreateNetworkChaos(ctx context.Context, r client.Client, chao chaos.NetworkChaos) error {
	networkChao := chao.(*chaosv1alpha1.NetworkChaos)
	if err := r.Create(ctx, networkChao); err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	}

	return nil
}

func (c *chaosMeshHandler) NewPodChaos(ssChao *v1alpha1.ShardingSphereChaos) chaos.PodChaos {
	pcb := NewPodChaosBuilder(ssChao.GetObjectMeta(), ssChao.GetObjectKind().GroupVersionKind())
	pcb.SetName(ssChao.Name).SetNamespace(ssChao.Namespace).SetLabels(ssChao.Labels)

	chao := ssChao.Spec.PodChaos
	pcb.SetAction(chao.Action)

	psb := NewPodSelectorBuilder()

	psb.SetNamespaces(chao.Namespaces).
		SetExpressionSelectors(chao.ExpressionSelectors).
		SetNodes(chao.Nodes).
		SetNodeSelector(chao.NodeSelectors).
		SetAnnotationSelectors(chao.AnnotationSelectors).
		SetLabelSelector(chao.LabelSelectors).
		SetPods(chao.Pods)

	psb.SetSelectMode(ssChao.Annotations[podSelectorMode]).
		SetValue(ssChao.Annotations[podSelectorValue])
	containerSelector := &chaosv1alpha1.ContainerSelector{
		PodSelector: *psb.Build(),
	}

	if chao.Action == v1alpha1.PodFailureAction {
		pcb.SetDuration(chao.PodActionParam.PodFailure.Duration)
	}

	if chao.Action == v1alpha1.ContainerKillAction {
		containerSelector.ContainerNames = chao.PodActionParam.ContainerKill.ContainerNames
	}

	pcb.SetContainerSelector(*containerSelector)

	return pcb.Build()
}

func (c *chaosMeshHandler) NewNetworkPodChaos(ssChao *v1alpha1.ShardingSphereChaos) chaos.NetworkChaos {
	ncb := NewNetworkChaosBuilder(ssChao.GetObjectMeta(), ssChao.GetObjectKind().GroupVersionKind())

	ncb.SetName(ssChao.Name).SetNamespace(ssChao.Namespace).SetLabels(ssChao.Labels)
	chao := ssChao.Spec.NetworkChaos
	ncb.SetAction(chao.Action).SetDuration(*chao.Duration).SetDirection(string(chao.Direction))

	psb := NewPodSelectorBuilder()

	psb.SetNamespaces(chao.Source.Namespaces).
		SetExpressionSelectors(chao.Source.ExpressionSelectors).
		SetNodes(chao.Source.Nodes).
		SetNodeSelector(chao.Source.NodeSelectors).
		SetAnnotationSelectors(chao.Source.AnnotationSelectors).
		SetLabelSelector(chao.Source.LabelSelectors).
		SetPods(chao.Source.Pods)

	psb.SetSelectMode(ssChao.Annotations[podSelectorMode]).
		SetValue(ssChao.Annotations[podSelectorValue])

	ncb.SetPodSelector(psb.Build())

	tpsb := NewPodSelectorBuilder()

	tpsb.SetNamespaces(chao.Target.Namespaces).
		SetExpressionSelectors(chao.Target.ExpressionSelectors).
		SetNodes(chao.Target.Nodes).
		SetNodeSelector(chao.Target.NodeSelectors).
		SetAnnotationSelectors(chao.Target.AnnotationSelectors).
		SetLabelSelector(chao.Target.LabelSelectors).
		SetPods(chao.Target.Pods)

	tpsb.SetSelectMode(ssChao.Annotations[targetPodSelectMode]).
		SetValue(ssChao.Annotations[targetPodSelectValue])

	ncb.SetTarget(tpsb.Build())

	ncb.SetDevice(ssChao.Annotations[device]).
		SetTargetDevice(ssChao.Annotations[targetDevice])

	tcParams := &chaosv1alpha1.TcParameter{}

	if chao.Action == v1alpha1.DelayAction {
		tcParams.Delay = &chaosv1alpha1.DelaySpec{
			Latency:     chao.NetWorkParams.Delay.Latency,
			Correlation: chao.NetWorkParams.Delay.Correlation,
			Jitter:      chao.NetWorkParams.Delay.Jitter,
		}
	}

	if chao.Action == v1alpha1.CorruptAction {
		tcParams.Corrupt = &chaosv1alpha1.CorruptSpec{
			Corrupt:     chao.NetWorkParams.Corrupt.Corrupt,
			Correlation: chao.NetWorkParams.Corrupt.Correlation,
		}
	}

	if chao.Action == v1alpha1.DuplicateAction {
		tcParams.Duplicate = &chaosv1alpha1.DuplicateSpec{
			Duplicate:   chao.NetWorkParams.Duplicate.Duplicate,
			Correlation: chao.NetWorkParams.Duplicate.Correlation,
		}
	}

	if chao.Action == v1alpha1.LossAction {
		tcParams.Loss = &chaosv1alpha1.LossSpec{
			Loss:        chao.NetWorkParams.Loss.Loss,
			Correlation: chao.NetWorkParams.Loss.Correlation,
		}
	}

	ncb.SetTcParameter(*tcParams)

	return ncb.Build()
}

func (c *chaosMeshHandler) UpdateNetworkChaos(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos, r client.Client, cur chaos.NetworkChaos) error {
	Recur := cur.(*chaosv1alpha1.NetworkChaos)
	exp := &chaosv1alpha1.NetworkChaos{}
	exp.ObjectMeta = Recur.ObjectMeta
	exp.ObjectMeta.ResourceVersion = ""
	exp.Labels = Recur.Labels
	exp.Annotations = Recur.Annotations
	ReExp := (c.NewNetworkPodChaos(ssChaos)).(chaosv1alpha1.NetworkChaos)
	exp.Spec = ReExp.Spec

	return r.Update(ctx, exp)
}

func (c *chaosMeshHandler) UpdatePodChaos(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos, r client.Client, cur chaos.PodChaos) error {
	Recur := cur.(*chaosv1alpha1.PodChaos)
	exp := &chaosv1alpha1.PodChaos{}
	exp.ObjectMeta = Recur.ObjectMeta
	exp.ObjectMeta.ResourceVersion = ""
	exp.Labels = Recur.Labels
	exp.Annotations = Recur.Annotations
	ReExp := (c.NewPodChaos(ssChaos)).(chaosv1alpha1.PodChaos)
	exp.Spec = ReExp.Spec
	return r.Update(ctx, exp)
}

type PodChaosBuilder interface {
	SetNamespace(string) PodChaosBuilder
	SetName(string) PodChaosBuilder
	SetLabels(map[string]string) PodChaosBuilder
	SetAnnotations(map[string]string) PodChaosBuilder
	SetContainerSelector(chaosv1alpha1.ContainerSelector) PodChaosBuilder
	SetAction(v1alpha1.PodChaosAction) PodChaosBuilder
	SetDuration(string) PodChaosBuilder
	SetGracePeriod(int64) PodChaosBuilder
	Build() *chaosv1alpha1.PodChaos
}

func NewPodChaosBuilder(meta metav1.Object, gvk schema.GroupVersionKind) PodChaosBuilder {
	return &podChaosBuilder{
		podChaos: DefaultPodChaos(meta, gvk),
	}
}

type podChaosBuilder struct {
	podChaos *chaosv1alpha1.PodChaos
}

func (p *podChaosBuilder) SetNamespace(namespace string) PodChaosBuilder {
	p.podChaos.Namespace = namespace
	return p
}

func (p *podChaosBuilder) SetName(name string) PodChaosBuilder {
	p.podChaos.Name = name
	return p
}

func (p *podChaosBuilder) SetLabels(labels map[string]string) PodChaosBuilder {
	p.podChaos.Labels = labels
	return p
}

func (p *podChaosBuilder) SetAnnotations(annotations map[string]string) PodChaosBuilder {
	p.podChaos.Annotations = annotations
	return p
}

func (p *podChaosBuilder) SetContainerSelector(selector chaosv1alpha1.ContainerSelector) PodChaosBuilder {
	p.podChaos.Spec.ContainerSelector = selector
	return p
}

func (p *podChaosBuilder) SetAction(action v1alpha1.PodChaosAction) PodChaosBuilder {
	if action == v1alpha1.PodFailureAction {
		p.podChaos.Spec.Action = chaosv1alpha1.PodFailureAction
	}

	if action == v1alpha1.ContainerKillAction {
		p.podChaos.Spec.Action = chaosv1alpha1.ContainerKillAction
	}
	return p
}

func (p *podChaosBuilder) SetDuration(duration string) PodChaosBuilder {
	if duration == "" {
		//todo: change to default
		ret := "1m"
		p.podChaos.Spec.Duration = &ret
	} else {
		p.podChaos.Spec.Duration = &duration
	}
	return p
}

func (p *podChaosBuilder) SetGracePeriod(gracePeriod int64) PodChaosBuilder {
	p.podChaos.Spec.GracePeriod = gracePeriod
	return p
}

func (p *podChaosBuilder) Build() *chaosv1alpha1.PodChaos {
	return p.podChaos
}

type NetworkChaosBuilder interface {
	SetNamespace(string) NetworkChaosBuilder
	SetName(string) NetworkChaosBuilder
	SetLabels(map[string]string) NetworkChaosBuilder
	SetAnnotations(map[string]string) NetworkChaosBuilder
	SetPodSelector(*chaosv1alpha1.PodSelector) NetworkChaosBuilder
	SetAction(v1alpha1.NetworkChaosAction) NetworkChaosBuilder
	SetDevice(string) NetworkChaosBuilder
	SetDuration(string) NetworkChaosBuilder
	SetDirection(string) NetworkChaosBuilder
	SetTarget(*chaosv1alpha1.PodSelector) NetworkChaosBuilder
	SetTargetDevice(string) NetworkChaosBuilder
	SetTcParameter(chaosv1alpha1.TcParameter) NetworkChaosBuilder
	Build() *chaosv1alpha1.NetworkChaos
}

type netWorkChaosBuilder struct {
	netWorkChaos *chaosv1alpha1.NetworkChaos
}

func (n *netWorkChaosBuilder) SetNamespace(namespace string) NetworkChaosBuilder {
	n.netWorkChaos.Namespace = namespace
	return n
}

func (n *netWorkChaosBuilder) SetName(name string) NetworkChaosBuilder {
	n.netWorkChaos.Name = name
	return n
}

func (n *netWorkChaosBuilder) SetLabels(labels map[string]string) NetworkChaosBuilder {
	n.netWorkChaos.Labels = labels
	return n
}

func (n *netWorkChaosBuilder) SetAnnotations(annotations map[string]string) NetworkChaosBuilder {
	n.netWorkChaos.Annotations = annotations
	return n
}

func (n *netWorkChaosBuilder) SetPodSelector(selector *chaosv1alpha1.PodSelector) NetworkChaosBuilder {
	n.netWorkChaos.Spec.PodSelector = *selector
	return n
}

func (n *netWorkChaosBuilder) SetAction(action v1alpha1.NetworkChaosAction) NetworkChaosBuilder {
	if action == v1alpha1.CorruptAction {
		n.netWorkChaos.Spec.Action = chaosv1alpha1.CorruptAction
	}

	if action == v1alpha1.PartitionAction {
		n.netWorkChaos.Spec.Action = chaosv1alpha1.PartitionAction
	}

	if action == v1alpha1.LossAction {
		n.netWorkChaos.Spec.Action = chaosv1alpha1.LossAction
	}

	if action == v1alpha1.DuplicateAction {
		n.netWorkChaos.Spec.Action = chaosv1alpha1.DuplicateAction
	}

	return n
}

func (n *netWorkChaosBuilder) SetDevice(device string) NetworkChaosBuilder {
	n.netWorkChaos.Spec.Device = device
	return n
}

func (n *netWorkChaosBuilder) SetDuration(duration string) NetworkChaosBuilder {
	n.netWorkChaos.Spec.Duration = &duration
	return n
}

func (n *netWorkChaosBuilder) SetDirection(direction string) NetworkChaosBuilder {
	n.netWorkChaos.Spec.Direction = chaosv1alpha1.Direction(direction)
	return n
}

func (n *netWorkChaosBuilder) SetTarget(selector *chaosv1alpha1.PodSelector) NetworkChaosBuilder {
	n.netWorkChaos.Spec.Target = selector
	return n
}

func (n *netWorkChaosBuilder) SetTargetDevice(targetDevice string) NetworkChaosBuilder {
	n.netWorkChaos.Spec.TargetDevice = targetDevice
	return n
}

func (n *netWorkChaosBuilder) SetTcParameter(parameter chaosv1alpha1.TcParameter) NetworkChaosBuilder {
	n.netWorkChaos.Spec.TcParameter = parameter
	return n
}

func (n *netWorkChaosBuilder) Build() *chaosv1alpha1.NetworkChaos {
	return n.netWorkChaos
}

func NewNetworkChaosBuilder(meta metav1.Object, gvk schema.GroupVersionKind) NetworkChaosBuilder {
	return &netWorkChaosBuilder{
		netWorkChaos: DefaultNetworkChaos(meta, gvk),
	}
}

type PodSelectorBuilder interface {
	SetNamespaces([]string) PodSelectorBuilder
	SetSelectMode(string) PodSelectorBuilder
	SetValue(string) PodSelectorBuilder
	SetNodes([]string) PodSelectorBuilder
	SetPods(map[string][]string) PodSelectorBuilder
	SetNodeSelector(map[string]string) PodSelectorBuilder
	SetPodPhaseSelectors([]string) PodSelectorBuilder
	SetFieldSelectors(map[string]string) PodSelectorBuilder
	SetLabelSelector(map[string]string) PodSelectorBuilder
	SetExpressionSelectors([]metav1.LabelSelectorRequirement) PodSelectorBuilder
	SetAnnotationSelectors(map[string]string) PodSelectorBuilder
	Build() *chaosv1alpha1.PodSelector
}

func NewPodSelectorBuilder() PodSelectorBuilder {
	return &podSelectorBuilder{
		podSelector: &chaosv1alpha1.PodSelector{},
	}
}

type podSelectorBuilder struct {
	podSelector *chaosv1alpha1.PodSelector
}

func (p *podSelectorBuilder) SetNamespaces(namespaces []string) PodSelectorBuilder {
	p.podSelector.Selector.Namespaces = namespaces
	return p
}

func (p *podSelectorBuilder) SetSelectMode(mode string) PodSelectorBuilder {
	p.podSelector.Mode = chaosv1alpha1.SelectorMode(mode)
	return p
}

func (p *podSelectorBuilder) SetValue(value string) PodSelectorBuilder {
	p.podSelector.Value = value
	return p
}

func (p *podSelectorBuilder) SetNodes(nodes []string) PodSelectorBuilder {
	p.podSelector.Selector.Nodes = nodes
	return p
}

func (p *podSelectorBuilder) SetPods(pods map[string][]string) PodSelectorBuilder {
	p.podSelector.Selector.Pods = pods
	return p
}

func (p *podSelectorBuilder) SetNodeSelector(nodeSelector map[string]string) PodSelectorBuilder {
	p.podSelector.Selector.NodeSelectors = nodeSelector
	return p
}

func (p *podSelectorBuilder) SetPodPhaseSelectors(podPhaseSelectors []string) PodSelectorBuilder {
	p.podSelector.Selector.PodPhaseSelectors = podPhaseSelectors
	return p
}

func (p *podSelectorBuilder) SetFieldSelectors(fieldSelectors map[string]string) PodSelectorBuilder {
	p.podSelector.Selector.FieldSelectors = fieldSelectors
	return p
}

func (p *podSelectorBuilder) SetLabelSelector(labelSelectors map[string]string) PodSelectorBuilder {
	p.podSelector.Selector.LabelSelectors = labelSelectors
	return p
}

func (p *podSelectorBuilder) SetExpressionSelectors(requirements []metav1.LabelSelectorRequirement) PodSelectorBuilder {
	p.podSelector.Selector.ExpressionSelectors = requirements
	return p
}

func (p *podSelectorBuilder) SetAnnotationSelectors(annotationSelectors map[string]string) PodSelectorBuilder {
	p.podSelector.Selector.AnnotationSelectors = annotationSelectors
	return p
}

func (p *podSelectorBuilder) Build() *chaosv1alpha1.PodSelector {
	return p.podSelector
}

func DefaultPodChaos(meta metav1.Object, gvk schema.GroupVersionKind) *chaosv1alpha1.PodChaos {
	return &chaosv1alpha1.PodChaos{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "shardingsphere-proxy",
			Namespace: "default",
			Labels:    map[string]string{},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(meta, gvk),
			},
		},
		Spec: chaosv1alpha1.PodChaosSpec{
			Action: chaosv1alpha1.ContainerKillAction,
		},
	}
}

func DefaultNetworkChaos(meta metav1.Object, gvk schema.GroupVersionKind) *chaosv1alpha1.NetworkChaos {
	return &chaosv1alpha1.NetworkChaos{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "shardingsphere-proxy",
			Namespace: "default",
			Labels:    map[string]string{},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(meta, gvk),
			},
		},
		Spec: chaosv1alpha1.NetworkChaosSpec{
			Action:    chaosv1alpha1.PartitionAction,
			Direction: "to",
		},
	}
}
