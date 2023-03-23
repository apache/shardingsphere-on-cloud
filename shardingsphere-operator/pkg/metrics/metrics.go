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

package metrics

import (
	"context"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

const (
	metricsNamespace = "shardingsphere_operator"
	leaderLabel      = "is_leader"
)

var (
	isLeader = false
)

// LeaderElectionMetrics represents metrics about leader election
type LeaderElectionMetric struct {
	elected <-chan struct{}
	status  *prometheus.GaugeVec
}

var _ manager.LeaderElectionRunnable = &LeaderElectionMetric{}

// Start this metrics
func (l *LeaderElectionMetric) Start(ctx context.Context) error {
	// Set default label
	l.status.WithLabelValues(strconv.FormatBool(isLeader)).Set(1)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-l.elected:
				// The first, delete old label
				l.status.Delete(prometheus.Labels{leaderLabel: strconv.FormatBool(isLeader)})
				isLeader = true
				// The second, recreate new label
				l.status.WithLabelValues(strconv.FormatBool(isLeader)).Set(1)
				isLeader = false
				return
			}
		}
	}()

	return nil
}

// NeedLeaderElection implements controller-runtime's manager.LeaderElectionRunnable.
func (l *LeaderElectionMetric) NeedLeaderElection() bool {
	return false
}

// NewLeaderElectionMetric creates a new leader election metric
func NewLeaderElectionMetric(elected <-chan struct{}) manager.Runnable {
	isLeaderGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsNamespace,
		Subsystem: "runtime",
		Name:      "is_leader",
		Help:      "This operator pod whether is the leader",
	}, []string{leaderLabel})
	metrics.Registry.MustRegister(isLeaderGauge)

	return &LeaderElectionMetric{
		elected: elected,
		status:  isLeaderGauge,
	}
}
