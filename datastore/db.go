package datastore

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbFileName             = "app.db"
	imageBaseURL           = "http://images.mattbutterfield.com/"
	insertImageQuery       = "INSERT INTO images (id, caption) VALUES (?, ?)"
	latestIDQuery          = "SELECT id FROM images ORDER BY id DESC LIMIT 1"
	selectRandomImageQuery = "SELECT id, caption FROM images WHERE id = (SELECT id FROM images ORDER BY RANDOM() LIMIT 1)"
)

var (
	db *sql.DB
)

type Page struct {
	Caption  string
	ImageURL string
}

func InitDB() (err error) {
	db, err = sql.Open("sqlite3", dbFileName)
	return
}

func GetLatestID() (id string, err error) {
	err = db.QueryRow(latestIDQuery).Scan(&id)
	return id, err
}

func GetRandomPage() (*Page, error) {
	var (
		imageID string
		caption sql.NullString
	)
	row := db.QueryRow(selectRandomImageQuery)
	err := row.Scan(&imageID, &caption)
	if err != nil {
		return nil, err
	}
	return &Page{
		Caption:  caption.String,
		ImageURL: imageBaseURL + imageID,
	}, nil
}

func SaveImage(keyName string, caption string) error {
	_, err := db.Exec(insertImageQuery, keyName, caption)
	return err
}
