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
	"strings"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/apps/v1"
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

	download_script = `wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/${MYSQL_CONNECTOR_VERSION}/mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar;
wget https://repo1.maven.org/maven2/mysql/mysql-connector-java/${MYSQL_CONNECTOR_VERSION}/mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar.md5;
if [ $(md5sum /mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar | cut -d ' ' -f1) = $(cat /mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar.md5) ];
then echo success;
else echo failed;exit 1;fi;mv /mysql-connector-java-${MYSQL_CONNECTOR_VERSION}.jar /opt/shardingsphere-proxy/ext-lib`
)

func relativeMySQLDriverMountName(v string) string {
	return fmt.Sprintf("mysql-connector-java-%s.jar", v)
}

func absoluteMySQLDriverMountName(p, v string) string {
	return fmt.Sprintf("%s/%s", p, relativeMySQLDriverMountName(v))
}

type ShardingSphereProxyContainerBuilder interface {
	ContainerBuilder
	SetVersion(version string) ShardingSphereProxyContainerBuilder
}

type shardingSphereProxyContainerBuilder struct {
	container *corev1.Container
	ContainerBuilder
}

func (c *shardingSphereProxyContainerBuilder) SetVersion(version string) ShardingSphereProxyContainerBuilder {
	if len(strings.Split(c.container.Image, ":")) > 0 {
		repo := strings.TrimSuffix(c.container.Image, "5.3.0")
		c.container.Image = fmt.Sprintf("%s%s", repo, version)
	}
	return c
}

func NewShardingSphereProxyContainerBuilder() ShardingSphereProxyContainerBuilder {
	pb := NewContainerBuilder().
		SetName(defaultContainerName).
		SetImage(defaultImage)

	return &shardingSphereProxyContainerBuilder{
		ContainerBuilder: NewContainerBuilder(),
		container:        pb.Build(),
	}
}

func (b *shardingSphereProxyContainerBuilder) Build() *corev1.Container {
	return b.container
}

type BootstrapContainerBuilder interface {
	ContainerBuilder
}

type bootstrapContainerBuilder struct {
	container *corev1.Container
	ContainerBuilder
}

func NewBootstrapContainerBuilder() BootstrapContainerBuilder {
	pb := NewContainerBuilder().
		SetName(defaultContainerName).
		SetImage(defaultImage)

	return &bootstrapContainerBuilder{
		container:        pb.Build(),
		ContainerBuilder: NewContainerBuilder(),
	}
}

func (b *bootstrapContainerBuilder) Build() *corev1.Container {
	return b.container
}

type ContainerBuilder interface {
	SetName(name string) ContainerBuilder
	SetImage(image string) ContainerBuilder
	SetPorts(ports []corev1.ContainerPort) ContainerBuilder
	SetResources(res corev1.ResourceRequirements) ContainerBuilder
	SetLivenessProbe(probe *corev1.Probe) ContainerBuilder
	SetReadinessProbe(probe *corev1.Probe) ContainerBuilder
	SetStartupProbe(probe *corev1.Probe) ContainerBuilder
	SetEnv(envs []corev1.EnvVar) ContainerBuilder
	SetVolumeMount(mount *corev1.VolumeMount) ContainerBuilder
	Build() *corev1.Container
}

func NewContainerBuilder() ContainerBuilder {
	return &containerBuilder{
		container: DefaultContainer(),
	}
}

type containerBuilder struct {
	container *corev1.Container
}

func (c *containerBuilder) SetName(name string) ContainerBuilder {
	c.container.Name = name
	return c
}

func (c *containerBuilder) SetImage(image string) ContainerBuilder {
	c.container.Image = image
	return c
}

func (c *containerBuilder) SetPorts(ports []corev1.ContainerPort) ContainerBuilder {
	if ports != nil {
		c.container.Ports = ports
	}
	return c
}

func (c *containerBuilder) SetResources(res corev1.ResourceRequirements) ContainerBuilder {
	c.container.Resources = res
	return c
}

func (c *containerBuilder) SetLivenessProbe(probe *corev1.Probe) ContainerBuilder {
	if probe != nil {
		if c.container.LivenessProbe == nil {
			c.container.LivenessProbe = &corev1.Probe{}
		}
		c.container.LivenessProbe = probe
	}
	return c
}

func (c *containerBuilder) SetReadinessProbe(probe *corev1.Probe) ContainerBuilder {
	if probe != nil {
		if c.container.ReadinessProbe == nil {
			c.container.ReadinessProbe = &corev1.Probe{}
		}
		c.container.ReadinessProbe = probe
	}
	return c
}

func (c *containerBuilder) SetStartupProbe(probe *corev1.Probe) ContainerBuilder {
	if probe != nil {
		if c.container.StartupProbe == nil {
			c.container.StartupProbe = &corev1.Probe{}
		}
		c.container.StartupProbe = probe
	}
	return c
}

func (c *containerBuilder) SetEnv(envs []corev1.EnvVar) ContainerBuilder {
	if envs != nil {
		c.container.Env = envs
	}
	return c
}

func (c *containerBuilder) SetCommand(cmds []string) ContainerBuilder {
	if cmds != nil {
		c.container.Command = cmds
	}
	return c
}

func (c *containerBuilder) SetVolumeMount(mount *corev1.VolumeMount) ContainerBuilder {
	var found bool
	if c.container.VolumeMounts != nil {
		for idx, v := range c.container.VolumeMounts {
			if v.Name == mount.Name {
				c.container.VolumeMounts[idx] = *mount
				found = true
				break
			}
		}
		if !found {
			c.container.VolumeMounts = append(c.container.VolumeMounts, *mount)
		}
	}

	return c
}

func (c *containerBuilder) Build() *corev1.Container {
	return c.container
}

func DefaultContainer() *corev1.Container {
	con := &corev1.Container{
		Name:  "default",
		Image: "busybox:1.35.0",
		Ports: []corev1.ContainerPort{},
		Env:   []corev1.EnvVar{},
		// LivenessProbe:  &corev1.Probe{},
		// ReadinessProbe: &corev1.Probe{},
		// StartupProbe:   &corev1.Probe{},
		// VolumeMounts: []corev1.VolumeMount{},
	}
	return con
}

type DeploymentBuilder interface {
	SetName(name string) DeploymentBuilder
	SetNamespace(namespace string) DeploymentBuilder
	SetLabelsAndSelectors(labels map[string]string, selectors *metav1.LabelSelector) DeploymentBuilder
	SetShardingSphereProxyContainer(*corev1.Container) DeploymentBuilder
	SetInitContainer(*corev1.Container) DeploymentBuilder
	SetVolume(*corev1.Volume) DeploymentBuilder
	Build() *appsv1.Deployment
}

func NewDeploymentBuilder(meta metav1.Object, gvk schema.GroupVersionKind) DeploymentBuilder {
	return &deploymentBuilder{
		deployment: DefaultDeployment(meta, gvk),
	}
}

type deploymentBuilder struct {
	deployment *appsv1.Deployment
}

func (d *deploymentBuilder) SetName(name string) DeploymentBuilder {
	d.deployment.Name = name
	return d
}

func (d *deploymentBuilder) SetNamespace(namespace string) DeploymentBuilder {
	d.deployment.Namespace = namespace
	return d
}

func (d *deploymentBuilder) SetLabelsAndSelectors(labels map[string]string, selectors *metav1.LabelSelector) DeploymentBuilder {
	d.deployment.Labels = labels
	d.deployment.Spec.Selector = selectors
	d.deployment.Spec.Template.Labels = labels
	return d
}

func (d *deploymentBuilder) SetShardingSphereProxyContainer(proxy *corev1.Container) DeploymentBuilder {
	var found bool
	if d.deployment.Spec.Template.Spec.Containers != nil {
		for idx, c := range d.deployment.Spec.Template.Spec.Containers {
			if c.Name == defaultContainerName {
				found = true
				d.deployment.Spec.Template.Spec.Containers[idx] = *proxy
				break
			}
		}

		if !found {
			d.deployment.Spec.Template.Spec.Containers = append(d.deployment.Spec.Template.Spec.Containers, *proxy)
		}
	}

	return d
}

func (d *deploymentBuilder) SetInitContainer(init *corev1.Container) DeploymentBuilder {
	var found bool
	if d.deployment.Spec.Template.Spec.InitContainers != nil {
		for idx, c := range d.deployment.Spec.Template.Spec.InitContainers {
			if c.Name == defaultContainerName {
				found = true
				d.deployment.Spec.Template.Spec.InitContainers[idx] = *init
				break
			}
		}

		if !found {
			d.deployment.Spec.Template.Spec.InitContainers = append(d.deployment.Spec.Template.Spec.InitContainers, *init)
		}
	}

	return d
}

type SharedVolumeAndMountBuilder interface {
	SetName(name string) SharedVolumeAndMountBuilder
	SetMountPath(idx int, path string) SharedVolumeAndMountBuilder
	SetSubPath(idx int, subpath string) SharedVolumeAndMountBuilder
	SetVolumeMountSize(size int) SharedVolumeAndMountBuilder
	SetVolumeSourceEmptyDir() SharedVolumeAndMountBuilder
	SetVolumeSourceConfigMap(name string) SharedVolumeAndMountBuilder
	Build() (*corev1.Volume, []*corev1.VolumeMount)
}
type sharedVolumeAndMountBuilder struct {
	volume       *corev1.Volume
	volumeMounts []*corev1.VolumeMount
}

func NewSharedVolumeAndMountBuilder() SharedVolumeAndMountBuilder {
	return &sharedVolumeAndMountBuilder{
		volume:       &corev1.Volume{},
		volumeMounts: []*corev1.VolumeMount{},
	}
}

func (b *sharedVolumeAndMountBuilder) SetName(name string) SharedVolumeAndMountBuilder {
	b.volume.Name = name
	for vm := range b.volumeMounts {
		b.volumeMounts[vm].Name = name
	}
	return b
}

func (b *sharedVolumeAndMountBuilder) SetVolumeMountSize(size int) SharedVolumeAndMountBuilder {
	b.volumeMounts = make([]*corev1.VolumeMount, size)
	return b
}

func (b *sharedVolumeAndMountBuilder) SetMountPath(idx int, path string) SharedVolumeAndMountBuilder {
	if b.volumeMounts[idx] == nil {
		b.volumeMounts[idx] = &corev1.VolumeMount{}
	}
	b.volumeMounts[idx].MountPath = path
	return b
}

func (b *sharedVolumeAndMountBuilder) SetSubPath(idx int, subpath string) SharedVolumeAndMountBuilder {
	if b.volumeMounts[idx] == nil {
		b.volumeMounts[idx] = &corev1.VolumeMount{}
	}
	b.volumeMounts[idx].SubPath = subpath
	return b
}

func (b *sharedVolumeAndMountBuilder) SetVolumeSourceEmptyDir() SharedVolumeAndMountBuilder {
	b.volume.EmptyDir = &corev1.EmptyDirVolumeSource{}
	return b
}

func (b *sharedVolumeAndMountBuilder) SetVolumeSourceConfigMap(name string) SharedVolumeAndMountBuilder {
	b.volume.ConfigMap.LocalObjectReference.Name = name
	return b
}

func (b *sharedVolumeAndMountBuilder) Build() (*corev1.Volume, []*corev1.VolumeMount) {
	return b.volume, b.volumeMounts
}

type VolumeAndMountBuilder interface {
	SetName(string) VolumeAndMountBuilder
	Build() (*corev1.Volume, *corev1.VolumeMount)
}

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

func (b *volumeAndMountBuilder) SetName(name string) VolumeAndMountBuilder {
	b.volume.Name = name
	b.volumemount.Name = name
	return b
}

func (b *volumeAndMountBuilder) SetMountPath(path string) VolumeAndMountBuilder {
	b.volumemount.MountPath = path
	return b
}

func (b *volumeAndMountBuilder) SetSubPath(subpath string) VolumeAndMountBuilder {
	b.volumemount.SubPath = subpath
	return b
}

func (b *volumeAndMountBuilder) SetVolumeSourceEmptyDir() VolumeAndMountBuilder {
	b.volume.EmptyDir = &corev1.EmptyDirVolumeSource{}
	return b
}

func (b *volumeAndMountBuilder) SetVolumeSourceConfigMap(name string) VolumeAndMountBuilder {
	b.volume.ConfigMap.LocalObjectReference.Name = name
	return b
}

func (b *volumeAndMountBuilder) Build() (*corev1.Volume, *corev1.VolumeMount) {
	return b.volume, b.volumemount
}

func (d *deploymentBuilder) SetVolume(v *corev1.Volume) DeploymentBuilder {
	var found bool
	if d.deployment.Spec.Template.Spec.Volumes != nil {
		for idx := range d.deployment.Spec.Template.Spec.Volumes {
			if d.deployment.Spec.Template.Spec.Volumes[idx].Name == v.Name {
				d.deployment.Spec.Template.Spec.Volumes[idx] = *v
				found = true
				break
			}
		}
		if !found {
			d.deployment.Spec.Template.Spec.Volumes = append(d.deployment.Spec.Template.Spec.Volumes, *v)
		}
	}

	return d
}

func (d *deploymentBuilder) Build() *appsv1.Deployment {
	return d.deployment
}

func NewDeployment(cn *v1alpha1.ComputeNode) *v1.Deployment {
	deploy := DefaultDeployment(cn.GetObjectMeta(), cn.GroupVersionKind())

	// basic information
	deploy.Name = cn.Name
	deploy.Namespace = cn.Namespace
	deploy.Labels = cn.Labels
	deploy.Spec.Selector = cn.Spec.Selector
	deploy.Spec.Replicas = &cn.Spec.Replicas
	deploy.Spec.Template.Labels = cn.Labels
	deploy.Spec.Template.Spec.Containers[0].Image = fmt.Sprintf("%s:%s", defaultImage, cn.Spec.ServerVersion)

	deploy.Spec.Template.Spec.Containers[0].Ports = []corev1.ContainerPort{}
	for _, pb := range cn.Spec.PortBindings {
		deploy.Spec.Template.Spec.Containers[0].Ports = append(deploy.Spec.Template.Spec.Containers[0].Ports, corev1.ContainerPort{
			Name:          pb.Name,
			HostIP:        pb.HostIP,
			ContainerPort: pb.ContainerPort,
			Protocol:      pb.Protocol,
		})
	}

	// additional information
	deploy.Spec.Template.Spec.Containers[0].Resources = cn.Spec.Resources
	for _, v := range deploy.Spec.Template.Spec.Volumes {
		if v.Name == defaultConfigVolumeName {
			v.ConfigMap.LocalObjectReference.Name = cn.Name
		}
	}

	if cn.Spec.Probes != nil {
		if cn.Spec.Probes.StartupProbe != nil {
			deploy.Spec.Template.Spec.Containers[0].StartupProbe = cn.Spec.Probes.StartupProbe.DeepCopy()
		}
		if cn.Spec.Probes.LivenessProbe != nil {
			deploy.Spec.Template.Spec.Containers[0].LivenessProbe = cn.Spec.Probes.LivenessProbe.DeepCopy()
		}
		if cn.Spec.Probes.ReadinessProbe != nil {
			deploy.Spec.Template.Spec.Containers[0].ReadinessProbe = cn.Spec.Probes.ReadinessProbe.DeepCopy()
		}
	}
	if len(cn.Spec.ImagePullSecrets) > 0 {
		deploy.Spec.Template.Spec.ImagePullSecrets = cn.Spec.ImagePullSecrets
	}
	if cn.Spec.StorageNodeConnector != nil {
		if cn.Spec.StorageNodeConnector.Type == v1alpha1.ConnectorTypeMySQL {
			// add or update initContainer
			if len(deploy.Spec.Template.Spec.InitContainers) > 0 {
				for idx, v := range deploy.Spec.Template.Spec.InitContainers[0].Env {
					if v.Name == defaultMySQLDriverEnvName {
						deploy.Spec.Template.Spec.InitContainers[0].Env[idx].Value = cn.Spec.StorageNodeConnector.Version
					}
				}
			} else {
				deploy.Spec.Template.Spec.InitContainers = []corev1.Container{
					{
						Name:    "boostrap",
						Image:   "busybox:1.35.0",
						Command: []string{"/bin/sh", "-c", download_script},
						Env: []corev1.EnvVar{
							{
								Name:  defaultMySQLDriverEnvName,
								Value: cn.Spec.StorageNodeConnector.Version,
							},
						},
						VolumeMounts: []corev1.VolumeMount{
							{
								Name:      defaultMySQLDriverVolumeName,
								MountPath: defaultExtlibPath,
							},
						},
					},
				}

				deploy.Spec.Template.Spec.Containers[0].VolumeMounts = append(deploy.Spec.Template.Spec.Containers[0].VolumeMounts, corev1.VolumeMount{
					Name:      defaultMySQLDriverVolumeName,
					SubPath:   relativeMySQLDriverMountName(cn.Spec.StorageNodeConnector.Version),
					MountPath: absoluteMySQLDriverMountName(defaultExtlibPath, cn.Spec.StorageNodeConnector.Version),
				})

				deploy.Spec.Template.Spec.Volumes = append(deploy.Spec.Template.Spec.Volumes, corev1.Volume{
					Name: defaultMySQLDriverVolumeName,
					VolumeSource: corev1.VolumeSource{
						EmptyDir: &corev1.EmptyDirVolumeSource{},
					},
				})
			}
		}
	}

	return deploy
}

func DefaultDeployment(meta metav1.Object, gvk schema.GroupVersionKind) *v1.Deployment {
	defaultMaxUnavailable := intstr.FromInt(0)
	defaultMaxSurge := intstr.FromInt(3)

	return &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "shardingsphere-proxy",
			Namespace: "default",
			Labels:    map[string]string{},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(meta, gvk),
			},
		},
		Spec: v1.DeploymentSpec{
			Strategy: v1.DeploymentStrategy{
				Type: v1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &v1.RollingUpdateDeployment{
					MaxUnavailable: &defaultMaxUnavailable,
					MaxSurge:       &defaultMaxSurge,
				},
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            defaultContainerName,
							Image:           defaultImage,
							ImagePullPolicy: corev1.PullIfNotPresent,
							Ports: []corev1.ContainerPort{
								{
									Name:          "proxy",
									ContainerPort: 3307,
								},
							},
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

func UpdateDeployment(cn *v1alpha1.ComputeNode, cur *v1.Deployment) *v1.Deployment {
	exp := &v1.Deployment{}
	exp.ObjectMeta = cur.ObjectMeta
	exp.ObjectMeta.ResourceVersion = ""
	exp.Labels = cur.Labels
	exp.Annotations = cur.Annotations
	exp.Spec = NewDeployment(cn).Spec
	return exp
}
