package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Caption  string
	ImageURL string
}

var indexTemplate = template.New("index.html")

func serve(res http.ResponseWriter, req *http.Request) {
	page := Page{
		Caption:  "Hello dude",
		ImageURL: "http://images.mattbutterfield.com/20070909_010.jpg",
	}
	err := indexTemplate.Execute(res, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	indexTemplate.ParseFiles("index.html")
	http.HandleFunc("/", serve)
	fmt.Println("Serving...")
	http.ListenAndServe(":8000", nil)
}
