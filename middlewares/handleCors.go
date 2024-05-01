package middlewares

import (
	"net/http"
	"os"

	"github.com/rs/cors"
)

func HandleCors(router http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("DEVELOPMENT_UI_URL"), os.Getenv("PRODUCTION_UI_URL")},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	return c.Handler(router)
}
