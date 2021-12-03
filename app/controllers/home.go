package controllers

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"html/template"
	"net/http"
	"strings"
	"time"
)

var homeTemplatePath = append([]string{templatePath + "index.gohtml"}, baseTemplatePaths...)

type homePage struct {
	*basePage
	*imageInfo
	ImageCaption  string
	ImageDate     string
	ImageLocation string
	NextImagePath string
}

func makeHomePage(image *data.Image, nextImageID string) homePage {
	return homePage{
		basePage:      makeBasePage(),
		imageInfo:     getImageInfo(image),
		ImageCaption:  image.Caption,
		ImageDate:     getImageTimeStr(image.Date),
		ImageLocation: image.Location,
		NextImagePath: makeImagePath(nextImageID),
	}
}

func getImageTimeStr(date time.Time) string {
	return date.Format(dateDisplayLayout)
}

func Home(w http.ResponseWriter, r *http.Request) {
	id, err := decodeImageID(strings.TrimSuffix(mux.Vars(r)["id"], "/"))
	if err != nil {
		http.Error(w, "invalid image id", http.StatusBadRequest)
		return
	}
	var image *data.Image
	if image, err = db.GetImage(id); err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		lib.InternalError(err, w)
		return
	}
	nextImage, err := db.GetRandomImage()
	if err != nil {
		lib.InternalError(err, w)
		return
	}
	tmpl, err := template.ParseFS(&static.FlexFS{}, homeTemplatePath...)
	if err != nil {
		lib.InternalError(err, w)
		return
	}
	imagePage := makeHomePage(image, nextImage.ID)
	tmpl.Execute(w, imagePage)
}
