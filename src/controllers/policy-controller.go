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
	body := context.Request().Body
	err := json.NewDecoder(body).Decode(&policyDto)

	if err != nil || policyDto.Role == "" || policyDto.Action == "" || policyDto.Domain == "" || policyDto.Resource == "" {
		return context.JSON(http.StatusBadRequest, errors.Error{Message: "Invalid object"})
	}

	result, err := p.policyService.AddPolicy(policyDto)
	if err != nil || result == false {
		return context.JSON(http.StatusInternalServerError, errors.Error{Message: "This policy could not be created"})
	}

	return context.JSON(http.StatusCreated, policyDto)
}

func NewPolicyController(service services.IPolicyService) IPolicyController {
	return &policyController{
		policyService: service,
	}
}
