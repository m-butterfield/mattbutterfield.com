package controllers

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLoginUser(t *testing.T) {
	w := httptest.NewRecorder()
	form := url.Values{}
	form.Add("auth", "1234")
	authArray = []byte("1234")
	r, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	testRouter().ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
	cookies := w.Result().Cookies()
	assert.Equal(t, len(cookies), 1)
	assert.Equal(t, cookies[0].Value, "1234")
}

func TestLoginUserBadAuth(t *testing.T) {
	w := httptest.NewRecorder()
	form := url.Values{}
	form.Add("auth", "12345")
	authArray = []byte("1234")
	r, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	testRouter().ServeHTTP(w, r)

	assert.Equal(t, 400, w.Code)
	respBody, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "invalid auth", string(respBody))
}
