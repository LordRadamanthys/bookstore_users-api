package services

import (
	"github.com/LordRadamanthys/bookstore_users-api/domain/users"
	"github.com/LordRadamanthys/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userID int) (*users.User, *errors.RestErr) {

	result := users.User{Id: userID}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}
