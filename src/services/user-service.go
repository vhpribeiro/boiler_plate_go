package services

import (
	"boiler_plate.com/src/repositories"
	"boiler_plate.com/src/services/dtos"
)

type IUserService interface {
	CreateUser(loginDto dtos.LoginDto) dtos.UserDto
}

type userService struct {
	userRepository repositories.IUserRepository
}

func (u *userService) CreateUser(loginDto dtos.LoginDto) dtos.UserDto {
	user := u.userRepository.CreateUser(loginDto.Username, loginDto.Password)

	userDto := dtos.UserDto{
		Username:  user.Username,
		Password:  user.Password,
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	return userDto
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}
