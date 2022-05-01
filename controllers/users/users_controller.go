package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/italo-carvalho/bookstore_users-api/domain/users"
	"github.com/italo-carvalho/bookstore_users-api/service"
	"github.com/italo-carvalho/bookstore_users-api/utils/errors"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadResquestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := service.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadResquestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := service.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)

}

func SearhUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
