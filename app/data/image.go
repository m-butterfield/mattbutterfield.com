package data

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

const (
	baseSelectImageQuery = "SELECT id, caption, location FROM images "
	getImageByIDQuery    = baseSelectImageQuery + "WHERE id = ?"
	getLatestImageQuery  = baseSelectImageQuery + "ORDER BY id DESC LIMIT 1"
	getNextQuery         = baseSelectImageQuery + "WHERE id > ? ORDER BY id LIMIT 1"
	getPreviousQuery     = baseSelectImageQuery + "WHERE id < ? ORDER BY id DESC LIMIT 1"
	getRandomImageQuery  = baseSelectImageQuery + "WHERE id = (SELECT id FROM images ORDER BY RANDOM() LIMIT 1)"
	insertImageQuery     = "INSERT INTO images (id, caption, location) VALUES (?, ?, ?)"
	updateImageQuery     = "UPDATE images SET location = ?, caption = ? WHERE id = ?"
)

const (
	imageIDDateLayout = "20060102"
)

type ImageStore interface {
	GetImage(string) (*Image, error)
	GetLatestImage() (*Image, error)
	GetPrevNextImages(string) (*Image, *Image, error)
	GetRandomImage() (*Image, error)
	SaveImage(Image) error
	UpdateImage(string, string, string) error
}

type Image struct {
	ID       string
	Caption  string
	Location string
}

func (i Image) TimeFromID() (*time.Time, error) {
	if len(i.ID) < len(imageIDDateLayout) {
		return nil, fmt.Errorf("Unexpected id format: %s", i.ID)
	}
	t, err := time.Parse(imageIDDateLayout, i.ID[:len(imageIDDateLayout)])
	if err != nil {
		return nil, err
	}
	return &t, nil
}

type DBImageStore struct {
	DB *sql.DB
}

func NewDBImageStore(db *sql.DB) DBImageStore {
	return DBImageStore{DB: db}
}

func (store DBImageStore) GetImage(id string) (*Image, error) {
	return makeImageFromRow(store.DB.QueryRow(getImageByIDQuery, id))
}

func (store DBImageStore) GetLatestImage() (*Image, error) {
	return makeImageFromRow(store.DB.QueryRow(getLatestImageQuery))
}

func (store DBImageStore) GetPrevNextImages(id string) (*Image, *Image, error) {
	previous, err := makeImageFromRow(store.DB.QueryRow(getPreviousQuery, id))
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, err
	}
	next, err := makeImageFromRow(store.DB.QueryRow(getNextQuery, id))
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, err
	}
	return previous, next, nil
}

func (store DBImageStore) GetRandomImage() (*Image, error) {
	return makeImageFromRow(store.DB.QueryRow(getRandomImageQuery))
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

func (store DBImageStore) UpdateImage(id, location, caption string) error {
	captionPtr, locationPtr := &caption, &location
	if *captionPtr == "" {
		captionPtr = nil
	}
	if *locationPtr == "" {
		locationPtr = nil
	}
	_, err := store.DB.Exec(updateImageQuery, locationPtr, captionPtr, id)
	return err
}

func makeImageFromRow(row *sql.Row) (*Image, error) {
	var (
		id       string
		caption  sql.NullString
		location sql.NullString
	)
	err := row.Scan(&id, &caption, &location)
	if err != nil {
		return nil, err
	}
	return &Image{
		ID:       id,
		Caption:  caption.String,
		Location: location.String,
	}, nil
}
