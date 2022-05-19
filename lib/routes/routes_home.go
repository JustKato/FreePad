package routes

import (
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

		// Sanitize the postName
		newPostName, err := url.QueryUnescape(postName)
		if err == nil {
			postName = newPostName
		}
		postName = sanitize.AlphaNumeric(postName, true)

		post := objects.GetPost(postName)

		c.HTML(200, "page.html", gin.H{
			"title":         postName,
			"post_content":  post.Content,
			"last_modified": post.LastModified,
			"domain_base":   helper.GetDomainBase(),
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
		postName = sanitize.AlphaNumeric(postName, true)

		p := objects.Post{
			Name:         postName,
			Content:      postContent,
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
