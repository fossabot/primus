// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/raba-jp/primus/pkg/operations/file/handlers (interfaces: SymlinkHandler)

// Package mock_handlers is a generated GoMock package.
package mock_handlers

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	handlers "github.com/raba-jp/primus/pkg/operations/file/handlers"
	reflect "reflect"
)

// MockSymlinkHandler is a mock of SymlinkHandler interface
type MockSymlinkHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSymlinkHandlerMockRecorder
}

// MockSymlinkHandlerMockRecorder is the mock recorder for MockSymlinkHandler
type MockSymlinkHandlerMockRecorder struct {
	mock *MockSymlinkHandler
}

// NewMockSymlinkHandler creates a new mock instance
func NewMockSymlinkHandler(ctrl *gomock.Controller) *MockSymlinkHandler {
	mock := &MockSymlinkHandler{ctrl: ctrl}
	mock.recorder = &MockSymlinkHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSymlinkHandler) EXPECT() *MockSymlinkHandlerMockRecorder {
	return m.recorder
}

// Symlink mocks base method
func (m *MockSymlinkHandler) Symlink(arg0 context.Context, arg1 bool, arg2 *handlers.SymlinkParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Symlink", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Symlink indicates an expected call of Symlink
func (mr *MockSymlinkHandlerMockRecorder) Symlink(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Symlink", reflect.TypeOf((*MockSymlinkHandler)(nil).Symlink), arg0, arg1, arg2)
}
