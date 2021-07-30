package workers_manager

import (
	"github.com/topfreegames/maestro/internal/core/monitoring"
)

var (
	currentWorkersGaugeMetric = monitoring.CreateGaugeMetric(&monitoring.MetricOpts{
		Namespace: monitoring.Namespace,
		Subsystem: monitoring.SubsystemWorker,
		Name:      "current_workers",
		Help:      "Current number of alive workers",
		Labels: []string{
			monitoring.LabelScheduler,
		},
	})

	restartedWorkersCounterMetric = monitoring.CreateCounterMetric(&monitoring.MetricOpts{
		Namespace: monitoring.Namespace,
		Subsystem: monitoring.SubsystemWorker,
		Name:      "restarted_workers",
		Help:      "Number of restarted workers",
		Labels: []string{
			monitoring.LabelScheduler,
		},
	})

	workersSyncCounterMetric = monitoring.CreateCounterMetric(&monitoring.MetricOpts{
		Namespace: monitoring.Namespace,
		Subsystem: monitoring.SubsystemWorker,
		Name:      "workers_sync",
		Help:      "Times of the workers sync processes",
		Labels:    []string{},
	})
)

func reportWorkerStart(schedulerName string) {
	currentWorkersGaugeMetric.WithLabelValues(schedulerName).Inc()
}

func reportWorkerStop(schedulerName string) {
	currentWorkersGaugeMetric.WithLabelValues(schedulerName).Dec()
}

func reportWorkerRestart(schedulerName string) {
	restartedWorkersCounterMetric.WithLabelValues(schedulerName).Inc()
}

func reportWorkersSynced() {
	workersSyncCounterMetric.WithLabelValues().Inc()
}