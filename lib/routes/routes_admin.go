package routes

import (
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/JustKato/FreePad/lib/controllers"
	"github.com/JustKato/FreePad/lib/helper"
	"github.com/JustKato/FreePad/lib/objects"
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

			// Set the cookie
			ctx.SetCookie("admin_token", hashToken, 60*60, "/", helper.GetDomainBase(), true, true)

			ctx.Request.Method = "GET"

			// Redirect the user to the admin page
			ctx.Redirect(http.StatusFound, "/admin/view")
			return
		} else {
			ctx.Request.Method = "GET"

			// Redirect the user to the admin page
			ctx.Redirect(http.StatusFound, "/admin/login?fail")
			return
		}

	})

	router.GET("/delete/:padname", func(ctx *gin.Context) {
		// Get the pad name that we bout' to delete
		padName := ctx.Param("padname")

		// Try and get the pad, check if valid
		pad := objects.GetPost(padName, false)

		// Delete the pad
		err := pad.Delete()
		fmt.Println(err)

		// Redirect the user to the admin page
		ctx.Redirect(http.StatusFound, "/admin/view")
	})

	// Admin view route
	router.GET("/view", func(ctx *gin.Context) {

		// Get all of the pads as a listing
		padList := objects.GetAllPosts()

		ctx.HTML(200, "admin_view.html", gin.H{
			"title":       "Admin",
			"padList":     padList,
			"domain_base": helper.GetDomainBase(),
		})

	})

}
