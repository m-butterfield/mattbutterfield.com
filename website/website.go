package website

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/m-butterfield/mattbutterfield.com/datastore"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DBFileName    = "app.db"
	imageBaseURL  = "http://images.mattbutterfield.com/"
	imagePathBase = "/img/"
	port          = "8000"
)

var (
	store             datastore.ImageStore
	imageTemplateName = "website/templates/image.html"
)

type ImagePage struct {
	ImageCaption  string
	ImageURL      string
	NextImagePath string
}

func NewImagePage(image, nextImage *datastore.Image) *ImagePage {
	return &ImagePage{
		ImageCaption:  image.Caption,
		ImageURL:      imageBaseURL + image.ID,
		NextImagePath: imagePathBase + encodeImageID(nextImage.ID),
	}
}

func Run() error {
	db, err := datastore.InitDB(DBFileName)
	if err != nil {
		return err
	}
	store = datastore.DBImageStore{DB: db}
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc(imagePathBase+"{id}", img)
	fmt.Println("Serving on port: ", port)
	err = http.ListenAndServe(net.JoinHostPort("", port), r)
	if err != nil {
		return err
	}
	return nil
}

func index(w http.ResponseWriter, r *http.Request) {
	image, err := store.GetRandomImage()
	if err != nil {
		http.Error(w, "error fetching image", http.StatusInternalServerError)
	}
	http.Redirect(w, r, makeImagePath(image.ID), http.StatusFound)
}

func img(w http.ResponseWriter, r *http.Request) {
	id, err := decodeImageID(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid image id", http.StatusInternalServerError)
		return
	}
	image, err := store.GetImage(id)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "error fetching image", http.StatusInternalServerError)
		return
	}
	nextImage, err := store.GetRandomImage()
	if err != nil {
		http.Error(w, "error fetching next image", http.StatusInternalServerError)
	}
	tmpl, err := template.ParseFiles(imageTemplateName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error fetching template", http.StatusInternalServerError)
		return
	}
	imagePage := NewImagePage(image, nextImage)
	tmpl.Execute(w, imagePage)
}

func makeImagePath(imageID string) string {
	return imagePathBase + encodeImageID(imageID)
}

func decodeImageID(encodedID string) (string, error) {
	imageID, err := base64.StdEncoding.DecodeString(encodedID)
	if err != nil {
		return "", err
	}
	return string(imageID), nil
}

func encodeImageID(imageID string) string {
	return base64.StdEncoding.EncodeToString([]byte(imageID))
}
