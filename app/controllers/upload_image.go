package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
)

type uploadPage struct {
	*basePage
	ImageTypes []data.ImageTypeName
	Tags       []*data.Tag
}

func uploadImage(c *gin.Context) {
	tags, err := ds.GetAllTags()
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	body, err := templateRender("admin/upload_image", uploadPage{
		basePage: makeBasePage(c),
		ImageTypes: []data.ImageTypeName{
			data.PhotoADayImageType,
		},
		Tags: tags,
	})
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}
