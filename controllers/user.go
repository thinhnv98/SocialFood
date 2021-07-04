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

	isSuccess, userID, err := _self.IUserService.Login(userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"IsSuccess": isSuccess,
		"ID":        userID,
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
			"ID":        0,
			"IsExisted": true,
			"IsSuccess": false,
		})
		return
	}

	ID, isSuccess, _ := _self.IUserService.Register(userRequest)
	ctx.JSON(http.StatusOK, gin.H{
		"ID":        ID,
		"IsExisted": false,
		"IsSuccess": isSuccess,
	})
	return
}
