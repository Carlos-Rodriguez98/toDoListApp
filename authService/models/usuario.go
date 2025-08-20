package models

import (
	"time"

	"gorm.io/gorm"
)

type Usuario struct {
	ID         uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Username   string         `json:"username" gorm:"size:50;unique;not null"`
	Password   string         `json:"-" gorm:"size:100;not null"` //This information isn't returned.
	AvatarPath string         `json:"avatarPath" gorm:"size:255;default:'/static/default.jpg'"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
