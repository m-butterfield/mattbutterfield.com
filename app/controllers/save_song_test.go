package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveSong(t *testing.T) {
	body, err := json.Marshal(&saveSongRequest{
		FileName:    "test.wav?123456",
		SongName:    "test song!",
		Description: "test description",
	})
	r, err := http.NewRequest(http.MethodPost, "/admin/save_song", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	r.AddCookie(&http.Cookie{Name: "auth", Value: "1234"})
	authArray = []byte("1234")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
