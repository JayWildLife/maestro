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

//go:build wireinject
// +build wireinject

package managementapi

import (
	"context"

	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/topfreegames/maestro/internal/api/handlers"
	"github.com/topfreegames/maestro/internal/config"
	"github.com/topfreegames/maestro/internal/core/operations/providers"
	"github.com/topfreegames/maestro/internal/service"
	api "github.com/topfreegames/maestro/pkg/api/v1"
)

func initializeManagementMux(ctx context.Context, conf config.Config) (*runtime.ServeMux, error) {
	wire.Build(
		// ports + adapters
		service.NewClockTime,
		service.NewOperationFlowRedis,
		service.NewOperationStorageRedis,
		service.NewOperationLeaseStorageRedis,
		service.NewSchedulerStoragePg,
		service.NewRoomStorageRedis,
		service.NewSchedulerCacheRedis,

		// scheduler operations
		providers.ProvideDefinitionConstructors,

		// services
		service.NewSchedulerManager,
		service.NewOperationManager,

		// api handlers
		handlers.ProvideSchedulersHandler,
		handlers.ProvideOperationsHandler,
		provideManagementMux,

		// config
		service.NewOperationManagerConfig,
	)

	return &runtime.ServeMux{}, nil
}

func provideManagementMux(ctx context.Context, schedulersHandler *handlers.SchedulersHandler, operationsHandler *handlers.OperationsHandler) *runtime.ServeMux {
	mux := runtime.NewServeMux()
	_ = api.RegisterSchedulersServiceHandlerServer(ctx, mux, schedulersHandler)
	_ = api.RegisterOperationsServiceHandlerServer(ctx, mux, operationsHandler)

	return mux
}
