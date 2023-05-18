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

package deployment

import (
	"fmt"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/common"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
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
)

func relativeMySQLDriverMountName(v string) string {
	return fmt.Sprintf("mysql-connector-java-%s.jar", v)
}

func absoluteMySQLDriverMountName(p, v string) string {
	return fmt.Sprintf("%s/%s", p, relativeMySQLDriverMountName(v))
}

// ShardingSphereProxyContainerBuilder contains a common container builder
// and several different Proxy related attributes
type ShardingSphereProxyContainerBuilder interface {
	// A default container builder
	common.ContainerBuilder

	// set the version of ShardingSphere Proxy
	SetVersion(version string) ShardingSphereProxyContainerBuilder
}

type shardingSphereProxyContainerBuilder struct {
	common.ContainerBuilder
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
		ContainerBuilder: common.NewContainerBuilder().
			SetName(defaultContainerName),
	}
}

// BootstrapContainerBuilder returns a Container for initialization
// The container will handle initilialization in Pod's InitContainer
type BootstrapContainerBuilder interface {
	common.ContainerBuilder
}

type bootstrapContainerBuilder struct {
	common.ContainerBuilder
}

// NewBootstrapContainerBuilderForMysqlJar will return a builder for MysqlJar download container
// This will set the default container name, image and commands
func NewBootstrapContainerBuilderForMysqlJar() BootstrapContainerBuilder {
	return &bootstrapContainerBuilder{
		ContainerBuilder: common.NewContainerBuilder().
			SetName("download-mysql-jar").
			SetImage("busybox:1.35.0").
			SetCommand([]string{"/bin/sh", "-c", downloadMysqlJarScript}),
	}
}

// NewBootstrapContainerBuilderForAgentBin will return a builder for ShardingSphere-Agent bin jar download container
// This will set the default container name, image and commands
func NewBootstrapContainerBuilderForAgentBin() BootstrapContainerBuilder {
	return &bootstrapContainerBuilder{
		ContainerBuilder: common.NewContainerBuilder().
			SetName("download-agent-bin-jar").
			SetImage("busybox:1.35.0").
			SetCommand([]string{"/bin/sh", "-c", downloadAgentJarScript}),
	}
}

// Build returns a Container
func (b *bootstrapContainerBuilder) Build() *corev1.Container {
	return b.ContainerBuilder.Build()
}

// DeploymentBuilder returns a deployment builder
type DeploymentBuilder interface {
	SetName(name string) DeploymentBuilder
	SetNamespace(namespace string) DeploymentBuilder
	SetLabelsAndSelectors(labels map[string]string, selectors *metav1.LabelSelector) DeploymentBuilder
	SetAnnotations(annos map[string]string) DeploymentBuilder
	SetShardingSphereProxyPodTemplate(tpl corev1.PodTemplateSpec) DeploymentBuilder
	SetShardingSphereProxyPodTemplateLabels(labels map[string]string) DeploymentBuilder
	SetShardingSphereProxyPodTemplateAnnotations(annos map[string]string) DeploymentBuilder
	SetShardingSphereProxyContainer(con *corev1.Container) DeploymentBuilder
	SetMySQLConnector(scb common.ContainerBuilder, cn *v1alpha1.ComputeNode) DeploymentBuilder
	SetAgentBin(scb common.ContainerBuilder, cn *v1alpha1.ComputeNode) DeploymentBuilder
	SetInitContainer(con *corev1.Container) DeploymentBuilder
	SetVolume(volume *corev1.Volume) DeploymentBuilder
	SetReplicas(r *int32) DeploymentBuilder
	Build() *appsv1.Deployment
}

// NewDeploymentBuilder creates a new DeploymentBuilder
func NewDeploymentBuilder(meta metav1.Object, gvk schema.GroupVersionKind) DeploymentBuilder {
	return &deploymentBuilder{
		deployment: DefaultDeployment(meta, gvk),
	}
}

type deploymentBuilder struct {
	deployment *appsv1.Deployment
}

// SetName sets Deployment name
func (d *deploymentBuilder) SetName(name string) DeploymentBuilder {
	d.deployment.Name = name
	return d
}

// SetNamespace sets Deployment namespace
func (d *deploymentBuilder) SetNamespace(namespace string) DeploymentBuilder {
	d.deployment.Namespace = namespace
	return d
}

// SetLabelsAndSelectors sets labels and selectors to Deployment labels, spec.selectors
// and spec.template.labels
func (d *deploymentBuilder) SetLabelsAndSelectors(labels map[string]string, selectors *metav1.LabelSelector) DeploymentBuilder {
	d.deployment.Labels = labels
	d.deployment.Spec.Selector = selectors
	d.deployment.Spec.Template.Labels = labels
	return d
}

// SetAnnotations sets Deployment annotations
func (d *deploymentBuilder) SetAnnotations(annos map[string]string) DeploymentBuilder {
	d.deployment.Annotations = annos
	return d
}

// SetReplicas sets Deployment replicas
func (d *deploymentBuilder) SetReplicas(r *int32) DeploymentBuilder {
	d.deployment.Spec.Replicas = r
	return d
}

// SetShardingSphereProxyPodTemplate sets Deployment PodTemplateSpec for ShardingSphereProxy Pod
func (d *deploymentBuilder) SetShardingSphereProxyPodTemplate(tpl corev1.PodTemplateSpec) DeploymentBuilder {
	d.deployment.Spec.Template = tpl
	return d
}

// SetShardingSphereProxyPodTemplateAnnotations sets annotations for ShardingSphereProxy Pod
func (d *deploymentBuilder) SetShardingSphereProxyPodTemplateAnnotations(annotations map[string]string) DeploymentBuilder {
	d.deployment.Spec.Template.Annotations = annotations
	return d
}

// SetShardingSphereProxyPodTemplateLabels sets labels for ShardingSphereProxy Pod
func (d *deploymentBuilder) SetShardingSphereProxyPodTemplateLabels(labels map[string]string) DeploymentBuilder {
	d.deployment.Spec.Template.Labels = labels
	return d
}

// SetShardingSphereProxyContainer sets a container for ShardingSphereProxy
func (d *deploymentBuilder) SetShardingSphereProxyContainer(proxy *corev1.Container) DeploymentBuilder {
	if d.deployment.Spec.Template.Spec.Containers == nil {
		d.deployment.Spec.Template.Spec.Containers = []corev1.Container{*proxy}
	}

	for idx := range d.deployment.Spec.Template.Spec.Containers {
		if d.deployment.Spec.Template.Spec.Containers[idx].Name == defaultContainerName {
			d.deployment.Spec.Template.Spec.Containers[idx] = *proxy
			return d
		}
	}

	d.deployment.Spec.Template.Spec.Containers = append(d.deployment.Spec.Template.Spec.Containers, *proxy)
	return d
}

// SetInitContainer sets the a init container for bootstrapping
func (d *deploymentBuilder) SetInitContainer(init *corev1.Container) DeploymentBuilder {
	if d.deployment.Spec.Template.Spec.InitContainers == nil {
		d.deployment.Spec.Template.Spec.InitContainers = []corev1.Container{}
	}

	for idx := range d.deployment.Spec.Template.Spec.InitContainers {
		if d.deployment.Spec.Template.Spec.InitContainers[idx].Name == init.Name {
			d.deployment.Spec.Template.Spec.InitContainers[idx] = *init
			return d
		}
	}

	d.deployment.Spec.Template.Spec.InitContainers = append(d.deployment.Spec.Template.Spec.InitContainers, *init)

	return d
}

// SharedVolumeAndMountBuilder build a Volume which could be mounted by different containers
type SharedVolumeAndMountBuilder interface {
	SetName(name string) SharedVolumeAndMountBuilder
	SetMountPath(idx int, path string) SharedVolumeAndMountBuilder
	SetSubPath(idx int, subpath string) SharedVolumeAndMountBuilder
	SetVolumeMountSize(size int) SharedVolumeAndMountBuilder
	SetVolumeSourceEmptyDir() SharedVolumeAndMountBuilder
	SetVolumeSourceConfigMap(name string, kps ...corev1.KeyToPath) SharedVolumeAndMountBuilder
	Build() (*corev1.Volume, []*corev1.VolumeMount)
}

type sharedVolumeAndMountBuilder struct {
	volume       *corev1.Volume
	volumeMounts []*corev1.VolumeMount
}

// NewSharedVolumeAndMountBuilder returns a new SharedVolumeAndMountBuilder
func NewSharedVolumeAndMountBuilder() SharedVolumeAndMountBuilder {
	return &sharedVolumeAndMountBuilder{
		volume:       &corev1.Volume{},
		volumeMounts: []*corev1.VolumeMount{},
	}
}

// SetName sets Volume and VolumeMounts name
func (b *sharedVolumeAndMountBuilder) SetName(name string) SharedVolumeAndMountBuilder {
	b.volume.Name = name
	for vm := range b.volumeMounts {
		b.volumeMounts[vm].Name = name
	}
	return b
}

// SetVolumeMountSize sets size of VolumeMounts
func (b *sharedVolumeAndMountBuilder) SetVolumeMountSize(size int) SharedVolumeAndMountBuilder {
	if len(b.volumeMounts) != size {
		vms := make([]*corev1.VolumeMount, size)
		for vm := range b.volumeMounts {
			vms[vm] = b.volumeMounts[vm].DeepCopy()
		}
		b.volumeMounts = vms
	}

	for vm := range b.volumeMounts {
		if b.volumeMounts[vm] == nil {
			b.volumeMounts[vm] = &corev1.VolumeMount{}
		}
	}
	return b
}

// SetMountPath sets mountPath of a Volume
func (b *sharedVolumeAndMountBuilder) SetMountPath(idx int, path string) SharedVolumeAndMountBuilder {
	if b.volumeMounts[idx] == nil {
		b.volumeMounts[idx] = &corev1.VolumeMount{}
	}
	b.volumeMounts[idx].MountPath = path
	return b
}

// SetSubPath sets subPath of a Volume
func (b *sharedVolumeAndMountBuilder) SetSubPath(idx int, subpath string) SharedVolumeAndMountBuilder {
	if b.volumeMounts[idx] == nil {
		b.volumeMounts[idx] = &corev1.VolumeMount{}
	}
	b.volumeMounts[idx].SubPath = subpath
	return b
}

// SetVolumeSourceEmptyDir sets a EmptyDir as Volume
func (b *sharedVolumeAndMountBuilder) SetVolumeSourceEmptyDir() SharedVolumeAndMountBuilder {
	if b.volume.EmptyDir == nil {
		b.volume.EmptyDir = &corev1.EmptyDirVolumeSource{}
	}
	return b
}

// SetVolumeSourceConfigMap sets a ConfigMap as Volume
func (b *sharedVolumeAndMountBuilder) SetVolumeSourceConfigMap(name string, kps ...corev1.KeyToPath) SharedVolumeAndMountBuilder {
	if b.volume.ConfigMap == nil {
		b.volume.ConfigMap = &corev1.ConfigMapVolumeSource{}
	}
	b.volume.ConfigMap.LocalObjectReference.Name = name

	if len(kps) > 0 {
		b.volume.ConfigMap.Items = kps
	}
	return b
}

// Build creates a new volume and volumeMounts
func (b *sharedVolumeAndMountBuilder) Build() (*corev1.Volume, []*corev1.VolumeMount) {
	return b.volume, b.volumeMounts
}

// VolumeAndMountBuilder build a Volume and related VolumeMount
type VolumeAndMountBuilder interface {
	SetName(string) VolumeAndMountBuilder
	Build() (*corev1.Volume, *corev1.VolumeMount)
}

// NewVolumeAndMountBuilder returns a VolumeAndMountBuilder
func NewVolumeAndMountBuilder() VolumeAndMountBuilder {
	return &volumeAndMountBuilder{
		volume:      &corev1.Volume{},
		volumemount: &corev1.VolumeMount{},
	}
}

type volumeAndMountBuilder struct {
	volume      *corev1.Volume
	volumemount *corev1.VolumeMount
}

// SetName sets Volume and VolumeMount name
func (b *volumeAndMountBuilder) SetName(name string) VolumeAndMountBuilder {
	b.volume.Name = name
	b.volumemount.Name = name
	return b
}

// SetMountPath sets mountPath of VolumeMount
func (b *volumeAndMountBuilder) SetMountPath(path string) VolumeAndMountBuilder {
	b.volumemount.MountPath = path
	return b
}

// SetSubPath sets subPath of VolumeMount
func (b *volumeAndMountBuilder) SetSubPath(subpath string) VolumeAndMountBuilder {
	b.volumemount.SubPath = subpath
	return b
}

// SetVolumeSourceEmptyDir sets EmptyDir as a Volume
func (b *volumeAndMountBuilder) SetVolumeSourceEmptyDir() VolumeAndMountBuilder {
	b.volume.EmptyDir = &corev1.EmptyDirVolumeSource{}
	return b
}

// SetVolumeSourceConfigMap sets ConfigMap as a Volume
func (b *volumeAndMountBuilder) SetVolumeSourceConfigMap(name string) VolumeAndMountBuilder {
	b.volume.ConfigMap.LocalObjectReference.Name = name
	return b
}

// Build builds a Volume and VolumeMount
func (b *volumeAndMountBuilder) Build() (*corev1.Volume, *corev1.VolumeMount) {
	return b.volume, b.volumemount
}

// SetVolume sets a volume for Deployment
func (d *deploymentBuilder) SetVolume(vol *corev1.Volume) DeploymentBuilder {
	if d.deployment.Spec.Template.Spec.Volumes == nil {
		d.deployment.Spec.Template.Spec.Volumes = []corev1.Volume{*vol}
	}

	for idx := range d.deployment.Spec.Template.Spec.Volumes {
		if d.deployment.Spec.Template.Spec.Volumes[idx].Name == vol.Name {
			d.deployment.Spec.Template.Spec.Volumes[idx] = *vol
			return d
		}
	}

	d.deployment.Spec.Template.Spec.Volumes = append(d.deployment.Spec.Template.Spec.Volumes, *vol)
	return d
}

// Build returns a Deployment
func (d *deploymentBuilder) Build() *appsv1.Deployment {
	return d.deployment
}

// NewDeployment creates a new Deployment
func NewDeployment(cn *v1alpha1.ComputeNode) *appsv1.Deployment {
	builder := NewDeploymentBuilder(cn.GetObjectMeta(), cn.GetObjectKind().GroupVersionKind())
	builder.SetName(cn.Name).SetNamespace(cn.Namespace).SetLabelsAndSelectors(cn.Labels, cn.Spec.Selector).SetAnnotations(cn.Annotations).SetReplicas(&cn.Spec.Replicas)

	ports := []corev1.ContainerPort{}
	for idx := range cn.Spec.PortBindings {
		ports = append(ports, corev1.ContainerPort{
			Name:          cn.Spec.PortBindings[idx].Name,
			HostIP:        cn.Spec.PortBindings[idx].HostIP,
			ContainerPort: cn.Spec.PortBindings[idx].ContainerPort,
			Protocol:      cn.Spec.PortBindings[idx].Protocol,
		})
	}

	scb := NewShardingSphereProxyContainerBuilder().
		SetVersion(cn.Spec.ServerVersion).
		SetPorts(ports).
		SetResources(cn.Spec.Resources)

	setProbes(scb, cn)

	vcb := NewSharedVolumeAndMountBuilder().
		SetVolumeMountSize(1).
		SetName(defaultConfigVolumeName).
		SetVolumeSourceConfigMap(cn.Name).
		SetMountPath(0, defaultConfigVolumeMountPath)
	vc, vmc := vcb.Build()

	builder.SetVolume(vc)
	scb.SetVolumeMount(vmc[0])

	// set agent for proxy
	if enabled, ok := cn.Annotations[defaultAnnotationJavaAgentEnabled]; ok && enabled == "true" {
		builder.SetAgentBin(scb, cn)

		metricsAnnos := map[string]string{}
		metricsAnnos[commonAnnotationPrometheusMetricsPath] = cn.Annotations[commonAnnotationPrometheusMetricsPath]
		metricsAnnos[commonAnnotationPrometheusMetricsPort] = cn.Annotations[commonAnnotationPrometheusMetricsPort]
		metricsAnnos[commonAnnotationPrometheusMetricsScrape] = cn.Annotations[commonAnnotationPrometheusMetricsScrape]
		metricsAnnos[commonAnnotationPrometheusMetricsScheme] = cn.Annotations[commonAnnotationPrometheusMetricsScheme]

		builder.SetShardingSphereProxyPodTemplateAnnotations(metricsAnnos)
	}

	if cn.Spec.StorageNodeConnector != nil {
		switch cn.Spec.StorageNodeConnector.Type {
		case v1alpha1.ConnectorTypeMySQL:
			builder.SetMySQLConnector(scb, cn)
		case v1alpha1.ConnectorTypePostgreSQL:
			sc := scb.Build()
			builder.SetShardingSphereProxyContainer(sc)
		}
	}

	return builder.Build()
}

func setProbes(scb common.ContainerBuilder, cn *v1alpha1.ComputeNode) {
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

// SetMySQLConnector will set an init container to download mysql jar and mount files for proxy container.
func (d *deploymentBuilder) SetMySQLConnector(scb common.ContainerBuilder, cn *v1alpha1.ComputeNode) DeploymentBuilder {
	scb.SetEnv([]corev1.EnvVar{
		{
			Name:  defaultMySQLDriverEnvName,
			Value: cn.Spec.StorageNodeConnector.Version,
		},
	})

	vb := NewSharedVolumeAndMountBuilder().
		SetVolumeMountSize(2).
		SetName(defaultMySQLDriverVolumeName).
		SetVolumeSourceEmptyDir().
		SetMountPath(0, defaultExtlibPath).
		SetMountPath(1, absoluteMySQLDriverMountName(defaultExtlibPath, cn.Spec.StorageNodeConnector.Version)).
		SetSubPath(1, relativeMySQLDriverMountName(cn.Spec.StorageNodeConnector.Version))

	v, vms := vb.Build()
	d.SetVolume(v)
	scb.SetVolumeMount(vms[1])

	cb := NewBootstrapContainerBuilderForMysqlJar().SetVolumeMount(vms[0]).SetEnv([]corev1.EnvVar{
		{
			Name:  defaultMySQLDriverEnvName,
			Value: cn.Spec.StorageNodeConnector.Version,
		},
	})
	con := cb.Build()
	d.SetInitContainer(con)

	sc := scb.Build()
	d.SetShardingSphereProxyContainer(sc)

	return d
}

// SetAgentBin set `agent bin` for ShardingSphereProxy with [observability](https://shardingsphere.apache.org/document/current/en/user-manual/shardingsphere-proxy/observability/)
func (d *deploymentBuilder) SetAgentBin(scb common.ContainerBuilder, cn *v1alpha1.ComputeNode) DeploymentBuilder {
	// set env JAVA_TOOL_OPTIONS to proxy container, make sure proxy will apply agent-bin.jar
	// agent-bin's version is always equals to shardingsphere proxy image's version
	scb.SetEnv([]corev1.EnvVar{
		{
			Name:  defaultJavaToolOptionsName,
			Value: fmt.Sprintf(defaultJavaAgentEnvValue, cn.Spec.ServerVersion),
		},
	})

	// mount agent-bin dir
	vbAgent := NewSharedVolumeAndMountBuilder().
		SetVolumeMountSize(1).
		SetName(defaultJavaAgentVolumeName).
		SetVolumeSourceEmptyDir().
		SetMountPath(0, defaultJavaAgentVolumeMountPath)
	va, vma := vbAgent.Build()
	d.SetVolume(va)
	scb.SetVolumeMount(vma[0])

	// mount agent config to overwrite agent-bin's config
	vbAgentConf := NewSharedVolumeAndMountBuilder().
		SetVolumeMountSize(1).
		SetName(defaultJavaAgentConfigVolumeName).
		SetVolumeSourceConfigMap(cn.Name, corev1.KeyToPath{Key: configmap.ConfigDataKeyForAgent, Path: configmap.ConfigDataKeyForAgent}).
		SetMountPath(0, defaultJavaAgentConfigVolumeMountPath)
	vc, vmc := vbAgentConf.Build()
	d.SetVolume(vc)
	scb.SetVolumeMount(vmc[0])

	cb := NewBootstrapContainerBuilderForAgentBin().SetVolumeMount(vma[0]).SetEnv([]corev1.EnvVar{
		{
			Name:  defaultAgentBinVersionEnvName,
			Value: cn.Spec.ServerVersion,
		},
	})
	con := cb.Build()
	d.SetInitContainer(con)

	sc := scb.Build()
	d.SetShardingSphereProxyContainer(sc)

	return d
}

// DefaultDeployment describes the default deployment
func DefaultDeployment(meta metav1.Object, gvk schema.GroupVersionKind) *appsv1.Deployment {
	defaultMaxUnavailable := intstr.FromInt(0)
	defaultMaxSurge := intstr.FromInt(3)

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "shardingsphere-proxy",
			Namespace:   "default",
			Labels:      map[string]string{},
			Annotations: map[string]string{},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(meta, gvk),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &appsv1.RollingUpdateDeployment{
					MaxUnavailable: &defaultMaxUnavailable,
					MaxSurge:       &defaultMaxSurge,
				},
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels:      map[string]string{},
					Annotations: map[string]string{},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            defaultContainerName,
							Image:           defaultImage,
							ImagePullPolicy: corev1.PullIfNotPresent,
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      defaultConfigVolumeName,
									MountPath: defaultConfigVolumeMountPath,
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: defaultConfigVolumeName,
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: defaultConfigVolumeName,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
