package main

import (
	"flag"

	"github.com/m-butterfield/mattbutterfield.com/website"
)

func main() {
	withAdmin := flag.Bool("admin", false, "Run with admin routes")
	flag.Parse()
	err := website.Run(*withAdmin)
	if err != nil {
		panic(err)
	}
}
