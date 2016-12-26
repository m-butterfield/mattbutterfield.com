package website

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/m-butterfield/mattbutterfield.com/datastore"
	_ "github.com/mattn/go-sqlite3"
)

const (
	indexFileName = "index.html"
	port          = 8000
)

var (
	indexTemplate *template.Template
)

func Run() error {
	indexTemplate = template.New(indexFileName)
	_, err := indexTemplate.ParseFiles(indexFileName)
	if err != nil {
		return err
	}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	fmt.Println("Serving on port: ", port)
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		return err
	}
	return nil
}

func home(w http.ResponseWriter, r *http.Request) {
	page, err := datastore.GetRandomPage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = indexTemplate.Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
