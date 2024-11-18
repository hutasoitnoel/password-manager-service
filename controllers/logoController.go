package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type LogoResponse struct {
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	Image  string `json:"image"`
}

func FindLogoByName(c *gin.Context) {
	websiteName := c.Query("name")

	if websiteName == "" {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "name query parameter is required"},
		)
		return
	}

	if len(strings.Fields(websiteName)) > 1 {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "only single words are allowed"},
		)
		return
	}

	apikey := os.Getenv("API_NINJA_API_KEY")
	url := fmt.Sprintf("https://api.api-ninjas.com/v1/logo?name=%v", websiteName)

	// Create URL
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("Error finding all credentials: %v", err)},
		)
		return
	}
	request.Header.Set("x-api-key", apikey)

	// Hit endpoint URL
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("Error finding logo: %v", err)},
		)
		return
	}

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "Error reading response body"},
		)
		return
	}

	// JSON parse the body
	var logos []LogoResponse
	if err := json.Unmarshal(body, &logos); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "Error parsing body"},
		)
		return
	}

	if len(logos) == 0 {
		c.IndentedJSON(
			http.StatusOK,
			gin.H{"message": "Website name not found"},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		logos[0],
	)
}
