package services

import (
	"errors"
	"toDoListApp/authService/config"
	"toDoListApp/authService/models"
	"toDoListApp/authService/utils"
)

func UserAuthenticator(userName, password string) (string, models.Usuario, error) {
	//Busqueda del usuario en la BD
	var user models.Usuario
	if err := config.DB.Where("username = ?", userName).First(&user).Error; err != nil {
		return "", models.Usuario{}, errors.New("usuario incorrecto")
	}

	// Comparación de contraseñas usando Hash
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", models.Usuario{}, errors.New("contraseña incorrecta")
	}

	//Generación de Token
	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		return "", models.Usuario{}, errors.New("error durante la generacion del token")
	}

	return token, user, nil
}
