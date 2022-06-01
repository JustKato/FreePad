package routes

import (
	"net/http"
	"net/url"
	"time"

	"github.com/JustKato/FreePad/lib/helper"
	"github.com/JustKato/FreePad/lib/objects"
	"github.com/gin-gonic/gin"
	"github.com/mrz1836/go-sanitize"
)

func HomeRoutes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title":       "HomePage",
			"domain_base": helper.GetDomainBase(),
		})
	})

	router.GET("/:post", func(c *gin.Context) {
		// Get the post we are looking for.
		postName := c.Param("post")

		if postName == `views_storage.json` {
			// Redirect the user to the homepage as this is a reserved keyword
			c.Redirect(http.StatusPermanentRedirect, "/")
			// Do not proceed further
			return
		}

		// Get the maximum pad size, so that we may notify the client-side to match server-side
		maximumPadSize := helper.GetMaximumPadSize()

		// Sanitize the postName
		newPostName, err := url.QueryUnescape(postName)
		if err == nil {
			postName = newPostName
		}
		postName = sanitize.XSS(sanitize.SingleLine(postName))

		post := objects.GetPost(postName)

		c.HTML(200, "page.html", gin.H{
			"title":          postName,
			"post_content":   post.Content,
			"maximumPadSize": maximumPadSize,
			"last_modified":  post.LastModified,
			"views":          post.Views,
			"domain_base":    helper.GetDomainBase(),
		})
	})

	router.POST("/:post", func(c *gin.Context) {
		// Get the post we are looking for.
		postName := c.Param("post")
		postContent := c.PostForm("content")

		// Sanitize the postName
		newPostName, err := url.QueryUnescape(postName)
		if err == nil {
			postName = newPostName
		}
		postName = sanitize.XSS(sanitize.SingleLine(postName))

		p := objects.Post{
			Name:         postName,
			Content:      postContent,
			Views:        0, // This can just be ignored
			LastModified: time.Now().Format("02/01/2006 03:04:05 PM"),
		}

		// Write the post
		err = objects.WritePost(p)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err,
			})

			// End
			return
		}

		// Return the success message
		c.JSON(200, gin.H{
			"pad": p,
		})
	})

}
