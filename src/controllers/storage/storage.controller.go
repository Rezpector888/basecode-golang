package storage

import (
	"net/http"

	"example.com/src/common/config"
	"example.com/src/common/utils"
	"example.com/src/services"
	"github.com/gin-gonic/gin"
)

func UploadPublic(c *gin.Context) {

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		utils.ResponseError(c.Writer, http.StatusBadRequest, "file is required")
	}

	result, err := services.StorageUploadPublic(file, *fileHeader)

	if err != nil {
		config.LogMessage("ERROR", err.Error())
		utils.ResponseError(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseSuccess(c.Writer, http.StatusCreated, "success", result)

}
