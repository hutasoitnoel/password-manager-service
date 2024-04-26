package controllers

import (
	"fmt"
	"net/http"
	config "password-manager-service/initializers"
	"password-manager-service/models"

	"github.com/gin-gonic/gin"
)

func FindAllCredentials(c *gin.Context) {
	var credentials []models.Credential
	if err := config.DB.Find(&credentials).Error; err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"Message": fmt.Sprintf("Error finding all credentials: %v", err)},
		)
	}

	c.IndentedJSON(
		http.StatusOK,
		credentials,
	)
}

func FindCredentialsByUserId(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "Please log in"},
		)
	}

	var credentials []models.Credential
	if err := config.DB.Where("user_id = ?", user.(models.User).ID).Find(&credentials).Error; err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error fetching credentials: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		credentials,
	)
}

func CreateCredential(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "Please log in"},
		)
	}

	var credential models.Credential
	if err := c.ShouldBindJSON(&credential); err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error binding JSON: %v", err)},
		)
		return
	}

	credential.UserId = user.(models.User).ID

	if err := config.DB.Create(&credential).Error; err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error creating credential: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		credential,
	)
}
