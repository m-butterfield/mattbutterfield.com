package tasks

import (
	"encoding/json"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"net/http"
)

func SaveSong(w http.ResponseWriter, r *http.Request) {
	body := &lib.SaveSongRequest{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
