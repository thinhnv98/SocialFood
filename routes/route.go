package routes

import (
	"database/sql"

	"SocialFood/controllers"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Db     *sql.DB
	Server *gin.Engine
}

func (_self Route) Register() {
	// Common Controller
	common := controllers.Common{}
	commonRoutes := _self.Server.Group("/api")
	{
		commonRoutes.GET("/ping", common.Ping)
	}

	// Add More Controller
}
