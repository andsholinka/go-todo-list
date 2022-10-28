package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	// gorm.Model
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `json:"title" form:"title"`
	Email     string         `json:"email" form:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
