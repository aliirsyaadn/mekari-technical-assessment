package repository

import (
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aliirsyaadn/mekari-technical-assessment/config"
	"github.com/aliirsyaadn/mekari-technical-assessment/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}

func TestEmployeeRepository_NewEmployeeRepository(t *testing.T) {
	db, _ := NewMockDB()
	employeeRepository := NewEmployeeRepository(config.InitLogger(), db)
	assert.NotNil(t, employeeRepository)
}

func TestEmployeeRepository_FindAll(t *testing.T) {
	db, mock := NewMockDB()
	employeeRepository := NewEmployeeRepository(config.InitLogger(), db)

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email"}).
		AddRow(1, "Ali", "Irsyaad", "aliirsyaadn@gmail.com").
		AddRow(2, "Ali2", "Irsyaad", "aliirsyaadn2@gmail.com")

	mock.ExpectQuery(`^SELECT (.+) FROM "employees"$`).WillReturnRows(rows)

	employees, err := employeeRepository.FindAll()
	assert.Nil(t, err)
	assert.NotNil(t, employees)
	assert.Equal(t, 2, len(employees))
}

func TestEmployeeRepository_FindByID(t *testing.T) {
	db, mock := NewMockDB()
	employeeRepository := NewEmployeeRepository(config.InitLogger(), db)

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email"}).
		AddRow(1, "Ali", "Irsyaad", "aliirsyaadn@gmail.com")

	mock.ExpectQuery(`^SELECT (.+) FROM "employees" WHERE "employees"."id" = (.+) ORDER BY "employees"."id" LIMIT (.+)`).WithArgs(1, 1).WillReturnRows(rows)

	employee, err := employeeRepository.FindByID(1)
	assert.Nil(t, err)
	assert.NotNil(t, employee)
	assert.Equal(t, 1, employee.ID)

	// Not Found
	rows = sqlmock.NewRows([]string{"id", "first_name", "last_name", "email"})
	mock.ExpectQuery(`^SELECT (.+) FROM "employees" WHERE "employees"."id" = (.+) ORDER BY "employees"."id" LIMIT (.+)`).WithArgs(2, 1).WillReturnRows(rows)

	employee, err = employeeRepository.FindByID(2)
	assert.Nil(t, err)
	assert.NotNil(t, employee)
	assert.Equal(t, 0, employee.ID)
}

func TestEmployeeRepository_Save(t *testing.T) {
	db, mock := NewMockDB()
	employeeRepository := NewEmployeeRepository(config.InitLogger(), db)

	now := time.Now()

	addRow := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email"}).
		AddRow(1, "Ali", "Irsyaad", "aliirsyaadn@gmail.com")

	mock.ExpectBegin()
	mock.ExpectQuery(`^INSERT INTO "employees" (.+) VALUES (.+) RETURNING "id"$`).WillReturnRows(addRow)
	mock.ExpectCommit()

	err := employeeRepository.Save(&model.Employee{
		FirstName: "Ali",
		LastName:  "Irsyaad",
		Email:     "aliirsyaadn@gmail.com",
		HireDate:  model.Date{Time: now},
	})

	assert.Nil(t, err)
}

func TestEmployeeRepository_Update(t *testing.T) {
	db, mock := NewMockDB()
	employeeRepository := NewEmployeeRepository(config.InitLogger(), db)

	mock.ExpectBegin()
	mock.ExpectExec(`^UPDATE "employees" SET (.+) WHERE "id" = (.+)$`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := employeeRepository.Update(&model.Employee{
		ID:        1,
		FirstName: "Ali",
		LastName:  "Irsyaad",
		Email:     "aliirsyaadn@gmail.com",
	})

	assert.Nil(t, err)
}

func TestEmployeeRepository_Delete(t *testing.T) {
	db, mock := NewMockDB()
	employeeRepository := NewEmployeeRepository(config.InitLogger(), db)

	mock.ExpectBegin()
	mock.ExpectExec(`^DELETE FROM "employees" WHERE id = (.+)$`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := employeeRepository.Delete(1)
	assert.Nil(t, err)
}
