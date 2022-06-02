package routes

import (
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/JustKato/FreePad/lib/controllers"
	"github.com/JustKato/FreePad/lib/helper"
	"github.com/gin-gonic/gin"

	"crypto/sha512"
)

var adminLoginToken string = ""

func AdminRoutes(router *gin.RouterGroup) {

	adminLoginToken = helper.GetAdminToken()

	// Apply the admin middleware for identification
	controllers.AdminMiddleware(router)

	// Admin login route
	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(200, "admin_login.html", gin.H{
			"title":       "Login Login",
			"domain_base": helper.GetDomainBase(),
		})
	})

	router.POST("/login", func(ctx *gin.Context) {

		// Get the value of the admin token
		adminToken := ctx.PostForm("admin-token")

		// Check if the input admin token matches our admin token
		if adminLoginToken != "" && adminLoginToken == adminToken {

			sha512Hasher := sha512.New()
			sha512Hasher.Write([]byte(adminToken))

			// Set the cookie to be an admin
			hashHexToken := sha512Hasher.Sum(nil)
			hashToken := hex.EncodeToString(hashHexToken)

			fmt.Println(hashToken)

			// Set the cookie
			ctx.SetCookie("admin_token", hashToken, 60*60, "/", helper.GetDomainBase(), true, true)

			ctx.Request.Method = "GET"

			// Redirect the user to the admin page
			ctx.Redirect(http.StatusTemporaryRedirect, "/admin")
			return
		} else {
			ctx.Request.Method = "GET"

			// Redirect the user to the admin page
			ctx.Redirect(http.StatusTemporaryRedirect, "/admin/login?fail")
			return
		}

	})

	// Admin view route
	router.GET("/", func(ctx *gin.Context) {

		adminToken, err := ctx.Cookie("admin_token")
		if err != nil {
			adminToken = ""
		}

		ctx.JSON(200, gin.H{
			`adminToken`: adminToken,
		})
	})

}
