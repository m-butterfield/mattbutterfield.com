package data

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"time"
)

const (
	baseSelectImageQuery  = "SELECT id, caption, location, width, height FROM images "
	getImageByIDQuery     = baseSelectImageQuery + "WHERE id = $1"
	getLatestImageQuery   = baseSelectImageQuery + "ORDER BY id DESC LIMIT 1"
	getNextImageQuery     = baseSelectImageQuery + "WHERE id > $1 ORDER BY id LIMIT 1"
	getPreviousImageQuery = baseSelectImageQuery + "WHERE id < $1 ORDER BY id DESC LIMIT 1"
	getRandomImageQuery   = baseSelectImageQuery + "WHERE id = (SELECT id FROM images ORDER BY RANDOM() LIMIT 1)"
	insertImageQuery      = "INSERT INTO images (id, caption, location) VALUES ($1, $2, $3)"
	updateImageQuery      = "UPDATE images SET location = $1, caption = $2 WHERE id = $3"
)

const (
	imageIDDateLayout = "20060102"
)

type DBStore interface {
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
	Width    int
	Height   int
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

type dbStore struct {
	db *sql.DB
}

func MakeDBStore(dbPath string) (DBStore, error) {
	db, err := sql.Open("pgx", dbPath)
	if err != nil {
		return nil, err
	}
	return dbStore{db: db}, nil
}

func (store dbStore) GetImage(id string) (*Image, error) {
	return makeImageFromRow(store.db.QueryRow(getImageByIDQuery, id))
}

func (store dbStore) GetLatestImage() (*Image, error) {
	return makeImageFromRow(store.db.QueryRow(getLatestImageQuery))
}

func (store dbStore) GetPrevNextImages(id string) (*Image, *Image, error) {
	previous, err := makeImageFromRow(store.db.QueryRow(getPreviousImageQuery, id))
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, err
	}
	next, err := makeImageFromRow(store.db.QueryRow(getNextImageQuery, id))
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, err
	}
	return previous, next, nil
}

func (store dbStore) GetRandomImage() (*Image, error) {
	return makeImageFromRow(store.db.QueryRow(getRandomImageQuery))
}

func (store dbStore) SaveImage(image Image) error {
	captionPtr, locationPtr := &image.Caption, &image.Location
	if *captionPtr == "" {
		captionPtr = nil
	}
	if *locationPtr == "" {
		locationPtr = nil
	}
	_, err := store.db.Exec(insertImageQuery, image.ID, captionPtr, locationPtr)
	return err
}

func (store dbStore) UpdateImage(id, location, caption string) error {
	captionPtr, locationPtr := &caption, &location
	if *captionPtr == "" {
		captionPtr = nil
	}
	if *locationPtr == "" {
		locationPtr = nil
	}
	_, err := store.db.Exec(updateImageQuery, locationPtr, captionPtr, id)
	return err
}

func makeImageFromRow(row *sql.Row) (*Image, error) {
	var (
		caption  sql.NullString
		location sql.NullString
	)
	image := &Image{}
	err := row.Scan(&image.ID, &caption, &location, &image.Width, &image.Height)
	if err != nil {
		return nil, err
	}
	image.Caption = caption.String
	image.Location = location.String
	return image, nil
}
