package controllers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFavicon(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/favicon.ico", nil)
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 301, w.Code)
}
