package controllers

import (
	"encoding/json"
	"net/http"

	"controle_acesso_core.com/src/services"
	"controle_acesso_core.com/src/services/dtos"
	"controle_acesso_core.com/src/utils/errors"
	"github.com/labstack/echo/v4"
)

type IPolicyController interface {
	AddPolicy(context echo.Context) error
}

type policyController struct {
	policyService services.IPolicyService
}

func (p *policyController) AddPolicy(context echo.Context) error {
	var policyDto dtos.PolicyDto
	responseWriter := context.Response().Writer
	body := context.Request().Body
	err := json.NewDecoder(body).Decode(&policyDto)

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return err
	}

	result, err := p.policyService.AddPolicy(policyDto)
	if err != nil {
		return err
	}

	if result {
		return context.JSON(http.StatusOK, policyDto)
	}

	return context.JSON(http.StatusInternalServerError, errors.Error{Message: "Não foi possível adicionar a política"})
}

func NewPolicyController(service services.IPolicyService) IPolicyController {
	return &policyController{
		policyService: service,
	}
}
