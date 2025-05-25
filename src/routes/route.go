package routes

import (
	"example.com/src/common/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.JSON(200, gin.H{
			"status":  true,
			"message": "API is ready",
		})
	})

	AuthRouter(&router.RouterGroup)

	auth := router.Group("/")
	auth.Use(middlewares.JwtAuthMiddleware())

	UserRoutes(auth)

}
