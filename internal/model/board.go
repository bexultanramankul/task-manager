package model

import "time"

type Board struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"` // Владелец доски
	Name      string    `json:"name"`
	IsPrivate bool      `json:"is_private"`
	CreatedAt time.Time `json:"created_at"`
}
