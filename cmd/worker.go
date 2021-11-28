package main

import (
	"github.com/m-butterfield/mattbutterfield.com/app/tasks"
	"log"
	"net"
	"net/http"
)

func main() {
	port := "8000"
	log.Println("Listening on port: ", port)
	if err := http.ListenAndServe(net.JoinHostPort("", port), tasks.Router()); err != nil {
		log.Fatal(err)
	}
}
