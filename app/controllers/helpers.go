package controllers

import (
	"cloud.google.com/go/pubsub"
	"context"
	"embed"
	"encoding/base64"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"log"
	"net/http"
	"time"
)

const (
	dateDisplayLayout = "January 2006"
	homeImage         = "20150615_002.jpg"
	imageBaseURL      = "https://images.mattbutterfield.com/"
)

var (
	//go:embed templates
	templatesFS embed.FS
	//go:embed css
	cssFS embed.FS
	//go:embed js
	jsFS embed.FS

	templatePath      = "templates/"
	baseTemplatePaths = []string{
		templatePath + "image.gohtml",
		templatePath + "base.gohtml",
	}

	db     data.Store
	pubSub *pubsub.Client
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
