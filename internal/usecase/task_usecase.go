package usecase

import (
	"task-manager/internal/models"
)

type TaskRepository interface {
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id int) (*models.Task, error)
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task) error
	DeleteTask(id int) error
}

type TaskUsecase interface {
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id int) (*models.Task, error)
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task) error
	DeleteTask(id int) error
}

type taskUsecase struct {
	repo TaskRepository
}

func NewTaskUsecase(repo TaskRepository) TaskUsecase {
	return &taskUsecase{repo}
}

func (u *taskUsecase) GetAllTasks() ([]models.Task, error) {
	return u.repo.GetAllTasks()
}

func (u *taskUsecase) GetTaskByID(id int) (*models.Task, error) {
	return u.repo.GetTaskByID(id)
}

func (u *taskUsecase) CreateTask(task *models.Task) error {
	return u.repo.CreateTask(task)
}

func (u *taskUsecase) UpdateTask(task *models.Task) error {
	return u.repo.UpdateTask(task)
}

func (u *taskUsecase) DeleteTask(id int) error {
	return u.repo.DeleteTask(id)
}
