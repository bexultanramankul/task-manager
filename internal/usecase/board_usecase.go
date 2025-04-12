package usecase

import (
	"task-manager/internal/model"
)

type BoardRepository interface {
	GetAllBoards() ([]model.Board, error)
	GetBoardByID(id int) (*model.Board, error)
	CreateBoard(board *model.Board) error
	UpdateBoard(board *model.Board) error
	DeleteBoard(id int) error
}

type BoardUsecase interface {
	GetAllBoards() ([]model.Board, error)
	GetBoardByID(id int) (*model.Board, error)
	CreateBoard(board *model.Board) error
	UpdateBoard(board *model.Board) error
	DeleteBoard(id int) error
}

type boardUsecase struct {
	repo BoardRepository
}

func NewBoardUsecase(repo BoardRepository) BoardUsecase {
	return &boardUsecase{repo}
}

func (u *boardUsecase) GetAllBoards() ([]model.Board, error) {
	return u.repo.GetAllBoards()
}

func (u *boardUsecase) GetBoardByID(id int) (*model.Board, error) {
	return u.repo.GetBoardByID(id)
}

func (u *boardUsecase) CreateBoard(board *model.Board) error {
	return u.repo.CreateBoard(board)
}

func (u *boardUsecase) UpdateBoard(board *model.Board) error {
	return u.repo.UpdateBoard(board)
}

func (u *boardUsecase) DeleteBoard(id int) error {
	return u.repo.DeleteBoard(id)
}
