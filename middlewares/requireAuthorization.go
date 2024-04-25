package middlewares

import (
	"fmt"
	"net/http"
	"os"
	config "password-manager-service/initializers"
	"password-manager-service/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuthorization(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": fmt.Sprintf("No cookie found: %v", err)},
		)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error parsing cookie: %v", err)},
		)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"message": fmt.Sprintf("cookie expired, please login: %v", err)},
			)
			return
		}

		var user models.User
		if err := config.DB.First(&user, claims["sub"]).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}

		if user.ID == 0 {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{"message": fmt.Sprintf("No user found: %v", err)},
			)
			return
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Error fetching claims: %v", err)},
		)
		return
	}
}
