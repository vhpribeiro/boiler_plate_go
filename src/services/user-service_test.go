package services

import (
	"testing"
	"time"

	"boiler_plate.com/src/models"
	"boiler_plate.com/src/services/dtos"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateUser(t *testing.T) {
	username := "Vitor Ribeiro"
	password := "fake_password"
	userModel := models.User{
		gorm.Model{
			ID:        1,
			CreatedAt: time.Now().AddDate(0, 0, 1),
			UpdatedAt: time.Now().AddDate(0, 0, 2),
		},
		username,
		password,
	}
	expectedResult := dtos.UserDto{
		Username:  userModel.Username,
		Password:  userModel.Password,
		ID:        userModel.ID,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
		DeletedAt: userModel.DeletedAt,
	}
	loginDto := dtos.LoginDto{
		Username: username,
		Password: password,
	}
	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("CreateUser").Return(userModel)
	userService := NewUserService(mockUserRepository)

	result := userService.CreateUser(loginDto)

	assert.Equal(t, expectedResult, result)
}
