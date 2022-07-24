package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCookieSet(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.AddCookie(&http.Cookie{Name: "auth", Value: "1234"})
	authArray = []byte("1234")
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = r

	authRequired(c)

	assert.Equal(t, w.Result().StatusCode, 200)
}

func TestCookieInvalid(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.AddCookie(&http.Cookie{Name: "auth", Value: "12345"})
	authArray = []byte("1234")
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = r

	authRequired(c)

	if w.Code != http.StatusFound {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
