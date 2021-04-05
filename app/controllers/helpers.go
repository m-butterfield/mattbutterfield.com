package controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"net/http"
	"path/filepath"
	"runtime"
	"time"
)

const (
	dateDisplayLayout = "January 2006"
	homeImage         = "20150615_002.jpg"
	imageBaseURL      = "https://images.mattbutterfield.com/"
)

var (
	_, b, _, _       = runtime.Caller(0)
	basePath         = filepath.Join(filepath.Dir(b), "../..")
	templatePath     = basePath + "/app/templates/"
	baseTemplatePath = templatePath + "base.gohtml"
)

var db data.Store

func Initialize() error {
	store, err := data.Connect()
	if err != nil {
		return err
	}
	db = store
	return nil
}

func internalError(err error, w http.ResponseWriter) {
	fmt.Println(err)
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
