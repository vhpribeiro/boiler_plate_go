package controllers

import (
	"net/http"

	"boiler_plate.com/src/services"
	"boiler_plate.com/src/services/dtos"
	"boiler_plate.com/src/utils/errors"
	"github.com/labstack/echo/v4"
)

type ILoginController interface {
	Login(c echo.Context) error
}

type loginController struct {
	loginService services.ILoginService
}

func (l *loginController) Login(c echo.Context) error {
	var loginDto dtos.LoginDto
	err := c.Bind(&loginDto)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &errors.InvalidObjectError{})
	}

	result, err := l.loginService.Login(loginDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &errors.InternalError{Message: "It was not possible get the login"})
	}

	return c.JSON(http.StatusOK, result)
}

func NewLoginController(loginService services.ILoginService) ILoginController {
	return &loginController{
		loginService: loginService,
	}
}
