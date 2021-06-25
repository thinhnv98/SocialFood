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
	common := controllers.UserController{
		IUserService: services.UserService{
			Db: _self.Db,
		},
	}
	commonRoutes := _self.Server.Group("/api")
	{
		commonRoutes.POST("/login", common.Login)
		commonRoutes.POST("/register", common.Register)
	}

	// Add More Controller
}
