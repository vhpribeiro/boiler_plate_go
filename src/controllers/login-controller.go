package controllers

import (
	"net/http"

	"boiler_plate.com/src/services"
	"boiler_plate.com/src/services/dtos"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type ILoginController interface {
	Login(c echo.Context) error
	RestrictedAccess(c echo.Context) error
}

type loginController struct {
	loginService services.ILoginService
}

func (l *loginController) Login(c echo.Context) error {
	var loginDto dtos.LoginDto
	err := c.Bind(&loginDto)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := l.loginService.Login(loginDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (l *loginController) RestrictedAccess(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func NewLoginController(loginService services.ILoginService) ILoginController {
	return &loginController{
		loginService: loginService,
	}
}
