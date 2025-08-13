package controllers

import (
	"net/http"
	"toDoListApp/authService/dto"
	"toDoListApp/authService/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	RegistrationService *services.RegistrationService
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		RegistrationService: services.NewRegistrationService(db),
	}
}

func (ctrl *UserController) RegisterUser(c *gin.Context) {
	var request dto.UserRegisterRequest

	// Bind datos del formulario (texto)
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtener archivo de imagen del formulario
	imageFileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al recibir imagen: " + err.Error()})
		return
	}

	// Llamar al servicio
	user, err := ctrl.RegistrationService.RegisterUser(request, imageFileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado con éxito",
		"user":    user,
	})
}
