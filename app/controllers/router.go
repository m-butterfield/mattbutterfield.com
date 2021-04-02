package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Index).Methods(http.MethodGet)
	r.HandleFunc("/img/{id:.*\\/?}", Home).Methods(http.MethodGet)
	r.HandleFunc("/{blog:blog\\/?}", Blog).Methods(http.MethodGet)
	r.HandleFunc("/blog/{entryName:.*\\/?}", BlogEntry).Methods(http.MethodGet)
	return r
}
