package controllers

import (
	"net/http"
	"toDoListApp/authService/dto"
	"toDoListApp/authService/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request dto.LoginRequest

	//Validar que se reciban los datos en el body
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Debe enviar su nombre de usuario y contraseña",
		})
		return
	}

	// Llamado al servicio para validar credenciales y generar token
	token, user, err := services.UserAuthenticator(request.UserName, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Respuesta exitosa con Token JWT
	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"usuario": user,
	})
}
