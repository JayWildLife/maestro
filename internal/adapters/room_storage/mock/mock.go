// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/room_storage.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	game_room "github.com/topfreegames/maestro/internal/core/entities/game_room"
	ports "github.com/topfreegames/maestro/internal/core/ports"
)

// MockRoomStorage is a mock of RoomStorage interface.
type MockRoomStorage struct {
	ctrl     *gomock.Controller
	recorder *MockRoomStorageMockRecorder
}

// MockRoomStorageMockRecorder is the mock recorder for MockRoomStorage.
type MockRoomStorageMockRecorder struct {
	mock *MockRoomStorage
}

// NewMockRoomStorage creates a new mock instance.
func NewMockRoomStorage(ctrl *gomock.Controller) *MockRoomStorage {
	mock := &MockRoomStorage{ctrl: ctrl}
	mock.recorder = &MockRoomStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomStorage) EXPECT() *MockRoomStorageMockRecorder {
	return m.recorder
}

// CreateRoom mocks base method.
func (m *MockRoomStorage) CreateRoom(ctx context.Context, room *game_room.GameRoom) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", ctx, room)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockRoomStorageMockRecorder) CreateRoom(ctx, room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockRoomStorage)(nil).CreateRoom), ctx, room)
}

// DeleteRoom mocks base method.
func (m *MockRoomStorage) DeleteRoom(ctx context.Context, scheduler, roomID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoom", ctx, scheduler, roomID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoom indicates an expected call of DeleteRoom.
func (mr *MockRoomStorageMockRecorder) DeleteRoom(ctx, scheduler, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoom", reflect.TypeOf((*MockRoomStorage)(nil).DeleteRoom), ctx, scheduler, roomID)
}

// GetAllRoomIDs mocks base method.
func (m *MockRoomStorage) GetAllRoomIDs(ctx context.Context, scheduler string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRoomIDs", ctx, scheduler)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRoomIDs indicates an expected call of GetAllRoomIDs.
func (mr *MockRoomStorageMockRecorder) GetAllRoomIDs(ctx, scheduler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRoomIDs", reflect.TypeOf((*MockRoomStorage)(nil).GetAllRoomIDs), ctx, scheduler)
}

// GetRoom mocks base method.
func (m *MockRoomStorage) GetRoom(ctx context.Context, scheduler, roomID string) (*game_room.GameRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoom", ctx, scheduler, roomID)
	ret0, _ := ret[0].(*game_room.GameRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoom indicates an expected call of GetRoom.
func (mr *MockRoomStorageMockRecorder) GetRoom(ctx, scheduler, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoom", reflect.TypeOf((*MockRoomStorage)(nil).GetRoom), ctx, scheduler, roomID)
}

// GetRoomCount mocks base method.
func (m *MockRoomStorage) GetRoomCount(ctx context.Context, scheduler string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomCount", ctx, scheduler)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomCount indicates an expected call of GetRoomCount.
func (mr *MockRoomStorageMockRecorder) GetRoomCount(ctx, scheduler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomCount", reflect.TypeOf((*MockRoomStorage)(nil).GetRoomCount), ctx, scheduler)
}

// GetRoomCountByStatus mocks base method.
func (m *MockRoomStorage) GetRoomCountByStatus(ctx context.Context, scheduler string, status game_room.GameRoomStatus) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomCountByStatus", ctx, scheduler, status)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomCountByStatus indicates an expected call of GetRoomCountByStatus.
func (mr *MockRoomStorageMockRecorder) GetRoomCountByStatus(ctx, scheduler, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomCountByStatus", reflect.TypeOf((*MockRoomStorage)(nil).GetRoomCountByStatus), ctx, scheduler, status)
}

// GetRoomIDsByLastPing mocks base method.
func (m *MockRoomStorage) GetRoomIDsByLastPing(ctx context.Context, scheduler string, threshold time.Time) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomIDsByLastPing", ctx, scheduler, threshold)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomIDsByLastPing indicates an expected call of GetRoomIDsByLastPing.
func (mr *MockRoomStorageMockRecorder) GetRoomIDsByLastPing(ctx, scheduler, threshold interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomIDsByLastPing", reflect.TypeOf((*MockRoomStorage)(nil).GetRoomIDsByLastPing), ctx, scheduler, threshold)
}

// SetRoomStatus mocks base method.
func (m *MockRoomStorage) SetRoomStatus(ctx context.Context, scheduler, roomID string, status game_room.GameRoomStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRoomStatus", ctx, scheduler, roomID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRoomStatus indicates an expected call of SetRoomStatus.
func (mr *MockRoomStorageMockRecorder) SetRoomStatus(ctx, scheduler, roomID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRoomStatus", reflect.TypeOf((*MockRoomStorage)(nil).SetRoomStatus), ctx, scheduler, roomID, status)
}

// UpdateRoom mocks base method.
func (m *MockRoomStorage) UpdateRoom(ctx context.Context, room *game_room.GameRoom) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoom", ctx, room)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRoom indicates an expected call of UpdateRoom.
func (mr *MockRoomStorageMockRecorder) UpdateRoom(ctx, room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoom", reflect.TypeOf((*MockRoomStorage)(nil).UpdateRoom), ctx, room)
}

// WatchRoomStatus mocks base method.
func (m *MockRoomStorage) WatchRoomStatus(ctx context.Context, room *game_room.GameRoom) (ports.RoomStorageStatusWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchRoomStatus", ctx, room)
	ret0, _ := ret[0].(ports.RoomStorageStatusWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchRoomStatus indicates an expected call of WatchRoomStatus.
func (mr *MockRoomStorageMockRecorder) WatchRoomStatus(ctx, room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRoomStatus", reflect.TypeOf((*MockRoomStorage)(nil).WatchRoomStatus), ctx, room)
}

// MockRoomStorageStatusWatcher is a mock of RoomStorageStatusWatcher interface.
type MockRoomStorageStatusWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockRoomStorageStatusWatcherMockRecorder
}

// MockRoomStorageStatusWatcherMockRecorder is the mock recorder for MockRoomStorageStatusWatcher.
type MockRoomStorageStatusWatcherMockRecorder struct {
	mock *MockRoomStorageStatusWatcher
}

// NewMockRoomStorageStatusWatcher creates a new mock instance.
func NewMockRoomStorageStatusWatcher(ctrl *gomock.Controller) *MockRoomStorageStatusWatcher {
	mock := &MockRoomStorageStatusWatcher{ctrl: ctrl}
	mock.recorder = &MockRoomStorageStatusWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomStorageStatusWatcher) EXPECT() *MockRoomStorageStatusWatcherMockRecorder {
	return m.recorder
}

// ResultChan mocks base method.
func (m *MockRoomStorageStatusWatcher) ResultChan() chan game_room.StatusEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResultChan")
	ret0, _ := ret[0].(chan game_room.StatusEvent)
	return ret0
}

// ResultChan indicates an expected call of ResultChan.
func (mr *MockRoomStorageStatusWatcherMockRecorder) ResultChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResultChan", reflect.TypeOf((*MockRoomStorageStatusWatcher)(nil).ResultChan))
}

// Stop mocks base method.
func (m *MockRoomStorageStatusWatcher) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockRoomStorageStatusWatcherMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRoomStorageStatusWatcher)(nil).Stop))
}
