package controllers

import (
	"encoding/json"
	"net/http"
)

type saveSongRequest struct {
	FileName    string `json:"fileName"`
	SongName    string `json:"songName"`
	Description string `json:"description"`
}

func SaveSong(w http.ResponseWriter, r *http.Request) {
	body := &saveSongRequest{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(201)
}
