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

}
