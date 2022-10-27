package controllers

import (
	"net/http"

	"github.com/andsholinka/go-todo-list/initializers"
	"github.com/andsholinka/go-todo-list/models"
	"github.com/gin-gonic/gin"
)

func ActivityCreate(c *gin.Context) {
	// Get data off re body
	var body struct {
		Title string
		Email string
	}

	c.Bind(&body)

	// Validate request body
	if body.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "title cannot be null",
		})
		return
	}

	// Create a post
	activity := models.Activity{Title: body.Title, Email: body.Email}
	result := initializers.DB.Create(&activity)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(http.StatusOK, gin.H{
		"status":  "Sukses",
		"message": "Sukses",
		"data":    activity,
	})
}

func ActivityIndex(c *gin.Context) {
	// Get the posts
	var activities []models.Activity
	initializers.DB.Limit(3).Find(&activities)

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"status":  "Sukses",
		"message": "Sukses",
		"data":    activities,
	})
}

func ActivityShow(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Get the activity
	var activity models.Activity
	initializers.DB.First(&activity, id)

	// Validate
	if activity.Title == "" {
		// fmt.Println("record not found")
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": `Activity with ID Not Found`,
		})
		return
	}

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"status":  "Sukses",
		"message": "Sukses",
		"data":    activity,
	})
}

func ActivityUpdate(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Get data off re body
	var body struct {
		Title string
		Email string
	}

	c.Bind(&body)

	// Find the data were updating
	var activity models.Activity
	initializers.DB.First(&activity, id)

	// Validate
	if activity.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": `Activity with ID Not Found`,
		})
		return
	}

	// Update it
	initializers.DB.Model(&activity).Updates(models.Activity{
		Title: body.Title,
		Email: body.Email,
	})

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"status":  "Sukses",
		"message": "Sukses",
		"data":    activity,
	})
}

func ActivityDelete(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Find the data were updating
	var activity models.Activity
	initializers.DB.First(&activity, id)

	// Validate
	if activity.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": `Activity with ID Not Found`,
		})
		return
	}

	// Delete the data
	initializers.DB.Delete(&models.Activity{}, id)

	// Respond
	c.Status(200)
}
