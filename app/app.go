package app

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"html/template"
	"net"
	"net/http"
	"time"
)

const (
	adminPathBase     = "/admin/"
	dateDisplayLayout = "January 2006"
	dbFileName        = "app.db"
	homeImage         = "20150615_002.jpg"
	imageBaseURL      = "https://images.mattbutterfield.com/"
	imagePathBase     = "/img/"
	port              = "8000"
	templatePath      = "app/templates/"
)

var (
	adminTemplateName = templatePath + "admin.html"
	imageTemplateName = templatePath + "image.html"
)

var dbStore data.DBStore

type imagePage struct {
	ImageCaption  string
	ImageDate     string
	ImageLocation string
	ImageURL      string
	ImageWidth    int
	ImageHeight   int
	NextImagePath string
	Year          string
}

func makeImagePage(image *data.Image, nextImageID string) imagePage {
	return imagePage{
		ImageCaption:  image.Caption,
		ImageDate:     getImageTimeStr(image),
		ImageLocation: image.Location,
		ImageURL:      imageBaseURL + image.ID,
		ImageWidth:    image.Width,
		ImageHeight:   image.Height,
		NextImagePath: makeImagePath(nextImageID),
		Year:          time.Now().Format("2006"),
	}
}

type adminPage struct {
	imagePage
	PreviousURL string
	NextURL     string
}

func makeAdminPage(image *data.Image, prevImageID, nextImageID string) adminPage {
	var prevURL, nextURL string
	if prevImageID != "" {
		prevURL = makeAdminPath(prevImageID)
	}
	if nextImageID != "" {
		nextURL = makeAdminPath(nextImageID)
	}
	return adminPage{
		imagePage:   makeImagePage(image, ""),
		PreviousURL: prevURL,
		NextURL:     nextURL,
	}
}

func getImageTimeStr(image *data.Image) string {
	t, err := image.TimeFromID()
	if err != nil {
		return ""
	}
	return t.Format(dateDisplayLayout)
}

func Run(withAdmin bool) error {
	var err error
	dbStore, err = data.MakeDBStore(dbFileName)
	if err != nil {
		return err
	}
	fmt.Println("Listening on port: ", port)
	err = http.ListenAndServe(net.JoinHostPort("", port), buildRouter(withAdmin))
	if err != nil {
		return err
	}
	return nil
}

func buildRouter(withAdmin bool) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc(imagePathBase+"{id}", img).Methods(http.MethodGet)
	if withAdmin {
		r.HandleFunc(adminPathBase+"{id}", admin).Methods(http.MethodGet, http.MethodPost)
	}
	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, makeImagePath(homeImage), http.StatusFound)
}

func img(w http.ResponseWriter, r *http.Request) {
	id, err := decodeImageID(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid image id", http.StatusBadRequest)
		return
	}
	image, err := dbStore.GetImage(id)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "error fetching image", http.StatusInternalServerError)
		return
	}
	nextImage, err := dbStore.GetRandomImage()
	if err != nil {
		http.Error(w, "error fetching next image", http.StatusInternalServerError)
	}
	tmpl, err := template.ParseFiles(imageTemplateName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error fetching template", http.StatusInternalServerError)
		return
	}
	imagePage := makeImagePage(image, nextImage.ID)
	tmpl.Execute(w, imagePage)
}

func admin(w http.ResponseWriter, r *http.Request) {
	id, err := decodeImageID(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid image id", http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		dbStore.UpdateImage(id, r.PostForm.Get("location"), r.PostForm.Get("caption"))
	}
	image, err := dbStore.GetImage(id)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "error fetching image", http.StatusInternalServerError)
		return
	}
	previous, next, err := dbStore.GetPrevNextImages(image.ID)
	if err != nil {
		http.Error(w, "error fetching previous and next images", http.StatusInternalServerError)
	}
	tmpl, err := template.ParseFiles(adminTemplateName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error fetching template", http.StatusInternalServerError)
		return
	}
	var previousID, nextID string
	if previous != nil {
		previousID = previous.ID
	}
	if next != nil {
		nextID = next.ID
	}
	if r.Method == http.MethodPost && previousID != "" {
		http.Redirect(w, r, makeAdminPath(previousID), http.StatusFound)
	} else {
		adminPage := makeAdminPage(image, previousID, nextID)
		tmpl.Execute(w, adminPage)
	}
}

func makeImagePath(imageID string) string {
	return imagePathBase + encodeImageID(imageID)
}

func makeAdminPath(imageID string) string {
	return adminPathBase + encodeImageID(imageID)
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
