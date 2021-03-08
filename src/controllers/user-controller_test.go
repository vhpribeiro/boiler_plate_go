package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"boiler_plate.com/src/services/dtos"
	"boiler_plate.com/src/utils/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(loginDto dtos.LoginDto) dtos.UserDto {
	args := m.Called(loginDto)
	result := args.Get(0)
	return result.(dtos.UserDto)
}

func TestShouldGetBadRequestWhenPassInvalidObjectInput(t *testing.T) {
	var result errors.InvalidObjectError
	expectedResult := errors.InvalidObjectError{}
	testCases := []struct {
		username string
		password string
	}{
		{"", "12345648"},
		{"       ", "12345648"},
		{"jon", ""},
		{"jon", "     "},
	}
	for _, tc := range testCases {
		t.Run("test invalid object", func(t *testing.T) {
			requestBody := createRequestBody(tc.username, tc.password)
			c, rec := createContextAndRecorder(requestBody)
			mockUserService := new(MockUserService)
			userController := NewUserController(mockUserService)

			userController.CreateUser(c)

			_ = json.Unmarshal(rec.Body.Bytes(), &result)
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, expectedResult, result)
		})
	}
}

func TestShouldGetBadRequestWhenBodyRequestIsInvalid(t *testing.T) {
	var result errors.InvalidObjectError
	expectedResult := errors.InvalidObjectError{}
	requestBody := `{"TESTE": "teste_falho"`
	c, rec := createContextAndRecorder(requestBody)
	mockUserService := new(MockUserService)
	userController := NewUserController(mockUserService)

	userController.CreateUser(c)

	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, expectedResult, result)
}

func TestShoulGetTheCreatedUser(t *testing.T) {
	var result dtos.UserDto
	username := "Vitor Ribeiro"
	password := "senha_123"
	expectedResult := dtos.UserDto{
		Username:  username,
		Password:  password,
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: &time.Time{},
	}
	loginDto := dtos.LoginDto{
		Username: username,
		Password: password,
	}
	requestBody := fmt.Sprintf("{\"username\": \"%s\", \"password\": \"%s\"}", username, password)
	c, rec := createContextAndRecorder(requestBody)
	mockUserService := new(MockUserService)
	mockUserService.On("CreateUser", loginDto).Return(expectedResult)
	userController := NewUserController(mockUserService)

	userController.CreateUser(c)

	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, expectedResult, result)
}

func createRequestBody(username, password string) string {
	return fmt.Sprintf("{\"username\": \"%s\", \"password\": \"%s\"}", username, password)
}

func createContextAndRecorder(requestBody string) (echo.Context, *httptest.ResponseRecorder) {
	url := "/login"
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec
}
