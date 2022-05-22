package controllers

import "github.com/gin-gonic/gin"

func ApplyHeaders(router *gin.Engine) {

	router.Use(func(ctx *gin.Context) {
		// Apply the header
		ctx.Header("FreePad-Version", "1.1.1")

		// Move on
		ctx.Next()
	})

}
