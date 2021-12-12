package main

import (
	"github.com/m-butterfield/mattbutterfield.com/app/controllers"
	"log"
	"net"
	"net/http"
)

func main() {
	err := controllers.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	port := "8000"
	log.Println("Listening on port: ", port)
	if err = http.ListenAndServe(net.JoinHostPort("", port), controllers.Router()); err != nil {
		log.Fatal(err)
	}
}
