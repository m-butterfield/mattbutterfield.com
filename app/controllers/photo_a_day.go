package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
)

func photoADay(c *gin.Context) {
	body, err := templateRender("photos/photo_a_day/index", makeBasePage())
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}
