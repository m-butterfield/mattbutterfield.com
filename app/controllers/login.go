package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
)

type loginPage struct {
	*basePage
	LoggedIn bool
}

func login(c *gin.Context) {
	_, loggedIn := c.Get("loggedIn")
	body, err := templateRender("login", loginPage{
		basePage: makeBasePage(),
		LoggedIn: loggedIn,
	})
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}
