package data

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"os"
	"time"
)

type Store interface {
	GetImage(string) (*Image, error)
	GetImages(time.Time, int) ([]*Image, error)
	GetRandomImage() (*Image, error)
	GetSongs() ([]*Song, error)
	SaveSong(string, string, time.Time) error
	SaveImage(string, string, string, int, int, time.Time) error
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

func nullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
