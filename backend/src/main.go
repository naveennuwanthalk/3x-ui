package main

import (
	"log"
	"os"
	"x-ui/backend/src/controllers"
	"x-ui/backend/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set up Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Initialize services
	services.InitializeServices()

	// Set up routes
	controllers.SetupRoutes(r)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}