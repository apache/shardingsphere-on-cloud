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

package chaosmesh

import (
	"context"
	"errors"
	"strconv"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"

	chaosmeshv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	AnnoPodSelectorMode      = "selector.chaos-mesh.org/mode"
	AnnoPodSelectorValue     = "selector.chaos-mesh.org/value"
	AnnoTargetPodSelectMode  = "target-selector.chaos-mesh.org/mode"
	AnnoTargetPodSelectValue = "target-selector.chaos-mesh.org/value"

	AnnoPodAction   = "podchaos.chaos-mesh.org/action"
	AnnoGracePeriod = "podchaos.chaos-mesh.org/gracePeriod"

	AnnoNetworkAction     = "networkchaos.chaos-mesh.org/action"
	AnnoDevice            = "networkchaos.chaos-mesh.org/device"
	AnnoTargetDevice      = "networkchaos.chaos-mesh.org/targetDevice"
	AnnoBandwidthRate     = "networkchaos.chaos-mesh.org/bandwidth:rate"
	AnnoBandwidthLimit    = "networkchaos.chaos-mesh.org/bandwidth:limit"
	AnnoBandwidthBuffer   = "networkchaos.chaos-mesh.org/bandwidth:buffer"
	AnnoBandwidthPeakrate = "networkchaos.chaos-mesh.org/bandwidth:peakrate"
	AnnoBandwidthMinBurst = "networkchaos.chaos-mesh.org/bandwidth:minburst"
)

var (
	ErrConvert     = errors.New("can not convert chaos interface to specify struct")
	ErrNotChanged  = errors.New("object not changed")
	ErrChangedSpec = errors.New("change spec")
)

type GenericChaos interface{}

func ConvertChaosStatus(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos, chaos GenericChaos) v1alpha1.ChaosCondition {
	var status chaosmeshv1alpha1.ChaosStatus
	if ssChaos.Spec.EmbedChaos.PodChaos != nil {
		if podChao, ok := chaos.(*chaosmeshv1alpha1.PodChaos); ok && podChao != nil {
			status = *podChao.GetStatus()
		} else {
			return v1alpha1.Unknown
		}
	} else if ssChaos.Spec.EmbedChaos.NetworkChaos != nil {
		if networkChaos, ok := chaos.(*chaosmeshv1alpha1.NetworkChaos); ok && networkChaos != nil {
			status = *networkChaos.GetStatus()
		}
		return v1alpha1.Unknown
	}
	var conditions = map[chaosmeshv1alpha1.ChaosConditionType]bool{}
	for i := range status.Conditions {
		conditions[status.Conditions[i].Type] = status.Conditions[i].Status == corev1.ConditionTrue
	}

	return judgeCondition(conditions, status.Experiment.DesiredPhase)
}

func judgeCondition(condition map[chaosmeshv1alpha1.ChaosConditionType]bool, phase chaosmeshv1alpha1.DesiredPhase) v1alpha1.ChaosCondition {

	if condition[chaosmeshv1alpha1.ConditionPaused] {
		if !condition[chaosmeshv1alpha1.ConditionSelected] {
			return v1alpha1.NoTarget
		}
		return v1alpha1.Paused
	}

	if condition[chaosmeshv1alpha1.ConditionSelected] {
		if condition[chaosmeshv1alpha1.ConditionAllRecovered] && phase == chaosmeshv1alpha1.StoppedPhase {
			return v1alpha1.AllRecovered
		}

		if condition[chaosmeshv1alpha1.ConditionAllInjected] && phase == chaosmeshv1alpha1.RunningPhase {
			return v1alpha1.AllInjected
		}
	}

	return v1alpha1.Unknown
}

func NewPodChaos(ssChao *v1alpha1.ShardingSphereChaos) (PodChaos, error) {
	pcb := NewPodChaosBuilder()
	pcb.SetName(ssChao.Name).SetNamespace(ssChao.Namespace).SetLabels(ssChao.Labels)

	chao := ssChao.Spec.PodChaos
	if act, ok := ssChao.Annotations[AnnoPodAction]; ok {
		pcb.SetAction(act)
		if gp, ok := ssChao.Annotations[AnnoGracePeriod]; chaosmeshv1alpha1.PodChaosAction(act) == chaosmeshv1alpha1.PodKillAction && ok {
			gpInt, err := strconv.ParseInt(gp, 10, 64)
			if err != nil {
				return nil, err
			}
			pcb.SetGracePeriod(gpInt)
		}
	} else {
		pcb.SetAction(string(chao.Action))
	}

	psb := NewPodSelectorBuilder()

	psb.SetNamespaces(chao.Namespaces).
		SetExpressionSelectors(chao.ExpressionSelectors).
		SetNodes(chao.Nodes).
		SetNodeSelector(chao.NodeSelectors).
		SetAnnotationSelectors(chao.AnnotationSelectors).
		SetLabelSelector(chao.LabelSelectors).
		SetPods(chao.Pods)

	psb.SetSelectMode(ssChao.Annotations[AnnoTargetPodSelectMode]).
		SetValue(ssChao.Annotations[AnnoTargetPodSelectValue])

	containerSelector := &chaosmeshv1alpha1.ContainerSelector{
		PodSelector: *psb.Build(),
	}

	if chao.Action == v1alpha1.PodFailure {
		pcb.SetDuration(chao.Params.PodFailure.Duration)
	}

	if chao.Action == v1alpha1.ContainerKill {
		containerSelector.ContainerNames = chao.Params.ContainerKill.ContainerNames
	}

	pcb.SetContainerSelector(containerSelector)
	podChao := pcb.Build()

	return podChao, nil
}

func getAnnotation(anno map[string]string, k string) string {
	if v, ok := anno[k]; ok {
		return v
	}
	return ""
}

func NewNetworkChaos(ssChao *v1alpha1.ShardingSphereChaos) (NetworkChaos, error) {
	ncb := NewNetworkChaosBuilder()
	ncb.SetName(ssChao.Name).SetNamespace(ssChao.Namespace).SetLabels(ssChao.Labels)

	tcParams := &chaosmeshv1alpha1.TcParameter{}

	chao := ssChao.Spec.NetworkChaos
	if act, ok := ssChao.Annotations[AnnoNetworkAction]; ok {
		ncb.SetAction(act)

		if chaosmeshv1alpha1.NetworkChaosAction(act) == chaosmeshv1alpha1.BandwidthAction {
			bwab := NewBandWidthActionBuilder()
			bwab.SetRate(getAnnotation(ssChao.Annotations, AnnoBandwidthRate))
			bwab.SetLimit(getAnnotation(ssChao.Annotations, AnnoBandwidthLimit))
			bwab.SetBuffer(getAnnotation(ssChao.Annotations, AnnoBandwidthBuffer))
			bwab.SetPeakRate(getAnnotation(ssChao.Annotations, AnnoBandwidthPeakrate))
			bwab.SetMinBurst(getAnnotation(ssChao.Annotations, AnnoBandwidthMinBurst))
			tcParams.Bandwidth = bwab.Build()
		}
	} else {
		ncb.SetAction(string(chao.Action))
	}

	ncb.SetDuration(chao.Duration).SetDirection(string(chao.Direction))

	psb := NewPodSelectorBuilder()
	psb.SetNamespaces(chao.Source.Namespaces).
		SetExpressionSelectors(chao.Source.ExpressionSelectors).
		SetNodes(chao.Source.Nodes).
		SetNodeSelector(chao.Source.NodeSelectors).
		SetAnnotationSelectors(chao.Source.AnnotationSelectors).
		SetLabelSelector(chao.Source.LabelSelectors).
		SetPods(chao.Source.Pods)

	psb.SetSelectMode(ssChao.Annotations[AnnoPodSelectorMode]).
		SetValue(ssChao.Annotations[AnnoPodSelectorValue])

	ncb.SetPodSelector(psb.Build())

	tpsb := NewPodSelectorBuilder()
	tpsb.SetNamespaces(chao.Target.Namespaces).
		SetExpressionSelectors(chao.Target.ExpressionSelectors).
		SetNodes(chao.Target.Nodes).
		SetNodeSelector(chao.Target.NodeSelectors).
		SetAnnotationSelectors(chao.Target.AnnotationSelectors).
		SetLabelSelector(chao.Target.LabelSelectors).
		SetPods(chao.Target.Pods)

	tpsb.SetSelectMode(ssChao.Annotations[AnnoTargetPodSelectMode]).
		SetValue(ssChao.Annotations[AnnoTargetPodSelectValue])

	ncb.SetTarget(tpsb.Build())
	ncb.SetDevice(ssChao.Annotations[AnnoDevice]).
		SetTargetDevice(ssChao.Annotations[AnnoTargetDevice])

	switch chao.Action {
	case v1alpha1.Delay:
		tcParams.Delay = &chaosmeshv1alpha1.DelaySpec{
			Latency: chao.Params.Delay.Latency,
			Jitter:  chao.Params.Delay.Jitter,
		}
	case v1alpha1.Corruption:
		tcParams.Corrupt = &chaosmeshv1alpha1.CorruptSpec{
			Corrupt: chao.Params.Corruption.Corruption,
		}
	case v1alpha1.Duplication:
		tcParams.Duplicate = &chaosmeshv1alpha1.DuplicateSpec{
			Duplicate: chao.Params.Duplication.Duplication,
		}
	case v1alpha1.Loss:
		tcParams.Loss = &chaosmeshv1alpha1.LossSpec{
			Loss: chao.Params.Loss.Loss,
		}
	}

	ncb.SetTcParameter(*tcParams)
	networkChao := ncb.Build()

	return networkChao, nil
}

type PodChaosBuilder interface {
	SetNamespace(string) PodChaosBuilder
	SetName(string) PodChaosBuilder
	SetLabels(map[string]string) PodChaosBuilder
	SetAnnotations(map[string]string) PodChaosBuilder
	SetContainerSelector(*chaosmeshv1alpha1.ContainerSelector) PodChaosBuilder
	SetAction(string) PodChaosBuilder
	SetDuration(*string) PodChaosBuilder
	SetGracePeriod(int64) PodChaosBuilder
	Build() *chaosmeshv1alpha1.PodChaos
}

func NewPodChaosBuilder() PodChaosBuilder {
	return &podChaosBuilder{
		podChaos: DefaultPodChaos(),
	}
}

type podChaosBuilder struct {
	podChaos *chaosmeshv1alpha1.PodChaos
}

type BandWidthActionBuilder interface {
	SetRate(string) BandWidthActionBuilder
	SetLimit(string) BandWidthActionBuilder
	SetBuffer(string) BandWidthActionBuilder
	SetPeakRate(string) BandWidthActionBuilder
	SetMinBurst(string) BandWidthActionBuilder
	Build() *chaosmeshv1alpha1.BandwidthSpec
}

func NewBandWidthActionBuilder() BandWidthActionBuilder {
	return &bandWidthActionBuilder{
		bandwidth: &chaosmeshv1alpha1.BandwidthSpec{},
	}
}

type bandWidthActionBuilder struct {
	bandwidth *chaosmeshv1alpha1.BandwidthSpec
}

func (b *bandWidthActionBuilder) SetRate(s string) BandWidthActionBuilder {
	b.bandwidth.Rate = s
	return b
}

func (b *bandWidthActionBuilder) SetLimit(s string) BandWidthActionBuilder {
	lim, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		b.bandwidth.Limit = uint32(lim)
	}
	return b
}

func (b *bandWidthActionBuilder) SetBuffer(s string) BandWidthActionBuilder {
	buf, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		b.bandwidth.Buffer = uint32(buf)
	}
	return b
}

func (b *bandWidthActionBuilder) SetPeakRate(s string) BandWidthActionBuilder {
	pr, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		ret := uint64(pr)
		b.bandwidth.Peakrate = &ret
	}
	return b
}

func (b *bandWidthActionBuilder) SetMinBurst(s string) BandWidthActionBuilder {
	burst, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		ret := uint32(burst)
		b.bandwidth.Minburst = &ret
	}
	return b
}

func (b *bandWidthActionBuilder) Build() *chaosmeshv1alpha1.BandwidthSpec {
	return b.bandwidth
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

func (p *podChaosBuilder) SetContainerSelector(selector *chaosmeshv1alpha1.ContainerSelector) PodChaosBuilder {
	p.podChaos.Spec.ContainerSelector = *selector
	return p
}

func (p *podChaosBuilder) SetAction(action string) PodChaosBuilder {
	if v1alpha1.PodChaosAction(action) == v1alpha1.PodFailure {
		p.podChaos.Spec.Action = chaosmeshv1alpha1.PodFailureAction
	}

	if v1alpha1.PodChaosAction(action) == v1alpha1.ContainerKill {
		p.podChaos.Spec.Action = chaosmeshv1alpha1.ContainerKillAction
	}

	if chaosmeshv1alpha1.PodChaosAction(action) == chaosmeshv1alpha1.PodKillAction {
		p.podChaos.Spec.Action = chaosmeshv1alpha1.PodKillAction
	}
	return p
}

func (p *podChaosBuilder) SetDuration(duration *string) PodChaosBuilder {
	p.podChaos.Spec.Duration = duration
	return p
}

func (p *podChaosBuilder) SetGracePeriod(gracePeriod int64) PodChaosBuilder {
	p.podChaos.Spec.GracePeriod = gracePeriod
	return p
}

func (p *podChaosBuilder) Build() *chaosmeshv1alpha1.PodChaos {
	return p.podChaos
}

type NetworkChaosBuilder interface {
	SetNamespace(string) NetworkChaosBuilder
	SetName(string) NetworkChaosBuilder
	SetLabels(map[string]string) NetworkChaosBuilder
	SetAnnotations(map[string]string) NetworkChaosBuilder
	SetPodSelector(*chaosmeshv1alpha1.PodSelector) NetworkChaosBuilder
	SetAction(string) NetworkChaosBuilder
	SetDevice(string) NetworkChaosBuilder
	SetDuration(*string) NetworkChaosBuilder
	SetDirection(string) NetworkChaosBuilder
	SetTarget(*chaosmeshv1alpha1.PodSelector) NetworkChaosBuilder
	SetTargetDevice(string) NetworkChaosBuilder
	SetTcParameter(chaosmeshv1alpha1.TcParameter) NetworkChaosBuilder
	Build() *chaosmeshv1alpha1.NetworkChaos
}

type netWorkChaosBuilder struct {
	netWorkChaos *chaosmeshv1alpha1.NetworkChaos
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

func (n *netWorkChaosBuilder) SetPodSelector(selector *chaosmeshv1alpha1.PodSelector) NetworkChaosBuilder {
	n.netWorkChaos.Spec.PodSelector = *selector
	return n
}

func (n *netWorkChaosBuilder) SetAction(action string) NetworkChaosBuilder {
	if chaosmeshv1alpha1.NetworkChaosAction(action) == chaosmeshv1alpha1.BandwidthAction {
		n.netWorkChaos.Spec.Action = chaosmeshv1alpha1.BandwidthAction
	}

	if v1alpha1.NetworkChaosAction(action) == v1alpha1.Corruption {
		n.netWorkChaos.Spec.Action = chaosmeshv1alpha1.CorruptAction
	}

	if v1alpha1.NetworkChaosAction(action) == v1alpha1.Partition {
		n.netWorkChaos.Spec.Action = chaosmeshv1alpha1.PartitionAction
	}

	if v1alpha1.NetworkChaosAction(action) == v1alpha1.Loss {
		n.netWorkChaos.Spec.Action = chaosmeshv1alpha1.LossAction
	}

	if v1alpha1.NetworkChaosAction(action) == v1alpha1.Duplication {
		n.netWorkChaos.Spec.Action = chaosmeshv1alpha1.DuplicateAction
	}

	return n
}

func (n *netWorkChaosBuilder) SetDevice(device string) NetworkChaosBuilder {
	n.netWorkChaos.Spec.Device = device
	return n
}

func (n *netWorkChaosBuilder) SetDuration(duration *string) NetworkChaosBuilder {
	n.netWorkChaos.Spec.Duration = duration
	return n
}

func (n *netWorkChaosBuilder) SetDirection(direction string) NetworkChaosBuilder {
	n.netWorkChaos.Spec.Direction = chaosmeshv1alpha1.Direction(direction)
	return n
}

func (n *netWorkChaosBuilder) SetTarget(selector *chaosmeshv1alpha1.PodSelector) NetworkChaosBuilder {
	n.netWorkChaos.Spec.Target = selector
	return n
}

func (n *netWorkChaosBuilder) SetTargetDevice(targetDevice string) NetworkChaosBuilder {
	n.netWorkChaos.Spec.TargetDevice = targetDevice
	return n
}

func (n *netWorkChaosBuilder) SetTcParameter(parameter chaosmeshv1alpha1.TcParameter) NetworkChaosBuilder {
	n.netWorkChaos.Spec.TcParameter = parameter
	return n
}

func (n *netWorkChaosBuilder) Build() *chaosmeshv1alpha1.NetworkChaos {
	return n.netWorkChaos
}

func NewNetworkChaosBuilder() NetworkChaosBuilder {
	return &netWorkChaosBuilder{
		netWorkChaos: DefaultNetworkChaos(),
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
	Build() *chaosmeshv1alpha1.PodSelector
}

func NewPodSelectorBuilder() PodSelectorBuilder {
	return &podSelectorBuilder{
		podSelector: &chaosmeshv1alpha1.PodSelector{},
	}
}

type podSelectorBuilder struct {
	podSelector *chaosmeshv1alpha1.PodSelector
}

func (p *podSelectorBuilder) SetNamespaces(namespaces []string) PodSelectorBuilder {
	p.podSelector.Selector.Namespaces = namespaces
	return p
}

func (p *podSelectorBuilder) SetSelectMode(mode string) PodSelectorBuilder {
	p.podSelector.Mode = chaosmeshv1alpha1.SelectorMode(mode)
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

func (p *podSelectorBuilder) Build() *chaosmeshv1alpha1.PodSelector {
	return p.podSelector
}

func DefaultPodChaos() *chaosmeshv1alpha1.PodChaos {
	d := "1m"
	return &chaosmeshv1alpha1.PodChaos{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "shardingsphere-proxy",
			Namespace: "default",
			Labels:    map[string]string{},
		},
		Spec: chaosmeshv1alpha1.PodChaosSpec{
			Action:   chaosmeshv1alpha1.ContainerKillAction,
			Duration: &d,
		},
	}
}

func DefaultNetworkChaos() *chaosmeshv1alpha1.NetworkChaos {
	return &chaosmeshv1alpha1.NetworkChaos{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "shardingsphere-proxy",
			Namespace: "default",
			Labels:    map[string]string{},
		},
		Spec: chaosmeshv1alpha1.NetworkChaosSpec{
			Action:    chaosmeshv1alpha1.PartitionAction,
			Direction: "to",
		},
	}
}
