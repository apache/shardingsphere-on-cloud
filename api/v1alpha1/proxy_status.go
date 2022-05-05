/*
 * Copyright (c) 2022.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PhaseStatus string

const (
	StatusReady    PhaseStatus = "Ready"
	StatusNotReady PhaseStatus = "NotReady"
)

type ConditionType string

const (
	ConditionInitialized ConditionType = "Initialized"
	ConditionStarted     ConditionType = "Started"
	ConditionReady       ConditionType = "Ready"
	ConditionUnknown     ConditionType = "Unknown"
)

// ProxyStatus defines the observed state of Proxy
type ProxyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// TODO:description
	Phase PhaseStatus `json:"phase"`
	// TODO:description
	Conditions Conditions `json:"conditions"`
	// TODO:description
	ReadyNodes int32 `json:"readyNodes"`
}

type Conditions []Condition

type Condition struct {
	Type           ConditionType      `json:"type"`
	Status         v1.ConditionStatus `json:"status"`
	LastUpdateTime metav1.Time        `json:"lastUpdateTime,omitempty"`
}

func (p *Proxy) SetInitialized() {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionInitialized,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
}

func (p *Proxy) SetInitializationFailed() {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionInitialized,
		Status:         v1.ConditionFalse,
		LastUpdateTime: metav1.Now(),
	})
}

func (p *Proxy) SetPodStarted(readyNodes int32) {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append(p.Status.Conditions, Condition{
		Type:           ConditionStarted,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
	p.Status.ReadyNodes = readyNodes
}

func (p *Proxy) SetPodNotStarted() {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append(p.Status.Conditions, Condition{
		Type:           ConditionStarted,
		Status:         v1.ConditionFalse,
		LastUpdateTime: metav1.Now(),
	})
}

func (p *Proxy) SetReady(readyNodes int32) {
	p.Status.Phase = StatusReady
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionReady,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
	p.Status.ReadyNodes = readyNodes

}

func (p *Proxy) SetFailed() {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionUnknown,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
}
