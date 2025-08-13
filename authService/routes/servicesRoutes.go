package routes

import (
	"toDoListApp/authService/config"
	"toDoListApp/authService/controllers"

	"github.com/gin-gonic/gin"
)

func ServiceRoutes(route *gin.Engine) {
	//Creación de instancia del controllador con la DB
	userController := controllers.NewUserController(config.DB)

	//Endpoint para registro de usuario
	route.POST("/usuarios", userController.RegisterUser)
}
