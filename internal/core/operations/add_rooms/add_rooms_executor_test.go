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

package add_rooms

import (
	"time"

	mockeventsservice "github.com/topfreegames/maestro/internal/core/services/interfaces/mock/events_service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	clock_mock "github.com/topfreegames/maestro/internal/adapters/clock/mock"
	instance_storage_mock "github.com/topfreegames/maestro/internal/adapters/instance_storage/mock"
	port_allocator_mock "github.com/topfreegames/maestro/internal/adapters/port_allocator/mock"

	runtime_mock "github.com/topfreegames/maestro/internal/adapters/runtime/mock"
	schedulerStorageMock "github.com/topfreegames/maestro/internal/adapters/scheduler_storage/mock"
	"github.com/topfreegames/maestro/internal/core/entities"
	"github.com/topfreegames/maestro/internal/core/entities/game_room"
	"github.com/topfreegames/maestro/internal/core/entities/operation"
	"github.com/topfreegames/maestro/internal/core/ports/errors"
	mockports "github.com/topfreegames/maestro/internal/core/ports/mock"
	"github.com/topfreegames/maestro/internal/core/services/room_manager"

	"context"
	"testing"
)

func TestAddRoomsExecutor_Execute(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	clockMock := clock_mock.NewFakeClock(time.Now())
	portAllocatorMock := port_allocator_mock.NewMockPortAllocator(mockCtrl)
	roomStorageMock := mockports.NewMockRoomStorage(mockCtrl)
	instanceStorageMock := instance_storage_mock.NewMockGameRoomInstanceStorage(mockCtrl)
	runtimeMock := runtime_mock.NewMockRuntime(mockCtrl)
	eventsForwarderService := mockeventsservice.NewMockEventsService(mockCtrl)
	schedulerStorage := schedulerStorageMock.NewMockSchedulerStorage(mockCtrl)
	config := room_manager.RoomManagerConfig{RoomInitializationTimeout: time.Millisecond * 1000}
	roomStorageStatusWatcher := mockports.NewMockRoomStorageStatusWatcher(mockCtrl)

	definition := AddRoomsDefinition{Amount: 10}

	container1 := game_room.Container{
		Name: "container1",
		Ports: []game_room.ContainerPort{
			{Protocol: "tcp"},
		},
	}

	container2 := game_room.Container{
		Name: "container2",
		Ports: []game_room.ContainerPort{
			{Protocol: "udp"},
		},
	}

	scheduler := entities.Scheduler{
		Name: "zooba_blue:1.0.0",
		Spec: game_room.Spec{
			Version:    "1.0.0",
			Containers: []game_room.Container{container1, container2},
		},
		PortRange: nil,
	}

	gameRoom := game_room.GameRoom{
		ID:          "game-1",
		SchedulerID: "zooba_blue:1.0.0",
		Version:     "1.0.0",
		Status:      game_room.GameStatusPending,
		LastPingAt:  clockMock.Now(),
	}

	gameRoomInstance := game_room.Instance{
		ID:          "game-1",
		SchedulerID: "game",
	}

	operation := operation.Operation{
		ID:             "some-op-id",
		SchedulerName:  "zooba_blue:1.0.0",
		Status:         operation.StatusPending,
		DefinitionName: "zooba_blue:1.0.0",
	}

	t.Run("when all room creations succeed then it returns nil without error", func(t *testing.T) {
		roomsManager := room_manager.New(clockMock, portAllocatorMock, roomStorageMock, instanceStorageMock, runtimeMock, eventsForwarderService, config)

		schedulerStorage.EXPECT().GetScheduler(context.Background(), operation.SchedulerName).Return(&scheduler, nil)

		portAllocatorMock.EXPECT().Allocate(nil, 2).
			Return([]int32{5000, 6000}, nil).
			Times(10)
		runtimeMock.EXPECT().CreateGameRoomInstance(gomock.Any(), scheduler.Name, game_room.Spec{
			Version:    "1.0.0",
			Containers: []game_room.Container{container1, container2},
		}).
			Return(&gameRoomInstance, nil).
			Times(10)

		gameRoomReady := gameRoom
		gameRoomReady.Status = game_room.GameStatusReady

		roomStorageMock.EXPECT().CreateRoom(gomock.Any(), &gameRoom).Times(10)
		roomStorageMock.EXPECT().GetRoom(gomock.Any(), gameRoom.SchedulerID, gameRoom.ID).Return(&gameRoomReady, nil).Times(10)
		roomStorageMock.EXPECT().WatchRoomStatus(gomock.Any(), &gameRoom).Return(roomStorageStatusWatcher, nil).Times(10)

		roomStorageStatusWatcher.EXPECT().Stop().Times(10)

		err := NewExecutor(roomsManager, schedulerStorage).Execute(context.Background(), &operation, &definition)

		require.NoError(t, err)
	})

	t.Run("when some room creation fail and others succeed then it returns error", func(t *testing.T) {
		roomsManager := mockports.NewMockRoomManager(mockCtrl)

		schedulerStorage.EXPECT().GetScheduler(gomock.Any(), operation.SchedulerName).Return(&scheduler, nil)

		portAllocatorMock.EXPECT().Allocate(nil, 2).
			Return([]int32{5000, 6000}, nil).
			Times(10)

		runtimeMock.EXPECT().CreateGameRoomInstance(gomock.Any(), scheduler.Name, game_room.Spec{
			Version:    "1.0.0",
			Containers: []game_room.Container{container1, container2},
		}).Return(&gameRoomInstance, nil).
			Times(5)
		runtimeMock.EXPECT().CreateGameRoomInstance(gomock.Any(), scheduler.Name, game_room.Spec{
			Version:    "1.0.0",
			Containers: []game_room.Container{container1, container2},
		}).Return(nil, errors.NewErrUnexpected("error create game room instance")).
			Times(5)

		gameRoomReady := gameRoom
		gameRoomReady.Status = game_room.GameStatusReady

		roomStorageMock.EXPECT().CreateRoom(gomock.Any(), &gameRoom).Times(5)
		roomStorageMock.EXPECT().GetRoom(gomock.Any(), gameRoom.SchedulerID, gameRoom.ID).Return(&gameRoomReady, nil).Times(5)
		roomStorageMock.EXPECT().WatchRoomStatus(gomock.Any(), &gameRoom).Return(roomStorageStatusWatcher, nil).Times(5)

		roomStorageStatusWatcher.EXPECT().Stop().Times(5)

		err := NewExecutor(roomsManager, schedulerStorage).Execute(context.Background(), &operation, &definition)

		require.Error(t, err)
	})

	t.Run("when no scheduler is found then it returns the proper error", func(t *testing.T) {
		roomsManager := room_manager.New(clockMock, portAllocatorMock, roomStorageMock, instanceStorageMock, runtimeMock, eventsForwarderService, config)

		schedulerStorage.EXPECT().GetScheduler(context.Background(), operation.SchedulerName).Return(nil, errors.NewErrNotFound("scheduler not found"))

		err := NewExecutor(roomsManager, schedulerStorage).Execute(context.Background(), &operation, &definition)
		require.Error(t, err)
	})
}

func TestAddRoomsExecutor_OnError(t *testing.T) {

	mockCtrl := gomock.NewController(t)

	clockMock := clock_mock.NewFakeClock(time.Now())
	portAllocatorMock := port_allocator_mock.NewMockPortAllocator(mockCtrl)
	roomStorageMock := mockports.NewMockRoomStorage(mockCtrl)
	instanceStorageMock := instance_storage_mock.NewMockGameRoomInstanceStorage(mockCtrl)
	runtimeMock := runtime_mock.NewMockRuntime(mockCtrl)
	eventsForwarderService := mockeventsservice.NewMockEventsService(mockCtrl)
	schedulerStorage := schedulerStorageMock.NewMockSchedulerStorage(mockCtrl)
	config := room_manager.RoomManagerConfig{RoomInitializationTimeout: time.Millisecond * 1000}
	roomStorageStatusWatcher := mockports.NewMockRoomStorageStatusWatcher(mockCtrl)

	definition := AddRoomsDefinition{Amount: 10}

	container1 := game_room.Container{
		Name: "container1",
		Ports: []game_room.ContainerPort{
			{Protocol: "tcp"},
		},
	}

	container2 := game_room.Container{
		Name: "container2",
		Ports: []game_room.ContainerPort{
			{Protocol: "udp"},
		},
	}

	scheduler := entities.Scheduler{
		Name: "zooba_blue:1.0.0",
		Spec: game_room.Spec{
			Version:    "1.0.0",
			Containers: []game_room.Container{container1, container2},
		},
		PortRange: nil,
	}

	gameRoom := game_room.GameRoom{
		ID:          "game-1",
		SchedulerID: "zooba_blue:1.0.0",
		Version:     "1.0.0",
		Status:      game_room.GameStatusPending,
		LastPingAt:  clockMock.Now(),
	}

	gameRoomInstance := game_room.Instance{
		ID:          "game-1",
		SchedulerID: "game",
	}

	operation := operation.Operation{
		ID:             "some-op-id",
		SchedulerName:  "zooba_blue:1.0.0",
		Status:         operation.StatusPending,
		DefinitionName: "zooba_blue:1.0.0",
	}

	t.Run("when no error occurs it deletes previously created rooms and return without error", func(t *testing.T) {
		roomsManager := room_manager.New(clockMock, portAllocatorMock, roomStorageMock, instanceStorageMock, runtimeMock, eventsForwarderService, config)

		schedulerStorage.EXPECT().GetScheduler(gomock.Any(), operation.SchedulerName).Return(&scheduler, nil)

		portAllocatorMock.EXPECT().Allocate(nil, 2).
			Return([]int32{5000, 6000}, nil).
			Times(10)

		runtimeMock.EXPECT().CreateGameRoomInstance(gomock.Any(), scheduler.Name, game_room.Spec{
			Version:    "1.0.0",
			Containers: []game_room.Container{container1, container2},
		}).
			Return(&gameRoomInstance, nil).
			Times(5)
		runtimeMock.EXPECT().CreateGameRoomInstance(gomock.Any(), scheduler.Name, game_room.Spec{
			Version:    "1.0.0",
			Containers: []game_room.Container{container1, container2},
		}).
			Return(nil, errors.NewErrUnexpected("error create game room instance")).
			Times(5)

		gameRoomReady := gameRoom
		gameRoomReady.Status = game_room.GameStatusReady
		gameRoomTerminating := gameRoom
		gameRoomTerminating.Status = game_room.GameStatusTerminating

		roomStorageMock.EXPECT().CreateRoom(gomock.Any(), &gameRoom).Times(5)
		roomStorageMock.EXPECT().GetRoom(gomock.Any(), gameRoom.SchedulerID, gameRoom.ID).Return(&gameRoomReady, nil).Times(5)
		roomStorageMock.EXPECT().WatchRoomStatus(gomock.Any(), &gameRoom).Return(roomStorageStatusWatcher, nil).Times(5)

		roomStorageStatusWatcher.EXPECT().Stop().Times(5)

		executor := NewExecutor(roomsManager, schedulerStorage)

		err := executor.Execute(context.Background(), &operation, &definition)

		require.Error(t, err)

		instanceStorageMock.EXPECT().GetInstance(gomock.Any(), gameRoom.SchedulerID, gameRoom.ID).Return(&gameRoomInstance, nil).Times(5)
		runtimeMock.EXPECT().DeleteGameRoomInstance(gomock.Any(), &gameRoomInstance).Times(5)
		roomStorageMock.EXPECT().GetRoom(gomock.Any(), gameRoom.SchedulerID, gameRoom.ID).Return(&gameRoomTerminating, nil).Times(5)
		roomStorageMock.EXPECT().WatchRoomStatus(gomock.Any(), &gameRoom).Return(roomStorageStatusWatcher, nil).Times(5)
		roomStorageStatusWatcher.EXPECT().Stop().Times(5)

		err = executor.OnError(context.Background(), &operation, &definition, nil)
		require.NoError(t, err)
	})

	t.Run("when some error occurs while deleting rooms it returns error", func(t *testing.T) {
		roomsManager := room_manager.New(clockMock, portAllocatorMock, roomStorageMock, instanceStorageMock, runtimeMock, eventsForwarderService, config)

		schedulerStorage.EXPECT().GetScheduler(gomock.Any(), operation.SchedulerName).Return(&scheduler, nil)

		portAllocatorMock.EXPECT().Allocate(nil, 2).
			Return([]int32{5000, 6000}, nil).
			Times(10)

		runtimeMock.EXPECT().CreateGameRoomInstance(gomock.Any(), scheduler.Name, game_room.Spec{
			Version:    "1.0.0",
			Containers: []game_room.Container{container1, container2},
		}).
			Return(&gameRoomInstance, nil).
			Times(5)
		runtimeMock.EXPECT().CreateGameRoomInstance(gomock.Any(), scheduler.Name, game_room.Spec{
			Version:    "1.0.0",
			Containers: []game_room.Container{container1, container2},
		}).
			Return(nil, errors.NewErrUnexpected("error create game room instance")).
			Times(5)

		gameRoomReady := gameRoom
		gameRoomReady.Status = game_room.GameStatusReady
		gameRoomTerminating := gameRoom
		gameRoomTerminating.Status = game_room.GameStatusTerminating

		roomStorageMock.EXPECT().CreateRoom(gomock.Any(), &gameRoom).Times(5)
		roomStorageMock.EXPECT().GetRoom(gomock.Any(), gameRoom.SchedulerID, gameRoom.ID).Return(&gameRoomReady, nil).Times(5)
		roomStorageMock.EXPECT().WatchRoomStatus(gomock.Any(), &gameRoom).Return(roomStorageStatusWatcher, nil).Times(5)

		roomStorageStatusWatcher.EXPECT().Stop().Times(5)

		executor := NewExecutor(roomsManager, schedulerStorage)

		err := executor.Execute(context.Background(), &operation, &definition)

		require.Error(t, err)

		instanceStorageMock.EXPECT().GetInstance(gomock.Any(), gameRoom.SchedulerID, gameRoom.ID).Return(&gameRoomInstance, nil).Times(1)
		runtimeMock.EXPECT().DeleteGameRoomInstance(gomock.Any(), &gameRoomInstance).Times(1).Return(errors.NewErrUnexpected("error deleting game room instance"))

		err = executor.OnError(context.Background(), &operation, &definition, nil)
		require.Error(t, err)
	})

}
