// main.go
package main

import (
	"log"
	"os"

	"github.com/Tekalig/job-finder-go/config"
	"github.com/Tekalig/job-finder-go/hasura"
	"github.com/Tekalig/job-finder-go/middleware"
	"github.com/Tekalig/job-finder-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Initialize Hasura client
	hasuraClient := hasura.NewClient(cfg.HasuraEndpoint, cfg.HasuraAdminSecret)

	// Initialize router
	r := gin.Default()

	// Add middleware
	r.Use(middleware.CORS())

	// Setup routes
	routes.SetupRoutes(r, cfg, hasuraClient)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
