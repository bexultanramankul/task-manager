package usecase

import (
	"task-manager/internal/models"
)

type BoardRepository interface {
	GetAllBoards() ([]models.Board, error)
	GetBoardByID(id int) (*models.Board, error)
	CreateBoard(board *models.Board) error
	UpdateBoard(board *models.Board) error
	DeleteBoard(id int) error
}

type BoardUsecase interface {
	GetAllBoards() ([]models.Board, error)
	GetBoardByID(id int) (*models.Board, error)
	CreateBoard(board *models.Board) error
	UpdateBoard(board *models.Board) error
	DeleteBoard(id int) error
}

type boardUsecase struct {
	repo BoardRepository
}

func NewBoardUsecase(repo BoardRepository) BoardUsecase {
	return &boardUsecase{repo}
}

func (u *boardUsecase) GetAllBoards() ([]models.Board, error) {
	return u.repo.GetAllBoards()
}

func (u *boardUsecase) GetBoardByID(id int) (*models.Board, error) {
	return u.repo.GetBoardByID(id)
}

func (u *boardUsecase) CreateBoard(board *models.Board) error {
	return u.repo.CreateBoard(board)
}

func (u *boardUsecase) UpdateBoard(board *models.Board) error {
	return u.repo.UpdateBoard(board)
}

func (u *boardUsecase) DeleteBoard(id int) error {
	return u.repo.DeleteBoard(id)
}
