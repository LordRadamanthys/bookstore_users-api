package users_db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LordRadamanthys/bookstore_users-api/utils/environment"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client   *sql.DB
	username = environment.GoDotEnvVariable(mysql_users_username)
	password = environment.GoDotEnvVariable(mysql_users_password)
	host     = environment.GoDotEnvVariable(mysql_users_host)
	schema   = environment.GoDotEnvVariable(mysql_users_schema)
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema)
	var err error
	Client, err = sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
