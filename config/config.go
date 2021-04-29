package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "root"
	password string = "root"
	database string = "movie"
)

var (
	dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
)

//connection to db
func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
