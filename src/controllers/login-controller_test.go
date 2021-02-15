package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"boiler_plate.com/src/services"
	"boiler_plate.com/src/services/dtos"
	"boiler_plate.com/src/utils/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLoginSerevice struct {
	mock.Mock
}

func (m *MockLoginSerevice) Login(loginDto dtos.LoginDto) (map[string]string, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(map[string]string), args.Error(1)
}

func TestShouldGetSuccessWhenGetLogin(t *testing.T) {
	var result map[string]string
	expectedResult := map[string]string{
		"token": "testesadahjsuahdhuaeyOk",
	}
	url := "/login"
	e := echo.New()
	requestBody := `{"username": "jon", "password": "password"}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockLoginService := new(MockLoginSerevice)
	mockLoginService.On("Login").Return(expectedResult, nil)
	loginController := NewLoginController(mockLoginService)

	loginController.Login(c)

	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expectedResult, result)
}

func TestShouldGetInvalidObjectWhenGetLogin(t *testing.T) {
	var result errors.InvalidObjectError
	expectedResult := errors.InvalidObjectError{}
	url := "/login"
	e := echo.New()
	requestBody := `{"TESTE": "teste_falho"`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockLoginService := new(MockLoginSerevice)
	loginController := NewLoginController(mockLoginService)

	loginController.Login(c)

	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, expectedResult, result)
}

func TestShouldGetInternalServerErrorWhenGetLogin(t *testing.T) {
	var result errors.InternalError
	expectedResult := errors.InternalError{Message: "It was not possible get the login"}
	mapResult := map[string]string{
		"token": "Invalido",
	}
	url := "/login"
	e := echo.New()
	requestBody := `{"username": "jon", "password": "password"}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockLoginService := new(MockLoginSerevice)
	mockLoginService.On("Login").Return(mapResult, &expectedResult)
	loginController := NewLoginController(mockLoginService)

	loginController.Login(c)

	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, expectedResult, result)
}

func TestShouldAccessRestrictedArea(t *testing.T) {
	var result string
	expectedResult := "You can pass young jon!"
	mockLoginService := new(MockLoginSerevice)
	loginController := NewLoginController(mockLoginService)
	token := getTheToken()
	bearer := "Bearer " + token
	url := "/login/restricted"
	e := echo.New()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", bearer)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	loginController.RestrictedAccess(c)

	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, expectedResult, result)
}

func getTheToken() string {
	var result map[string]string
	url := "/login"
	e := echo.New()
	requestBody := `{"username": "jon", "password": "password"}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	loginService := services.NewLoginService()
	loginController := NewLoginController(loginService)

	loginController.Login(c)
	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	return result["token"]
}
