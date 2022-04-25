package users

import (
	"net/http"

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
	c.String(http.StatusNotImplemented, "implement me!")
}

func SearhUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
