package controllers

import (
	"encoding/json"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"log"
	"net/http"
)

type saveImageResult struct {
	URL string `json:"url"`
}

func SaveImage(w http.ResponseWriter, r *http.Request) {
	body := &lib.SaveImageRequest{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		lib.InternalError(err, w)
		return
	}

	if task, err := taskCreator.CreateTask("save_image", "save-image-uploads", body); err != nil {
		lib.InternalError(err, w)
		return
	} else {
		log.Println("Created task: " + task.Name)
	}
	w.WriteHeader(201)
	if err = json.NewEncoder(w).Encode(saveImageResult{URL: "hey"}); err != nil {
		lib.InternalError(err, w)
		return
	}
}
