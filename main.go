package main

import (
	"os"

	"github.com/JustKato/FreePad/lib/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	_, isDevelopment := os.LookupEnv("DEV_MODE")
	if isDevelopment {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize the router
	router := gin.Default()

	// Read HTML Templates
	router.LoadHTMLGlob("templates/**/*.html")

	// Load in static path
	router.Static("/static", "static/")

	// Add Routes
	routes.HomeRoutes(router)

	router.Run(":8080")

}
