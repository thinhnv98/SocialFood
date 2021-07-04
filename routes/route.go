package routes

import (
	"SocialFood/services"
	"database/sql"

	"SocialFood/controllers"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Db     *sql.DB
	Server *gin.Engine
}

func (_self Route) Register() {
	// UserController Controller
	user := controllers.UserController{
		IUserService: services.UserService{
			Db: _self.Db,
		},
	}

	destination := controllers.DestinationController{
		IDestinationService: services.DestinationService{
			Db: _self.Db,
		},
	}

	restaurant := controllers.RestaurantController{
		IRestaurantService: services.RestaurantService{
			Db: _self.Db,
		},
	}

	commonRoutes := _self.Server.Group("/api")
	{
		commonRoutes.POST("/login", user.Login)
		commonRoutes.POST("/register", user.Register)
		commonRoutes.POST("/create-destination", destination.CreateDestination)

		commonRoutes.POST("/destination/vote", destination.Vote)
		commonRoutes.POST("/create-restaurant", restaurant.CreateRestaurant)
		commonRoutes.POST("/create-dish", restaurant.CreateNewDish)
		commonRoutes.POST("/review-restaurant", restaurant.Review)

		//*****************************************

		commonRoutes.GET("/destinations", destination.GetDestination)
		commonRoutes.GET("/destination-detail/:id", destination.GetDestinationDetail)

		commonRoutes.GET("/restaurants", restaurant.GetRestaurants)
		commonRoutes.GET("/restaurants-detail/:id", restaurant.GetRestaurantDetailByID)
		commonRoutes.GET("/restaurants-dishes/:id", restaurant.GetDishesByResID)
	}

	// Add More Controller
}
