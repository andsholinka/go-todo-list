package main

import (
	"github.com/andsholinka/go-todo-list/controllers"
	"github.com/andsholinka/go-todo-list/initializers"
	"github.com/andsholinka/go-todo-list/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.CoonetToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Activity{})
	r := gin.Default()

	r.GET("/", controllers.Home)
	r.POST("/activity-groups", controllers.ActivityCreate)
	r.GET("/activity-groups", controllers.ActivityIndex)
	r.GET("/activity-groups/:id", controllers.ActivityShow)
	r.PATCH("/activity-groups/:id", controllers.ActivityUpdate)
	r.DELETE("/activity-groups/:id", controllers.ActivityDelete)

	r.Run()
}
