package routes

import (
	"github.com/Dzaakk/gin-jwt-mongodb/controllers"
	"github.com/Dzaakk/gin-jwt-mongodb/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {

	routes.Use(middleware.Authenticate())
	routes.GET("/users", controllers.GetUsers())
	routes.GET("users/:user_id", controllers.GetUserById())
}
