package routes

import (
	"fmt"
	"net/url"

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

		fmt.Println("Sanitizing ", postName)

		// Sanitize the postName
		newPostName, err := url.QueryUnescape(postName)
		if err == nil {
			postName = newPostName
		}
		postName = sanitize.AlphaNumeric(postName, true)

		fmt.Println("Fetching ", postName)

		post := objects.GetPost(postName)

		c.HTML(200, "page.html", gin.H{
			"title":        postName,
			"post_content": post.Content,
			"domain_base":  helper.GetDomainBase(),
		})
	})

}
