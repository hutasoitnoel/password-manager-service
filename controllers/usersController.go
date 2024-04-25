package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	config "password-manager-service/initializers"
	"password-manager-service/models"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var payload models.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("Error binding JSON: %v", err)},
		)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("Error hashing password: %v", err)},
		)
		return
	}

	var result = models.User{Username: payload.Username, Password: string(hash)}
	if err := config.DB.Create(&result).Error; err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error creating credential: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		result,
	)
}

func LoginUser(c *gin.Context) {
	var payload models.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("Error binding JSON: %v", err)},
		)
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", payload.Username).First(&user).Error; err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("User not found: %v", err)},
		)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Incorrect password: %v", err)},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		user,
	)
}

func FindAllUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error fetching users: %v", err)},
		)
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func FindUserById(c *gin.Context) {
	idUint, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error parsing id: %v", err)},
		)
		return
	}

	var user models.User
	if err := config.DB.First(&user, idUint).Error; err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Failed to retrieve first user: %v", err)},
		)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
