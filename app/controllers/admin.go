package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
)

func admin(c *gin.Context) {
	body, err := templateRender("admin/index", makeBasePage())
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}
