package tasks

import (
	"github.com/m-butterfield/mattbutterfield.com/app/data"
)

var (
	db data.Store
)

func Initialize() error {
	store, err := data.Connect()
	if err != nil {
		return err
	}
	db = store
	return nil
}
