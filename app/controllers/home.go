package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"net/http"
)

type homePage struct {
	*basePage
	*imageInfo
	NextImagePath string
}

func makeHomePage(c *gin.Context, image *data.Image, nextImageID string) homePage {
	info := getImageInfo(image)
	info.LinkImage = false
	return homePage{
		basePage:      makeBasePage(c),
		imageInfo:     info,
		NextImagePath: makeImagePath(nextImageID),
	}
}

func home(c *gin.Context) {
	id, err := decodeImageID(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "invalid image id")
		return
	}
	var image *data.Image
	if image, err = ds.GetImage(id); err == sql.ErrNoRows {
		c.String(http.StatusNotFound, "not found")
		return
	}
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	nextImage, err := ds.GetRandomImage()
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	body, err := templateRender("index", makeHomePage(c, image, nextImage.ID))
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}
