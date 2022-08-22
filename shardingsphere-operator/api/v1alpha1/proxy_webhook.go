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

package v1alpha1

import (
	shardingwebhook "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/webhook"

	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var shardingsphereproxylog = logf.Log.WithName("shardingsphereproxy-resource")

func (r *ShardingSphereProxy) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return shardingwebhook.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/apis/admission.shardingsphere.sphere-ex.com/v1alpha1/mutate-shardingsphere-sphere-ex-com-v1alpha1-proxy,mutating=true,failurePolicy=fail,sideEffects=None,groups=shardingsphere.sphere-ex.com,resources=shardingsphereproxies,verbs=create;update,versions=v1alpha1,name=mproxy.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &ShardingSphereProxy{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *ShardingSphereProxy) Default() {
	shardingsphereproxylog.Info("default", "name", r.Name)
	if r.Spec.StartupProbe == nil {
		r.Spec.StartupProbe = &v1.Probe{
			ProbeHandler: v1.ProbeHandler{
				TCPSocket: &v1.TCPSocketAction{
					Port: intstr.FromInt(int(r.Spec.Port)),
				},
			},
			PeriodSeconds:    5,
			FailureThreshold: 2,
		}
	}
	if r.Spec.ReadinessProbe == nil {
		r.Spec.ReadinessProbe = &v1.Probe{
			ProbeHandler: v1.ProbeHandler{
				TCPSocket: &v1.TCPSocketAction{
					Port: intstr.FromInt(int(r.Spec.Port)),
				},
			},
			PeriodSeconds: 10,
		}
	}
	if r.Spec.LivenessProbe == nil {
		r.Spec.LivenessProbe = &v1.Probe{
			ProbeHandler: v1.ProbeHandler{
				TCPSocket: &v1.TCPSocketAction{
					Port: intstr.FromInt(int(r.Spec.Port)),
				},
			},

			PeriodSeconds: 10,
		}
	}
	if r.Spec.Resources == nil {
		// The application for resources is based on the upper limit of cpu: 4core memory: 2Gi to apply for computing resources.
		// The cpu is applied according to 10%, and the memory is applied according to 80%
		cpuLimits, _ := resource.ParseQuantity("4")
		memoryLimits, _ := resource.ParseQuantity("2Gi")
		cpuRequest, _ := resource.ParseQuantity("400m")
		memoryRequest, _ := resource.ParseQuantity("1600Mbi")
		r.Spec.Resources = &v1.ResourceRequirements{
			Limits: v1.ResourceList{
				"cpu":    cpuLimits,
				"memory": memoryLimits,
			},
			Requests: v1.ResourceList{
				"cpu":    cpuRequest,
				"memory": memoryRequest,
			},
		}
	}
}

// +kubebuilder:webhook:path=/apis/admission.shardingsphere.sphere-ex.com/v1alpha1/validate-shardingsphere-sphere-ex-com-v1alpha1-proxy,mutating=false,failurePolicy=fail,sideEffects=None,groups=shardingsphere.sphere-ex.com,resources=proxies,verbs=create;update,versions=v1alpha1,name=vproxy.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &ShardingSphereProxy{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ShardingSphereProxy) ValidateCreate() error {
	shardingsphereproxylog.Info("validate create", "name", r.Name)
	err := r.validateService()
	if err != nil {
		return err
	}
	err = r.validateInstance()
	if err != nil {
		return err
	}
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ShardingSphereProxy) ValidateUpdate(old runtime.Object) error {
	shardingsphereproxylog.Info("validate update", "name", r.Name)
	err := r.validateService()
	if err != nil {
		return err
	}
	err = r.validateInstance()
	if err != nil {
		return err
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ShardingSphereProxy) ValidateDelete() error {
	return nil
}
func (r *ShardingSphereProxy) validateService() error {
	var allErrs field.ErrorList
	field.NewPath("spec").Child("serviceType")
	if r.Spec.ServiceType.NodePort != 0 && r.Spec.ServiceType.Type != v1.ServiceTypeNodePort {
		allErrs = append(allErrs, field.Invalid(field.NewPath("spec").Child("serviceType"), r.Spec.ServiceType, errors.New("nodePort: Forbidden: may not be used when `type` is 'ClusterIP'").Error()))
		return apierrors.NewInvalid(schema.GroupKind{
			Group: "shardingsphere.sphere-ex.com",
			Kind:  "ShardingSphereProxy",
		}, r.Name, allErrs)
	}
	return nil
}

func (r *ShardingSphereProxy) validateInstance() error {
	var allErrs field.ErrorList
	if r.Spec.AutomaticScaling != nil && r.Spec.AutomaticScaling.MaxInstance < 1 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("spec").Child("MaxInstance"), r.Spec.ServiceType, errors.New("If automatic scaling is enabled, the number of instances must be filled in.").Error()))
		return apierrors.NewInvalid(schema.GroupKind{
			Group: "shardingsphere.sphere-ex.com",
			Kind:  "ShardingSphereProxy",
		}, r.Name, allErrs)
	}
	return nil
}
