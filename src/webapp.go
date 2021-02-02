package main

import (
	"controle_acesso_core.com/src/controllers"
	http_adapter "controle_acesso_core.com/src/http"
	"controle_acesso_core.com/src/repositorys"
	"controle_acesso_core.com/src/services"
)

func main() {
	//Instanciar os repositorys
	casbinRepository := repositorys.NewCasbinPostgressRepository()

	//Instanciar os serviços
	userService, err := services.NewUserService(casbinRepository)
	if err != nil {
		panic(err)
	}
	policyService, err := services.NewPolicyService(casbinRepository)
	if err != nil {
		panic(err)
	}

	//Instanciar as controller
	userController := controllers.NewUserController(userService)
	policyController := controllers.NewPolicyController(policyService)
	healthCheckController := controllers.NewHealtCheckController()

	handler := http_adapter.NewHandler(healthCheckController, userController, policyController)

	err = handler.Start()
	if err != nil {
		panic(err)
	}
}
