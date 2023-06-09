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

package computenode_test

import (
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/computenode"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
)

var _ = Describe("GetConditionFromPods", func() {
	Context("Empty Pod list", func() {
		podlist := &corev1.PodList{Items: []corev1.Pod{}}
		conditions := computenode.GetConditionFromPods(podlist)
		It("should be one unknown condition", func() {
			Expect(len(conditions)).To(Equal(1))
		})

		It("unknown condition should be true", func() {
			Expect(conditions[0].Type).To(Equal(v1alpha1.ComputeNodeConditionUnknown))
			Expect(conditions[0].Status).To(Equal(v1alpha1.ConditionStatusTrue))
		})
	})

	Context("One Pod Pending", func() {
		podlist := &corev1.PodList{Items: []corev1.Pod{
			{
				Status: corev1.PodStatus{
					Phase: corev1.PodPending,
				},
			},
		}}
		conditions := computenode.GetConditionFromPods(podlist)

		It("should be one pending condition", func() {
			Expect(len(conditions)).To(Equal(1))
		})

		It("unknown condition should be true", func() {
			Expect(conditions[0].Type).To(Equal(v1alpha1.ComputeNodeConditionPending))
			Expect(conditions[0].Status).To(Equal(v1alpha1.ConditionStatusTrue))
		})
	})

	Context("One Pod Scheduled", func() {
		podlist := &corev1.PodList{Items: []corev1.Pod{
			{
				Status: corev1.PodStatus{
					Phase: corev1.PodPending,
					Conditions: []corev1.PodCondition{
						{
							Type:   corev1.PodReady,
							Status: corev1.ConditionFalse,
						},
						{
							Type:   corev1.ContainersReady,
							Status: corev1.ConditionFalse,
						},
						{
							Type:   corev1.PodInitialized,
							Status: corev1.ConditionFalse,
						},
						{
							Type:   corev1.PodScheduled,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		}}
		conditions := computenode.GetConditionFromPods(podlist)

		It("should be two conditions", func() {
			Expect(len(conditions)).To(Equal(2))
		})

		It("condition status should be correct", func() {
			Expect(containConditionType(conditions, v1alpha1.ComputeNodeConditionPending, v1alpha1.ComputeNodeConditionDeployed)).To(BeTrue())
			for _, cond := range conditions {
				if cond.Type == v1alpha1.ComputeNodeConditionPending {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
				if cond.Type == v1alpha1.ComputeNodeConditionDeployed {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
			}
		})
	})

	Context("One Pod Initialized", func() {
		podlist := &corev1.PodList{Items: []corev1.Pod{
			{
				Status: corev1.PodStatus{
					Phase: corev1.PodPending,
					Conditions: []corev1.PodCondition{
						{
							Type:   corev1.PodReady,
							Status: corev1.ConditionFalse,
						},
						{
							Type:   corev1.ContainersReady,
							Status: corev1.ConditionFalse,
						},
						{
							Type:   corev1.PodInitialized,
							Status: corev1.ConditionTrue,
						},
						{
							Type:   corev1.PodScheduled,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		}}
		conditions := computenode.GetConditionFromPods(podlist)

		It("should be three conditions", func() {
			Expect(len(conditions)).To(Equal(3))
		})

		It("condition status should be correct", func() {
			Expect(containConditionType(conditions, v1alpha1.ComputeNodeConditionPending, v1alpha1.ComputeNodeConditionDeployed, v1alpha1.ComputeNodeConditionInitialized)).To(BeTrue())
			for _, cond := range conditions {
				if cond.Type == v1alpha1.ComputeNodeConditionPending {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
				if cond.Type == v1alpha1.ComputeNodeConditionDeployed {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
				if cond.Type == v1alpha1.ComputeNodeConditionInitialized {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
			}
		})
	})

	Context("One Pod Ready", func() {
		podlist := &corev1.PodList{Items: []corev1.Pod{
			{
				Status: corev1.PodStatus{
					Phase: corev1.PodPending,
					Conditions: []corev1.PodCondition{
						{
							Type:   corev1.PodReady,
							Status: corev1.ConditionTrue,
						},
						{
							Type:   corev1.ContainersReady,
							Status: corev1.ConditionTrue,
						},
						{
							Type:   corev1.PodInitialized,
							Status: corev1.ConditionTrue,
						},
						{
							Type:   corev1.PodScheduled,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		}}
		conditions := computenode.GetConditionFromPods(podlist)
		It("should be five conditions", func() {
			Expect(len(conditions)).To(Equal(5))
		})

		It("condition status should be correct", func() {
			Expect(containConditionType(conditions, v1alpha1.ComputeNodeConditionPending, v1alpha1.ComputeNodeConditionDeployed, v1alpha1.ComputeNodeConditionInitialized, v1alpha1.ComputeNodeConditionStarted, v1alpha1.ComputeNodeConditionReady)).To(BeTrue())
			for _, cond := range conditions {
				if cond.Type == v1alpha1.ComputeNodeConditionPending {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
				if cond.Type == v1alpha1.ComputeNodeConditionDeployed {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
				if cond.Type == v1alpha1.ComputeNodeConditionInitialized {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
				if cond.Type == v1alpha1.ComputeNodeConditionStarted {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
				if cond.Type == v1alpha1.ComputeNodeConditionReady {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
			}
		})
	})

	Context("One Pod Running", func() {
		podlist := &corev1.PodList{Items: []corev1.Pod{
			{
				Status: corev1.PodStatus{
					Phase: corev1.PodRunning,
					Conditions: []corev1.PodCondition{
						{
							Type:   corev1.PodReady,
							Status: corev1.ConditionTrue,
						},
						{
							Type:   corev1.ContainersReady,
							Status: corev1.ConditionTrue,
						},
						{
							Type:   corev1.PodInitialized,
							Status: corev1.ConditionTrue,
						},
						{
							Type:   corev1.PodScheduled,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		}}
		conditions := computenode.GetConditionFromPods(podlist)
		It("should be five conditions", func() {
			Expect(len(conditions)).To(Equal(5))
		})

		It("condition status should be correct", func() {
			Expect(containConditionType(conditions, v1alpha1.ComputeNodeConditionPending, v1alpha1.ComputeNodeConditionDeployed, v1alpha1.ComputeNodeConditionInitialized, v1alpha1.ComputeNodeConditionStarted, v1alpha1.ComputeNodeConditionReady)).To(BeTrue())
			for _, cond := range conditions {
				if cond.Type == v1alpha1.ComputeNodeConditionPending {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
				if cond.Type == v1alpha1.ComputeNodeConditionDeployed {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
				if cond.Type == v1alpha1.ComputeNodeConditionInitialized {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
				if cond.Type == v1alpha1.ComputeNodeConditionStarted {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
				if cond.Type == v1alpha1.ComputeNodeConditionReady {
					Expect(cond.Status).To(Equal(v1alpha1.ConditionStatusTrue))
				}
			}
		})
	})
})

func containConditionType(conds []v1alpha1.ComputeNodeCondition, ts ...v1alpha1.ComputeNodeConditionType) bool {
	if len(conds) != len(ts) {
		return false
	}

	contains := map[v1alpha1.ComputeNodeConditionType]bool{}
	for _, t := range ts {
		contains[t] = true
	}

	for _, c := range conds {
		if !contains[c.Type] {
			return false
		}
	}
	return true
}
