package datastore

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	baseSelectImageQuery = "SELECT id, caption FROM images "
	getImageByIDQuery    = baseSelectImageQuery + "WHERE id = ? "
	getLatestImageQuery  = baseSelectImageQuery + "ORDER BY id DESC LIMIT 1"
	getRandomImageQuery  = baseSelectImageQuery + "WHERE id = (SELECT id FROM images ORDER BY RANDOM() LIMIT 1) "
	insertImageQuery     = "INSERT INTO images (id, caption) VALUES (?, ?)"
)

type Image struct {
	ID      string
	Caption string
}

func NewImage(id, caption string) *Image {
	return &Image{
		ID:      id,
		Caption: caption,
	}
}

func (I Image) Save() error {
	if db == nil {
		err := initDB()
		if err != nil {
			return err
		}
	}
	captionPtr := &I.Caption
	if *captionPtr == "" {
		captionPtr = nil
	}
	_, err := db.Exec(insertImageQuery, I.ID, captionPtr)
	return err
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
	return NewImage(id, caption.String), nil
}
