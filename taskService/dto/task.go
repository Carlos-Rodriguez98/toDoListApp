package dto

import "time"

type TaskCreateRequest struct {
	Text       string     `json:"text" binding:"required"`
	CategoryID *uint      `json:"category_id"`
	DueDate    *time.Time `json:"due_date"` // formato ISO-8601 recomendado
}

type TaskCreateResponse struct {
	ID         uint       `json:"id"`
	UserID     uint       `json:"user_id"`
	CategoryID *uint      `json:"category_id"`
	Text       string     `json:"text"`
	DueDate    *time.Time `json:"due_date"`
	CreatedAt  time.Time  `json:"created_at"`
	Status     string     `json:"status"`
}

type TaskUpdateRequest struct {
	Text    *string     `json:"text"`     // opcional
	DueDate *time.Time  `json:"due_date"` // opcional
	Status  *string     `json:"status"`   // opcional
}

type TaskResponse struct {
	ID         uint       `json:"id"`
	UserID     uint       `json:"user_id"`
	CategoryID *uint      `json:"category_id"`
	Text       string     `json:"text"`
	DueDate    *time.Time `json:"due_date"`
	CreatedAt  time.Time  `json:"created_at"`
	Status     string     `json:"status"`
}
