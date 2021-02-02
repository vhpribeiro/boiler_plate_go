package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IHealthCheckController interface {
	GetHealthCheck(context echo.Context) error
}

type healthCheckController struct{}

func (*healthCheckController) GetHealthCheck(context echo.Context) error {
	return context.JSON(http.StatusOK, "Aplicação disponível")
}

func NewHealtCheckController() IHealthCheckController {
	return &healthCheckController{}
}
