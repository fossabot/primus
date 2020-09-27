// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/raba-jp/primus/pkg/operations/file/handlers (interfaces: MoveHandler)

// Package mock_handlers is a generated GoMock package.
package mock_handlers

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	handlers "github.com/raba-jp/primus/pkg/operations/file/handlers"
	reflect "reflect"
)

// MockMoveHandler is a mock of MoveHandler interface
type MockMoveHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMoveHandlerMockRecorder
}

// MockMoveHandlerMockRecorder is the mock recorder for MockMoveHandler
type MockMoveHandlerMockRecorder struct {
	mock *MockMoveHandler
}

// NewMockMoveHandler creates a new mock instance
func NewMockMoveHandler(ctrl *gomock.Controller) *MockMoveHandler {
	mock := &MockMoveHandler{ctrl: ctrl}
	mock.recorder = &MockMoveHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMoveHandler) EXPECT() *MockMoveHandlerMockRecorder {
	return m.recorder
}

// Move mocks base method
func (m *MockMoveHandler) Move(arg0 context.Context, arg1 bool, arg2 *handlers.MoveParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Move", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Move indicates an expected call of Move
func (mr *MockMoveHandlerMockRecorder) Move(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockMoveHandler)(nil).Move), arg0, arg1, arg2)
}