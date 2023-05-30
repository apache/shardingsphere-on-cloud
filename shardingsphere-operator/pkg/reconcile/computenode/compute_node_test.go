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
	"testing"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func Test_getPreferedConditionFromPodConditions(t *testing.T) {
	cases := []struct {
		pcs     []corev1.PodCondition
		expect  []v1alpha1.ComputeNodeCondition
		message string
	}{
		{
			pcs: []corev1.PodCondition{
				corev1.PodCondition{
					Type:   corev1.PodScheduled,
					Status: corev1.ConditionTrue,
				},
			},
			expect: []v1alpha1.ComputeNodeCondition{
				{
					Type: v1alpha1.ComputeNodeConditionDeployed,
				},
			},
			message: "scheduled pod condition should get deployed condition",
		},
	}
	for _, c := range cases {
		act := getPreferedConditionFromPodConditions(c.pcs)
		assert.Equal(t, len(c.expect), len(act), c.message)
		for i := range act {
			assert.Equal(t, c.expect[i].Type, act[i].Type, c.message)
		}
	}
}

func Test_getPreferedConditionFromPod(t *testing.T) {
	cases := []struct {
		pod     corev1.Pod
		expect  []v1alpha1.ComputeNodeCondition
		message string
	}{
		{
			pod: corev1.Pod{
				Status: corev1.PodStatus{
					Phase: corev1.PodPending,
				},
			},
			expect: []v1alpha1.ComputeNodeCondition{
				{
					Type:   v1alpha1.ComputeNodeConditionPending,
					Status: v1alpha1.ConditionStatusTrue,
				},
			},
			message: "pending pod should get pending condition",
		},
	}

	for _, c := range cases {
		act := getPreferedConditionFromPod(&c.pod)
		assert.Equal(t, len(c.expect), len(act), c.message)
		for i := range act {
			assert.Equal(t, c.expect[i].Type, act[i].Type, c.message)
		}
	}
}

func Test_GetConditionFromPods(t *testing.T) {
	cases := []struct {
		podlist *corev1.PodList
		expect  []v1alpha1.ComputeNodeCondition
		message string
	}{
		{
			podlist: &corev1.PodList{Items: []corev1.Pod{}},
			expect: []v1alpha1.ComputeNodeCondition{
				{
					Type:   v1alpha1.ComputeNodeConditionUnknown,
					Status: v1alpha1.ConditionStatusTrue,
				},
			},
			message: "empty podlist should get unknown condition",
		},
	}

	for _, c := range cases {
		act := GetConditionFromPods(c.podlist)
		assert.Equal(t, len(c.expect), len(act), c.message)
		for i := range act {
			assert.Equal(t, c.expect[i].Type, act[i].Type, c.message)
			assert.Equal(t, c.expect[i].Status, act[i].Status, c.message)
		}
	}

}
