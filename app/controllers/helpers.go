package controllers

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/base64"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	dateDisplayLayout = "January 2006"
	homeImage         = "20150615_002.jpg"
	imageBaseURL      = "https://images.mattbutterfield.com/"
	templatePath      = "templates/"
)

var (
	baseTemplatePaths = []string{
		templatePath + "image.gohtml",
		templatePath + "base.gohtml",
	}

	authArray []byte
	db        data.Store
	pubSub    *pubsub.Client
)

func Initialize() error {
	store, err := data.Connect()
	if err != nil {
		return err
	}
	db = store
	pubSub, err = pubsub.NewClient(context.Background(), "mattbutterfield")
	if err != nil {
		return err
	}
	authArray = []byte(os.Getenv("AUTH_TOKEN"))
	if len(authArray) == 0 {
		log.Fatal("No value set for AUTH_TOKEN")
	}
	return nil
}

func internalError(err error, w http.ResponseWriter) {
	log.Println(err)
	http.Error(w, "internal error", http.StatusInternalServerError)
}

type imageInfo struct {
	ImageURL    string
	ImageWidth  int
	ImageHeight int
}

func getImageInfo(image *data.Image) imageInfo {
	return imageInfo{
		ImageURL:    imageBaseURL + image.ID,
		ImageWidth:  image.Width,
		ImageHeight: image.Height,
	}
}

type singleImagePage struct {
	imageInfo
	Year string
}

func makeSingleImagePage(image *data.Image) singleImagePage {
	return singleImagePage{imageInfo: getImageInfo(image), Year: time.Now().Format("2006")}
}

func makeImagePath(imageID string) string {
	return "/img/" + encodeImageID(imageID)
}

func decodeImageID(encodedID string) (string, error) {
	imageID, err := base64.URLEncoding.DecodeString(encodedID)
	if err != nil {
		return "", err
	}
	return string(imageID), nil
}

func encodeImageID(imageID string) string {
	return base64.URLEncoding.EncodeToString([]byte(imageID))
}
