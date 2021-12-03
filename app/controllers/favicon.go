package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"net/http"
)

func Favicon(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, lib.ImagesBaseURL+"/favicon.ico", http.StatusMovedPermanently)
}
