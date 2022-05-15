package main

import (
	"github.com/JustKato/FreePad/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the router
	router := gin.Default()

	// Read HTML Templates
	router.LoadHTMLGlob("templates/*")

	// Add Routes
	// Bind health checks
	routes.HealthRoutes(router)
	// Bind /api
	routes.ApiRoutes(router.Group("/api"))

	router.Run(":8080")

}
