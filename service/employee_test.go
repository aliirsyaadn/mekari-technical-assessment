package service

import (
	"testing"

	"github.com/aliirsyaadn/mekari-technical-assessment/config"
	"github.com/aliirsyaadn/mekari-technical-assessment/dummy"
	mockRepo "github.com/aliirsyaadn/mekari-technical-assessment/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestEmployeeService_NewEmployeeService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeRepo := mockRepo.NewMockEmployeeRepository(ctrl)

	employeeService := NewEmployeeService(config.InitLogger(), employeeRepo)

	assert.NotNil(t, employeeService)
}

func TestEmployeeService_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeRepo := mockRepo.NewMockEmployeeRepository(ctrl)

	employeeService := NewEmployeeService(config.InitLogger(), employeeRepo)

	employeeRepo.EXPECT().FindAll().Return(dummy.DummyEmployees, nil)

	employees, err := employeeService.FindAll()
	assert.Nil(t, err)
	assert.NotNil(t, employees)
	assert.Equal(t, 2, len(employees))
}

func TestEmployeeService_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeRepo := mockRepo.NewMockEmployeeRepository(ctrl)

	employeeService := NewEmployeeService(config.InitLogger(), employeeRepo)

	employeeRepo.EXPECT().FindByID(1).Return(dummy.DummyEmployees[0], nil)

	employee, err := employeeService.FindByID(1)
	assert.Nil(t, err)
	assert.NotNil(t, employee)
	assert.Equal(t, 1, employee.ID)
}

func TestEmployeeService_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeRepo := mockRepo.NewMockEmployeeRepository(ctrl)

	employeeService := NewEmployeeService(config.InitLogger(), employeeRepo)

	employeeRepo.EXPECT().Save(&dummy.DummyEmployees[0]).Return(nil)

	err := employeeService.Save(&dummy.DummyEmployees[0])
	assert.Nil(t, err)
}

func TestEmployeeService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeRepo := mockRepo.NewMockEmployeeRepository(ctrl)

	employeeService := NewEmployeeService(config.InitLogger(), employeeRepo)

	employeeRepo.EXPECT().FindByID(1).Return(dummy.DummyEmployees[0], nil)
	employeeRepo.EXPECT().Update(&dummy.DummyEmployees[0]).Return(nil)

	err := employeeService.Update(&dummy.DummyEmployees[0])
	assert.Nil(t, err)
}

func TestEmployeeService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeRepo := mockRepo.NewMockEmployeeRepository(ctrl)

	employeeService := NewEmployeeService(config.InitLogger(), employeeRepo)

	employeeRepo.EXPECT().FindByID(1).Return(dummy.DummyEmployees[0], nil)
	employeeRepo.EXPECT().Delete(1).Return(nil)

	err := employeeService.Delete(1)
	assert.Nil(t, err)
}
