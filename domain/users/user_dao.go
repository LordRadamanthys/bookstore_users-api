package users

import (
	"fmt"
	"strings"

	"github.com/LordRadamanthys/bookstore_users-api/datasources/mysql/users_db"
	"github.com/LordRadamanthys/bookstore_users-api/utils/date"
	"github.com/LordRadamanthys/bookstore_users-api/utils/errors"
	_ "github.com/go-sql-driver/mysql"
)

var (
	userDB = make(map[int]*User)
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date.GetDateNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.BadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.InternalServerError(err.Error())
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = int(userID)
	return nil
}

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
