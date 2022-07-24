package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBlogEntry(t *testing.T) {
	getRandomImageCalled := 0
	ds = &testStore{
		getRandomImage: func() (*data.Image, error) {
			getRandomImageCalled += 1
			return &data.Image{ID: lib.HomeImage}, nil
		},
	}

	r, err := http.NewRequest(http.MethodGet, "/blog/2021-04-05-migrating-to-gcp", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter().ServeHTTP(w, r)
	if getRandomImageCalled != 1 {
		t.Errorf("Unexpected call count for GetRandomImage(): %d", getRandomImageCalled)
	}
	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
