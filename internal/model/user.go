package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null;unique"` // Добавлено unique
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"column:password_hash"`
	Role      string `gorm:"default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
