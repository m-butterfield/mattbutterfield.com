package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"io/fs"
	"net/http"
	"strconv"
	"time"
)

func photoADayYear(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	yearPath := fmt.Sprintf("photos/photo_a_day/%d", year)
	ffs := &static.FS{}
	if list, err := fs.Glob(ffs, templatePath+yearPath+".gohtml"); err != nil {
		lib.InternalError(err, c)
		return
	} else if len(list) == 0 {
		c.String(http.StatusNotFound, "not found")
		return
	}

	images, err := getYearImages(c, year)
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

	body, err := templateRender(yearPath, &photosPage{
		basePage:   makeBasePage(),
		ImagesInfo: imagesInfo,
		NextURL:    nextURL,
	})
	c.Render(200, body)
}

func getYearImages(c *gin.Context, year int) ([]*data.Image, error) {
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

	images, err := ds.GetYearImages(year, before, 5)
	if err != nil {
		lib.InternalError(err, c)
		return nil, err
	}
	return images, nil
}
