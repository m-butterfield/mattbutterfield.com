package controllers

import (
	"net/http"
)

func Favicon(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "dont call this yet", http.StatusInternalServerError)
}
