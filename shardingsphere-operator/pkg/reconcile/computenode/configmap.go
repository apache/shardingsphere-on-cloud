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
	"encoding/json"
	"reflect"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// ConfigDataKeyForLogback refers to the configuration file name of logback
	ConfigDataKeyForLogback = "logback.xml"
	// ConfigDataKeyForServer refers to the configuration file name of server
	ConfigDataKeyForServer = "server.yaml"
	// ConfigDataKeyForAgent refers to the configuration file name of agent
	ConfigDataKeyForAgent = "agent.yaml"

	// AnnoClusterRepoConfig refers to the content of cluster repo config in server config
	AnnoClusterRepoConfig = "computenode.shardingsphere.org/server-config-mode-cluster"
	// AnnoClusterRepoConfig refers to the content of logback.xml
	AnnoLogbackConfig = "computenode.shardingsphere.org/logback"
)

// NewConfigMap returns a new ConfigMap
func NewConfigMap(cn *v1alpha1.ComputeNode) *v1.ConfigMap {
	builder := NewConfigMapBuilder(cn.GetObjectMeta(), cn.GetObjectKind().GroupVersionKind())
	builder.SetName(cn.Name).SetNamespace(cn.Namespace).SetLabels(cn.Labels).SetAnnotations(cn.Annotations)

	cluster := cn.Annotations[AnnoClusterRepoConfig]
	logback := cn.Annotations[AnnoLogbackConfig]

	if len(logback) > 0 {
		builder.SetLogback(logback)
	} else {
		builder.SetLogback(string(DefaultLogback))
	}

	// NOTE: ShardingSphere Proxy 5.3.0 needs a server.yaml no matter if it is empty
	if !reflect.DeepEqual(cn.Spec.Bootstrap.ServerConfig, v1alpha1.ServerConfig{}) {
		servconf := cn.Spec.Bootstrap.ServerConfig.DeepCopy()
		if cn.Spec.Bootstrap.ServerConfig.Mode.Type == v1alpha1.ModeTypeCluster {
			if len(cluster) > 0 {
				json.Unmarshal([]byte(cluster), &servconf.Mode.Repository)
			}
		}
		if y, err := yaml.Marshal(servconf); err == nil {
			builder.SetServerConfig(string(y))
		}
	} else {
		builder.SetServerConfig("# Empty file is needed")
	}

	// load java agent config to configmap if needed
	if !reflect.DeepEqual(cn.Spec.Bootstrap.AgentConfig, v1alpha1.AgentConfig{}) {
		agentConf := cn.Spec.Bootstrap.AgentConfig.DeepCopy()
		if y, err := yaml.Marshal(agentConf); err == nil {
			builder.SetAgentConfig(string(y))
		}
	}

	return builder.Build()
}

// ConfigMapBuilder is a builder for ConfigMap by ComputeNode
type ConfigMapBuilder interface {
	SetName(name string) ConfigMapBuilder
	SetNamespace(namespace string) ConfigMapBuilder
	SetLabels(labels map[string]string) ConfigMapBuilder
	SetAnnotations(annos map[string]string) ConfigMapBuilder
	SetLogback(logback string) ConfigMapBuilder
	SetServerConfig(serverConfig string) ConfigMapBuilder
	SetAgentConfig(agentConfig string) ConfigMapBuilder
	Build() *v1.ConfigMap
}

type configmapBuilder struct {
	configmap *v1.ConfigMap
}

// NewConfigMapBuilder returns a ConfigMapBuilder
func NewConfigMapBuilder(meta metav1.Object, gvk schema.GroupVersionKind) ConfigMapBuilder {
	return &configmapBuilder{
		configmap: DefaultConfigMap(meta, gvk),
	}
}

// SetName set the ConfigMap name
func (c *configmapBuilder) SetName(name string) ConfigMapBuilder {
	c.configmap.Name = name
	return c
}

// SetNamespace set the ConfigMap namespace
func (c *configmapBuilder) SetNamespace(namespace string) ConfigMapBuilder {
	c.configmap.Namespace = namespace
	return c
}

// SetLabels set the ConfigMap labels
func (c *configmapBuilder) SetLabels(labels map[string]string) ConfigMapBuilder {
	c.configmap.Labels = labels
	return c
}

// SetAnnotations set the ConfigMap annotations
func (c *configmapBuilder) SetAnnotations(annos map[string]string) ConfigMapBuilder {
	c.configmap.Annotations = annos
	return c
}

// SetLogback set the ConfigMap data logback
func (c *configmapBuilder) SetLogback(logback string) ConfigMapBuilder {
	c.configmap.Data[ConfigDataKeyForLogback] = logback
	return c
}

// SetServerConfig set the ConfigMap data server config
func (c *configmapBuilder) SetServerConfig(serviceConfig string) ConfigMapBuilder {
	c.configmap.Data[ConfigDataKeyForServer] = serviceConfig
	return c
}

// SetAgentConfig set the ConfigMap data agent config
func (c *configmapBuilder) SetAgentConfig(agentConfig string) ConfigMapBuilder {
	c.configmap.Data[ConfigDataKeyForAgent] = agentConfig
	return c
}

// Build returns a ConfigMap
func (c *configmapBuilder) Build() *v1.ConfigMap {
	return c.configmap
}

// DefaultConfigMap returns a ConfigMap filling with default expected values
func DefaultConfigMap(meta metav1.Object, gvk schema.GroupVersionKind) *v1.ConfigMap {
	return &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "shardingsphere-proxy",
			Namespace: "default",
			Labels:    map[string]string{},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(meta, gvk),
			},
		},
		Data: map[string]string{},
	}
}

// FIXME: check if changed first, then decide if need to respawn the Pods
func UpdateConfigMap(cn *v1alpha1.ComputeNode, cur *v1.ConfigMap) *v1.ConfigMap {
	exp := &v1.ConfigMap{}
	exp.ObjectMeta = cur.ObjectMeta
	exp.ObjectMeta.ResourceVersion = ""
	exp.Labels = cur.Labels
	exp.Annotations = cur.Annotations
	exp.Data = NewConfigMap(cn).Data
	return exp
}

// DefaultLogback contains the default logback config
const DefaultLogback = `<?xml version="1.0"?>
<configuration>
    <appender name="console" class="ch.qos.logback.core.ConsoleAppender">
        <encoder>
            <pattern>[%-5level] %d{yyyy-MM-dd HH:mm:ss.SSS} [%thread] %logger{36} - %msg%n</pattern>
        </encoder>
    </appender>
    <appender name="sqlConsole" class="ch.qos.logback.core.ConsoleAppender">
        <encoder>
            <pattern>[%-5level] %d{yyyy-MM-dd HH:mm:ss.SSS} [%thread] [%X{database}] [%X{user}] [%X{host}] %logger{36} - %msg%n</pattern>
        </encoder>
    </appender>
    
    <logger name="ShardingSphere-SQL" level="info" additivity="false">
        <appender-ref ref="sqlConsole" />
    </logger>
    <logger name="org.apache.shardingsphere" level="info" additivity="false">
        <appender-ref ref="console" />
    </logger>
    
    <logger name="com.zaxxer.hikari" level="error" />
    
    <logger name="com.atomikos" level="error" />
    
    <logger name="io.netty" level="error" />
    
    <root>
        <level value="info" />
        <appender-ref ref="console" />
    </root>
</configuration> 
`
