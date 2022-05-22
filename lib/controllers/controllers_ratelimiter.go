package controllers

import (
	"time"

	"github.com/JustKato/FreePad/lib/helper"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"

	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
)

// Apply the rate limiter to the gin Engine
func DoRateLimit(router *gin.Engine) {

	// Initialize the rate limiter
	rate := limiter.Rate{
		Period: 5 * time.Minute,
		Limit:  int64(helper.GetApiBanLimit()),
	}

	// Initialize the memory storage
	store := memory.NewStore()

	// Initialize the limiter instance
	instance := limiter.New(store, rate)

	// Create the gin middleware
	middleWare := mgin.NewMiddleware(instance)

	// use the middleware in gin
	router.Use(middleWare)
}
