package datastore

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbFileName = "app.db"
)

func OpenDB() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", dbFileName)
	return
}
