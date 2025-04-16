package model

import "time"

type Board struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Name      string `gorm:"not null"`
	IsPrivate bool   `gorm:"default:false"`
	IsBlocked bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
