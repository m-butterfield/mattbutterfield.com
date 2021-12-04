package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"html/template"
	"net/http"
)

var videoTemplatePath = append([]string{templatePath + "video.gohtml"}, baseTemplatePaths...)

func Video(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(&static.FlexFS{}, videoTemplatePath...)
	if err != nil {
		lib.InternalError(err, w)
		return
	}
	if err = tmpl.Execute(w, makeBasePage()); err != nil {
		lib.InternalError(err, w)
		return
	}
}
