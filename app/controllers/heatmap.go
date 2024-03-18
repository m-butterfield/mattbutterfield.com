package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
)

func heatmap(c *gin.Context) {
	body, err := templateRender("heatmap", makeBasePage())
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}
