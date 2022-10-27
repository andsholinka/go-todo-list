package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ActivityGroupId string `json:"activity_group_id" form:"activity_group_id"`
	Title           string `json:"title" form:"title"`
	IsActive        bool   `json:"is_active" form:"is_active"`
	Priority        string `json:"priority" form:"priority"`
}
