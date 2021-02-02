package http

import (
	"fmt"

	"controle_acesso_core.com/src/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type IHandler interface {
	Start() error
}

type apiHandler struct {
	echo        *echo.Echo
	healthCheck controllers.IHealthCheckController
}

func (handler *apiHandler) Start() error {
	handler.echo.Use(middleware.Logger())
	handler.echo.Use(middleware.Recover())
	api := handler.echo.Group("/api")

	healthCheck := api.Group("/healthcheck")
	healthCheck.GET("", handler.healthCheck.GetHealthCheck)

	return handler.echo.Start(fmt.Sprintf("%s:%d", "0.0.0.0", 8000))
}

func NewHandler(healthCheckController controllers.IHealthCheckController) IHandler {
	e := echo.New()
	return &apiHandler{
		echo:        e,
		healthCheck: healthCheckController,
	}
}
