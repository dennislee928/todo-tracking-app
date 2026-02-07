// Package main Todo Tracking App Backend Server
// @title Todo Tracking API
// @version 1.0
// @description API for Todo Tracking App - tasks, projects, and labels
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/todo-tracking-app
// @contact.email support@todo-tracking-app.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/todo-tracking-app/web-be/docs"
	"github.com/todo-tracking-app/web-be/internal/config"
	"github.com/todo-tracking-app/web-be/internal/database"
	"github.com/todo-tracking-app/web-be/internal/middleware"
	"github.com/todo-tracking-app/web-be/api/rest"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("connect database: %v", err)
	}

	r := gin.Default()

	// CORS
	r.Use(middleware.CORS())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")
	{
		// Auth routes (no auth middleware)
		authGroup := v1.Group("/auth")
		rest.RegisterAuthRoutes(authGroup, db, cfg)

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.Auth(cfg))
		{
			rest.RegisterProjectRoutes(protected, db)
			rest.RegisterTaskRoutes(protected, db)
			rest.RegisterLabelRoutes(protected, db)
		}
	}

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("run server: %v", err)
	}
}
