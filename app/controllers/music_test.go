package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMusic(t *testing.T) {
	getSongsCalled := 0
	db = &testStore{
		getSongs: func() ([]*data.Song, error) {
			getSongsCalled += 1
			createdAt := time.Date(2021, time.Month(9), 6, 13, 11, 0, 0, time.UTC)
			return []*data.Song{{
				ID:          " 20901202",
				Description: "drone",
				CreatedAt:   &createdAt,
			}}, nil
		},
	}
	r, err := http.NewRequest(http.MethodGet, "/music", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, r)
	if getSongsCalled != 1 {
		t.Errorf("Unexpected call count for GetSongs(): %d", getSongsCalled)
	}
	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
