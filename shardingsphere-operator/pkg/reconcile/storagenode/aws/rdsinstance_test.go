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

package aws

import (
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// test validCreateInstanceParams with ginkgo and gomega.
var _ = Describe("validCreateInstanceParams", func() {
	Context("validCreateInstanceParams", func() {
		node := &v1alpha1.StorageNode{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "test-instance",
				Namespace:   "test-namespace",
				Annotations: map[string]string{},
			},
		}
		params := map[string]string{
			"instanceClass":      "db.t3.micro",
			"engine":             "mysql",
			"engineVersion":      "5.7",
			"masterUsername":     "root",
			"masterUserPassword": "root123456",
			"allocatedStorage":   "20",
		}

		It("should return true", func() {

			node.Annotations[v1alpha1.AnnotationsInstanceIdentifier] = "test-instance"
			Expect(validCreateInstanceParams(node, &params)).To(BeNil())
			Expect(node.Annotations[v1alpha1.AnnotationsMasterUserPassword]).To(Equal("root123456"))
		})
		It("should return username contains invalid characters", func() {
			params["masterUsername"] = "@masterUser"
			node.Annotations[v1alpha1.AnnotationsInstanceIdentifier] = "test-instance"
			Expect(validCreateInstanceParams(node, &params)).To(MatchError(ContainSubstring("username contains invalid characters")))
		})
		It("should handle multiple characters correctly", func() {
			params["masterUsername"] = "test__test--"
			node.Annotations[v1alpha1.AnnotationsInstanceIdentifier] = "test-instance"
			Expect(validCreateInstanceParams(node, &params)).To(BeNil())
			Expect(params["masterUsername"]).To(Equal("test_test-"))
		})
	})
})
