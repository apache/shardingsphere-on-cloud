package metrics

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"strconv"
)

const (
	metricsNamespace = "shardingsphere_proxy_operator"
	leaderLabel      = "is_leader"
)

var (
	isLeader = false
)

type LeaderElectionMetric struct {
	elected <-chan struct{}
	status  *prometheus.GaugeVec
}

var _ manager.LeaderElectionRunnable = &LeaderElectionMetric{}

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
