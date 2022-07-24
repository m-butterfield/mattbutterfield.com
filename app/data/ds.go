package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

type ds struct {
	db *gorm.DB
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
