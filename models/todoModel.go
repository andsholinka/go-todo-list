package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ActivityGroupId string `json:"activity_group_id" form:"activity_group_id" binding:"required" validate:"required"`
	Title           string `json:"title" form:"title" binding:"required" validate:"required"`
	IsActive        bool   `json:"is_active" form:"is_active" binding:"required" validate:"required"`
	Priority        string `json:"priority" form:"priority" binding:"required" validate:"required"`
}
