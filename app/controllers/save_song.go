package controllers

import (
	"encoding/json"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"log"
	"net/http"
)

func SaveSong(w http.ResponseWriter, r *http.Request) {
	body := &lib.SaveSongRequest{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		internalError(err, w)
		return
	}

	if task, err := taskCreator.CreateTask("save_song", "save-song-uploads", body); err != nil {
		internalError(err, w)
		return
	} else {
		log.Println("Created task: " + task.Name)
	}
	w.WriteHeader(201)
}
