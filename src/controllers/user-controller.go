package controllers

import (
	"net/http"

	"boiler_plate.com/src/errors"
	"boiler_plate.com/src/helpers"
	"boiler_plate.com/src/services"
	"boiler_plate.com/src/services/dtos"
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

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if helpers.IsEmptyOrWhiteSpace(loginDto.Username) || helpers.IsEmptyOrWhiteSpace(loginDto.Password) {
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
