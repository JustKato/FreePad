package main

import (
	"fmt"
	"os"

	"github.com/JustKato/FreePad/models/database"
	"github.com/JustKato/FreePad/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	_, isDevelopment := os.LookupEnv("IS_DEV")
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
	// Bind /api
	routes.ApiRoutes(router.Group("/api"))

	// TODO: Sockets: https://gist.github.com/supanadit/f6de65fc5896e8bb0c4656e451387d0f

	// Try and run migrations
	err := database.MigrateMysql()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("Error")
	}

	router.Run(":8080")

}
