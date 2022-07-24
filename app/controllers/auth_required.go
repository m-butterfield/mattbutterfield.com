package controllers

import (
	"crypto/subtle"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"net/http"
)

func authRequired(c *gin.Context) {
	authValue, err := c.Cookie("auth")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			redirectToLogin(c)
			return
		}
		lib.InternalError(err, c)
		return
	}
	if subtle.ConstantTimeCompare([]byte(authValue), authArray) != 1 {
		redirectToLogin(c)
		return
	}
	c.Set("loggedIn", true)
}

func redirectToLogin(c *gin.Context) {
	c.Redirect(http.StatusFound, "/login?next="+c.FullPath())
	c.Abort()
}
