package data

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"os"
)

type Store interface {
	GetImage(string) (*Image, error)
	GetRandomImage() (*Image, error)
	GetSongs() ([]*Song, error)
	SaveSong(string, string) error
}

type dbStore struct {
	db *sql.DB
}

func Connect() (Store, error) {
	db, err := sql.Open("pgx", os.Getenv("DB_SOCKET"))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &dbStore{db: db}, nil
}
