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
	"golang.org/x/mod/semver"
)

var (
	supportedShardingSphereVersion = []string{"5.3.0", "5.3.1", "5.3.2", "5.4.0"}
)

func IsSupportedShardingSphereVersion(v string) bool {
	if !semver.IsValid(v) {
		return false
	}

	for i := range supportedShardingSphereVersion {
		if supportedShardingSphereVersion[i] == v {
			return true
		}
	}

	return false
}

func VersionBetween(version, left, right string) bool {
	if !IsSupportedShardingSphereVersion(version) || !IsSupportedShardingSphereVersion(left) || !IsSupportedShardingSphereVersion(right) {
		return false
	}

	if semver.Compare(version, left) >= 0 && semver.Compare(version, right) <= 0 {
		return true
	}
	return false
}

func VersionGreaterAndEqualThan(version, target string) bool {
	if !IsSupportedShardingSphereVersion(target) || !IsSupportedShardingSphereVersion(version) {
		return false
	}

	return semver.Compare(version, target) >= 0
}

func VersionExactEqualTo(version, target string) bool {
	if !IsSupportedShardingSphereVersion(target) || !IsSupportedShardingSphereVersion(version) {
		return false
	}

	return semver.Compare(version, target) == 0
}
