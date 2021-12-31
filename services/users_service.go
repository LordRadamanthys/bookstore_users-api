package services

import (
	"github.com/LordRadamanthys/bookstore_users-api/domain/users"
	"github.com/LordRadamanthys/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
	// return nil, errors.BadRequestError("error to create")
}
