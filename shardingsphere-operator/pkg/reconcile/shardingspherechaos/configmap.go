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
	"reflect"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/common"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	configExperimental = "experimental.sh"
	configPressure     = "pressure.sh"
	configVerify       = "verify.sh"
)

const (
	DefaultConfigMapName = "ssChaos-configmap"
)

// NewSSConfigMap returns a new ConfigMap
func NewSSConfigMap(chaos *v1alpha1.ShardingSphereChaos) *corev1.ConfigMap {
	cmb := NewSSConfigMapBuilder()
	cmb.SetName(chaos.Name).SetNamespace(chaos.Namespace).SetLabels(chaos.Labels)
	cmb.SetExperimental(chaos.Spec.InjectJob.Experimental).SetPressure(chaos.Spec.InjectJob.Pressure).SetVerify(chaos.Spec.InjectJob.Verify)

	return cmb.Build()
}

// SSConfigMapBuilder is a builder for ConfigMap by ComputeNode
type SSConfigMapBuilder interface {
	common.ConfigMapBuilder

	SetExperimental(v1alpha1.Script) SSConfigMapBuilder
	SetPressure(v1alpha1.Script) SSConfigMapBuilder
	SetVerify(v1alpha1.Script) SSConfigMapBuilder
}

type configmapBuilder struct {
	common.ConfigMapBuilder
	configmap *corev1.ConfigMap
}

// NewSSConfigMapBuilder returns a new SSConfigMapBuilder
func NewSSConfigMapBuilder() SSConfigMapBuilder {
	configMap := defaultConfigMap()

	return &configmapBuilder{
		common.NewCommonConfigMapBuilder(configMap),
		configMap,
	}
}

// SetExperimental sets the experimental command
func (c *configmapBuilder) SetExperimental(s v1alpha1.Script) SSConfigMapBuilder {
	c.configmap.Data[configExperimental] = string(s)
	return c
}

// SetPressure sets the pressure command
func (c *configmapBuilder) SetPressure(s v1alpha1.Script) SSConfigMapBuilder {
	c.configmap.Data[configPressure] = string(s)
	return c
}

// SetVerify sets the verify scripts
func (c *configmapBuilder) SetVerify(s v1alpha1.Script) SSConfigMapBuilder {
	c.configmap.Data[configVerify] = string(s)
	return c
}

// defaultConfigMap returns a ConfigMap filling with default expected values
func defaultConfigMap() *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      DefaultConfigMapName,
			Namespace: "default",
			Labels:    map[string]string{},
		},
		Data: map[string]string{},
	}
}

// UpdateConfigMap returns a new ConfigMap
func UpdateConfigMap(ssChaos *v1alpha1.ShardingSphereChaos, cur *corev1.ConfigMap) *corev1.ConfigMap {
	exp := &corev1.ConfigMap{}
	exp.ObjectMeta = cur.ObjectMeta
	exp.Labels = cur.Labels
	exp.Annotations = cur.Annotations
	now := NewSSConfigMap(ssChaos)
	if reflect.DeepEqual(now.Data, cur.Data) {
		return nil
	}
	exp.Data = now.Data
	return exp
}
