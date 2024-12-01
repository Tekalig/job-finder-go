// main.go
package main

import (
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/Tekalig/job-finder-go/config"
    "github.com/Tekalig/job-finder-go/routes"
    "github.com/Tekalig/job-finder-go/middleware"
)

func main() {
    // Load configuration
    cfg, err := config.Load()
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }

    // Initialize router
    r := gin.Default()

    // Add middleware
    r.Use(middleware.CORS())

    // Setup routes
    routes.SetupRoutes(r, cfg)

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}
