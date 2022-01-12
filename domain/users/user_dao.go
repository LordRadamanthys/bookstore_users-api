package users

import (
	"github.com/LordRadamanthys/bookstore_users-api/datasources/mysql/users_db"
	"github.com/LordRadamanthys/bookstore_users-api/utils/date"
	"github.com/LordRadamanthys/bookstore_users-api/utils/errors"
	"github.com/LordRadamanthys/bookstore_users-api/utils/mysql_utils"
	_ "github.com/go-sql-driver/mysql"
)

const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	queryUpdateUser  = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created from users WHERE id = ?;"
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date.GetDateNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}
	user.Id = int(userID)
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(int(user.Id))

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}
