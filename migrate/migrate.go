package main

import (
	"github.com/andsholinka/go-todo-list/initializers"
	"github.com/andsholinka/go-todo-list/models"
)

func init() {
	initializers.CoonetToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Activity{})
}
