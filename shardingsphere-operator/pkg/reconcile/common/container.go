package common

import (
	v1 "k8s.io/api/core/v1"
)

// ContainerBuilder is a common builder for Container
type ContainerBuilder interface {
	SetName(name string) ContainerBuilder
	SetImage(image string) ContainerBuilder
	SetPorts(ports []v1.ContainerPort) ContainerBuilder
	SetResources(res v1.ResourceRequirements) ContainerBuilder
	SetLivenessProbe(probe *v1.Probe) ContainerBuilder
	SetReadinessProbe(probe *v1.Probe) ContainerBuilder
	SetStartupProbe(probe *v1.Probe) ContainerBuilder
	SetEnv(envs []v1.EnvVar) ContainerBuilder
	SetCommand(cmds []string) ContainerBuilder
	SetVolumeMount(mount *v1.VolumeMount) ContainerBuilder
	Build() *v1.Container
}

// NewContainerBuilder return a builder for Container
func NewContainerBuilder() ContainerBuilder {
	return &containerBuilder{
		container: DefaultContainer(),
	}
}

type containerBuilder struct {
	container *v1.Container
}

// SetName sets the name of the container
func (c *containerBuilder) SetName(name string) ContainerBuilder {
	c.container.Name = name
	return c
}

// SetImage sets the name of the container
func (c *containerBuilder) SetImage(image string) ContainerBuilder {
	c.container.Image = image
	return c
}

// SetPorts set the container port of the container
func (c *containerBuilder) SetPorts(ports []v1.ContainerPort) ContainerBuilder {
	if ports == nil {
		c.container.Ports = []v1.ContainerPort{}
	}
	if ports != nil {
		c.container.Ports = ports
	}
	return c
}

// SetResources set the resources of the container
func (c *containerBuilder) SetResources(res v1.ResourceRequirements) ContainerBuilder {
	c.container.Resources = res
	return c
}

// SetLivenessProbe set the livenessProbe of the container
func (c *containerBuilder) SetLivenessProbe(probe *v1.Probe) ContainerBuilder {
	if probe != nil {
		if c.container.LivenessProbe == nil {
			c.container.LivenessProbe = &v1.Probe{}
		}
		c.container.LivenessProbe = probe
	}
	return c
}

// SetReadinessProbe set the readinessProbe of the container
func (c *containerBuilder) SetReadinessProbe(probe *v1.Probe) ContainerBuilder {
	if probe != nil {
		if c.container.ReadinessProbe == nil {
			c.container.ReadinessProbe = &v1.Probe{}
		}
		c.container.ReadinessProbe = probe
	}
	return c
}

// SetStartupProbe set the startupProbe of the container
func (c *containerBuilder) SetStartupProbe(probe *v1.Probe) ContainerBuilder {
	if probe != nil {
		if c.container.StartupProbe == nil {
			c.container.StartupProbe = &v1.Probe{}
		}
		c.container.StartupProbe = probe
	}
	return c
}

// SetEnv set the env of the container
func (c *containerBuilder) SetEnv(envs []v1.EnvVar) ContainerBuilder {
	if envs == nil {
		c.container.Env = []v1.EnvVar{}
	}
	if envs != nil {
		c.container.Env = envs
	}
	return c
}

// SetCommand set the command of the container
func (c *containerBuilder) SetCommand(cmds []string) ContainerBuilder {
	if cmds != nil {
		c.container.Command = cmds
	}
	return c
}

// SetVolumeMount set the mount point of the container
func (c *containerBuilder) SetVolumeMount(mount *v1.VolumeMount) ContainerBuilder {
	if c.container.VolumeMounts == nil {
		c.container.VolumeMounts = []v1.VolumeMount{*mount}
	} else {
		for idx := range c.container.VolumeMounts {
			if c.container.VolumeMounts[idx].Name == mount.Name {
				c.container.VolumeMounts[idx] = *mount
				return c
			}
		}
		c.container.VolumeMounts = append(c.container.VolumeMounts, *mount)
	}

	return c
}

// Build returns a Container
func (c *containerBuilder) Build() *v1.Container {
	return c.container
}

// DefaultContainer returns a container with busybox
func DefaultContainer() *v1.Container {
	con := &v1.Container{
		Name:  "default",
		Image: "busybox:1.35.0",
	}
	return con
}
