package routes

import "github.com/gin-gonic/gin"

func HealthRoutes(route *gin.Engine) {

	// Add in
	route.GET("/health", healthCheck)

}

func healthCheck(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Healthy",
	})

}
