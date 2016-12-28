package datastore

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	baseSelectImageQuery = "SELECT id, caption FROM images "
	getImageByIDQuery    = baseSelectImageQuery + "WHERE id = ?"
	getLatestImageQuery  = baseSelectImageQuery + "ORDER BY id DESC LIMIT 1"
	getRandomImageQuery  = baseSelectImageQuery + "WHERE id = (SELECT id FROM images ORDER BY RANDOM() LIMIT 1)"
	insertImageQuery     = "INSERT INTO images (id, caption) VALUES (?, ?)"
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
}

type DBImageStore struct {
	DB *sql.DB
}

func (store DBImageStore) SaveImage(image Image) error {
	captionPtr := &image.Caption
	if *captionPtr == "" {
		captionPtr = nil
	}
	_, err := store.DB.Exec(insertImageQuery, image.ID, captionPtr)
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
	)
	err := row.Scan(&id, &caption)
	if err != nil {
		return nil, err
	}
	return &Image{
		ID:      id,
		Caption: caption.String,
	}, nil
}
