package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"html/template"
	"net/http"
)

var uploadTemplatePath = append([]string{templatePath + "upload.gohtml"}, baseTemplatePaths...)

func Upload(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(&static.FlexFS{}, uploadTemplatePath...)
	if err != nil {
		lib.InternalError(err, w)
		return
	}
	if err = tmpl.Execute(w, makeBasePage()); err != nil {
		lib.InternalError(err, w)
		return
	}
}
