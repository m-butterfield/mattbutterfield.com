package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddlewareRedirectsToLogin(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/admin/upload", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, r)
	if w.Code != http.StatusFound {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
	if w.Header()["Location"][0] != "/login?next=/admin/upload" {
		t.Errorf("Unexpected redirect location header")
	}
}

func TestAuthMiddlewareAllowsValidAuth(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/admin/upload", nil)
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

func TestAuthMiddlewareRedirectsInvalidAuth(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/admin/upload", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.AddCookie(&http.Cookie{Name: "auth", Value: "12345"})
	authArray = []byte("1234")
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, r)
	if w.Code != http.StatusFound {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
