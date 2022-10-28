package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	// gorm.Model
	ID              uint           `gorm:"primaryKey" json:"id"`
	ActivityGroupId uint           `json:"activity_group_id" form:"activity_group_id"`
	Title           string         `json:"title" form:"title"`
	IsActive        bool           `gorm:"default:true" json:"is_active" form:"is_active"`
	Priority        string         `gorm:"default:very-high" json:"priority" form:"priority"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
