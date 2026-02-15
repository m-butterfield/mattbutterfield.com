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
	Years      []*yearInfo
	Filter     string
}

type yearInfo struct {
	Year   int
	Months []*monthInfo
}

type monthInfo struct {
	Month     int
	MonthName string
	Count     int
}

func photos(c *gin.Context) {
	images, filter, err := getImages(c)
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
		nextURL = fmt.Sprintf("/photos?before=%d", images[len(images)-1].CreatedAt.Unix())
		if filter != "" {
			nextURL += fmt.Sprintf("&filter=%s", filter)
		}
		nextURL += "#photos"
	}

	ymc, err := ds.GetImageYearsMonths()
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	var years []*yearInfo
	var currentYear *yearInfo
	for _, y := range ymc {
		if currentYear == nil || currentYear.Year != y.Year {
			currentYear = &yearInfo{
				Year: y.Year,
			}
			years = append(years, currentYear)
		}
		currentYear.Months = append(currentYear.Months, &monthInfo{
			Month:     int(y.Month),
			MonthName: y.Month.String()[:3],
			Count:     y.Count,
		})
	}

	body, err := templateRender("photos/index", &photosPage{
		basePage:   makeBasePage(),
		ImagesInfo: imagesInfo,
		NextURL:    nextURL,
		Years:      years,
		Filter:     filter,
	})
	c.Render(200, body)
}

func getImages(c *gin.Context) ([]*data.Image, string, error) {
	var before time.Time
	beforeStr := c.Query("before")
	if beforeStr != "" {
		beforeInt, err := strconv.ParseInt(beforeStr, 10, 64)
		if err != nil {
			lib.InternalError(err, c)
		}
		before = time.Unix(beforeInt, 0)
	} else {
		yearStr := c.Query("year")
		monthStr := c.Query("month")
		if yearStr != "" {
			year, err := strconv.Atoi(yearStr)
			if err != nil {
				return nil, "", err
			}
			if monthStr != "" {
				month, err := strconv.Atoi(monthStr)
				if err != nil {
					return nil, "", err
				}
				before = time.Date(year, time.Month(month)+1, 1, 0, 0, 0, 0, time.UTC)
			} else {
				before = time.Date(year+1, 1, 1, 0, 0, 0, 0, time.UTC)
			}
		} else {
			before = time.Now()
		}
	}

	filter := c.Query("filter")
	images, err := ds.GetImages(before, 20, filter)
	if err != nil {
		lib.InternalError(err, c)
		return nil, "", err
	}
	return images, filter, nil
}
