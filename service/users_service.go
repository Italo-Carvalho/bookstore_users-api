package service

import (
	"net/http"

	"github.com/italo-carvalho/bookstore_users-api/domain/users"
	"github.com/italo-carvalho/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.ResErr) {
	return &user, nil
	return &user, &errors.ResErr{
		Status: http.StatusInternalServerError,
	}
}
