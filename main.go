package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbFileName             = "app.db"
	imageBaseURL           = "http://images.mattbutterfield.com/"
	indexFileName          = "index.html"
	port                   = 8000
	selectRandomImageQuery = "SELECT id, caption FROM images WHERE id = (SELECT id FROM images ORDER BY RANDOM() LIMIT 1)"
)

var (
	db            *sql.DB
	indexTemplate *template.Template
)

type Page struct {
	Caption  string
	ImageURL string
}

func serve(res http.ResponseWriter, req *http.Request) {
	err := indexTemplate.Execute(res, getRandomPage())
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	db, err = sql.Open("sqlite3", dbFileName)
	if err != nil {
		panic(err)
	}
	indexTemplate = template.New(indexFileName)
	indexTemplate.ParseFiles(indexFileName)
	http.HandleFunc("/", serve)
	fmt.Printf("Serving on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getRandomPage() Page {
	var imageID, caption string
	row := db.QueryRow(selectRandomImageQuery)
	err := row.Scan(&imageID, &caption)
	if err != nil {
		panic(err)
	}
	return Page{
		Caption:  caption,
		ImageURL: imageBaseURL + imageID,
	}
}
