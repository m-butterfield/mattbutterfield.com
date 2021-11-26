package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"html/template"
	"net/http"
)

var blogTemplatePath = append([]string{templatePath + "blog/index.gohtml"}, baseTemplatePaths...)

func Blog(w http.ResponseWriter, _ *http.Request) {
	image, err := db.GetRandomImage()
	if err != nil {
		internalError(err, w)
		return
	}
	var tmpl *template.Template
	if tmpl, err = template.ParseFS(&static.FlexFS{}, blogTemplatePath...); err != nil {
		internalError(err, w)
		return
	}
	if err = tmpl.Execute(w, makeSingleImagePage(image)); err != nil {
		internalError(err, w)
		return
	}
}
