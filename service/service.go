package service

//go:generate mockgen -source=service.go -destination=mock/service.go

import (
	"github.com/aliirsyaadn/mekari-technical-assessment/model"
)

type EmployeeService interface {
	FindAll() ([]model.Employee, error)
	FindByID(id int) (model.Employee, error)
	Save(employee *model.Employee) error
	Update(employee *model.Employee) error
	Delete(id int) error
}
