package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"task-manager/internal/models"
)

type BoardRepoImpl struct {
	db *sql.DB
}

func NewBoardRepository(db *sql.DB) *BoardRepoImpl {
	return &BoardRepoImpl{db}
}

func (r *BoardRepoImpl) GetAllBoards() ([]models.Board, error) {
	const query = "SELECT id, user_id, name, is_private, created_at FROM boards"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	boards := make([]models.Board, 0)
	for rows.Next() {
		var board models.Board
		if err := rows.Scan(&board.ID, &board.UserID, &board.Name, &board.IsPrivate, &board.CreatedAt); err != nil {
			return nil, err
		}
		boards = append(boards, board)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return boards, nil
}

func (r *BoardRepoImpl) GetBoardByID(id int) (*models.Board, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid board ID: %d", id)
	}

	const query = "SELECT id, user_id, name, is_private, created_at FROM boards WHERE id = $1"

	var board models.Board
	err := r.db.QueryRow(query, id).Scan(&board.ID, &board.UserID, &board.Name, &board.IsPrivate, &board.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("board not found")
		}
		return nil, fmt.Errorf("failed to get board: %w", err)
	}

	return &board, nil
}

func (r *BoardRepoImpl) CreateBoard(board *models.Board) error {
	const query = `
		INSERT INTO boards (user_id, name, is_private) 
		VALUES ($1, $2, COALESCE($3, FALSE)) 
		RETURNING id, created_at`

	err := r.db.QueryRow(query, board.UserID, board.Name, board.IsPrivate).Scan(&board.ID, &board.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create board: %w", err)
	}

	return nil
}

func (r *BoardRepoImpl) UpdateBoard(board *models.Board) error {
	if board.ID <= 0 {
		return fmt.Errorf("invalid board ID: %d", board.ID)
	}

	const query = `
		UPDATE boards 
		SET name = $1, is_private = $2, updated_at = NOW() 
		WHERE id = $3`

	result, err := r.db.Exec(query, board.Name, board.IsPrivate, board.ID)
	if err != nil {
		return fmt.Errorf("failed to update board %d: %w", board.ID, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for board %d: %w", board.ID, err)
	}

	if rowsAffected == 0 {
		return errors.New("board not found")
	}

	log.Printf("Board %d updated successfully", board.ID)

	return nil
}

func (r *BoardRepoImpl) DeleteBoard(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid board ID: %d", id)
	}

	const query = "DELETE FROM boards WHERE id = $1"

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete board %d: %w", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for board %d: %w", id, err)
	}

	if rowsAffected == 0 {
		return errors.New("board not found")
	}

	log.Printf("Board %d deleted successfully", id)

	return nil
}
