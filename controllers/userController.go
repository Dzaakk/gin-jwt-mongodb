package controllers

import (
	"net/http"

	"github.com/Dzaakk/gin-jwt-mongodb/database"
	"github.com/Dzaakk/gin-jwt-mongodb/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword() {

}

func VerifyPassword() {

}
func Signup() {

}

func Login() {

}

func GetUsers() {

}

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUserId(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
}
