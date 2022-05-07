/*
 *   Copyright © 2022，Beijing Sifei Software Technology Co., LTD.
 *   All Rights Reserved.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
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

// ConditionType shows some states during the startup process of ShardingSphere-Proxy
const (
	ConditionInitialized ConditionType = "Initialized"
	ConditionStarted     ConditionType = "Started"
	ConditionReady       ConditionType = "Ready"
	ConditionUnknown     ConditionType = "Unknown"
)

// ProxyStatus defines the observed state of Proxy
type ProxyStatus struct {
	//ShardingSphere-Proxy phase are a brief summary of the ShardingSphere-Proxy life cycle
	//There are two possible phase values:
	//Ready: ShardingSphere-Proxy can already provide external services
	//NotReady: ShardingSphere-Proxy cannot provide external services
	Phase PhaseStatus `json:"phase"`

	//Conditions The conditions array, the reason and message fields
	Conditions Conditions `json:"conditions"`
	//ReadyNodes shows the number of replicas that ShardingSphere-Proxy is running normally
	ReadyNodes int32 `json:"readyNodes"`
}

type Conditions []Condition

//| **condition** | **status** | **directions**|
//| ------------- | ---------- | ---------------------------------------------------- |
//| Initialized   | true       | Initialization successful|
//| Initialized   | false      | initialization failed|
//| Started       | true       | pod started successfully but not ready|
//| Started       | false      | pod started failed|
//| Ready         | true       | The pod is ready and can provide external services|
//| Unknown       | true       | ShardingSphere-Proxy failed to start correctly due to some problems |
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
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionStarted,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
	p.Status.ReadyNodes = readyNodes
}

func (p *Proxy) SetPodNotStarted(readyNodes int32) {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionStarted,
		Status:         v1.ConditionFalse,
		LastUpdateTime: metav1.Now(),
	})
	p.Status.ReadyNodes = readyNodes
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
