package users

import (
	"fmt"

	"github.com/LordRadamanthys/bookstore_users-api/utils/date"
	"github.com/LordRadamanthys/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int]*User)
)

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.BadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.BadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	user.DateCreated = date.GetDateNowString()
	userDB[user.Id] = user
	return nil
}

func (user *User) Get() *errors.RestErr {

	result := userDB[user.Id]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated
	return nil
}
