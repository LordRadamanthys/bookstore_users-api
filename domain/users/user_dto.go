package users

import (
	"strings"

	"github.com/LordRadamanthys/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.BadRequestError("Invalid email address!")
	}
	return nil
}
