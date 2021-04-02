package main

import (
	"github.com/m-butterfield/mattbutterfield.com/app"
	"net/http"
)

func main() {
	if err := app.Run(http.ListenAndServe, "8000"); err != nil {
		panic(err)
	}
}
