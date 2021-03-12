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

type MockRedisRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUser(username, password string) models.User {
	args := m.Called(username, password)
	result := args.Get(0)
	return result.(models.User)
}

func (m *MockUserRepository) CreateUser(username, password string) models.User {
	args := m.Called()
	result := args.Get(0)
	return result.(models.User)
}

func (m *MockRedisRepository) Save(key, value string) error {
	args := m.Called(key, value)
	return args.Error(0)
}

func (m *MockRedisRepository) Get(key string) (string, error) {
	args := m.Called(key)
	result := args.Get(0)
	return result.(string), args.Error(1)
}

func TestShouldGetErrorWhenTryLoginAndUserWasNotFound(t *testing.T) {
	expectedError := &errors.InternalError{Message: "User not found!"}
	loginDto := dtos.LoginDto{
		Username: "",
		Password: "",
	}
	invalidUser := models.User{
		Model:    gorm.Model{ID: 0},
		Username: "",
		Password: "",
	}
	userNotFoundInCache := ""
	mockUserRepository := new(MockUserRepository)
	mockRedisRepository := new(MockRedisRepository)
	mockUserRepository.On("GetUser", loginDto.Username, loginDto.Password).Return(invalidUser)
	mockRedisRepository.On("Get", loginDto.Username).Return(userNotFoundInCache, nil)
	loginService := NewLoginService(mockUserRepository, mockRedisRepository)

	_, resultError := loginService.Login(loginDto)

	assert.Equal(t, expectedError, resultError)
}

func TestShouldGetErrorWhenTryLoginAndCanNotSaveInCacheRepository(t *testing.T) {
	expectedError := &errors.InternalError{Message: "Can not save in cache repository!"}
	username := "Vitor Ribeiro"
	password := "123458678"
	loginDto := dtos.LoginDto{Username: username, Password: password}
	validUser := models.User{
		Model:    gorm.Model{ID: 1},
		Username: username,
		Password: password,
	}
	userNotFoundInCache := ""
	mockUserRepository := new(MockUserRepository)
	mockRedisRepository := new(MockRedisRepository)
	mockUserRepository.On("GetUser", loginDto.Username, loginDto.Password).Return(validUser)
	mockRedisRepository.On("Get", loginDto.Username).Return(userNotFoundInCache, nil)
	mockRedisRepository.On("Save", validUser.Username, validUser.Password).Return(expectedError)
	loginService := NewLoginService(mockUserRepository, mockRedisRepository)

	_, err := loginService.Login(loginDto)

	assert.Equal(t, expectedError, err)
}

func TestShouldGetTheTokenWhenDoLoginAndUserIsInCache(t *testing.T) {
	username := "Vitor Ribeiro"
	password := "123458678"
	loginDto := dtos.LoginDto{Username: username, Password: password}
	tokenGenerated, _ := generateValidToken(username)
	expectedResult := map[string]string{
		"token": tokenGenerated,
	}
	mockUserRepository := new(MockUserRepository)
	mockRedisRepository := new(MockRedisRepository)
	mockRedisRepository.On("Get", loginDto.Username).Return(username, nil)
	loginService := NewLoginService(mockUserRepository, mockRedisRepository)

	result, _ := loginService.Login(loginDto)

	assert.Equal(t, expectedResult, result)
}

func TestShouldGetTheTokenWhenDoLoginAndUserIsNotInCache(t *testing.T) {
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
	userNotFoundInCache := ""
	mockUserRepository := new(MockUserRepository)
	mockRedisRepository := new(MockRedisRepository)
	mockUserRepository.On("GetUser", loginDto.Username, loginDto.Password).Return(validUser)
	mockRedisRepository.On("Get", loginDto.Username).Return(userNotFoundInCache, nil)
	mockRedisRepository.On("Save", validUser.Username, validUser.Password).Return(nil)
	loginService := NewLoginService(mockUserRepository, mockRedisRepository)

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
