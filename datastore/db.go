package datastore

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbFileName = "app.db"
)

var (
	db *sql.DB
)

func initDB() (err error) {
	db, err = sql.Open("sqlite3", dbFileName)
	return
}
