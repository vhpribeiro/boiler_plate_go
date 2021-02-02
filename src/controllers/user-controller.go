package controllers

import (
	"encoding/json"
	"net/http"

	"controle_acesso_core.com/src/errors"
	"controle_acesso_core.com/src/services"
	"controle_acesso_core.com/src/services/dtos"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	AddRoleForUserInDomain(context echo.Context) error
	CheckIfUserHasPermission(context echo.Context) error
}

type userController struct {
	userService services.IUserService
}

func (userController *userController) AddRoleForUserInDomain(context echo.Context) error {
	var roleDomainDto dtos.RoleDomainDto
	user := context.Param("userName")
	responseWriter := context.Response().Writer
	body := context.Request().Body
	err := json.NewDecoder(body).Decode(&roleDomainDto)

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return err
	}

	result, err := userController.userService.AddRoleForUserInDomain(user, roleDomainDto.Domain, roleDomainDto.Role)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return err
	}

	if result {
		return context.JSON(http.StatusOK, roleDomainDto)
	}
	return context.JSON(http.StatusInternalServerError, errors.Error{Message: "Não foi possível adicionar o papel ao usuário"})
}

func (userController *userController) CheckIfUserHasPermission(context echo.Context) error {
	responseWriter := context.Response().Writer
	user := context.Request().Header.Get("user")
	domain := context.QueryParam("domain")
	resource := context.QueryParam("resource")
	action := context.QueryParam("action")

	if user == "" || domain == "" || resource == "" || action == "" {
		return context.JSON(http.StatusInternalServerError,
			errors.Error{Message: "É necessário informar o usuário e o domínio"})
	}

	result, err := userController.userService.CheckIfUserHasPermission(user, domain, resource, action)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return err
	}

	if result {
		return context.JSON(http.StatusOK, true)
	}
	return context.JSON(http.StatusUnauthorized, errors.Error{Message: "Usuário não tem permissão"})
}

func NewUserController(userService services.IUserService) IUserController {
	return &userController{
		userService: userService,
	}
}
