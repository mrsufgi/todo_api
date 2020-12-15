// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mrsufgi/todo_api/internal/domain (interfaces: TodosService)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/mrsufgi/todo_api/internal/domain"
	reflect "reflect"
)

// MockTodosService is a mock of TodosService interface
type MockTodosService struct {
	ctrl     *gomock.Controller
	recorder *MockTodosServiceMockRecorder
}

// MockTodosServiceMockRecorder is the mock recorder for MockTodosService
type MockTodosServiceMockRecorder struct {
	mock *MockTodosService
}

// NewMockTodosService creates a new mock instance
func NewMockTodosService(ctrl *gomock.Controller) *MockTodosService {
	mock := &MockTodosService{ctrl: ctrl}
	mock.recorder = &MockTodosServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTodosService) EXPECT() *MockTodosServiceMockRecorder {
	return m.recorder
}

// CreateTodo mocks base method
func (m *MockTodosService) CreateTodo(arg0 domain.Todo) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodo indicates an expected call of CreateTodo
func (mr *MockTodosServiceMockRecorder) CreateTodo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockTodosService)(nil).CreateTodo), arg0)
}

// DeleteTodo mocks base method
func (m *MockTodosService) DeleteTodo(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTodo indicates an expected call of DeleteTodo
func (mr *MockTodosServiceMockRecorder) DeleteTodo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodo", reflect.TypeOf((*MockTodosService)(nil).DeleteTodo), arg0)
}

// ReadTodo mocks base method
func (m *MockTodosService) ReadTodo(arg0 int) (*domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadTodo", arg0)
	ret0, _ := ret[0].(*domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadTodo indicates an expected call of ReadTodo
func (mr *MockTodosServiceMockRecorder) ReadTodo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadTodo", reflect.TypeOf((*MockTodosService)(nil).ReadTodo), arg0)
}

// SearchTodos mocks base method
func (m *MockTodosService) SearchTodos() (*[]domain.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchTodos")
	ret0, _ := ret[0].(*[]domain.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchTodos indicates an expected call of SearchTodos
func (mr *MockTodosServiceMockRecorder) SearchTodos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchTodos", reflect.TypeOf((*MockTodosService)(nil).SearchTodos))
}

// UpdateTodo mocks base method
func (m *MockTodosService) UpdateTodo(arg0 int, arg1 domain.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTodo indicates an expected call of UpdateTodo
func (mr *MockTodosServiceMockRecorder) UpdateTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodo", reflect.TypeOf((*MockTodosService)(nil).UpdateTodo), arg0, arg1)
}
