package services

import (
	"toDoListApp/authService/dto"
	"toDoListApp/authService/models"
	"toDoListApp/authService/utils"

	"mime/multipart"

	"gorm.io/gorm"
)

type RegistrationService struct {
	DB *gorm.DB
}

func NewRegistrationService(db *gorm.DB) *RegistrationService {
	return &RegistrationService{DB: db}
}

func (s *RegistrationService) RegisterUser(input dto.UserRegisterRequest, imageFileHeader *multipart.FileHeader) (*models.Usuario, error) {
	//Abrir archivo antes de pasarlo a SaveProfileImage
	file, err := imageFileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Guardar imagen en el almacenamiento
	imagePath, err := utils.SaveProfileImage(file, imageFileHeader)
	if err != nil {
		return nil, err
	}

	// Hashear la contraseña
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	// Crear modelo de usuario
	usuario := models.Usuario{
		UserName: input.UserName,
		Password: hashedPassword,
		Image:    imagePath,
	}

	// Guardar usuario en la base de datos
	if err := s.DB.Create(&usuario).Error; err != nil {
		return nil, err
	}

	return &usuario, nil
}
