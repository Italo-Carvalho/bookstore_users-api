package app

import (
	"github.com/italo-carvalho/bookstore_users-api/controllers/ping"
	"github.com/italo-carvalho/bookstore_users-api/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users/", users.CreateUser)

}
