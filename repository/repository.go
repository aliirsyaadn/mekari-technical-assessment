package repository

//go:generate mockgen -source=repository.go -destination=mock/repository.go

import (
	"github.com/aliirsyaadn/mekari-technical-assessment/model"
)

type EmployeeRepository interface {
	FindAll() ([]model.Employee, error)
	FindByID(id int) (model.Employee, error)
	Save(employee *model.Employee) error
	Update(employee *model.Employee) error
	Delete(id int) error
}
