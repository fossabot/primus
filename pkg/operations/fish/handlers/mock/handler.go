// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/raba-jp/primus/pkg/operations/fish/handlers (interfaces: SetVariableHandler,SetPathHandler)

// Package mock_handlers is a generated GoMock package.
package mock_handlers

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	handlers "github.com/raba-jp/primus/pkg/operations/fish/handlers"
	reflect "reflect"
)

// MockSetVariableHandler is a mock of SetVariableHandler interface
type MockSetVariableHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSetVariableHandlerMockRecorder
}

// MockSetVariableHandlerMockRecorder is the mock recorder for MockSetVariableHandler
type MockSetVariableHandlerMockRecorder struct {
	mock *MockSetVariableHandler
}

// NewMockSetVariableHandler creates a new mock instance
func NewMockSetVariableHandler(ctrl *gomock.Controller) *MockSetVariableHandler {
	mock := &MockSetVariableHandler{ctrl: ctrl}
	mock.recorder = &MockSetVariableHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSetVariableHandler) EXPECT() *MockSetVariableHandlerMockRecorder {
	return m.recorder
}

// SetVariable mocks base method
func (m *MockSetVariableHandler) SetVariable(arg0 context.Context, arg1 bool, arg2 *handlers.SetVariableParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVariable", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVariable indicates an expected call of SetVariable
func (mr *MockSetVariableHandlerMockRecorder) SetVariable(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVariable", reflect.TypeOf((*MockSetVariableHandler)(nil).SetVariable), arg0, arg1, arg2)
}

// MockSetPathHandler is a mock of SetPathHandler interface
type MockSetPathHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSetPathHandlerMockRecorder
}

// MockSetPathHandlerMockRecorder is the mock recorder for MockSetPathHandler
type MockSetPathHandlerMockRecorder struct {
	mock *MockSetPathHandler
}

// NewMockSetPathHandler creates a new mock instance
func NewMockSetPathHandler(ctrl *gomock.Controller) *MockSetPathHandler {
	mock := &MockSetPathHandler{ctrl: ctrl}
	mock.recorder = &MockSetPathHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSetPathHandler) EXPECT() *MockSetPathHandlerMockRecorder {
	return m.recorder
}

// SetPath mocks base method
func (m *MockSetPathHandler) SetPath(arg0 context.Context, arg1 bool, arg2 *handlers.SetPathParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPath", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPath indicates an expected call of SetPath
func (mr *MockSetPathHandlerMockRecorder) SetPath(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPath", reflect.TypeOf((*MockSetPathHandler)(nil).SetPath), arg0, arg1, arg2)
}
