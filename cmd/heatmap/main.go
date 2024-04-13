package main

import (
	"github.com/m-butterfield/mattbutterfield.com/app/heatmap"
	"log"
)

func main() {
	if err := heatmap.BuildHeatMap(); err != nil {
		log.Fatal(err)
	}
}
