package datastore

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	baseSelectImageQuery = "SELECT id, caption, location FROM images "
	getImageByIDQuery    = baseSelectImageQuery + "WHERE id = ?"
	getLatestImageQuery  = baseSelectImageQuery + "ORDER BY id DESC LIMIT 1"
	getRandomImageQuery  = baseSelectImageQuery + "WHERE id = (SELECT id FROM images ORDER BY RANDOM() LIMIT 1)"
	insertImageQuery     = "INSERT INTO images (id, caption, location) VALUES (?, ?, ?)"
)

type ImageStore interface {
	SaveImage(image Image) error
	GetImage(id string) (*Image, error)
	GetLatestImage() (*Image, error)
	GetRandomImage() (*Image, error)
}

type Image struct {
	ID      string
	Caption string
	Location string
}

type DBImageStore struct {
	DB *sql.DB
}

func (store DBImageStore) SaveImage(image Image) error {
	captionPtr, locationPtr := &image.Caption, &image.Location
	if *captionPtr == "" {
		captionPtr = nil
	}
	if *locationPtr == "" {
		locationPtr = nil
	}
	_, err := store.DB.Exec(insertImageQuery, image.ID, captionPtr, locationPtr)
	return err
}

func (store DBImageStore) GetImage(id string) (*Image, error) {
	return makeImageFromRow(store.DB.QueryRow(getImageByIDQuery, id))
}

func (store DBImageStore) GetLatestImage() (*Image, error) {
	return makeImageFromRow(store.DB.QueryRow(getLatestImageQuery))
}

func (store DBImageStore) GetRandomImage() (*Image, error) {
	return makeImageFromRow(store.DB.QueryRow(getRandomImageQuery))
}

func makeImageFromRow(row *sql.Row) (*Image, error) {
	var (
		id      string
		caption sql.NullString
		location sql.NullString
	)
	err := row.Scan(&id, &caption, &location)
	if err != nil {
		return nil, err
	}
	return &Image{
		ID:      id,
		Caption: caption.String,
		Location: location.String,
	}, nil
}
