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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	configExperimental = "experimental.sh"
	configPressure     = "pressure.sh"
)

const (
	DefaultConfigMapName = "ssChaos-configmap"
)

func NewSSConfigMap(chaos *v1alpha1.ShardingSphereChaos) *v1.ConfigMap {
	cmb := NewSSConfigMapBuilder()

	cmb.SetName(chaos.Name).SetNamespace(chaos.Namespace).SetLabels(chaos.Labels)

	cmb.SetExperimental(chaos.Spec.InjectJob.Experimental).SetPressure(chaos.Spec.InjectJob.Pressure)

	return cmb.Build()
}

// SSConfigMapBuilder is a builder for ConfigMap by ComputeNode
type SSConfigMapBuilder interface {
	common.ConfigMapBuilder
	SetExperimental(string) SSConfigMapBuilder
	SetPressure(string) SSConfigMapBuilder
}

type configmapBuilder struct {
	common.ConfigMapBuilder
	configmap *v1.ConfigMap
}

func NewSSConfigMapBuilder() SSConfigMapBuilder {
	configMap := defaultConfigMap()
	return &configmapBuilder{
		common.NewCommonConfigMapBuilder(configMap),
		configMap,
	}
}

func (c *configmapBuilder) SetExperimental(cmd string) SSConfigMapBuilder {
	c.configmap.Data[configExperimental] = cmd
	return c
}

func (c *configmapBuilder) SetPressure(cmd string) SSConfigMapBuilder {
	c.configmap.Data[configPressure] = cmd
	return c
}

// defaultConfigMap returns a ConfigMap filling with default expected values
func defaultConfigMap() *v1.ConfigMap {
	return &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      DefaultConfigMapName,
			Namespace: "default",
			Labels:    map[string]string{},
		},
		Data: map[string]string{},
	}
}

func UpdateConfigMap(ssChaos *v1alpha1.ShardingSphereChaos, cur *v1.ConfigMap) *v1.ConfigMap {
	exp := &v1.ConfigMap{}
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
