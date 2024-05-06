package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aliirsyaadn/mekari-technical-assessment/config"
	"github.com/aliirsyaadn/mekari-technical-assessment/dummy"
	"github.com/aliirsyaadn/mekari-technical-assessment/entity"
	"github.com/aliirsyaadn/mekari-technical-assessment/model"
	mockService "github.com/aliirsyaadn/mekari-technical-assessment/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestEmployeeHandler_FindAll_success_with_data(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeService := mockService.NewMockEmployeeService(ctrl)
	employeeHandler := NewEmployeeHandler(config.InitLogger(), employeeService)

	employeeService.EXPECT().FindAll().Return(dummy.DummyEmployees, nil)

	r := SetUpRouter()
	r.GET("/employees", employeeHandler.FindAll)
	req, _ := http.NewRequest("GET", "/employees", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "success", responseData.Message)
	dataBytes, _ := json.Marshal(responseData.Data)
	var employees []model.Employee
	json.Unmarshal(dataBytes, &employees)

	assert.Equal(t, employees, dummy.DummyEmployees)
}

func TestEmployeeHandler_FindAll_success_no_data(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeService := mockService.NewMockEmployeeService(ctrl)
	employeeHandler := NewEmployeeHandler(config.InitLogger(), employeeService)

	employeeService.EXPECT().FindAll().Return([]model.Employee{}, nil)

	r := SetUpRouter()
	r.GET("/employees", employeeHandler.FindAll)
	req, _ := http.NewRequest("GET", "/employees", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "success no data", responseData.Message)
}

func TestEmployeeHandler_FindAll_failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeService := mockService.NewMockEmployeeService(ctrl)
	employeeHandler := NewEmployeeHandler(config.InitLogger(), employeeService)

	employeeService.EXPECT().FindAll().Return(nil, assert.AnError)

	r := SetUpRouter()
	r.GET("/employees", employeeHandler.FindAll)
	req, _ := http.NewRequest("GET", "/employees", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 400, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "Failed to find all employees", responseData.Message)
}

func TestEmployeeHandler_FindByID_success_with_data(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeService := mockService.NewMockEmployeeService(ctrl)
	employeeHandler := NewEmployeeHandler(config.InitLogger(), employeeService)

	employeeService.EXPECT().FindByID(1).Return(dummy.DummyEmployees[0], nil)

	r := SetUpRouter()
	r.GET("/employees/:id", employeeHandler.FindByID)
	req, _ := http.NewRequest("GET", "/employees/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "success", responseData.Message)

	dataBytes, _ := json.Marshal(responseData.Data)
	var employee model.Employee
	json.Unmarshal(dataBytes, &employee)

	assert.Equal(t, employee, dummy.DummyEmployees[0])
}

func TestEmployeeHandler_FindByID_success_no_data(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeService := mockService.NewMockEmployeeService(ctrl)
	employeeHandler := NewEmployeeHandler(config.InitLogger(), employeeService)

	employeeService.EXPECT().FindByID(1).Return(model.Employee{}, nil)

	r := SetUpRouter()
	r.GET("/employees/:id", employeeHandler.FindByID)
	req, _ := http.NewRequest("GET", "/employees/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "success no data", responseData.Message)
}

func TestEmployeeHandler_FindByID_failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeService := mockService.NewMockEmployeeService(ctrl)
	employeeHandler := NewEmployeeHandler(config.InitLogger(), employeeService)

	employeeService.EXPECT().FindByID(1).Return(model.Employee{}, assert.AnError)

	r := SetUpRouter()
	r.GET("/employees/:id", employeeHandler.FindByID)
	req, _ := http.NewRequest("GET", "/employees/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 400, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "Failed to find employee by id", responseData.Message)

}

func TestEmployeeHandler_FindByID_failed_id_not_number(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeHandler := NewEmployeeHandler(config.InitLogger(), nil)

	r := SetUpRouter()
	r.GET("/employees/:id", employeeHandler.FindByID)
	req, _ := http.NewRequest("GET", "/employees/abc", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 400, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "id must be a number", responseData.Message)
}

func TestEmployeeHandler_Save_success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeService := mockService.NewMockEmployeeService(ctrl)
	employeeHandler := NewEmployeeHandler(config.InitLogger(), employeeService)

	employeeService.EXPECT().Save(&dummy.DummyEmployees[0]).Return(nil)

	r := SetUpRouter()
	r.POST("/employees", employeeHandler.Save)

	jsonValue, _ := json.Marshal(dummy.DummyEmployees[0])
	req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "success", responseData.Message)

	dataBytes, _ := json.Marshal(responseData.Data)
	var employee model.Employee
	json.Unmarshal(dataBytes, &employee)

	assert.Equal(t, employee, dummy.DummyEmployees[0])

}

func TestEmployeeHandler_Save_failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeService := mockService.NewMockEmployeeService(ctrl)
	employeeHandler := NewEmployeeHandler(config.InitLogger(), employeeService)

	employeeService.EXPECT().Save(&dummy.DummyEmployees[0]).Return(assert.AnError)

	r := SetUpRouter()
	r.POST("/employees", employeeHandler.Save)

	jsonValue, _ := json.Marshal(dummy.DummyEmployees[0])
	req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 400, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "Failed to save employee", responseData.Message)
}

func TestEmployeeHandler_Save_failed_bind_error(t *testing.T) {
	employeeHandler := NewEmployeeHandler(config.InitLogger(), nil)

	r := SetUpRouter()
	r.POST("/employees", employeeHandler.Save)

	// email is required
	data := map[string]any{
		"first_name": "Ali",
		"last_name":  "Irsyaad",
	}
	jsonValue, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 400, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "invalid request", responseData.Message)
}

func TestEmployeeHandler_Update_success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeService := mockService.NewMockEmployeeService(ctrl)
	employeeHandler := NewEmployeeHandler(config.InitLogger(), employeeService)

	employeeService.EXPECT().Update(&dummy.DummyEmployees[0]).Return(nil)

	r := SetUpRouter()
	r.PUT("/employees/:id", employeeHandler.Update)

	jsonValue, _ := json.Marshal(dummy.DummyEmployees[0])
	req, _ := http.NewRequest("PUT", "/employees/1", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "success", responseData.Message)

	dataBytes, _ := json.Marshal(responseData.Data)
	var employee model.Employee
	json.Unmarshal(dataBytes, &employee)

	assert.Equal(t, employee, dummy.DummyEmployees[0])
}

func TestEmployeeHandler_Update_failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeService := mockService.NewMockEmployeeService(ctrl)
	employeeHandler := NewEmployeeHandler(config.InitLogger(), employeeService)

	employeeService.EXPECT().Update(&dummy.DummyEmployees[0]).Return(assert.AnError)

	r := SetUpRouter()
	r.PUT("/employees/:id", employeeHandler.Update)

	jsonValue, _ := json.Marshal(dummy.DummyEmployees[0])
	req, _ := http.NewRequest("PUT", "/employees/1", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 400, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "Failed to update employee", responseData.Message)
}

func TestEmployeeHandler_Update_failed_bind_error(t *testing.T) {
	employeeHandler := NewEmployeeHandler(config.InitLogger(), nil)

	r := SetUpRouter()
	r.PUT("/employees/:id", employeeHandler.Update)

	// email is required
	data := map[string]any{
		"first_name": "Ali",
		"last_name":  "Irsyaad",
	}
	jsonValue, _ := json.Marshal(data)
	req, _ := http.NewRequest("PUT", "/employees/1", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 400, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "invalid request", responseData.Message)
}

func TestEmployeeHandler_Delete_failed_id_not_number(t *testing.T) {
	employeeHandler := NewEmployeeHandler(config.InitLogger(), nil)

	r := SetUpRouter()
	r.DELETE("/employees/:id", employeeHandler.Delete)
	req, _ := http.NewRequest("DELETE", "/employees/abc", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 400, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "id must be a number", responseData.Message)
}

func TestEmployeeHandler_Delete_success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeService := mockService.NewMockEmployeeService(ctrl)
	employeeHandler := NewEmployeeHandler(config.InitLogger(), employeeService)

	employeeService.EXPECT().Delete(1).Return(nil)

	r := SetUpRouter()
	r.DELETE("/employees/:id", employeeHandler.Delete)
	req, _ := http.NewRequest("DELETE", "/employees/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "success", responseData.Message)
}

func TestEmployeeHandler_Delete_failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	employeeService := mockService.NewMockEmployeeService(ctrl)
	employeeHandler := NewEmployeeHandler(config.InitLogger(), employeeService)

	employeeService.EXPECT().Delete(1).Return(assert.AnError)

	r := SetUpRouter()
	r.DELETE("/employees/:id", employeeHandler.Delete)
	req, _ := http.NewRequest("DELETE", "/employees/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseData entity.Response
	json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.Equal(t, 400, w.Code)
	assert.NotEmpty(t, responseData)
	assert.Equal(t, "Failed to delete employee", responseData.Message)
}
