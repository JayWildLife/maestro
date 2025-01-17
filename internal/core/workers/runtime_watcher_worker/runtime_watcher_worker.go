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

package runtime_watcher_worker

import (
	"context"
	"fmt"
	"sync"

	"github.com/topfreegames/maestro/internal/core/logs"

	"github.com/topfreegames/maestro/internal/core/entities"
	"github.com/topfreegames/maestro/internal/core/entities/game_room"
	"github.com/topfreegames/maestro/internal/core/ports"
	"github.com/topfreegames/maestro/internal/core/workers"
	"go.uber.org/zap"
)

var _ workers.Worker = (*runtimeWatcherWorker)(nil)

const WorkerName = "runtime_watcher"

// runtimeWatcherWorker is a work that will watch for changes on the Runtime
// and apply to keep the Maestro state up-to-date. It is not expected to perform
// any action over the game rooms or the scheduler.
type runtimeWatcherWorker struct {
	scheduler   *entities.Scheduler
	roomManager ports.RoomManager
	// TODO(gabrielcorado): should we access the port directly? do we need to
	// provide the same `Watcher` interface but on the RoomManager?
	runtime    ports.Runtime
	logger     *zap.Logger
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewRuntimeWatcherWorker(scheduler *entities.Scheduler, opts *workers.WorkerOptions) workers.Worker {
	return &runtimeWatcherWorker{
		scheduler:   scheduler,
		roomManager: opts.RoomManager,
		runtime:     opts.Runtime,
		logger:      zap.L().With(zap.String(logs.LogFieldServiceName, WorkerName), zap.String(logs.LogFieldSchedulerName, scheduler.Name)),
	}
}

func (w *runtimeWatcherWorker) Start(ctx context.Context) error {
	watcher, err := w.runtime.WatchGameRoomInstances(ctx, w.scheduler)
	if err != nil {
		return fmt.Errorf("failed to start watcher: %w", err)
	}

	w.ctx, w.cancelFunc = context.WithCancel(ctx)
	defer w.cancelFunc()

	resultChan := watcher.ResultChan()

	workerWaitGroup := &sync.WaitGroup{}

	for i := 0; i < 200; i++ {
		workerWaitGroup.Add(1)
		go func(goroutineNumber int) {
			defer workerWaitGroup.Done()
			goroutineLogger := w.logger.With(zap.Int("goroutine", goroutineNumber))
			goroutineLogger.Info("Starting event processing goroutine")
			for {
				select {
				case event, ok := <-resultChan:
					if !ok {
						w.logger.Warn("resultChan closed, finishing worker goroutine", zap.Error(err))
						return
					}
					err = w.processEvent(w.ctx, event)
					if err != nil {
						w.logger.Warn("failed to process event", zap.Error(err))
						reportEventProcessingStatus(event, false)
					} else {
						reportEventProcessingStatus(event, true)
					}
				case <-w.ctx.Done():
					return
				}
			}
		}(i)
	}

	workerWaitGroup.Wait()
	watcher.Stop()
	return nil
}

func (w *runtimeWatcherWorker) Stop(_ context.Context) {
	if w.cancelFunc != nil {
		w.cancelFunc()
	}
}

func (w *runtimeWatcherWorker) IsRunning() bool {
	return w.ctx != nil && w.ctx.Err() == nil
}

func (w *runtimeWatcherWorker) processEvent(ctx context.Context, event game_room.InstanceEvent) error {
	eventLogger := w.logger.With(zap.String(logs.LogFieldInstanceID, event.Instance.ID))
	switch event.Type {
	case game_room.InstanceEventTypeAdded, game_room.InstanceEventTypeUpdated:
		eventLogger.Info(fmt.Sprintf("processing %s event. Updating rooms instance", event.Type.String()))
		if event.Instance == nil {
			return fmt.Errorf("cannot process event since instance is nil")
		}
		err := w.roomManager.UpdateRoomInstance(ctx, event.Instance)
		if err != nil {
			eventLogger.Error(fmt.Sprintf("failed to process %s event.", event.Type.String()), zap.Error(err))
			return fmt.Errorf("failed to update room instance %s: %w", event.Instance.ID, err)
		}
	case game_room.InstanceEventTypeDeleted:
		eventLogger.Info(fmt.Sprintf("processing %s event. Cleaning Room state", event.Type.String()))
		if event.Instance == nil {
			return fmt.Errorf("cannot process event since instance is nil")
		}
		err := w.roomManager.CleanRoomState(ctx, event.Instance.SchedulerID, event.Instance.ID)
		if err != nil {
			eventLogger.Error(fmt.Sprintf("failed to process %s event.", event.Type.String()), zap.Error(err))
			return fmt.Errorf("failed to clean room %s state: %w", event.Instance.ID, err)
		}
	default:
		return fmt.Errorf("not implemented")
	}
	return nil
}
