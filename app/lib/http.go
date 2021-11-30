package lib

import (
	"log"
	"net/http"
)

func InternalError(err error, w http.ResponseWriter) {
	log.Println(err)
	http.Error(w, "internal error", http.StatusInternalServerError)
}
