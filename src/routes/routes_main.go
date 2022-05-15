package routes

import (
	"github.com/JustKato/FreePad/controllers/post"
	"github.com/JustKato/FreePad/helper"
	"github.com/gin-gonic/gin"
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

		// Try and get this post's data
		postData, err := post.Retrieve(postName)
		if err != nil {
			postData = &post.Post{
				Name:    postName,
				Content: "",
			}
		}

		c.HTML(200, "page.html", gin.H{
			"title":        postName,
			"post_content": postData.Content,
			"domain_base":  helper.GetDomainBase(),
		})
	})

}
