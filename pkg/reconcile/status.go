package reconcile

import v1 "k8s.io/api/core/v1"

func IsRunning(podList *v1.PodList) bool {
	status := true
	for _, pod := range podList.Items {
		if pod.Status.Phase != v1.PodRunning {
			status = false
		}
	}

	return status
}

func ReadyCount(podList *v1.PodList) int32 {
	var readyPods int32
	readyPods = 0
	for _, pod := range podList.Items {
		if pod.Status.ContainerStatuses[0].Ready && pod.ObjectMeta.DeletionTimestamp == nil {
			readyPods++
		}
	}
	return readyPods
}
