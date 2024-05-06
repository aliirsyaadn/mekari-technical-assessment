package config

import (
	"github.com/aliirsyaadn/mekari-technical-assessment/model"
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Employee{},
	)
	if err != nil {
		panic(err)
	}
}
