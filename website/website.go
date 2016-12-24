package website

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/m-butterfield/mattbutterfield.com/datastore"
	_ "github.com/mattn/go-sqlite3"
)

const (
	port                   = 8000
	indexFileName          = "index.html"
	imageBaseURL           = "http://images.mattbutterfield.com/"
	selectRandomImageQuery = "SELECT id, caption FROM images WHERE id = (SELECT id FROM images ORDER BY RANDOM() LIMIT 1)"
)

var (
	indexTemplate *template.Template
	db            *sql.DB
)

type Page struct {
	Caption  string
	ImageURL string
}

func Run() error {
	var err error
	db, err = datastore.OpenDB()
	if err != nil {
		return err
	}
	indexTemplate = template.New(indexFileName)
	indexTemplate.ParseFiles(indexFileName)
	http.HandleFunc("/", serve)
	fmt.Printf("Serving on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return nil
}

func serve(res http.ResponseWriter, req *http.Request) {
	err := indexTemplate.Execute(res, getRandomPage())
	if err != nil {
		panic(err)
	}
}

func getRandomPage() Page {
	var (
		imageID string
		caption sql.NullString
	)
	row := db.QueryRow(selectRandomImageQuery)
	err := row.Scan(&imageID, &caption)
	if err != nil {
		panic(err)
	}
	return Page{
		Caption:  caption.String,
		ImageURL: imageBaseURL + imageID,
	}
}
