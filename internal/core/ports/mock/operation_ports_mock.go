// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/core/ports/operation_ports.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	operation "github.com/topfreegames/maestro/internal/core/entities/operation"
	operations "github.com/topfreegames/maestro/internal/core/operations"
	ports "github.com/topfreegames/maestro/internal/core/ports"
)

// MockOperationManager is a mock of OperationManager interface.
type MockOperationManager struct {
	ctrl     *gomock.Controller
	recorder *MockOperationManagerMockRecorder
}

// MockOperationManagerMockRecorder is the mock recorder for MockOperationManager.
type MockOperationManagerMockRecorder struct {
	mock *MockOperationManager
}

// NewMockOperationManager creates a new mock instance.
func NewMockOperationManager(ctrl *gomock.Controller) *MockOperationManager {
	mock := &MockOperationManager{ctrl: ctrl}
	mock.recorder = &MockOperationManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationManager) EXPECT() *MockOperationManagerMockRecorder {
	return m.recorder
}

// AppendOperationEventToExecutionHistory mocks base method.
func (m *MockOperationManager) AppendOperationEventToExecutionHistory(ctx context.Context, op *operation.Operation, eventMessage string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AppendOperationEventToExecutionHistory", ctx, op, eventMessage)
}

// AppendOperationEventToExecutionHistory indicates an expected call of AppendOperationEventToExecutionHistory.
func (mr *MockOperationManagerMockRecorder) AppendOperationEventToExecutionHistory(ctx, op, eventMessage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendOperationEventToExecutionHistory", reflect.TypeOf((*MockOperationManager)(nil).AppendOperationEventToExecutionHistory), ctx, op, eventMessage)
}

// CreateOperation mocks base method.
func (m *MockOperationManager) CreateOperation(ctx context.Context, schedulerName string, definition operations.Definition) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOperation", ctx, schedulerName, definition)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOperation indicates an expected call of CreateOperation.
func (mr *MockOperationManagerMockRecorder) CreateOperation(ctx, schedulerName, definition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOperation", reflect.TypeOf((*MockOperationManager)(nil).CreateOperation), ctx, schedulerName, definition)
}

// CreatePriorityOperation mocks base method.
func (m *MockOperationManager) CreatePriorityOperation(ctx context.Context, schedulerName string, definition operations.Definition) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePriorityOperation", ctx, schedulerName, definition)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePriorityOperation indicates an expected call of CreatePriorityOperation.
func (mr *MockOperationManagerMockRecorder) CreatePriorityOperation(ctx, schedulerName, definition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePriorityOperation", reflect.TypeOf((*MockOperationManager)(nil).CreatePriorityOperation), ctx, schedulerName, definition)
}

// EnqueueOperationCancellationRequest mocks base method.
func (m *MockOperationManager) EnqueueOperationCancellationRequest(ctx context.Context, schedulerName, operationID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnqueueOperationCancellationRequest", ctx, schedulerName, operationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnqueueOperationCancellationRequest indicates an expected call of EnqueueOperationCancellationRequest.
func (mr *MockOperationManagerMockRecorder) EnqueueOperationCancellationRequest(ctx, schedulerName, operationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqueueOperationCancellationRequest", reflect.TypeOf((*MockOperationManager)(nil).EnqueueOperationCancellationRequest), ctx, schedulerName, operationID)
}

// FinishOperation mocks base method.
func (m *MockOperationManager) FinishOperation(ctx context.Context, op *operation.Operation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishOperation", ctx, op)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinishOperation indicates an expected call of FinishOperation.
func (mr *MockOperationManagerMockRecorder) FinishOperation(ctx, op interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishOperation", reflect.TypeOf((*MockOperationManager)(nil).FinishOperation), ctx, op)
}

// GetOperation mocks base method.
func (m *MockOperationManager) GetOperation(ctx context.Context, schedulerName, operationID string) (*operation.Operation, operations.Definition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperation", ctx, schedulerName, operationID)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(operations.Definition)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOperation indicates an expected call of GetOperation.
func (mr *MockOperationManagerMockRecorder) GetOperation(ctx, schedulerName, operationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockOperationManager)(nil).GetOperation), ctx, schedulerName, operationID)
}

// GrantLease mocks base method.
func (m *MockOperationManager) GrantLease(ctx context.Context, operation *operation.Operation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantLease", ctx, operation)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantLease indicates an expected call of GrantLease.
func (mr *MockOperationManagerMockRecorder) GrantLease(ctx, operation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantLease", reflect.TypeOf((*MockOperationManager)(nil).GrantLease), ctx, operation)
}

// ListSchedulerActiveOperations mocks base method.
func (m *MockOperationManager) ListSchedulerActiveOperations(ctx context.Context, schedulerName string) ([]*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedulerActiveOperations", ctx, schedulerName)
	ret0, _ := ret[0].([]*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedulerActiveOperations indicates an expected call of ListSchedulerActiveOperations.
func (mr *MockOperationManagerMockRecorder) ListSchedulerActiveOperations(ctx, schedulerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulerActiveOperations", reflect.TypeOf((*MockOperationManager)(nil).ListSchedulerActiveOperations), ctx, schedulerName)
}

// ListSchedulerFinishedOperations mocks base method.
func (m *MockOperationManager) ListSchedulerFinishedOperations(ctx context.Context, schedulerName string) ([]*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedulerFinishedOperations", ctx, schedulerName)
	ret0, _ := ret[0].([]*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedulerFinishedOperations indicates an expected call of ListSchedulerFinishedOperations.
func (mr *MockOperationManagerMockRecorder) ListSchedulerFinishedOperations(ctx, schedulerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulerFinishedOperations", reflect.TypeOf((*MockOperationManager)(nil).ListSchedulerFinishedOperations), ctx, schedulerName)
}

// ListSchedulerPendingOperations mocks base method.
func (m *MockOperationManager) ListSchedulerPendingOperations(ctx context.Context, schedulerName string) ([]*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedulerPendingOperations", ctx, schedulerName)
	ret0, _ := ret[0].([]*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedulerPendingOperations indicates an expected call of ListSchedulerPendingOperations.
func (mr *MockOperationManagerMockRecorder) ListSchedulerPendingOperations(ctx, schedulerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulerPendingOperations", reflect.TypeOf((*MockOperationManager)(nil).ListSchedulerPendingOperations), ctx, schedulerName)
}

// PendingOperationsChan mocks base method.
func (m *MockOperationManager) PendingOperationsChan(ctx context.Context, schedulerName string) <-chan string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingOperationsChan", ctx, schedulerName)
	ret0, _ := ret[0].(<-chan string)
	return ret0
}

// PendingOperationsChan indicates an expected call of PendingOperationsChan.
func (mr *MockOperationManagerMockRecorder) PendingOperationsChan(ctx, schedulerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingOperationsChan", reflect.TypeOf((*MockOperationManager)(nil).PendingOperationsChan), ctx, schedulerName)
}

// RevokeLease mocks base method.
func (m *MockOperationManager) RevokeLease(ctx context.Context, operation *operation.Operation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeLease", ctx, operation)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeLease indicates an expected call of RevokeLease.
func (mr *MockOperationManagerMockRecorder) RevokeLease(ctx, operation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeLease", reflect.TypeOf((*MockOperationManager)(nil).RevokeLease), ctx, operation)
}

// StartLeaseRenewGoRoutine mocks base method.
func (m *MockOperationManager) StartLeaseRenewGoRoutine(ctx context.Context, op *operation.Operation) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartLeaseRenewGoRoutine", ctx, op)
}

// StartLeaseRenewGoRoutine indicates an expected call of StartLeaseRenewGoRoutine.
func (mr *MockOperationManagerMockRecorder) StartLeaseRenewGoRoutine(ctx, op interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartLeaseRenewGoRoutine", reflect.TypeOf((*MockOperationManager)(nil).StartLeaseRenewGoRoutine), ctx, op)
}

// StartOperation mocks base method.
func (m *MockOperationManager) StartOperation(ctx context.Context, op *operation.Operation, cancelFunction context.CancelFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartOperation", ctx, op, cancelFunction)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartOperation indicates an expected call of StartOperation.
func (mr *MockOperationManagerMockRecorder) StartOperation(ctx, op, cancelFunction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartOperation", reflect.TypeOf((*MockOperationManager)(nil).StartOperation), ctx, op, cancelFunction)
}

// WatchOperationCancellationRequests mocks base method.
func (m *MockOperationManager) WatchOperationCancellationRequests(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchOperationCancellationRequests", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// WatchOperationCancellationRequests indicates an expected call of WatchOperationCancellationRequests.
func (mr *MockOperationManagerMockRecorder) WatchOperationCancellationRequests(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchOperationCancellationRequests", reflect.TypeOf((*MockOperationManager)(nil).WatchOperationCancellationRequests), ctx)
}

// MockOperationFlow is a mock of OperationFlow interface.
type MockOperationFlow struct {
	ctrl     *gomock.Controller
	recorder *MockOperationFlowMockRecorder
}

// MockOperationFlowMockRecorder is the mock recorder for MockOperationFlow.
type MockOperationFlowMockRecorder struct {
	mock *MockOperationFlow
}

// NewMockOperationFlow creates a new mock instance.
func NewMockOperationFlow(ctrl *gomock.Controller) *MockOperationFlow {
	mock := &MockOperationFlow{ctrl: ctrl}
	mock.recorder = &MockOperationFlowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationFlow) EXPECT() *MockOperationFlowMockRecorder {
	return m.recorder
}

// EnqueueOperationCancellationRequest mocks base method.
func (m *MockOperationFlow) EnqueueOperationCancellationRequest(ctx context.Context, request ports.OperationCancellationRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnqueueOperationCancellationRequest", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnqueueOperationCancellationRequest indicates an expected call of EnqueueOperationCancellationRequest.
func (mr *MockOperationFlowMockRecorder) EnqueueOperationCancellationRequest(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqueueOperationCancellationRequest", reflect.TypeOf((*MockOperationFlow)(nil).EnqueueOperationCancellationRequest), ctx, request)
}

// InsertOperationID mocks base method.
func (m *MockOperationFlow) InsertOperationID(ctx context.Context, schedulerName, operationID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOperationID", ctx, schedulerName, operationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOperationID indicates an expected call of InsertOperationID.
func (mr *MockOperationFlowMockRecorder) InsertOperationID(ctx, schedulerName, operationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOperationID", reflect.TypeOf((*MockOperationFlow)(nil).InsertOperationID), ctx, schedulerName, operationID)
}

// ListSchedulerPendingOperationIDs mocks base method.
func (m *MockOperationFlow) ListSchedulerPendingOperationIDs(ctx context.Context, schedulerName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedulerPendingOperationIDs", ctx, schedulerName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedulerPendingOperationIDs indicates an expected call of ListSchedulerPendingOperationIDs.
func (mr *MockOperationFlowMockRecorder) ListSchedulerPendingOperationIDs(ctx, schedulerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulerPendingOperationIDs", reflect.TypeOf((*MockOperationFlow)(nil).ListSchedulerPendingOperationIDs), ctx, schedulerName)
}

// NextOperationID mocks base method.
func (m *MockOperationFlow) NextOperationID(ctx context.Context, schedulerName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextOperationID", ctx, schedulerName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextOperationID indicates an expected call of NextOperationID.
func (mr *MockOperationFlowMockRecorder) NextOperationID(ctx, schedulerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextOperationID", reflect.TypeOf((*MockOperationFlow)(nil).NextOperationID), ctx, schedulerName)
}

// RemoveNextOperation mocks base method.
func (m *MockOperationFlow) RemoveNextOperation(ctx context.Context, schedulerName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveNextOperation", ctx, schedulerName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNextOperation indicates an expected call of RemoveNextOperation.
func (mr *MockOperationFlowMockRecorder) RemoveNextOperation(ctx, schedulerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNextOperation", reflect.TypeOf((*MockOperationFlow)(nil).RemoveNextOperation), ctx, schedulerName)
}

// WatchOperationCancellationRequests mocks base method.
func (m *MockOperationFlow) WatchOperationCancellationRequests(ctx context.Context) chan ports.OperationCancellationRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchOperationCancellationRequests", ctx)
	ret0, _ := ret[0].(chan ports.OperationCancellationRequest)
	return ret0
}

// WatchOperationCancellationRequests indicates an expected call of WatchOperationCancellationRequests.
func (mr *MockOperationFlowMockRecorder) WatchOperationCancellationRequests(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchOperationCancellationRequests", reflect.TypeOf((*MockOperationFlow)(nil).WatchOperationCancellationRequests), ctx)
}

// MockOperationStorage is a mock of OperationStorage interface.
type MockOperationStorage struct {
	ctrl     *gomock.Controller
	recorder *MockOperationStorageMockRecorder
}

// MockOperationStorageMockRecorder is the mock recorder for MockOperationStorage.
type MockOperationStorageMockRecorder struct {
	mock *MockOperationStorage
}

// NewMockOperationStorage creates a new mock instance.
func NewMockOperationStorage(ctrl *gomock.Controller) *MockOperationStorage {
	mock := &MockOperationStorage{ctrl: ctrl}
	mock.recorder = &MockOperationStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationStorage) EXPECT() *MockOperationStorageMockRecorder {
	return m.recorder
}

// CreateOperation mocks base method.
func (m *MockOperationStorage) CreateOperation(ctx context.Context, operation *operation.Operation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOperation", ctx, operation)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOperation indicates an expected call of CreateOperation.
func (mr *MockOperationStorageMockRecorder) CreateOperation(ctx, operation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOperation", reflect.TypeOf((*MockOperationStorage)(nil).CreateOperation), ctx, operation)
}

// GetOperation mocks base method.
func (m *MockOperationStorage) GetOperation(ctx context.Context, schedulerName, operationID string) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperation", ctx, schedulerName, operationID)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperation indicates an expected call of GetOperation.
func (mr *MockOperationStorageMockRecorder) GetOperation(ctx, schedulerName, operationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockOperationStorage)(nil).GetOperation), ctx, schedulerName, operationID)
}

// ListSchedulerActiveOperations mocks base method.
func (m *MockOperationStorage) ListSchedulerActiveOperations(ctx context.Context, schedulerName string) ([]*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedulerActiveOperations", ctx, schedulerName)
	ret0, _ := ret[0].([]*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedulerActiveOperations indicates an expected call of ListSchedulerActiveOperations.
func (mr *MockOperationStorageMockRecorder) ListSchedulerActiveOperations(ctx, schedulerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulerActiveOperations", reflect.TypeOf((*MockOperationStorage)(nil).ListSchedulerActiveOperations), ctx, schedulerName)
}

// ListSchedulerFinishedOperations mocks base method.
func (m *MockOperationStorage) ListSchedulerFinishedOperations(ctx context.Context, schedulerName string) ([]*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedulerFinishedOperations", ctx, schedulerName)
	ret0, _ := ret[0].([]*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedulerFinishedOperations indicates an expected call of ListSchedulerFinishedOperations.
func (mr *MockOperationStorageMockRecorder) ListSchedulerFinishedOperations(ctx, schedulerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulerFinishedOperations", reflect.TypeOf((*MockOperationStorage)(nil).ListSchedulerFinishedOperations), ctx, schedulerName)
}

// UpdateOperationExecutionHistory mocks base method.
func (m *MockOperationStorage) UpdateOperationExecutionHistory(ctx context.Context, op *operation.Operation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOperationExecutionHistory", ctx, op)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOperationExecutionHistory indicates an expected call of UpdateOperationExecutionHistory.
func (mr *MockOperationStorageMockRecorder) UpdateOperationExecutionHistory(ctx, op interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOperationExecutionHistory", reflect.TypeOf((*MockOperationStorage)(nil).UpdateOperationExecutionHistory), ctx, op)
}

// UpdateOperationStatus mocks base method.
func (m *MockOperationStorage) UpdateOperationStatus(ctx context.Context, schedulerName, operationID string, status operation.Status) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOperationStatus", ctx, schedulerName, operationID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOperationStatus indicates an expected call of UpdateOperationStatus.
func (mr *MockOperationStorageMockRecorder) UpdateOperationStatus(ctx, schedulerName, operationID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOperationStatus", reflect.TypeOf((*MockOperationStorage)(nil).UpdateOperationStatus), ctx, schedulerName, operationID, status)
}

// MockOperationLeaseStorage is a mock of OperationLeaseStorage interface.
type MockOperationLeaseStorage struct {
	ctrl     *gomock.Controller
	recorder *MockOperationLeaseStorageMockRecorder
}

// MockOperationLeaseStorageMockRecorder is the mock recorder for MockOperationLeaseStorage.
type MockOperationLeaseStorageMockRecorder struct {
	mock *MockOperationLeaseStorage
}

// NewMockOperationLeaseStorage creates a new mock instance.
func NewMockOperationLeaseStorage(ctrl *gomock.Controller) *MockOperationLeaseStorage {
	mock := &MockOperationLeaseStorage{ctrl: ctrl}
	mock.recorder = &MockOperationLeaseStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationLeaseStorage) EXPECT() *MockOperationLeaseStorageMockRecorder {
	return m.recorder
}

// FetchLeaseTTL mocks base method.
func (m *MockOperationLeaseStorage) FetchLeaseTTL(ctx context.Context, schedulerName, operationID string) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchLeaseTTL", ctx, schedulerName, operationID)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchLeaseTTL indicates an expected call of FetchLeaseTTL.
func (mr *MockOperationLeaseStorageMockRecorder) FetchLeaseTTL(ctx, schedulerName, operationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchLeaseTTL", reflect.TypeOf((*MockOperationLeaseStorage)(nil).FetchLeaseTTL), ctx, schedulerName, operationID)
}

// FetchOperationsLease mocks base method.
func (m *MockOperationLeaseStorage) FetchOperationsLease(ctx context.Context, schedulerName string, operationIDs ...string) ([]*operation.OperationLease, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, schedulerName}
	for _, a := range operationIDs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchOperationsLease", varargs...)
	ret0, _ := ret[0].([]*operation.OperationLease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchOperationsLease indicates an expected call of FetchOperationsLease.
func (mr *MockOperationLeaseStorageMockRecorder) FetchOperationsLease(ctx, schedulerName interface{}, operationIDs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, schedulerName}, operationIDs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOperationsLease", reflect.TypeOf((*MockOperationLeaseStorage)(nil).FetchOperationsLease), varargs...)
}

// GrantLease mocks base method.
func (m *MockOperationLeaseStorage) GrantLease(ctx context.Context, schedulerName, operationID string, initialTTL time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantLease", ctx, schedulerName, operationID, initialTTL)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantLease indicates an expected call of GrantLease.
func (mr *MockOperationLeaseStorageMockRecorder) GrantLease(ctx, schedulerName, operationID, initialTTL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantLease", reflect.TypeOf((*MockOperationLeaseStorage)(nil).GrantLease), ctx, schedulerName, operationID, initialTTL)
}

// ListExpiredLeases mocks base method.
func (m *MockOperationLeaseStorage) ListExpiredLeases(ctx context.Context, schedulerName string, maxLease time.Time) ([]operation.OperationLease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListExpiredLeases", ctx, schedulerName, maxLease)
	ret0, _ := ret[0].([]operation.OperationLease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExpiredLeases indicates an expected call of ListExpiredLeases.
func (mr *MockOperationLeaseStorageMockRecorder) ListExpiredLeases(ctx, schedulerName, maxLease interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExpiredLeases", reflect.TypeOf((*MockOperationLeaseStorage)(nil).ListExpiredLeases), ctx, schedulerName, maxLease)
}

// RenewLease mocks base method.
func (m *MockOperationLeaseStorage) RenewLease(ctx context.Context, schedulerName, operationID string, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenewLease", ctx, schedulerName, operationID, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenewLease indicates an expected call of RenewLease.
func (mr *MockOperationLeaseStorageMockRecorder) RenewLease(ctx, schedulerName, operationID, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewLease", reflect.TypeOf((*MockOperationLeaseStorage)(nil).RenewLease), ctx, schedulerName, operationID, ttl)
}

// RevokeLease mocks base method.
func (m *MockOperationLeaseStorage) RevokeLease(ctx context.Context, schedulerName, operationID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeLease", ctx, schedulerName, operationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeLease indicates an expected call of RevokeLease.
func (mr *MockOperationLeaseStorageMockRecorder) RevokeLease(ctx, schedulerName, operationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeLease", reflect.TypeOf((*MockOperationLeaseStorage)(nil).RevokeLease), ctx, schedulerName, operationID)
}
