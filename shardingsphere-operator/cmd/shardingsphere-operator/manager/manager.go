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
	"os"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/deployment"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/service"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/metrics"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/controllers"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	logger = ctrl.Log.WithName("setup")
)

type Manager struct {
	manager.Manager
}

func SetupWithOptions(opts *Options) *Manager {
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts.ZapOptions)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), opts.Options)
	if err != nil {
		logger.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&controllers.ProxyReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		Log:    mgr.GetLogger(),
	}).SetupWithManager(mgr); err != nil {
		logger.Error(err, "unable to create controller", "controller", "ShardingSphereProxy")
		os.Exit(1)
	}
	if err = (&controllers.ProxyConfigReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		Log:    mgr.GetLogger(),
	}).SetupWithManager(mgr); err != nil {
		logger.Error(err, "unable to create controller", "controller", "ShardingSphereProxyServerConfig")
		os.Exit(1)
	}

	if opts.ComputeNode {
		if err = (&controllers.ComputeNodeReconciler{
			Client:     mgr.GetClient(),
			Scheme:     mgr.GetScheme(),
			Log:        mgr.GetLogger(),
			Deployment: deployment.NewDeployment(mgr.GetClient()),
			Service:    service.NewService(mgr.GetClient()),
			ConfigMap:  configmap.NewConfigMap(mgr.GetClient()),
		}).SetupWithManager(mgr); err != nil {
			logger.Error(err, "unable to create controller", "controller", "ComputeNode")
			os.Exit(1)
		}
	}

	return &Manager{
		Manager: mgr,
	}
}

func (mgr *Manager) SetHealthzCheck(path string, check healthz.Checker) *Manager {
	if err := mgr.Manager.AddHealthzCheck(path, check); err != nil {
		logger.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	return mgr
}

func (mgr *Manager) SetReadyzCheck(path string, check healthz.Checker) *Manager {
	if err := mgr.Manager.AddReadyzCheck(path, check); err != nil {
		logger.Error(err, "unable to set up ready check")
		os.Exit(1)
	}
	return mgr
}

func (mgr *Manager) SetMetrics() *Manager {
	if err := mgr.Add(metrics.NewLeaderElectionMetric(mgr.Elected())); err != nil {
		logger.Error(err, "unable to add LeaderElection Metric")
		os.Exit(1)
	}

	return mgr
}

func (mgr *Manager) Start(ctx context.Context) error {
	logger.Info("starting operator")
	return mgr.Manager.Start(ctx)
}
