package main

import (
	"net/http"

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
	initializers.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to API TODO",
		})
	})

	r.POST("/activity-groups", controllers.ActivityCreate)
	r.GET("/activity-groups", controllers.ActivityIndex)
	r.GET("/activity-groups/:id", controllers.ActivityShow)
	r.PATCH("/activity-groups/:id", controllers.ActivityUpdate)
	r.DELETE("/activity-groups/:id", controllers.ActivityDelete)

	r.POST("/todo-items", controllers.TodoCreate)
	r.GET("/todo-items", controllers.TodoIndex)
	r.GET("/todo-items/:id", controllers.TodoShow)
	r.PATCH("/todo-items/:id", controllers.TodoUpdate)
	r.DELETE("/todo-items/:id", controllers.TodoDelete)

	r.Run()
}
