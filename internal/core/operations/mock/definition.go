// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/core/operations/definition.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	operation "github.com/topfreegames/maestro/internal/core/entities/operation"
)

// MockDefinition is a mock of Definition interface.
type MockDefinition struct {
	ctrl     *gomock.Controller
	recorder *MockDefinitionMockRecorder
}

// MockDefinitionMockRecorder is the mock recorder for MockDefinition.
type MockDefinitionMockRecorder struct {
	mock *MockDefinition
}

// NewMockDefinition creates a new mock instance.
func NewMockDefinition(ctrl *gomock.Controller) *MockDefinition {
	mock := &MockDefinition{ctrl: ctrl}
	mock.recorder = &MockDefinitionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDefinition) EXPECT() *MockDefinitionMockRecorder {
	return m.recorder
}

// Marshal mocks base method.
func (m *MockDefinition) Marshal() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshal")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Marshal indicates an expected call of Marshal.
func (mr *MockDefinitionMockRecorder) Marshal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockDefinition)(nil).Marshal))
}

// Name mocks base method.
func (m *MockDefinition) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockDefinitionMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockDefinition)(nil).Name))
}

// ShouldExecute mocks base method.
func (m *MockDefinition) ShouldExecute(ctx context.Context, currentOperations []*operation.Operation) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldExecute", ctx, currentOperations)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldExecute indicates an expected call of ShouldExecute.
func (mr *MockDefinitionMockRecorder) ShouldExecute(ctx, currentOperations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldExecute", reflect.TypeOf((*MockDefinition)(nil).ShouldExecute), ctx, currentOperations)
}

// Unmarshal mocks base method.
func (m *MockDefinition) Unmarshal(raw []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", raw)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmarshal indicates an expected call of Unmarshal.
func (mr *MockDefinitionMockRecorder) Unmarshal(raw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockDefinition)(nil).Unmarshal), raw)
}
