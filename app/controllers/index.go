package controllers

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, makeImagePath(homeImage), http.StatusFound)
}
