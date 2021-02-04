package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"controle_acesso_core.com/src/services/dtos"
	"controle_acesso_core.com/src/utils/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPolicyService struct {
	mock.Mock
}

func (mock *MockPolicyService) AddPolicy(policyDto dtos.PolicyDto) (bool, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(bool), args.Error(1)
}

func TestShouldGetPolicyCreatedWhenAddPolicy(t *testing.T) {
	//Arrange
	var result dtos.PolicyDto
	expectedResult := dtos.PolicyDto{Role: "cadastro_lider", Domain: "administrativo", Action: "consultar", Resource: "cliente"}
	url := "/policies"
	e := echo.New()
	requestBody := `{"role": "cadastro_lider", "domain": "administrativo", "action": "consultar", "resource": "cliente"}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(requestBody))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	MockPolicyService := new(MockPolicyService)
	MockPolicyService.On("AddPolicy").Return(true, nil)
	policyController := NewPolicyController(MockPolicyService)

	//Action
	policyController.AddPolicy(c)

	//Assert
	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, expectedResult, result)
}

func TestShouldGetInternalErrorWhenAddPolicy(t *testing.T) {
	//Arrange
	var result errors.Error
	expectedResult := errors.Error{Message: "This policy could not be created"}
	url := "/policies"
	e := echo.New()
	requestBody := `{"role": "cadastro_lider", "domain": "administrativo", "action": "consultar", "resource": "cliente"}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(requestBody))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	MockPolicyService := new(MockPolicyService)
	MockPolicyService.On("AddPolicy").Return(false, nil)
	policyController := NewPolicyController(MockPolicyService)

	//Action
	policyController.AddPolicy(c)

	//Assert
	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, expectedResult, result)
}

func TestShouldGetBadRequestWhenAddPolicy(t *testing.T) {
	//Arrange
	var result errors.Error
	expectedResult := errors.Error{Message: "Invalid object"}
	url := "/policies"
	e := echo.New()
	requestBody := `{"role": "cadastro_lider", "domain": "administrativo"}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(requestBody))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	MockPolicyService := new(MockPolicyService)
	MockPolicyService.On("AddPolicy").Return(false, nil)
	policyController := NewPolicyController(MockPolicyService)

	//Action
	policyController.AddPolicy(c)

	//Assert
	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, expectedResult, result)
}
