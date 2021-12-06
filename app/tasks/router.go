package tasks

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{save_song:save_song\\/?}", SaveSong).Methods(http.MethodPost)
	r.HandleFunc("/{save_image:save_image\\/?}", SaveImage).Methods(http.MethodPost)
	return r
}
