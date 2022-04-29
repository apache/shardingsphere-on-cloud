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

package reconcile

import (
	v1 "k8s.io/api/core/v1"
)

func IsRunning(podList *v1.PodList) bool {
	status := true
	for _, pod := range podList.Items {
		if pod.Status.Phase != v1.PodRunning {
			status = false
		}
	}

	return status
}

func CountingReadyPods(podList *v1.PodList) int32 {
	var readyPods int32
	readyPods = 0
	for _, pod := range podList.Items {
		if pod.Status.ContainerStatuses[0].Ready && pod.ObjectMeta.DeletionTimestamp == nil {
			readyPods++
		}
	}
	return readyPods
}

func CountingPodMaxRestartTimes(podList *v1.PodList) int32 {
	var podRestartCount int32 = math.MinInt32
	for _, pod := range podList.Items {
		if podRestartCount < pod.Status.ContainerStatuses[0].RestartCount {
			podRestartCount = pod.Status.ContainerStatuses[0].RestartCount
		}
	}
	return podRestartCount
}
