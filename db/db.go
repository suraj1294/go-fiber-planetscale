package db

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var localUrl = ""

func NewDatabaseConnection() *sqlx.DB {

	dsn := os.Getenv("DSN")

	if dsn == "" {
		dsn = localUrl
	}

	Db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		panic(err)
	}

	return Db

}
