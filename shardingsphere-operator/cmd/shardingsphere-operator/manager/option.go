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
	"strings"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/controllers"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/deployment"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/service"
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
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
}

type Options struct {
	ctrl.Options
	FeatureGates string
	ZapOptions   zap.Options
}

func ParseOptionsFromCmdFlags() *Options {
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

	flag.StringVar(&opt.MetricsBindAddress, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&opt.HealthProbeBindAddress, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&opt.LeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.StringVar(&opt.FeatureGates, "feature-gates", "", "A set of key=value pairs that describe feature gates for alpha/experimental features.")

	opt.ZapOptions.BindFlags(flag.CommandLine)

	flag.Parse()

	return opt
}

func (opts *Options) ParseFeatureGates() []FeatureGateHandler {
	handlers := []FeatureGateHandler{}
	if len(opts.FeatureGates) == 0 {
		return handlers
	}
	if gatesVal := strings.Split(opts.FeatureGates, ","); len(gatesVal) > 0 {
		for i := range gatesVal {
			gate, enable := func() (string, bool) {
				if gval := strings.Split(gatesVal[i], "="); len(gval) == 2 {
					return gval[0], gval[1] == "true"
				}
				return "", false
			}()

			if h, ok := featureGatesHandlers[gate]; ok && enable {
				handlers = append(handlers, h)
			}
		}
	}
	return handlers
}

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
}
