package heatmap

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"fmt"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"log"
	"os"
	"path/filepath"
)

func saveGeoJSONResult(coordinates [][][]float64, name string) error {
	geoJSONData, err := makeGeoJSONData(coordinates)
	if err != nil {
		return err
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer func(client *storage.Client) {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}(client)

	object := client.Bucket(lib.FilesBucket).Object(fmt.Sprintf("heatmap-geoJSONData/%s.geojson", name))
	w := object.NewWriter(ctx)
	w.ContentType = "application/json"
	if _, err = w.Write(geoJSONData); err != nil {
		return err
	}
	if err = w.Close(); err != nil {
		return err
	}

	return nil
}

func saveGeoJSONResultLocal(coordinates [][][]float64, fileName string) error {
	geoJSONData, err := makeGeoJSONData(coordinates)
	if err != nil {
		return err
	}

	newpath := filepath.Join(".", "tmp")
	err = os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, geoJSONData, 0644)
}

func makeGeoJSONData(coordinates [][][]float64) ([]byte, error) {
	geoJSON := geoJSONResult{
		Type:        "MultiLineString",
		Coordinates: coordinates,
	}

	return json.Marshal(geoJSON)
}
