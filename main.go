package main

import (
	"fmt"
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

	fmt.Println("poom")

	// Test endpoints
	router.GET("/ping", controllers.Ping)

	// Super user endpoints
	router.GET("/users", controllers.FindAllUsers)
	router.GET("/users/:id", controllers.FindUserById)
	router.GET("/all-passwords", controllers.FindAllCredentials)
	router.DELETE("/users/:user_id", controllers.DeleteUserById)

	// Client endpoints
	router.GET("/check-auth", controllers.CheckAuthentication)
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// Authorized endpoints
	router.GET("/passwords", middlewares.RequireAuthorization, controllers.FindCredentialsByUserId)
	router.POST("/passwords", middlewares.RequireAuthorization, controllers.CreateCredential)
	router.PATCH("/passwords/:credential_id", middlewares.RequireAuthorization, controllers.UpdateCredential)
	router.DELETE("/passwords/:credential_id", middlewares.RequireAuthorization, controllers.DeleteCredential)

	log.Fatal((http.ListenAndServe(":8080", middlewares.HandleCors(router))))
}
