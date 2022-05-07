package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/italo-carvalho/bookstore_users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.ResErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database repsonse")
	}

	switch sqlErr.Number {
	case 1062: // Duplicate entry
		return errors.NewBadResquestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")

}
