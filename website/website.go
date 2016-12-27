package website

import (
	"fmt"
	"html/template"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/m-butterfield/mattbutterfield.com/datastore"
	_ "github.com/mattn/go-sqlite3"
)

const (
	imageTemplateName = "website/templates/image.html"
	port              = "8000"
)

func Run() error {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/img/{id}", img)
	fmt.Println("Serving on port: ", port)
	err := http.ListenAndServe(net.JoinHostPort("", port), r)
	if err != nil {
		return err
	}
	return nil
}

func index(w http.ResponseWriter, r *http.Request) {
	image, err := datastore.GetRandomImage()
	if err != nil {
		http.Error(w, "error fetching image", http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/img/"+image.EncodeID(), http.StatusFound)
}

func img(w http.ResponseWriter, r *http.Request) {
	id, err := datastore.DecodeImageID(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid image id", http.StatusInternalServerError)
		return
	}
	image, err := datastore.GetImage(id)
	if err != nil {
		http.Error(w, "error fetching image", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles(imageTemplateName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error fetching template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, image)
}
