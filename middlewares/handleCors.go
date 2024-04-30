package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

func HandleCors(router http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "PUT"},
	})

	return c.Handler(router)
}
