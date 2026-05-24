package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUploadImage(t *testing.T) {
	ds = &testStore{
		getAllTags: func() ([]*data.Tag, error) {
			return []*data.Tag{}, nil
		},
	}

	r, err := http.NewRequest(http.MethodGet, "/admin/upload_image", nil)
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
