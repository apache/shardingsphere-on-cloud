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

package manager

import (
	"flag"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/job"
	"strings"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaos"
	batchV1 "k8s.io/api/batch/v1"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/controllers"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/deployment"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/service"
	chaosv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"go.uber.org/zap/zapcore"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(chaosv1alpha1.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(batchV1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

// Options represents common options for the controller
type Options struct {
	ctrl.Options
	FeatureGates string
	ZapOptions   zap.Options
}

// ParseOptionsFromFlags parses options from flags
func ParseOptionsFromCmdFlags() *Options {
	// Declare and initialize the options struct
	opt := &Options{
		Options: ctrl.Options{
			Scheme:           scheme,
			LeaderElectionID: "shardingsphere.apache.org",
		},

		ZapOptions: zap.Options{
			Development: true,
			TimeEncoder: zapcore.RFC3339TimeEncoder,
		},
	}

	// Declaring flags for command-line arguments
	flag.StringVar(&opt.MetricsBindAddress, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&opt.HealthProbeBindAddress, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&opt.LeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.StringVar(&opt.FeatureGates, "feature-gates", "ShardingSphereChaos=true", "A set of key=value pairs that describe feature gates for alpha/experimental features.")

	opt.ZapOptions.BindFlags(flag.CommandLine)
	flag.Parse()
	return opt
}

// ParseFeatureGates parse options from command line to build features
func (opts *Options) ParseFeatureGates() []FeatureGateHandler {
	handlers := []FeatureGateHandler{}
	if len(opts.FeatureGates) == 0 {
		return handlers
	}
	for _, gateVal := range strings.Split(opts.FeatureGates, ",") {
		gate, enable := func() (string, bool) {
			gval := strings.Split(gateVal, "=")
			if len(gval) == 2 {
				return gval[0], gval[1] == "true"
			}
			return "", false
		}()
		if h, ok := featureGatesHandlers[gate]; ok && enable {
			handlers = append(handlers, h)
		}
	}
	return handlers
}

// FeatureGateHandler returns a Manager for the given crd
type FeatureGateHandler func(mgr manager.Manager) error

var featureGatesHandlers = map[string]FeatureGateHandler{
	"ComputeNode": func(mgr manager.Manager) error {
		if err := (&controllers.ComputeNodeReconciler{
			Client:     mgr.GetClient(),
			Scheme:     mgr.GetScheme(),
			Log:        mgr.GetLogger(),
			Deployment: deployment.NewDeployment(mgr.GetClient()),
			Service:    service.NewService(mgr.GetClient()),
			ConfigMap:  configmap.NewConfigMap(mgr.GetClient()),
		}).SetupWithManager(mgr); err != nil {
			logger.Error(err, "unable to create controller", "controller", "ComputeNode")
			return err
		}
		return nil
	},
	"StorageNode": func(mgr manager.Manager) error {
		if err := (&controllers.StorageNodeReconciler{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
			Log:    mgr.GetLogger(),
		}).SetupWithManager(mgr); err != nil {
			logger.Error(err, "unable to create controller", "controller", "StorageNode")
			return err
		}
		return nil
	},
	"ShardingSphereChaos": func(mgr manager.Manager) error {
		if err := (&controllers.ShardingSphereChaosReconciler{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
			Log:    mgr.GetLogger(),
			Chaos:  chaos.NewChaos(mgr.GetClient()),
			Job:    job.NewJob(mgr.GetClient()),
		}).SetupWithManager(mgr); err != nil {
			logger.Error(err, "unable to create controller", "controller", "ShardingSphereChaos")
			return err
		}
		return nil
	},
}
