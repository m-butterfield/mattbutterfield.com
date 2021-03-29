package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, r)
	if w.Code != http.StatusFound {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
