package router

import (
	"github.com/aliirsyaadn/mekari-technical-assessment/handler"
	"github.com/gin-gonic/gin"
)

type employeeRouter struct {
	routerGroup     *gin.RouterGroup
	employeeHandler handler.EmployeeHandler
}

func NewEmployeeRouter(
	routerGroup *gin.RouterGroup,
	employeeHandler handler.EmployeeHandler,
) Router {
	return &employeeRouter{
		routerGroup:     routerGroup,
		employeeHandler: employeeHandler,
	}
}

func (r *employeeRouter) Routers() {
	r.routerGroup.GET("", r.employeeHandler.FindAll)
	r.routerGroup.GET(":id", r.employeeHandler.FindByID)
	r.routerGroup.POST("", r.employeeHandler.Save)
	r.routerGroup.PUT(":id", r.employeeHandler.Update)
	r.routerGroup.DELETE(":id", r.employeeHandler.Delete)
}
