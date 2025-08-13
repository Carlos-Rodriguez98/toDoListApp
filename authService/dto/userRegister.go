package dto

import (
	"mime/multipart"
)

type UserRegisterRequest struct {
	UserName string                `form:"username" binding:"required"`
	Password string                `form:"password" binding:"required"`
	Image    *multipart.FileHeader `form:"image"`
}
