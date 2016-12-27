package datastore

import (
	"database/sql"
	"os"
	"testing"
)

const (
	testDBFileName = "../app_test.db"
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
}
