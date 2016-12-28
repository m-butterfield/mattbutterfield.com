package website

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/m-butterfield/mattbutterfield.com/datastore"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestIndex(t *testing.T) {
	var err error
	var db_mock sqlmock.Sqlmock
	db, db_mock, err = sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	imageID := "1234"
	db_mock.ExpectQuery(datastore.SelectRandomImageRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption"}).AddRow(imageID, ""))

	r, err := http.NewRequest(http.MethodGet, imagePathBase+encodeImageID(imageID), nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	index(w, r)
	if w.Code != http.StatusFound {
		t.Errorf("Unexpected return code: %s", w.Code)
	}
	if err := db_mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}


func TestImg(t *testing.T) {
	var err error
	var db_mock sqlmock.Sqlmock
	db, db_mock, err = sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	imageTemplateName = cwd + "/" + "templates/image.html"

	imageID := "1234"
	db_mock.ExpectQuery(datastore.SelectImageByIDRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption"}).AddRow(imageID, ""))
	db_mock.ExpectQuery(datastore.SelectRandomImageRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id", "caption"}).AddRow(imageID, ""))

	r, err := http.NewRequest(http.MethodGet, imagePathBase+encodeImageID(imageID), nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	img(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %s", w.Code)
	}
	if err := db_mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
