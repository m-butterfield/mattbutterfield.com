package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	testRouter().ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
