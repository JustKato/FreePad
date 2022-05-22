package main

import (
	"os"

	"github.com/JustKato/FreePad/lib/controllers"
	"github.com/JustKato/FreePad/lib/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables, ignore if any errors come up
	_ = godotenv.Load()

	dm, isDevelopment := os.LookupEnv("DEV_MODE")
	if !isDevelopment && dm == "0" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Run the TaskManager
	go controllers.TaskManager()

	// Initialize the router
	router := gin.Default()

	// Read HTML Templates
	router.LoadHTMLGlob("templates/**/*.html")

	// Load in static path
	router.Static("/static", "static/")

	// Implement the rate limiter
	controllers.DoRateLimit(router)

	// Add Routes
	routes.HomeRoutes(router)

	router.Run(":8080")

}
