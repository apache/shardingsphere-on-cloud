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

package main

import (
	"os"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/cmd/shardingsphere-operator/manager"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
)

func main() {
	if err := manager.SetupWithOptions(manager.ParseOptionsFromCmdFlags()).
		SetHealthzCheck("healthz", healthz.Ping).
		SetReadyzCheck("readyz", healthz.Ping).
		SetMetrics().
		Start(ctrl.SetupSignalHandler()); err != nil {
		os.Exit(1)
	}
}
