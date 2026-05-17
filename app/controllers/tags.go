package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
)

func tagImages(c *gin.Context) {
	slug := c.Param("slug")
	tags, err := ds.GetTags()
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	var tagName string
	for _, t := range tags {
		if t.Slug == slug {
			tagName = t.Name
			break
		}
	}
	if tagName == "" {
		c.String(http.StatusNotFound, "tag not found")
		return
	}

	var before time.Time
	beforeStr := c.Query("before")
	if beforeStr == "" {
		before = time.Now()
	} else {
		beforeInt, err := strconv.ParseInt(beforeStr, 10, 64)
		if err != nil {
			lib.InternalError(err, c)
			return
		}
		before = time.Unix(beforeInt, 0)
	}

	images, err := ds.GetImagesByTag(slug, before, 20)
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
		nextURL = fmt.Sprintf("/tag/%s?before=%d#photos", slug, images[len(images)-1].CreatedAt.Unix())
	}

	body, err := templateRender("photos/index", &photosPage{
		basePage:   makeBasePage(),
		ImagesInfo: imagesInfo,
		TagName:    tagName,
		NextURL:    nextURL,
	})
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}
