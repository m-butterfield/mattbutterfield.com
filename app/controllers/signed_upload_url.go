package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

type signedUploadURLRequest struct {
	FileName    string `json:"fileName"`
	ContentType string `json:"contentType"`
}

type signedUploadURLResponse struct {
	URL string `json:"url"`
}

func SignedUploadURL(w http.ResponseWriter, r *http.Request) {
	body := &signedUploadURLRequest{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fileName := lib.UploadsPrefix + body.FileName
	contentType := body.ContentType

	conf, err := google.JWTConfigFromJSON(uploaderServiceAccount())
	if err != nil {
		return
	}
	opts := &storage.SignedURLOptions{
		Scheme: storage.SigningSchemeV4,
		Method: "PUT",
		Headers: []string{
			"Content-Type: " + contentType,
		},
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().UTC().Add(15 * time.Minute),
	}

	url, err := storage.SignedURL(lib.FilesBucket, fileName, opts)
	if err != nil {
		if _, err = fmt.Fprintf(w, err.Error()); err != nil {
			log.Fatal("error:", err)
		}
		return
	}
	result, err := json.Marshal(&signedUploadURLResponse{URL: url})
	if err != nil {
		log.Fatal("error: ", err)
	}
	if _, err = fmt.Fprint(w, string(result)); err != nil {
		log.Fatal("error:", err)
	}
}

func uploaderServiceAccount() []byte {
	data, err := base64.StdEncoding.DecodeString(os.Getenv("UPLOADER_SERVICE_ACCOUNT"))
	if err != nil {
		log.Fatal("error:", err)
	}
	return data
}
