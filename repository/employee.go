package repository

import (
	"github.com/aliirsyaadn/mekari-technical-assessment/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type employeeRepository struct {
	log *zap.SugaredLogger
	db  *gorm.DB
}

func NewEmployeeRepository(
	log *zap.SugaredLogger,
	db *gorm.DB,
) EmployeeRepository {
	return &employeeRepository{
		log: log,
		db:  db,
	}
}

func (r *employeeRepository) FindAll() ([]model.Employee, error) {
	var employees []model.Employee
	err := r.db.Find(&employees).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.log.Info("Employees not found")
			return nil, nil
		}
		r.log.Error("Failed to find all employees: ", err.Error())
		return nil, err
	}
	return employees, nil
}

func (r *employeeRepository) FindByID(id int) (model.Employee, error) {
	var employee model.Employee
	err := r.db.First(&employee, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.log.Info("Employee not found")
			return model.Employee{}, nil
		}
		r.log.Error("Failed to find employee by id: ", err)
		return model.Employee{}, err
	}
	return employee, nil
}

func (r *employeeRepository) Save(employee *model.Employee) error {
	err := r.db.Create(employee).Error
	if err != nil {
		r.log.Error("Failed to save employee: ", err)
		return err
	}
	return nil
}

func (r *employeeRepository) Update(employee *model.Employee) error {
	err := r.db.Omit("CreatedAt").Updates(employee).Error
	if err != nil {
		r.log.Error("Failed to update employee: ", err)
		return err
	}
	return nil
}

func (r *employeeRepository) Delete(id int) error {
	result := r.db.Where("id = ?", id).Unscoped().Delete(&model.Employee{})
	if result.Error != nil {
		r.log.Error("Failed to delete employee: ", result.Error)
		return result.Error
	}
	return nil
}
