package controllers

import (
	"fmt"
	"net/http"
	config "password-manager-service/initializers"
	"password-manager-service/models"
	"strconv"

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
	userIdUint, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error parsing id: %v", err)},
		)
		return
	}

	var credentials []models.Credential
	if err := config.DB.Where("user_id = ?", uint(userIdUint)).Find(&credentials).Error; err != nil {
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
	var credential models.Credential
	if err := c.ShouldBindJSON(&credential); err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error binding JSON: %v", err)},
		)
		return
	}

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
