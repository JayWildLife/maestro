// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package metricsreporter

import (
	"github.com/google/wire"
	"github.com/topfreegames/maestro/internal/config"
	"github.com/topfreegames/maestro/internal/core/services/workers_manager"
	"github.com/topfreegames/maestro/internal/core/workers"
	config2 "github.com/topfreegames/maestro/internal/core/workers/config"
	"github.com/topfreegames/maestro/internal/core/workers/metricsreporter"
	"github.com/topfreegames/maestro/internal/service"
)

// Injectors from wire.go:

func initializeMetricsReporter(c config.Config) (*workers_manager.WorkersManager, error) {
	workerBuilder := provideMetricsReporterBuilder()
	schedulerStorage, err := service.NewSchedulerStoragePg(c)
	if err != nil {
		return nil, err
	}
	roomStorage, err := service.NewRoomStorageRedis(c)
	if err != nil {
		return nil, err
	}
	gameRoomInstanceStorage, err := service.NewGameRoomInstanceStorageRedis(c)
	if err != nil {
		return nil, err
	}
	metricsReporterConfig := provideMetricsReporterConfig(c)
	workerOptions := &workers.WorkerOptions{
		RoomStorage:           roomStorage,
		InstanceStorage:       gameRoomInstanceStorage,
		MetricsReporterConfig: metricsReporterConfig,
	}
	workersManager := workers_manager.NewWorkersManager(workerBuilder, c, schedulerStorage, workerOptions)
	return workersManager, nil
}

// wire.go:

func provideMetricsReporterBuilder() *workers.WorkerBuilder {
	return &workers.WorkerBuilder{
		Func:          metricsreporter.NewMetricsReporterWorker,
		ComponentName: metricsreporter.WorkerName,
	}
}

func provideMetricsReporterConfig(c config.Config) *config2.MetricsReporterConfig {
	return &config2.MetricsReporterConfig{MetricsReporterIntervalMillis: c.GetDuration("reporter.metrics.intervalMillis")}

}

var WorkerOptionsSet = wire.NewSet(service.NewRoomStorageRedis, service.NewGameRoomInstanceStorageRedis, provideMetricsReporterConfig, wire.Struct(new(workers.WorkerOptions), "RoomStorage", "InstanceStorage", "MetricsReporterConfig"))
