package controllers

import (
	"fmt"
	"net/http"
	"password-manager-service/helpers"
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
		return
	}

	var credentials []models.Credential
	// Fetch user's credentials
	if err := config.DB.Where("user_id = ?", user.(models.User).ID).Find(&credentials).Error; err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error fetching credentials: %v", err)},
		)
		return
	}

	// Decrypt password
	for i := range credentials {
		decryptedPassword, err := helpers.Decrypt(credentials[i].Password)
		if err != nil {
			c.IndentedJSON(
				http.StatusInternalServerError,
				gin.H{"message": fmt.Sprintf("Error decrypting credentials: %v", err)},
			)
			return
		}
		credentials[i].Password = decryptedPassword
	}

	// Return decrypted credentials
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
		return
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

	if err := helpers.Validator.Struct(credential); err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error missing payload: %v", err)},
		)
		return
	}

	// Encrypt password
	encryptedPassword, err := helpers.Encrypt(credential.Password)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error encrypting password: %v", err)},
		)
		return
	}
	credential.Password = encryptedPassword

	// Save to DB
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

func UpdateCredential(c *gin.Context) {
	credentialId := c.Param("credential_id")

	var payload models.Credential
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error binding JSON: %v", err)},
		)
		return
	}

	var credential models.Credential
	data := config.DB.Find(&credential, credentialId)
	if data.Error != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error finding credential: %v", data.Error)},
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

	// Encrypt password
	encryptedPassword, err := helpers.Encrypt(credential.Password)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error encrypting password: %v", err)},
		)
		return
	}
	credential.Password = encryptedPassword

	if err := config.DB.Model(&credential).Updates(&payload).Error; err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error updating record: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		credential,
	)
}

func DeleteCredential(c *gin.Context) {
	credentialId := c.Param("credential_id")

	var credential models.Credential
	data := config.DB.Find(&credential, credentialId)
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

	if err := config.DB.Delete(&models.Credential{}, credentialId).Error; err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error deleting credential ID: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		gin.H{"message": fmt.Sprintf("Successfully deleted credential ID: %v", credentialId)},
	)
}
