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

package e2e

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/controllers"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes"
	mockChaos "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/chaosmesh/mocks"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/configmap"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/kubernetes/service"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/reconcile/computenode"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes/scheme"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	cfg       *rest.Config
	k8sClient client.Client // You'll be using this client in your tests.
	testEnv   *envtest.Environment
	ctx       context.Context
	cancel    context.CancelFunc
	err       error
	mockchaos *mockChaos.MockChaos
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	ctx, cancel = context.WithCancel(context.TODO())

	By("bootstrapping test environment for controllers")

	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join("..", "..", "config", "crd", "bases"),
		},
		ErrorIfCRDPathMissing: true,
		CRDInstallOptions: envtest.CRDInstallOptions{
			MaxTime: 60 * time.Second,
		},
	}

	// cfg is defined in this file globally.
	cfg, err = testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())

	Expect(v1alpha1.AddToScheme(scheme.Scheme)).NotTo(HaveOccurred())
	Expect(clientgoscheme.AddToScheme(scheme.Scheme)).NotTo(HaveOccurred())
	//+kubebuilder:scaffold:scheme

	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClient).NotTo(BeNil())

	k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{})
	Expect(err).ToNot(HaveOccurred())
	// print k8sManager Options
	err = (&controllers.StorageNodeReconciler{
		Client:             k8sManager.GetClient(),
		Scheme:             k8sManager.GetScheme(),
		Log:                ctrl.Log.WithName("controllers").WithName("StorageNode"),
		Recorder:           k8sManager.GetEventRecorderFor("StorageNode"),
		AwsRegion:          "AwsRegion",
		AwsAccessKeyID:     "AwsAccessKeyID",
		AwsSecretAccessKey: "AwsSecretAccessKey",
		Service:            service.NewServiceClient(k8sManager.GetClient()),
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&controllers.ComputeNodeReconciler{
		Client:    k8sManager.GetClient(),
		Scheme:    k8sManager.GetScheme(),
		Log:       logf.Log.WithName("controllers").WithName("ComputeNode"),
		Builder:   computenode.NewBuilder(),
		Resources: kubernetes.NewResources(k8sManager.GetClient()),
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	ctl := gomock.NewController(GinkgoT())
	mockchaos = mockChaos.NewMockChaos(ctl)

	err = (&controllers.ChaosReconciler{
		Client:    k8sManager.GetClient(),
		Scheme:    k8sManager.GetScheme(),
		Log:       logf.Log,
		Events:    k8sManager.GetEventRecorderFor("chaos-controller"),
		Chaos:     mockchaos,
		ExecCtrls: make([]*controllers.ExecCtrl, 0),
		ConfigMap: configmap.NewConfigMapClient(k8sManager.GetClient()),
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	go func() {
		defer GinkgoRecover()
		err = k8sManager.Start(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to run manager")
	}()
})

var _ = AfterSuite(func() {
	cancel()
	By("tearing down the test environment for controllers")
	err := testEnv.Stop()
	Expect(err).NotTo(HaveOccurred())
})
