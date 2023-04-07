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

package job

import (
	"context"
	batchV1 "k8s.io/api/batch/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewJob creates a new Job
func NewJob(c client.Client) Job {
	return JobClient{
		JobGetter: jobGetter{
			Client: c,
		},
		JobSetter: jobSetter{
			Client: c,
		},
	}
}

// Job interface contains setter and getter
type Job interface {
	JobGetter
	JobSetter
}

type JobClient struct {
	JobGetter
	JobSetter
}

// JobGetter get Job from different parameters
type JobGetter interface {
	GetByNamespacedName(context.Context, types.NamespacedName) (*batchV1.Job, error)
}

type jobGetter struct {
	client.Client
}

// GetByNamespacedName returns Job from given namespaced name
func (jg jobGetter) GetByNamespacedName(ctx context.Context, namespacedName types.NamespacedName) (*batchV1.Job, error) {
	dp := &batchV1.Job{}
	if err := jg.Get(ctx, namespacedName, dp); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	} else {
		return dp, nil
	}
}

// JobMapGetter get Job from different parameters
type JobSetter interface {
}

type jobSetter struct {
	client.Client
}
