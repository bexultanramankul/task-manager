package repository

import (
	"errors"
	"fmt"
	"task-manager/internal/model"
	"task-manager/pkg/logger"

	"gorm.io/gorm"
)

type boardRepoGorm struct {
	db *gorm.DB
}

func NewBoardRepository(db *gorm.DB) *boardRepoGorm {
	return &boardRepoGorm{db: db}
}

func (r *boardRepoGorm) GetAllBoards() ([]model.Board, error) {
	var boards []model.Board
	if err := r.db.Find(&boards).Error; err != nil {
		return nil, err
	}
	return boards, nil
}

func (r *boardRepoGorm) GetBoardByID(id uint) (*model.Board, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid board ID: %d", id)
	}

	var board model.Board
	if err := r.db.First(&board, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("board not found")
		}
		return nil, err
	}
	return &board, nil
}

func (r *boardRepoGorm) CreateBoard(board *model.Board) error {
	if err := r.db.Create(board).Error; err != nil {
		return fmt.Errorf("failed to create board: %w", err)
	}
	return nil
}

func (r *boardRepoGorm) UpdateBoard(board *model.Board) error {
	if board.ID == 0 {
		return fmt.Errorf("invalid board ID: %d", board.ID)
	}

	result := r.db.Model(&model.Board{}).
		Where("id = ?", board.ID).
		Updates(map[string]interface{}{
			"name":       board.Name,
			"is_private": board.IsPrivate,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("board not found")
	}

	logger.Log.Printf("Board %d updated successfully", board.ID)
	return nil
}

func (r *boardRepoGorm) DeleteBoard(id uint) error {
	if id == 0 {
		return fmt.Errorf("invalid board ID: %d", id)
	}

	result := r.db.Delete(&model.Board{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete board: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("board not found")
	}

	logger.Log.Printf("Board %d deleted successfully", id)
	return nil
}

func (r *boardRepoGorm) BlockBoard(id uint, adminID uint) error {
	if id == 0 {
		return fmt.Errorf("invalid board ID: %d", id)
	}

	result := r.db.Model(&model.Board{}).
		Where("id = ?", id).
		Update("is_blocked", true)

	if result.Error != nil {
		return fmt.Errorf("failed to block board %d: %w", id, result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("board not found or already blocked")
	}

	logger.Log.Printf("Board %d has been blocked by admin %d", id, adminID)
	return nil
}
