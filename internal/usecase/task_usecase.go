package usecase

import (
	"task-manager/internal/model"
)

type TaskRepository interface {
	GetAllTasks() ([]model.Task, error)
	GetTaskByID(id uint) (*model.Task, error)
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task) error
	DeleteTask(id uint) error
	BlockTask(id uint, adminID uint) error
}

type taskUsecase struct {
	repo TaskRepository
}

func NewTaskUsecase(repo TaskRepository) *taskUsecase {
	return &taskUsecase{repo}
}

func (u *taskUsecase) GetAllTasks() ([]model.Task, error) {
	return u.repo.GetAllTasks()
}

func (u *taskUsecase) GetTaskByID(id uint) (*model.Task, error) {
	return u.repo.GetTaskByID(id)
}

func (u *taskUsecase) CreateTask(task *model.Task) error {
	return u.repo.CreateTask(task)
}

func (u *taskUsecase) UpdateTask(task *model.Task) error {
	return u.repo.UpdateTask(task)
}

func (u *taskUsecase) DeleteTask(id uint) error {
	return u.repo.DeleteTask(id)
}

func (u *taskUsecase) BlockTask(id uint, adminID uint) error {
	return u.repo.BlockTask(id, adminID)
}
