package app

import (
	"fmt"
	"github.com/m-butterfield/mattbutterfield.com/app/controllers"
	"net"
	"net/http"
)

func Run(listenAndServe func(string, http.Handler) error, port string) error {
	err := controllers.Initialize()
	if err != nil {
		return err
	}
	fmt.Println("Listening on port: ", port)
	if err = listenAndServe(net.JoinHostPort("", port), controllers.Router()); err != nil {
		return err
	}
	return nil
}
