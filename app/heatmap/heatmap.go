package heatmap

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"fmt"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"io"
	"log"
	"os/exec"
)

func BuildHeatMap() error {
	ds, err := data.Connect()
	if err != nil {
		return err
	}

	//if err := getLatestActivities(ds); err != nil {
	//	return err
	//}

	mbtilesFileName, err := buildHeatmap(ds)
	if err != nil {
		return err
	}

	log.Print("Updating Mapbox")
	err = updateMapbox(ds, mbtilesFileName)
	if err != nil {
		return err
	}

	return nil
}

func buildHeatmap(ds data.Store) (string, error) {
	//coordinates, err := getActivityCoordinates(ds)
	//if err != nil {
	//	return "", err
	//}

	log.Print("Saving heatmap geoJSON")
	geoJSONFileName := "/tmp/heatmap.geojson"
	//if err = saveGeoJSONResultLocal(coordinates, geoJSONFileName); err != nil {
	//	return "", err
	//}

	log.Print("Converting geoJSON to EPSG:4326")
	convertedGeoJSONFileName := "/tmp/heatmap_4326.geojson"
	cmd := exec.Command("ogr2ogr", "-f", "GeoJSON", convertedGeoJSONFileName, "-t_srs", "EPSG:4326", geoJSONFileName)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	log.Print("Converting geoJSON to MBTiles")
	mbtilesFileName := "/tmp/heatmap.mbtiles"
	cmd = exec.Command("tippecanoe", "--force", "-o", mbtilesFileName, convertedGeoJSONFileName)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return mbtilesFileName, nil
}

func getActivityCoordinates(ds data.Store) ([][][]float64, error) {
	activities, err := ds.GetStravaActivities()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer func(client *storage.Client) {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}(client)

	coordinates := make([][][]float64, 0)

	for i, activity := range activities {
		log.Printf("Getting GEOJSON for activity %d", activity.ID)
		object := client.Bucket(lib.FilesBucket).Object(fmt.Sprintf("heatmap-geoJSONData/%d.geojson", activity.ID))
		reader, err := object.NewReader(ctx)
		if err != nil {
			return nil, err
		}

		jsonData, err := io.ReadAll(reader)
		if err != nil {
			return nil, err
		}

		jsonResult := geoJSONResult{}

		if err = json.Unmarshal(jsonData, &jsonResult); err != nil {
			return nil, err
		}

		if err = reader.Close(); err != nil {
			return nil, err
		}

		coordinates = append(coordinates, jsonResult.Coordinates[0])

		if i > 10 {
			break
		}
	}

	return coordinates, nil
}
