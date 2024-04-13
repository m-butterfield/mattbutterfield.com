package controllers

import (
	"fmt"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPhotos(t *testing.T) {
	getImagesCalled := 0
	expectedBefore := time.Unix(time.Now().Unix(), 0)
	expectedLimit := 20
	ds = &testStore{
		getImages: func(before time.Time, limit int) ([]*data.Image, error) {
			getImagesCalled += 1
			if before != expectedBefore {
				t.Errorf("Unexpected before: %s != %s", before, expectedBefore)
			}
			if limit != expectedLimit {
				t.Errorf("Unexpected limit: %d != %d", limit, expectedLimit)
			}
			return []*data.Image{{
				ID:       "12345",
				Caption:  "test caption",
				Location: "test location",
				Width:    100,
				Height:   200,
			}}, nil
		},
	}
	r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/photos?before=%d", expectedBefore.Unix()), nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter().ServeHTTP(w, r)
	if getImagesCalled != 1 {
		t.Errorf("Unexpected call count for GetImages(): %d", getImagesCalled)
	}
	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
