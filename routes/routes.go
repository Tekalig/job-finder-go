// routes/routes.go
package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/Tekalig/job-finder-go/handlers"
    "github.com/Tekalig/job-finder-go/middleware"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
    api := r.Group("/api")
    
    // Public routes
    auth := api.Group("/auth")
    {
        auth.POST("/employer/signup", handlers.EmployerSignup)
        auth.POST("/expert/signup", handlers.ExpertSignup)
        auth.POST("/login", handlers.Login)
    }
    
    // Protected routes
    protected := api.Group("/")
    protected.Use(middleware.Auth(cfg.JWTSecret))
    {
        protected.POST("/jobs", handlers.CreateJob)
        protected.GET("/jobs", handlers.GetJobs)
        protected.PUT("/jobs/:id", handlers.UpdateJob)
        protected.DELETE("/jobs/:id", handlers.DeleteJob)
    }
}
