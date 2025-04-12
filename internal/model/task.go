package model

import "time"

type Task struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`                    // Владелец задачи (кто создал)
	BoardID        int       `json:"board_id"`                   // Доска, к которой принадлежит задача
	AssignedUserID *int      `json:"assigned_user_id,omitempty"` // Исполнитель (может быть NULL)
	Title          string    `json:"title"`
	Description    *string   `json:"description,omitempty"` // Описание (может быть NULL)
	Status         string    `json:"status"`                // "todo", "in_progress", "done"
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
