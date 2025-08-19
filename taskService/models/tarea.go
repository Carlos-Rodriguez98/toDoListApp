package models

import "time"

type Tarea struct {
	ID         uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     uint       `json:"user_id" gorm:"not null;index"`
	CategoryID *uint      `json:"category_id" gorm:"index"`
	Text       string     `json:"text" gorm:"type:text;not null"`
	DueDate    *time.Time `json:"due_date"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	Status     string     `json:"status" gorm:"size:50;default:'Pendiente'"`
}

func (Tarea) TableName() string { return "tareas" }
