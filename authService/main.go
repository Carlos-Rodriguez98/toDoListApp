package main

import (
	"log"
	"time"
	"toDoListApp/authService/config"
	"toDoListApp/authService/routes"

	//"toDoListApp/authService/routes"

	"github.com/gin-contrib/cors"
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

	// Configurar CORS para permitir cualquier origen
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//Servir carpeta "static" para datos estáticos
	r.Static("/static", "./static")

	//Registrar rutas
	routes.ServiceRoutes(r)

	//Iniciar servidor
	log.Println("Servidor corriendo en http://localhost:8080")
	if error := r.Run(":8080"); error != nil {
		log.Fatal(error)
	}

}
