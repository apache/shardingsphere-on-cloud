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
	"fmt"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/container"
	corev1 "k8s.io/api/core/v1"
)

// ShardingSphereProxyContainerBuilder contains a common container builder
// and several different Proxy related attributes
type ShardingSphereProxyContainerBuilder interface {
	// A default container builder
	container.ContainerBuilder

	// set the version of ShardingSphere Proxy
	SetVersion(version string) ShardingSphereProxyContainerBuilder
}

type shardingSphereProxyContainerBuilder struct {
	container.ContainerBuilder
}

// SetVersion sets the version of ShardingSphere Proxy
func (c *shardingSphereProxyContainerBuilder) SetVersion(version string) ShardingSphereProxyContainerBuilder {
	c.SetImage(fmt.Sprintf("%s:%s", defaultImageName, version))
	return c
}

// NewShardingSphereProxyContainerBuilder returns a builder for ShardingSphereContainer
// This will set default container name
func NewShardingSphereProxyContainerBuilder() ShardingSphereProxyContainerBuilder {
	return &shardingSphereProxyContainerBuilder{
		ContainerBuilder: container.NewContainerBuilder().
			SetName(defaultContainerName),
	}
}

// BootstrapContainerBuilder returns a Container for initialization
// The container will handle initilialization in Pod's InitContainer
type BootstrapContainerBuilder interface {
	container.ContainerBuilder
}

type bootstrapContainerBuilder struct {
	container.ContainerBuilder
}

// NewBootstrapContainerBuilderForMysqlJar will return a builder for MysqlJar download container
// This will set the default container name, image and commands
func NewBootstrapContainerBuilderForMysqlJar() BootstrapContainerBuilder {
	return &bootstrapContainerBuilder{
		ContainerBuilder: container.NewContainerBuilder().
			SetName("download-mysql-jar").
			SetImage("busybox:1.36").
			SetCommand([]string{"/bin/sh", "-c", downloadMysqlJarScript}),
	}
}

// NewBootstrapContainerBuilderForAgentBin will return a builder for ShardingSphere-Agent bin jar download container
// This will set the default container name, image and commands
func NewBootstrapContainerBuilderForAgentBin() BootstrapContainerBuilder {
	return &bootstrapContainerBuilder{
		ContainerBuilder: container.NewContainerBuilder().
			SetName("download-agent-bin-jar").
			SetImage("busybox:1.36").
			SetCommand([]string{"/bin/sh", "-c", downloadAgentJarScript}),
	}
}

// NewBootstrapContainerBuilderForStartScript will return a builder for ShardingSphere-Proxy modify container start.sh
func NewBootstrapContainerBuilderForStartScripts() BootstrapContainerBuilder {
	return &bootstrapContainerBuilder{
		ContainerBuilder: container.NewContainerBuilder().
			SetName("replace-start-script").
			SetImage(fmt.Sprintf("%s:%s", defaultImageName, "5.3.2")).
			SetCommand([]string{"/bin/sh", "-c", replaceStartScript}),
	}
}

// Build returns a Container
func (b *bootstrapContainerBuilder) Build() *corev1.Container {
	return b.ContainerBuilder.BuildContainer()
}
