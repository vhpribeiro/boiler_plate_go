package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckEndpoint(t *testing.T) {
	var result string
	url := "/healthcheck"
	mensagemEsperada := "Aplicação disponível"
	e := echo.New()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	healthCheckController := NewHealtCheckController()

	healthCheckController.GetHealthCheck(c)

	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, mensagemEsperada, result)
	assert.Equal(t, http.StatusOK, rec.Code)
}
