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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("ConfigMap", func() {
	var (
		expect = &corev1.ConfigMap{}
		cn     = &v1alpha1.ComputeNode{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test_name",
				Namespace: "test_namespace",
				Labels: map[string]string{
					"test_key": "test_value",
				},
				Annotations: map[string]string{
					computenode.AnnoLogbackConfig:     "test_logback",
					computenode.AnnoClusterRepoConfig: "test_cluster_repo_config",
				},
			},
		}
	)

	BeforeEach(func() {
		expect.Name = "test_name"
		expect.Namespace = "test_namespace"
		expect.Labels = map[string]string{
			"test_key": "test_value",
		}
		expect.Data = map[string]string{}
		expect.Data[computenode.ConfigDataKeyForLogback] = "test_logback"
		expect.Data[computenode.ConfigDataKeyForServer] = "test_cluster_repo_config"
	})

	Context("Assert ObjectMeta", func() {
		cm := computenode.NewConfigMap(cn)
		It("name should be equal", func() {
			Expect(expect.Name).To(Equal(cm.Name))
		})
		It("namespace should be equal", func() {
			Expect(expect.Namespace).To(Equal(cm.Namespace))
		})
		It("labels should be equal", func() {
			Expect(expect.Labels).To(Equal(cm.Labels))
		})
	})

	Context("Assert Default Spec Data", func() {
		cm := computenode.NewConfigMap(cn)
		It("default logback should be equal", func() {
			Expect(expect.Data[computenode.AnnoLogbackConfig]).To(Equal(cm.Data[computenode.AnnoLogbackConfig]))
		})
		It("default cluster repo config should be equal", func() {
			Expect(expect.Data[computenode.AnnoClusterRepoConfig]).To(Equal(cm.Data[computenode.AnnoClusterRepoConfig]))
		})
	})
})
