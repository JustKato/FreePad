package main

import (
	"fmt"
	"os"

	"github.com/JustKato/FreePad/lib/controllers"
	"github.com/JustKato/FreePad/lib/objects"
	"github.com/JustKato/FreePad/lib/routes"
	"github.com/JustKato/FreePad/lib/socketmanager"
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

	// Load in the views data from storage
	err := objects.LoadViewsCache()
	if err != nil {
		fmt.Println("Failed to load views from cache")
		fmt.Println(err)
	}

	// Initialize the router
	router := gin.Default()

	// Apply the FreePad Headers
	controllers.ApplyHeaders(router)

	// Read HTML Templates
	router.LoadHTMLGlob("templates/**/*.html")

	// Load in static path
	router.Static("/static", "static/")

	// Implement the rate limiter
	controllers.DoRateLimit(router)

	// Admin Routing
	routes.AdminRoutes(router.Group("/admin"))

	// Add Routes
	routes.HomeRoutes(router)

	// Bind the Web Sockets
	socketmanager.BindSocket(router.Group("/ws"))

	router.Run(":8080")

}
