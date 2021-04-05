package controllers

import (
	"html/template"
	"net/http"
)

var blogTemplatePath = []string{templatePath + "blog/index.gohtml", baseTemplatePath}

func Blog(w http.ResponseWriter, _ *http.Request) {
	image, err := db.GetRandomImage()
	if err != nil {
		internalError(err, w)
		return
	}
	var tmpl *template.Template
	if tmpl, err = template.ParseFiles(blogTemplatePath...); err != nil {
		internalError(err, w)
		return
	}
	if err = tmpl.Execute(w, makeSingleImagePage(image)); err != nil {
		internalError(err, w)
		return
	}
}
