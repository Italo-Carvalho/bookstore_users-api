package users_db

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client *sql.DB
)

func loadEnv(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	return os.Getenv(key)
}

func init() { // call in the first time to import this package
	// DONT PRINT [dataSourceName] IN PROD!!!
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		loadEnv(mysql_users_username),
		loadEnv(mysql_users_password),
		loadEnv(mysql_users_host),
		loadEnv(mysql_users_schema),
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
