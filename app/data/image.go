package data

import (
	"database/sql"
	"time"
)

const (
	baseSelectImageQuery = "SELECT id, caption, location, width, height, created_at FROM images "
	getImageByIDQuery    = baseSelectImageQuery + "WHERE id = $1"
	getRandomImageQuery  = baseSelectImageQuery + "WHERE id = (SELECT id FROM images ORDER BY RANDOM() LIMIT 1)"
	getImagesQuery       = baseSelectImageQuery + "WHERE created_at < $1 ORDER BY created_at DESC LIMIT $2"
	insertImageQuery     = "INSERT INTO images (id, caption, location, width, height, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
)

type Image struct {
	ID        string
	Caption   string
	Location  string
	Width     int
	Height    int
	CreatedAt time.Time
}

func (s *dbStore) GetImage(id string) (*Image, error) {
	return makeImageFromRow(s.db.QueryRow(getImageByIDQuery, id))
}

func (s *dbStore) GetImages(before time.Time, limit int) ([]*Image, error) {
	rows, err := s.db.Query(getImagesQuery, before, limit)
	if err != nil {
		return nil, err
	}
	var images []*Image
	for rows.Next() {
		image, err := makeImageFromRow(rows)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	return images, nil
}

func (s *dbStore) GetRandomImage() (*Image, error) {
	return makeImageFromRow(s.db.QueryRow(getRandomImageQuery))
}

func makeImageFromRow(row interface{ Scan(...interface{}) error }) (*Image, error) {
	var (
		caption  sql.NullString
		location sql.NullString
	)
	image := &Image{}
	if err := row.Scan(&image.ID, &caption, &location, &image.Width, &image.Height, &image.CreatedAt); err != nil {
		return nil, err
	}
	image.Caption = caption.String
	image.Location = location.String
	return image, nil
}

func (s *dbStore) SaveImage(id, caption, location string, width, height int, createdDate time.Time) error {
	stmt, err := s.db.Prepare(insertImageQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, nullString(caption), nullString(location), width, height, createdDate)
	if err != nil {
		return err
	}
	return nil
}
