package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
)

func blog(c *gin.Context) {
	image, err := ds.GetRandomImage()
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	body, err := templateRender("blog/index", makeSingleImagePage(image))
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}
