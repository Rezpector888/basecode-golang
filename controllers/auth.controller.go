package controllers

import (
	"net/http"

	"example.com/common/utils"
	"example.com/services"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func AuthController(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c.Writer, 400, err.Error())
		return
	}

	result, err := services.AuthLogin(input.Email, input.Password)
	if err != nil {
		utils.ResponseError(c.Writer, 400, err.Error())
		return
	}

	token, err := services.AuthCreateToken(result)
	if err != nil {
		utils.ResponseError(c.Writer, 400, err.Error())
		return
	}

	response := gin.H{
		"token": token,
	}

	utils.ResponseSuccess(c.Writer, 200, "success", response)
}

func AuthMeController(c *gin.Context) {
	userId, exist := c.Get("user_id")

	if !exist {
		utils.ResponseError(c.Writer, http.StatusUnauthorized, "User id not found")
		return
	}

	userIdStr, ok := userId.(string)
	if !ok {
		utils.ResponseError(c.Writer, http.StatusUnauthorized, "Invalid user id type")
		return
	}

	result, err := services.UserFindById(userIdStr)

	if err != nil {
		utils.ResponseError(c.Writer, 400, err.Error())
	}

	utils.ResponseSuccess(c.Writer, 200, "success", result)
}
