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

//go:build unit
// +build unit

package newschedulerversion_test

import (
	serviceerrors "github.com/topfreegames/maestro/internal/core/services/errors"

	"context"
	"fmt"
	"testing"
	"time"

	"github.com/topfreegames/maestro/internal/core/operations"

	"github.com/topfreegames/maestro/internal/core/operations/add_rooms"
	"github.com/topfreegames/maestro/internal/core/ports/errors"

	"github.com/topfreegames/maestro/internal/validations"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/topfreegames/maestro/internal/core/entities"
	"github.com/topfreegames/maestro/internal/core/entities/game_room"
	"github.com/topfreegames/maestro/internal/core/entities/operation"
	"github.com/topfreegames/maestro/internal/core/operations/newschedulerversion"
	mockports "github.com/topfreegames/maestro/internal/core/ports/mock"
)

func TestCreateNewSchedulerVersionExecutor_Execute(t *testing.T) {
	err := validations.RegisterValidations()
	if err != nil {
		t.Errorf("unexpected error %d'", err)
	}

	t.Run("should succeed - major version update, game room is valid, greatest major version is v1, returns no error -> enqueue switch active version op", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v2")
		newSchedulerExpectedVersion := "v2.0.0"
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		switchOpID := "switch-active-version-op-id"
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)

		schedulerVersions := []*entities.SchedulerVersion{{Version: "v1.0.0"}, {Version: "v1.1.0"}, {Version: "v1.2.0"}}
		gameRoom := &game_room.GameRoom{ID: "id-1"}

		roomManager.EXPECT().CreateRoom(gomock.Any(), gomock.Any(), true).Return(gameRoom, nil, nil)
		roomManager.EXPECT().WaitRoomStatus(gomock.Any(), gameRoom, game_room.GameStatusReady).Return(nil)
		roomManager.EXPECT().DeleteRoom(gomock.Any(), gomock.Any()).Return(nil)

		schedulerManager.
			EXPECT().
			CreateNewSchedulerVersionAndEnqueueSwitchVersion(gomock.Any(), gomock.Any()).
			DoAndReturn(
				func(ctx context.Context, scheduler *entities.Scheduler) (string, error) {
					require.Equal(t, newSchedulerExpectedVersion, scheduler.Spec.Version)
					return switchOpID, nil
				})
		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return(schedulerVersions, nil)
		operationsManager.EXPECT().AppendOperationEventToExecutionHistory(gomock.Any(), op, fmt.Sprintf("enqueued switch active version operation with id: %s", switchOpID))

		result := executor.Execute(context.Background(), op, operationDef)

		require.Nil(t, result)
	})

	t.Run("should succeed - major version update, game room is valid, greatest major version is v3, returns no error -> enqueue switch active version op", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v2")
		newSchedulerExpectedVersion := "v4.0.0"
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		switchOpID := "switch-active-version-op-id"
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)

		schedulerVersions := []*entities.SchedulerVersion{{Version: "v2.0.0"}, {Version: "v3.1.0"}, {Version: "v1.2.0"}}
		gameRoom := &game_room.GameRoom{ID: "id-1"}

		roomManager.EXPECT().CreateRoom(gomock.Any(), gomock.Any(), true).Return(gameRoom, nil, nil)
		roomManager.EXPECT().WaitRoomStatus(gomock.Any(), gameRoom, game_room.GameStatusReady).Return(nil)
		roomManager.EXPECT().DeleteRoom(gomock.Any(), gomock.Any()).Return(nil)

		schedulerManager.
			EXPECT().
			CreateNewSchedulerVersionAndEnqueueSwitchVersion(gomock.Any(), gomock.Any()).
			DoAndReturn(
				func(ctx context.Context, scheduler *entities.Scheduler) (string, error) {
					require.Equal(t, newSchedulerExpectedVersion, scheduler.Spec.Version)
					return switchOpID, nil
				})
		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return(schedulerVersions, nil)
		operationsManager.EXPECT().AppendOperationEventToExecutionHistory(gomock.Any(), op, fmt.Sprintf("enqueued switch active version operation with id: %s", switchOpID))

		result := executor.Execute(context.Background(), op, operationDef)

		require.Nil(t, result)
	})

	t.Run("should succeed - major version update, game room is valid, greatest major version is the current one, returns no error -> enqueue switch active version op", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v2")
		newSchedulerExpectedVersion := "v2.0.0"
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		switchOpID := "switch-active-version-op-id"
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)

		schedulerVersions := []*entities.SchedulerVersion{{Version: "v1.1.0"}, {Version: "v1.2.0"}, {Version: "v1.3.0"}}
		gameRoom := &game_room.GameRoom{ID: "id-1"}

		roomManager.EXPECT().CreateRoom(gomock.Any(), gomock.Any(), true).Return(gameRoom, nil, nil)
		roomManager.EXPECT().WaitRoomStatus(gomock.Any(), gameRoom, game_room.GameStatusReady).Return(nil)
		roomManager.EXPECT().DeleteRoom(gomock.Any(), gomock.Any()).Return(nil)

		schedulerManager.
			EXPECT().
			CreateNewSchedulerVersionAndEnqueueSwitchVersion(gomock.Any(), gomock.Any()).
			DoAndReturn(
				func(ctx context.Context, scheduler *entities.Scheduler) (string, error) {
					require.Equal(t, newSchedulerExpectedVersion, scheduler.Spec.Version)
					return switchOpID, nil
				})
		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return(schedulerVersions, nil)
		operationsManager.EXPECT().AppendOperationEventToExecutionHistory(gomock.Any(), op, fmt.Sprintf("enqueued switch active version operation with id: %s", switchOpID))

		result := executor.Execute(context.Background(), op, operationDef)

		require.Nil(t, result)
	})

	t.Run("should succeed - major version update, game room is valid, fail to delete validation room, returns no error -> enqueue switch active version op", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v2")
		newSchedulerExpectedVersion := "v2.0.0"
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		switchOpID := "switch-active-version-op-id"
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)

		schedulerVersions := []*entities.SchedulerVersion{{Version: "v1.1.0"}, {Version: "v1.2.0"}, {Version: "v1.3.0"}}
		gameRoom := &game_room.GameRoom{ID: "id-1"}

		roomManager.EXPECT().CreateRoom(gomock.Any(), gomock.Any(), true).Return(gameRoom, nil, nil)
		roomManager.EXPECT().WaitRoomStatus(gomock.Any(), gameRoom, game_room.GameStatusReady).Return(nil)
		roomManager.EXPECT().DeleteRoom(gomock.Any(), gomock.Any()).Return(errors.NewErrUnexpected("some_error"))

		schedulerManager.
			EXPECT().
			CreateNewSchedulerVersionAndEnqueueSwitchVersion(gomock.Any(), gomock.Any()).
			DoAndReturn(
				func(ctx context.Context, scheduler *entities.Scheduler) (string, error) {
					require.Equal(t, newSchedulerExpectedVersion, scheduler.Spec.Version)
					return switchOpID, nil
				})
		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return(schedulerVersions, nil)
		operationsManager.EXPECT().AppendOperationEventToExecutionHistory(gomock.Any(), op, fmt.Sprintf("enqueued switch active version operation with id: %s", switchOpID))

		result := executor.Execute(context.Background(), op, operationDef)

		require.Nil(t, result)
	})

	t.Run("should fail - major version update, game room is valid, fail when loading scheduler versions -> returns error, don't create new version/switch to it", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v2")
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)

		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return([]*entities.SchedulerVersion{}, errors.NewErrUnexpected("some_error"))

		result := executor.Execute(context.Background(), op, operationDef)

		require.NotNil(t, result)
		require.Equal(t, result.Kind(), operations.ErrKindUnexpected)
		require.EqualError(t, result.Error(), "failed to calculate new major version: failed to load scheduler versions: some_error")
	})

	t.Run("should fail - major version update, game room is valid, fail parsing scheduler versions -> returns error, don't create new version/switch to it", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v2")
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		schedulerVersions := []*entities.SchedulerVersion{{Version: "v-----"}}
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)

		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return(schedulerVersions, nil)

		result := executor.Execute(context.Background(), op, operationDef)

		require.NotNil(t, result)
		require.Equal(t, result.Kind(), operations.ErrKindUnexpected)
		require.EqualError(t, result.Error(), "failed to calculate new major version: failed to parse scheduler version v-----: Invalid Semantic Version")
	})

	t.Run("should fail - major version update, fail creating test room", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v2")
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)
		schedulerVersions := []*entities.SchedulerVersion{{Version: "v2.0.0"}, {Version: "v3.1.0"}, {Version: "v1.2.0"}}

		newSchedulerWithNewVersion := newScheduler
		newSchedulerWithNewVersion.Spec.Version = "v2.0.0"
		newSchedulerWithNewVersion.RollbackVersion = "v1.0.0"

		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return(schedulerVersions, nil)

		roomManager.EXPECT().CreateRoom(gomock.Any(), gomock.Any(), true).Return(nil, nil, fmt.Errorf("error creating test game room"))

		operationExecutionError := executor.Execute(context.Background(), op, operationDef)

		require.NotNil(t, operationExecutionError)
		require.Equal(t, operations.ErrKindUnexpected, operationExecutionError.Kind())
		require.ErrorContains(t, operationExecutionError.Error(), "error creating new game room for validating new version")
	})

	t.Run("should fail - major version update, game room is invalid, timeout error -> returns error, don't create new version/switch to it", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v2")
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)
		schedulerVersions := []*entities.SchedulerVersion{{Version: "v2.0.0"}, {Version: "v3.1.0"}, {Version: "v1.2.0"}}

		newSchedulerWithNewVersion := newScheduler
		newSchedulerWithNewVersion.Spec.Version = "v2.0.0"
		newSchedulerWithNewVersion.RollbackVersion = "v1.0.0"
		gameRoom := &game_room.GameRoom{ID: "id-1", SchedulerID: "some-scheduler"}

		roomManager.EXPECT().CreateRoom(gomock.Any(), gomock.Any(), true).Return(gameRoom, nil, nil)
		roomManager.EXPECT().WaitRoomStatus(gomock.Any(), gameRoom, game_room.GameStatusReady).Return(serviceerrors.NewErrGameRoomStatusWaitingTimeout("some error"))
		roomManager.EXPECT().DeleteRoom(gomock.Any(), gameRoom).Return(nil)

		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return(schedulerVersions, nil)

		operationExecutionError := executor.Execute(context.Background(), op, operationDef)

		require.NotNil(t, operationExecutionError)
		require.Equal(t, operations.ErrKindInvalidGru, operationExecutionError.Kind())
		require.ErrorContains(t, operationExecutionError.Error(), "error validating game room with ID")
	})

	t.Run("should fail - major version update, game room is invalid, unexpected error-> returns error, don't create new version/switch to it", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v2")
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)
		schedulerVersions := []*entities.SchedulerVersion{{Version: "v2.0.0"}, {Version: "v3.1.0"}, {Version: "v1.2.0"}}

		newSchedulerWithNewVersion := newScheduler
		newSchedulerWithNewVersion.Spec.Version = "v2.0.0"
		newSchedulerWithNewVersion.RollbackVersion = "v1.0.0"
		gameRoom := &game_room.GameRoom{ID: "id-1", SchedulerID: "some-scheduler"}

		roomManager.EXPECT().CreateRoom(gomock.Any(), gomock.Any(), true).Return(gameRoom, nil, nil)
		roomManager.EXPECT().WaitRoomStatus(gomock.Any(), gameRoom, game_room.GameStatusReady).Return(fmt.Errorf("some error"))
		roomManager.EXPECT().DeleteRoom(gomock.Any(), gameRoom).Return(nil)

		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return(schedulerVersions, nil)

		operationExecutionError := executor.Execute(context.Background(), op, operationDef)

		require.NotNil(t, operationExecutionError)
		require.Equal(t, operationExecutionError.Kind(), operations.ErrKindUnexpected)
		require.ErrorContains(t, operationExecutionError.Error(), "unexpected error validating game room with ID")
	})

	t.Run("should succeed - given a minor version update it, when the greatest minor version is v1.0 returns no error and enqueue switch active version op", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v1")
		newSchedulerExpectedVersion := "v1.1.0"
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		switchOpID := "switch-active-version-op-id"
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)

		schedulerVersions := []*entities.SchedulerVersion{{Version: "v2.0.0"}, {Version: "v3.1.0"}, {Version: "v4.2.0"}}

		schedulerManager.
			EXPECT().
			CreateNewSchedulerVersionAndEnqueueSwitchVersion(gomock.Any(), gomock.Any()).
			DoAndReturn(
				func(ctx context.Context, scheduler *entities.Scheduler) (string, error) {
					require.Equal(t, newSchedulerExpectedVersion, scheduler.Spec.Version)
					return switchOpID, nil
				})
		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return(schedulerVersions, nil)
		operationsManager.EXPECT().AppendOperationEventToExecutionHistory(gomock.Any(), op, fmt.Sprintf("enqueued switch active version operation with id: %s", switchOpID))

		result := executor.Execute(context.Background(), op, operationDef)

		require.Nil(t, result)
	})

	t.Run("should succeed - given a minor version update it, when the greatest minor version is v1.5 returns no error and enqueue switch active version op", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v1")
		newSchedulerExpectedVersion := "v1.6.0"
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		switchOpID := "switch-active-version-op-id"
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)

		schedulerVersions := []*entities.SchedulerVersion{{Version: "v2.0.0"}, {Version: "v1.3.0"}, {Version: "v1.5.0"}}

		schedulerManager.
			EXPECT().
			CreateNewSchedulerVersionAndEnqueueSwitchVersion(gomock.Any(), gomock.Any()).
			DoAndReturn(
				func(ctx context.Context, scheduler *entities.Scheduler) (string, error) {
					require.Equal(t, newSchedulerExpectedVersion, scheduler.Spec.Version)
					return switchOpID, nil
				})
		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return(schedulerVersions, nil)
		operationsManager.EXPECT().AppendOperationEventToExecutionHistory(gomock.Any(), op, fmt.Sprintf("enqueued switch active version operation with id: %s", switchOpID))

		result := executor.Execute(context.Background(), op, operationDef)

		require.Nil(t, result)
	})

	t.Run("should succeed - given a minor version update it, when the greatest minor version is the current one returns no error and enqueue switch active version op", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v1")
		newSchedulerExpectedVersion := "v1.1.0"
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		switchOpID := "switch-active-version-op-id"
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)

		schedulerVersions := []*entities.SchedulerVersion{{Version: "v2.0.0"}, {Version: "v2.1.0"}, {Version: "v3.5.0"}}

		schedulerManager.
			EXPECT().
			CreateNewSchedulerVersionAndEnqueueSwitchVersion(gomock.Any(), gomock.Any()).
			DoAndReturn(
				func(ctx context.Context, scheduler *entities.Scheduler) (string, error) {
					require.Equal(t, newSchedulerExpectedVersion, scheduler.Spec.Version)
					return switchOpID, nil
				})
		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return(schedulerVersions, nil)
		operationsManager.EXPECT().AppendOperationEventToExecutionHistory(gomock.Any(), op, fmt.Sprintf("enqueued switch active version operation with id: %s", switchOpID))

		result := executor.Execute(context.Background(), op, operationDef)

		require.Nil(t, result)
	})

	t.Run("should fail - minor version update, fail when loading scheduler versions -> returns error, don't create new version/switch to it", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v1")
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)

		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return([]*entities.SchedulerVersion{}, errors.NewErrUnexpected("some_error"))

		result := executor.Execute(context.Background(), op, operationDef)

		require.NotNil(t, result)
		require.Equal(t, result.Kind(), operations.ErrKindUnexpected)
		require.EqualError(t, result.Error(), "failed to calculate new minor version: failed to load scheduler versions: some_error")
	})

	t.Run("should fail - minor version update, fail parsing scheduler versions -> returns error, don't create new version/switch to it", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("image-v1")
		newScheduler := *newValidSchedulerWithImageVersion("image-v1")
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		schedulerVersions := []*entities.SchedulerVersion{{Version: "v-----"}}
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)

		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return(schedulerVersions, nil)

		result := executor.Execute(context.Background(), op, operationDef)

		require.NotNil(t, result)
		require.Equal(t, result.Kind(), operations.ErrKindUnexpected)
		require.EqualError(t, result.Error(), "failed to calculate new minor version: failed to parse scheduler version v-----: Invalid Semantic Version")
	})

	t.Run("should fail - valid scheduler, error occurs (creating new version in db or enqueueing switch op) -> returns error, don't create new version/switch to it", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		currentActiveScheduler := newValidSchedulerWithImageVersion("v1.0")
		newScheduler := *newValidSchedulerWithImageVersion("v1.0")
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)
		newSchedulerWithNewVersion := newScheduler
		newSchedulerWithNewVersion.Spec.Version = "v1.1.0"
		newSchedulerWithNewVersion.RollbackVersion = "v1.0.0"

		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)
		schedulerManager.EXPECT().GetSchedulerVersions(gomock.Any(), newScheduler.Name).Return([]*entities.SchedulerVersion{}, nil)
		schedulerManager.
			EXPECT().
			CreateNewSchedulerVersionAndEnqueueSwitchVersion(gomock.Any(), gomock.Any()).
			Return("", errors.NewErrUnexpected("some_error"))

		result := executor.Execute(context.Background(), op, operationDef)

		require.NotNil(t, result)
		require.Equal(t, result.Kind(), operations.ErrKindUnexpected)
		require.EqualError(t, result.Error(), "error creating new scheduler version in db: some_error")
	})

	t.Run("should fail - valid scheduler, some error occurs (retrieving current active scheduler), returns error, don't create new version", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		newScheduler := *newValidSchedulerWithImageVersion("v1.0")
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)
		newSchedulerWithNewVersion := newScheduler
		newSchedulerWithNewVersion.Spec.Version = "v1.1.0"
		newSchedulerWithNewVersion.RollbackVersion = "v1.0.0"

		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(nil, errors.NewErrUnexpected("some_error"))

		result := executor.Execute(context.Background(), op, operationDef)

		require.NotNil(t, result)
		require.Equal(t, result.Kind(), operations.ErrKindUnexpected)
		require.EqualError(t, result.Error(), "error getting active scheduler: some_error")
	})

	t.Run("should fail - valid scheduler when provided operation definition != CreateNewSchedulerVersionDefinition, returns error, don't create new version", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		newScheduler := *newValidSchedulerWithImageVersion("v1.0")
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &add_rooms.AddRoomsDefinition{}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)
		newSchedulerWithNewVersion := newScheduler
		newSchedulerWithNewVersion.Spec.Version = "v1.1.0"
		newSchedulerWithNewVersion.RollbackVersion = "v1.0.0"

		result := executor.Execute(context.Background(), op, operationDef)

		require.NotNil(t, result)
		require.Equal(t, result.Kind(), operations.ErrKindUnexpected)
		require.EqualError(t, result.Error(), "invalid operation definition for create_new_scheduler_version operation")
	})

	t.Run("given a invalid scheduler when the version parse fails it returns error and don't create new version", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		newScheduler := newInValidScheduler()
		currentActiveScheduler := newInValidScheduler()

		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)

		// mocks for SchedulerManager GetActiveScheduler method
		schedulerManager.EXPECT().GetActiveScheduler(gomock.Any(), newScheduler.Name).Return(currentActiveScheduler, nil)

		result := executor.Execute(context.Background(), op, operationDef)

		require.NotNil(t, result)
		require.Equal(t, result.Kind(), operations.ErrKindUnexpected)
		require.EqualError(t, result.Error(), "failed to parse scheduler current version: Invalid Semantic Version")
	})

}

func TestCreateNewSchedulerVersionExecutor_Rollback(t *testing.T) {
	err := validations.RegisterValidations()
	if err != nil {
		t.Errorf("unexpected error %d'", err)
	}

	t.Run("when some game room were created during execution, it deletes the room and return no error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		newScheduler := *newValidSchedulerWithImageVersion("v1.2")
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)
		executor.AddValidationRoomID(newScheduler.Name, &game_room.GameRoom{ID: "room1"})
		roomManager.EXPECT().DeleteRoom(gomock.Any(), gomock.Any()).Return(nil)
		result := executor.Rollback(context.Background(), op, operationDef, nil)

		require.Nil(t, result)
	})

	t.Run("when some game room were created during execution, it returns error if some error occur in deleting the game room", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		newScheduler := *newValidSchedulerWithImageVersion("v1.2")
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)
		executor.AddValidationRoomID(newScheduler.Name, &game_room.GameRoom{ID: "room1"})
		roomManager.EXPECT().DeleteRoom(gomock.Any(), gomock.Any()).Return(errors.NewErrUnexpected("some error"))
		result := executor.Rollback(context.Background(), op, operationDef, nil)

		require.EqualError(t, result, "error in Rollback function execution: some error")
	})

	t.Run("when no game room were created during execution, it does nothing", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)

		newScheduler := *newValidSchedulerWithImageVersion("v1.2")
		op := &operation.Operation{
			ID:             "123",
			Status:         operation.StatusInProgress,
			DefinitionName: newschedulerversion.OperationName,
			SchedulerName:  newScheduler.Name,
		}
		operationDef := &newschedulerversion.CreateNewSchedulerVersionDefinition{NewScheduler: &newScheduler}
		roomManager := mockports.NewMockRoomManager(mockCtrl)
		schedulerManager := mockports.NewMockSchedulerManager(mockCtrl)
		operationsManager := mockports.NewMockOperationManager(mockCtrl)
		config := newschedulerversion.Config{
			RoomInitializationTimeout: time.Duration(120000),
		}

		executor := newschedulerversion.NewExecutor(roomManager, schedulerManager, operationsManager, config)
		result := executor.Rollback(context.Background(), op, operationDef, nil)

		require.Nil(t, result)
	})

}

func newValidSchedulerWithImageVersion(imageVersion string) *entities.Scheduler {
	return &entities.Scheduler{
		Name:            "scheduler",
		Game:            "game",
		State:           entities.StateCreating,
		MaxSurge:        "5",
		RollbackVersion: "",
		Spec: game_room.Spec{
			Version:                "v1.0.0",
			TerminationGracePeriod: 60,
			Toleration:             "toleration",
			Affinity:               "affinity",
			Containers: []game_room.Container{
				{
					Name:            "default",
					Image:           fmt.Sprintf("some-image:%s", imageVersion),
					ImagePullPolicy: "IfNotPresent",
					Command:         []string{"hello"},
					Ports: []game_room.ContainerPort{
						{Name: "tcp", Protocol: "tcp", Port: 80},
					},
					Requests: game_room.ContainerResources{
						CPU:    "10m",
						Memory: "100Mi",
					},
					Limits: game_room.ContainerResources{
						CPU:    "10m",
						Memory: "100Mi",
					},
				},
			},
		},
		PortRange: &entities.PortRange{
			Start: 40000,
			End:   60000,
		},
	}
}

func newInValidScheduler() *entities.Scheduler {
	scheduler := newValidSchedulerWithImageVersion("v1.0.0")
	scheduler.Spec.Version = "R1.0.0"
	return scheduler
}
