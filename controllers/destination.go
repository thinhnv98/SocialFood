package controllers

import (
	"SocialFood/custom_models"
	"SocialFood/models"
	"SocialFood/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DestinationController struct {
	IDestinationService services.IDestinationService
}

func (_self DestinationController) GetDestination(ctx *gin.Context) {
	//Call service
	results, err := _self.IDestinationService.GetDestination()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Err": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": results,
	})
}

func (_self DestinationController) CreateDestination(ctx *gin.Context) {
	var destination custom_models.DestinationFully

	if err := ctx.ShouldBindJSON(&destination); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Call service
	err := _self.IDestinationService.CreateDestination(destination)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
	})
}

func (_self DestinationController) GetDestinationDetail(ctx *gin.Context) {
	idString := ctx.Param("id")

	ID, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	result, err := _self.IDestinationService.GetDestinationDetailByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
	return
}

func (_self DestinationController) Vote(ctx *gin.Context) {
	var vote models.DestinationRank
	if err := ctx.ShouldBindJSON(&vote); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Call service
	currentRank, err := _self.IDestinationService.Vote(vote)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"rank":      currentRank,
	})
	return
}
