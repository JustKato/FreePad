package routes

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/JustKato/FreePad/controllers/post"
	"github.com/JustKato/FreePad/helper"
	"github.com/JustKato/FreePad/models/database"
	"github.com/JustKato/FreePad/types"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func ApiRoutes(route *gin.RouterGroup) {

	// Bind the rate limiter
	helper.BindRateLimiter(route)

	route.POST("/post", func(ctx *gin.Context) {
		// Get the name of the post
		postName := ctx.PostForm("name")
		// Get the content of the post
		postContent := ctx.PostForm("content")

		// Try and run migrations
		err := database.MigrateMysql()
		if err != nil {
			fmt.Println("Error")
			fmt.Println(err)
			fmt.Println("Error")
		}

		// Create my post
		myPost, err := post.Create(postName, postContent)
		if err != nil {
			fmt.Println("Error", err)
			ctx.JSON(400, types.FreeError{
				Error:   err.Error(),
				Message: "There has been an error processing your request",
			})
			return
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
			fmt.Println("Error", err)
			ctx.JSON(400, types.FreeError{
				Error:   err.Error(),
				Message: "There has been an error processing your request",
			})
			return
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

	route.POST("/qr", func(ctx *gin.Context) {

		// Get the name of the post
		link := ctx.PostForm("link")

		// store the png somewhere
		var png []byte

		// Encode the link into a qr code
		png, err := qrcode.Encode(link, qrcode.High, 512)
		if err != nil {
			ctx.JSON(200, types.FreeError{
				Error:   fmt.Sprint(err),
				Message: "Failed to convert qr Code",
			})
			return
		}

		// Write the png to the response
		ctx.JSON(200, gin.H{
			"message": "Succesfully generated the QR",
			"qr":      "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(png),
		})
	})

}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Healthy",
	})
}
