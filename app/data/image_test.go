package data

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

const (
	baseSelectImageRegex   = "^SELECT id, caption, location, width, height FROM images "
	SelectImageByIDRegex   = baseSelectImageRegex + "WHERE id = \\$1$"
	SelectRandomImageRegex = baseSelectImageRegex + "WHERE id = \\(SELECT id FROM images ORDER BY RANDOM\\(\\) LIMIT 1\\)$"
)

func TestGetImage(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	store := &dbStore{db: db}

	id, caption, location, width, height := "20040901_001.jpg", "hello", "NYC", 100, 200
	dbMock.ExpectQuery(SelectImageByIDRegex).WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption", "location", "width", "height"}).AddRow(id, caption, location, width, height))

	image, err := store.GetImage(id)
	if err != nil {
		t.Fatal(err)
	}
	err = dbMock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled database expectations: %s", err)
	}
	if image.ID != id {
		t.Errorf("Unexpected image id: %s != %s", id, image.ID)
	}
	if image.Caption != caption {
		t.Errorf("Unexpected image caption: %s != %s", caption, image.Caption)
	}
	if image.Location != location {
		t.Errorf("Unexpected image caption: %s != %s", caption, image.Location)
	}
	if image.Width != width {
		t.Errorf("Unexpected image width: %d != %d", image.Width, width)
	}
	if image.Height != height {
		t.Errorf("Unexpected image height: %d != %d", image.Height, height)
	}
}

func TestGetRandomImage(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	store := &dbStore{db: db}

	dbMock.ExpectQuery(SelectRandomImageRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption", "location", "width", "height"}).AddRow("20040901_001.jpg", nil, nil, 100, 200))

	_, err = store.GetRandomImage()
	if err != nil {
		t.Fatal(err)
	}
	err = dbMock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled database expectations: %s", err)
	}
}
