package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"task-manager/internal/models"
)

type TaskRepoImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepoImpl {
	return &TaskRepoImpl{db}
}

func (r *TaskRepoImpl) GetAllTasks() ([]models.Task, error) {
	const query = `
        SELECT id, user_id, board_id, assigned_user_id, title, description, 
               status, created_at, updated_at 
        FROM tasks
    `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]models.Task, 0)
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.UserID, &task.BoardID, &task.AssignedUserID,
			&task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepoImpl) GetTaskByID(id int) (*models.Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid task ID: %d", id)
	}

	const query = `
		SELECT id, user_id, board_id, assigned_user_id, title, description, 
		       status, created_at, updated_at 
		FROM tasks 
		WHERE id = $1
	`

	var task models.Task
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.UserID, &task.BoardID, &task.AssignedUserID,
		&task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("task not found")
		}
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return &task, nil
}

func (r *TaskRepoImpl) CreateTask(task *models.Task) error {
	const query = `
		INSERT INTO tasks (user_id, board_id, assigned_user_id, title, description, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(query, task.UserID, task.BoardID, task.AssignedUserID,
		task.Title, task.Description, task.Status).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}

func (r *TaskRepoImpl) UpdateTask(task *models.Task) error {
	if task.ID <= 0 {
		return fmt.Errorf("invalid task ID: %d", task.ID)
	}

	const query = `
		UPDATE tasks 
		SET user_id = $1, board_id = $2, assigned_user_id = $3, 
		    title = $4, description = $5, status = $6, updated_at = NOW()
		WHERE id = $7
	`

	result, err := r.db.Exec(query, task.UserID, task.BoardID, task.AssignedUserID,
		task.Title, task.Description, task.Status, task.ID)
	if err != nil {
		return fmt.Errorf("failed to update task %d: %w", task.ID, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for task %d: %w", task.ID, err)
	}

	if rowsAffected == 0 {
		return errors.New("task not found")
	}

	log.Printf("Task %d updated successfully", task.ID)

	return nil
}

func (r *TaskRepoImpl) DeleteTask(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid task ID: %d", id)
	}

	const query = `
		DELETE FROM tasks 
		WHERE id = $1
	`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task %d: %w", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for task %d: %w", id, err)
	}

	if rowsAffected == 0 {
		return errors.New("task not found")
	}

	log.Printf("Task %d deleted, rows affected: %d", id, rowsAffected)

	return nil
}
