package users

import (
	"fmt"

	"github.com/LordRadamanthys/bookstore_users-api/datasources/mysql/users_db"
	"github.com/LordRadamanthys/bookstore_users-api/utils/mysql_utils"
	"github.com/LordRadamanthys/bookstore_utils-go/rest_errors"
	_ "github.com/go-sql-driver/mysql"
)

const (
	queryInsertUser           = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES (?,?,?,?,?,?);"
	queryUpdateUser           = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	queryGetUser              = "SELECT id, first_name, last_name, email, status, date_created from users WHERE id = ?;"
	queryDeleteUser           = "DELETE FROM users WHERE id =?;"
	queryFindUserByStatus     = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
	indexUniqueEmail          = "email_UNIQUE"
	errorNoRows               = "no rows in result set"
)

func (user *User) Save() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return rest_errors.InternalServerError(err.Error(), err)
	}
	defer stmt.Close()

	fmt.Printf(user.Status)
	insertResult, saveErr := stmt.Exec(user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
		user.Password,
		user.Status)

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

func (user *User) Update() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		return rest_errors.InternalServerError(err.Error(), err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Get() *rest_errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		return rest_errors.InternalServerError(err.Error(), err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(int(user.Id))

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Delete() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return rest_errors.InternalServerError(err.Error(), err)
	}
	defer stmt.Close()

	fmt.Println(user)
	if _, deleteErr := stmt.Exec(user.Id); deleteErr != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *rest_errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)

	if err != nil {
		return nil, rest_errors.InternalServerError(err.Error(), err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, rest_errors.InternalServerError(err.Error(), err)
	}

	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User

		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, rest_errors.NotFoundError(fmt.Sprintf("no users matching status %s", status), nil)
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindEmailAndPassword)

	if err != nil {
		return rest_errors.InternalServerError(err.Error(), err)
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}
