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
	defer emptyImagesTable()
	id, caption := "1234", "hello"
	if err := NewImage(id, caption).SaveToDB(test_db); err != nil {
		t.Fatal(err)
	}
	if image, err := GetImage(test_db, id); err != nil {
		t.Fatal(err)
	} else {
		if image.ID != id {
			t.Errorf("Unexpected image id: %s != %s", id, image.ID)
		}
		if image.Caption != caption {
			t.Errorf("Unexpected image caption: %s != %s", caption, image.Caption)
		}
	}
}

func TestGetLatestImage(t *testing.T) {
	defer emptyImagesTable()
	if err := NewImage("1234", "").SaveToDB(test_db); err != nil {
		t.Fatal(err)
	}
	secondImage := NewImage("5678", "")
	if err := secondImage.SaveToDB(test_db); err != nil {
		t.Fatal(err)
	}
	if image, err := GetLatestImage(test_db); err != nil {
		t.Fatal(err)
	} else {
		if image.ID != secondImage.ID {
			t.Errorf("Unexpected image id: %s != %s", secondImage.ID, image.ID)
		}
		if image.Caption != secondImage.Caption {
			t.Errorf("Unexpected image caption: %s != %s", secondImage.Caption, image.Caption)
		}
	}
}

func TestGetRandomImage(t *testing.T) {
	defer emptyImagesTable()
	if err := NewImage("1234", "").SaveToDB(test_db); err != nil {
		t.Fatal(err)
	}
	if err := NewImage("5678", "").SaveToDB(test_db); err != nil {
		t.Fatal(err)
	}
	if image, err := GetLatestImage(test_db); err != nil || image == nil {
		t.Fatal(err)
	}
}

func emptyImagesTable() {
	_, err := test_db.Exec(emptyImagesTableQuery)
	if err != nil {
		panic(err)
	}
}
