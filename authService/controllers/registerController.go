package controllers

import (
	"fmt"            //Para formatear las cadenas
	"mime/multipart" //Para manejo de archivos subidos
	"net/http"       //Para interactuar con servidores web y clientes HTTP
	"time"           //Para obtener timestamps

	"toDoListApp/authService/models" //Ruta donde se encuentra el modelo del usuario.
	"toDoListApp/authService/utils"

	"github.com/gin-gonic/gin" //Framework Gin
)

func CrearUsuario(c *gin.Context) {

	//Estuctura temporal llamada userData donde se capturan los datos enviados
	var userData struct {
		userName string                `form:"userName" binding:"required"`
		password string                `form:"password" binding:"required"`
		image    *multipart.FileHeader `form:"imagen"`
	}

	//Toma los datos recibidos y los almacena en la estructura y si se genera un err lo retorna como respuesta
	if err := c.ShouldBind(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Almacenar imagen si es enviada
	var rutaImagen string
	if userData.image != nil {
		nombreArchivo := fmt.Sprintf("%d_%s", time.Now().Unix(), userData.image.Filename)
		ruta := fmt.Sprintf("uploads/profileImages/%s", nombreArchivo)

		if err := c.SaveUploadedFile(userData.image, ruta); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error guardando imagen"})
			return
		}
		rutaImagen = "/" + ruta
	}

	//Hash del password y validar que no se tiene un error
	hashPassword, err := utils.HashPassword(userData.password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando hash de contraseña"})
		return
	}

	usuario := models.Usuario{
		UserName: userData.userName,
		Password: hashPassword,
		Image:    rutaImagen,
	}

	//if err := database.DB.Create(&usuario).Error; err != nill {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando usuario"})
	//	return
	//}

	c.JSON(http.StatusCreated, gin.H{
		"id":       usuario.ID,
		"userName": usuario.UserName,
		"image":    usuario.Image,
	})
}
