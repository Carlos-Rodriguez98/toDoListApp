package controllers

import (
	"fmt"
	"net/http"
	"toDoListApp/taskService/dto"
	"toDoListApp/taskService/services"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Service *services.TaskService
}

func NewTaskController(s *services.TaskService) *TaskController {
	return &TaskController{Service: s}
}

// POST /tareas
func (ctl *TaskController) CreateTask(c *gin.Context) {
	var req dto.TaskCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Body inválido: " + err.Error()})
		return
	}
	uid := c.GetUint("user_id")
	task, err := ctl.Service.Create(uid, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la tarea"})
		return
	}
	c.JSON(http.StatusCreated, task)
}

// PUT /tareas/:id
func (ctl *TaskController) UpdateTask(c *gin.Context) {
	var req dto.TaskUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Body inválido: " + err.Error()})
		return
	}
	var id uint
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	uid := c.GetUint("user_id")
	task, err := ctl.Service.Update(uid, id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// DELETE /tareas/:id
func (ctl *TaskController) DeleteTask(c *gin.Context) {
	var id uint
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	uid := c.GetUint("user_id")
	if err := ctl.Service.Delete(uid, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tarea eliminada", "id": id})
}

// GET /tareas/usuario
func (ctl *TaskController) ListMyTasks(c *gin.Context) {
	uid := c.GetUint("user_id")
	tasks, err := ctl.Service.ListByUser(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las tareas"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tareas": tasks})
}

// GET /tareas/:id
func (ctl *TaskController) GetByID(c *gin.Context) {
	var id uint
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	uid := c.GetUint("user_id")
	task, err := ctl.Service.GetByID(uid, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
		return
	}
	c.JSON(http.StatusOK, task)
}
