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

package shardingspherechaos

import (
	"context"
	"errors"
	"reflect"
	"strconv"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	chaosv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	podSelectorMode      = "selector.chaos-mesh.org/mode"
	podSelectorValue     = "selector.chaos-mesh.org/value"
	targetPodSelectMode  = "target-selector.chaos-mesh.org/mode"
	targetPodSelectValue = "target-selector.chaos-mesh.org/value"
	podAction            = "podchaos.chaos-mesh.org/action"
	gracePeriod          = "podchaos.chaos-mesh.org/gracePeriod"
	device               = "networkchaos.chaos-mesh.org/device"
	targetDevice         = "networkchaos.chaos-mesh.org/targetDevice"
	networkAction        = "networkchaos.chaos-mesh.org/action"
	rate                 = "networkchaos.chaos-mesh.org/bandwidth:rate"
	limit                = "networkchaos.chaos-mesh.org/bandwidth:limit"
	buffer               = "networkchaos.chaos-mesh.org/bandwidth:buffer"
	peakrate             = "networkchaos.chaos-mesh.org/bandwidth:peakrate"
	minburst             = "networkchaos.chaos-mesh.org/bandwidth:minburst"
)

var (
	ErrConvert     = errors.New("can not convert chaos interface to specify struct")
	ErrNotChanged  = errors.New("object not changed")
	ErrChangedSpec = errors.New("change spec")
)

type chaosMeshHandler struct {
	r client.Client
}

func NewChaosMeshHandler(r client.Client) ChaosHandler {
	return &chaosMeshHandler{r}
}

func (c *chaosMeshHandler) ConvertChaosStatus(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos, chaos GenericChaos) v1alpha1.ChaosCondition {
	var status chaosv1alpha1.ChaosStatus
	if ssChaos.Spec.EmbedChaos.PodChaos != nil {
		if podChao, ok := chaos.(*chaosv1alpha1.PodChaos); ok && podChao != nil {
			status = *podChao.GetStatus()
		} else {
			return v1alpha1.Unknown
		}
	} else if ssChaos.Spec.EmbedChaos.NetworkChaos != nil {
		if networkChaos, ok := chaos.(*chaosv1alpha1.NetworkChaos); ok && networkChaos != nil {
			status = *networkChaos.GetStatus()
		}
		return v1alpha1.Unknown
	}
	var conditions = map[chaosv1alpha1.ChaosConditionType]bool{}
	for i := range status.Conditions {
		conditions[status.Conditions[i].Type] = status.Conditions[i].Status == corev1.ConditionTrue
	}

	return judgeCondition(conditions, status.Experiment.DesiredPhase)
}

func judgeCondition(condition map[chaosv1alpha1.ChaosConditionType]bool, phase chaosv1alpha1.DesiredPhase) v1alpha1.ChaosCondition {

	if condition[chaosv1alpha1.ConditionPaused] {
		if !condition[chaosv1alpha1.ConditionSelected] {
			return v1alpha1.NoTarget
		}
		return v1alpha1.Paused
	}

	if condition[chaosv1alpha1.ConditionSelected] {
		if condition[chaosv1alpha1.ConditionAllRecovered] && phase == chaosv1alpha1.StoppedPhase {
			return v1alpha1.AllRecovered
		}

		if condition[chaosv1alpha1.ConditionAllInjected] && phase == chaosv1alpha1.RunningPhase {
			return v1alpha1.AllInjected
		}
	}

	return v1alpha1.Unknown
}

func (c *chaosMeshHandler) CreatePodChaos(ctx context.Context, chao PodChaos) error {
	podChao, ok := chao.(*chaosv1alpha1.PodChaos)
	if !ok {
		return ErrConvert
	}
	if err := c.r.Create(ctx, podChao); err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	}

	return nil
}

func (c *chaosMeshHandler) CreateNetworkChaos(ctx context.Context, chao NetworkChaos) error {
	networkChao, ok := chao.(*chaosv1alpha1.NetworkChaos)
	if !ok {
		return ErrConvert
	}
	if err := c.r.Create(ctx, networkChao); err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	}

	return nil
}

func (c *chaosMeshHandler) NewPodChaos(ssChao *v1alpha1.ShardingSphereChaos) (PodChaos, error) {
	pcb := NewPodChaosBuilder()
	pcb.SetName(ssChao.Name).SetNamespace(ssChao.Namespace).SetLabels(ssChao.Labels)

	chao := ssChao.Spec.PodChaos
	if act, ok := ssChao.Annotations[podAction]; ok {
		pcb.SetAction(act)
		if gp, ok := ssChao.Annotations[gracePeriod]; chaosv1alpha1.PodChaosAction(act) == chaosv1alpha1.PodKillAction && ok {
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

	pcb.SetContainerSelector(containerSelector)
	podChao := pcb.Build()

	if err := ctrl.SetControllerReference(ssChao, podChao, c.r.Scheme()); err != nil {
		return nil, err
	}
	return podChao, nil
}

func (c *chaosMeshHandler) NewNetworkPodChaos(ssChao *v1alpha1.ShardingSphereChaos) (NetworkChaos, error) {
	ncb := NewNetworkChaosBuilder()
	ncb.SetName(ssChao.Name).SetNamespace(ssChao.Namespace).SetLabels(ssChao.Labels)
	chao := ssChao.Spec.NetworkChaos
	act, ok := ssChao.Annotations[networkAction]
	if ok {
		ncb.SetAction(act)
	} else {
		ncb.SetAction(string(chao.Action))
	}

	ncb.SetDuration(*chao.Duration).SetDirection(string(chao.Direction))

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
			Latency:     chao.Network.Delay.Latency,
			Correlation: chao.Network.Delay.Correlation,
			Jitter:      chao.Network.Delay.Jitter,
		}
	}

	if chao.Action == v1alpha1.CorruptAction {
		tcParams.Corrupt = &chaosv1alpha1.CorruptSpec{
			Corrupt:     chao.Network.Corrupt.Corrupt,
			Correlation: chao.Network.Corrupt.Correlation,
		}
	}

	if chao.Action == v1alpha1.DuplicateAction {
		tcParams.Duplicate = &chaosv1alpha1.DuplicateSpec{
			Duplicate:   chao.Network.Duplicate.Duplicate,
			Correlation: chao.Network.Duplicate.Correlation,
		}
	}

	if chao.Action == v1alpha1.LossAction {
		tcParams.Loss = &chaosv1alpha1.LossSpec{
			Loss:        chao.Network.Loss.Loss,
			Correlation: chao.Network.Loss.Correlation,
		}
	}

	if chaosv1alpha1.NetworkChaosAction(act) == chaosv1alpha1.BandwidthAction {
		bwab := NewBandWidthActionBuilder()
		if ind1, ok := ssChao.Annotations[rate]; ok {
			bwab.SetRate(ind1)
		}

		if ind2, ok := ssChao.Annotations[limit]; ok {
			bwab.SetLimit(ind2)
		}
		if ind3, ok := ssChao.Annotations[buffer]; ok {
			bwab.SetBuffer(ind3)
		}
		if ind4, ok := ssChao.Annotations[peakrate]; ok {
			bwab.SetPeakRate(ind4)
		}
		if ind5, ok := ssChao.Annotations[minburst]; ok {
			bwab.SetMinBurst(ind5)
		}
		tcParams.Bandwidth = bwab.Build()
	}
	ncb.SetTcParameter(*tcParams)

	networkChao := ncb.Build()
	if err := ctrl.SetControllerReference(ssChao, networkChao, c.r.Scheme()); err != nil {
		return nil, err
	}
	return networkChao, nil
}

func (c *chaosMeshHandler) UpdateNetworkChaos(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos, cur NetworkChaos) error {
	networkChao, err := c.NewNetworkPodChaos(ssChaos)
	if err != nil {
		return err
	}

	reExp, ok := networkChao.(*chaosv1alpha1.NetworkChaos)
	if !ok {
		return ErrConvert
	}
	reCur, ok := cur.(*chaosv1alpha1.NetworkChaos)
	if !ok {
		return ErrConvert
	}
	isEqual := reflect.DeepEqual(reExp.Spec, reCur.Spec)
	if isEqual {
		return ErrNotChanged
	}

	if err := c.r.Create(ctx, reCur); err != nil {
		return err
	}

	if err := c.r.Update(ctx, reExp); err != nil {
		return err
	}

	return nil
}

func (c *chaosMeshHandler) UpdatePodChaos(ctx context.Context, ssChaos *v1alpha1.ShardingSphereChaos, cur PodChaos) error {
	podChao, err := c.NewPodChaos(ssChaos)
	if err != nil {
		return err
	}
	reExp, ok := (podChao).(*chaosv1alpha1.PodChaos)
	if !ok {
		return ErrConvert
	}
	reCur, ok := cur.(*chaosv1alpha1.PodChaos)
	if !ok {
		return ErrConvert
	}
	isEqual := reflect.DeepEqual(reExp.Spec, reCur.Spec)
	if isEqual {
		return ErrNotChanged
	}

	if err := c.r.Delete(ctx, reCur); err != nil {
		return err
	}

	if err := c.CreatePodChaos(ctx, reExp); err != nil {
		return err
	}

	return nil
}

type PodChaosBuilder interface {
	SetNamespace(string) PodChaosBuilder
	SetName(string) PodChaosBuilder
	SetLabels(map[string]string) PodChaosBuilder
	SetAnnotations(map[string]string) PodChaosBuilder
	SetContainerSelector(*chaosv1alpha1.ContainerSelector) PodChaosBuilder
	SetAction(string) PodChaosBuilder
	SetDuration(string) PodChaosBuilder
	SetGracePeriod(int64) PodChaosBuilder
	Build() *chaosv1alpha1.PodChaos
}

func NewPodChaosBuilder() PodChaosBuilder {
	return &podChaosBuilder{
		podChaos: DefaultPodChaos(),
	}
}

type podChaosBuilder struct {
	podChaos *chaosv1alpha1.PodChaos
}

type BandWidthActionBuilder interface {
	SetRate(string) BandWidthActionBuilder
	SetLimit(string) BandWidthActionBuilder
	SetBuffer(string) BandWidthActionBuilder
	SetPeakRate(string) BandWidthActionBuilder
	SetMinBurst(string) BandWidthActionBuilder
	Build() *chaosv1alpha1.BandwidthSpec
}

func NewBandWidthActionBuilder() BandWidthActionBuilder {
	return &bandWidthActionBuilder{
		bandwidth: &chaosv1alpha1.BandwidthSpec{},
	}
}

type bandWidthActionBuilder struct {
	bandwidth *chaosv1alpha1.BandwidthSpec
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

func (b *bandWidthActionBuilder) Build() *chaosv1alpha1.BandwidthSpec {
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

func (p *podChaosBuilder) SetContainerSelector(selector *chaosv1alpha1.ContainerSelector) PodChaosBuilder {
	p.podChaos.Spec.ContainerSelector = *selector
	return p
}

func (p *podChaosBuilder) SetAction(action string) PodChaosBuilder {
	if v1alpha1.PodChaosAction(action) == v1alpha1.PodFailureAction {
		p.podChaos.Spec.Action = chaosv1alpha1.PodFailureAction
	}

	if v1alpha1.PodChaosAction(action) == v1alpha1.ContainerKillAction {
		p.podChaos.Spec.Action = chaosv1alpha1.ContainerKillAction
	}

	if chaosv1alpha1.PodChaosAction(action) == chaosv1alpha1.PodKillAction {
		p.podChaos.Spec.Action = chaosv1alpha1.PodKillAction
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
	SetAction(string) NetworkChaosBuilder
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

func (n *netWorkChaosBuilder) SetAction(action string) NetworkChaosBuilder {

	if chaosv1alpha1.NetworkChaosAction(action) == chaosv1alpha1.BandwidthAction {
		n.netWorkChaos.Spec.Action = chaosv1alpha1.BandwidthAction
	}

	if v1alpha1.NetworkChaosAction(action) == v1alpha1.CorruptAction {
		n.netWorkChaos.Spec.Action = chaosv1alpha1.CorruptAction
	}

	if v1alpha1.NetworkChaosAction(action) == v1alpha1.PartitionAction {
		n.netWorkChaos.Spec.Action = chaosv1alpha1.PartitionAction
	}

	if v1alpha1.NetworkChaosAction(action) == v1alpha1.LossAction {
		n.netWorkChaos.Spec.Action = chaosv1alpha1.LossAction
	}

	if v1alpha1.NetworkChaosAction(action) == v1alpha1.DuplicateAction {
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

func DefaultPodChaos() *chaosv1alpha1.PodChaos {
	return &chaosv1alpha1.PodChaos{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "shardingsphere-proxy",
			Namespace: "default",
			Labels:    map[string]string{},
		},
		Spec: chaosv1alpha1.PodChaosSpec{
			Action: chaosv1alpha1.ContainerKillAction,
		},
	}
}

func DefaultNetworkChaos() *chaosv1alpha1.NetworkChaos {
	return &chaosv1alpha1.NetworkChaos{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "shardingsphere-proxy",
			Namespace: "default",
			Labels:    map[string]string{},
		},
		Spec: chaosv1alpha1.NetworkChaosSpec{
			Action:    chaosv1alpha1.PartitionAction,
			Direction: "to",
		},
	}
}
