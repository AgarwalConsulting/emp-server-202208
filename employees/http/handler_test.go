package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

func TestCreateV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)
	sut := empHTTP.New(mockSvc)

	expectedEmp := entities.Employee{Name: "Gaurav", Department: "LnD", ProjectID: 1001}

	mockSvc.EXPECT().Create(expectedEmp).Return(&expectedEmp, nil)

	respRec := httptest.NewRecorder()

	jsonBody := `{"name":"Gaurav", "speciality": "LnD", "project": 1001}`
	req := httptest.NewRequest("POST", "/v1/employees", strings.NewReader(jsonBody))

	// sut.CreateV1(respRec, req)
	sut.ServeHTTP(respRec, req)

	res := respRec.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var actualEmp entities.Employee

	err := json.NewDecoder(res.Body).Decode(&actualEmp)

	assert.Nil(t, err)

	assert.Equal(t, expectedEmp.Name, actualEmp.Name)
}
