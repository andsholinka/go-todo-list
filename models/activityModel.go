package models

import "gorm.io/gorm"

type Activity struct {
	// ID uint `gorm:"primaryKey" json:"id"`
	gorm.Model
	Title string `json:"title" form:"title" binding:"required" validate:"required"`
	Email string `json:"email" form:"email" binding:"required" validate:"required"`
}
