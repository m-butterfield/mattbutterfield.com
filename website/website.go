package website

import (
	"fmt"
	"html/template"
	"net/http"

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
	indexTemplate.ParseFiles(indexFileName)
	http.HandleFunc("/", serve)
	fmt.Printf("Serving on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return nil
}

func serve(res http.ResponseWriter, req *http.Request) {
	page, err := datastore.GetRandomPage()
	if err != nil {
		panic(err)
	}
	err = indexTemplate.Execute(res, page)
	if err != nil {
		panic(err)
	}
}
