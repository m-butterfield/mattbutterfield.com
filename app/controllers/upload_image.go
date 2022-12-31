package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
)

type uploadPage struct {
	*basePage
	ImageTypes []data.ImageTypeName
}

func uploadImage(c *gin.Context) {
	body, err := templateRender("admin/upload_image", uploadPage{
		basePage: makeBasePage(),
		ImageTypes: []data.ImageTypeName{
			data.PhotoADayImageType,
		},
	})
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}
