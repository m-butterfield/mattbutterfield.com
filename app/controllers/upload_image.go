package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"html/template"
	"net/http"
)

var uploadImageTemplatePath = append([]string{templatePath + "admin/upload_image.gohtml"}, baseTemplatePaths...)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(&static.FlexFS{}, uploadImageTemplatePath...)
	if err != nil {
		lib.InternalError(err, w)
		return
	}
	if err = tmpl.Execute(w, makeBasePage()); err != nil {
		lib.InternalError(err, w)
		return
	}
}
