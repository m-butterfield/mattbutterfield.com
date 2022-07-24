package controllers

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin/render"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"html/template"
	"net"
	"os"
	"time"
)

const (
	dateDisplayLayout = "January 2006"
	templatePath      = "templates/"
)

var (
	baseTemplatePaths = []string{
		templatePath + "image.gohtml",
		templatePath + "base.gohtml",
	}

	authArray []byte
	ds        data.Store
	pubSub    *pubsub.Client
	tc        lib.TaskCreator
)

func Run(port string) error {
	var err error
	if ds, err = data.Connect(); err != nil {
		return err
	}
	if tc, err = lib.NewTaskCreator(); err != nil {
		return err
	}
	pubSub, err = pubsub.NewClient(context.Background(), lib.ProjectID)
	if err != nil {
		return err
	}
	authArray = []byte(os.Getenv("AUTH_TOKEN"))
	if len(authArray) == 0 {
		return errors.New("no value set for AUTH_TOKEN")
	}
	r, err := router()
	if err != nil {
		return err
	}
	return r.Run(net.JoinHostPort("", port))
}

type imageInfo struct {
	ImageURL      string
	ImageWidth    int
	ImageHeight   int
	ImageCaption  string
	ImageDate     string
	ImageLocation string
}

func getImageInfo(image *data.Image) *imageInfo {
	return &imageInfo{
		ImageURL:      lib.ImagesBaseURL + image.ID,
		ImageWidth:    image.Width,
		ImageHeight:   image.Height,
		ImageCaption:  image.Caption,
		ImageDate:     image.CreatedAt.Format(dateDisplayLayout),
		ImageLocation: image.Location,
	}
}

type basePage struct {
	ImagesBaseURL string
	Year          string
}

type singleImagePage struct {
	*basePage
	*imageInfo
}

func makeSingleImagePage(image *data.Image) *singleImagePage {
	return &singleImagePage{
		basePage:  makeBasePage(),
		imageInfo: getImageInfo(image),
	}
}

func makeBasePage() *basePage {
	return &basePage{
		ImagesBaseURL: lib.ImagesBaseURL,
		Year:          time.Now().Format("2006"),
	}
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

func templateRender(name string, data interface{}) (render.Render, error) {
	paths := append([]string{templatePath + name + ".gohtml"}, baseTemplatePaths...)
	tmpl, err := template.ParseFS(static.FS{}, paths...)
	if err != nil {
		return nil, err
	}
	return render.HTML{
		Template: tmpl,
		Data:     data,
	}, nil
}
