// Copyright 2023 SphereEx Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"testing"

	v1 "k8s.io/api/core/v1"
)

func Test_GetReadyProxyInstances(t *testing.T) {
	// create sample PodList
	podlist := v1.PodList{
		Items: []v1.Pod{
			{
				Status: v1.PodStatus{
					Phase: v1.PodRunning,
					Conditions: []v1.PodCondition{
						{
							Type:   v1.PodReady,
							Status: v1.ConditionTrue,
						},
					},
					ContainerStatuses: []v1.ContainerStatus{
						{
							Name:  "shardingsphere-proxy",
							Ready: true,
						},
					},
				},
			},
			{
				Status: v1.PodStatus{
					Phase: v1.PodRunning,
					Conditions: []v1.PodCondition{
						{
							Type:   v1.PodReady,
							Status: v1.ConditionTrue,
						},
					},
					ContainerStatuses: []v1.ContainerStatus{
						{
							Name:  "another-container",
							Ready: true,
						},
					},
				},
			},
			{
				Status: v1.PodStatus{
					Phase: v1.PodRunning,
					Conditions: []v1.PodCondition{
						{
							Type:   v1.PodReady,
							Status: v1.ConditionFalse,
						},
					},
					ContainerStatuses: []v1.ContainerStatus{
						{
							Name:  "shardingsphere-proxy",
							Ready: false,
						},
					},
				},
			},
			{
				Status: v1.PodStatus{
					Phase: v1.PodPending,
					Conditions: []v1.PodCondition{
						{
							Type:   v1.PodReady,
							Status: v1.ConditionTrue,
						},
					},
					ContainerStatuses: []v1.ContainerStatus{
						{
							Name:  "shardingsphere-proxy",
							Ready: true,
						},
					},
				},
			},
		},
	}

	// expected result is 1 because only one pod has a ready shardingsphere-proxy container
	expected := int32(1)

	// call the function to get the actual result
	actual := getReadyProxyInstances(podlist)

	// compare the expected and actual results
	if actual != expected {
		t.Errorf("getReadyInstances returned %d, expected %d", actual, expected)
	}
}
