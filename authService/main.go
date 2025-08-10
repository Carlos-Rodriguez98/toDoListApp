package authService

import (
	"log"
	"toDoListApp/authService/config"
	"toDoListApp/authService/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//Inicializar configuraciones (DB, Variables de entorno, etc)
	config.InitConfig()

	//Inicializar Gin
	r := gin.Default()

	//Registrar rutas
	routes.RegisterRoutes(r)

	//Iniciar servidor
	log.Println("Servidore corriendo en http://localhost:8080")
	if error := r.Run(":8080"); error != nil {
		log.Fatal(error)
	}

}
