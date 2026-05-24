package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
)

func video(c *gin.Context) {
	body, err := templateRender("video", makeBasePage(c))
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}
