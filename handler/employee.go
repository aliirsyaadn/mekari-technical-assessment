package handler

import (
	"strconv"

	"github.com/aliirsyaadn/mekari-technical-assessment/entity"
	"github.com/aliirsyaadn/mekari-technical-assessment/model"
	"github.com/aliirsyaadn/mekari-technical-assessment/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type employeeHandler struct {
	log             *zap.SugaredLogger
	employeeService service.EmployeeService
}

func NewEmployeeHandler(
	log *zap.SugaredLogger,
	employeeService service.EmployeeService,
) EmployeeHandler {
	return &employeeHandler{
		log:             log,
		employeeService: employeeService,
	}
}

func (h *employeeHandler) FindAll(c *gin.Context) {
	employees, err := h.employeeService.FindAll()
	if err != nil {
		h.log.Error("Failed to find all employees:", err.Error())
		c.JSONP(400, entity.Response{
			Message: "Failed to find all employees",
		})
		return
	}

	if len(employees) == 0 {
		h.log.Info("No data found")
		c.JSONP(200, entity.Response{
			Message: "success no data",
		})
		return
	}

	c.JSONP(200, entity.Response{
		Message: "success",
		Data:    employees,
	})
}

func (h *employeeHandler) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("Failed to convert id to int:", err.Error())
		c.JSONP(400, entity.Response{
			Message: "id must be a number",
		})
		return
	}

	employee, err := h.employeeService.FindByID(id)
	if err != nil {
		h.log.Error("Failed to find employee by id:", err.Error())
		c.JSONP(400, entity.Response{
			Message: "Failed to find employee by id",
		})
		return
	}

	if employee.ID == 0 {
		c.JSONP(200, entity.Response{
			Message: "success no data",
		})
		return
	}

	c.JSONP(200, entity.Response{
		Message: "success",
		Data:    employee,
	})
}

func (h *employeeHandler) Save(c *gin.Context) {
	var employee model.Employee
	err := c.BindJSON(&employee)
	if err != nil {
		h.log.Error("Failed to bind json:", err.Error())
		c.JSONP(400, entity.Response{
			Message: "invalid request",
		})
		return
	}

	err = h.employeeService.Save(&employee)
	if err != nil {
		h.log.Error("Failed to save employee:", err.Error())
		c.JSONP(400, entity.Response{
			Message: "Failed to save employee",
		})
		return
	}

	c.JSONP(200, entity.Response{
		Message: "success",
		Data:    employee,
	})

}

func (h *employeeHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("Failed to convert id to int:", err.Error())
		c.JSONP(400, entity.Response{
			Message: "id must be a number",
		})
		return
	}

	var employee model.Employee
	err = c.BindJSON(&employee)
	if err != nil {
		h.log.Error("Failed to bind json:", err.Error())
		c.JSONP(400, entity.Response{
			Message: "invalid request",
		})
		return
	}

	employee.ID = id

	err = h.employeeService.Update(&employee)
	if err != nil {
		h.log.Error("Failed to update employee:", err.Error())
		c.JSONP(400, entity.Response{
			Message: "Failed to update employee",
		})
		return
	}

	c.JSONP(200, entity.Response{
		Message: "success",
		Data:    employee,
	})
}

func (h *employeeHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("Failed to convert id to int:", err.Error())
		c.JSONP(400, entity.Response{
			Message: "id must be a number",
		})
		return
	}

	err = h.employeeService.Delete(id)
	if err != nil {
		h.log.Error("Failed to delete employee:", err.Error())
		c.JSONP(400, entity.Response{
			Message: "Failed to delete employee",
		})
		return
	}

	c.JSONP(200, entity.Response{
		Message: "success",
	})
}
