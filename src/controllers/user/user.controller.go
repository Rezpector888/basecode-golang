package user

import (
	"net/http"
	"strconv"

	"example.com/src/common/config"
	"example.com/src/common/utils"
	"example.com/src/controllers/user/dto"
	"example.com/src/services"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	search := c.DefaultQuery("search", "")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset <= 0 {
		offset = 0
	}

	users, total, err := services.UserFindAll(search, limit, offset)
	if err != nil {
		config.LogMessage("ERROR", err.Error())
		utils.ResponsePagination(c.Writer, http.StatusInternalServerError, "Failed to fetch user", nil, 0, limit, offset)
		return
	}

	utils.ResponsePagination(c.Writer, http.StatusOK, "success", users, total, limit, offset)

}

func FindOne(c *gin.Context) {
	userId := c.Param("id")

	user, err := services.UserFindById(userId)

	if err != nil {
		config.LogMessage("ERROR", err.Error())
		utils.ResponseError(c.Writer, http.StatusNotFound, "user not found")
		return
	}

	utils.ResponseSuccess(c.Writer, http.StatusOK, "success", user)
}

func Create(c *gin.Context) {
	var payload dto.CreateUserInput

	if err := c.ShouldBindJSON(&payload); err != nil {
		config.LogMessage("ERROR", err.Error())
		utils.ResponseError(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	user, err := services.UserCreate(payload)

	if err != nil {
		config.LogMessage("ERROR", err.Error())
		utils.ResponseError(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseSuccess(c.Writer, http.StatusCreated, "success", user)
}

func Update(c *gin.Context) {
	userId := c.Param("id")
	var payload dto.UpdateUserInput

	if err := c.ShouldBindJSON(&payload); err != nil {
		config.LogMessage("ERROR", err.Error())
		utils.ResponseError(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	user, err := services.UserUpdate(userId, payload)

	if err != nil {
		config.LogMessage("ERROR", err.Error())
		utils.ResponseError(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseSuccess(c.Writer, http.StatusOK, "success", user)

}

func Delete(c *gin.Context) {
	userId := c.Param("id")

	user, err := services.UserDelete(userId)

	if err != nil {
		config.LogMessage("ERROR", err.Error())
		utils.ResponseError(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseSuccess(c.Writer, http.StatusOK, "success", user)
}
