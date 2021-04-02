package controllers

import (
	"html/template"
	"net/http"
)

var blogTemplateName = templatePath + "blog/index.html"

func Blog(w http.ResponseWriter, _ *http.Request) {
	image, err := db.GetRandomImage()
	if err != nil {
		internalError(err, w)
		return
	}
	var tmpl *template.Template
	if tmpl, err = template.ParseFiles(blogTemplateName); err != nil {
		internalError(err, w)
		return
	}
	if err = tmpl.Execute(w, makeSingleImagePage(image)); err != nil {
		internalError(err, w)
		return
	}
}
