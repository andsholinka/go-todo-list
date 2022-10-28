package controllers

import (
	"fmt"
	"net/http"

	"github.com/andsholinka/go-todo-list/initializers"
	"github.com/andsholinka/go-todo-list/models"
	"github.com/gin-gonic/gin"
)

func TodoCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		ActivityGroupId uint   `json:"activity_group_id"`
		Title           string `json:"title"`
		IsActive        bool
		Priority        string
	}

	c.Bind(&body)

	// Validate request body
	if !body.IsActive {
		body.IsActive = true
	}

	if body.Priority == "" {
		body.Priority = "very-high"
	}

	if body.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "title cannot be null",
		})
		return
	}

	if body.ActivityGroupId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "activity_group_id cannot be null",
		})
		return
	}

	// Create a post
	todo := models.Todo{ActivityGroupId: body.ActivityGroupId, Title: body.Title, IsActive: body.IsActive, Priority: body.Priority}
	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})
}

func TodoIndex(c *gin.Context) {

	// Get query param
	qr := c.Request.URL.Query()
	newId := qr.Get("activity_group_id")

	if newId == "" {
		var todos []models.Todo
		initializers.DB.Limit(3).Find(&todos)

		fmt.Println("tanpa query param id")

		// Respond with them
		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Success",
			"data":    todos,
		})
		return
	}

	fmt.Println("ada query param id")

	var todos []models.Todo
	initializers.DB.Limit(3).Find(&todos, "activity_group_id = ?", newId)

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    todos,
	})
}

func TodoShow(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Get the todo
	var todo models.Todo
	initializers.DB.First(&todo, id)

	// Validate
	if todo.ID <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Todo with ID " + id + " Not Found",
		})
		return
	}

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})
}

func TodoUpdate(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Find the data were updating
	var todo models.Todo
	initializers.DB.First(&todo, id)

	// Validate
	if todo.ID <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Todo with ID " + id + " Not Found",
		})
		return
	}

	// Get data off req body
	var body struct {
		ActivityGroupId uint   `json:"activity_group_id"`
		Title           string `json:"title"`
		IsActive        bool   `json:"is_active" form:"is_active"`
		Priority        string `gorm:"default:very-high" json:"priority" form:"priority"`
	}

	c.Bind(&body)

	// Update it
	initializers.DB.Model(&todo).Updates(models.Todo{
		ActivityGroupId: body.ActivityGroupId,
		Title:           body.Title,
		IsActive:        body.IsActive,
		Priority:        body.Priority,
	})

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})
}

func TodoDelete(c *gin.Context) {
	var data struct {
	}

	// Get id from url
	id := c.Param("id")

	// Find the data were updating
	var todo models.Todo
	initializers.DB.First(&todo, id)

	// Validate
	if todo.ID <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Todo with ID " + id + " Not Found",
			"data":    data,
		})
		return
	}

	// Delete the data
	initializers.DB.Delete(&models.Todo{}, id)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}
