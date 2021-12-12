package data

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

const (
	baseSelectImageRegex   = "^SELECT id, caption, location, width, height, created_at FROM images "
	selectImageByIDRegex   = baseSelectImageRegex + "WHERE id = \\$1$"
	selectRandomImageRegex = baseSelectImageRegex + "WHERE id = \\(SELECT id FROM images ORDER BY RANDOM\\(\\) LIMIT 1\\)$"
	saveImageRegex         = "^INSERT INTO images \\(id, caption, location, width, height, created_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6\\)"
)

func TestGetImage(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	store := &dbStore{db: db}

	id, caption, location, width, height, createdAt := "ab23ce7b39649ad4380349578829d5786a9f29bcfca17bc786f2869351fc339b.jpg", "hello", "NYC", 100, 200, time.Now()
	dbMock.ExpectQuery(selectImageByIDRegex).WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption", "location", "width", "height", "created_at"}).AddRow(id, caption, location, width, height, createdAt))

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
	if image.CreatedAt != createdAt {
		t.Errorf("Unexpected created_at: %s ! %s", image.CreatedAt, createdAt)
	}
}

func TestGetRandomImage(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	store := &dbStore{db: db}

	dbMock.ExpectQuery(selectRandomImageRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption", "location", "width", "height", "created_at"}).AddRow("ab23ce7b39649ad4380349578829d5786a9f29bcfca17bc786f2869351fc339b.jpg", nil, nil, 100, 200, time.Now()))

	_, err = store.GetRandomImage()
	if err != nil {
		t.Fatal(err)
	}
	err = dbMock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled database expectations: %s", err)
	}
}

func TestSaveImage(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	store := &dbStore{db: db}
	now := time.Now()
	id := "ab23ce7b39649ad4380349578829d5786a9f29bcfca17bc786f2869351fc339b.jpg"
	caption := "test caption"
	location := "NYC"
	width, height := 100, 200
	dbMock.ExpectPrepare(saveImageRegex).ExpectExec().WithArgs(id, caption, location, width, height, now).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err = store.SaveImage(id, caption, location, width, height, now)
	if err != nil {
		t.Fatal(err)
	}
}
