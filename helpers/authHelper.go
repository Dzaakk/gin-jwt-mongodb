package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}
func MatchUserTypeToUserId(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	UID := c.GetString("uid")
	err = nil

	if userType == "USER" && UID != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	err = CheckUserType(c, userType)
	return err
}
