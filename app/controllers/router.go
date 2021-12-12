package controllers

import (
	"github.com/gorilla/mux"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"net/http"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Use(AuthMiddleware)
	r.HandleFunc("/{login:login\\/?}", Login).Methods(http.MethodGet, http.MethodPost)

	ffs := &static.FlexFS{}
	r.PathPrefix("/css/").Handler(http.FileServer(http.FS(ffs)))
	r.PathPrefix("/js/").Handler(http.FileServer(http.FS(ffs)))

	r.HandleFunc("/favicon.ico", Favicon).Methods(http.MethodGet)
	r.HandleFunc("/", Index).Methods(http.MethodGet)
	r.HandleFunc("/img/{id:.*\\/?}", Home).Methods(http.MethodGet)

	r.HandleFunc("/{blog:blog\\/?}", Blog).Methods(http.MethodGet)
	r.HandleFunc("/blog/{entryName:.*\\/?}", BlogEntry).Methods(http.MethodGet)

	r.HandleFunc("/{music:music\\/?}", Music).Methods(http.MethodGet)

	r.HandleFunc("/{photos:photos\\/?}", Photos).Methods(http.MethodGet)

	r.HandleFunc("/{video:video\\/?}", Video).Methods(http.MethodGet)
	r.HandleFunc("/video/{connections:connections\\/?}", VideoConnections).Methods(http.MethodGet)

	r.HandleFunc("/{admin:admin\\/?}", Admin).Methods(http.MethodGet)
	r.HandleFunc("/admin/{upload_music:upload_music\\/?}", UploadMusic).Methods(http.MethodGet)
	r.HandleFunc("/admin/{upload_image:upload_image\\/?}", UploadImage).Methods(http.MethodGet)
	r.HandleFunc("/admin/{signed_upload_url:signed_upload_url\\/?}", SignedUploadURL).Methods(http.MethodPost)
	r.HandleFunc("/admin/{save_song:save_song\\/?}", SaveSong).Methods(http.MethodPost)
	r.HandleFunc("/admin/{song_image:save_image\\/?}", SaveImage).Methods(http.MethodPost)
	return r
}
