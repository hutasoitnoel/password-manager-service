package controllers

import (
	"fmt"
	"net/http"
	"password-manager-service/helpers"
	config "password-manager-service/initializers"
	"password-manager-service/models"

	"github.com/gin-gonic/gin"
)

func FindSavingsByUserId(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "Please log in"},
		)
	}

	var savings []models.Saving
	if err := config.DB.Where("user_id = ?", user.(models.User).ID).Find(&savings).Error; err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error fetching savings: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		savings,
	)
}

func CreateSaving(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "Please log in"},
		)
	}

	var saving models.Saving
	if err := c.ShouldBindJSON(&saving); err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error binding JSON: %v", err)},
		)
		return
	}

	saving.UserId = user.(models.User).ID

	if err := helpers.Validator.Struct(saving); err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error missing payload: %v", err)},
		)
		return
	}

	if err := config.DB.Create(&saving).Error; err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error creating saving: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		saving,
	)
}

func UpdateSaving(c *gin.Context) {
	savingId := c.Param("saving_id")

	var payload models.Saving
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error binding JSON: %v", err)},
		)
		return
	}

	var saving models.Saving
	if err := config.DB.Find(&saving, savingId).Error; err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error finding saving: %v", err)},
		)
		return
	}

	if err := config.DB.Model(&saving).Updates(&payload).Error; err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error updating record: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		saving,
	)
}

func DeleteSaving(c *gin.Context) {
	savingId := c.Param("saving_id")

	if err := config.DB.Delete(&models.Saving{}, savingId).Error; err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error deleting saving ID: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		gin.H{"message": fmt.Sprintf("Successfully deleted saving ID: %v", savingId)},
	)
}
