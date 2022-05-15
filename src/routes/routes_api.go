package routes

import (
	"fmt"
	"net/url"

	"github.com/JustKato/FreePad/controllers/post"
	"github.com/JustKato/FreePad/helper"
	"github.com/JustKato/FreePad/types"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(route *gin.RouterGroup) {

	// Add in
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ok!",
		})
	})

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

	route.GET("/posts", func(ctx *gin.Context) {
		// Return the post list
		ctx.JSON(200, post.GetPostList())
	})

}
