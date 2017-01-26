package data

import (
	"database/sql"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

const (
	baseSelectImageRegex   = "^SELECT id, caption, location FROM images"
	InsertImageRegex       = "^INSERT INTO images \\(id, caption, location\\) VALUES \\(\\?, \\?, \\?\\)$"
	SelectImageByIDRegex   = baseSelectImageRegex + " WHERE id = \\?$"
	SelectLatestImageRegex = baseSelectImageRegex + " ORDER BY id DESC LIMIT 1$"
	SelectRandomImageRegex = baseSelectImageRegex + " WHERE id = \\(SELECT id FROM images ORDER BY RANDOM\\(\\) LIMIT 1\\)$"
	SelectPrevImageRegex   = baseSelectImageRegex + " WHERE id \\< \\? ORDER BY id DESC LIMIT 1$"
	SelectNextImageRegex   = baseSelectImageRegex + " WHERE id \\> \\? ORDER BY id LIMIT 1$"
	UpdateImageRegex       = "^UPDATE images SET location = \\?, caption = \\? WHERE id = \\?$"
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
		t.Errorf("Unfulfilled database expectations: %s", err)
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
		t.Errorf("Unfulfilled database expectations: %s", err)
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
		t.Errorf("Unfulfilled database expectations: %s", err)
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
		t.Errorf("Unfulfilled database expectations: %s", err)
	}
}

func TestUpdateImage(t *testing.T) {
	db, db_mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	id, caption, location := "20040901_001.jpg", "hello", "NYC"
	db_mock.ExpectExec(UpdateImageRegex).WithArgs(location, caption, id).WillReturnResult(sqlmock.NewResult(1, 1))

	store := DBImageStore{DB: db}
	err = store.UpdateImage(id, location, caption)
	if err != nil {
		t.Fatal(err)
	}
	if err := db_mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled database expectations: %s", err)
	}
}

func TestUpdateImageNilLocationCaption(t *testing.T) {
	db, db_mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	id := "20040901_001.jpg"
	db_mock.ExpectExec(UpdateImageRegex).WithArgs(nil, nil, id).WillReturnResult(sqlmock.NewResult(1, 1))

	store := DBImageStore{DB: db}
	err = store.UpdateImage(id, "", "")
	if err != nil {
		t.Fatal(err)
	}
	if err := db_mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled database expectations: %s", err)
	}
}

func TestImageTimeFromID(t *testing.T) {
	img := Image{ID: "20040901_001.jpg"}
	imgTime, err := img.TimeFromID()
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	expectedFormat := "20060102"
	expectedTime, err := time.Parse(expectedFormat, img.ID[:len(expectedFormat)])
	if err != nil {
		panic(err)
	}
	if *imgTime != expectedTime {
		t.Errorf("Unexpected time returned: %v", *imgTime)
	}
	img.ID = "blerpityblerpityboo"
	_, err = img.TimeFromID()
	if err == nil {
		t.Errorf("Expected error when image id = %s", img.ID)
	}
	img.ID = "blah"
	if err == nil {
		t.Errorf("Expected error when image id = %s", img.ID)
	}
}

func TestGetNextPrevious(t *testing.T) {
	db, db_mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	prevID, id, nextID := "20040901_001.jpg", "20040901_002.jpg", "20040901_003.jpg"

	db_mock.ExpectQuery(SelectPrevImageRegex).WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption", "location"}).AddRow(prevID, nil, nil))
	db_mock.ExpectQuery(SelectNextImageRegex).WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption", "location"}).AddRow(nextID, nil, nil))

	store := DBImageStore{db}
	prev, next, err := store.GetPrevNextImages(id)
	if err != nil {
		t.Fatal(err)
	}
	if err := db_mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled database expectations: %s", err)
	}
	if prevID != prev.ID {
		t.Errorf("Unexpected id for 'previous': %s != %s", prevID, prev.ID)
	}
	if nextID != next.ID {
		t.Errorf("Unexpected id for 'next': %s != %s", nextID, next.ID)
	}
}

func TestGetNextPreviousNoRowsPrev(t *testing.T) {
	db, db_mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	id, nextID := "20040901_002.jpg", "20040901_003.jpg"

	db_mock.ExpectQuery(SelectPrevImageRegex).WithArgs(id).
		WillReturnError(sql.ErrNoRows)
	db_mock.ExpectQuery(SelectNextImageRegex).WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption", "location"}).AddRow(nextID, nil, nil))

	store := DBImageStore{db}
	prev, next, err := store.GetPrevNextImages(id)
	if err != nil {
		t.Fatal(err)
	}
	if err := db_mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled database expectations: %s", err)
	}
	if prev != nil {
		t.Errorf("Unexpected return for 'previous': %s", prev.ID)
	}
	if nextID != next.ID {
		t.Errorf("Unexpected id for 'next': %s != %s", nextID, next.ID)
	}
}

func TestGetNextPreviousNoRowsNext(t *testing.T) {
	db, db_mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	prevID, id := "20040901_001.jpg", "20040901_002.jpg"

	db_mock.ExpectQuery(SelectPrevImageRegex).WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption", "location"}).AddRow(prevID, nil, nil))
	db_mock.ExpectQuery(SelectNextImageRegex).WithArgs(id).
		WillReturnError(sql.ErrNoRows)

	store := DBImageStore{db}
	prev, next, err := store.GetPrevNextImages(id)
	if err != nil {
		t.Fatal(err)
	}
	if err := db_mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled database expectations: %s", err)
	}
	if prevID != prev.ID {
		t.Errorf("Unexpected id for 'prev': %s != %s", prevID, prev.ID)
	}
	if next != nil {
		t.Errorf("Unexpected return for 'next': %s", next.ID)
	}
}
