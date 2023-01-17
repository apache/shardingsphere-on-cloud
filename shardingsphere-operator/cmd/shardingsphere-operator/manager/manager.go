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
	"context"
	"flag"
	"os"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/metrics"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/controllers"
	"go.uber.org/zap/zapcore"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
}

type Options struct {
	ctrl.Options
	FeatureGateOptions
}

type FeatureGateOptions struct {
	ComputeNode bool
}

func ParseOptionsFromFlags() *Options {
	opt := &Options{}
	flag.StringVar(&opt.MetricsBindAddress, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&opt.HealthProbeBindAddress, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&opt.LeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.BoolVar(&opt.ComputeNode, "feature-gate-compute-node", false, "Enable support for CustomResourceDefinition ComputeNode.")

	opts := zap.Options{
		Development: true,
		TimeEncoder: zapcore.RFC3339TimeEncoder,
	}

	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	opt.Scheme = scheme
	opt.LeaderElectionID = "shardingsphere.apache.org"

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))
	return opt
}

type Manager struct {
	manager.Manager
}

func New(opts *Options) *Manager {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), opts.Options)
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&controllers.ProxyReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		Log:    mgr.GetLogger(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ShardingSphereProxy")
		os.Exit(1)
	}
	if err = (&controllers.ProxyConfigReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		Log:    mgr.GetLogger(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ShardingSphereProxyServerConfig")
		os.Exit(1)
	}

	if opts.ComputeNode {
		if err = (&controllers.ComputeNodeReconciler{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
			Log:    mgr.GetLogger(),
			// Deployment: controllers.NewDeployment(mgr.GetClient()),
		}).SetupWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "ComputeNode")
			os.Exit(1)
		}
	}

	return &Manager{
		Manager: mgr,
	}
}

func (mgr *Manager) SetHealthzCheck(path string, check healthz.Checker) *Manager {
	if err := mgr.Manager.AddHealthzCheck(path, check); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	return mgr
}

func (mgr *Manager) SetReadyzCheck(path string, check healthz.Checker) *Manager {
	if err := mgr.Manager.AddReadyzCheck(path, check); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}
	return mgr
}

func (mgr *Manager) SetMetrics() *Manager {
	if err := mgr.Add(metrics.NewLeaderElectionMetric(mgr.Elected())); err != nil {
		setupLog.Error(err, "unable to add LeaderElection Metric")
		os.Exit(1)
	}

	return mgr
}

func (mgr *Manager) Start(ctx context.Context) error {
	setupLog.Info("starting operator")
	return mgr.Manager.Start(ctx)
}
