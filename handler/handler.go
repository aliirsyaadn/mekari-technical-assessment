package handler

//go:generate mockgen -source=handler.go -destination=mock/handler.go

import (
	"github.com/gin-gonic/gin"
)

type EmployeeHandler interface {
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
	Save(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
