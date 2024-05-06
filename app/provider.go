package app

import (
	"github.com/aliirsyaadn/mekari-technical-assessment/config"
	"github.com/aliirsyaadn/mekari-technical-assessment/handler"
	"github.com/aliirsyaadn/mekari-technical-assessment/repository"
	"github.com/aliirsyaadn/mekari-technical-assessment/router"
	"github.com/aliirsyaadn/mekari-technical-assessment/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var serverSet = wire.NewSet(
	config.InitDatabase,
	config.InitLogger,
	repository.NewEmployeeRepository,
	service.NewEmployeeService,
	handler.NewEmployeeHandler,
	provideRouter,
)

func provideRouter(
	employeeHandler handler.EmployeeHandler,
) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	router.NewEmployeeRouter(r.Group("/employees"), employeeHandler).Routers()

	return r
}
