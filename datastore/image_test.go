package datastore

import (
	"database/sql"
	"os"
	"testing"
)

const (
	emptyImagesTableQuery = "DELETE FROM images"
	testDBFileName        = "../app_test.db"
)

var (
	test_db *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	test_db, err = InitDB(testDBFileName)
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestGetImage(t *testing.T) {
	id, caption := "1234", "hello"
	image := NewImage(id, caption)
	err := image.SaveToDB(test_db)
	if err != nil {
		t.Fatal(err)
	}
	image, err = GetImage(test_db, id)
	if err != nil {
		t.Fatal(err)
	}
	if image.ID != id {
		t.Errorf("Unexpected image id: %s != %s", id, image.ID)
	}
	if image.Caption != caption {
		t.Errorf("Unexpected image caption: %s != %s", caption, image.Caption)
	}
	emptyImagesTable()
}

func TestGetLatestImage(t *testing.T) {
	firstImage := NewImage("1234", "")
	err := firstImage.SaveToDB(test_db)
	if err != nil {
		t.Fatal(err)
	}
	secondImage := NewImage("5678", "")
	err = secondImage.SaveToDB(test_db)
	if err != nil {
		t.Fatal(err)
	}
	image, err := GetLatestImage(test_db)
	if err != nil {
		t.Fatal(err)
	}
	if image.ID != secondImage.ID {
		t.Errorf("Unexpected image id: %s != %s", secondImage.ID, image.ID)
	}
	if image.Caption != secondImage.Caption {
		t.Errorf("Unexpected image caption: %s != %s", secondImage.Caption, image.Caption)
	}
	emptyImagesTable()
}

func emptyImagesTable() {
	_, err := test_db.Exec(emptyImagesTableQuery)
	if err != nil {
		panic(err)
	}
}
