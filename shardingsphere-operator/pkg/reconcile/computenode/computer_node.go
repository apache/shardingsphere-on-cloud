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

package computenode

import (
	"time"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetConditionFromPods(podlist *corev1.PodList) v1alpha1.ComputeNodeCondition {
	if len(podlist.Items) == 0 {
		return newCondition(v1alpha1.ComputeNodeConditionUnknown, "PodNotFound", "No pod was found")
	}
	var cond v1alpha1.ComputeNodeCondition
	result := map[v1alpha1.ComputeNodeConditionType]int{}
	for _, p := range podlist.Items {
		pc := getPreferedConditionFromPod(p)
		result[pc.Type]++
	}

	if result[v1alpha1.ComputeNodeConditionUnknown] == len(podlist.Items) {
		return newCondition(v1alpha1.ComputeNodeConditionUnknown, "PodUnknown", "All pods are unknown")
	}

	if result[v1alpha1.ComputeNodeConditionReady] > 0 {
		return newCondition(v1alpha1.ComputeNodeConditionReady, "PodReady", "Some pods are ready")
	}

	if result[v1alpha1.ComputeNodeConditionStarted] > 0 {
		return newCondition(v1alpha1.ComputeNodeConditionStarted, "PodStarted", "Some pods are started")
	}

	if result[v1alpha1.ComputeNodeConditionInitialized] > 0 {
		return newCondition(v1alpha1.ComputeNodeConditionInitialized, "PodInitialized", "Some pods are initialized")
	}

	if result[v1alpha1.ComputeNodeConditionDeployed] > 0 {
		return newCondition(v1alpha1.ComputeNodeConditionDeployed, "PodDeployed", "Some pods are deployed")
	}

	if result[v1alpha1.ComputeNodeConditionPending] > 0 {
		return newCondition(v1alpha1.ComputeNodeConditionPending, "PodPending", "Some pods are pending")
	}

	if result[v1alpha1.ComputeNodeConditionFailed] > 0 {
		return newCondition(v1alpha1.ComputeNodeConditionFailed, "PodFailed", "Some pods are failed")
	}

	return cond
}

func getPreferedConditionFromPod(pod corev1.Pod) v1alpha1.ComputeNodeCondition {
	if pod.Status.Phase == corev1.PodUnknown {
		return v1alpha1.ComputeNodeCondition{
			Type: v1alpha1.ComputeNodeConditionUnknown,
		}
	}

	if pod.Status.Phase == corev1.PodPending {
		return v1alpha1.ComputeNodeCondition{
			Type: v1alpha1.ComputeNodeConditionPending,
		}
	}

	if pod.Status.Phase == corev1.PodFailed {
		return v1alpha1.ComputeNodeCondition{
			Type: v1alpha1.ComputeNodeConditionFailed,
		}
	}

	return getPreferedConditionFromPodConditions(pod.Status.Conditions)
}

func getPreferedConditionFromPodConditions(conditions []corev1.PodCondition) v1alpha1.ComputeNodeCondition {
	var (
		sched       bool
		initialized bool
		conReady    bool
		ready       bool
	)

	for _, c := range conditions {
		if c.Type == corev1.PodScheduled && c.Status == corev1.ConditionTrue {
			sched = true
		}
		if c.Type == corev1.PodInitialized && c.Status == corev1.ConditionTrue {
			initialized = true
		}
		if c.Type == corev1.ContainersReady && c.Status == corev1.ConditionTrue {
			conReady = true
		}
		if c.Type == corev1.PodReady && c.Status == corev1.ConditionTrue {
			ready = true
		}
	}

	if ready {
		return v1alpha1.ComputeNodeCondition{
			Type: v1alpha1.ComputeNodeConditionReady,
		}
	}

	if conReady {
		return v1alpha1.ComputeNodeCondition{
			Type: v1alpha1.ComputeNodeConditionStarted,
		}
	}

	if initialized {
		return v1alpha1.ComputeNodeCondition{
			Type: v1alpha1.ComputeNodeConditionInitialized,
		}
	}

	if sched {
		return v1alpha1.ComputeNodeCondition{
			Type: v1alpha1.ComputeNodeConditionDeployed,
		}
	}

	return v1alpha1.ComputeNodeCondition{}
}

func newCondition(t v1alpha1.ComputeNodeConditionType, reason, message string) v1alpha1.ComputeNodeCondition {
	return v1alpha1.ComputeNodeCondition{
		Type:               t,
		Status:             v1alpha1.ConditionStatusTrue,
		LastUpdateTime:     metav1.NewTime(time.Now()),
		LastTransitionTime: metav1.NewTime(time.Now()),
		Reason:             reason,
		Message:            message,
	}
}

func setCondition(conditions []v1alpha1.ComputeNodeCondition, t v1alpha1.ComputeNodeConditionType, reason, message string, exlusive bool) {
	cond := v1alpha1.ComputeNodeCondition{
		Type:               t,
		Status:             v1alpha1.ConditionStatusTrue,
		LastUpdateTime:     metav1.NewTime(time.Now()),
		LastTransitionTime: metav1.NewTime(time.Now()),
		Reason:             reason,
		Message:            message,
	}

	var found bool
	for i := range conditions {
		if conditions[i].Type == cond.Type {
			found = true
			conditions[i] = cond
		} else {
			if cond.Type != v1alpha1.ComputeNodeConditionUnknown {

			}
			if exlusive {
				conditions[i].LastUpdateTime = cond.LastUpdateTime
				conditions[i].Status = v1alpha1.ConditionStatusFalse
			}
		}
	}

	// check current conditions
	if len(conditions) == 0 || !found {
		conditions = append(conditions, cond)
	}
}
