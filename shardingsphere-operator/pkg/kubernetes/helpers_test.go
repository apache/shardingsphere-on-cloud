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

package kubernetes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsSupportedShardingSphereVersion(t *testing.T) {
	cases := []struct {
		version string
		exp     bool
	}{
		{
			version: "v5.3.0",
			exp:     false,
		},
		{
			version: "5.3.0",
			exp:     true,
		},
		{
			version: "5.3.2",
			exp:     true,
		},
	}

	for _, c := range cases {
		act := IsSupportedShardingSphereVersion(c.version)
		assert.Equal(t, c.exp, act, fmt.Sprintf("%s should be valid", c.version))
	}
}

func Test_VersionBetween(t *testing.T) {
	cases := []struct {
		version      string
		versionRange []string
		exp          bool
	}{
		{
			version:      "5.3.0",
			exp:          true,
			versionRange: []string{"5.3.0", "5.3.2"},
		},
	}

	for _, c := range cases {
		act := VersionBetween(c.version, c.versionRange[0], c.versionRange[1])
		assert.Equal(t, c.exp, act, fmt.Sprintf("%s should be valid", c.version))
	}
}
