package controllers

import (
	"fmt"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

var photosTemplatePath = append([]string{templatePath + "photos/index.gohtml"}, baseTemplatePaths...)

type photosPage struct {
	*basePage
	ImagesInfo []*imageInfo
	NextURL    string
}

func Photos(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(&static.FlexFS{}, photosTemplatePath...)
	if err != nil {
		lib.InternalError(err, w)
		return
	}

	images, err := getImages(w, r)
	if err != nil {
		lib.InternalError(err, w)
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

	if err = tmpl.Execute(w, &photosPage{
		basePage:   makeBasePage(),
		ImagesInfo: imagesInfo,
		NextURL:    nextURL,
	}); err != nil {
		lib.InternalError(err, w)
		return
	}
}

func getImages(w http.ResponseWriter, r *http.Request) ([]*data.Image, error) {
	var before time.Time
	beforeStr := r.URL.Query().Get("before")
	if beforeStr == "" {
		before = time.Now()
	} else {
		beforeInt, err := strconv.ParseInt(beforeStr, 10, 64)
		if err != nil {
			lib.InternalError(err, w)
		}
		before = time.Unix(beforeInt, 0)
	}

	images, err := db.GetImages(before, 5)
	if err != nil {
		lib.InternalError(err, w)
		return nil, err
	}
	return images, nil
}
