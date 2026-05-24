package controllers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/m-butterfield/mattbutterfield.com/app/data"
)

func TestEditImage(t *testing.T) {
	imageID := "test.jpg"
	ds = &testStore{
		getImage: func(id string) (*data.Image, error) {
			return &data.Image{
				ID:      imageID,
				Caption: "test caption",
				Width:   100,
				Height:  200,
			}, nil
		},
		getAllTags: func() ([]*data.Tag, error) {
			return []*data.Tag{}, nil
		},
	}

	r, err := http.NewRequest(http.MethodGet, "/admin/edit_image/"+encodeImageID(imageID), nil)
	if err != nil {
		t.Fatal(err)
	}
	r.AddCookie(&http.Cookie{Name: "auth", Value: "1234"})
	authArray = []byte("1234")

	w := httptest.NewRecorder()
	testRouter().ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}

func TestEditImageNotFound(t *testing.T) {
	ds = &testStore{
		getImage: func(id string) (*data.Image, error) {
			return nil, sql.ErrNoRows
		},
	}

	r, err := http.NewRequest(http.MethodGet, "/admin/edit_image/"+encodeImageID("nonexistent.jpg"), nil)
	if err != nil {
		t.Fatal(err)
	}
	r.AddCookie(&http.Cookie{Name: "auth", Value: "1234"})
	authArray = []byte("1234")

	w := httptest.NewRecorder()
	testRouter().ServeHTTP(w, r)

	if w.Code != http.StatusNotFound {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
