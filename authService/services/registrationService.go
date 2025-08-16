package services

import (
	"toDoListApp/authService/dto"
	"toDoListApp/authService/models"
	"toDoListApp/authService/utils"

	"gorm.io/gorm"
)

type RegistrationService struct {
	DB *gorm.DB
}

func NewRegistrationService(db *gorm.DB) *RegistrationService {
	return &RegistrationService{DB: db}
}

func (s *RegistrationService) RegisterUser(input dto.UserRegisterRequest) (*models.Usuario, error) {
	// Definir ruta de imagen por defecto
	avatar := "/static/avatar.jpg"

	// Hashear la contraseña
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	// Crear modelo de usuario
	usuario := models.Usuario{
		UserName: input.UserName,
		Password: hashedPassword,
		Image:    avatar,
	}

	// Guardar usuario en la base de datos
	if err := s.DB.Create(&usuario).Error; err != nil {
		return nil, err
	}

	return &usuario, nil
}
