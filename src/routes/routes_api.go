package routes

import (
	"fmt"
	"net/url"

	"github.com/JustKato/FreePad/controllers/post"
	"github.com/JustKato/FreePad/helper"
	"github.com/JustKato/FreePad/models/database"
	"github.com/JustKato/FreePad/types"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(route *gin.RouterGroup) {

	// Bind the rate limiter
	helper.BindRateLimiter(route)

	route.POST("/post", func(ctx *gin.Context) {
		// Get the name of the post
		postName := ctx.PostForm("name")
		// Get the content of the post
		postContent := ctx.PostForm("content")

		// Create my post
		myPost, err := post.Create(postName, postContent)
		if err != nil {
			ctx.JSON(400, types.FreeError{
				Error:   err.Error(),
				Message: "There has been an error processing your request",
			})
		}

		ctx.JSON(200, gin.H{
			"message": "Post succesfully created",
			"post":    myPost,
			"link":    fmt.Sprintf("%s/%s", helper.GetDomainBase(), url.QueryEscape(myPost.Name)),
		})

	})

	route.GET("/post", func(ctx *gin.Context) {
		// Get the name of the post
		postName := ctx.PostForm("name")

		myPost, err := post.Retrieve(postName)
		if err != nil {
			ctx.JSON(400, types.FreeError{
				Error:   err.Error(),
				Message: "There has been an error processing your request",
			})
		}

		// Return the post list
		ctx.JSON(200, myPost)
	})

	route.GET("/posts", func(ctx *gin.Context) {
		// Return the post list
		ctx.JSON(200, post.GetPostList())
	})

	// Add in health checks
	route.GET("/health", healthCheck)

	route.POST("/test", func(ctx *gin.Context) {
		ctx.JSON(200, database.MigrationUpdate())
	})
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Healthy",
	})
}
