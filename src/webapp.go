package main

import (
	"boiler_plate.com/src/controllers"
	http_adapter "boiler_plate.com/src/http"
	"boiler_plate.com/src/services"
)

func main() {

	//Instanciar os repositorys

	//Instanciar os serviços
	loginService := services.NewLoginService()

	//Instanciar controllers
	healthCheckController := controllers.NewHealtCheckController()
	loginController := controllers.NewLoginController(loginService)

	handler := http_adapter.NewHandler(healthCheckController, loginController)

	err := handler.Start()
	if err != nil {
		panic(err)
	}
}
