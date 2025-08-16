package models

import (
	"time"

	"gorm.io/gorm"
)

type Usuario struct {
	ID         uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Username   string         `json:"username" gorm:"unique;not null"`
	Password   string         `json:"-" gorm:"not null"` //This information isn't returned.
	AvatarPath string         `json:"avatarPath"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdateAt   time.Time      `json:"updateAt"`
	DeleteAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
