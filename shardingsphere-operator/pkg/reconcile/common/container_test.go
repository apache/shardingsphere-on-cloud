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
