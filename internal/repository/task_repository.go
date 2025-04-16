package repository

import (
	"errors"
	"fmt"
	"task-manager/internal/model"
	"task-manager/pkg/logger"

	"gorm.io/gorm"
)

type taskRepoGorm struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepoGorm {
	return &taskRepoGorm{db}
}

func (r *taskRepoGorm) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepoGorm) GetTaskByID(id uint) (*model.Task, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid task ID: %d", id)
	}

	var task model.Task
	if err := r.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &task, nil
}

func (r *taskRepoGorm) CreateTask(task *model.Task) error {
	if err := r.db.Create(task).Error; err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	return nil
}

func (r *taskRepoGorm) UpdateTask(task *model.Task) error {
	if task.ID == 0 {
		return fmt.Errorf("invalid task ID: %d", task.ID)
	}

	result := r.db.Model(&model.Task{}).
		Where("id = ?", task.ID).
		Updates(map[string]interface{}{
			"board_id":         task.BoardID,
			"assigned_user_id": task.AssignedUserID,
			"title":            task.Title,
			"description":      task.Description,
			"status":           task.Status,
		})

	if result.Error != nil {
		return fmt.Errorf("failed to update task %d: %w", task.ID, result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}

	logger.Log.Printf("Task %d updated successfully", task.ID)
	return nil
}

func (r *taskRepoGorm) DeleteTask(id uint) error {
	if id == 0 {
		return fmt.Errorf("invalid task ID: %d", id)
	}

	result := r.db.Delete(&model.Task{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete task %d: %w", id, result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}

	logger.Log.Printf("Task %d deleted, rows affected: %d", id, result.RowsAffected)
	return nil
}

func (r *taskRepoGorm) BlockTask(id uint, adminID uint) error {
	if id == 0 {
		return fmt.Errorf("invalid task ID: %d", id)
	}

	result := r.db.Model(&model.Task{}).
		Where("id = ?", id).
		Update("is_blocked", true)

	if result.Error != nil {
		return fmt.Errorf("failed to block task %d: %w", id, result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("task not found or already blocked")
	}

	logger.Log.Printf("Task %d has been blocked by admin %d", id, adminID)
	return nil
}
