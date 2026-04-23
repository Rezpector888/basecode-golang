package routes

import (
	"example.com/controllers/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(c *gin.RouterGroup) {
	userGroup := c.Group("/user")

	userGroup.GET("/", user.FindAll)
	userGroup.GET("/:id", user.FindOne)
	userGroup.POST("/", user.Create)
	userGroup.PATCH("/:id", user.Update)
	userGroup.DELETE("/:id", user.Delete)
}
