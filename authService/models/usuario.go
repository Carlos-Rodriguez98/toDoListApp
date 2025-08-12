package models

import (
	"time"

	"gorm.io/gorm"
)

type Usuario struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserName  string         `json:"userName" gorm:"unique;not null"`
	Password  string         `json:"-" gorm:"not null"` //This information isn't returned.
	Image     string         `json:"image"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdateAt  time.Time      `json:"updateAt"`
	DeleteAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
