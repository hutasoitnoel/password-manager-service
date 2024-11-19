package controllers

import (
	"fmt"
	"net/http"
	"password-manager-service/helpers"
	config "password-manager-service/initializers"
	"password-manager-service/models"

	"github.com/gin-gonic/gin"
)

func FindAllIdentifications(c *gin.Context) {
	var identificationCards []models.IdentificationCard
	if err := config.DB.Find(&identificationCards).Error; err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"Message": fmt.Sprintf("Error finding all identificationCards: %v", err)},
		)
	}

	c.IndentedJSON(
		http.StatusOK,
		identificationCards,
	)
}

func FindIdentificationsByUserId(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "Please log in"},
		)
		return
	}

	var identificationCards []models.IdentificationCard
	// Fetch user's identificationCards
	if err := config.DB.Where("user_id = ?", user.(models.User).ID).Find(&identificationCards).Error; err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error fetching identificationCards: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		identificationCards,
	)
}

func CreateIdentification(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "Please log in"},
		)
		return
	}

	var identificationCard models.IdentificationCard
	if err := c.ShouldBindJSON(&identificationCard); err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error binding JSON: %v", err)},
		)
		return
	}

	identificationCard.UserId = user.(models.User).ID

	if err := helpers.Validator.Struct(identificationCard); err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error missing payload: %v", err)},
		)
		return
	}

	// Save to DB
	if err := config.DB.Create(&identificationCard).Error; err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error creating identificationCard: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		identificationCard,
	)
}

func UpdateIdentification(c *gin.Context) {
	identificationCardId := c.Param("identification_id")

	var payload models.IdentificationCard
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error binding JSON: %v", err)},
		)
		return
	}

	var identificationCard models.IdentificationCard
	data := config.DB.Find(&identificationCard, identificationCardId)
	if data.Error != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error finding identificationCard: %v", data.Error)},
		)
		return
	}

	if data.RowsAffected == 0 {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "Record not found"},
		)
		return
	}

	if err := config.DB.Model(&identificationCard).Updates(&payload).Error; err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error updating record: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		identificationCard,
	)
}

func DeleteIdentification(c *gin.Context) {
	identificationCardId := c.Param("identification_id")

	var identificationCard models.IdentificationCard
	data := config.DB.Find(&identificationCard, identificationCardId)
	if data.Error != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error finding saving: %v", data.Error)},
		)
		return
	}

	if data.RowsAffected == 0 {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": "Record not found"},
		)
		return
	}

	if err := config.DB.Delete(&models.IdentificationCard{}, identificationCardId).Error; err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error deleting identificationCard ID: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		gin.H{"message": fmt.Sprintf("Successfully deleted identificationCard ID: %v", identificationCardId)},
	)
}
