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

package proxyserver

import (
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const imageName = "apache/shardingsphere-proxy"

const defaultLogback = `<?xml version="1.0"?>
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

// ConstructCascadingConfigmap Construct spec resources to Configmap
func ConstructCascadingConfigmap(proxyConfig *v1alpha1.ShardingSphereProxyServerConfig) *v1.ConfigMap {
	y := toYaml(proxyConfig)
	return &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proxyConfig.Name,
			Namespace: proxyConfig.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(proxyConfig.GetObjectMeta(), proxyConfig.GroupVersionKind()),
			},
		},
		Data: map[string]string{
			"server.yaml": y,
			"logback.xml": defaultLogback,
		},
	}
}

// ToYaml Convert ShardingSphereProxyServerConfig spec content to yaml format
func toYaml(proxyConfig *v1alpha1.ShardingSphereProxyServerConfig) string {
	y, _ := yaml.Marshal(proxyConfig.Spec)
	return string(y)
}

func fromInt32(val int32) intstr.IntOrString {
	return intstr.IntOrString{Type: intstr.Int, IntVal: val}
}
