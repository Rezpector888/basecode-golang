package routes

import (
	"example.com/controllers/storage"
	"github.com/gin-gonic/gin"
)

func StorageRoute(c *gin.RouterGroup) {
	storageGroup := c.Group("storage")

	storageGroup.GET("download")
	storageGroup.POST("public", storage.UploadPublic)
}
