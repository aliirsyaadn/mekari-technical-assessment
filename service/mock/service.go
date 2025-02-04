// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=mock/service.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	model "github.com/aliirsyaadn/mekari-technical-assessment/model"
	gomock "go.uber.org/mock/gomock"
)

// MockEmployeeService is a mock of EmployeeService interface.
type MockEmployeeService struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeServiceMockRecorder
}

// MockEmployeeServiceMockRecorder is the mock recorder for MockEmployeeService.
type MockEmployeeServiceMockRecorder struct {
	mock *MockEmployeeService
}

// NewMockEmployeeService creates a new mock instance.
func NewMockEmployeeService(ctrl *gomock.Controller) *MockEmployeeService {
	mock := &MockEmployeeService{ctrl: ctrl}
	mock.recorder = &MockEmployeeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeService) EXPECT() *MockEmployeeServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockEmployeeService) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEmployeeServiceMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEmployeeService)(nil).Delete), id)
}

// FindAll mocks base method.
func (m *MockEmployeeService) FindAll() ([]model.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockEmployeeServiceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockEmployeeService)(nil).FindAll))
}

// FindByID mocks base method.
func (m *MockEmployeeService) FindByID(id int) (model.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(model.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockEmployeeServiceMockRecorder) FindByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockEmployeeService)(nil).FindByID), id)
}

// Save mocks base method.
func (m *MockEmployeeService) Save(employee *model.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", employee)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockEmployeeServiceMockRecorder) Save(employee any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockEmployeeService)(nil).Save), employee)
}

// Update mocks base method.
func (m *MockEmployeeService) Update(employee *model.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", employee)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockEmployeeServiceMockRecorder) Update(employee any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEmployeeService)(nil).Update), employee)
}
