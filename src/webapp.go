package main

import (
	"boiler_plate.com/src/controllers"
	http_adapter "boiler_plate.com/src/http"
	"boiler_plate.com/src/repositories"
	"boiler_plate.com/src/services"
)

func main() {

	//Instanciar os repositories
	userRepository, err := repositories.NewUserRepository()
	if err != nil {
		panic(err)
	}
	redisRepository := repositories.NewRedisRepository()

	//Instanciar os serviços
	loginService := services.NewLoginService(userRepository, redisRepository)
	userService := services.NewUserService(userRepository)

	//Instanciar controllers
	healthCheckController := controllers.NewHealtCheckController()
	loginController := controllers.NewLoginController(loginService)
	userController := controllers.NewUserController(userService)

	handler := http_adapter.NewHandler(healthCheckController, loginController, userController)

	err = handler.Start()
	if err != nil {
		panic(err)
	}
}
