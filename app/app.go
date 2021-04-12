package app

import (
	"github.com/m-butterfield/mattbutterfield.com/app/controllers"
	"log"
	"net"
	"net/http"
)

func Run(listenAndServe func(string, http.Handler) error, port string) error {
	err := controllers.Initialize()
	if err != nil {
		return err
	}
	log.Println("Listening on port: ", port)
	if err = listenAndServe(net.JoinHostPort("", port), controllers.Router()); err != nil {
		return err
	}
	return nil
}
