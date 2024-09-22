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
	// pancingggg

	router := gin.Default()

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
	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.Logout)
	router.GET("/website-logo", controllers.FindLogoByName)

	// Authorized endpoints
	router.GET("/passwords", middlewares.RequireAuthorization, controllers.FindCredentialsByUserId)
	router.POST("/passwords", middlewares.RequireAuthorization, controllers.CreateCredential)
	router.PATCH("/passwords/:credential_id", middlewares.RequireAuthorization, controllers.UpdateCredential)
	router.DELETE("/passwords/:credential_id", middlewares.RequireAuthorization, controllers.DeleteCredential)

	router.GET("/savings", middlewares.RequireAuthorization, controllers.FindSavingsByUserId)
	router.POST("/savings", middlewares.RequireAuthorization, controllers.CreateSaving)
	router.PATCH("/savings/:saving_id", middlewares.RequireAuthorization, controllers.UpdateSaving)
	router.DELETE("/savings/:saving_id", middlewares.RequireAuthorization, controllers.DeleteSaving)

	log.Fatal((http.ListenAndServe(":8080", middlewares.HandleCors(router))))
}
