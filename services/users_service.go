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

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email

	}

	updateErr := current.Update()
	if updateErr != nil {
		return nil, err
	}
	return current, nil
}

func GetUser(userID int) (*users.User, *errors.RestErr) {

	result := users.User{Id: userID}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}

func DeleteUser(userID int) *errors.RestErr {
	user, err := GetUser(userID)

	if err != nil {
		return err
	}
	if deleteErr := user.Delete(); deleteErr != nil {
		return deleteErr
	}

	return nil
}

func FindByStatus(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
