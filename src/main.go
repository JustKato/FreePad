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
	routes.HomeRoutes(router)
	// Bind /api
	routes.ApiRoutes(router.Group("/api"))

	router.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8080")

}
