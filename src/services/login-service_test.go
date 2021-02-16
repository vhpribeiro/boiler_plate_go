package services

import (
	"testing"
	"time"

	"boiler_plate.com/src/models"
	"boiler_plate.com/src/services/dtos"
	"boiler_plate.com/src/utils/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUser(username, password string) models.User {
	args := m.Called()
	result := args.Get(0)
	return result.(models.User)
}

func TestShouldGetErrorWhenTryLoginAndUserWasNotFound(t *testing.T) {
	expectedError := &errors.InternalError{Message: "User not found!"}
	loginDto := dtos.LoginDto{Username: ""}
	invalidUser := models.User{
		Model:    gorm.Model{ID: 0},
		Username: "",
		Password: "",
	}
	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("GetUser").Return(invalidUser)
	loginService := NewLoginService(mockUserRepository)

	_, resultError := loginService.Login(loginDto)

	assert.Equal(t, expectedError, resultError)
}

func TestShouldGetTheTokenWhenDoLogin(t *testing.T) {
	username := "Vitor Ribeiro"
	password := "123458678"
	loginDto := dtos.LoginDto{Username: username, Password: password}
	validUser := models.User{
		Model:    gorm.Model{ID: 1},
		Username: username,
		Password: password,
	}
	tokenGenerated, _ := generateValidToken(username)
	expectedResult := map[string]string{
		"token": tokenGenerated,
	}
	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("GetUser").Return(validUser)
	loginService := NewLoginService(mockUserRepository)

	result, _ := loginService.Login(loginDto)

	assert.Equal(t, expectedResult, result)
}

func generateValidToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte("secret"))
}
