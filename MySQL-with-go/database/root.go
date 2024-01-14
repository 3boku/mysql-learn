package database

import (
	"database/sql"
)

type Database struct {
	db *sql.DB
}

func NewDatabse(dataSourceName string) *Database {
	d := new(Database)

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	return d
}
