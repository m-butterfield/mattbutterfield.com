package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLoginSuccess(t *testing.T) {
	form := url.Values{}
	form.Add("auth", "1234")
	r, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	authArray = []byte("1234")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}

func TestLoginFail(t *testing.T) {
	form := url.Values{}
	form.Add("auth", "12345")
	r, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	authArray = []byte("1234")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, r)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
