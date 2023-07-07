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

const (
	defaultImageName             = "apache/shardingsphere-proxy"
	defaultImage                 = "apache/shardingsphere-proxy:5.3.0"
	defaultContainerName         = "shardingsphere-proxy"
	defaultConfigVolumeName      = "shardingsphere-proxy-config"
	defaultConfigVolumeMountPath = "/opt/shardingsphere-proxy/conf"

	defaultExtlibPath            = "/opt/shardingsphere-proxy/ext-lib"
	defaultMySQLDriverEnvName    = "MYSQL_CONNECTOR_VERSION"
	defaultMySQLDriverVolumeName = "mysql-connector-java"

	defaultJavaAgentVolumeName            = "java-agent-bin"
	defaultJavaAgentVolumeMountPath       = "/opt/shardingsphere-proxy/agent"
	defaultJavaAgentConfigVolumeName      = "java-agent-config"
	defaultJavaAgentConfigVolumeMountPath = "/opt/shardingsphere-proxy/agent/conf"
	defaultJavaToolOptionsName            = "JAVA_TOOL_OPTIONS"
	defaultJavaAgentEnvValue              = "-javaagent:/opt/shardingsphere-proxy/agent/shardingsphere-agent-%s.jar"
	defaultAgentBinVersionEnvName         = "AGENT_BIN_VERSION"
)

const (
	DefaultAnnotationJavaAgentEnabled       = "shardingsphere.apache.org/java-agent-enabled"
	commonAnnotationPrometheusMetricsPath   = "prometheus.io/path"
	commonAnnotationPrometheusMetricsPort   = "prometheus.io/port"
	commonAnnotationPrometheusMetricsScrape = "prometheus.io/scrape"
	commonAnnotationPrometheusMetricsScheme = "prometheus.io/scheme"
)

const (
	downloadMysqlJarScript = `wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/${MYSQL_CONNECTOR_VERSION}/mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar;
 wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/${MYSQL_CONNECTOR_VERSION}/mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar.md5;
 if [ $(md5sum /mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar | cut -d ' ' -f1) = $(cat /mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar.md5) ];
 then echo success;
 else echo failed;exit 1;fi;mv /mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar /opt/shardingsphere-proxy/ext-lib`
	downloadAgentJarScript = `wget https://archive.apache.org/dist/shardingsphere/${AGENT_BIN_VERSION}/apache-shardingsphere-${AGENT_BIN_VERSION}-shardingsphere-agent-bin.tar.gz;
 tar -zxvf apache-shardingsphere-${AGENT_BIN_VERSION}-shardingsphere-agent-bin.tar.gz -C /opt/shardingsphere-proxy/agent --strip-component 1;`
	replaceStartScript = `sed -i 's#exec \$JAVA \${JAVA_OPTS} \${JAVA_MEM_OPTS} -classpath \${CLASS_PATH} \${MAIN_CLASS}#exec \$JAVA \${JAVA_OPTS} \${JAVA_MEM_OPTS} -classpath \${CLASS_PATH} \${AGENT_PARAM} \${MAIN_CLASS}#g' /opt/shardingsphere-proxy/bin/start.sh;
	cp /opt/shardingsphere-proxy/bin/start.sh /opt/shardingsphere-proxy/tmpbin/start.sh;`
)
