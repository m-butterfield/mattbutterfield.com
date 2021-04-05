package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var blogEntryTemplateBase = templatePath + "blog/%s.gohtml"

func BlogEntry(w http.ResponseWriter, r *http.Request) {
	image, err := db.GetRandomImage()
	if err != nil {
		internalError(err, w)
		return
	}
	entryName := strings.TrimSuffix(mux.Vars(r)["entryName"], "/")
	entryPath := fmt.Sprintf(blogEntryTemplateBase, entryName)
	if _, err := os.Stat(entryPath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}
	var tmpl *template.Template
	if tmpl, err = template.ParseFiles([]string{entryPath, baseTemplatePath}...); err != nil {
		internalError(err, w)
		return
	}
	if err = tmpl.Execute(w, makeSingleImagePage(image)); err != nil {
		internalError(err, w)
		return
	}
}
