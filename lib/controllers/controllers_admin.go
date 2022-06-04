package controllers

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/JustKato/FreePad/lib/helper"
	"github.com/gin-gonic/gin"
)

func AdminMiddleware(router *gin.RouterGroup) {

	// Handl
	router.Use(func(ctx *gin.Context) {

		// Check which route we are accessing
		fmt.Println(`Accesing: `, ctx.Request.RequestURI)

		// Check if the request is other than the login request
		if ctx.Request.RequestURI != "/admin/login" {
			// Check if the user is logged-in

			fmt.Println(`Checking if admin`)

			if !IsAdmin(ctx) {
				// Not an admin, redirect to homepage
				ctx.Redirect(http.StatusTemporaryRedirect, "/")
				ctx.Abort()

				fmt.Println(`Not an admin!`)
				return
			}

		}

	})

}

func IsAdmin(ctx *gin.Context) bool {
	adminToken, err := ctx.Cookie("admin_token")
	if err != nil {
		return false
	}

	// Encode the real token
	sha512Hasher := sha512.New()
	sha512Hasher.Write([]byte(helper.GetAdminToken()))
	hashHexToken := sha512Hasher.Sum(nil)
	trueToken := hex.EncodeToString(hashHexToken)

	// Check if the user's admin token matches the token
	if adminToken != "" && adminToken == trueToken {
		// Yep, it's the admin!
		return true
	}

	// Definitely not an admin
	return false
}
