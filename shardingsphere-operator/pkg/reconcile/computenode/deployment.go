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
	"context"
	"fmt"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/container"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/deployment"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/metadata"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	defaultExtlibPath            = "/opt/shardingsphere-proxy/ext-lib"
	defaultImageName             = "apache/shardingsphere-proxy"
	defaultImage                 = "apache/shardingsphere-proxy:5.3.0"
	defaultContainerName         = "shardingsphere-proxy"
	defaultConfigVolumeName      = "shardingsphere-proxy-config"
	defaultConfigVolumeMountPath = "/opt/shardingsphere-proxy/conf"
	defaultMySQLDriverEnvName    = "MYSQL_CONNECTOR_VERSION"
	defaultMySQLDriverVolumeName = "mysql-connector-java"

	DefaultAnnotationJavaAgentEnabled       = "shardingsphere.apache.org/java-agent-enabled"
	commonAnnotationPrometheusMetricsPath   = "prometheus.io/path"
	commonAnnotationPrometheusMetricsPort   = "prometheus.io/port"
	commonAnnotationPrometheusMetricsScrape = "prometheus.io/scrape"
	commonAnnotationPrometheusMetricsScheme = "prometheus.io/scheme"

	defaultJavaAgentVolumeName            = "java-agent-bin"
	defaultJavaAgentVolumeMountPath       = "/opt/shardingsphere-proxy/agent"
	defaultJavaAgentConfigVolumeName      = "java-agent-config"
	defaultJavaAgentConfigVolumeMountPath = "/opt/shardingsphere-proxy/agent/conf"
	defaultJavaToolOptionsName            = "JAVA_TOOL_OPTIONS"
	defaultJavaAgentEnvValue              = "-javaagent:/opt/shardingsphere-proxy/agent/shardingsphere-agent-%s.jar"
	defaultAgentBinVersionEnvName         = "AGENT_BIN_VERSION"

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

// Builder build Deployment from given ComputeNode
type Builder interface {
	Build(context.Context, *v1alpha1.ComputeNode) *appsv1.Deployment
}

func NewBuilder() Builder {
	return &builder{}
}

type builder struct{}

func (b builder) buildProbes(scb container.ContainerBuilder, cn *v1alpha1.ComputeNode) {
	if cn.Spec.Probes == nil {
		return
	}

	if cn.Spec.Probes.LivenessProbe != nil {
		scb.SetLivenessProbe(cn.Spec.Probes.LivenessProbe)
	}
	if cn.Spec.Probes.ReadinessProbe != nil {
		scb.SetReadinessProbe(cn.Spec.Probes.ReadinessProbe)
	}
	if cn.Spec.Probes.StartupProbe != nil {
		scb.SetStartupProbe(cn.Spec.Probes.StartupProbe)
	}
}

func (b builder) buildMetadata(ssbuilder ShardingSphereDeploymentBuilder, cn *v1alpha1.ComputeNode) {
	ssbuilder.SetName(cn.Name).
		SetNamespace(cn.Namespace).
		SetLabels(cn.Labels).
		SetAnnotations(cn.Annotations)
}

func getContainerPortsFromComputeNode(cn *v1alpha1.ComputeNode) []corev1.ContainerPort {
	ports := []corev1.ContainerPort{}
	for idx := range cn.Spec.PortBindings {
		ports = append(ports, corev1.ContainerPort{
			Name:          cn.Spec.PortBindings[idx].Name,
			HostIP:        cn.Spec.PortBindings[idx].HostIP,
			ContainerPort: cn.Spec.PortBindings[idx].ContainerPort,
			Protocol:      cn.Spec.PortBindings[idx].Protocol,
		})
	}
	return ports
}

func (b builder) buildSpec(ssbuilder ShardingSphereDeploymentBuilder, cn *v1alpha1.ComputeNode) {
	ssbuilder.SetSelectors(cn.Spec.Selector)
	ssbuilder.SetReplicas(&cn.Spec.Replicas)
	ssbuilder.SetRollingUpdateStrategy(0, 3)

	tpl := &corev1.PodTemplateSpec{}
	tm := metadata.NewMetadataBuilder()
	tm.SetLabels(cn.Labels)

	ports := getContainerPortsFromComputeNode(cn)

	scb := NewShardingSphereProxyContainerBuilder().
		SetVersion(cn.Spec.ServerVersion).
		SetPorts(ports).
		SetResources(cn.Spec.Resources)

	b.buildProbes(scb, cn)

	vcb := deployment.NewSharedVolumeAndMountBuilder().
		SetVolumeMountSize(1).
		SetName(defaultConfigVolumeName).
		SetVolumeSourceConfigMap(cn.Name).
		SetMountPath(0, defaultConfigVolumeMountPath)
	vc, vmc := vcb.Build()

	ssbuilder.AppendVolumes([]corev1.Volume{
		*vc,
	})
	scb.AppendVolumeMounts([]corev1.VolumeMount{*vmc[0]})

	sc := scb.BuildContainer()
	ssbuilder.AppendContainers([]corev1.Container{
		*sc,
	})

	if enabled, ok := cn.Annotations[DefaultAnnotationJavaAgentEnabled]; ok && enabled == "true" {
		ssbuilder.SetAgentBin(cn)
	}

	if cn.Spec.StorageNodeConnector != nil {
		if cn.Spec.StorageNodeConnector.Type == v1alpha1.ConnectorTypeMySQL {
			ssbuilder.SetMySQLConnector(cn)
		}
	}

	tpl.ObjectMeta = *tm.BuildMetadata()
	tpl.Spec = *ssbuilder.BuildPodSpec()
	ssbuilder.SetPodTemplateSpec(tpl)
}

// Build returns a new Deployment
func (b builder) Build(ctx context.Context, cn *v1alpha1.ComputeNode) *appsv1.Deployment {
	ssbuilder := NewShardingSphereDeploymentBuilder(cn.GetObjectMeta(), cn.GetObjectKind().GroupVersionKind())

	b.buildMetadata(ssbuilder, cn)
	b.buildSpec(ssbuilder, cn)

	return ssbuilder.BuildShardingSphereDeployment()
}

type ShardingSphereDeploymentBuilder interface {
	deployment.DeploymentBuilder

	SetMySQLConnector(cn *v1alpha1.ComputeNode) ShardingSphereDeploymentBuilder
	SetAgentBin(cn *v1alpha1.ComputeNode) ShardingSphereDeploymentBuilder
	SetAgentScript(cn *v1alpha1.ComputeNode) ShardingSphereDeploymentBuilder

	BuildShardingSphereDeployment() *appsv1.Deployment
}

// NewShardingSphereDeploymentBuilder creates a new ShardingSphereDeploymentBuilder
func NewShardingSphereDeploymentBuilder(meta metav1.Object, gvk schema.GroupVersionKind) ShardingSphereDeploymentBuilder {
	// dp := deployment.DefaultDeployment(meta, gvk)
	db := deployment.NewDeploymentBuilder()
	dp := db.BuildDeployment()

	return &shardingsphereDeploymentBuilder{
		DeploymentBuilder: db,
		deployment:        dp,
	}
}

type shardingsphereDeploymentBuilder struct {
	deployment.DeploymentBuilder
	deployment *appsv1.Deployment
}

// SetMySQLConnector will set an init container to download mysql jar and mount files for proxy container.
func (d *shardingsphereDeploymentBuilder) SetMySQLConnector(cn *v1alpha1.ComputeNode) ShardingSphereDeploymentBuilder {
	proxy := d.FindContainerByName("shardingsphere-proxy")
	proxy.AppendEnv([]corev1.EnvVar{
		{
			Name:  defaultMySQLDriverEnvName,
			Value: cn.Spec.StorageNodeConnector.Version,
		},
	})

	cb := d.FindInitContainerByName("download-mysql-jar")
	if cb == nil {
		cb = NewBootstrapContainerBuilderForMysqlJar()
	}

	cb.AppendEnv([]corev1.EnvVar{
		{
			Name:  defaultMySQLDriverEnvName,
			Value: cn.Spec.StorageNodeConnector.Version,
		},
	})

	vb := deployment.NewSharedVolumeAndMountBuilder().
		SetVolumeMountSize(2).
		SetName(defaultMySQLDriverVolumeName).
		SetVolumeSourceEmptyDir().
		SetMountPath(0, defaultExtlibPath).
		SetMountPath(1, absoluteMySQLDriverMountName(defaultExtlibPath, cn.Spec.StorageNodeConnector.Version)).
		SetSubPath(1, relativeMySQLDriverMountName(cn.Spec.StorageNodeConnector.Version))
	v, vms := vb.Build()

	cb.AppendVolumeMounts([]corev1.VolumeMount{*vms[0]})
	proxy.AppendVolumeMounts([]corev1.VolumeMount{*vms[1]})

	d.UpdateContainerByName(proxy.BuildContainer())
	d.UpdateInitContainerByName(cb.BuildContainer())
	d.AppendVolumes([]corev1.Volume{*v})

	return d
}

// SetAgentBin set `agent bin` for ShardingSphereProxy with [observability](https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-proxy/observability/)
func (d *shardingsphereDeploymentBuilder) SetAgentBin(cn *v1alpha1.ComputeNode) ShardingSphereDeploymentBuilder {
	// set env JAVA_TOOL_OPTIONS to proxy container, make sure proxy will apply agent-bin.jar
	// agent-bin's version is always equals to shardingsphere proxy image's version

	metricsAnnos := map[string]string{}
	metricsAnnos[commonAnnotationPrometheusMetricsPath] = cn.Annotations[commonAnnotationPrometheusMetricsPath]
	metricsAnnos[commonAnnotationPrometheusMetricsPort] = cn.Annotations[commonAnnotationPrometheusMetricsPort]
	metricsAnnos[commonAnnotationPrometheusMetricsScrape] = cn.Annotations[commonAnnotationPrometheusMetricsScrape]
	metricsAnnos[commonAnnotationPrometheusMetricsScheme] = cn.Annotations[commonAnnotationPrometheusMetricsScheme]

	if d.deployment.Spec.Template.Annotations == nil {
		d.deployment.Spec.Template.Annotations = map[string]string{}
	}
	d.deployment.Spec.Template.Annotations = metricsAnnos

	proxy := d.FindContainerByName("shardingsphere-proxy")
	proxy.AppendEnv([]corev1.EnvVar{
		{
			Name:  defaultJavaToolOptionsName,
			Value: fmt.Sprintf(defaultJavaAgentEnvValue, cn.Spec.ServerVersion),
		},
	})

	vbAgentConf := deployment.NewSharedVolumeAndMountBuilder().
		SetVolumeMountSize(1).
		SetName(defaultJavaAgentConfigVolumeName).
		SetVolumeSourceConfigMap(cn.Name, corev1.KeyToPath{Key: configmap.ConfigDataKeyForAgent, Path: configmap.ConfigDataKeyForAgent}).
		SetMountPath(0, defaultJavaAgentConfigVolumeMountPath)
	vc, vmc := vbAgentConf.Build()

	vbAgent := deployment.NewSharedVolumeAndMountBuilder().
		SetVolumeMountSize(1).
		SetName(defaultJavaAgentVolumeName).
		SetVolumeSourceEmptyDir().
		SetMountPath(0, defaultJavaAgentVolumeMountPath)
	va, vma := vbAgent.Build()

	d.AppendVolumes([]corev1.Volume{*vc, *va})

	cb := d.FindInitContainerByName("download-agent-bin-jar")
	if cb == nil {
		cb = NewBootstrapContainerBuilderForAgentBin()
	}
	cb.AppendVolumeMounts([]corev1.VolumeMount{*vma[0]}).
		AppendEnv([]corev1.EnvVar{
			{
				Name:  defaultAgentBinVersionEnvName,
				Value: cn.Spec.ServerVersion,
			},
		})

	d.UpdateInitContainerByName(cb.BuildContainer())

	proxy.AppendVolumeMounts([]corev1.VolumeMount{*vmc[0], *vma[0]})

	if cn.Spec.ServerVersion == "5.3.2" {
		d.SetAgentScript(cn)
	}
	d.UpdateContainerByName(proxy.BuildContainer())
	return d
}

func (d *shardingsphereDeploymentBuilder) SetAgentScript(cn *v1alpha1.ComputeNode) ShardingSphereDeploymentBuilder {
	proxy := d.FindContainerByName("shardingsphere-proxy")

	sv := deployment.NewSharedVolumeAndMountBuilder().
		SetVolumeMountSize(1).
		SetName("replace-start-script").
		SetVolumeSourceEmptyDir().
		SetMountPath(0, "/opt/shardingsphere-proxy/bin")
	va, vma := sv.Build()
	d.AppendVolumes([]corev1.Volume{
		*va,
	})

	proxy.AppendVolumeMounts([]corev1.VolumeMount{*vma[0]})

	// NOTE: This mountpath is not same with init container
	vma[0].MountPath = "/opt/shardingsphere-proxy/tmpbin"

	cb := d.FindInitContainerByName("replace-start-script")
	if cb == nil {
		cb = NewBootstrapContainerBuilderForStartScripts()
	}
	cb.AppendVolumeMounts([]corev1.VolumeMount{*vma[0]})

	d.UpdateInitContainerByName(cb.BuildContainer())
	return d
}

func (d *shardingsphereDeploymentBuilder) BuildShardingSphereDeployment() *appsv1.Deployment {
	dp := d.DeploymentBuilder.BuildDeployment()

	return dp
}
