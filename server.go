package main

import (
	"flag"

	"github.com/m-butterfield/mattbutterfield.com/app"
)

func main() {
	withAdmin := flag.Bool("admin", false, "Run with admin routes")
	flag.Parse()
	err := app.Run(*withAdmin)
	if err != nil {
		panic(err)
	}
}
