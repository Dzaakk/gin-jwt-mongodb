package routes

import (
	"github.com/Dzaakk/gin-jwt-mongodb/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.Engine) {
	routes.POST("users/signup", controllers.Signup())
	routes.POST("users/login", controllers.Login())
}
