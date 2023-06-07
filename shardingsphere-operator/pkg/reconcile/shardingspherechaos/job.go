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

package shardingspherechaos

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/common"

	v1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	DefaultImageName     = "agoiyanzsa/tools-runtime:2.0"
	DefaultContainerName = "tools-runtime"
	DefaultWorkPath      = "/app/start"
	DefaultConfigName    = "cmd-conf"
)

var (
	DefaultTTLSecondsAfterFinished int32 = 300
)

var DefaultFileMode int32 = 493

const (
	AnnoJobCompletions             = "jobs.batch/completions"
	AnnoJobActiveDeadlineSeconds   = "jobs.batch/activeDeadlineSeconds"
	AnnoJobParallelism             = "job.batch/parallelism"
	AnnoJobBackoffLimit            = "job.batch/backoffLimit"
	AnnoJobTTLSecondsAfterFinished = "job.batch/ttlSecondsAfterFinished"
	AnnoJobSuspend                 = "job.batch/suspend"
)

type JobType string

var (
	InSteady JobType = "steady"
	InChaos  JobType = "chaos"
)

func MakeJobName(name string, requirement JobType) string {
	return fmt.Sprintf("%s-%s", name, string(requirement))
}

func NewJob(ssChaos *v1alpha1.Chaos, requirement JobType) (*v1.Job, error) {
	jbd := NewJobBuilder()
	jbd.SetNamespace(ssChaos.Namespace).SetLabels(ssChaos.Labels).SetName(MakeJobName(ssChaos.Name, requirement))

	c, _ := MustInt32(ssChaos.Annotations[AnnoJobCompletions])
	jbd.SetCompletions(c)

	c, _ = MustInt32(ssChaos.Annotations[AnnoJobParallelism])
	jbd.SetParallelism(c)

	c, _ = MustInt32(ssChaos.Annotations[AnnoJobBackoffLimit])
	jbd.SetBackoffLimit(c)

	c, _ = MustInt32(ssChaos.Annotations[AnnoJobTTLSecondsAfterFinished])
	jbd.SetTTLSecondsAfterFinished(c)

	t, _ := MustInt64(ssChaos.Annotations[AnnoJobActiveDeadlineSeconds])
	jbd.SetActiveDeadlineSeconds(t)

	if v, ok := ssChaos.Annotations[AnnoJobSuspend]; ok {
		if v == "true" {
			jbd.SetSuspend(true)
		}

		if v == "false" {
			jbd.SetSuspend(false)
		}
	}

	v := &corev1.Volume{Name: DefaultConfigName}

	v.ConfigMap = &corev1.ConfigMapVolumeSource{}
	v.ConfigMap.LocalObjectReference.Name = ssChaos.Name
	v.ConfigMap.DefaultMode = &DefaultFileMode
	jbd.SetVolumes(v)

	vm := &corev1.VolumeMount{Name: DefaultConfigName, MountPath: DefaultWorkPath}

	cbd := common.NewContainerBuilder()
	cbd.SetImage(DefaultImageName)
	cbd.SetName(DefaultContainerName)
	cbd.SetVolumeMount(vm)
	cbd.SetCommand([]string{"sh", "-c"})

	container := cbd.Build()
	container.Args = NewCmds(requirement)
	jbd.SetContainers(container)

	rjob := jbd.Build()

	return rjob, nil
}

func NewCmds(requirement JobType) []string {
	var cmds []string
	if requirement == InSteady {
		cmds = append(cmds, fmt.Sprintf("%s/%s", DefaultWorkPath, configExperimental))
	}
	if requirement == InChaos {
		cmds = append(cmds, fmt.Sprintf("%s/%s;%s/%s", DefaultWorkPath, configPressure, DefaultWorkPath, configExperimental))
	}
	return cmds
}

func MustInt32(s string) (int32, error) {
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

func MustInt64(s string) (int64, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(v), nil
}

func IsJobChanged(ssChaos *v1alpha1.Chaos, requirement JobType, cur *v1.Job) (bool, error) {
	now, err := NewJob(ssChaos, requirement)
	if err != nil {
		return false, err
	}
	isEqual := judgeJobEqual(cur, now)
	if isEqual {
		return true, nil
	}

	return false, nil
}

func judgeJobEqual(now *v1.Job, exp *v1.Job) bool {
	if !judgeJobConfigEqual(now, exp) {
		return false
	}
	if !judgeContainerEqual(&now.Spec.Template.Spec.Containers[0], &exp.Spec.Template.Spec.Containers[0]) {
		return false
	}
	return true
}

func judgeJobConfigEqual(now *v1.Job, exp *v1.Job) bool {
	if !judgeTTLSecondsAfterFinished(now.Spec.TTLSecondsAfterFinished, exp.Spec.TTLSecondsAfterFinished) {
		return false
	}
	if exp.Spec.BackoffLimit != nil && *now.Spec.BackoffLimit != *exp.Spec.BackoffLimit {
		return false
	}
	if exp.Spec.Suspend != nil && *now.Spec.Suspend != *exp.Spec.Suspend {
		return false
	}
	if exp.Spec.Parallelism != nil && *now.Spec.Parallelism != *exp.Spec.Parallelism {
		return false
	}
	if exp.Spec.Completions != nil && *now.Spec.Completions != *exp.Spec.Completions {
		return false
	}
	if !judgeActiveDeadlineSeconds(now.Spec.ActiveDeadlineSeconds, exp.Spec.ActiveDeadlineSeconds) {
		return false
	}
	return true
}
func judgeTTLSecondsAfterFinished(cur *int32, exp *int32) bool {
	if cur == nil && exp == nil {
		return true
	}
	if cur != nil && exp != nil {
		if *cur == *exp {
			return true
		}
	}
	return false
}
func judgeActiveDeadlineSeconds(cur *int64, exp *int64) bool {
	if exp != nil && *cur != *exp {
		return false
	}
	return true
}
func judgeContainerEqual(now *corev1.Container, exp *corev1.Container) bool {
	if now.Name != exp.Name {
		return false
	}
	if !reflect.DeepEqual(now.Command, exp.Command) {
		return false
	}
	if !reflect.DeepEqual(now.Args, exp.Args) {
		return false
	}
	if now.Image != exp.Image {
		return false
	}

	if !reflect.DeepEqual(now.VolumeMounts, now.VolumeMounts) {
		return false
	}

	return true
}

type JobBuilder interface {
	SetName(string) JobBuilder
	SetNamespace(string) JobBuilder
	SetLabels(map[string]string) JobBuilder
	SetCompletions(int32) JobBuilder
	SetActiveDeadlineSeconds(int64) JobBuilder
	SetParallelism(int32) JobBuilder
	SetBackoffLimit(int32) JobBuilder
	SetContainers(*corev1.Container) JobBuilder
	SetTTLSecondsAfterFinished(int32) JobBuilder
	SetSuspend(bool) JobBuilder
	SetVolumes(*corev1.Volume) JobBuilder
	Build() *v1.Job
}

func NewJobBuilder() JobBuilder {
	return &jobBuilder{
		defaultJob(),
	}
}

type jobBuilder struct {
	job *v1.Job
}

func (j *jobBuilder) SetName(name string) JobBuilder {
	j.job.ObjectMeta.Name = name
	return j
}

func (j *jobBuilder) SetNamespace(namespace string) JobBuilder {
	j.job.ObjectMeta.Namespace = namespace
	return j
}

func (j *jobBuilder) SetLabels(labels map[string]string) JobBuilder {
	j.job.ObjectMeta.Labels = labels
	return j
}

func (j *jobBuilder) SetCompletions(i int32) JobBuilder {
	j.job.Spec.Completions = &i
	return j
}

func (j *jobBuilder) SetActiveDeadlineSeconds(i int64) JobBuilder {
	j.job.Spec.ActiveDeadlineSeconds = &i
	return j
}

func (j *jobBuilder) SetParallelism(i int32) JobBuilder {
	j.job.Spec.Parallelism = &i
	return j
}

func (j *jobBuilder) SetBackoffLimit(i int32) JobBuilder {
	j.job.Spec.BackoffLimit = &i
	return j
}

func (j *jobBuilder) SetContainers(container *corev1.Container) JobBuilder {
	if j.job.Spec.Template.Spec.Containers == nil {
		j.job.Spec.Template.Spec.Containers = []corev1.Container{*container}
	}

	for i := range j.job.Spec.Template.Spec.Containers {
		if j.job.Spec.Template.Spec.Containers[i].Name == DefaultContainerName {
			j.job.Spec.Template.Spec.Containers[i] = *container
			return j
		}
	}

	j.job.Spec.Template.Spec.Containers = append(j.job.Spec.Template.Spec.Containers, *container)
	return j
}

func (j *jobBuilder) SetTTLSecondsAfterFinished(i int32) JobBuilder {
	ret := i
	j.job.Spec.TTLSecondsAfterFinished = &ret
	return j
}

func (j *jobBuilder) SetSuspend(b bool) JobBuilder {
	j.job.Spec.Suspend = &b
	return j
}

func (j *jobBuilder) SetVolumes(volume *corev1.Volume) JobBuilder {
	if j.job.Spec.Template.Spec.Volumes == nil || len(j.job.Spec.Template.Spec.Volumes) == 0 {
		j.job.Spec.Template.Spec.Volumes = []corev1.Volume{}
	}

	for i := range j.job.Spec.Template.Spec.Volumes {
		if j.job.Spec.Template.Spec.Volumes[i].Name == volume.Name {
			j.job.Spec.Template.Spec.Volumes[i] = *volume
			return j
		}
	}

	j.job.Spec.Template.Spec.Volumes = append(j.job.Spec.Template.Spec.Volumes, *volume)
	return j
}

func (j *jobBuilder) Build() *v1.Job {
	return j.job
}

func defaultJob() *v1.Job {
	return &v1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "shardingsphere-proxy",
		},
		Spec: v1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers:    []corev1.Container{},
					RestartPolicy: corev1.RestartPolicyOnFailure,
				},
			},
		},
	}
}
