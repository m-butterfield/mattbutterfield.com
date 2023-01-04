package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"strconv"
	"time"
)

type photosPage struct {
	*basePage
	ImagesInfo []*imageInfo
	NextURL    string
}

func photos(c *gin.Context) {
	images, err := getImages(c)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	var imagesInfo []*imageInfo
	for _, image := range images {
		imagesInfo = append(imagesInfo, getImageInfo(image))
	}

	nextURL := ""
	if len(images) > 0 {
		nextURL = fmt.Sprintf("/photos?before=%d#photos", images[len(images)-1].CreatedAt.Unix())
	}

	body, err := templateRender("photos/index", &photosPage{
		basePage:   makeBasePage(),
		ImagesInfo: imagesInfo,
		NextURL:    nextURL,
	})
	c.Render(200, body)
}

func getImages(c *gin.Context) ([]*data.Image, error) {
	var before time.Time
	beforeStr := c.Query("before")
	if beforeStr == "" {
		before = time.Now()
	} else {
		beforeInt, err := strconv.ParseInt(beforeStr, 10, 64)
		if err != nil {
			lib.InternalError(err, c)
		}
		before = time.Unix(beforeInt, 0)
	}

	images, err := ds.GetImages(before, 20)
	if err != nil {
		lib.InternalError(err, c)
		return nil, err
	}
	return images, nil
}
