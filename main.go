package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	Caption  string
	ImageURL string
}

var (
	db *sql.DB
	indexTemplate *template.Template
)

func serve(res http.ResponseWriter, req *http.Request) {
	err := indexTemplate.Execute(res, getRandomPage())
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	db, err = sql.Open("sqlite3", "app.db")
	if err != nil {
		panic(err)
	}
	indexTemplate = template.New("index.html")
	indexTemplate.ParseFiles("index.html")
	http.HandleFunc("/", serve)
	fmt.Println("Serving...")
	http.ListenAndServe(":8000", nil)
}

func getRandomPage() Page {
	var imageID, caption string
	row := db.QueryRow("SELECT id, caption FROM images ORDER BY RANDOM() LIMIT 1")
	err := row.Scan(&imageID, &caption)
	if err != nil {
		panic(err)
	}
	return Page{
		Caption: caption,
		ImageURL: "http://images.mattbutterfield.com/" + imageID,
	}
}
