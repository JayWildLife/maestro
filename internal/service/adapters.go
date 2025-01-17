// MIT License
//
// Copyright (c) 2021 TFG Co
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package service

import (
	"fmt"
	"time"

	"github.com/topfreegames/maestro/internal/core/services/scheduler_manager"

	"github.com/topfreegames/maestro/internal/core/entities/autoscaling"
	"github.com/topfreegames/maestro/internal/core/operations/add_rooms"
	"github.com/topfreegames/maestro/internal/core/operations/healthcontroller"
	"github.com/topfreegames/maestro/internal/core/operations/remove_rooms"

	operationadapters "github.com/topfreegames/maestro/internal/adapters/operation"

	eventsadapters "github.com/topfreegames/maestro/internal/adapters/events"

	scheduleradapters "github.com/topfreegames/maestro/internal/adapters/scheduler"

	autoscalerports "github.com/topfreegames/maestro/internal/core/ports/autoscaler"

	"github.com/topfreegames/maestro/internal/core/operations"
	"github.com/topfreegames/maestro/internal/core/services/autoscaler"
	"github.com/topfreegames/maestro/internal/core/services/autoscaler/policies/roomoccupancy"
	"github.com/topfreegames/maestro/internal/core/services/operation_manager"
	"github.com/topfreegames/maestro/internal/core/services/room_manager"

	"github.com/go-pg/pg"
	"github.com/go-redis/redis/v8"
	clockTime "github.com/topfreegames/maestro/internal/adapters/clock/time"
	instanceStorageRedis "github.com/topfreegames/maestro/internal/adapters/instance_storage/redis"
	portAllocatorRandom "github.com/topfreegames/maestro/internal/adapters/port_allocator/random"
	roomStorageRedis "github.com/topfreegames/maestro/internal/adapters/room_storage/redis"
	kubernetesRuntime "github.com/topfreegames/maestro/internal/adapters/runtime/kubernetes"
	"github.com/topfreegames/maestro/internal/config"
	"github.com/topfreegames/maestro/internal/core/entities"
	"github.com/topfreegames/maestro/internal/core/ports"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// configurations paths for the adapters
const (
	// Kubernetes runtime
	runtimeKubernetesMasterURLPath  = "adapters.runtime.kubernetes.masterUrl"
	runtimeKubernetesKubeconfigPath = "adapters.runtime.kubernetes.kubeconfig"
	runtimeKubernetesInClusterPath  = "adapters.runtime.kubernetes.inCluster"
	// Redis operation storage
	operationStorageRedisURLPath      = "adapters.operationStorage.redis.url"
	operationLeaseStorageRedisURLPath = "adapters.operationLeaseStorage.redis.url"
	// Redis room storage
	roomStorageRedisURLPath = "adapters.roomStorage.redis.url"
	// Redis scheduler cache
	schedulerCacheRedisURLPath = "adapters.schedulerCache.redis.url"
	// Redis instance storage
	instanceStorageRedisURLPath      = "adapters.instanceStorage.redis.url"
	instanceStorageRedisScanSizePath = "adapters.instanceStorage.redis.scanSize"
	// Redis operation flow
	operationFlowRedisURLPath = "adapters.operationFlow.redis.url"
	// Redis configs
	redisPoolSizePath = "adapters.redis.poolSize"
	// Random port allocator
	portAllocatorRandomRangePath = "adapters.portAllocator.random.range"
	// Postgres scheduler storage
	schedulerStoragePostgresURLPath = "adapters.schedulerStorage.postgres.url"

	// operation TTL
	operationsTTLPath = "workers.redis.operationsTtl"
)

// NewSchedulerManager instantiates a new scheduler manager.
func NewSchedulerManager(schedulerStorage ports.SchedulerStorage, schedulerCache ports.SchedulerCache, operationManager ports.OperationManager, roomStorage ports.RoomStorage) ports.SchedulerManager {
	return scheduler_manager.NewSchedulerManager(schedulerStorage, schedulerCache, operationManager, roomStorage)
}

// NewOperationManager instantiates a new operation manager
func NewOperationManager(flow ports.OperationFlow, storage ports.OperationStorage, operationDefinitionConstructors map[string]operations.DefinitionConstructor, leaseStorage ports.OperationLeaseStorage, config operation_manager.OperationManagerConfig, schedulerStorage ports.SchedulerStorage) ports.OperationManager {
	return operation_manager.New(flow, storage, operationDefinitionConstructors, leaseStorage, config, schedulerStorage)
}

// NewRoomManager instantiates a room manager.
func NewRoomManager(clock ports.Clock, portAllocator ports.PortAllocator, roomStorage ports.RoomStorage, instanceStorage ports.GameRoomInstanceStorage, runtime ports.Runtime, eventsService ports.EventsService, config room_manager.RoomManagerConfig) ports.RoomManager {
	return room_manager.New(clock, portAllocator, roomStorage, instanceStorage, runtime, eventsService, config)
}

// NewEventsForwarder instantiates GRPC as events forwarder.
func NewEventsForwarder(c config.Config) (ports.EventsForwarder, error) {
	forwarderGrpc := eventsadapters.NewForwarderClient()
	return eventsadapters.NewEventsForwarder(forwarderGrpc), nil
}

// NewRuntimeKubernetes instantiates kubernetes as runtime.
func NewRuntimeKubernetes(c config.Config) (ports.Runtime, error) {
	var masterURL string
	var kubeConfigPath string

	inCluster := c.GetBool(runtimeKubernetesInClusterPath)
	if !inCluster {
		masterURL = c.GetString(runtimeKubernetesMasterURLPath)
		kubeConfigPath = c.GetString(runtimeKubernetesKubeconfigPath)
	}

	clientSet, err := createKubernetesClient(masterURL, kubeConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Kubernetes runtime: %w", err)
	}

	return kubernetesRuntime.New(clientSet), nil
}

// NewOperationStorageRedis instantiates redis as operation storage.
func NewOperationStorageRedis(clock ports.Clock, c config.Config) (ports.OperationStorage, error) {
	client, err := createRedisClient(c, c.GetString(operationStorageRedisURLPath))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Redis operation storage: %w", err)
	}

	operationsTTLPathMap := map[operationadapters.Definition]time.Duration{
		healthcontroller.OperationName: c.GetDuration(operationsTTLPath),
		add_rooms.OperationName:        c.GetDuration(operationsTTLPath),
		remove_rooms.OperationName:     c.GetDuration(operationsTTLPath),
	}

	return operationadapters.NewRedisOperationStorage(client, clock, operationsTTLPathMap), nil
}

// NewOperationLeaseStorageRedis instantiates redis as operation lease storage.
func NewOperationLeaseStorageRedis(clock ports.Clock, c config.Config) (ports.OperationLeaseStorage, error) {
	client, err := createRedisClient(c, c.GetString(operationLeaseStorageRedisURLPath))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Redis operation lease storage: %w", err)
	}

	return operationadapters.NewRedisOperationLeaseStorage(client, clock), nil
}

// NewRoomStorageRedis instantiates redis as room storage.
func NewRoomStorageRedis(c config.Config) (ports.RoomStorage, error) {
	client, err := createRedisClient(c, c.GetString(roomStorageRedisURLPath))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Redis room storage: %w", err)
	}

	return roomStorageRedis.NewRedisStateStorage(client), nil
}

// NewGameRoomInstanceStorageRedis instantiates redis as instance storage.
func NewGameRoomInstanceStorageRedis(c config.Config) (ports.GameRoomInstanceStorage, error) {
	client, err := createRedisClient(c, c.GetString(instanceStorageRedisURLPath))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Redis instance storage: %w", err)
	}

	return instanceStorageRedis.NewRedisInstanceStorage(client, c.GetInt(instanceStorageRedisScanSizePath)), nil
}

// NewSchedulerCacheRedis instantiates redis as scheduler cache.
func NewSchedulerCacheRedis(c config.Config) (ports.SchedulerCache, error) {
	client, err := createRedisClient(c, c.GetString(schedulerCacheRedisURLPath))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Redis scheduler cache: %w", err)
	}

	return scheduleradapters.NewRedisSchedulerCache(client), nil
}

// NewClockTime instantiates a new clock.
func NewClockTime() ports.Clock {
	return clockTime.NewClock()
}

// NewPortAllocatorRandom instantiates a new port allocator.
func NewPortAllocatorRandom(c config.Config) (ports.PortAllocator, error) {
	portRange, err := entities.ParsePortRange(c.GetString(portAllocatorRandomRangePath))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize random port allocator: %w", err)
	}

	return portAllocatorRandom.NewRandomPortAllocator(portRange), nil
}

// NewSchedulerStoragePg instanteates a postgres connection as scheduler storage.
func NewSchedulerStoragePg(c config.Config) (ports.SchedulerStorage, error) {
	opts, err := connectToPostgres(GetSchedulerStoragePostgresURL(c))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize postgres scheduler storage: %w", err)
	}

	return scheduleradapters.NewSchedulerStorage(opts), nil
}

// GetSchedulerStoragePostgresURL get scheduler storage postgres URL.
func GetSchedulerStoragePostgresURL(c config.Config) string {
	return c.GetString(schedulerStoragePostgresURLPath)
}

func createRedisClient(c config.Config, url string) (*redis.Client, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("invalid redis URL: %w", err)
	}
	opts.PoolSize = c.GetInt(redisPoolSizePath)
	return redis.NewClient(opts), nil
}

// NewPolicyMap instantiates a new policy to be used by autoscaler expecting a room storage as parameter.
func NewPolicyMap(roomStorage ports.RoomStorage) autoscaler.PolicyMap {
	return autoscaler.PolicyMap{
		autoscaling.RoomOccupancy: roomoccupancy.NewPolicy(roomStorage),
	}
}

// NewAutoscaler instantiates  a new autoscaler expecting a Policy Map as parameter.
func NewAutoscaler(policies autoscaler.PolicyMap) autoscalerports.Autoscaler {
	return autoscaler.NewAutoscaler(policies)
}

// NewOperationFlowRedis instantiates a new operation flow using redis as backend.
func NewOperationFlowRedis(c config.Config) (ports.OperationFlow, error) {
	client, err := createRedisClient(c, c.GetString(operationFlowRedisURLPath))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Redis operation storage: %w", err)
	}

	return operationadapters.NewRedisOperationFlow(client), nil
}

func connectToPostgres(url string) (*pg.Options, error) {
	opts, err := pg.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("invalid postgres URL: %w", err)
	}

	return opts, nil
}

func createKubernetesClient(masterURL, kubeconfigPath string) (kubernetes.Interface, error) {
	// NOTE: if neither masterURL or kubeconfigPath are not passed, this will
	// fallback to in cluster config.
	kubeconfig, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to construct kubernetes config: %w", err)
	}

	client, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes client: %w", err)
	}

	return client, nil
}
