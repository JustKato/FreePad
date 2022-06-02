package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware(router *gin.RouterGroup) {

	// Handl
	router.Use(func(ctx *gin.Context) {

		// Check which route we are accessing
		fmt.Println(`Accesing: `, ctx.Request.RequestURI)

	})

}
