package routes

import (
	"toDoListApp/taskService/config"
	"toDoListApp/taskService/controllers"
	"toDoListApp/taskService/services"
	"toDoListApp/taskService/utils"

	"github.com/gin-gonic/gin"
)

func ServiceRoutes(r *gin.Engine) {
	taskSrv := services.NewTaskService(config.DB)
	taskCtl := controllers.NewTaskController(taskSrv)

	// Todas requieren JWT (para conocer el user_id)
	protected := r.Group("/")
	protected.Use(utils.AuthMiddleware()) // lee user_id del token

	// 1. Crear
	protected.POST("/tareas", taskCtl.CreateTask)
	// 2. Actualizar
	protected.PUT("/tareas/:id", taskCtl.UpdateTask)
	// 3. Eliminar
	protected.DELETE("/tareas/:id", taskCtl.DeleteTask)
	// 4. Lista por usuario
	protected.GET("/tareas/usuario", taskCtl.ListMyTasks)
	// 5. Obtener por ID
	protected.GET("/tareas/:id", taskCtl.GetByID)
}
