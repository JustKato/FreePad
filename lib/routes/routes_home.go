package routes

import (
	"github.com/JustKato/FreePad/lib/helper"
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

		c.HTML(200, "page.html", gin.H{
			"title":        postName,
			"post_content": "",
			"domain_base":  helper.GetDomainBase(),
		})
	})

}
