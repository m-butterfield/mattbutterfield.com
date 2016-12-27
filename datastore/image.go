package datastore

import (
	"database/sql"
	"encoding/base64"

	_ "github.com/mattn/go-sqlite3"
)

const (
	baseSelectImageQuery = "SELECT id, caption FROM images "
	getImageByIDQuery    = baseSelectImageQuery + "WHERE id = ? "
	getLatestImageQuery  = baseSelectImageQuery + "ORDER BY id DESC LIMIT 1"
	getRandomImageQuery  = baseSelectImageQuery + "WHERE id = (SELECT id FROM images ORDER BY RANDOM() LIMIT 1) "
	imageBaseURL         = "http://images.mattbutterfield.com/"
	insertImageQuery     = "INSERT INTO images (id, caption) VALUES (?, ?)"
)

type Image struct {
	ID      string
	URL     string
	Caption string
}

func (I Image) EncodeID() string {
	return base64.StdEncoding.EncodeToString([]byte(I.ID))
}

func GetImage(id string) (*Image, error) {
	if db == nil {
		err := initDB()
		if err != nil {
			return nil, err
		}
	}
	return makeImageFromRow(db.QueryRow(getImageByIDQuery, id))
}

func DecodeImageID(encodedID string) (string, error) {
	imageID, err := base64.StdEncoding.DecodeString(encodedID)
	if err != nil {
		return "", err
	}
	return string(imageID), nil
}

func GetLatestImage() (*Image, error) {
	if db == nil {
		err := initDB()
		if err != nil {
			return nil, err
		}
	}
	return makeImageFromRow(db.QueryRow(getLatestImageQuery))
}

func GetRandomImage() (*Image, error) {
	if db == nil {
		err := initDB()
		if err != nil {
			return nil, err
		}
	}
	return makeImageFromRow(db.QueryRow(getRandomImageQuery))
}

func makeImageFromRow(row *sql.Row) (*Image, error) {
	var (
		id      string
		caption sql.NullString
	)
	err := row.Scan(&id, &caption)
	if err != nil {
		return nil, err
	}
	return &Image{
		ID:      id,
		URL:     imageBaseURL + id,
		Caption: caption.String,
	}, nil
}

func SaveImage(keyName string, caption *string) error {
	if db == nil {
		err := initDB()
		if err != nil {
			return err
		}
	}
	_, err := db.Exec(insertImageQuery, keyName, caption)
	return err
}
