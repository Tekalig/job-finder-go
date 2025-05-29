package routes

import (
	"github.com/Tekalig/job-finder-go/config"
	"github.com/Tekalig/job-finder-go/handlers"
	"github.com/Tekalig/job-finder-go/hasura"
	"github.com/Tekalig/job-finder-go/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config, hasuraClient *hasura.Client) {
	api := r.Group("/api")

	// Public routes
	auth := api.Group("/auth")
	{
		auth.POST("/employer/signup", func(c *gin.Context) {
			handlers.EmployerSignup(c, hasuraClient)
		})
		auth.POST("/expert/signup", func(c *gin.Context) {
			handlers.ExpertSignup(c, hasuraClient)
		})
		auth.POST("/login", func(c *gin.Context) {
			handlers.Login(c, hasuraClient, cfg.JWTSecret)
		})
	}

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.Auth(cfg.JWTSecret))
	{
		protected.POST("/jobs", func(c *gin.Context) {
			handlers.CreateJobHandler(c, hasuraClient)
		})
		protected.GET("/jobs", func(c *gin.Context) {
			handlers.GetJobsHandler(c, hasuraClient)
		})
		protected.PUT("/jobs/:id", func(c *gin.Context) {
			handlers.UpdateJobHandler(c, hasuraClient)
		})
		protected.DELETE("/jobs/:id", func(c *gin.Context) {
			handlers.DeleteJobHandler(c, hasuraClient)
		})
	}
}
