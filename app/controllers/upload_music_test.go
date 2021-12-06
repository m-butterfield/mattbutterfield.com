package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUploadMusic(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/admin/upload_music", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.AddCookie(&http.Cookie{Name: "auth", Value: "1234"})
	authArray = []byte("1234")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
