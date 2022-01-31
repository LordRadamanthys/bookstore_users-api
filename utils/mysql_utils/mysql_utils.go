package mysql_utils

import (
	"strings"

	"github.com/LordRadamanthys/bookstore_utils-go/rest_errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return rest_errors.NotFoundError("no record matching given id", sqlErr)
		}
		return rest_errors.InternalServerError("error parsing database response", sqlErr)
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.BadRequestError("invalid data", sqlErr)
	}
	return rest_errors.InternalServerError("error processing request", sqlErr)
}
