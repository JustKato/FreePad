package helper

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"

	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
)

func BindRateLimiter(router *gin.RouterGroup) {
	// Setup rate limitng
	rate := limiter.Rate{
		Period: 5 * time.Minute,
		Limit:  150,
	}

	// Initialize the memory storage
	store := memory.NewStore()

	// Generate the middleware
	middleware := mgin.NewMiddleware(limiter.New(store, rate))

	// Use the middleware
	router.Use(middleware)
}
