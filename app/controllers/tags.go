package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
)

func tagImages(c *gin.Context) {
	raw := c.Param("slugs")
	if raw == "" {
		c.String(http.StatusNotFound, "tag not found")
		return
	}
	slugs := strings.Split(raw, ",")

	tags, err := ds.GetTagsBySlugs(slugs)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	tagNames := make([]string, len(tags))
	for i, t := range tags {
		tagNames[i] = t.Name
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

	images, err := ds.GetImagesByTag(slugs, before, 20)
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
		nextURL = fmt.Sprintf("/tag/%s?before=%d#photos", raw, images[len(images)-1].CreatedAt.Unix())
	}

	body, err := templateRender("photos/index", &photosPage{
		basePage:   makeBasePage(),
		ImagesInfo: imagesInfo,
		TagNames:   strings.Join(tagNames, ", "),
		NextURL:    nextURL,
	})
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}
