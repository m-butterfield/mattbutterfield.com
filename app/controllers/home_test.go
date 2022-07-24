package controllers

import (
	"database/sql"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	imageID := lib.HomeImage
	randImageID := "blerp"
	getImageCalled, randomCalled := 0, 0
	ds = &testStore{
		getImage: func(id string) (*data.Image, error) {
			getImageCalled += 1
			if id != imageID {
				t.Errorf("GetImage called with unexpected image id: %s", id)
			}
			return &data.Image{ID: imageID}, nil
		},
		getRandomImage: func() (*data.Image, error) {
			randomCalled += 1
			return &data.Image{ID: randImageID}, nil
		},
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/img/"+encodeImageID(imageID), nil)
	testRouter().ServeHTTP(w, req)

	if getImageCalled != 1 {
		t.Errorf("Unexpected call count for GetImage(): %d", getImageCalled)
	}
	if randomCalled != 1 {
		t.Errorf("Unexpected call count for GetRandomImage(): %d", randomCalled)
	}
	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}

func TestHomeInvalidID(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "/img/"+"MjAwO", nil)
	w := httptest.NewRecorder()

	testRouter().ServeHTTP(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}

func TestHomeImageNotFound(t *testing.T) {
	getImageCalled := 0
	ds = &testStore{
		getImage: func(id string) (*data.Image, error) {
			getImageCalled += 1
			return nil, sql.ErrNoRows
		},
	}

	r, _ := http.NewRequest(http.MethodGet, "/img/"+encodeImageID("1234"), nil)
	w := httptest.NewRecorder()

	testRouter().ServeHTTP(w, r)
	if getImageCalled != 1 {
		t.Errorf("Unexpected call count for GetImage(): %d", getImageCalled)
	}
	if w.Code != http.StatusNotFound {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
