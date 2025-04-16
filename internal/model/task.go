package model

import "time"

type Task struct {
	ID             uint   `gorm:"primaryKey"`
	UserID         uint   `gorm:"not null"`
	BoardID        uint   `gorm:"not null"`
	BoardAdminID   uint   `gorm:"not null"`
	AssignedUserID uint   `gorm:"not null"`
	Title          string `gorm:"not null"`
	Description    string
	Status         string `gorm:"default:todo"`
	IsBlocked      bool   `gorm:"default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
