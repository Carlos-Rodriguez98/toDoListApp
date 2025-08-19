package main

import (
	"log"
	"toDoListApp/taskService/config"
	"toDoListApp/taskService/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDatabase()

	r := gin.Default()

	// estáticos si quieres
	r.Static("/static", "./static")

	routes.ServiceRoutes(r)

	addr := ":" + config.AppConfig.ServerPort
	log.Println("TaskService corriendo en http://localhost" + addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
