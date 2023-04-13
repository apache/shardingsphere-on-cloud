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

package awsaurora

import (
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Awsaurora", func() {
	Context("IsValid", func() {
		node := &v1alpha1.StorageNode{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "test-storage-node",
				Namespace:   "test-namespace",
				Annotations: map[string]string{},
			},
			Spec: v1alpha1.StorageNodeSpec{
				DatabaseClassName: "aws-aurora",
			},
		}

		It("username empty", func() {
			aurora := AwsAurora{}
			err := aurora.IsValid(node)
			Expect(err).To(HaveOccurred())
			// username
			Expect(err.Error()).To(ContainSubstring("username is required"))
			// password
			Expect(err.Error()).To(ContainSubstring("password is required"))
			// vpcSecurityGroupIds
			Expect(err.Error()).To(ContainSubstring("vpcSecurityGroupIds is required"))
			// subnetGroupName
			Expect(err.Error()).To(ContainSubstring("subnetGroupName is required"))
			// dbClusterIdentifier
			Expect(err.Error()).To(ContainSubstring("dbClusterIdentifier is required"))
		})
		It("password empty", func() {
			aurora := AwsAurora{}
			node.Annotations[AnnoUsername] = "username"
			err := aurora.IsValid(node)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("password is required"))
			// vpcSecurityGroupIds
			Expect(err.Error()).To(ContainSubstring("vpcSecurityGroupIds is required"))
			// subnetGroupName
			Expect(err.Error()).To(ContainSubstring("subnetGroupName is required"))
			// dbClusterIdentifier
			Expect(err.Error()).To(ContainSubstring("dbClusterIdentifier is required"))
		})
		It("vpcSecurityGroupIds empty", func() {
			aurora := AwsAurora{}
			node.Annotations[AnnoUsername] = "username"
			node.Annotations[AnnoPassword] = "password"
			err := aurora.IsValid(node)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("vpcSecurityGroupIds is required"))
			// subnetGroupName
			Expect(err.Error()).To(ContainSubstring("subnetGroupName is required"))
			// dbClusterIdentifier
			Expect(err.Error()).To(ContainSubstring("dbClusterIdentifier is required"))
		})
		It("subnetGroupName empty", func() {
			aurora := AwsAurora{}
			node.Annotations[AnnoUsername] = "username"
			node.Annotations[AnnoPassword] = "password"
			node.Annotations[AnnoVpcSecurityGroupIds] = "vpcSecurityGroupIds"
			err := aurora.IsValid(node)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("subnetGroupName is required"))
			// dbClusterIdentifier
			Expect(err.Error()).To(ContainSubstring("dbClusterIdentifier is required"))
		})
		It("dbClusterIdentifier empty", func() {
			aurora := AwsAurora{}
			node.Annotations[AnnoUsername] = "username"
			node.Annotations[AnnoPassword] = "password"
			node.Annotations[AnnoVpcSecurityGroupIds] = "vpcSecurityGroupIds"
			node.Annotations[AnnoSubnetGroupName] = "subnetGroupName"
			err := aurora.IsValid(node)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("dbClusterIdentifier is required"))
		})
		It("all required fields are set", func() {
			aurora := AwsAurora{}
			node.Annotations[AnnoUsername] = "username"
			node.Annotations[AnnoPassword] = "password"
			node.Annotations[AnnoVpcSecurityGroupIds] = "vpcSecurityGroupIds"
			node.Annotations[AnnoSubnetGroupName] = "subnetGroupName"
			node.Annotations[AnnoDBClusterIdentifier] = "dbClusterIdentifier"
			err := aurora.IsValid(node)
			Expect(err).ToNot(HaveOccurred())
		})

	})
})
