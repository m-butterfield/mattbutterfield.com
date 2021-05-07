package controllers

import (
	"html/template"
	"net/http"
	"time"
)

var multiVideoTemplatePath = append([]string{templatePath + "multivideo.gohtml"}, baseTemplatePaths...)

func MultiVideo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(ffs, multiVideoTemplatePath...)
	if err != nil {
		internalError(err, w)
		return
	}
	if err = tmpl.Execute(w, struct{ Year string }{Year: time.Now().Format("2006")}); err != nil {
		internalError(err, w)
		return
	}
}
