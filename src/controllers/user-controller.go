package controllers

import (
	"net/http"

	"boiler_plate.com/src/services"
	"boiler_plate.com/src/services/dtos"
	"boiler_plate.com/src/utils/errors"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	CreateUser(c echo.Context) error
}

type userController struct {
	userService services.IUserService
}

func (u *userController) CreateUser(c echo.Context) error {
	var loginDto dtos.LoginDto
	err := c.Bind(&loginDto)

	if err != nil || loginDto.Username == "" || loginDto.Password == "" {
		return c.JSON(http.StatusBadRequest, &errors.InvalidObjectError{})
	}

	createdUser := u.userService.CreateUser(loginDto)
	return c.JSON(http.StatusCreated, createdUser)
}

func NewUserController(userService services.IUserService) IUserController {
	return &userController{
		userService: userService,
	}
}
