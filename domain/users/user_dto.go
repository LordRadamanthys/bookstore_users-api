package users

import (
	"fmt"
	"strings"

	"github.com/LordRadamanthys/bookstore_utils-go/rest_errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate() *rest_errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))

	if user.Email == "" {
		return rest_errors.BadRequestError("Invalid email address!", nil)
	}
	user.Password = strings.TrimSpace(user.Password)

	fmt.Println(user.Password)
	if user.Password == "" {
		return rest_errors.BadRequestError("invalid password!", nil)
	}
	return nil
}
