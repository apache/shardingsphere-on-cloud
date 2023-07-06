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

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestContainerBuilder_BuildDefaultContainer(t *testing.T) {
	// create a new container builder
	c := NewContainerBuilder()

	// check if the container is created
	if c == nil {
		t.Errorf("NewContainerBuilder() failed to create a new container builder")
	}

	con := c.BuildContainer()

	// check if the container is initialized
	if con == nil {
		t.Errorf("Build() failed to initialize the container")
	}

	assert.EqualValues(t, con, DefaultContainer(), "Build() failed to initialize the container")
}

func TestContainerBUilder_BuildFullContainer(t *testing.T) {

	resourceQuantity := func(value string) *resource.Quantity {
		q, e := resource.ParseQuantity(value)
		if e == nil {
			return &resource.Quantity{}
		}
		return &q
	}

	c := NewContainerBuilder()
	c.SetName("test").
		SetImage("nginx:1.9.7").
		SetPorts([]v1.ContainerPort{
			{
				Name:          "http",
				ContainerPort: 80,
			},
		}).
		SetResources(v1.ResourceRequirements{
			Limits: v1.ResourceList{
				"cpu":    *resourceQuantity("2"),
				"memory": *resourceQuantity("2Gi"),
			},
			Requests: v1.ResourceList{
				"cpu":    *resourceQuantity("1"),
				"memory": *resourceQuantity("1Gi"),
			},
		}).
		SetLivenessProbe(&v1.Probe{
			ProbeHandler: v1.ProbeHandler{
				HTTPGet: &v1.HTTPGetAction{
					Path: "/health",
				},
			},
			PeriodSeconds:    10,
			TimeoutSeconds:   5,
			SuccessThreshold: 1,
			FailureThreshold: 3,
		}).
		SetReadinessProbe(&v1.Probe{
			ProbeHandler: v1.ProbeHandler{
				HTTPGet: &v1.HTTPGetAction{
					Path: "/metrics",
				},
			},
			PeriodSeconds:    10,
			TimeoutSeconds:   5,
			SuccessThreshold: 1,
			FailureThreshold: 3,
		}).
		SetStartupProbe(&v1.Probe{
			ProbeHandler: v1.ProbeHandler{
				HTTPGet: &v1.HTTPGetAction{
					Path: "/startup",
				},
			},
			PeriodSeconds:    10,
			TimeoutSeconds:   5,
			SuccessThreshold: 1,
			FailureThreshold: 3,
		}).
		SetEnv([]v1.EnvVar{
			{
				Name:  "test_key",
				Value: "test_value",
			},
		}).
		SetCommand([]string{
			"echo",
		}).
		AppendVolumeMounts([]v1.VolumeMount{
			{
				Name:      "test-mount",
				MountPath: "/test",
			},
		})

	con := c.BuildContainer()
	assert.Equal(t, "test", con.Name, "Build() failed to set the container name")
	assert.Equal(t, "nginx:1.9.7", con.Image, "Build() failed to set the container image")
	assert.EqualValues(t, &v1.Probe{
		ProbeHandler: v1.ProbeHandler{
			HTTPGet: &v1.HTTPGetAction{
				Path: "/health",
			},
		},
		PeriodSeconds:    10,
		TimeoutSeconds:   5,
		SuccessThreshold: 1,
		FailureThreshold: 3,
	}, con.LivenessProbe, "Build() failed to set the container liveness probe")
	assert.EqualValues(t, &v1.Probe{
		ProbeHandler: v1.ProbeHandler{
			HTTPGet: &v1.HTTPGetAction{
				Path: "/metrics",
			},
		},
		PeriodSeconds:    10,
		TimeoutSeconds:   5,
		SuccessThreshold: 1,
		FailureThreshold: 3,
	}, con.ReadinessProbe, "Build() failed to set the container readiness probe")
	assert.EqualValues(t, &v1.Probe{
		ProbeHandler: v1.ProbeHandler{
			HTTPGet: &v1.HTTPGetAction{
				Path: "/startup",
			},
		},
		PeriodSeconds:    10,
		TimeoutSeconds:   5,
		SuccessThreshold: 1,
		FailureThreshold: 3,
	}, con.StartupProbe, "Build() failed to set the container startup probe")

	assert.EqualValues(t, []v1.ContainerPort{
		{
			Name:          "http",
			ContainerPort: 80,
		},
	}, con.Ports, "Build() failed to set the container ports")
	assert.EqualValues(t, v1.ResourceRequirements{
		Limits: v1.ResourceList{
			"cpu":    *resourceQuantity("2"),
			"memory": *resourceQuantity("2Gi"),
		},
		Requests: v1.ResourceList{
			"cpu":    *resourceQuantity("1"),
			"memory": *resourceQuantity("1Gi"),
		},
	}, con.Resources, "Build() failed to set the container limits")
	assert.EqualValues(t, []v1.EnvVar{
		{
			Name:  "test_key",
			Value: "test_value",
		},
	}, con.Env, "Build() failed to set the container env")
	assert.EqualValues(t, []string{
		"echo",
	}, con.Command, "Build() failed to set the container command")
	assert.EqualValues(t, []v1.VolumeMount{
		{
			Name:      "test-mount",
			MountPath: "/test",
		},
	}, con.VolumeMounts, "Build() failed to set the container volume mounts")

}

func TestContainerBuilder_SetVolumeMount(t *testing.T) {
	var found bool

	// create a new container builder
	c := &containerBuilder{
		container: &corev1.Container{
			VolumeMounts: []corev1.VolumeMount{},
		},
	}

	// add a new volume mount
	mount := &corev1.VolumeMount{
		Name:      "test-mount",
		MountPath: "/test",
	}
	c.AppendVolumeMounts([]corev1.VolumeMount{*mount})

	// check if the volume mount has been added
	for _, v := range c.container.VolumeMounts {
		if v.Name == mount.Name {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("SetVolumeMount() failed to add the VolumeMount")
	}

	// update an existing volume mount
	updatedMount := &corev1.VolumeMount{
		Name:      "test-mount",
		MountPath: "/new-test",
	}
	c.UpdateVolumeMountByName(*updatedMount)

	// check if the volume mount has been updated
	for _, v := range c.container.VolumeMounts {
		if v.MountPath == updatedMount.MountPath {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("SetVolumeMount() failed to update the VolumeMount")
	}
}
