package services

import (
	"github.com/Reyshal/task-manager-api/config"
	"github.com/Reyshal/task-manager-api/models"
	"gorm.io/gorm"
)

type TaskService struct {
	db *gorm.DB
}

func NewTaskService() *TaskService {
	return &TaskService{
		db: config.DB,
	}
}

func (s *TaskService) GetTasks(userID uint, filter string) ([]models.Task, error) {
	var tasks []models.Task

	query := s.db.Where("user_id = ?", userID)

	// Filter tasks
	if filter != "all" {
		query = query.Where("completed = ?", filter == "completed")
	}

	if err := query.Order("created_at").Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) GetByID(id string) (*models.Task, error) {
	var task models.Task
	if err := s.db.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) Create(task *models.Task) error {
	return s.db.Create(task).Error
}

func (s *TaskService) Update(task *models.Task) error {
	return s.db.Save(task).Error
}

func (s *TaskService) Delete(task *models.Task) error {
	return s.db.Delete(task).Error
}
