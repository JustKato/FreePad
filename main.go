package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	_, isDevelopment := os.LookupEnv("DEV_MODE")
	if isDevelopment {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize the router
	router := gin.Default()

	router.Run(":8080")

}
