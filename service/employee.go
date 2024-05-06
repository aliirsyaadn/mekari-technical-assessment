package service

import (
	"errors"

	"github.com/aliirsyaadn/mekari-technical-assessment/model"
	"github.com/aliirsyaadn/mekari-technical-assessment/repository"
	"go.uber.org/zap"
)

type employeeService struct {
	log                *zap.SugaredLogger
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(
	log *zap.SugaredLogger,
	employeeRepository repository.EmployeeRepository,
) EmployeeService {
	return &employeeService{
		log:                log,
		employeeRepository: employeeRepository,
	}
}

func (s *employeeService) FindAll() ([]model.Employee, error) {
	return s.employeeRepository.FindAll()
}

func (s *employeeService) FindByID(id int) (model.Employee, error) {
	return s.employeeRepository.FindByID(id)
}

func (s *employeeService) Save(employee *model.Employee) error {
	return s.employeeRepository.Save(employee)
}

func (s *employeeService) Update(employee *model.Employee) error {
	prevData, err := s.employeeRepository.FindByID(employee.ID)
	if err != nil {
		s.log.Error("Failed to get employee with id", employee.ID)
		return err
	}

	if prevData.ID == 0 {
		s.log.Error("Failed employee not found")
		return errors.New("employee not found")
	}

	err = s.employeeRepository.Update(employee)
	if err != nil {
		s.log.Error("Failed to update employee", employee.ID)
		return err
	}

	employee.CreatedAt = prevData.CreatedAt

	return nil
}

func (s *employeeService) Delete(id int) error {
	prevData, err := s.employeeRepository.FindByID(id)
	if err != nil {
		s.log.Error("Failed to get employee with id", id)
		return err
	}

	if prevData.ID == 0 {
		s.log.Error("Failed employee not found")
		return errors.New("employee not found")
	}

	return s.employeeRepository.Delete(id)
}
