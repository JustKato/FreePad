package main

import (
	"os"

	"github.com/JustKato/FreePad/models/database"
	"github.com/JustKato/FreePad/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	_, isRelease := os.LookupEnv("RELEASE_MODE")
	if isRelease {
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
	// Bind /api
	routes.ApiRoutes(router.Group("/api"))

	// TODO: Sockets: https://gist.github.com/supanadit/f6de65fc5896e8bb0c4656e451387d0f

	// Try and run migrations
	database.MigrateMysql()

	router.Run(":8080")

}
