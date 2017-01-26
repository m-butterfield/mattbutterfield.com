package datastore

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbPath string) (*sql.DB, error) {
	return sql.Open("sqlite3", dbPath)
}
