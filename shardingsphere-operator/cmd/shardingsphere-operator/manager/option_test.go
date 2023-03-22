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
	"os"
	"testing"
)

func Test_ParseFeatureGates(t *testing.T) {
	testCases := []struct {
		desc        string
		opts        Options
		expectedLen int
	}{
		{
			desc:        "Returns empty handlers if feature gates are not provided",
			opts:        Options{},
			expectedLen: 0,
		},
		{
			desc: "Returns correct handlers if feature gates are provided",
			opts: Options{
				FeatureGates: "ComputeNode=true,Cluster=false,StorageNode=false",
			},
			expectedLen: 1,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			handlers := tC.opts.ParseFeatureGates()
			if len(handlers) != tC.expectedLen {
				t.Errorf("Expected %d feature gate handlers, but got %d", tC.expectedLen, len(handlers))
			}
		})
	}
}

func Test_ParseOptionsFromCmdFlags(t *testing.T) {
	// Create a mock set of command line arguments
	os.Args = []string{
		"app",
		"--metrics-bind-address=localhost:8888",
		"--health-probe-bind-address=localhost:9999",
		"--leader-elect=true",
		"--feature-gates=foo=true,bar=false",
	}

	// Call the function being tested
	opt := ParseOptionsFromCmdFlags()

	// Validate the returned options
	if opt.MetricsBindAddress != "localhost:8888" {
		t.Errorf("Expected opt.MetricsBindAddress to equal 'localhost:8888', but got '%s'", opt.MetricsBindAddress)
	}

	if opt.HealthProbeBindAddress != "localhost:9999" {
		t.Errorf("Expected opt.HealthProbeBindAddress to equal 'localhost:9999', but got '%s'", opt.HealthProbeBindAddress)
	}

	if !opt.LeaderElection {
		t.Errorf("Expected opt.LeaderElection to equal 'true', but got '%v'", opt.LeaderElection)
	}

	if opt.FeatureGates != "foo=true,bar=false" {
		t.Errorf("Expected opt.FeatureGates to equal 'foo=true,bar=false', but got '%s'", opt.FeatureGates)
	}
}
