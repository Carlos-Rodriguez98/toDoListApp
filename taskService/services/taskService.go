package services

import (
	"errors"
	"toDoListApp/taskService/dto"
	"toDoListApp/taskService/models"

	"gorm.io/gorm"
)

type TaskService struct {
	DB *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{DB: db}
}

func (s *TaskService) Create(userID uint, req dto.TaskCreateRequest) (*models.Tarea, error) {
	task := models.Tarea{
		UserID:     userID,
		CategoryID: req.CategoryID,
		Text:       req.Text,
		DueDate:    req.DueDate,
		Status:     "Pendiente",
	}
	if err := s.DB.Create(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) Update(userID, id uint, req dto.TaskUpdateRequest) (*models.Tarea, error) {
	var task models.Tarea
	if err := s.DB.First(&task, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	if req.Text != nil {
		task.Text = *req.Text
	}
	if req.DueDate != nil {
		task.DueDate = req.DueDate
	}
	if req.Status != nil {
		task.Status = *req.Status
	}

	if err := s.DB.Save(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) Delete(userID, id uint) error {
	return s.DB.
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&models.Tarea{}).Error
}

func (s *TaskService) GetByID(userID, id uint) (*models.Tarea, error) {
	var task models.Tarea
	if err := s.DB.First(&task, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) ListByUser(userID uint) ([]models.Tarea, error) {
	var tasks []models.Tarea
	if err := s.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
