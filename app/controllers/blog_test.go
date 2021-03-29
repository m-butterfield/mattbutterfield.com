package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBlog(t *testing.T) {
	getRandomImageCalled := 0
	db = &testStore{
		getRandomImage: func() (*data.Image, error) {
			getRandomImageCalled += 1
			return &data.Image{ID: "20040901_001.jpg"}, nil
		},
	}

	r, err := http.NewRequest(http.MethodGet, "/blog", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, r)
	if getRandomImageCalled != 1 {
		t.Errorf("Unexpected call count for GetImage(): %d", getRandomImageCalled)
	}
	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
