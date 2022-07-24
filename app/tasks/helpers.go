package tasks

import (
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"net"
)

var (
	ds data.Store
)

func Run(port string) error {
	var err error
	if ds, err = data.Connect(); err != nil {
		return err
	}
	r, err := router()
	if err != nil {
		return err
	}
	return r.Run(net.JoinHostPort("", port))
}
