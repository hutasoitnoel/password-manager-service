package main

import (
	"net/http"
	"password-manager-service/controllers"
	config "password-manager-service/initializers"
	"password-manager-service/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func validateToken(c *gin.Context) {
	user, _ := c.Get("user")

	c.IndentedJSON(http.StatusOK, gin.H{"message": "I am logged in!", "user": user})
}

func main() {
	router := gin.Default()

	router.GET("/ping", controllers.Ping)
	router.GET("/users", controllers.FindAllUsers)
	router.GET("/users/:id", controllers.FindUserById)
	router.GET("/all-passwords", controllers.FindAllCredentials)

	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	router.POST("/passwords", controllers.CreateCredential)

	// Authorized endpoints
	router.GET("/passwords", middlewares.RequireAuthorization, controllers.FindCredentialsByUserId)
	router.GET("/validate", middlewares.RequireAuthorization, validateToken)

	router.Run(":8080")
}
