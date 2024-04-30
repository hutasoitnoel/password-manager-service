package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "Pong!"},
	)
}

func ValidateToken(c *gin.Context) {
	user, _ := c.Get("user")

	c.IndentedJSON(http.StatusOK, gin.H{"message": "I am logged in!", "user": user})
}
