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

package configmap

import (
	"reflect"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"

	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// ConfigDataKeyForLogback refers to the configuration file name of logback
	ConfigDataKeyForLogback = "logback.xml"
	// ConfigDataKeyForServer refers to the configuration file name of server
	ConfigDataKeyForServer = "server.yaml"
	// ConfigDataKeyForAgent refers to the configuration file name of agent
	ConfigDataKeyForAgent = "agent.yaml"

	// AnnoClusterRepoConfig refers to the content of logback.xml
	AnnoLogbackConfig = "computenode.shardingsphere.org/logback"

	// DefaultLogback contains the default logback config
	DefaultLogback = `<?xml version="1.0"?>
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
	// DefaultServerConfig contains the default server config
	DefaultServerConfig = "# Empty file is needed"
)

// DefaultConfigMap returns a ConfigMap filling with default expected values
func DefaultConfigMap(meta metav1.Object, gvk schema.GroupVersionKind) *corev1.ConfigMap {
	return &corev1.ConfigMap{
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

// NewConfigMap returns a new ConfigMap
func NewConfigMap(obj runtime.Object) *corev1.ConfigMap {
	factory := NewConfigMapFactory(obj)
	gvk := obj.GetObjectKind().GroupVersionKind()

	var (
		cn   *v1alpha1.ComputeNode
		sc   *v1alpha1.Chaos
		meta metav1.Object
		ok   bool
	)

	if cn, ok = obj.(*v1alpha1.ComputeNode); ok {
		meta = cn.GetObjectMeta()
		return factory.NewConfigMapBuilder(meta, gvk).Build()
	}

	if sc, ok = obj.(*v1alpha1.Chaos); ok {
		meta = sc.GetObjectMeta()
		return factory.NewConfigMapBuilder(meta, gvk).Build()
	}

	return &corev1.ConfigMap{}
}

// UpdateConfigMap returns a new ConfigMap
func UpdateComputeNodeConfigMap(cn *v1alpha1.ComputeNode, cur *corev1.ConfigMap) *corev1.ConfigMap {
	exp := &corev1.ConfigMap{}
	exp.ObjectMeta = cur.ObjectMeta
	exp.Labels = cur.Labels
	exp.Annotations = cur.Annotations
	exp.Data = NewConfigMap(cn).Data
	return exp
}

type computeNodeConfigMapBuilder struct {
	configMapBuilder
	obj runtime.Object
}

// SetData sets the data of ConfigMap
func (c *computeNodeConfigMapBuilder) SetData(data map[string]string) ConfigMapBuilder {
	if c.configmap.Data == nil {
		c.configmap.Data = map[string]string{}
	}

	if val, ok := data[ConfigDataKeyForServer]; ok {
		c.configmap.Data[ConfigDataKeyForServer] = val
	}

	if val, ok := data[ConfigDataKeyForLogback]; ok {
		c.configmap.Data[ConfigDataKeyForLogback] = val
	}

	if val, ok := data[ConfigDataKeyForAgent]; ok {
		c.configmap.Data[ConfigDataKeyForAgent] = val
	}

	return c
}

// SetBinaryData sets the binary data of ConfigMap
func (c *computeNodeConfigMapBuilder) SetBinaryData(binary map[string][]byte) ConfigMapBuilder {
	if c.configmap.BinaryData == nil {
		c.configmap.BinaryData = map[string][]byte{}
	}
	if val, ok := binary[ConfigDataKeyForServer]; ok {
		c.configmap.BinaryData[ConfigDataKeyForServer] = val
	}

	if val, ok := binary[ConfigDataKeyForLogback]; ok {
		c.configmap.BinaryData[ConfigDataKeyForLogback] = val
	}

	if val, ok := binary[ConfigDataKeyForAgent]; ok {
		c.configmap.BinaryData[ConfigDataKeyForAgent] = val
	}

	return c
}

// Build builds the ConfigMap
func (c *computeNodeConfigMapBuilder) Build() *corev1.ConfigMap {
	var (
		cn *v1alpha1.ComputeNode
		ok bool
	)
	if cn, ok = c.obj.(*v1alpha1.ComputeNode); !ok {
		return nil
	}
	c.SetName(cn.Name).SetNamespace(cn.Namespace).SetLabels(cn.Labels).SetAnnotations(cn.Annotations)

	data := map[string]string{}

	logback := cn.Annotations[AnnoLogbackConfig]
	if len(logback) > 0 {
		data[ConfigDataKeyForLogback] = logback
	} else {
		data[ConfigDataKeyForLogback] = DefaultLogback
	}

	// NOTE: ShardingSphere Proxy 5.3.0 needs a server.yaml no matter if it is empty
	if !reflect.DeepEqual(cn.Spec.Bootstrap.ServerConfig, v1alpha1.ServerConfig{}) {
		servconf := cn.Spec.Bootstrap.ServerConfig.DeepCopy()
		if y, err := yaml.Marshal(servconf); err == nil {
			data[ConfigDataKeyForServer] = string(y)
		}
	} else {
		data[ConfigDataKeyForServer] = DefaultServerConfig
	}

	// load java agent config to configmap if needed
	if !reflect.DeepEqual(cn.Spec.Bootstrap.AgentConfig, v1alpha1.AgentConfig{}) {
		agentConf := cn.Spec.Bootstrap.AgentConfig.DeepCopy()
		if y, err := yaml.Marshal(agentConf); err == nil {
			data[ConfigDataKeyForAgent] = string(y)
		}
	}

	c.SetData(data)

	return c.configmap
}

type shardingsphereChaosConfigMapBuilder struct {
	configMapBuilder
	obj runtime.Object
}

// SetData sets the data of ConfigMap
func (c *shardingsphereChaosConfigMapBuilder) SetData(data map[string]string) ConfigMapBuilder {
	if c.configmap.Data == nil {
		c.configmap.Data = map[string]string{}
	}
	if val, ok := data[configExperimental]; ok {
		c.configmap.Data[configExperimental] = val
	}

	if val, ok := data[configVerify]; ok {
		c.configmap.Data[configVerify] = val
	}

	if val, ok := data[configPressure]; ok {
		c.configmap.Data[configPressure] = val
	}

	return c
}

// SetBinaryData sets the binary data of ConfigMap
func (c *shardingsphereChaosConfigMapBuilder) SetBinaryData(binary map[string][]byte) ConfigMapBuilder {
	if c.configmap.BinaryData == nil {
		c.configmap.BinaryData = map[string][]byte{}
	}
	if val, ok := binary[configExperimental]; ok {
		c.configmap.BinaryData[configExperimental] = val
	}

	if val, ok := binary[configVerify]; ok {
		c.configmap.BinaryData[configVerify] = val
	}

	if val, ok := binary[configPressure]; ok {
		c.configmap.BinaryData[configPressure] = val
	}

	return c
}

// Build builds the ConfigMap
func (c *shardingsphereChaosConfigMapBuilder) Build() *corev1.ConfigMap {
	var (
		chaos *v1alpha1.Chaos
		ok    bool
	)
	if chaos, ok = c.obj.(*v1alpha1.Chaos); !ok {
		return nil
	}
	c.SetName(chaos.Name).SetNamespace(chaos.Namespace).SetLabels(chaos.Labels).SetAnnotations(chaos.Annotations)

	data := map[string]string{}
	data[configExperimental] = string(chaos.Spec.InjectJob.Experimental)
	data[configPressure] = string(chaos.Spec.InjectJob.Pressure)
	data[configVerify] = string(chaos.Spec.InjectJob.Verify)

	c.SetData(data)

	return c.configmap
}

const (
	configExperimental = "experimental.sh"
	configPressure     = "pressure.sh"
	configVerify       = "verify.sh"
)

const (
	// DefaultConfigMapName is the data key name
	DefaultConfigMapName = "ssChaos-configmap"
)

// UpdateConfigMap returns a new ConfigMap
func UpdateShardingSphereChaosConfigMap(ssChaos *v1alpha1.Chaos, cur *corev1.ConfigMap) *corev1.ConfigMap {
	exp := &corev1.ConfigMap{}
	exp.ObjectMeta = cur.ObjectMeta
	exp.Labels = cur.Labels
	exp.Annotations = cur.Annotations
	now := NewConfigMap(ssChaos)
	exp.Data = now.Data
	return exp
}
