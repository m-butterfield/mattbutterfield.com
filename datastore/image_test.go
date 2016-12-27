package datastore

import (
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetImage(t *testing.T) {
	db, db_mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	id, caption := "1234", "hello"
	db_mock.ExpectQuery("^SELECT id, caption FROM images WHERE id = \\?$").WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption"}).AddRow(id, caption))

	image, err := GetImage(db, id)
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
}

func TestGetLatestImage(t *testing.T) {
	db, db_mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	db_mock.ExpectQuery("^SELECT id, caption FROM images ORDER BY id DESC LIMIT 1$").
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption"}).AddRow("1234", ""))

	_, err = GetLatestImage(db)
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

	db_mock.ExpectQuery("^SELECT id, caption FROM images WHERE id = \\(SELECT id FROM images ORDER BY RANDOM\\(\\) LIMIT 1\\)$").
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption"}).AddRow("1234", ""))

	_, err = GetRandomImage(db)
	if err != nil {
		t.Fatal(err)
	}
	if err := db_mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
