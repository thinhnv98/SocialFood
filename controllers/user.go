package controllers

import (
	"SocialFood/models"
	"SocialFood/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	IUserService services.IUserService
}

func (_self UserController) Login(ctx *gin.Context) {
	var userRequest models.Account
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	isSuccess, err := _self.IUserService.Login(userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"IsSuccess": isSuccess,
	})
	return
}

func (_self UserController) Register(ctx *gin.Context) {
	var userRequest models.Account

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	isExisted, _ := _self.IUserService.IsExisted(userRequest)
	if isExisted {
		ctx.JSON(http.StatusOK, gin.H{
			"IsExisted": true,
			"IsSuccess": false,
		})
		return
	}

	isSuccess, _ := _self.IUserService.Register(userRequest)
	ctx.JSON(http.StatusOK, gin.H{
		"IsExisted": false,
		"IsSuccess": isSuccess,
	})
	return
}
