// Copyright 2022 SphereEx Authors
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

package reconcile

import (
	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func Test_IsRunning(t *testing.T) {
	ts := metav1.Time{}
	cases := []struct {
		podlist *v1.PodList
		exp     bool
		message string
	}{
		{
			podlist: &v1.PodList{
				Items: []v1.Pod{},
			},
			exp:     false,
			message: "Empty PodList should be false",
		},
		{
			podlist: &v1.PodList{
				Items: []v1.Pod{
					{
						Status: v1.PodStatus{
							Phase: v1.PodRunning,
						},
					},
					{
						Status: v1.PodStatus{
							Phase: v1.PodFailed,
						},
					},
				},
			},
			// At least one Pod is Running considered the Cluster be availabe
			exp:     true,
			message: "First Pod is Running and second Pod is not Running should be false",
		},
		{
			podlist: &v1.PodList{
				Items: []v1.Pod{
					{
						Status: v1.PodStatus{
							Phase: v1.PodFailed,
						},
					},
					{
						Status: v1.PodStatus{
							Phase: v1.PodRunning,
						},
					},
				},
			},
			// At least one Pod is Running considered the Cluster be availabe
			exp:     true,
			message: "First Pod is not Running and second Pod is Running should be false",
		},
		{
			podlist: &v1.PodList{
				Items: []v1.Pod{
					{
						Status: v1.PodStatus{
							Phase: v1.PodRunning,
						},
					},
					{
						Status: v1.PodStatus{
							Phase: v1.PodRunning,
						},
					},
				},
			},
			exp:     true,
			message: "All Pods are running should be true",
		},
		{
			podlist: &v1.PodList{
				Items: []v1.Pod{
					{
						Status: v1.PodStatus{
							Phase: v1.PodFailed,
						},
					},
					{
						Status: v1.PodStatus{
							Phase: v1.PodFailed,
						},
					},
				},
			},
			exp:     false,
			message: "All Pods are not running should be false",
		},
		{
			podlist: &v1.PodList{
				Items: []v1.Pod{
					{
						ObjectMeta: metav1.ObjectMeta{
							DeletionTimestamp: &ts,
						},
						Status: v1.PodStatus{
							Phase: v1.PodRunning,
						},
					},
				},
			},
			exp:     false,
			message: "Pod is running with deletion timestamp should be false",
		},
	}

	for _, c := range cases {
		act := IsRunning(c.podlist)
		if len(c.podlist.Items) != 0 {
			assert.Equal(t, c.exp, act, c.message)
		}
	}
}

func Test_CountingReadyPods(t *testing.T) {
	ts := metav1.Time{}
	cases := []struct {
		podlist *v1.PodList
		exp     int32
		message string
	}{
		{
			podlist: &v1.PodList{
				Items: []v1.Pod{},
			},
			exp:     0,
			message: "Empty PodList should be 0",
		},
		{
			podlist: &v1.PodList{
				Items: []v1.Pod{
					{
						Status: v1.PodStatus{
							ContainerStatuses: []v1.ContainerStatus{},
						},
					},
				},
			},
			exp:     0,
			message: "Only one Pod without any container statuses should be 0",
		},
		{
			podlist: &v1.PodList{
				Items: []v1.Pod{
					{
						Status: v1.PodStatus{
							ContainerStatuses: []v1.ContainerStatus{
								{
									Ready: true,
								},
							},
						},
					},
				},
			},
			exp:     1,
			message: "Only one Pod is running should be 1",
		},
		{
			podlist: &v1.PodList{
				Items: []v1.Pod{
					{
						ObjectMeta: metav1.ObjectMeta{
							DeletionTimestamp: &ts,
						},
						Status: v1.PodStatus{
							ContainerStatuses: []v1.ContainerStatus{
								{
									Ready: true,
								},
							},
						},
					},
				},
			},
			exp:     0,
			message: "Pod has ready container but with deletion timestamp should be 0",
		},
		{
			podlist: &v1.PodList{
				Items: []v1.Pod{
					{
						ObjectMeta: metav1.ObjectMeta{
							DeletionTimestamp: &ts,
						},
						Status: v1.PodStatus{
							ContainerStatuses: []v1.ContainerStatus{
								{
									Ready: true,
								},
							},
						},
					},
					{
						Status: v1.PodStatus{
							ContainerStatuses: []v1.ContainerStatus{
								{
									Ready: true,
								},
							},
						},
					},
				},
			},
			exp:     1,
			message: "First Pod has ready container, second Pod has ready container but with deletion timestamp should be 0",
		},
		{
			podlist: &v1.PodList{
				Items: []v1.Pod{
					{
						Status: v1.PodStatus{
							ContainerStatuses: []v1.ContainerStatus{
								{
									Ready: true,
								},
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							DeletionTimestamp: &ts,
						},
						Status: v1.PodStatus{
							ContainerStatuses: []v1.ContainerStatus{
								{
									Ready: true,
								},
							},
						},
					},
				},
			},
			exp:     1,
			message: "First Pod has ready container but with deletion timestamp, second Pod has ready container should be 0",
		},
	}

	for _, c := range cases {
		act := CountingReadyPods(c.podlist)
		assert.Equal(t, c.exp, act, c.message)
	}
}
