package controllers

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"html/template"
	"net/http"
	"strings"
	"time"
)

var homeTemplatePath = append([]string{templatePath + "index.gohtml"}, baseTemplatePaths...)

type homePage struct {
	imageInfo
	ImageCaption  string
	ImageDate     string
	ImageLocation string
	NextImagePath string
	Year          string
}

func makeHomePage(image *data.Image, nextImageID string) homePage {
	return homePage{
		imageInfo: imageInfo{
			ImageURL:    imageBaseURL + image.ID,
			ImageWidth:  image.Width,
			ImageHeight: image.Height,
		},
		ImageCaption:  image.Caption,
		ImageDate:     getImageTimeStr(image),
		ImageLocation: image.Location,
		NextImagePath: makeImagePath(nextImageID),
		Year:          time.Now().Format("2006"),
	}
}

func getImageTimeStr(image *data.Image) string {
	t, err := image.TimeFromID()
	if err != nil {
		return ""
	}
	return t.Format(dateDisplayLayout)
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
		internalError(err, w)
		return
	}
	nextImage, err := db.GetRandomImage()
	if err != nil {
		internalError(err, w)
		return
	}
	tmpl, err := template.ParseFS(templatesFS, homeTemplatePath...)
	if err != nil {
		internalError(err, w)
		return
	}
	imagePage := makeHomePage(image, nextImage.ID)
	tmpl.Execute(w, imagePage)
}
