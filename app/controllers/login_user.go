package controllers

import (
	"crypto/subtle"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"net/http"
)

const yearInSeconds = 60 * 60 * 24 * 365

func loginUser(c *gin.Context) {
	auth := c.PostFormArray("auth")
	if len(auth) != 1 {
		c.String(http.StatusBadRequest, "auth field required")
		return
	}
	if subtle.ConstantTimeCompare([]byte(auth[0]), authArray) == 1 {
		c.SetCookie("auth", auth[0], yearInSeconds, "/", "mattbutterfield.com", true, true)
		next := c.Query("next")
		if next != "" {
			c.Redirect(http.StatusFound, next)
		} else {
			body, err := templateRender("login", loginPage{
				basePage: makeBasePage(),
				LoggedIn: true,
			})
			if err != nil {
				lib.InternalError(err, c)
				return
			}
			c.Render(200, body)
		}
	} else {
		c.String(http.StatusBadRequest, "invalid auth")
	}
}
