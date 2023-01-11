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
	ConditionDeployed    ConditionType = "Deployed"
	ConditionInitialized ConditionType = "Initialized"
	ConditionStarted     ConditionType = "Started"
	ConditionReady       ConditionType = "Ready"
	ConditionUnknown     ConditionType = "Unknown"
	ConditionFailed      ConditionType = "Failed"
)

// ProxyStatus defines the observed state of ShardingSphereProxy
type ProxyStatus struct {
	//ShardingSphere-Proxy phase are a brief summary of the ShardingSphere-Proxy life cycle
	//There are two possible phase values:
	//Ready: ShardingSphere-Proxy can already provide external services
	//NotReady: ShardingSphere-Proxy cannot provide external services
	Phase PhaseStatus `json:"phase"`

	//Conditions The conditions array, the reason and message fields
	Conditions []Condition `json:"conditions"`
	//ReadyNodes shows the number of replicas that ShardingSphere-Proxy is running normally
	ReadyNodes int32 `json:"readyNodes"`
}

type Conditions []Condition

// Condition
// | **phase** | **condition**  | **descriptions**|
// | ------------- | ---------- | ---------------------------------------------------- |
// | NotReady      | Deployed   | pods are deployed but are not created or currently pending|
// | NotReady      | Started    | pods are started but not satisfy ready requirements|
// | Ready         | Ready      | minimum pods satisfy ready requirements|
// | NotReady      | Unknown    | can not locate the status of pods |
// | NotReady      | Failed     | ShardingSphere-Proxy failed to start correctly due to some problems|

type Condition struct {
	Type           ConditionType      `json:"type"`
	Status         v1.ConditionStatus `json:"status"`
	LastUpdateTime metav1.Time        `json:"lastUpdateTime,omitempty"`
}

func (p *ShardingSphereProxy) SetInitialized() {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionInitialized,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
}

func (p *ShardingSphereProxy) SetInitializationFailed() {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionInitialized,
		Status:         v1.ConditionFalse,
		LastUpdateTime: metav1.Now(),
	})
}

func (p *ShardingSphereProxy) SetPodStarted(readyNodes int32) {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionStarted,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
	p.Status.ReadyNodes = readyNodes
}

func (p *ShardingSphereProxy) SetPodNotStarted(readyNodes int32) {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionStarted,
		Status:         v1.ConditionFalse,
		LastUpdateTime: metav1.Now(),
	})
	p.Status.ReadyNodes = readyNodes
}

func (p *ShardingSphereProxy) SetReady(readyNodes int32) {
	p.Status.Phase = StatusReady
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionReady,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
	p.Status.ReadyNodes = readyNodes

}

func (p *ShardingSphereProxy) SetFailed() {
	p.Status.Phase = StatusNotReady
	p.Status.Conditions = append([]Condition{}, Condition{
		Type:           ConditionUnknown,
		Status:         v1.ConditionTrue,
		LastUpdateTime: metav1.Now(),
	})
}
func (p *ShardingSphereProxy) UpdateReadyNodes(readyNodes int32) {
	p.Status.ReadyNodes = readyNodes
}
