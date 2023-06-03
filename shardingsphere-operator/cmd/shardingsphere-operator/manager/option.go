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
	sschaos "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaosmesh"
	cloudnativepg "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/cloudnative-pg"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/deployment"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/job"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/service"

	chaosv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	cnpgv1 "github.com/cloudnative-pg/cloudnative-pg/api/v1"
	"github.com/database-mesh/golang-sdk/aws"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
	dbmeshv1alpha1 "github.com/database-mesh/golang-sdk/kubernetes/api/v1alpha1"
	"go.uber.org/zap/zapcore"
	batchV1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientset "k8s.io/client-go/kubernetes"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	scheme *runtime.Scheme
)

func init() {
	scheme = runtime.NewScheme()
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(chaosv1alpha1.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(batchV1.AddToScheme(scheme))
	utilruntime.Must(dbmeshv1alpha1.AddToScheme(scheme))
	utilruntime.Must(cnpgv1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

// Options represents common options for the controller
type Options struct {
	ctrl.Options
	FeatureGates string
	ZapOptions   zap.Options
}

var (
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	AwsRegion          string
)

// ParseOptionsFromCmdFlags parses options from flags
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
	flag.StringVar(&opt.FeatureGates, "feature-gates", "", "A set of key=value pairs that describe feature gates for alpha/experimental features.")
	// aws client options
	flag.StringVar(&AwsAccessKeyID, "aws-access-key-id", "", "The AWS access key ID.")
	flag.StringVar(&AwsSecretAccessKey, "aws-secret-access-key", "", "The AWS secret access key.")
	flag.StringVar(&AwsRegion, "aws-region", "", "The AWS region.")

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
			Deployment: deployment.NewDeploymentClient(mgr.GetClient()),
			Service:    service.NewServiceClient(mgr.GetClient()),
			ConfigMap:  configmap.NewConfigMapClient(mgr.GetClient()),
		}).SetupWithManager(mgr); err != nil {
			logger.Error(err, "unable to create controller", "controller", "ComputeNode")
			return err
		}

		return nil
	},
	"StorageNode": func(mgr manager.Manager) error {
		reconciler := &controllers.StorageNodeReconciler{
			Client:   mgr.GetClient(),
			Scheme:   mgr.GetScheme(),
			Log:      mgr.GetLogger(),
			Recorder: mgr.GetEventRecorderFor(controllers.StorageNodeControllerName),
			Service:  service.NewServiceClient(mgr.GetClient()),
			CNPG:     cloudnativepg.NewCloudNativePGClient(mgr.GetClient()),
		}

		// init aws client if aws credentials are provided
		if AwsRegion != "" && AwsAccessKeyID != "" && AwsSecretAccessKey != "" {
			sess := aws.NewSessions().SetCredential(AwsRegion, AwsAccessKeyID, AwsSecretAccessKey).Build()
			reconciler.AwsRDS = rds.NewService(sess[AwsRegion])
		}

		if err := reconciler.SetupWithManager(mgr); err != nil {
			logger.Error(err, "unable to create controller", "controller", "StorageNode")
			return err
		}
		return nil
	},
	"ShardingSphereChaos": func(mgr manager.Manager) error {
		clientset, err := clientset.NewForConfig(mgr.GetConfig())
		if err != nil {
			return err
		}
		if err := (&controllers.ShardingSphereChaosReconciler{
			Client:    mgr.GetClient(),
			Scheme:    mgr.GetScheme(),
			Log:       mgr.GetLogger(),
			Chaos:     sschaos.NewChaos(mgr.GetClient()),
			Job:       job.NewJob(mgr.GetClient()),
			ExecCtrls: make([]*controllers.ExecCtrl, 0),
			ConfigMap: configmap.NewConfigMapClient(mgr.GetClient()),
			Events:    mgr.GetEventRecorderFor("shardingsphere-chaos-controller"),
			ClientSet: clientset,
		}).SetupWithManager(mgr); err != nil {
			logger.Error(err, "unable to create controller", "controller", "ShardingSphereChaos")
			return err
		}
		return nil
	},
}
