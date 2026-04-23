package routes

import (
	"example.com/common/middlewares"
	"example.com/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(c *gin.RouterGroup) {
	authGroup := c.Group("/auth")
	authGroup.POST("login", controllers.AuthController)
	authGroup.GET("me", middlewares.JwtAuthMiddleware(), controllers.AuthMeController)

}
