package main

import (
	"log"
	"time"
	"toDoListApp/taskService/config"
	"toDoListApp/taskService/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDatabase()

	r := gin.Default()

	// Middleware CORS (permite cualquier origen)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// estáticos si quieres
	r.Static("/static", "./static")

	routes.ServiceRoutes(r)

	addr := ":" + config.AppConfig.ServerPort
	log.Println("TaskService corriendo en http://localhost" + addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
