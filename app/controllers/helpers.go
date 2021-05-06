package controllers

import (
	"cloud.google.com/go/pubsub"
	"context"
	"embed"
	"encoding/base64"
	"errors"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	dateDisplayLayout = "January 2006"
	homeImage         = "20150615_002.jpg"
	imageBaseURL      = "https://images.mattbutterfield.com/"
)

type flexFS struct{}

var (
	//go:embed templates
	templatesEmbedFS embed.FS
	//go:embed css
	cssEmbedFS embed.FS
	//go:embed js
	jsEmbedFS embed.FS

	ffs = &flexFS{}

	templatePath      = "templates/"
	baseTemplatePaths = []string{
		templatePath + "image.gohtml",
		templatePath + "base.gohtml",
	}

	db     data.Store
	pubSub *pubsub.Client
)

func (f *flexFS) Open(name string) (fs.File, error) {
	if os.Getenv("USE_LOCAL_FS") != "" {
		return os.Open("./app/controllers/" + name)
	}
	if strings.HasPrefix(name, "js/") {
		return jsEmbedFS.Open(name)
	}
	if strings.HasPrefix(name, "css/") {
		return cssEmbedFS.Open(name)
	}
	if strings.HasPrefix(name, "templates/") {
		return templatesEmbedFS.Open(name)
	}
	return nil, errors.New("could not find file")
}

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
