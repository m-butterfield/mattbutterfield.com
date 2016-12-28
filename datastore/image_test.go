package datastore

import (
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"time"
	"fmt"
)

const (
	InsertImageRegex       = "^INSERT INTO images \\(id, caption, location\\) VALUES \\(\\?, \\?, \\?\\)$"
	SelectImageByIDRegex   = "^SELECT id, caption, location FROM images WHERE id = \\?$"
	SelectLatestImageRegex = "^SELECT id, caption, location FROM images ORDER BY id DESC LIMIT 1$"
	SelectRandomImageRegex = "^SELECT id, caption, location FROM images WHERE id = \\(SELECT id FROM images ORDER BY RANDOM\\(\\) LIMIT 1\\)$"
)

func TestGetImage(t *testing.T) {
	db, db_mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	id, caption, location := "20040901_001.jpg", "hello", "NYC"
	db_mock.ExpectQuery(SelectImageByIDRegex).WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption", "location"}).AddRow(id, caption, location))

	store := DBImageStore{DB: db}
	image, err := store.GetImage(id)
	if err != nil {
		t.Fatal(err)
	}
	if err := db_mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
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
}

func TestGetLatestImage(t *testing.T) {
	db, db_mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	db_mock.ExpectQuery(SelectLatestImageRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption", "location"}).AddRow("20040901_001.jpg", nil, nil))

	store := DBImageStore{DB: db}
	_, err = store.GetLatestImage()
	if err != nil {
		t.Fatal(err)
	}
	if err := db_mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestGetRandomImage(t *testing.T) {
	db, db_mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	db_mock.ExpectQuery(SelectRandomImageRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption", "location"}).AddRow("20040901_001.jpg", nil, nil))

	store := DBImageStore{db}
	_, err = store.GetRandomImage()
	if err != nil {
		t.Fatal(err)
	}
	if err := db_mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestSaveImage(t *testing.T) {
	db, db_mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	id, caption, location := "20040901_001.jpg", "hello", "NYC"
	db_mock.ExpectExec(InsertImageRegex).WithArgs(id, caption, location).WillReturnResult(sqlmock.NewResult(1, 1))

	image := Image{ID: id, Caption: caption, Location: location}
	store := DBImageStore{DB: db}
	err = store.SaveImage(image)
	if err != nil {
		t.Fatal(err)
	}
	if err := db_mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestSaveImageNilLocationCaption(t *testing.T) {
	db, db_mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	id := "20040901_001.jpg"
	db_mock.ExpectExec(InsertImageRegex).WithArgs(id, nil, nil).WillReturnResult(sqlmock.NewResult(1, 1))

	image := Image{ID: id}
	store := DBImageStore{DB: db}
	err = store.SaveImage(image)
	if err != nil {
		t.Fatal(err)
	}
	if err := db_mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestImageTimeFromID(t *testing.T) {
	id := "20040901_001.jpg"
	img := Image{ID: id}
	imgTime, err := img.TimeFromID()
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	expectedFormat := "20060102"
	expectedTime, err := time.Parse(expectedFormat, id[:len(expectedFormat)])
	if err != nil {
		panic(err)
	}
	if *imgTime != expectedTime {
		t.Errorf("Unexpected time returned: %v", *imgTime)
	}
	img.ID = "blerpityblerpityboo"
	_, err = img.TimeFromID()
	if err == nil {
		fmt.Errorf("Expected error when image id = %s", img.ID)
	}
	img.ID = "blah"
	if err == nil {
		fmt.Errorf("Expected error when image id = %s", img.ID)
	}
}
