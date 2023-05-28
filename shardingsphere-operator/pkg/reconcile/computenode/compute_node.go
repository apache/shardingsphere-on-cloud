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

// GetConditionFromPods returns the condition for a pod
func GetConditionFromPods(podlist *corev1.PodList) []v1alpha1.ComputeNodeCondition {
	conds := []v1alpha1.ComputeNodeCondition{}

	if len(podlist.Items) == 0 {
		conds = append(conds, newCondition(v1alpha1.ComputeNodeConditionUnknown, "PodNotFound", "No pod was found"))
		return conds
	}

	result := map[v1alpha1.ComputeNodeConditionType]int{}
	for i := range podlist.Items {
		pcs := getPreferedConditionFromPod(&podlist.Items[i])
		for idx := range pcs {
			result[pcs[idx].Type]++
		}
	}

	if result[v1alpha1.ComputeNodeConditionUnknown] == len(podlist.Items) {
		conds = append(conds, newCondition(v1alpha1.ComputeNodeConditionUnknown, "PodUnknown", "All pods are unknown"))
	}

	if result[v1alpha1.ComputeNodeConditionReady] > 0 {
		conds = append(conds, newCondition(v1alpha1.ComputeNodeConditionReady, "PodReady", "Some pods are ready"))
	}

	if result[v1alpha1.ComputeNodeConditionStarted] > 0 {
		conds = append(conds, newCondition(v1alpha1.ComputeNodeConditionStarted, "PodStarted", "Some pods are started"))
	}

	if result[v1alpha1.ComputeNodeConditionInitialized] > 0 {
		conds = append(conds, newCondition(v1alpha1.ComputeNodeConditionInitialized, "PodInitialized", "Some pods are initialized"))
	}

	if result[v1alpha1.ComputeNodeConditionDeployed] > 0 {
		conds = append(conds, newCondition(v1alpha1.ComputeNodeConditionDeployed, "PodDeployed", "Some pods are deployed"))
	}

	if result[v1alpha1.ComputeNodeConditionPending] > 0 {
		conds = append(conds, newCondition(v1alpha1.ComputeNodeConditionPending, "PodPending", "Some pods are pending"))
	}

	if result[v1alpha1.ComputeNodeConditionFailed] > 0 {
		conds = append(conds, newCondition(v1alpha1.ComputeNodeConditionFailed, "PodFailed", "Some pods are failed"))
	}

	return conds
}

func getPreferedConditionFromPod(pod *corev1.Pod) []v1alpha1.ComputeNodeCondition {
	conds := []v1alpha1.ComputeNodeCondition{}
	if pod.Status.Phase == corev1.PodUnknown {
		conds = append(conds, v1alpha1.ComputeNodeCondition{
			Type: v1alpha1.ComputeNodeConditionUnknown,
		})
		return conds
	}

	if pod.Status.Phase == corev1.PodPending {
		conds = append(conds, v1alpha1.ComputeNodeCondition{
			Type: v1alpha1.ComputeNodeConditionPending,
		})
		return conds
	}

	if pod.Status.Phase == corev1.PodFailed {
		conds = append(conds, v1alpha1.ComputeNodeCondition{
			Type: v1alpha1.ComputeNodeConditionFailed,
		})
		return conds
	}

	return getPreferedConditionFromPodConditions(pod.Status.Conditions)
}

func getPreferedConditionFromPodConditions(pcs []corev1.PodCondition) []v1alpha1.ComputeNodeCondition {
	conditions := []v1alpha1.ComputeNodeCondition{}

	for i := range pcs {
		if pcs[i].Status != corev1.ConditionTrue {
			continue
		}

		if pcs[i].Type == corev1.PodScheduled {
			conditions = append(conditions, v1alpha1.ComputeNodeCondition{
				Type: v1alpha1.ComputeNodeConditionDeployed,
			})
		}
		if pcs[i].Type == corev1.PodInitialized {
			conditions = append(conditions, v1alpha1.ComputeNodeCondition{
				Type: v1alpha1.ComputeNodeConditionInitialized,
			})
		}
		if pcs[i].Type == corev1.ContainersReady {
			conditions = append(conditions, v1alpha1.ComputeNodeCondition{
				Type: v1alpha1.ComputeNodeConditionStarted,
			})
		}
		if pcs[i].Type == corev1.PodReady {
			conditions = append(conditions, v1alpha1.ComputeNodeCondition{
				Type: v1alpha1.ComputeNodeConditionReady,
			})
		}
	}

	return conditions
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
