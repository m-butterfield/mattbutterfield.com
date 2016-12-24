package main

import (
	"github.com/m-butterfield/mattbutterfield.com/website"
)

func main() {
	err := website.Run()
	if err != nil {
		panic(err)
	}
}
