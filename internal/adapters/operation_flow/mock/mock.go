// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/operation_flow.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOperationFlow is a mock of OperationFlow interface
type MockOperationFlow struct {
	ctrl     *gomock.Controller
	recorder *MockOperationFlowMockRecorder
}

// MockOperationFlowMockRecorder is the mock recorder for MockOperationFlow
type MockOperationFlowMockRecorder struct {
	mock *MockOperationFlow
}

// NewMockOperationFlow creates a new mock instance
func NewMockOperationFlow(ctrl *gomock.Controller) *MockOperationFlow {
	mock := &MockOperationFlow{ctrl: ctrl}
	mock.recorder = &MockOperationFlowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOperationFlow) EXPECT() *MockOperationFlowMockRecorder {
	return m.recorder
}

// InsertOperationID mocks base method
func (m *MockOperationFlow) InsertOperationID(ctx context.Context, schedulerName, operationID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOperationID", ctx, schedulerName, operationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOperationID indicates an expected call of InsertOperationID
func (mr *MockOperationFlowMockRecorder) InsertOperationID(ctx, schedulerName, operationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOperationID", reflect.TypeOf((*MockOperationFlow)(nil).InsertOperationID), ctx, schedulerName, operationID)
}

// NextOperationID mocks base method
func (m *MockOperationFlow) NextOperationID(ctx context.Context, schedulerName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextOperationID", ctx, schedulerName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextOperationID indicates an expected call of NextOperationID
func (mr *MockOperationFlowMockRecorder) NextOperationID(ctx, schedulerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextOperationID", reflect.TypeOf((*MockOperationFlow)(nil).NextOperationID), ctx, schedulerName)
}

// ListSchedulerPendingOperationIDs mocks base method
func (m *MockOperationFlow) ListSchedulerPendingOperationIDs(ctx context.Context, schedulerName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedulerPendingOperationIDs", ctx, schedulerName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedulerPendingOperationIDs indicates an expected call of ListSchedulerPendingOperationIDs
func (mr *MockOperationFlowMockRecorder) ListSchedulerPendingOperationIDs(ctx, schedulerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulerPendingOperationIDs", reflect.TypeOf((*MockOperationFlow)(nil).ListSchedulerPendingOperationIDs), ctx, schedulerName)
}
