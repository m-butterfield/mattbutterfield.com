package controllers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
