package controllers

import (
	"SocialFood/custom_models"
	"SocialFood/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RestaurantController struct {
	IRestaurantService services.IRestaurantService
}

func (_self RestaurantController) GetRestaurants(ctx *gin.Context) {
	//Call service
	results, err := _self.IRestaurantService.GetRestaurants()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": results,
	})
}

func (_self RestaurantController) GetRestaurantDetailByID(ctx *gin.Context) {
	idString := ctx.Param("id")

	ID, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	//Call service
	result, err := _self.IRestaurantService.GetRestaurantDetailByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
	return
}

func (_self RestaurantController) CreateRestaurant(ctx *gin.Context) {
	var restaurant custom_models.NewRestaurantObj

	if err := ctx.ShouldBindJSON(&restaurant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Call service
	err := _self.IRestaurantService.CreateRestaurant(restaurant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
	})
}

func (_self RestaurantController) CreateNewDish(ctx *gin.Context) {
	var dishRequest custom_models.DishWithResID

	if err := ctx.ShouldBindJSON(&dishRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Call service
	err := _self.IRestaurantService.CreateDish(dishRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
	})
}

func (_self RestaurantController) GetDishesByResID(ctx *gin.Context) {
	idString := ctx.Param("id")

	ID, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	//Call service
	results, err := _self.IRestaurantService.GetDishesByResID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	dishesResult := []custom_models.Dish{}
	for _, result := range results {
		dishesResult = append(dishesResult, custom_models.Dish{
			Name:      result.Name.String,
			Price:     result.Price.String,
			Photo:     result.Photo.String,
			NumPhotos: result.Numphotos.Int,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": dishesResult,
	})
	return
}

func (_self RestaurantController) Review(ctx *gin.Context) {
	var review custom_models.ReviewAndVote

	if err := ctx.ShouldBindJSON(&review); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Call service
	rank, reviews, err := _self.IRestaurantService.Review(review)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"rank":      rank,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"rank":      rank,
		"reviews":   reviews,
	})
}
