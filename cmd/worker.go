package main

import (
	"github.com/m-butterfield/mattbutterfield.com/app/tasks"
	"log"
	"net"
	"net/http"
)

func main() {
	err := tasks.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	port := "8001"
	log.Println("Listening on port: ", port)
	if err = http.ListenAndServe(net.JoinHostPort("", port), tasks.Router()); err != nil {
		log.Fatal(err)
	}
}
