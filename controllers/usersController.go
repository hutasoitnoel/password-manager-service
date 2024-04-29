package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"password-manager-service/helpers"
	config "password-manager-service/initializers"
	"password-manager-service/models"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
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

	if err := helpers.Validator.Struct(payload); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("Error missing payload: %v", err)},
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

	if err := helpers.Validator.Struct(payload); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("Error missing payload: %v", err)},
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error creating token: %v", err)},
		)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"data":  user,
			"token": tokenString,
		},
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
