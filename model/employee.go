package model

import (
	"time"
)

type Employee struct {
	ID        int       `gorm:"primaryKey;type:int;autoIncrement;" json:"id"`
	FirstName string    `gorm:"not null" json:"first_name" binding:"required"`
	LastName  string    `gorm:"not null" json:"last_name" binding:"required"`
	Email     string    `gorm:"not null;unique" json:"email" binding:"required"`
	HireDate  Date      `gorm:"not null;type:date" json:"hire_date" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
