package controllers

import (
	"net/http"
)

func Favicon(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://images.mattbutterfield.com/favicon.ico", http.StatusMovedPermanently)
}
