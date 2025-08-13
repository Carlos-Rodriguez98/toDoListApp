package main

import (
	"log"
	"toDoListApp/authService/config"
	"toDoListApp/authService/routes"

	//"toDoListApp/authService/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//Inicializar configuraciones
	//Variables de entorno
	config.LoadEnv()

	//Conectar a la base de datos usando AppConfig
	config.ConnectDatabase()

	//Inicializar Gin
	r := gin.Default()

	//Registrar rutas
	routes.ServiceRoutes(r)

	//Iniciar servidor
	log.Println("Servidore corriendo en http://localhost:8080")
	if error := r.Run(":8080"); error != nil {
		log.Fatal(error)
	}

}
