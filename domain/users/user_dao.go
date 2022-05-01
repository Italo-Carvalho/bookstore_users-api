package users

import (
	"fmt"

	"github.com/italo-carvalho/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.ResErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.ResErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadResquestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewBadResquestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}
