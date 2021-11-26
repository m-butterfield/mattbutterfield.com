package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"html/template"
	"io/fs"
	"net/http"
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
	ffs := &static.FlexFS{}
	if list, err := fs.Glob(ffs, entryPath); err != nil {
		internalError(err, w)
		return
	} else if len(list) == 0 {
		http.NotFound(w, r)
		return
	}
	var tmpl *template.Template
	if tmpl, err = template.ParseFS(ffs, append([]string{entryPath}, baseTemplatePaths...)...); err != nil {
		internalError(err, w)
		return
	}
	if err = tmpl.Execute(w, makeSingleImagePage(image)); err != nil {
		internalError(err, w)
		return
	}
}
