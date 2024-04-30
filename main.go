package main

import (
	"log"
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

func main() {
	router := gin.Default()

	// Test endpoints
	router.GET("/ping", controllers.Ping)
	router.GET("/validate", middlewares.RequireAuthorization, controllers.ValidateToken)

	// Super user endpoints
	router.GET("/users", controllers.FindAllUsers)
	router.GET("/users/:id", controllers.FindUserById)
	router.GET("/all-passwords", controllers.FindAllCredentials)

	// Client endpoints
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// Authorized endpoints
	router.GET("/passwords", middlewares.RequireAuthorization, controllers.FindCredentialsByUserId)
	router.POST("/passwords", middlewares.RequireAuthorization, controllers.CreateCredential)
	router.PATCH("/passwords/:credential_id", middlewares.RequireAuthorization, controllers.UpdateCredential)
	router.DELETE("/passwords/:credential_id", middlewares.RequireAuthorization, controllers.DeleteCredential)

	log.Fatal((http.ListenAndServe(":8080", middlewares.HandleCors(router))))
}
