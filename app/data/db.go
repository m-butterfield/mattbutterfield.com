package data

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

type ds struct {
	db *gorm.DB
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

func getDS() (*ds, error) {
	var logLevel logger.LogLevel
	if os.Getenv("SQL_LOGS") == "true" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Silent
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_SOCKET")), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}
	return &ds{db: db}, nil
}
