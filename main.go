package main

import (
	"password-manager-service/controllers"
	config "password-manager-service/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	router := gin.Default()

	router.GET("/ping", controllers.Ping)
	router.GET("/users", controllers.FindAllUsers)
	router.GET("/users/:id", controllers.FindUserById)
	router.GET("/passwords", controllers.FindAllCredentials)
	router.GET("/passwords/:user_id", controllers.FindCredentialsByUserId)

	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	router.POST("/passwords", controllers.CreateCredential)

	router.Run(":8080")
}
