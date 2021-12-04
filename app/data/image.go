package data

import (
	"database/sql"
	"time"
)

const (
	baseSelectImageQuery = "SELECT id, caption, location, width, height, date FROM images "
	getImageByIDQuery    = baseSelectImageQuery + "WHERE id = $1"
	getRandomImageQuery  = baseSelectImageQuery + "WHERE id = (SELECT id FROM images ORDER BY RANDOM() LIMIT 1)"
)

type Image struct {
	ID       string
	Caption  string
	Location string
	Width    int
	Height   int
	Date     time.Time
}

func (s *dbStore) GetImage(id string) (*Image, error) {
	return makeImageFromRow(s.db.QueryRow(getImageByIDQuery, id))
}

func (s *dbStore) GetRandomImage() (*Image, error) {
	return makeImageFromRow(s.db.QueryRow(getRandomImageQuery))
}

func makeImageFromRow(row *sql.Row) (*Image, error) {
	var (
		caption  sql.NullString
		location sql.NullString
	)
	image := &Image{}
	if err := row.Scan(&image.ID, &caption, &location, &image.Width, &image.Height, &image.Date); err != nil {
		return nil, err
	}
	image.Caption = caption.String
	image.Location = location.String
	return image, nil
}
