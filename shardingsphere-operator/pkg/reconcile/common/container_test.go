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
	corev1 "k8s.io/api/core/v1"
	"testing"
)

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
	c.SetVolumeMount(mount)

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
	c.SetVolumeMount(updatedMount)

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
