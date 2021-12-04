package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, makeImagePath(lib.HomeImage), http.StatusFound)
}
